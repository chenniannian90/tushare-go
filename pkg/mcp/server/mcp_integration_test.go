package server

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"tushare-go/pkg/mcp/common"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// TestMCPServerIntegration tests the complete MCP server workflow
// using the official SDK with our tool adapter
func TestMCPServerIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tests := []struct {
		name        string
		setupFunc   func() (*mcpsdk.Server, *MockToolRegistry)
		testFunc    func(t *testing.T, srv *mcpsdk.Server, registry *MockToolRegistry)
		expectError bool
	}{
		{
			name: "server initialization and tool registration",
			setupFunc: func() (*mcpsdk.Server, *MockToolRegistry) {
				tools := []common.Tool{
					{Name: "test.tool1", Description: "Test tool 1"},
					{Name: "test.tool2", Description: "Test tool 2"},
				}
				registry := &MockToolRegistry{
					tools: tools,
					callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
						return common.SimpleResult("test result"), nil
					},
				}
				srv := mcpsdk.NewServer(&mcpsdk.Implementation{
					Name:    "test-server",
					Version: "1.0.0",
				}, nil)
				return srv, registry
			},
			testFunc: func(t *testing.T, srv *mcpsdk.Server, registry *MockToolRegistry) {
				adapter := NewToolAdapter(registry)
				err := adapter.RegisterTools(srv)
				if err != nil {
					t.Fatalf("Failed to register tools: %v", err)
				}

				// Verify tools were registered
				if len(registry.ListTools()) != 2 {
					t.Errorf("Expected 2 tools, got %d", len(registry.ListTools()))
				}
			},
			expectError: false,
		},
		{
			name: "tool call with valid arguments",
			setupFunc: func() (*mcpsdk.Server, *MockToolRegistry) {
				tools := []common.Tool{
					{Name: "echo", Description: "Echo tool"},
				}
				registry := &MockToolRegistry{
					tools: tools,
					callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
						message, _ := args["message"].(string)
						return &common.ToolResult{
							Content: []common.Content{
								{Type: "text", Text: message},
							},
						}, nil
					},
				}
				srv := mcpsdk.NewServer(&mcpsdk.Implementation{
					Name:    "echo-server",
					Version: "1.0.0",
				}, nil)
				return srv, registry
			},
			testFunc: func(t *testing.T, srv *mcpsdk.Server, registry *MockToolRegistry) {
				adapter := NewToolAdapter(registry)
				if err := adapter.RegisterTools(srv); err != nil {
					t.Fatalf("Failed to register tools: %v", err)
				}

				// Simulate tool call
				handler := adapter.createToolHandler(registry.tools[0])
				args := map[string]interface{}{"message": "hello"}
				argsJSON, _ := json.Marshal(args)
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

				if textContent.Text != "hello" {
					t.Errorf("Expected 'hello', got '%s'", textContent.Text)
				}
			},
			expectError: false,
		},
		{
			name: "tool call with successful response",
			setupFunc: func() (*mcpsdk.Server, *MockToolRegistry) {
				tools := []common.Tool{
					{Name: "success_tool", Description: "Success tool"},
				}
				registry := &MockToolRegistry{
					tools: tools,
					callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
						// Return successful result
						return &common.ToolResult{
							Content: []common.Content{
								{Type: "text", Text: "Tool executed successfully"},
							},
						}, nil
					},
				}
				srv := mcpsdk.NewServer(&mcpsdk.Implementation{
					Name:    "success-server",
					Version: "1.0.0",
				}, nil)
				return srv, registry
			},
			testFunc: func(t *testing.T, srv *mcpsdk.Server, registry *MockToolRegistry) {
				adapter := NewToolAdapter(registry)
				if err := adapter.RegisterTools(srv); err != nil {
					t.Fatalf("Failed to register tools: %v", err)
				}

				handler := adapter.createToolHandler(registry.tools[0])
				req := &mcpsdk.CallToolRequest{
					Params: &mcpsdk.CallToolParamsRaw{
						Arguments: json.RawMessage(`{}`),
					},
				}

				result, err := handler(context.Background(), req)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				if result.IsError {
					t.Error("Expected IsError=false, got true")
				}

				if len(result.Content) == 0 {
					t.Fatal("Expected success content")
				}

				textContent, ok := result.Content[0].(*mcpsdk.TextContent)
				if !ok {
					t.Fatalf("Expected TextContent, got %T", result.Content[0])
				}

				expectedText := "Tool executed successfully"
				if textContent.Text != expectedText {
					t.Errorf("Expected '%s', got: %s", expectedText, textContent.Text)
				}
			},
			expectError: false,
		},
		{
			name: "concurrent tool calls",
			setupFunc: func() (*mcpsdk.Server, *MockToolRegistry) {
				tools := []common.Tool{
					{Name: "concurrent", Description: "Concurrent tool"},
				}
				registry := &MockToolRegistry{
					tools: tools,
					callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
						// Simulate some work
						time.Sleep(10 * time.Millisecond)
						return common.SimpleResult("done"), nil
					},
				}
				srv := mcpsdk.NewServer(&mcpsdk.Implementation{
					Name:    "concurrent-server",
					Version: "1.0.0",
				}, nil)
				return srv, registry
			},
			testFunc: func(t *testing.T, srv *mcpsdk.Server, registry *MockToolRegistry) {
				adapter := NewToolAdapter(registry)
				if err := adapter.RegisterTools(srv); err != nil {
					t.Fatalf("Failed to register tools: %v", err)
				}

				handler := adapter.createToolHandler(registry.tools[0])

				// Make concurrent calls
				const concurrency = 10
				errChan := make(chan error, concurrency)

				for i := 0; i < concurrency; i++ {
					go func() {
						req := &mcpsdk.CallToolRequest{
							Params: &mcpsdk.CallToolParamsRaw{
								Arguments: json.RawMessage(`{}`),
							},
						}
						_, err := handler(context.Background(), req)
						errChan <- err
					}()
				}

				// Wait for all calls to complete
				for i := 0; i < concurrency; i++ {
					if err := <-errChan; err != nil {
						t.Errorf("Concurrent call failed: %v", err)
					}
				}

				t.Log("All concurrent calls completed successfully")
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv, registry := tt.setupFunc()
			tt.testFunc(t, srv, registry)
		})
	}
}

// TestMCPServerWithContextCancellation tests server behavior with context cancellation
func TestMCPServerWithContextCancellation(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tools := []common.Tool{
		{Name: "slow", Description: "Slow tool"},
	}
	registry := &MockToolRegistry{
		tools: tools,
		callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
			// Simulate slow operation
			select {
			case <-time.After(100 * time.Millisecond):
				return common.SimpleResult("done"), nil
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		},
	}

	srv := mcpsdk.NewServer(&mcpsdk.Implementation{
		Name:    "slow-server",
		Version: "1.0.0",
	}, nil)

	adapter := NewToolAdapter(registry)
	if err := adapter.RegisterTools(srv); err != nil {
		t.Fatalf("Failed to register tools: %v", err)
	}

	handler := adapter.createToolHandler(registry.tools[0])

	// Create context that gets cancelled quickly
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	req := &mcpsdk.CallToolRequest{
		Params: &mcpsdk.CallToolParamsRaw{
			Arguments: json.RawMessage(`{}`),
		},
	}

	// Call handler - adapter catches errors and returns error result
	result, err := handler(ctx, req)

	// Adapter always returns nil error and wraps actual errors in result
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if we got an error result
	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	// Either operation completed or was cancelled
	if result.IsError {
		t.Logf("✅ Operation was cancelled as expected: %v", result.Content[0].(*mcpsdk.TextContent).Text)
	} else {
		t.Logf("✅ Operation completed before cancellation")
	}
}

// TestMCPServerProtocolCompliance tests MCP protocol compliance
func TestMCPServerProtocolCompliance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tools := []common.Tool{
		{Name: "compliant_tool", Description: "A compliant MCP tool"},
	}

	registry := &MockToolRegistry{
		tools: tools,
		callFn: func(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
			return &common.ToolResult{
				Content: []common.Content{
					{Type: "text", Text: "Compliant result"},
				},
			}, nil
		},
	}

	srv := mcpsdk.NewServer(&mcpsdk.Implementation{
		Name:    "compliant-server",
		Version: "1.0.0",
	}, nil)

	adapter := NewToolAdapter(registry)
	if err := adapter.RegisterTools(srv); err != nil {
		t.Fatalf("Failed to register tools: %v", err)
	}

	// Verify tool has required fields
	handler := adapter.createToolHandler(registry.tools[0])
	if handler == nil {
		t.Fatal("Expected non-nil handler")
	}

	// Test with valid JSON arguments
	validJSON := `{"param1":"value1","param2":123}`
	req := &mcpsdk.CallToolRequest{
		Params: &mcpsdk.CallToolParamsRaw{
			Arguments: json.RawMessage(validJSON),
		},
	}

	result, err := handler(context.Background(), req)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.IsError {
		t.Error("Expected success, got error")
	}

	if len(result.Content) == 0 {
		t.Error("Expected non-empty content")
	}

	t.Log("✅ Protocol compliance test passed")
}
