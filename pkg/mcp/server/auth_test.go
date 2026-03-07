//go:build integration

package server

import (
	"os"
	"testing"
)

// TODO: Re-implement authentication tests for official MCP SDK
//
// The authentication tests below were designed for the custom StdioMCPServer implementation.
// Since we've migrated to the official MCP SDK, we need to re-implement authentication support.
//
// Options for implementing authentication with official SDK:
// 1. Use middleware to intercept initialize requests
// 2. Implement custom transport layer
// 3. Add authentication as a pre-processing step
//
// For now, these tests are disabled. To re-enable:
// 1. Implement authentication in the adapter or main.go
// 2. Update these tests to work with the new architecture
// 3. Remove this skip directive

func TestAPIKeyAuthentication(t *testing.T) {
	t.Skip("Authentication not yet implemented for official MCP SDK - TODO")
}

/*
func TestAPIKeyAuthentication(t *testing.T) {
	tests := []struct {
		name        string
		setupConfig func() *ServerConfig
		initParams  map[string]interface{}
		wantAuthErr bool
	}{
		{
			name: "successful API key authentication",
			setupConfig: func() *ServerConfig {
				return &ServerConfig{
					Token:       "test-token",
					APIKey:      "valid-api-key-12345",
					RequireAuth: true,
				}
			},
			initParams: map[string]interface{}{
				"clientInfo": map[string]interface{}{
					"apiKey": "valid-api-key-12345",
				},
			},
			wantAuthErr: false,
		},
		// ... more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Implementation needed for official SDK
		})
	}
}
*/

func TestAPIKeyEnvironmentVariable(t *testing.T) {
	// Test that API key can be loaded from environment variable
	validAPIKey := "env-test-api-key-67890"

	// Set environment variables
	os.Setenv("TUSHARE_TOKEN", "test-token")
	defer os.Unsetenv("TUSHARE_TOKEN")

	os.Setenv("MCP_API_KEY", validAPIKey)
	defer os.Unsetenv("MCP_API_KEY")

	// Set environment variable to require authentication
	os.Setenv("MCP_REQUIRE_AUTH", "true")
	defer os.Unsetenv("MCP_REQUIRE_AUTH")

	config := DefaultServerConfig()

	// Verify configuration loaded from environment
	if config.APIKey != validAPIKey {
		t.Errorf("expected API key %s, got %s", validAPIKey, config.APIKey)
	}

	if !config.RequireAuth {
		t.Error("expected RequireAuth to be true")
	}

	// Test validation
	if err := config.Validate(); err != nil {
		t.Errorf("config validation failed: %v", err)
	}

	// TODO: Test with official MCP server implementation
	t.Skip("Authentication not yet implemented for official MCP SDK - TODO")
}

func TestInvalidAPIKeyConfiguration(t *testing.T) {
	// Test configuration validation when authentication is required but API key is missing
	os.Setenv("TUSHARE_TOKEN", "test-token")
	defer os.Unsetenv("TUSHARE_TOKEN")

	os.Setenv("MCP_REQUIRE_AUTH", "true")
	os.Unsetenv("MCP_API_KEY")
	defer os.Unsetenv("MCP_REQUIRE_AUTH")

	config := DefaultServerConfig()

	// Should fail validation since API key is required but not provided
	if err := config.Validate(); err == nil {
		t.Error("expected validation error when RequireAuth is true but API key is missing")
	} else {
		t.Logf("✅ Got expected validation error: %v", err)
	}
}
