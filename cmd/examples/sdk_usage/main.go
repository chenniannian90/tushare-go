package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	sdkapis "tushare-go/pkg/sdk/apis"
	stockboard "tushare-go/pkg/sdk/api/stock/stock_board"
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
	fmt.Println("SDK API 调用方式对比")
	fmt.Println("========================================")

	// 方式 1: 直接调用
	fmt.Println("\n✅ 方式 1: 直接调用（原有方式）")
	fmt.Println("\n调用示例：")
	fmt.Println("  import stockboard \".../pkg/sdk/api/stock/stock_board\"")
	fmt.Println("  stockboard.TopList(ctx, client, req)")

	_, err = stockboard.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Note: Expected error (empty response_fields): %v", err)
	}

	// 方式 2: apis 包
	fmt.Println("\n✅ 方式 2: apis 包（推荐）")
	fmt.Println("\n调用示例：")
	fmt.Println("  import sdkapis \".../pkg/sdk/apis\"")
	fmt.Println("  sdkapis.TopList(ctx, client, req)")

	_, err = sdkapis.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Note: Expected error (empty response_fields): %v", err)
	}

	fmt.Println("\n========================================")
	fmt.Println("两种调用方式对比")
	fmt.Println("========================================")

	fmt.Println("\n📝 方式 1: 直接调用")
	fmt.Println("   优点: 简单直接，无需额外包装")
	fmt.Println("   缺点: 需要导入具体 API 包")
	fmt.Println("   适合: 快速脚本，API 调用较少")
	fmt.Println("\n   示例:")
	fmt.Println("   import stockboard \".../pkg/sdk/api/stock/stock_board\"")
	fmt.Println("   import stockmarket \".../pkg/sdk/api/stock/stock_market\"")
	fmt.Println("   stockboard.TopList(ctx, client, req)")
	fmt.Println("   stockmarket.Daily(ctx, client, req)")

	fmt.Println("\n📦 方式 2: apis 包（推荐）")
	fmt.Println("   优点: 类型安全，IDE 提示友好，统一导入")
	fmt.Println("   缺点: 需要导入 apis 包")
	fmt.Println("   适合: 大型项目，需要类型安全")
	fmt.Println("\n   示例:")
	fmt.Println("   import sdkapis \".../pkg/sdk/apis\"")
	fmt.Println("   sdkapis.TopList(ctx, client, req)")
	fmt.Println("   sdkapis.Daily(ctx, client, req)")
	fmt.Println("   sdkapis.TradeCal(ctx, client, req)")

	fmt.Println("\n========================================")
	fmt.Println("推荐用法")
	fmt.Println("========================================")

	fmt.Println("\n🎯 推荐: apis 包")
	fmt.Println("\n在项目中统一使用 apis 包：")
	fmt.Println(`
	// 统一导入
	import sdkapis "tushare-go/pkg/sdk/apis"

	// 使用
	data, err := sdkapis.TopList(ctx, client, req)
	daily, err := sdkapis.Daily(ctx, client, req)
	cal, err := sdkapis.TradeCal(ctx, client, req)
	`)

	fmt.Println("\n💡 优势:")
	fmt.Println("   • 类型安全")
	fmt.Println("   • IDE 自动提示")
	fmt.Println("   • 无需手动解析结果")
	fmt.Println("   • 统一的 API 接口")
	fmt.Println("   • 更少的 import 语句")

	fmt.Println("\n========================================")
	fmt.Println("可用方法")
	fmt.Println("========================================")

	fmt.Println("\n📦 apis 包提供的方法:")
	methods := []string{
		"TopList", "LimitList", "DragonList", "TopInst",
		"ThsConcept", "EmHot", // 板块
		"Daily", "DailyBasic", "Weekly", "Monthly", // 市场
		"TradeCal", "StockBasicInfo", // 基础
		"Income", "Balancesheet", "FinaIndicator", "Dividend", // 财务
	}
	for _, method := range methods {
		fmt.Printf("   • sdkapis.%s()\n", method)
	}

	fmt.Println("\n🎯 快速开始:")
	fmt.Println(`
	// 1. 导入包
	import (
	    "tushare-go/pkg/sdk"
	    sdkapis "tushare-go/pkg/sdk/apis"
	)

	// 2. 创建客户端
	client := sdk.NewClient(config)

	// 3. 使用 apis 包调用（推荐）
	data, err := sdkapis.TopList(ctx, client, req)

	// 4. 处理数据
	if err != nil {
	    log.Fatal(err)
	}
	// 使用 data...
	`)

	fmt.Println("\n注意：当前 API spec 的 response_fields 为空，")
	fmt.Println("补充字段定义后，这些方法将返回完整的结构化数据。")
}
