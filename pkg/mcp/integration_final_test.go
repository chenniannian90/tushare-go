//go:build integration

package mcp

import (
	"context"
	"os"
	"testing"
	"time"

	"tushare-go/pkg/sdk"
)

// TestMCPIntegration_RealAPIThroughMCP tests real API calls through MCP
func TestMCPIntegration_RealAPIThroughMCP(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	registry := NewToolRegistry(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test tools through MCP
	t.Run("list_tools_through_mcp", func(t *testing.T) {
		tools := registry.ListTools()

		if len(tools) < 100 {
			t.Errorf("expected at least 100 tools, got %d", len(tools))
		}

		t.Logf("✓ All %d tools are registered in MCP server", len(tools))

		// Verify each tool has required fields
		for _, tool := range tools {
			if tool.Name == "" {
				t.Errorf("tool has empty name")
			}
			if tool.Description == "" {
				t.Errorf("tool %s has empty description", tool.Name)
			}
		}
	})

	t.Run("call_stock_basic_through_mcp", func(t *testing.T) {
		result, err := registry.CallTool(ctx, "stock_basic.stock_basic", map[string]interface{}{
			"ts_code": "000001.SZ",
		})

		if err != nil {
			t.Fatalf("failed to call tool through MCP: %v", err)
		}

		if result == nil {
			t.Fatal("expected result, got nil")
		}

		if len(result.Content) == 0 {
			t.Error("expected content in result")
		}

		t.Logf("✓ Successfully called stock_basic through MCP")
	})

	t.Run("call_trade_cal_through_mcp", func(t *testing.T) {
		result, err := registry.CallTool(ctx, "stock_market.trade_cal", map[string]interface{}{
			"exchange":  "SSE",
			"start_date": "20240101",
			"end_date":   "20240105",
		})

		if err != nil {
			t.Fatalf("failed to call tool through MCP: %v", err)
		}

		if result == nil {
			t.Fatal("expected result, got nil")
		}

		t.Logf("✓ Successfully called trade_cal through MCP")
	})
}

// TestMCPIntegration_ErrorHandlingThroughMCP tests error handling
func TestMCPIntegration_ErrorHandlingThroughMCP(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	registry := NewToolRegistry(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("invalid_tool_name", func(t *testing.T) {
		result, err := registry.CallTool(ctx, "invalid_tool_name", map[string]interface{}{})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected result for invalid tool")
		}

		if len(result.Content) == 0 {
			t.Error("expected error message in content")
		}

		t.Log("✓ Invalid tool name handled correctly")
	})

	t.Run("invalid_parameters", func(t *testing.T) {
		result, err := registry.CallTool(ctx, "stock_basic.stock_basic", map[string]interface{}{
			"ts_code": "INVALID_FORMAT_12345",
		})

		if err != nil {
			t.Logf("✓ Invalid parameters handled correctly: %v", err)
		} else if result != nil && len(result.Content) > 0 {
			t.Log("✓ Invalid parameters returned result")
		}
	})
}

// TestMCPIntegration_ConcurrentCallsThroughMCP tests concurrent tool calls
func TestMCPIntegration_ConcurrentCallsThroughMCP(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	registry := NewToolRegistry(client)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	results := make(chan error, 3)

	// Concurrent tool calls
	go func() {
		_, err := registry.CallTool(ctx, "stock_basic.stock_basic", map[string]interface{}{
			"ts_code": "000001.SZ",
		})
		results <- err
	}()

	go func() {
		_, err := registry.CallTool(ctx, "stock_market.trade_cal", map[string]interface{}{
			"exchange":   "SSE",
			"start_date": "20240101",
			"end_date":   "20240105",
		})
		results <- err
	}()

	go func() {
		_, err := registry.CallTool(ctx, "index.index_basic", map[string]interface{}{
			"market": "SSE",
		})
		results <- err
	}()

	// Collect results
	successCount := 0
	for i := 0; i < 3; i++ {
		if err := <-results; err != nil {
			t.Errorf("concurrent call %d failed: %v", i+1, err)
		} else {
			successCount++
		}
	}

	t.Logf("Concurrent MCP tool calls: %d out of 3 succeeded", successCount)

	if successCount != 3 {
		t.Errorf("expected all 3 concurrent calls to succeed, got %d", successCount)
	}
}

// TestMCPIntegration_ToolCount verifies the expected number of tools
func TestMCPIntegration_ToolCount(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	registry := NewToolRegistry(client)

	tools := registry.ListTools()

	// We expect at least 190 tools (195 was the count at migration time)
	if len(tools) < 190 {
		t.Errorf("expected at least 190 tools, got %d", len(tools))
	}

	t.Logf("✓ Tool registry contains %d tools", len(tools))

	// Verify no duplicate tool names
	toolNames := make(map[string]bool)
	duplicates := []string{}

	for _, tool := range tools {
		if toolNames[tool.Name] {
			duplicates = append(duplicates, tool.Name)
		}
		toolNames[tool.Name] = true
	}

	if len(duplicates) > 0 {
		t.Errorf("found duplicate tool names: %v", duplicates)
	} else {
		t.Log("✓ No duplicate tool names found")
	}
}
