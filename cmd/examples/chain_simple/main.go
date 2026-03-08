// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/apis"
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

	// 创建链式调用客户端
	_ = apis.NewTushareClient(client)

	fmt.Println("========================================")
	fmt.Println("链式调用客户端 - 简单示例")
	fmt.Println("========================================")

	_ = context.Background()

	// 示例：演示如何使用链式调用客户端
	// 注意：实际的 API 调用需要提供正确的参数
	// 这里只是展示使用方式

	fmt.Println("\n✅ 创建链式调用客户端成功！")
	fmt.Println("\n📋 可用的 API 接口：")

	// 打印所有可用的接口类型
	fmt.Println("  股票相关:")
	fmt.Println("    - StockBasic()      股票基本信息")
	fmt.Println("    - TradeCal()        交易日历")
	fmt.Println("    - Daily()           日线行情")
	fmt.Println("    - TopList()         龙虎榜")
	fmt.Println("    - LimitList()       涨跌停")

	fmt.Println("  基金相关:")
	fmt.Println("    - FundBasic()       基金基本信息")
	fmt.Println("    - FundNav()         基金净值")

	fmt.Println("  指数相关:")
	fmt.Println("    - IndexBasic()      指数基本信息")
	fmt.Println("    - IndexDaily()      指数日线")

	fmt.Println("  其他:")
	fmt.Println("    - HkBasic()         港股基本信息")
	fmt.Println("    - UsBasic()         美股基本信息")
	fmt.Println("    - FutBasic()        期货基本信息")

	fmt.Println("\n🎯 使用示例:")
	fmt.Println(`
	// 1. 导入包
	import "tushare-go/pkg/sdk/apis"

	// 2. 创建客户端
	client := sdk.NewClient(config)
	tushareClient := apis.NewTushareClient(client)

	// 3. 使用 API（示例）
	ctx := context.Background()

	// 获取股票信息
	stocks, err := tushareClient.StockBasic(ctx, &StockBasicRequest{...})
	if err != nil {
	    log.Fatal(err)
	}

	// 获取日线数据
	daily, err := tushareClient.Daily(ctx, &DailyRequest{...})
	if err != nil {
	    log.Fatal(err)
	}
	`)

	fmt.Println("\n💡 主要优势:")
	fmt.Println("  ✅ 统一的入口点 - 一个客户端访问所有 API")
	fmt.Println("  ✅ 类型安全 - 完整的类型检查和 IDE 提示")
	fmt.Println("  ✅ 简洁易用 - 无需手动导入多个 API 包")
	fmt.Println("  ✅ 易于维护 - 统一的代码风格和结构")

	fmt.Println("\n📚 更多示例:")
	fmt.Println("  查看其他示例程序了解更多用法:")
	fmt.Println("  - go run cmd/examples/daily/main.go")
	fmt.Println("  - go run cmd/examples/stock_basic/main.go")
	fmt.Println("  - go run cmd/examples/boards/main.go")

	fmt.Println("\n📖 文档:")
	fmt.Println("  完整的 API 文档请参考项目 README.md")
	fmt.Println("  查看具体的 API 定义：pkg/sdk/apis/ 目录")

	fmt.Println("\n========================================")
	fmt.Println("✅ 链式调用客户端使用演示完成")
	fmt.Println("========================================")
}
