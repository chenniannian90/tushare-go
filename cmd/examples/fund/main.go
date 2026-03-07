package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/fund"
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

	// Example 1: Get fund basic information
	fmt.Println("=== Example 1: Get fund basic information ===")
	fundBasic, err := fund.FundBasic(context.Background(), client, &fund.FundBasicRequest{})
	if err != nil {
		log.Fatalf("Failed to get fund basic: %v", err)
	}
	fmt.Printf("Found %d funds\n", len(fundBasic))
	fmt.Printf("Response structure: %+v\n", fundBasic)

	// Example 2: Get fund net value
	fmt.Println("\n=== Example 2: Get fund net value ===")
	fundNav, err := fund.FundNav(context.Background(), client, &fund.FundNavRequest{})
	if err != nil {
		log.Fatalf("Failed to get fund nav: %v", err)
	}
	fmt.Printf("Found %d nav records\n", len(fundNav))
	fmt.Printf("Response structure: %+v\n", fundNav)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
