package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
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

	// Example 1: Get limit up/down stocks
	fmt.Println("=== Example 1: Get limit list ===")
	limitList, err := stockboard.LimitList(context.Background(), client, &stockboard.LimitListRequest{})
	if err != nil {
		log.Fatalf("Failed to get limit list: %v", err)
	}
	fmt.Printf("Found %d limit list entries\n", len(limitList))
	fmt.Printf("Response structure: %+v\n", limitList)

	// Example 2: Get top list (dragon and tiger list)
	fmt.Println("\n=== Example 2: Get dragon and tiger list ===")
	topList, err := stockboard.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Fatalf("Failed to get top list: %v", err)
	}
	fmt.Printf("Found %d top list entries\n", len(topList))
	fmt.Printf("Response structure: %+v\n", topList)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
