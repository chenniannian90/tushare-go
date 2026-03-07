//go:build integration

package mcp

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
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
	registry := NewToolRegistry()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test tools through MCP
	t.Run("list_tools_through_mcp", func(t *testing.T) {
		tools := registry.GetTools()

		if len(tools) != 19 {
			t.Errorf("expected 19 tools, got %d", len(tools))
		}

		t.Logf("✓ All 19 tools are registered in MCP server")

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
		result, err := registry.CallTool(ctx, client, "stock_basic", map[string]interface{}{
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
		result, err := registry.CallTool(ctx, client, "trade_cal", map[string]interface{}{
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
	registry := NewToolRegistry()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("invalid_tool_name", func(t *testing.T) {
		result, err := registry.CallTool(ctx, client, "invalid_tool_name", map[string]interface{}{})

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
		result, err := registry.CallTool(ctx, client, "stock_basic", map[string]interface{}{
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
	registry := NewToolRegistry()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	results := make(chan error, 3)

	// Concurrent tool calls
	go func() {
		_, err := registry.CallTool(ctx, client, "stock_basic", map[string]interface{}{
			"ts_code": "000001.SZ",
		})
		results <- err
	}()

	go func() {
		_, err := registry.CallTool(ctx, client, "trade_cal", map[string]interface{}{
			"exchange":  "SSE",
			"start_date": "20240101",
			"end_date":   "20240105",
		})
		results <- err
	}()

	go func() {
		_, err := registry.CallTool(ctx, client, "index_basic", map[string]interface{}{
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

// TestMCPIntegration_ServerLifecycle tests server lifecycle management
func TestMCPIntegration_ServerLifecycle(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	config, _ := sdk.NewConfig(token)
	client := sdk.NewClient(config)
	server := NewStdioMCPServer(client)

	// Test health status
	t.Run("health_status", func(t *testing.T) {
		health := server.Health.GetStatus(19)

		if health.Status != "healthy" {
			t.Errorf("expected healthy status, got %s", health.Status)
		}

		if health.ToolCount != 19 {
			t.Errorf("expected 19 tools, got %d", health.ToolCount)
		}

		if health.Uptime == "" {
			t.Error("expected uptime string")
		}

		t.Logf("✓ Server health: %s, uptime: %s, tools: %d",
			health.Status, health.Uptime, health.ToolCount)
	})

	// Test lifecycle management
	t.Run("lifecycle_management", func(t *testing.T) {
		lifecycle := server.Lifecycle

		if lifecycle == nil {
			t.Fatal("expected lifecycle manager")
		}

		if lifecycle.IsShutdown() {
			t.Error("server should not be shutdown")
		}

		t.Log("✓ Lifecycle manager is properly initialized")
	})
}
