// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/hk_stock"
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

	// Example 1: Get HK stock basic information
	fmt.Println("=== Example 1: Get HK stock basic information ===")
	hkBasic, err := hk_stock.HkBasic(context.Background(), client, &hk_stock.HkBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get HK stock basic: %v", err)
	}
	fmt.Printf("Found %d HK stocks\n", len(hkBasic))
	fmt.Printf("Response structure: %+v\n", hkBasic)

	// Example 2: Get HK stock daily data
	fmt.Println("\n=== Example 2: Get HK stock daily data ===")
	hkDaily, err := hk_stock.HkDaily(context.Background(), client, &hk_stock.HkDailyRequest{})
	if err != nil {
		log.Fatalf("Failed to get HK daily: %v", err)
	}
	fmt.Printf("Found %d HK daily records\n", len(hkDaily))
	fmt.Printf("Response structure: %+v\n", hkDaily)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
