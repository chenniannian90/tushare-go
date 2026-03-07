package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	apiBasePath string
	mcpToolsPath string
	modules     []APIModule
}

func main() {
	// Define command-line flags
	apiPath := flag.String("api-path", "pkg/sdk/api", "Path to the API directory")
	mcpToolsPath := flag.String("mcp-tools-path", "pkg/mcp/tools", "Path to the MCP tools output directory")
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
		fmt.Println("  gen-mcp-tools -api-path ./api -mcp-tools-path ./tools")
		fmt.Println("  gen-mcp-tools -api-path /absolute/path/to/api")
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
	mcpToolsOutputPath := resolvePath(projectRoot, *mcpToolsPath)

	fmt.Printf("Project root: %s\n", projectRoot)
	fmt.Printf("API path: %s\n", apiBasePath)
	fmt.Printf("MCP tools path: %s\n", mcpToolsOutputPath)

	generator := &MCPGenerator{
		apiBasePath:  apiBasePath,
		mcpToolsPath: mcpToolsOutputPath,
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

	fmt.Println("\n🔨 Generating MCP tools...")

	// Generate main tools file
	if err := g.generateMainToolsFile(); err != nil {
		return fmt.Errorf("failed to generate main tools file: %w", err)
	}

	// Generate individual module tool files
	for _, module := range g.modules {
		if err := g.generateModuleTools(module); err != nil {
			return fmt.Errorf("failed to generate tools for module %s: %w", module.Name, err)
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

	"github.com/chenniannian90/tushare-go/pkg/sdk"

	// Import all tool modules
	bondtools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/bond"
	etftools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/etf"
	forextools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/forex"
	fundtools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/fund"
	futurestools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/futures"
	hk_stocktools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/hk_stock"
	indextools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/index"
	llm_corpustools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/llm_corpus"
	optionstools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/options"
	spottools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/spot"
	us_stocktools "github.com/chenniannian90/tushare-go/pkg/mcp/tools/us_stock"

	"github.com/chenniannian90/tushare-go/pkg/mcp/common"
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

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/mcp/common"
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
		content += fmt.Sprintf("\t\t{\n")
		content += fmt.Sprintf("\t\t\tName: \"%s\",\n", toolName)
		content += fmt.Sprintf("\t\t\tDescription: \"Access %s data from Tushare API\",\n", apiName)
		content += fmt.Sprintf("\t\t},\n")
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
	outputPath := filepath.Join(moduleDir, module.Name+".go")
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
	"fmt"

	%s "github.com/chenniannian90/tushare-go/pkg/sdk/api/%s"
	"github.com/chenniannian90/tushare-go/pkg/mcp/common"
)

// call%s handles %s tool calls
func (m *%sTools) call%s(ctx context.Context, args map[string]interface{}) (*common.ToolResult, error) {
	req := &%s.%s{}

	// Parse arguments into request
	// TODO: Implement proper argument parsing based on request struct fields
	// For now, this is a placeholder implementation

	items, err := %s.%s(ctx, m.client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call %s API: %%w", err)
	}

	// Format results
	result, err := common.APIResult(items, "%s", "%s")
	if err != nil {
		return nil, err
	}
	return result, nil
}
`,
		packageName,
		importAlias, actualPackagePath,
		fn.Name, fn.Name,
		className, fn.Name,
		importAlias, requestType,
		importAlias, fn.Name, apiName,
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
	mainModuleFile := filepath.Join(moduleDir, module.Name + ".go")
	if outputPath == mainModuleFile {
		// Add a prefix to avoid conflict
		filename = "call_" + g.toSnakeCase(fn.Name) + ".go"
		outputPath = filepath.Join(moduleDir, filename)
	}

	return os.WriteFile(outputPath, []byte(content), 0644)
}
