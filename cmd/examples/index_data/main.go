package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	indexapi "github.com/chenniannian90/tushare-go/pkg/sdk/api/index"
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

	// Example 1: Get index basic information
	fmt.Println("=== Example 1: Get index basic information ===")
	indexBasic, err := indexapi.IndexBasic(context.Background(), client, &indexapi.IndexBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get index basic: %v", err)
	}
	fmt.Printf("Found %d indices\n", len(indexBasic))
	fmt.Printf("Response structure: %+v\n", indexBasic)

	// Example 2: Get index daily data
	fmt.Println("\n=== Example 2: Get index daily data ===")
	indexDaily, err := indexapi.IndexDaily(context.Background(), client, &indexapi.IndexDailyRequest{})
	if err != nil {
		log.Fatalf("Failed to get index daily: %v", err)
	}
	fmt.Printf("Found %d index daily records\n", len(indexDaily))
	fmt.Printf("Response structure: %+v\n", indexDaily)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
