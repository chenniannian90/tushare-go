package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command-line flags
	apiDirFlag := flag.String("api-dir", "pkg/sdk/api/stock", "Path to the stock API directory")
	outputFileFlag := flag.String("output", "pkg/sdk/apis/stock.go", "Path to the output wrapper file")
	showHelp := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *showHelp {
		fmt.Println("Stock Wrapper Generator")
		fmt.Println("\nGenerates wrapper functions for stock API calls.")
		fmt.Println("\nUsage:")
		fmt.Println("  gen-stock-wrapper [options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  gen-stock-wrapper")
		fmt.Println("  gen-stock-wrapper -api-dir ./api/stock -output ./apis/stock.go")
		fmt.Println("  gen-stock-wrapper -api-dir /absolute/path/to/api/stock")
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
	apiDir := resolvePath(projectRoot, *apiDirFlag)
	outputFile := resolvePath(projectRoot, *outputFileFlag)

	fmt.Printf("Project root: %s\n", projectRoot)
	fmt.Printf("API directory: %s\n", apiDir)
	fmt.Printf("Output file: %s\n", outputFile)

	fset := token.NewFileSet()

	// Collect all API functions
	var functions []struct {
		Name       string
		Package    string
		ImportPath string
	}

	err = filepath.Walk(apiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil // Skip files that can't be parsed
		}

		// Get the package name
		pkgName := file.Name.Name

		// Skip if not in a subdirectory
		relPath, err := filepath.Rel(apiDir, path)
		if err != nil {
			return nil
		}

		dir := filepath.Dir(relPath)
		if dir == "." {
			return nil // Skip files in the root stock directory
		}

		// Build import path
		parts := strings.Split(dir, string(filepath.Separator))
		importPath := "tushare-go/pkg/sdk/api/stock/" + strings.Join(parts, "/")

		// Find exported functions
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}

			if fn.Recv == nil && fn.Name.IsExported() {
				functions = append(functions, struct {
					Name       string
					Package    string
					ImportPath string
				}{
					Name:       fn.Name.Name,
					Package:    pkgName,
					ImportPath: importPath,
				})
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Generate the wrapper file
	content := `// Package apis 提供类型化的链式调用方法
package apis

import (
	"context"

	"tushare-go/pkg/sdk"
`

	// Collect unique imports
	imports := make(map[string]struct{})
	for _, fn := range functions {
		imports[fn.ImportPath] = struct{}{}
	}

	for importPath := range imports {
		content += fmt.Sprintf("\t\"%s\"\n", importPath)
	}

	content += ")\n\n// ==================== Stock 方法 ====================\n\n"

	// Group functions by package
	pkgFunctions := make(map[string][]struct {
		Name       string
		Package    string
		ImportPath string
	})

	for _, fn := range functions {
		pkgFunctions[fn.Package] = append(pkgFunctions[fn.Package], fn)
	}

	// Generate wrapper functions
	for _, fn := range functions {
		// Create a unique function name by combining package and function name
		parts := strings.Split(fn.Package, "_")
		prefix := ""
		for _, part := range parts {
			if len(part) > 0 {
				prefix += strings.ToUpper(part[:1]) + part[1:]
			}
		}

		wrapperName := prefix + fn.Name
		content += fmt.Sprintf("// %s 包装函数\n", wrapperName)
		content += fmt.Sprintf("func %s(ctx context.Context, client *sdk.Client, req *%s.%sRequest) ([]%s.%sItem, error) {\n",
			wrapperName, fn.Package, fn.Name, fn.Package, fn.Name)
		content += fmt.Sprintf("\treturn %s.%s(ctx, client, req)\n", fn.Package, fn.Name)
		content += "}\n\n"
	}

	// Write to file
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d wrapper functions in %s\n", len(functions), outputFile)
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
