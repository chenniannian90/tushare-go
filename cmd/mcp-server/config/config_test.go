package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFileWithAPITokens(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")

	// Sample configuration with api_tokens
	configData := map[string]interface{}{
		"host":  "0.0.0.0",
		"port":  8080,
		"transport": "stdio",
		"services": map[string]interface{}{
			"all": map[string]interface{}{
				"name": "all",
				"path": "/",
				"description": "All services",
				"categories": []string{},
			},
		},
		"global_auth": map[string]interface{}{
			"type": "none",
			"required": false,
		},
		"api_tokens": []string{"token1", "token2", "token3"},
	}

	// Write config to file
	data, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the config
	config, err := LoadFile(configPath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	// Verify api_tokens are loaded correctly
	if len(config.APITokens) != 3 {
		t.Errorf("Expected 3 API tokens, got %d", len(config.APITokens))
	}

	expectedTokens := []string{"token1", "token2", "token3"}
	for i, token := range config.APITokens {
		if token != expectedTokens[i] {
			t.Errorf("Token %d: expected %q, got %q", i, expectedTokens[i], token)
		}
	}
}

func TestLoadFileWithoutAPITokens(t *testing.T) {
	// Create a temporary config file without api_tokens
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")

	// Sample configuration without api_tokens
	configData := map[string]interface{}{
		"host":  "0.0.0.0",
		"port":  8080,
		"transport": "stdio",
		"services": map[string]interface{}{
			"all": map[string]interface{}{
				"name": "all",
				"path": "/",
				"description": "All services",
				"categories": []string{},
			},
		},
		"global_auth": map[string]interface{}{
			"type": "none",
			"required": false,
		},
	}

	// Write config to file
	data, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the config
	config, err := LoadFile(configPath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	// Verify api_tokens is empty but not nil
	if config.APITokens != nil && len(config.APITokens) != 0 {
		t.Errorf("Expected empty API tokens, got %v", config.APITokens)
	}
}
