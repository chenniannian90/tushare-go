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

// getCurrentDir returns the directory of the file that calls this function
func getCurrentDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "."
	}
	return filepath.Dir(file)
}

// toGoType converts API type to Go type

// toGoType converts API type to Go type
func toGoType(apiType string) string {
	switch strings.ToLower(apiType) {
	case "string", "str": //nolint:goconst
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

	for i, word := range words {
		if word == "" {
			continue
		}

		// Convert string to rune slice for proper Unicode handling
		runes := []rune(word)
		if len(runes) == 0 {
			continue
		}

		// Handle words starting with numbers
		if runes[0] >= '0' && runes[0] <= '9' {
			// For first word starting with number, add "Field" prefix
			if i == 0 {
				result.WriteString("Field")
			}
			// Append the word as-is
			result.WriteString(word)
		} else {
			// Capitalize first rune, keep rest as is
			// Use Unicode-aware case conversion
			first := runes[0]
			rest := runes[1:]

			// Convert first rune to uppercase
			if first >= 'a' && first <= 'z' {
				first = first - 'a' + 'A'
			}

			result.WriteRune(first)
			for _, r := range rest {
				result.WriteRune(r)
			}
		}
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
	"toGoFieldName":  toGoFieldName,
	"hasStringFields": hasStringFields,
}

// hasStringFields checks if any field is a string type
func hasStringFields(fields []ParamField) bool {
	for _, f := range fields {
		if strings.ToLower(f.Type) == "str" || strings.ToLower(f.Type) == "string" {
			return true
		}
	}
	return false
}

// Generate generates API wrapper code from a spec
func Generate(spec *APISpec, outputPath string) error {
	// Get the directory where this file is located
	currentDir := getCurrentDir()

	// Template is in templates/ subdirectory relative to this file
	tmplPath := filepath.Join(currentDir, "templates", "api.go.tmpl")
	tmplContent, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read template from %s: %w", tmplPath, err)
	}

	// Extract package name from output path
	packageName := extractPackageName(outputPath)

	// Create template data with package name
	// Clean description to avoid code generation issues
	cleanSpec := *spec
	cleanSpec.Description = cleanDescription(spec.Description)

	templateData := struct {
		*APISpec
		PackageName string
	}{
		APISpec:     &cleanSpec,
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
	currentDir := getCurrentDir()
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
		// Use APICode (English) instead of APIName (Chinese) for filename
		filename := spec.APICode
		if filename == "" {
			filename = toSnakeCase(spec.APIName)
		}
		outputFile := filepath.Join(outputDir, subDir, filename+".go")

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
	// Handle format "中文名___英文代码" (e.g., "基础数据___stock_basic")
	if strings.Contains(category, "___") {
		parts := strings.Split(category, "___")
		if len(parts) >= 2 {
			// Return the English code part as directory name
			return parts[len(parts)-1]
		}
	}

	// Fallback mapping for old format (纯中文分类名)
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

// cleanDescription cleans the description string for use in code generation
func cleanDescription(desc string) string {
	if desc == "" {
		return desc
	}

	// Remove newlines and replace with space
	cleaned := strings.ReplaceAll(desc, "\n", " ")
	cleaned = strings.ReplaceAll(cleaned, "\r", " ")

	// Remove excessive whitespace
	cleaned = strings.TrimSpace(cleaned)

	// Truncate if too long (max 200 characters for code comments)
	// Ensure we don't cut in the middle of a UTF-8 character
	if len(cleaned) > 200 {
		// Convert to runes to handle multi-byte characters properly
		runes := []rune(cleaned)
		if len(runes) > 197 { // Leave room for "..."
			cleaned = string(runes[:197]) + "..."
		}
	}

	return cleaned
}

// toGoFieldName converts a field name to a valid Go identifier
// For Chinese field names, it uses the JSON field name (usually pinyin or English)
func toGoFieldName(fieldName string) string {
	if fieldName == "" {
		return fieldName
	}

	// Check if field name contains Chinese characters
	if hasChinese(fieldName) {
		// For Chinese field names, use "Field" prefix to make it valid Go identifier
		// and avoid UTF-8 encoding issues
		return "Field_" + fieldName
	}

	// For English field names, use standard PascalCase
	return toPascalCase(fieldName)
}

// hasChinese checks if a string contains Chinese characters
func hasChinese(s string) bool {
	for _, r := range s {
		if r >= 0x4E00 && r <= 0x9FFF {
			return true
		}
	}
	return false
}
