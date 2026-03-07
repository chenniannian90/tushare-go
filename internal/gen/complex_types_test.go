package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestComplexTypesSupport(t *testing.T) {
	// Create a temporary directory for test output
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "complex_example.go")

	// Load the complex types example spec
	spec, err := LoadSpec("specs/complex_types_example.json")
	if err != nil {
		t.Fatalf("failed to load spec: %v", err)
	}

	// Verify spec has complex types
	if len(spec.ResponseFields) == 0 {
		t.Fatal("spec has no response fields")
	}

	// Find complex type fields
	var arrayField, objectField *ParamField
	for i := range spec.ResponseFields {
		if spec.ResponseFields[i].Type == "array" {
			arrayField = &spec.ResponseFields[i]
		}
		if spec.ResponseFields[i].Type == "object" {
			objectField = &spec.ResponseFields[i]
		}
	}

	if arrayField == nil {
		t.Error("spec should have at least one array field")
	}
	if objectField == nil {
		t.Error("spec should have at least one object field")
	}

	// Verify array field has items
	if arrayField != nil && arrayField.Items == nil {
		t.Error("array field should have items defined")
	}

	// Generate code
	if err := Generate(spec, outputPath); err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("generated file does not exist: %s", outputPath)
	}

	// Read generated code
	generatedCode, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read generated code: %v", err)
	}

	codeStr := string(generatedCode)

	// Verify correct types were generated
	if !contains(codeStr, "Tags []string") {
		t.Error("generated code should have Tags []string")
	}
	if !contains(codeStr, "Scores []float64") {
		t.Error("generated code should have Scores []float64")
	}
	if !contains(codeStr, "Metadata map[string]interface{}") {
		t.Error("generated code should have Metadata map[string]interface{}")
	}

	// Verify array handling code was generated
	if !contains(codeStr, "处理 tags 的数组类型") {
		t.Error("generated code should handle array type for tags")
	}
	if !contains(codeStr, "处理 scores 的数组类型") {
		t.Error("generated code should handle array type for scores")
	}

	// Verify object handling code was generated
	if !contains(codeStr, "处理 metadata 的对象类型") {
		t.Error("generated code should handle object type for metadata")
	}

	// Verify boolean parameter handling
	if !contains(codeStr, "if req.IncludeNested") {
		t.Error("generated code should handle boolean parameter correctly")
	}

	t.Log("✓ Complex types support test passed")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
