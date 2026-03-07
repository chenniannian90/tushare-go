package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
	stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
)

func main() {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	client := sdk.NewClient(config)

	fmt.Println("========================================")
	fmt.Println("SDK 链式调用实际使用示例")
	fmt.Println("========================================")

	// 方式 1: 使用 apis 包的类型化方法（推荐）
	fmt.Println("\n✅ 方式 1: apis 包 - 类型安全，推荐使用")
	fmt.Println("\n调用示例：")
	fmt.Println("  apis.TopList(ctx, client, req)")
	fmt.Println("  apis.Daily(ctx, client, req)")
	fmt.Println("  apis.TradeCal(ctx, client, req)")

	// 实际调用
	_, err = sdkapis.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Note: Expected error (empty response_fields): %v", err)
	}

	// 方式 2: 使用 SDK 内置的链式调用 + CallAPI
	fmt.Println("\n✅ 方式 2: SDK 链式调用 - 通用方法")
	fmt.Println("\n调用示例：")
	fmt.Println("  client.StockBoard().CallAPI(ctx, \"top_list\", params, fields, &result)")
	fmt.Println("  client.StockMarket().CallAPI(ctx, \"daily\", params, fields, &result)")
	fmt.Println("  client.Index().CallAPI(ctx, \"index_basic\", params, fields, &result)")

	var result struct {
		Fields []string                 `json:"fields"`
		Items  []map[string]interface{} `json:"items"`
	}
	err = client.StockBoard().CallAPI(
		context.Background(),
		"top_list",
		map[string]interface{}{},
		[]string{},
		&result,
	)
	if err != nil {
		log.Printf("Note: Expected error (empty response_fields): %v", err)
	}

	fmt.Println("\n========================================")
	fmt.Println("对比总结")
	fmt.Println("========================================")

	fmt.Println("\n📝 方式 1 (apis 包):")
	fmt.Println("   ✅ 类型安全")
	fmt.Println("   ✅ IDE 自动提示")
	fmt.Println("   ✅ 无需手动解析结果")
	fmt.Println("   ⚠️  需要导入 apis 包")

	fmt.Println("\n📝 方式 2 (CallAPI):")
	fmt.Println("   ✅ 无需导入具体 API 包")
	fmt.Println("   ✅ 灵活调用任何 API")
	fmt.Println("   ⚠️  需要手动定义结果结构")
	fmt.Println("   ⚠️  需要手动解析字段")

	fmt.Println("\n💡 推荐:")
	fmt.Println("   日常使用：apis 包（方式 1）")
	fmt.Println("   特殊 API：CallAPI（方式 2）")

	fmt.Println("\n========================================")
	fmt.Println("可用的 API 分类")
	fmt.Println("========================================")

	fmt.Println("\n📊 所有可用分类:")
	apis := []string{
		"StockBoard", "StockMarket", "StockBasic", "StockFinancial",
		"Index", "Futures", "Fund", "HKStock", "Bond", "ETF",
		"Forex", "Options", "Spot", "USStock", "Wealth",
		"Industry", "LLMCorpus", "Macro",
	}
	for _, api := range apis {
		fmt.Printf("   • client.%s()\n", api)
	}

	fmt.Println("\n📦 apis 包提供的方法:")
	methods := []string{
		"TopList", "LimitList", "Daily", "DailyBasic",
		"TradeCal", "Income", "Balancesheet", "FinaIndicator", "Dividend",
	}
	for _, method := range methods {
		fmt.Printf("   • apis.%s()\n", method)
	}

	fmt.Println("\n🎯 快速开始:")
	fmt.Println(`
	// 1. 导入包
	import (
	    "github.com/chenniannian90/tushare-go/pkg/sdk"
	    sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
	)

	// 2. 创建客户端
	client := sdk.NewClient(config)

	// 3. 使用 apis 包调用（推荐）
	data, err := sdkapis.TopList(ctx, client, req)

	// 4. 或使用通用 CallAPI 方法
	var result struct{...}
	client.StockBoard().CallAPI(ctx, "api_name", params, fields, &result)
	`)

	fmt.Println("\n注意：当前 API spec 的 response_fields 为空，")
	fmt.Println("补充字段定义后，这些方法将返回完整的结构化数据。")
}
