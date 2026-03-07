//go:build integration

package mcp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/mcp/common"
)

func TestMCPTools_ToolRegistry(t *testing.T) {
	config, _ := sdk.NewConfig("test-token")
	client := sdk.NewClient(config)
	registry := NewToolRegistry(client)

	// Verify tools are registered
	tools := registry.ListTools()

	if len(tools) < 100 {
		t.Errorf("expected at least 100 tools, got %d", len(tools))
	}

	toolMap := make(map[string]common.Tool)
	for _, tool := range tools {
		toolMap[tool.Name] = tool
	}

	// Check some expected tools exist
	expectedTools := []string{
		"stock_basic.stock_basic",
		"stock_basic.trade_cal",
		"index.index_basic",
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

	t.Logf("✓ Verified %d tools are registered", len(tools))
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

	registry := NewToolRegistry(client)
	ctx := context.Background()

	// Test that errors are properly propagated
	result, err := registry.CallTool(ctx, "stock_basic.stock_basic", map[string]interface{}{})

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

	registry := NewToolRegistry(client)
	ctx := context.Background()

	// Test calling unknown tool
	result, err := registry.CallTool(ctx, "unknown_tool", map[string]interface{}{})

	// Should return error result, not Go error
	if err != nil {
		t.Errorf("unexpected Go error: %v", err)
	}

	if result == nil {
		t.Fatal("expected result for unknown tool")
	}

	if len(result.Content) == 0 {
		t.Error("expected error message in content")
	}

	t.Log("✓ Unknown tool handled correctly")
}
