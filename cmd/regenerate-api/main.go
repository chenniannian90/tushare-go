package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chenniannian90/tushare-go/internal/gen"
)

// extractNameFromPath 从路径组件中提取 ___ 后面的部分
func extractNameFromPath(component string) string {
	if idx := strings.Index(component, "___"); idx != -1 {
		return component[idx+3:] // "___" 的长度是 3
	}
	return component
}

// getOutputPath 根据 spec 文件路径生成输出文件路径
func getOutputPath(specFile, specsRoot, outputRoot string) (string, error) {
	// 获取相对路径
	relPath, err := filepath.Rel(specsRoot, specFile)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}

	// 分割路径
	components := strings.Split(relPath, string(filepath.Separator))

	// 处理每个组件
	newComponents := make([]string, 0, len(components))
	for i, component := range components {
		// 最后一个组件是文件名，去掉 .json 后缀
		if i == len(components)-1 {
			component = strings.TrimSuffix(component, ".json")
		}
		newComponent := extractNameFromPath(component)
		newComponents = append(newComponents, newComponent)
	}

	// 构建输出路径
	if len(newComponents) == 0 {
		return "", fmt.Errorf("no components in path")
	}

	// 所有组件除了最后一个都是目录
	outputDir := filepath.Join(outputRoot, filepath.Join(newComponents[:len(newComponents)-1]...))
	outputFile := filepath.Join(outputDir, newComponents[len(newComponents)-1]+".go")

	return outputFile, nil
}

func main() {
	// 路径配置
	specsRoot := "internal/gen/specs"
	outputRoot := "pkg/sdk/api"

	// 将相对路径转换为绝对路径
	specsRootAbs, err := filepath.Abs(specsRoot)
	if err != nil {
		fmt.Printf("❌ Error: failed to get absolute path for specs root: %v\n", err)
		os.Exit(1)
	}
	outputRootAbs, err := filepath.Abs(outputRoot)
	if err != nil {
		fmt.Printf("❌ Error: failed to get absolute path for output root: %v\n", err)
		os.Exit(1)
	}

	// 检查目录是否存在
	if _, err := os.Stat(specsRootAbs); os.IsNotExist(err) {
		fmt.Printf("❌ Error: specs directory does not exist: %s\n", specsRootAbs)
		os.Exit(1)
	}

	fmt.Printf("🔍 Scanning directory: %s\n", specsRootAbs)
	fmt.Printf("📁 Output directory: %s\n", outputRootAbs)
	fmt.Println()

	// 列出所有 spec 文件
	specs, err := gen.ListSpecs()
	if err != nil {
		fmt.Printf("❌ Error: failed to list specs: %v\n", err)
		os.Exit(1)
	}

	// 统计信息
	totalFiles := len(specs)
	successCount := 0
	errorCount := 0

	// 生成所有文件
	for i, specPath := range specs {
		// 加载 spec
		spec, err := gen.LoadSpec(specPath)
		if err != nil {
			fmt.Printf("❌ Error: failed to load spec %s: %v\n", specPath, err)
			errorCount++
			continue
		}

		// 生成输出路径
		outputPath, err := getOutputPath(specPath, specsRootAbs, outputRootAbs)
		if err != nil {
			fmt.Printf("❌ Error: failed to generate output path for %s: %v\n", specPath, err)
			errorCount++
			continue
		}

		// 生成代码
		if err := gen.Generate(spec, outputPath); err != nil {
			fmt.Printf("❌ Error: failed to generate %s: %v\n", spec.APIName, err)
			errorCount++
			continue
		}

		successCount++

		// 显示进度
		if (i+1)%50 == 0 {
			fmt.Printf("   Processed %d/%d files...\n", i+1, totalFiles)
		}
	}

	// 打印摘要
	fmt.Println()
	fmt.Println("📊 Summary:")
	fmt.Printf("   Total files: %d\n", totalFiles)
	fmt.Printf("   Success: %d\n", successCount)
	fmt.Printf("   Failed: %d\n", errorCount)
	fmt.Println()

	if errorCount == 0 {
		fmt.Println("✅ All files generated successfully!")
	} else {
		fmt.Printf("⚠️  %d file(s) failed to generate\n", errorCount)
		os.Exit(1)
	}
}
