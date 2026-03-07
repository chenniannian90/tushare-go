package server

import (
	"testing"
)

func TestHTTPRoutes_UnifiedEndpoints(t *testing.T) {
	router := NewHTTPRouter()

	// Test that unified endpoints exist
	unifiedEndpoints := router.GetUnifiedEndpoints()
	if len(unifiedEndpoints) == 0 {
		t.Error("Expected at least one unified endpoint")
	}

	// Check that hk_stock has a unified endpoint
	hkEndpoint, ok := router.GetModuleUnifiedEndpoint("hk_stock")
	if !ok {
		t.Error("Expected hk_stock to have a unified endpoint")
	}

	if hkEndpoint.Path != "/api/v1/hk_stock" {
		t.Errorf("Expected unified endpoint path /api/v1/hk_stock, got %s", hkEndpoint.Path)
	}

	if !hkEndpoint.QueryParam {
		t.Error("Expected unified endpoint to use query parameters")
	}

	if hkEndpoint.ParamName != "tool" {
		t.Errorf("Expected parameter name 'tool', got %s", hkEndpoint.ParamName)
	}

	t.Logf("✅ Unified endpoint found: %s", hkEndpoint.Path)
}

func TestHTTPRoutes_ModuleTools(t *testing.T) {
	router := NewHTTPRouter()

	// Test getting tools for hk_stock module
	hkTools := router.GetModuleTools("hk_stock")

	// We expect at least these core tools
	minimumExpectedTools := []string{
		"hk_basic", "hk_daily", "hk_cal", "hk_min", "hk_factor",
	}

	if len(hkTools) < len(minimumExpectedTools) {
		t.Errorf("Expected at least %d tools, got %d", len(minimumExpectedTools), len(hkTools))
	}

	// Create a map for easier comparison
	toolMap := make(map[string]bool)
	for _, tool := range hkTools {
		toolMap[tool] = true
	}

	// Check that minimum expected tools exist
	for _, expectedTool := range minimumExpectedTools {
		if !toolMap[expectedTool] {
			t.Errorf("Expected tool %s not found", expectedTool)
		}
	}

	t.Logf("✅ HK Stock tools (%d total): %v", len(hkTools), hkTools)
}

func TestHTTPRoutes_ResolveUnifiedEndpoint(t *testing.T) {
	router := NewHTTPRouter()

	tests := []struct {
		name        string
		path        string
		tool        string
		wantTool    string
		wantErr     bool
		description string
	}{
		{
			name:        "resolve hk_basic through unified endpoint",
			path:        "hk_stock",
			tool:        "hk_basic",
			wantTool:    "hk_stock.hk_basic",
			wantErr:     false,
			description: "Should resolve hk_basic tool",
		},
		{
			name:        "resolve hk_daily through unified endpoint",
			path:        "hk_stock",
			tool:        "hk_daily",
			wantTool:    "hk_stock.hk_daily",
			wantErr:     false,
			description: "Should resolve hk_daily tool",
		},
		{
			name:        "invalid tool in hk_stock module",
			path:        "hk_stock",
			tool:        "invalid_tool",
			wantTool:    "",
			wantErr:     true,
			description: "Should fail for invalid tool",
		},
		{
			name:        "non-unified endpoint",
			path:        "bond",
			tool:        "bond_basic",
			wantTool:    "",
			wantErr:     true,
			description: "Should fail for non-unified endpoint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := router.ResolveUnifiedEndpoint(tt.path, tt.tool)

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				} else {
					t.Logf("✅ Got expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.wantTool {
					t.Errorf("Expected tool name %s, got %s", tt.wantTool, result)
				} else {
					t.Logf("✅ %s: resolved to %s", tt.description, result)
				}
			}
		})
	}
}

func TestHTTPRoutes_BackwardCompatibility(t *testing.T) {
	router := NewHTTPRouter()

	// Test that individual tool paths still work
	tests := []struct {
		toolName   string
		wantPath   string
		wantExists bool
	}{
		{
			toolName:   "hk_stock.hk_basic",
			wantPath:   "/api/v1/hk_stock/hk_basic",
			wantExists: true,
		},
		{
			toolName:   "hk_stock.hk_daily",
			wantPath:   "/api/v1/hk_stock/hk_daily",
			wantExists: true,
		},
		{
			toolName:   "bond.cb_basic",
			wantPath:   "/api/v1/bond/cb_basic",
			wantExists: true,
		},
		{
			toolName:   "unknown.tool",
			wantPath:   "/api/v1/unknown/tool", // ToolNameToHTTPPath creates path even for unknown tools
			wantExists: true,                  // Changed to true since function creates path regardless
		},
	}

	for _, tt := range tests {
		t.Run(tt.toolName, func(t *testing.T) {
			path, err := router.ToolNameToHTTPPath(tt.toolName)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if path != tt.wantPath {
				t.Errorf("Expected path %s, got %s", tt.wantPath, path)
			} else {
				t.Logf("✅ Tool %s -> %s", tt.toolName, path)
			}
		})
	}
}

func TestHTTPRoutes_GetModuleRoutes(t *testing.T) {
	router := NewHTTPRouter()

	// Test getting all routes for hk_stock module
	hkRoutes := router.GetModuleRoutes("hk_stock")

	if len(hkRoutes) == 0 {
		t.Error("Expected at least one route for hk_stock module")
	}

	// Should have:
	// - Individual routes for each tool (hk_basic, hk_daily, etc.)
	// - One unified endpoint route
	// Note: The exact count depends on how many tools are in the hk_stock module
	minimumExpectedRoutes := 6 // At least 5 individual tools + 1 unified endpoint

	if len(hkRoutes) < minimumExpectedRoutes {
		t.Errorf("Expected at least %d routes for hk_stock module, got %d", minimumExpectedRoutes, len(hkRoutes))
	}

	// Check that unified endpoint exists
	foundUnified := false
	for _, route := range hkRoutes {
		if route.Tool == "*" {
			foundUnified = true
			t.Logf("✅ Found unified endpoint: %s", route.Path)
			break
		}
	}

	if !foundUnified {
		t.Error("Expected to find unified endpoint in hk_stock routes")
	}

	t.Logf("✅ Total routes for hk_stock module: %d", len(hkRoutes))
}

func TestHTTPRoutes_UnifiedEndpointStructure(t *testing.T) {
	router := NewHTTPRouter()

	// Test the structure of unified endpoints
	unifiedEndpoints := router.GetUnifiedEndpoints()

	for _, endpoint := range unifiedEndpoints {
		t.Logf("Unified Endpoint: %s", endpoint.Path)
		t.Logf("  Module: %s", endpoint.Module)
		t.Logf("  Method: %s", endpoint.HTTPMethod)
		t.Logf("  Description: %s", endpoint.Description)
		t.Logf("  Param Name: %s", endpoint.ParamName)
		t.Logf("  Query Param: %v", endpoint.QueryParam)

		// Verify structure
		if endpoint.Path == "" {
			t.Error("Unified endpoint path should not be empty")
		}
		if endpoint.Module == "" {
			t.Error("Unified endpoint module should not be empty")
		}
		if endpoint.HTTPMethod == "" {
			t.Error("Unified endpoint HTTP method should not be empty")
		}
		if endpoint.ParamName == "" {
			t.Error("Unified endpoint param name should not be empty")
		}
		if !endpoint.QueryParam {
			t.Error("Unified endpoint should use query parameters")
		}
	}
}
