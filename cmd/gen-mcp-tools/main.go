package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// APIModule represents a discovered API module
type APIModule struct {
	Name        string
	Path        string
	PackageName string
	Functions   []APIFunction
}

// APIFunction represents a discovered API function
type APIFunction struct {
	Name       string
	FileName   string
	StructName string
	Request    string
	Response   string
}

// MCPGenerator generates MCP tools from API modules
type MCPGenerator struct {
	apiBasePath   string
	mcpToolsPath  string
	specBasePath  string
	optimizedMode bool
	modules       []APIModule
}

// APISpec represents the structure of a Tushare API spec file
type APISpec struct {
	APIName       string       `json:"api_name"`
	APICode       string       `json:"api_code"`
	Description   string       `json:"description"`
	Describe      DescribeInfo `json:"__describe__"`
	RequestParams []FieldInfo  `json:"request_params"`
	ResponseFields []FieldInfo `json:"response_fields"`
}

// DescribeInfo contains metadata about the API
type DescribeInfo struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// FieldInfo represents a field in the API spec
type FieldInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// FieldWithGoType represents a field with Go type information
type FieldWithGoType struct {
	Name            string
	JSONName        string
	GoType          string
	Description     string
	StructFieldName string
	Required        bool
}

func main() {
	// Define command-line flags
	apiPath := flag.String("api-path", "pkg/sdk/api", "Path to the API directory")
	specPath := flag.String("spec-path", "internal/gen/specs", "Path to the spec files directory")
	mcpToolsPath := flag.String("mcp-tools-path", "pkg/mcp/tools", "Path to the MCP tools output directory")
	optimizedMode := flag.Bool("optimized", false, "Generate optimized MCP tools with JSON schema")
	showHelp := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *showHelp {
		fmt.Println("MCP Tools Generator")
		fmt.Println("\nGenerates MCP tool wrappers from API definitions.")
		fmt.Println("\nUsage:")
		fmt.Println("  gen-mcp-tools [options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  gen-mcp-tools")
		fmt.Println("  gen-mcp-tools -optimized")
		fmt.Println("  gen-mcp-tools -api-path ./api -mcp-tools-path ./tools")
		fmt.Println("  gen-mcp-tools -api-path /absolute/path/to/api -optimized")
		fmt.Println("\nOptimized Mode:")
		fmt.Println("  When -optimized is enabled, generates tools following the reference")
		fmt.Println("  implementation pattern with JSON schema support and better type safety.")
		os.Exit(0)
	}

	// Get the project root directory
	projectRoot, err := getProjectRoot()
	if err != nil {
		fmt.Printf("Error finding project root: %v\n", err)
		fmt.Println("Using current directory as project root")
		projectRoot = "."
	}

	// Resolve paths (support both relative and absolute paths)
	apiBasePath := resolvePath(projectRoot, *apiPath)
	specBasePath := resolvePath(projectRoot, *specPath)
	mcpToolsOutputPath := resolvePath(projectRoot, *mcpToolsPath)

	fmt.Printf("Project root: %s\n", projectRoot)
	fmt.Printf("API path: %s\n", apiBasePath)
	fmt.Printf("Spec path: %s\n", specBasePath)
	fmt.Printf("MCP tools path: %s\n", mcpToolsOutputPath)

	generator := &MCPGenerator{
		apiBasePath:   apiBasePath,
		specBasePath:  specBasePath,
		mcpToolsPath:  mcpToolsOutputPath,
		optimizedMode: *optimizedMode,
	}

	if err := generator.Generate(); err != nil {
		fmt.Printf("Error generating MCP tools: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ MCP tools generated successfully!")
}

// getProjectRoot finds the project root directory by looking for go.mod
func getProjectRoot() (string, error) {
 cwd, err := os.Getwd()
 if err != nil {
  return "", fmt.Errorf("failed to get working directory: %w", err)
 }

 // Search for go.mod file
 dir := cwd
 for {
  if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
   return dir, nil
  }

  parent := filepath.Dir(dir)
  if parent == dir {
   return "", fmt.Errorf("go.mod not found")
  }
  dir = parent
 }
}

// resolvePath resolves a path relative to the project root
func resolvePath(base, path string) string {
 if filepath.IsAbs(path) {
  return path
 }
 return filepath.Join(base, path)
}

// Generate scans API modules and generates MCP tools
func (g *MCPGenerator) Generate() error {
	fmt.Println("🔍 Scanning API modules...")

	// Scan API modules
	if err := g.scanAPIModules(); err != nil {
		return fmt.Errorf("failed to scan API modules: %w", err)
	}

	fmt.Printf("📦 Found %d API modules\n", len(g.modules))

	// Generate MCP tools for each module
	for _, module := range g.modules {
		fmt.Printf("  - %s: %d functions\n", module.Name, len(module.Functions))
	}

	if g.optimizedMode {
		fmt.Println("\n🔨 Generating optimized MCP tools with JSON schema...")
		fmt.Println("📝 Note: Standard mode generation skipped (optimized mode enabled)")
	} else {
		fmt.Println("\n🔨 Generating MCP tools...")
	}

	// Generate main tools file and standard tools (only in non-optimized mode)
	if !g.optimizedMode {
		if err := g.generateMainToolsFile(); err != nil {
			return fmt.Errorf("failed to generate main tools file: %w", err)
		}

		// Generate individual module tool files
		for _, module := range g.modules {
			if err := g.generateModuleTools(module); err != nil {
				return fmt.Errorf("failed to generate tools for module %s: %w", module.Name, err)
			}
		}
	}

	// Generate optimized tools for each module
	for _, module := range g.modules {
		if err := g.generateOptimizedModuleTools(module); err != nil {
			return fmt.Errorf("failed to generate optimized tools for module %s: %w", module.Name, err)
		}
	}

	fmt.Println("✅ All MCP tools generated successfully!")
	return nil
}

// scanAPIModules scans the API directory for modules
func (g *MCPGenerator) scanAPIModules() error {
	entries, err := os.ReadDir(g.apiBasePath)
	if err != nil {
		return fmt.Errorf("failed to read API directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		moduleName := entry.Name()
		modulePath := filepath.Join(g.apiBasePath, moduleName)

		module, err := g.scanAPIModule(moduleName, modulePath)
		if err != nil {
			return fmt.Errorf("failed to scan module %s: %w", moduleName, err)
		}

		if len(module.Functions) > 0 {
			g.modules = append(g.modules, module)
		}

		// Also scan subdirectories for modules that have nested structure
		if moduleName == "stock" || moduleName == "index" || moduleName == "wealth" || moduleName == "macro" || moduleName == "industry" {
			subEntries, err := os.ReadDir(modulePath)
			if err != nil {
				continue
			}

			for _, subEntry := range subEntries {
				if !subEntry.IsDir() {
					continue
				}

				subModuleName := subEntry.Name()
				subModulePath := filepath.Join(modulePath, subModuleName)

				subModule, err := g.scanAPIModule(subModuleName, subModulePath)
				if err != nil {
					continue
				}

				if len(subModule.Functions) > 0 {
					// Update the package name to include parent
					subModule.PackageName = moduleName + "/" + subModule.PackageName
					g.modules = append(g.modules, subModule)
				}

				// For macro module, scan one more level deep (macro_domestic has subdirectories)
				if moduleName == "macro" && subModuleName == "macro_domestic" {
					subSubEntries, err := os.ReadDir(subModulePath)
					if err != nil {
						continue
					}

					for _, subSubEntry := range subSubEntries {
						if !subSubEntry.IsDir() {
							continue
						}

						subSubModuleName := subSubEntry.Name()
						subSubModulePath := filepath.Join(subModulePath, subSubModuleName)

						subSubModule, err := g.scanAPIModule(subSubModuleName, subSubModulePath)
						if err != nil {
							continue
						}

						if len(subSubModule.Functions) > 0 {
							// Update the package name to include parent and grandparent
							subSubModule.PackageName = moduleName + "/" + subModuleName + "/" + subSubModule.PackageName
							g.modules = append(g.modules, subSubModule)
						}
					}
				}
			}
		}
	}

	return nil
}

// scanAPIModule scans a single API module
func (g *MCPGenerator) scanAPIModule(name, path string) (APIModule, error) {
	module := APIModule{
		Name:        name,
		Path:        path,
		PackageName: name,
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return module, fmt.Errorf("failed to read module directory: %w", err)
	}

	// First, try to find .go files in current directory
	hasGoFiles := false
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") && !strings.HasSuffix(entry.Name(), "_test.go") {
			hasGoFiles = true
			break
		}
	}

	// If current directory has .go files, scan them
	if hasGoFiles {
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
				continue
			}

			// Skip test files
			if strings.HasSuffix(entry.Name(), "_test.go") {
				continue
			}

			funcName := g.extractFunctionName(entry.Name())
			if funcName == "" {
				continue
			}

			module.Functions = append(module.Functions, APIFunction{
				Name:       funcName,
				FileName:   entry.Name(),
				StructName: g.toStructName(funcName),
			})
		}
	} else {
		// For modules with subdirectories, return empty to let parent handle subdirectories individually
		// The parent will scan each subdirectory as a separate module
		return module, nil
	}

	return module, nil
}

// extractFunctionName extracts the function name from filename
func (g *MCPGenerator) extractFunctionName(filename string) string {
	// Remove .go extension
	name := strings.TrimSuffix(filename, ".go")

	// Skip if it's a test file or special file
	if strings.HasSuffix(name, "_test") {
		return ""
	}

	// Convert filename to function name
	parts := strings.Split(name, "_")
	var funcParts []string
	for _, part := range parts {
		if len(part) > 0 {
			funcParts = append(funcParts, strings.ToUpper(part[:1])+part[1:])
		}
	}

	return strings.Join(funcParts, "")
}

// toStructName converts function name to struct name
func (g *MCPGenerator) toStructName(funcName string) string {
	return funcName + "Request"
}

// toToolName converts module and function name to tool name
func (g *MCPGenerator) toToolName(module, function string) string {
	return fmt.Sprintf("%s.%s", module, g.toSnakeCase(function))
}

// toToolDescription generates a better description for the tool
func (g *MCPGenerator) toToolDescription(module, function, apiName string) string {
	// Try to extract description from spec file first
	if spec := g.loadAPISpec(module, apiName); spec != nil && spec.Description != "" {
		return spec.Description
	}

	// Fallback to generated description
	functionName := strings.ReplaceAll(function, "_", " ")
	functionName = strings.ToLower(functionName)

	moduleName := strings.ReplaceAll(module, "_", " ")

	return fmt.Sprintf("Retrieve %s data from Tushare %s API", functionName, moduleName)
}

// loadAPISpec loads the API spec file for a given module and API code
func (g *MCPGenerator) loadAPISpec(module, apiCode string) *APISpec {
	// Use recursive search to find spec files
	// This handles paths with Chinese characters and nested structures
	targetFilename := "___" + apiCode + ".json"

	var foundSpecFile string
	err := filepath.WalkDir(g.specBasePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Continue walking even if there's an error
		}
		if d.IsDir() {
			return nil
		}

		// Check if the filename ends with the target pattern
		if strings.HasSuffix(filepath.Base(path), targetFilename) {
			foundSpecFile = path
			return filepath.SkipAll // Stop walking once we find a match
		}
		return nil
	})

	if err != nil {
		return nil
	}

	if foundSpecFile == "" {
		return nil
	}

	// Read and parse the spec file
	data, err := os.ReadFile(foundSpecFile)
	if err != nil {
		return nil
	}

	var spec APISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil
	}

	return &spec
}

// toAPIName converts function name to Tushare API name
func (g *MCPGenerator) toAPIName(function string) string {
	return g.toSnakeCase(function)
}

// toSnakeCase converts CamelCase to snake_case
func (g *MCPGenerator) toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

// generateMainToolsFile generates the main tools registry file
func (g *MCPGenerator) generateMainToolsFile() error {
	content := `// Code generated by gen-mcp-tools. DO NOT EDIT.

package mcp

import (
	"context"
	"fmt"

	"tushare-go/pkg/sdk"

	// Import all tool modules
`

	// Generate import statements for all modules dynamically
	for _, module := range g.modules {
		packageAlias := module.Name + "tools"
		content += fmt.Sprintf("\t%s \"tushare-go/pkg/mcp/tools/%s\"\n", packageAlias, module.Name)
	}

	content += `
	"tushare-go/pkg/mcp/common"
)

// ModuleTools is the interface for all API module tools
type ModuleTools interface {
	ListTools() []common.Tool
	HandlesTool(toolName string) bool
	CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error)
}

// ToolRegistry manages all available tools
type ToolRegistry struct {
	modules map[string]ModuleTools
	client  *sdk.Client
}

// NewToolRegistry creates a new tool registry
func NewToolRegistry(client *sdk.Client) *ToolRegistry {
	registry := &ToolRegistry{
		modules: make(map[string]ModuleTools),
		client:  client,
	}

	// Initialize all modules
	registry.initializeModules()

	return registry
}

// initializeModules initializes all API modules
func (r *ToolRegistry) initializeModules() {
`

	// Add module initialization for each discovered module
	for _, module := range g.modules {
		varName := module.Name + "Tools"
		className := strings.ToUpper(module.Name[:1]) + module.Name[1:]
		packageAlias := module.Name + "tools"
		content += fmt.Sprintf("\n\t// Initialize %s module\n\t%s := %s.New%sTools(r.client)\n\tr.modules[\"%s\"] = %s\n",
			module.Name, varName, packageAlias, className, module.Name, varName)
	}

	content += `}

// ListTools returns all available tools
func (r *ToolRegistry) ListTools() []common.Tool {
	var tools []common.Tool
	for _, module := range r.modules {
		tools = append(tools, module.ListTools()...)
	}
	return tools
}

// GetTools returns all available tools (alias for ListTools)
func (r *ToolRegistry) GetTools() []common.Tool {
	return r.ListTools()
}

// CallTool executes a tool call
func (r *ToolRegistry) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	// Find which module handles this tool
	for _, module := range r.modules {
		if module.HandlesTool(toolName) {
			return module.CallTool(ctx, toolName, args)
		}
	}

	return nil, fmt.Errorf("unknown tool: %s", toolName)
}

// GetModuleByPrefix returns the module that handles tools with the given prefix
func (r *ToolRegistry) GetModuleByPrefix(prefix string) (ModuleTools, bool) {
	if module, ok := r.modules[prefix]; ok {
		return module, true
	}
	return nil, false
}
`

	// Write to file
	outputPath := filepath.Join(g.mcpToolsPath, "../tools_registry.go")
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// generateModuleTools generates tools for a specific module
func (g *MCPGenerator) generateModuleTools(module APIModule) error {
	className := strings.ToUpper(module.Name[:1]) + module.Name[1:]
	packageName := module.Name + "tools"

	// Generate main module file
	if err := g.generateMainModuleFile(module, className, packageName); err != nil {
		return fmt.Errorf("failed to generate main module file: %w", err)
	}

	// Generate individual tool files
	for _, fn := range module.Functions {
		if err := g.generateToolFile(module, fn, className, packageName); err != nil {
			return fmt.Errorf("failed to generate tool file %s: %w", fn.Name, err)
		}
	}

	return nil
}

// generateMainModuleFile generates the main module file
func (g *MCPGenerator) generateMainModuleFile(module APIModule, className, packageName string) error {
	content := fmt.Sprintf(`// Code generated by gen-mcp-tools. DO NOT EDIT.

package %s

import (
	"context"
	"fmt"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/mcp/common"
)

// %sTools implements tools for %s API
type %sTools struct {
	client *sdk.Client
}

// New%sTools creates a new %s tools instance
func New%sTools(client *sdk.Client) *%sTools {
	return &%sTools{
		client: client,
	}
}

// GetAPIName returns the Tushare API name for this module
func (m *%sTools) GetAPIName() string {
	return "%s"
}

// ListTools returns all available tools in this module
func (m *%sTools) ListTools() []common.Tool {
	return []common.Tool{
`,
		packageName, className, module.Name, className,
		className, module.Name, className, className, className,
		className, module.Name, className)

	// Add tool definitions
	for _, fn := range module.Functions {
		toolName := g.toToolName(module.Name, fn.Name)
		apiName := g.toAPIName(fn.Name)
		description := g.toToolDescription(module.Name, fn.Name, apiName)
		content += "\t\t{\n"
		content += fmt.Sprintf("\t\t\tName: %q,\n", toolName)
		content += fmt.Sprintf("\t\t\tDescription: %q,\n", description)
		content += "\t\t},\n"
	}

	content += "\t}\n}\n\n"

	// Add HandlesTool method
	content += `// HandlesTool checks if this module handles a tool
func (m *` + className + `Tools) HandlesTool(toolName string) bool {
	tools := m.ListTools()
	for _, tool := range tools {
		if tool.Name == toolName {
			return true
		}
	}
	return false
}

// CallTool executes a tool call
func (m *` + className + `Tools) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	switch toolName {
`

	// Add case statements for each function
	for _, fn := range module.Functions {
		toolName := g.toToolName(module.Name, fn.Name)
		content += fmt.Sprintf("\tcase \"%s\":\n\t\treturn m.call%s(ctx, args)\n", toolName, fn.Name)
	}

	content += "\tdefault:\n\t\treturn nil, fmt.Errorf(\"unknown tool: %s\", toolName)\n\t}\n}\n"

	// Ensure the module directory exists
	moduleDir := filepath.Join(g.mcpToolsPath, module.Name)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Write to file
	outputPath := filepath.Join(moduleDir, "registry.go")
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// generateToolFile generates individual tool file
func (g *MCPGenerator) generateToolFile(module APIModule, fn APIFunction, className, packageName string) error {
	requestType := fn.StructName
	apiName := g.toAPIName(fn.Name)
	moduleName := module.Name

	// Use the actual package path from the module
	actualPackagePath := module.PackageName

	// Create a safe import alias by replacing slashes with underscores
	importAlias := strings.ReplaceAll(actualPackagePath, "/", "_")

	content := fmt.Sprintf(`// Code generated by gen-mcp-tools. DO NOT EDIT.

package %s

import (
	"context"

	%s "tushare-go/pkg/sdk/api/%s"
	"tushare-go/pkg/mcp/common"
)

// call%s handles %s tool calls
// This tool provides access to %s data from the Tushare API
func (m *%sTools) call%s(ctx context.Context, args map[string]interface{}) (*common.ToolResult, error) {
	req := &%s.%s{}

	// Parse arguments into request
	if err := common.ParseInput(args, req); err != nil {
		return common.ErrorResult(err), nil
	}

	items, err := %s.%s(ctx, m.client, req)
	if err != nil {
		return common.ErrorResult(err), nil
	}

	// Format results
	result, err := common.APIResult(items, "%s", "%s")
	if err != nil {
		return common.ErrorResult(err), nil
	}
	return result, nil
}
`,
		packageName,
		importAlias, actualPackagePath,
		fn.Name, fn.Name, apiName,
		className, fn.Name,
		importAlias, requestType,
		importAlias, fn.Name,
		moduleName, apiName)

	// Ensure the module directory exists
	moduleDir := filepath.Join(g.mcpToolsPath, module.Name)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Write to file (use snake_case for filename with prefix to avoid conflicts)
	filename := g.toSnakeCase(fn.Name) + ".go"
	outputPath := filepath.Join(moduleDir, filename)

	// Check if this would overwrite the main module file
	mainModuleFile := filepath.Join(moduleDir, "registry.go")
	if outputPath == mainModuleFile {
		// Add a prefix to avoid conflict
		filename = "call_" + g.toSnakeCase(fn.Name) + ".go"
		outputPath = filepath.Join(moduleDir, filename)
	}

	return os.WriteFile(outputPath, []byte(content), 0644)
}

// Template for optimized main module file
const optimizedMainModuleTemplate = `// Code generated by gen-mcp-tools. DO NOT EDIT.

package {{.PackageName}}

import (
	"tushare-go/pkg/sdk"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// {{.ClassName}}Tools represents {{.ModuleName}} tools
type {{.ClassName}}Tools struct {
	server *mcp.Server
	client *sdk.Client
}

// New{{.ClassName}}Tools creates a new instance
func New{{.ClassName}}Tools(server *mcp.Server, client *sdk.Client) *{{.ClassName}}Tools {
	return &{{.ClassName}}Tools{server: server, client: client}
}

// RegisterAll registers all tools
func (r *{{.ClassName}}Tools) RegisterAll() {
{{range .Functions}}	r.register{{.Name}}()
{{end}}
}
`

// Template for types.go file
const typesFileTemplate = `// Code generated by gen-mcp-tools. DO NOT EDIT.

package {{.PackageName}}

import (
{{range $import := .Imports}}	{{$import}}
{{end}}
)

{{range .Types}}
// {{.InputTypeName}} defines the input schema
type {{.InputTypeName}} struct {
{{range .Fields}}{{.FieldName}} {{.FieldType}} ` + "`json:{{.JSONName}},omitempty jsonschema:{{.Description}}`" + `
{{end}}
}

// {{.OutputTypeName}} defines the output schema
type {{.OutputTypeName}} struct {
	Data  []{{.DataType}} ` + "`json:data jsonschema:{{.APIName}} data list`" + `
	Total int              ` + "`json:total jsonschema:Total count`" + `
}
{{end}}
`

// Template for optimized tool registration (without type definitions)
const optimizedToolRegistrationTemplate = `// Code generated by gen-mcp-tools. DO NOT EDIT.

package {{.PackageName}}

import (
	"context"
	"encoding/json"
	"fmt"

	{{.ImportAlias}} "tushare-go/pkg/sdk/api/{{.ActualPackagePath}}"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// register{{.FunctionName}} registers the tool
func (r *{{.ClassName}}Tools) register{{.FunctionName}}() {
	inputSchema, _ := jsonschema.For[{{.FunctionName}}Input](nil)

	tool := &mcp.Tool{
		Name:        "{{.ToolName}}",
		Description: "{{.ToolDescription}}",
		InputSchema: inputSchema,
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var input {{.FunctionName}}Input
		if err := json.Unmarshal(req.Params.Arguments, &input); err != nil {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf(` + "`" + `{"error":"Invalid input: %v"}` + "`" + `, err)}},
			}, nil
		}

		apiReq := &{{.ImportAlias}}.{{.RequestType}}{
{{range .Fields}}{{.StructFieldName}}: input.{{.Name}},
{{end}}
		}

		items, err := {{.ImportAlias}}.{{.FunctionName}}(ctx, r.client, apiReq)
		if err != nil {
			return &mcp.CallToolResult{
				IsError: true,
				Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf(` + "`" + `{"error":"API call failed: %v"}` + "`" + `, err)}},
			}, nil
		}

		output := {{.FunctionName}}Output{
			Data:  items,
			Total: len(items),
		}

		outputJSON, _ := json.MarshalIndent(output, "", "  ")
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: string(outputJSON)}},
		}, nil
	}

	r.server.AddTool(tool, handler)
}
`

// generateOptimizedMainModuleFile generates the main optimized module file
func (g *MCPGenerator) generateOptimizedMainModuleFile(module APIModule, className, packageName string) error {
	// Create template data
	data := struct {
		PackageName string
		ClassName   string
		ModuleName  string
		Functions   []APIFunction
	}{
		PackageName: packageName,
		ClassName:   className,
		ModuleName:  module.Name,
		Functions:   module.Functions,
	}

	// Parse template
	tmpl, err := template.New("optimizedMainModule").Parse(optimizedMainModuleTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var content strings.Builder
	if err := tmpl.Execute(&content, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Ensure the module directory exists
	moduleDir := filepath.Join(g.mcpToolsPath, module.Name)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Write to file
	outputPath := filepath.Join(moduleDir, "registry.go")
	return os.WriteFile(outputPath, []byte(content.String()), 0644)
}

// generateOptimizedToolRegistration generates individual tool registration files
func (g *MCPGenerator) generateOptimizedToolRegistration(module APIModule, fn APIFunction, className, packageName string) error {
	requestType := fn.StructName
	apiName := g.toAPIName(fn.Name)
	moduleName := module.Name

	// Use the actual package path from the module
	actualPackagePath := module.PackageName

	// Create a safe import alias by replacing slashes with underscores
	importAlias := strings.ReplaceAll(actualPackagePath, "/", "_")

	// Extract fields from spec file
	specFields := g.extractFieldsFromSpec(module.Name, apiName)

	// Extract item type name from SDK API
	itemTypeName := g.extractItemTypeName(module.PackageName, fn.Name)

	// Convert spec fields to template fields
	fields := make([]struct {
		Name            string
		JSONName        string
		Type            string
		Description     string
		StructFieldName string
	}, len(specFields))
	for i, field := range specFields {
		fields[i] = struct {
			Name            string
			JSONName        string
			Type            string
			Description     string
			StructFieldName string
		}{
			Name:            field.Name,
			JSONName:        field.JSONName,
			Type:            field.GoType,
			Description:     field.Description,
			StructFieldName: field.StructFieldName,
		}
	}

	// Create template data
	data := struct {
		PackageName       string
		ImportAlias       string
		ActualPackagePath string
		FunctionName      string
		APIName           string
		ToolName          string
		ToolDescription   string
		ClassName         string
		ModuleName        string
		RequestType       string
		ItemTypeName       string
		Fields            []struct {
			Name            string
			JSONName        string
			Type            string
			Description     string
			StructFieldName string
		}
	}{
		PackageName:       packageName,
		ImportAlias:       importAlias,
		ActualPackagePath: actualPackagePath,
		FunctionName:      fn.Name,
		APIName:           apiName,
		ToolName:          g.toToolName(module.Name, fn.Name),
		ToolDescription:   g.toToolDescription(module.Name, fn.Name, apiName),
		ClassName:         className,
		ModuleName:        moduleName,
		RequestType:       requestType,
		ItemTypeName:       itemTypeName,
		Fields:            fields,
	}

	// Parse template
	tmpl, err := template.New("optimizedToolRegistration").Parse(optimizedToolRegistrationTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var content strings.Builder
	if err := tmpl.Execute(&content, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Ensure the module directory exists
	moduleDir := filepath.Join(g.mcpToolsPath, module.Name)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Write to file
	filename := g.toSnakeCase(fn.Name) + ".go"
	outputPath := filepath.Join(moduleDir, filename)

	return os.WriteFile(outputPath, []byte(content.String()), 0644)
}

// generateOptimizedModuleTools generates optimized tools for a specific module
func (g *MCPGenerator) generateOptimizedModuleTools(module APIModule) error {
	className := strings.ToUpper(module.Name[:1]) + module.Name[1:]
	packageName := module.Name + "tools"

	// Generate the types.go file with all Input/Output types
	if err := g.generateTypesFile(module, className, packageName); err != nil {
		return fmt.Errorf("failed to generate types file: %w", err)
	}

	// Generate the main optimized module file with tools struct
	if err := g.generateOptimizedMainModuleFile(module, className, packageName); err != nil {
		return fmt.Errorf("failed to generate optimized main module file: %w", err)
	}

	// Generate individual tool registration files (no struct definitions)
	for _, fn := range module.Functions {
		if err := g.generateOptimizedToolRegistration(module, fn, className, packageName); err != nil {
			return fmt.Errorf("failed to generate optimized tool registration %s: %w", fn.Name, err)
		}
	}

	return nil
}

// generateTypesFile generates the types.go file containing all Input/Output types
func (g *MCPGenerator) generateTypesFile(module APIModule, className, packageName string) error {
	// Collect unique imports needed for data types
	imports := make(map[string]string)
	types := make([]struct {
		InputTypeName  string
		OutputTypeName string
		APIName        string
		DataType       string
		Fields         []struct {
			FieldName    string
			FieldType    string
			JSONName     string
			Description  string
		}
	}, 0, len(module.Functions))

	for _, fn := range module.Functions {
		apiName := g.toAPIName(fn.Name)
		actualPackagePath := module.PackageName
		importAlias := strings.ReplaceAll(actualPackagePath, "/", "_")
		itemTypeName := g.extractItemTypeName(module.PackageName, fn.Name)

		// Add import for this API's item type
		importPath := fmt.Sprintf(`%s "tushare-go/pkg/sdk/api/%s"`, importAlias, actualPackagePath)
		imports[importPath] = importPath

		// Extract fields from spec
		specFields := g.extractFieldsFromSpec(module.Name, apiName)

		// Build fields list for types.go
		fields := make([]struct {
			FieldName   string
			FieldType   string
			JSONName    string
			Description string
		}, len(specFields))

		for i, field := range specFields {
			fields[i] = struct {
				FieldName   string
				FieldType   string
				JSONName    string
				Description string
			}{
				FieldName:   field.Name,
				FieldType:   field.GoType,
				JSONName:    field.JSONName,
				Description: field.Description,
			}
		}

		// Remove "Request" suffix from StructName for Input/Output type names
		baseTypeName := strings.TrimSuffix(fn.StructName, "Request")

		types = append(types, struct {
			InputTypeName  string
			OutputTypeName string
			APIName        string
			DataType       string
			Fields         []struct {
				FieldName    string
				FieldType    string
				JSONName     string
				Description  string
			}
		}{
			InputTypeName:  baseTypeName + "Input",
			OutputTypeName: baseTypeName + "Output",
			APIName:        apiName,
			DataType:       importAlias + "." + itemTypeName,
			Fields:         fields,
		})
	}

	// Convert imports map to slice
	importList := make([]string, 0, len(imports))
	for _, imp := range imports {
		importList = append(importList, imp)
	}

	// Create template data
	data := struct {
		PackageName string
		Imports     []string
		Types       []struct {
			InputTypeName  string
			OutputTypeName string
			APIName        string
			DataType       string
			Fields         []struct {
				FieldName    string
				FieldType    string
				JSONName     string
				Description  string
			}
		}
	}{
		PackageName: packageName,
		Imports:     importList,
		Types:       types,
	}

	// Parse template
	tmpl, err := template.New("typesFile").Parse(typesFileTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse types template: %w", err)
	}

	// Execute template
	var content strings.Builder
	if err := tmpl.Execute(&content, data); err != nil {
		return fmt.Errorf("failed to execute types template: %w", err)
	}

	// Ensure the module directory exists
	moduleDir := filepath.Join(g.mcpToolsPath, module.Name)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	// Write to types.go
	outputPath := filepath.Join(moduleDir, "types.go")
	return os.WriteFile(outputPath, []byte(content.String()), 0644)
}

// extractFieldsFromSpec extracts fields from API spec for generating Input structs
func (g *MCPGenerator) extractFieldsFromSpec(module, apiCode string) []FieldWithGoType {
	spec := g.loadAPISpec(module, apiCode)
	if spec == nil || len(spec.RequestParams) == 0 {
		// Return empty slice if no request params (API takes no input)
		return []FieldWithGoType{}
	}

	fields := make([]FieldWithGoType, 0, len(spec.RequestParams))
	for _, param := range spec.RequestParams {
		field := FieldWithGoType{
			Name:            g.toPascalCase(param.Name),
			JSONName:        param.Name,
			GoType:          g.specTypeToGoType(param.Type),
			Description:     param.Description,
			StructFieldName: g.toPascalCase(param.Name),
			Required:        param.Required,
		}
		fields = append(fields, field)
	}

	return fields
}

// specTypeToGoType converts spec type to Go type
func (g *MCPGenerator) specTypeToGoType(specType string) string {
	switch strings.ToLower(specType) {
	case "str", "string": //nolint:goconst
		return "string"
	case "int", "integer":
		return "int"
	case "float", "float64", "double":
		return "float64"
	case "bool", "boolean":
		return "bool"
	default:
		return "string"
	}
}

// toPascalCase converts snake_case to PascalCase
func (g *MCPGenerator) toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			result.WriteString(part[1:])
		}
	}
	return result.String()
}

// extractItemTypeName extracts the item type name from SDK API file
func (g *MCPGenerator) extractItemTypeName(packageName, functionName string) string {
	// Try to read the SDK API file
	apiFileName := g.toSnakeCase(functionName)
	apiFilePath := filepath.Join(g.apiBasePath, packageName, apiFileName+".go")

	data, err := os.ReadFile(apiFilePath)
	if err != nil {
		// Fallback to function name + Item
		return functionName + "Item"
	}

	// Search for Item type definition
	content := string(data)
	searchPattern := functionName + "Item struct"

	// Find the struct definition
	idx := strings.Index(content, searchPattern)
	if idx == -1 {
		// Fallback to function name + Item
		return functionName + "Item"
	}

	return functionName + "Item"
}
