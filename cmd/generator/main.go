package main

import (
	"fmt"
	"os"

	"github.com/chenniannian90/tushare-go/internal/gen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: generator <output-dir>")
		os.Exit(1)
	}

	outputDir := os.Args[1]

	count, err := gen.GenerateAll(outputDir)
	if err != nil {
		fmt.Printf("Error generating APIs: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d API wrappers in %s\n", count, outputDir)
}
