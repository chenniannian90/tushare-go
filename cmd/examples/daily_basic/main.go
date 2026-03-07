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

	// Example: Get daily basic metrics
	fmt.Println("=== Daily Basic Metrics Example ===")
	basicData, err := api.DailyBasic(context.Background(), client, &api.DailyBasicRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240101",
		EndDate:   "20240110",
	})
	if err != nil {
		log.Fatalf("Failed to get daily basic data: %v", err)
	}

	fmt.Printf("Found %d daily basic records\n", len(basicData))
	if len(basicData) > 0 {
		fmt.Println("\nDaily Basic Metrics:")
		for _, d := range basicData {
			fmt.Printf("\nDate: %s\n", d.TradeDate)
			fmt.Printf("  PE (TTM): %.2f\n", d.PeTtm)
			fmt.Printf("  PB: %.2f\n", d.Pb)
			fmt.Printf("  PS (TTM): %.2f\n", d.PsTtm)
			fmt.Printf("  Total MV: %.2f亿\n", d.TotalMv)
			fmt.Printf("  Circ MV: %.2f亿\n", d.CircMv)
			fmt.Printf("  Turnover Rate: %.2f%%\n", d.TurnoverRate)
			fmt.Printf("  Volume Ratio: %.2f\n", d.VolumeRatio)
		}
	}
}
