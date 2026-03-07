package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/futures"
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

	// Example 1: Get futures basic information
	fmt.Println("=== Example 1: Get futures basic information ===")
	futBasic, err := futures.FutBasic(context.Background(), client, &futures.FutBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get futures basic: %v", err)
	}
	fmt.Printf("Found %d futures contracts\n", len(futBasic))
	fmt.Printf("Response structure: %+v\n", futBasic)

	// Example 2: Get futures daily data
	fmt.Println("\n=== Example 2: Get futures daily data ===")
	futDaily, err := futures.FutDaily(context.Background(), client, &futures.FutDailyRequest{})
	if err != nil {
		log.Fatalf("Failed to get futures daily: %v", err)
	}
	fmt.Printf("Found %d futures daily records\n", len(futDaily))
	fmt.Printf("Response structure: %+v\n", futDaily)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
