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
	case "string", "str":
		return "string"
	case "int", "integer":
		return "int"
	case "float", "float64", "double":
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

	// Extract package name from output path
	packageName := extractPackageName(outputPath)

	// Create template data with package name
	templateData := struct {
		*APISpec
		PackageName string
	}{
		APISpec:     spec,
		PackageName: packageName,
	}

	// Create template with custom functions
	tmpl, err := template.New("api.go.tmpl").Funcs(templateFuncs).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
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

// extractPackageName extracts the package name from the output file path
func extractPackageName(outputPath string) string {
	// Get the directory containing the output file
	outputDir := filepath.Dir(outputPath)

	// Get the last component of the directory path
	dirName := filepath.Base(outputDir)

	// Convert to a valid Go package name
	// Replace hyphens and dots with underscores
	goPackageName := strings.Map(func(r rune) rune {
		if r == '-' || r == '.' {
			return '_'
		}
		return r
	}, dirName)

	// Ensure the package name is a valid Go identifier
	if !isValidGoIdentifier(goPackageName) {
		return "api"
	}

	return goPackageName
}

// isValidGoIdentifier checks if a string is a valid Go identifier
func isValidGoIdentifier(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != '_' && (i == 0 || r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// ListSpecs returns a list of all spec file paths (including subdirectories)
func ListSpecs() ([]string, error) {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	specsDir := filepath.Join(currentDir, "specs")

	var specs []string
	err := filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // 继续遍历子目录
		}
		if filepath.Ext(path) != ".json" {
			return nil // 跳过非 JSON 文件
		}
		specs = append(specs, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk specs directory: %w", err)
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

		// 确定输出子目录
		var subDir string
		if spec.Describe != nil && spec.Describe.Category != "" {
			subDir = categoryToDir(spec.Describe.Category)
		} else {
			subDir = "other"
		}

		// Generate output filename with subdirectory
		outputFile := filepath.Join(outputDir, subDir, toSnakeCase(spec.APIName)+".go")

		// Generate code
		if err := Generate(spec, outputFile); err != nil {
			return count, fmt.Errorf("failed to generate %s: %w", spec.APIName, err)
		}

		count++
	}

	return count, nil
}

// categoryToDir maps Chinese category names to English directory names
func categoryToDir(category string) string {
	categoryMap := map[string]string{
		"行情数据": "market_data",
		"股票信息": "stock_info",
		"财务数据": "financial_data",
		"指数数据": "index_data",
		"交易日历": "trading_calendar",
		"权益数据": "equity_data",
	}

	if dir, ok := categoryMap[category]; ok {
		return dir
	}
	// 默认返回 "other" 如果没有匹配的分类
	return "other"
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
