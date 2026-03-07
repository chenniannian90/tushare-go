package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

// toGoType converts API type to Go type
func toGoType(apiType string) string {
	switch strings.ToLower(apiType) {
	case "string":
		return "string"
	case "int", "integer":
		return "int"
	case "float", "double":
		return "float64"
	case "bool", "boolean":
		return "bool"
	case "array":
		return "[]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return "string" // default to string for unknown types
	}
}

// isArrayType checks if the type is an array
func isArrayType(field ParamField) bool {
	return strings.ToLower(field.Type) == "array"
}

// isObjectType checks if the type is an object
func isObjectType(field ParamField) bool {
	return strings.ToLower(field.Type) == "object"
}

// getArrayType returns the Go type for array elements
func getArrayType(field ParamField) string {
	if field.Items != nil {
		return toGoType(field.Items.Type)
	}
	return "interface{}"
}

// getFullArrayType returns the full Go array type
func getFullArrayType(field ParamField) string {
	elemType := "interface{}"
	if field.Items != nil {
		elemType = toGoType(field.Items.Type)
	}
	return "[]" + elemType
}

// toPascalCase converts snake_case to PascalCase
func toPascalCase(s string) string {
	if s == "" {
		return s
	}

	words := strings.Split(s, "_")
	var result strings.Builder

	for _, word := range words {
		if word == "" {
			continue
		}
		// Capitalize first letter, keep rest as is
		result.WriteString(strings.ToUpper(word[:1]))
		result.WriteString(word[1:])
	}

	return result.String()
}

// toLowerCamel converts a string to lowerCamelCase
func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	// Convert to PascalCase first, then lowercase first letter
	pascal := toPascalCase(s)
	return strings.ToLower(pascal[:1]) + pascal[1:]
}

// templateFuncs provides custom template functions
var templateFuncs = template.FuncMap{
	"Title":         toPascalCase,
	"goType":        toGoType,
	"lowerCamel":    toLowerCamel,
	"isArray":       isArrayType,
	"isObject":      isObjectType,
	"arrayElemType": getArrayType,
	"fullArrayType": getFullArrayType,
}

// Generate generates API wrapper code from a spec
func Generate(spec *APISpec, outputPath string) error {
	// Get the directory where this file is located
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)

	// Template is in templates/ subdirectory relative to this file
	tmplPath := filepath.Join(currentDir, "templates", "api.go.tmpl")
	tmplContent, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read template from %s: %w", tmplPath, err)
	}

	// Create template with custom functions
	tmpl, err := template.New("api.go.tmpl").Funcs(templateFuncs).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, spec); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Create output directory if needed
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write output file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// ListSpecs returns a list of all spec file paths
func ListSpecs() ([]string, error) {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	specsDir := filepath.Join(currentDir, "specs")

	entries, err := os.ReadDir(specsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read specs directory: %w", err)
	}

	var specs []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		specs = append(specs, filepath.Join(specsDir, entry.Name()))
	}

	return specs, nil
}

// GenerateAll generates API wrappers for all spec files
func GenerateAll(outputDir string) (int, error) {
	specs, err := ListSpecs()
	if err != nil {
		return 0, fmt.Errorf("failed to list specs: %w", err)
	}

	count := 0
	for _, specPath := range specs {
		// Load spec
		spec, err := LoadSpec(specPath)
		if err != nil {
			return count, fmt.Errorf("failed to load spec %s: %w", specPath, err)
		}

		// Generate output filename
		outputFile := filepath.Join(outputDir, toSnakeCase(spec.APIName)+".go")

		// Generate code
		if err := Generate(spec, outputFile); err != nil {
			return count, fmt.Errorf("failed to generate %s: %w", spec.APIName, err)
		}

		count++
	}

	return count, nil
}

// toSnakeCase converts PascalCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
