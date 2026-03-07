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

	// Example 1: Get daily data for a specific stock
	fmt.Println("=== Example 1: Get daily data for stock ===")
	dailyData, err := api.Daily(context.Background(), client, &api.DailyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240101",
		EndDate:   "20240131",
	})
	if err != nil {
		log.Fatalf("Failed to get daily data: %v", err)
	}
	fmt.Printf("Found %d daily records\n", len(dailyData))
	if len(dailyData) > 0 {
		fmt.Println("Last 5 records:")
		for i := len(dailyData) - 1; i >= 0 && i >= len(dailyData)-5; i-- {
			d := dailyData[i]
			fmt.Printf("  %s: Open=%.2f, High=%.2f, Low=%.2f, Close=%.2f, Vol=%.0f\n",
				d.TradeDate, d.Open, d.High, d.Low, d.Close, d.Vol)
		}
	}

	// Example 2: Get daily data for specific date
	fmt.Println("\n=== Example 2: Get daily data for specific date ===")
	dateData, err := api.Daily(context.Background(), client, &api.DailyRequest{
		TradeDate: "20240115",
	})
	if err != nil {
		log.Fatalf("Failed to get daily data: %v", err)
	}
	fmt.Printf("Found %d records for 20240115\n", len(dateData))
	if len(dateData) > 0 {
		fmt.Println("First 5 stocks:")
		for i, d := range dateData {
			if i >= 5 {
				break
			}
			fmt.Printf("  %s: Close=%.2f, Volume=%.0f\n", d.TsCode, d.Close, d.Vol)
		}
	}
}
