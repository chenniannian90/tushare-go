package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	stockmarket "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_market"
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

	// Example: Get daily basic metrics
	fmt.Println("=== Daily Basic Metrics Example ===")
	basicData, err := stockmarket.DailyBasic(context.Background(), client, &stockmarket.DailyBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get daily basic data: %v", err)
	}

	fmt.Printf("Found %d daily basic records\n", len(basicData))
	if len(basicData) > 0 {
		fmt.Println("\nDaily Basic Metrics:")
		for _, d := range basicData {
			fmt.Printf("\nData: %+v\n", d)
			break
		}
	}
}
