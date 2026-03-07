package main

import (
	"fmt"
	"os"

	"github.com/chenniannian90/tushare-go/internal/gen"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: generate-single <spec-file> <output-file>")
		os.Exit(1)
	}

	specFile := os.Args[1]
	outputFile := os.Args[2]

	// Load spec
	spec, err := gen.LoadSpec(specFile)
	if err != nil {
		fmt.Printf("Error loading spec: %v\n", err)
		os.Exit(1)
	}

	// Generate code
	if err := gen.Generate(spec, outputFile); err != nil {
		fmt.Printf("Error generating API: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s from %s\n", outputFile, specFile)
}
