// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	stockfinancial "tushare-go/pkg/sdk/api/stock/stock_financial"
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

	// Example 1: Get income statement
	fmt.Println("=== Example 1: Get income statement ===")
	incomeData, err := stockfinancial.Income(context.Background(), client, &stockfinancial.IncomeRequest{})
	if err != nil {
		log.Fatalf("Failed to get income data: %v", err)
	}
	fmt.Printf("Found %d income records\n", len(incomeData))
	fmt.Printf("Response structure: %+v\n", incomeData)

	// Example 2: Get balance sheet
	fmt.Println("\n=== Example 2: Get balance sheet ===")
	bsData, err := stockfinancial.Balancesheet(context.Background(), client, &stockfinancial.BalancesheetRequest{})
	if err != nil {
		log.Fatalf("Failed to get balance sheet data: %v", err)
	}
	fmt.Printf("Found %d balance sheet records\n", len(bsData))
	fmt.Printf("Response structure: %+v\n", bsData)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
