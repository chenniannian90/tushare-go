package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	stockbasic "tushare-go/pkg/sdk/api/stock/stock_basic"
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

	// Example: Get stock basic information
	fmt.Println("=== Example: Get stock basic information ===")
	stocks, err := stockbasic.StockBasic(context.Background(), client, &stockbasic.StockBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get stocks: %v", err)
	}
	fmt.Printf("Found %d stocks\n", len(stocks))
	fmt.Printf("Response structure: %+v\n", stocks)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
	fmt.Println("\nSDK 使用成功！一旦补充了 API 字段定义，数据将自动填充到响应结构体中。")
}
