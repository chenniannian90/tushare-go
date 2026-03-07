//go:build integration

package mcp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestMCPTools_ToolRegistry(t *testing.T) {
	registry := NewToolRegistry()

	// Verify all 19 tools are registered
	tools := registry.GetTools()

	expectedTools := []string{
		"stock_basic",
		"daily",
		"daily_basic",
		"weekly",
		"monthly",
		"pro_bar",
		"trade_cal",
		"income",
		"balancesheet",
		"fina_indicator",
		"moneyflow",
		"dividend",
		"top10_holders",
		"holder_number",
		"limit_list",
		"concept",
		"concept_detail",
		"index_basic",
		"index_daily",
	}

	if len(tools) != len(expectedTools) {
		t.Errorf("expected %d tools, got %d", len(expectedTools), len(tools))
	}

	toolMap := make(map[string]Tool)
	for _, tool := range tools {
		toolMap[tool.Name] = tool
	}

	for _, expectedTool := range expectedTools {
		tool, exists := toolMap[expectedTool]
		if !exists {
			t.Errorf("tool %s not registered", expectedTool)
			continue
		}

		if tool.Description == "" {
			t.Errorf("tool %s has empty description", expectedTool)
		}
	}
}

func TestMCPTools_ErrorHandling(t *testing.T) {
	// Test error handling when API returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 40203,
			"msg": "insufficient privileges"
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	registry := NewToolRegistry()
	ctx := context.Background()

	// Test that errors are properly propagated
	result, err := registry.CallTool(ctx, client, "stock_basic", map[string]interface{}{})

	if err == nil {
		t.Error("expected error from API, got nil")
	}

	if result != nil {
		t.Error("expected nil result when API returns error")
	}
}

func TestMCPTools_UnknownTool(t *testing.T) {
	config, _ := sdk.NewConfig("test-token")
	client := sdk.NewClient(config)

	registry := NewToolRegistry()
	ctx := context.Background()

	// Test calling unknown tool
	result, err := registry.CallTool(ctx, client, "unknown_tool", map[string]interface{}{})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Error("expected result for unknown tool")
	}

	if len(result.Content) == 0 {
		t.Error("expected error message in content")
	}

	// Verify the error message mentions available tools
	if result.Content[0].Text == "" {
		t.Error("expected non-empty error message")
	}
}

func TestMCPTools_APIConnection(t *testing.T) {
	// Test that MCP tools can connect to mock API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request to verify it's properly formatted
		if r.Method != "POST" {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		// Return successful response with complete data for stock_basic
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "symbol", "name", "area", "industry"],
				"items": [{"ts_code": "000001.SZ", "symbol": "000001", "name": "平安银行", "area": "深圳", "industry": "银行"}]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	registry := NewToolRegistry()
	ctx := context.Background()

	// Test a simple tool call
	result, err := registry.CallTool(ctx, client, "stock_basic", map[string]interface{}{
		"ts_code": "000001.SZ",
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if len(result.Content) == 0 {
		t.Error("expected content in result")
	}

	// Verify content contains data
	if result.Content[0].Text == "" {
		t.Error("expected non-empty text content")
	}

	t.Logf("Tool result: %s", result.Content[0].Text)
}

func TestMCPTools_AllToolsReachable(t *testing.T) {
	// Test that all 19 tools are callable (even if they fail due to mock data)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return minimal success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code"],
				"items": [{"ts_code": "000001.SZ"}]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	registry := NewToolRegistry()
	ctx := context.Background()

	tools := registry.GetTools()

	for _, tool := range tools {
		t.Run(tool.Name, func(t *testing.T) {
			// Try to call each tool with minimal params
			result, err := registry.CallTool(ctx, client, tool.Name, map[string]interface{}{
				"ts_code": "000001.SZ",
			})

			// We don't care if it succeeds (mock data may not match),
			// but we should get either a result or an error
			if result == nil && err == nil {
				t.Errorf("tool %s returned nil result and nil error", tool.Name)
			}

			// Log the result for debugging
			if err != nil {
				t.Logf("tool %s returned error (expected with mock data): %v", tool.Name, err)
			} else if result != nil && len(result.Content) > 0 {
				t.Logf("tool %s succeeded with content length: %d", tool.Name, len(result.Content[0].Text))
			}
		})
	}
}
