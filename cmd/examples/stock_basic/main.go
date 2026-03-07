package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api"
)

func main() {
	// Get Tushare token from environment
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	// Create SDK client
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	client := sdk.NewClient(config)

	// Example 1: Get all stocks
	fmt.Println("=== Example 1: Get all stocks ===")
	allStocks, err := api.StockBasic(context.Background(), client, &api.StockBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get stocks: %v", err)
	}
	fmt.Printf("Found %d stocks total\n", len(allStocks))

	// Example 2: Get specific stock by code
	fmt.Println("\n=== Example 2: Get stock by code ===")
	stock, err := api.StockBasic(context.Background(), client, &api.StockBasicRequest{
		TsCode: "000001.SZ",
	})
	if err != nil {
		log.Fatalf("Failed to get stock: %v", err)
	}
	if len(stock) > 0 {
		s := stock[0]
		fmt.Printf("Stock: %s (%s) - %s - %s\n", s.TsCode, s.Symbol, s.Name, s.Industry)
	}

	// Example 3: Get stocks by exchange
	fmt.Println("\n=== Example 3: Get SSE stocks ===")
	sseStocks, err := api.StockBasic(context.Background(), client, &api.StockBasicRequest{
		Exchange: "SSE",
	})
	if err != nil {
		log.Fatalf("Failed to get SSE stocks: %v", err)
	}
	fmt.Printf("Found %d SSE stocks\n", len(sseStocks))
	if len(sseStocks) > 0 {
		fmt.Println("First 5 SSE stocks:")
		for i, s := range sseStocks {
			if i >= 5 {
				break
			}
			fmt.Printf("  - %s (%s): %s\n", s.TsCode, s.Symbol, s.Name)
		}
	}

	// Example 4: Get listed stocks only
	fmt.Println("\n=== Example 4: Get listed stocks only ===")
	listedStocks, err := api.StockBasic(context.Background(), client, &api.StockBasicRequest{
		ListStatus: "L",
	})
	if err != nil {
		log.Fatalf("Failed to get listed stocks: %v", err)
	}
	fmt.Printf("Found %d listed stocks\n", len(listedStocks))
}
