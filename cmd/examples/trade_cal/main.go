package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	stockbasic "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_basic"
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

	// Example: Get trading calendar
	fmt.Println("=== Trading Calendar Example ===")
	calData, err := stockbasic.TradeCal(context.Background(), client, &stockbasic.TradeCalRequest{})
	if err != nil {
		log.Fatalf("Failed to get trading calendar: %v", err)
	}

	fmt.Printf("Found %d calendar days\n", len(calData))
	if len(calData) > 0 {
		fmt.Println("\nFirst 5 days:")
		for i, d := range calData {
			if i >= 5 {
				break
			}
			fmt.Printf("  Data: %+v\n", d)
		}
	}
}
