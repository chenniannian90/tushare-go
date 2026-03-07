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

	// Example: Get trading calendar
	fmt.Println("=== Trading Calendar Example ===")
	calData, err := api.TradeCal(context.Background(), client, &api.TradeCalRequest{
		Exchange:  "SSE",
		StartDate: "20240101",
		EndDate:   "20240131",
	})
	if err != nil {
		log.Fatalf("Failed to get trading calendar: %v", err)
	}

	fmt.Printf("Found %d calendar days\n", len(calData))

	// Count trading days vs non-trading days
	tradingDays := 0
	nonTradingDays := 0
	for _, d := range calData {
		if d.IsOpen == "1" {
			tradingDays++
		} else {
			nonTradingDays++
		}
	}

	fmt.Printf("Trading days: %d\n", tradingDays)
	fmt.Printf("Non-trading days: %d\n", nonTradingDays)

	// Show first 10 days
	fmt.Println("\nFirst 10 days:")
	for i, d := range calData {
		if i >= 10 {
			break
		}
		status := "Closed"
		if d.IsOpen == "1" {
			status = "Open"
		}
		fmt.Printf("  %s: %s\n", d.CalDate, status)
	}
}
