package server

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"tushare-go/pkg/mcp/common"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// MockToolRegistry is a mock implementation of ToolRegistryInterface for testing
type MockToolRegistry struct {
tools   []common.Tool
callFn  func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error)
}

func (m *MockToolRegistry) ListTools() []common.Tool {
	return m.tools
}

func (m *MockToolRegistry) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	if m.callFn != nil {
		return m.callFn(ctx, toolName, args)
	}
	return common.SimpleResult("mock result"), nil
}

func TestToolAdapter_RegisterTools(t *testing.T) {
	tests := []struct {
		name        string
		tools       []common.Tool
		expectError bool
	}{
		{
			name: "register single tool",
			tools: []common.Tool{
				{Name: "test_tool", Description: "A test tool"},
			},
			expectError: false,
		},
		{
			name: "register multiple tools",
			tools: []common.Tool{
				{Name: "tool1", Description: "Tool 1"},
				{Name: "tool2", Description: "Tool 2"},
				{Name: "tool3", Description: "Tool 3"},
			},
			expectError: false,
		},
		{
			name:        "register empty tools",
			tools:       []common.Tool{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock registry
			registry := &MockToolRegistry{
				tools: tt.tools,
			}

			// Create adapter
			adapter := NewToolAdapter(registry)

			// Create MCP server
			srv := mcpsdk.NewServer(&mcpsdk.Implementation{
				Name:    "test-server",
				Version: "1.0.0",
			}, nil)

			// Register tools
			err := adapter.RegisterTools(srv)

			// Check result
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestToolAdapter_ToolHandler(t *testing.T) {
	tests := []struct {
		name           string
		tool           common.Tool
		args           map[string]interface{}
		mockResult     *common.ToolResult
		mockError      error
		expectError    bool
		expectContent  string
		expectIsError  bool
		specialCase    string
	}{
		{
			name: "successful tool call",
			tool: common.Tool{Name: "test_tool", Description: "Test"},
			args: map[string]interface{}{
				"param1": "value1",
			},
			mockResult: common.SimpleResult("Success"),
			expectError: false,
			expectContent: "Success",
			expectIsError: false,
		},
		{
			name: "tool call with error",
			tool: common.Tool{Name: "error_tool", Description: "Error tool"},
			args: map[string]interface{}{},
			mockResult: nil, // Will return error from callFn
			mockError: fmt.Errorf("test error"),
			expectError: false,
			expectContent: "Error: test error",
			expectIsError: true,
		},
		{
			name: "tool call with nil result",
			tool: common.Tool{Name: "nil_tool", Description: "Nil tool"},
			args: map[string]interface{}{},
			mockResult: nil,
			mockError: nil,
			expectError: false,
			expectContent: "No results returned",
			expectIsError: false,
			specialCase: "returnNil",
		},
		{
			name: "tool call with empty content",
			tool: common.Tool{Name: "empty_tool", Description: "Empty tool"},
			args: map[string]interface{}{},
			mockResult: &common.ToolResult{Content: []common.Content{}},
			expectError: false,
			expectContent: "No results returned",
			expectIsError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock registry
			registry := &MockToolRegistry{
				tools: []common.Tool{tt.tool},
				callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
					if tt.mockError != nil {
						return nil, tt.mockError
					}
					if tt.specialCase == "returnNil" {
						return nil, nil
					}
					if tt.mockResult != nil {
						return tt.mockResult, nil
					}
					// Default: return simple result
					return common.SimpleResult("default result"), nil
				},
			}

			// Create adapter
			adapter := NewToolAdapter(registry)

			// Create MCP server and register tool
			srv := mcpsdk.NewServer(&mcpsdk.Implementation{
				Name:    "test-server",
				Version: "1.0.0",
			}, nil)

			if err := adapter.RegisterTools(srv); err != nil {
				t.Fatalf("Failed to register tools: %v", err)
			}

			// Get the tool handler
			handler := adapter.createToolHandler(tt.tool)

			// Prepare request arguments
			argsJSON, _ := json.Marshal(tt.args)
			req := &mcpsdk.CallToolRequest{
				Params: &mcpsdk.CallToolParamsRaw{
					Arguments: json.RawMessage(argsJSON),
				},
			}

			// Call handler
			result, err := handler(context.Background(), req)

			// Check error expectation
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Check result
			if result != nil {
				if result.IsError != tt.expectIsError {
					t.Errorf("Expected IsError=%v, got=%v", tt.expectIsError, result.IsError)
				}

				if len(result.Content) > 0 {
					textContent, ok := result.Content[0].(*mcpsdk.TextContent)
					if !ok {
						t.Errorf("Expected TextContent, got %T", result.Content[0])
					} else if textContent.Text != tt.expectContent {
						t.Errorf("Expected content=%q, got=%q", tt.expectContent, textContent.Text)
					}
				}
			}
		})
	}
}

func TestToolAdapter_ToolHandler_ErrorCases(t *testing.T) {
	tests := []struct {
		name          string
		tool          common.Tool
		args          map[string]interface{}
		mockError     error
		setupMock     func(*MockToolRegistry)
		expectContent string
	}{
		{
			name: "tool returns error",
			tool: common.Tool{Name: "error_tool", Description: "Error tool"},
			args: map[string]interface{}{},
			mockError: nil, // Will be set in test
			setupMock: func(m *MockToolRegistry) {
				m.callFn = func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
					return nil, nil
				}
			},
			expectContent: "Error: <nil>", // Will be updated
		},
		{
			name: "invalid JSON arguments",
			tool: common.Tool{Name: "json_tool", Description: "JSON tool"},
			args: nil, // Will use invalid JSON in test
			setupMock: func(m *MockToolRegistry) {
				m.callFn = func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
					return common.SimpleResult("called"), nil
				}
			},
			expectContent: "Error: failed to parse arguments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := &MockToolRegistry{
				tools: []common.Tool{tt.tool},
			}

			if tt.setupMock != nil {
				tt.setupMock(registry)
			}

			adapter := NewToolAdapter(registry)

			handler := adapter.createToolHandler(tt.tool)

			var req *mcpsdk.CallToolRequest

			if tt.name == "invalid JSON arguments" {
				// Create request with invalid JSON
				req = &mcpsdk.CallToolRequest{
					Params: &mcpsdk.CallToolParamsRaw{
						Arguments: json.RawMessage([]byte("{invalid json}")),
					},
				}
			} else {
				argsJSON, _ := json.Marshal(tt.args)
				req = &mcpsdk.CallToolRequest{
					Params: &mcpsdk.CallToolParamsRaw{
						Arguments: json.RawMessage(argsJSON),
					},
				}
			}

			result, err := handler(context.Background(), req)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("Expected result but got nil")
			}

			if len(result.Content) == 0 {
				t.Fatal("Expected content but got empty slice")
			}

			textContent, ok := result.Content[0].(*mcpsdk.TextContent)
			if !ok {
				t.Fatalf("Expected TextContent, got %T", result.Content[0])
			}

			// For error tool test
			if tt.name == "tool returns error" {
				registry.callFn = func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
					return nil, tt.mockError
				}
				handler = adapter.createToolHandler(tt.tool)
				argsJSON, _ := json.Marshal(tt.args)
				req = &mcpsdk.CallToolRequest{
					Params: &mcpsdk.CallToolParamsRaw{
						Arguments: json.RawMessage(argsJSON),
					},
				}
				result, err = handler(context.Background(), req)

				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				textContent, ok = result.Content[0].(*mcpsdk.TextContent)
				if !ok {
					t.Fatalf("Expected TextContent, got %T", result.Content[0])
				}
			}

			if textContent.Text == "" {
				t.Errorf("Expected non-empty content")
			}
		})
	}
}

// Test integration with real (but mocked) ToolRegistry
func TestToolAdapter_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This test can be expanded to use a real ToolRegistry with mock SDK client
	// For now, we use our mock

	tools := []common.Tool{
		{Name: "integration_tool", Description: "Integration test tool"},
	}

	registry := &MockToolRegistry{
		tools: tools,
		callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
			return &common.ToolResult{
				Content: []common.Content{
					{Type: "text", Text: "Integration test result"},
				},
			}, nil
		},
	}

	adapter := NewToolAdapter(registry)

	srv := mcpsdk.NewServer(&mcpsdk.Implementation{
		Name:    "integration-test-server",
		Version: "1.0.0",
	}, nil)

	err := adapter.RegisterTools(srv)
	if err != nil {
		t.Fatalf("Failed to register tools: %v", err)
	}

	// Verify tools were registered
	if len(registry.ListTools()) != 1 {
		t.Errorf("Expected 1 tool, got %d", len(registry.ListTools()))
	}

	// Test tool handler
	handler := adapter.createToolHandler(tools[0])
	argsJSON, _ := json.Marshal(map[string]interface{}{"test": "value"})
	req := &mcpsdk.CallToolRequest{
		Params: &mcpsdk.CallToolParamsRaw{
			Arguments: json.RawMessage(argsJSON),
		},
	}

	result, err := handler(context.Background(), req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result == nil || len(result.Content) == 0 {
		t.Fatal("Expected result with content")
	}

	textContent, ok := result.Content[0].(*mcpsdk.TextContent)
	if !ok {
		t.Fatalf("Expected TextContent, got %T", result.Content[0])
	}

	expectedText := "Integration test result"
	if textContent.Text != expectedText {
		t.Errorf("Expected text=%q, got=%q", expectedText, textContent.Text)
	}
}
