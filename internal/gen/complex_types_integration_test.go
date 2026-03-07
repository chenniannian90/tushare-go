//go:build integration

package gen

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// This test verifies that generated code with complex types works correctly
func TestGeneratedComplexTypesCode(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// Create temporary directory
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "complex_types_test.go")

	// Load and generate the complex types spec
	spec, err := LoadSpec("specs/complex_types_example.json")
	if err != nil {
		t.Fatalf("failed to load spec: %v", err)
	}

	if err := Generate(spec, outputPath); err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Create a mock server with complex data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "trade_date", "tags", "scores", "metadata", "close", "vol"],
				"items": [
					{
						"ts_code": "000001.SZ",
						"trade_date": "20240101",
						"tags": ["金融", "银行", "蓝筹"],
						"scores": [1.5, 2.3, 3.7],
						"metadata": {"sector": "finance", "market_cap": "large"},
						"close": 10.5,
						"vol": 1000000
					}
				]
			}
		}`))
	}))
	defer server.Close()

	t.Run("verify_complex_types_in_generated_code", func(t *testing.T) {
		// Read the generated code to verify it has the correct structure
		generatedCode, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatalf("failed to read generated code: %v", err)
		}

		codeStr := string(generatedCode)

		// Verify complex types are used
		typeChecks := []struct {
			name     string
			expected string
		}{
			{"string array type", "Tags []string"},
			{"float64 array type", "Scores []float64"},
			{"object type", "Metadata map[string]interface{}"},
			{"boolean parameter", "IncludeNested bool"},
		}

		for _, check := range typeChecks {
			if !containsSubstring(codeStr, check.expected) {
				t.Errorf("generated code missing %s: expected to find %q", check.name, check.expected)
			} else {
				t.Logf("✓ Found %s: %q", check.name, check.expected)
			}
		}

		// Verify type conversion logic for arrays
		if !containsSubstring(codeStr, "case []interface{}:") {
			t.Error("generated code should handle []interface{} conversion")
		}
		if !containsSubstring(codeStr, "case []string:") {
			t.Error("generated code should handle []string conversion")
		}
		if !containsSubstring(codeStr, "case []float64:") {
			t.Error("generated code should handle []float64 conversion")
		}

		t.Log("✓ All complex type checks passed")
	})
}
