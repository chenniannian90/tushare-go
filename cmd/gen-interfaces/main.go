package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// APIInfo holds information about a generated API function
type APIInfo struct {
	FunctionName string // e.g., "EtfShareSize"
	RequestType  string // e.g., "EtfShareSizeRequest"
	ItemType     string // e.g., "EtfShareSizeItem"
	Package      string // e.g., "bond", "etf", "stock"
	FileName     string // source file name
	Description  string // API description
}

// CategoryAPIs holds all APIs for a category
type CategoryAPIs struct {
	CategoryName string // e.g., "Bond", "ETF", "Stock"
	PackageName  string // e.g., "bond", "etf", "stock"
	APIs         []APIInfo
}

func main() {
	// 配置
	apiDir := "pkg/sdk/api"
	outputDir := "pkg/sdk/apis"

	// 扫描所有生成的 API
	categories, err := scanGeneratedAPIs(apiDir)
	if err != nil {
		fmt.Printf("❌ Error scanning APIs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📁 Found %d categories\n", len(categories))

	// 生成接口文件
	for _, category := range categories {
		if err := generateInterfaceFile(category, outputDir); err != nil {
			fmt.Printf("❌ Error generating %s: %v\n", category.CategoryName, err)
			continue
		}
		fmt.Printf("✅ Generated %s\n", category.CategoryName)
	}

	// 生成主文件
	if err := generateMainFile(categories, outputDir); err != nil {
		fmt.Printf("❌ Error generating main file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Generated apis.go\n")

	fmt.Printf("\n🎉 Generated %d interface files\n", len(categories)+1)
}

// scanGeneratedAPIs scans pkg/sdk/api for generated API functions
func scanGeneratedAPIs(apiDir string) ([]CategoryAPIs, error) {
	categoryMap := make(map[string][]APIInfo)

	err := filepath.Walk(apiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// 读取文件内容
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// 提取包名
		packageName := extractPackageName(path, apiDir)
		if packageName == "" {
			return nil
		}

		// 提取 API 函数信息
		apis := extractAPIInfo(string(content), packageName)
		categoryMap[packageName] = append(categoryMap[packageName], apis...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 转换为 CategoryAPIs 列表
	var categories []CategoryAPIs
	for pkgName, apis := range categoryMap {
		// 从包路径中提取最后一个组件作为类别名
		// 例如：macro/macro_domestic/macro_business -> macro_business -> MacroBusiness
		lastComponent := filepath.Base(pkgName)
		categoryName := toPascalCase(lastComponent)
		categories = append(categories, CategoryAPIs{
			CategoryName: categoryName,
			PackageName:  pkgName,
			APIs:         apis,
		})
	}

	return categories, nil
}

// extractPackageName extracts package name from file path
func extractPackageName(filePath string, baseDir string) string {
	// Get the directory containing the file
	dir := filepath.Dir(filePath)

	// Get the relative path from baseDir
	relPath, err := filepath.Rel(baseDir, dir)
	if err != nil {
		// Fallback to just the directory name if relative path fails
		return filepath.Base(dir)
	}

	// Convert to forward slashes for Go package paths
	return filepath.ToSlash(relPath)
}

// extractAPIInfo extracts API function information from Go source code
func extractAPIInfo(content, packageName string) []APIInfo {
	var apis []APIInfo

	// 匹配: func FunctionName(ctx context.Context, client *sdk.Client, req *RequestType) ([]ItemType, error)
	funcRegex := regexp.MustCompile(`func (\w+)\(ctx context\.Context, client \*sdk\.Client, req \*(\w+Request)\) \(\[\](\w+Item), error\)`)
	matches := funcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			functionName := match[1]
			// 提取函数上方的注释作为描述
			description := extractFunctionDescription(content, functionName)

			apis = append(apis, APIInfo{
				FunctionName: functionName,
				RequestType:  match[2],
				ItemType:     match[3],
				Package:      packageName,
				Description:  description,
			})
		}
	}

	return apis
}

// extractFunctionDescription extracts the description comment above a function
func extractFunctionDescription(content, functionName string) string {
	// 查找函数定义的位置
	funcPattern := regexp.MustCompile(`func ` + functionName + `\(`)
	funcIndex := funcPattern.FindStringIndex(content)
	if funcIndex == nil {
		return ""
	}

	// 从函数定义向前查找注释
	lines := strings.Split(content[:funcIndex[0]], "\n")
	var commentLines []string

	// 从后向前查找，直到遇到非注释行
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			// 移除 "// " 前缀
			commentText := strings.TrimPrefix(line, "//")
			commentText = strings.TrimSpace(commentText)
			commentLines = append([]string{commentText}, commentLines...)
		} else {
			// 遇到非注释行，停止
			break
		}
	}

	if len(commentLines) > 0 {
		return strings.Join(commentLines, " ")
	}
	return ""
}

// toPascalCase converts to PascalCase
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
		result.WriteString(strings.ToUpper(word[:1]))
		result.WriteString(word[1:])
	}
	return result.String()
}

// toLowerCamelCase converts to lowerCamelCase
func toLowerCamelCase(s string) string {
	if s == "" {
		return s
	}
	pascal := toPascalCase(s)
	return strings.ToLower(pascal[:1]) + pascal[1:]
}

// toSnakeCase converts to snake_case
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

// generateInterfaceFile generates an interface file for a category
func generateInterfaceFile(category CategoryAPIs, outputDir string) error {
	// 手动生成内容以避免模板上下文问题
	var builder strings.Builder

	// 生成包别名 - 将路径中的斜杠替换为下划线
	packageAlias := strings.ReplaceAll(category.PackageName, "/", "_")

	builder.WriteString("// Code generated by tushare-gen. DO NOT EDIT.\n\n")
	builder.WriteString("package apis\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("\t\"context\"\n")
	builder.WriteString("\t\"tushare-go/pkg/sdk\"\n")
	builder.WriteString("\t" + packageAlias + " \"tushare-go/pkg/sdk/api/" + category.PackageName + "\"\n")
	builder.WriteString(")\n\n")

	builder.WriteString("type " + category.CategoryName + " interface {\n")
	for _, api := range category.APIs {
		if api.Description != "" {
			builder.WriteString("\t// " + api.Description + "\n")
		}
		builder.WriteString("\t" + api.FunctionName + "(ctx context.Context, req *" + packageAlias + "." + api.RequestType + ") ([]" + packageAlias + "." + api.ItemType + ", error)\n")
	}
	builder.WriteString("}\n\n")

	builder.WriteString("type " + toLowerCamelCase(category.CategoryName) + "Impl struct {\n")
	builder.WriteString("\tclient *sdk.Client\n")
	builder.WriteString("}\n\n")

	for _, api := range category.APIs {
		if api.Description != "" {
			builder.WriteString("// " + api.Description + "\n")
		}
		builder.WriteString("func (impl *" + toLowerCamelCase(category.CategoryName) + "Impl) " + api.FunctionName + "(ctx context.Context, req *" + packageAlias + "." + api.RequestType + ") ([]" + packageAlias + "." + api.ItemType + ", error) {\n")
		builder.WriteString("\treturn " + packageAlias + "." + api.FunctionName + "(ctx, impl.client, req)\n")
		builder.WriteString("}\n\n")
	}

	builder.WriteString("func new" + toLowerCamelCase(category.CategoryName) + "Impl(client *sdk.Client) " + category.CategoryName + " {\n")
	builder.WriteString("\treturn &" + toLowerCamelCase(category.CategoryName) + "Impl{client: client}\n")
	builder.WriteString("}\n")

	// 写入文件 - 所有接口文件都在同一目录
	outputFile := filepath.Join(outputDir, toSnakeCase(category.CategoryName)+".go")
	// 确保输出目录存在
	outputDirPath := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDirPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	if err := os.WriteFile(outputFile, []byte(builder.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// generateMainFile generates the main apis.go file
func generateMainFile(categories []CategoryAPIs, outputDir string) error {
	var builder strings.Builder

	builder.WriteString("// Package apis 提供类型化的链式调用方法\n")
	builder.WriteString("// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法\n")
	builder.WriteString("package apis\n\n")
	builder.WriteString("import \"tushare-go/pkg/sdk\"\n\n")

	builder.WriteString("// TushareClient 提供链式调用客户端，集成所有 API 接口\n")
	builder.WriteString("type TushareClient struct {\n")
	builder.WriteString("\tclient *sdk.Client\n")
	for _, category := range categories {
		builder.WriteString("\t" + category.CategoryName + " " + category.CategoryName + "\n")
	}
	builder.WriteString("}\n\n")

	builder.WriteString("// NewTushareClient 创建一个新的链式调用客户端\n")
	builder.WriteString("func NewTushareClient(client *sdk.Client) *TushareClient {\n")
	builder.WriteString("\treturn &TushareClient{\n")
	builder.WriteString("\t\tclient: client,\n")
	for _, category := range categories {
		builder.WriteString("\t\t" + category.CategoryName + ": new" + toLowerCamelCase(category.CategoryName) + "Impl(client),\n")
	}
	builder.WriteString("\t}\n")
	builder.WriteString("}\n")

	// 写入文件
	outputFile := filepath.Join(outputDir, "apis.go")
	if err := os.WriteFile(outputFile, []byte(builder.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
