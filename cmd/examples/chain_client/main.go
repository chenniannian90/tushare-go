package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/apis"
	stockbasic "tushare-go/pkg/sdk/api/stock/stock_basic"
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

	// 创建链式调用客户端
	tushareClient := apis.NewTushareClient(client)

	fmt.Println("========================================")
	fmt.Println("链式调用客户端示例")
	fmt.Println("========================================")

	ctx := context.Background()

	// 示例1：获取股票基本信息
	fmt.Println("\n✅ 示例1: 获取股票基本信息")
	fmt.Println("调用: tushareClient.StockBasic()")
	stocks, err := tushareClient.StockBasic(ctx, &stockbasic.StockBasicRequest{
		TsCode: "000001.SZ",
	})
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("获取到 %d 只股票信息\n", len(stocks))
		if len(stocks) > 0 {
			fmt.Printf("示例: %s - %s\n", stocks[0].TsCode, stocks[0].Name)
		}
	}

	// 示例2：获取板块数据
	fmt.Println("\n✅ 示例2: 获取板块数据")
	fmt.Println("调用: tushareClient.TopList()")
	topList, err := tushareClient.TopList(ctx, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("获取到 %d 条板块数据\n", len(topList))
	}

	// 示例3：获取交易日历
	fmt.Println("\n✅ 示例3: 获取交易日历")
	fmt.Println("调用: tushareClient.TradeCal()")
	calendar, err := tushareClient.TradeCal(ctx, &stockbasic.TradeCalRequest{
		Exchange:  "SSE",
		StartDate: "20240101",
		EndDate:   "20240105",
	})
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("获取到 %d 条交易日历数据\n", len(calendar))
		for i, item := range calendar {
			if i >= 3 {
				break
			}
			status := "休市"
			if item.IsOpen == 1 {
				status = "交易"
			}
			fmt.Printf("  %s: %s\n", item.CalendarDate, status)
		}
	}

	// 示例4：使用不同的 API 接口
	fmt.Println("\n✅ 示例4: 可用的 API 接口")
	fmt.Println("链式调用客户端支持以下接口：")
	fmt.Println("  - tushareClient.Stock      (股票 API)")
	fmt.Println("  - tushareClient.Bond       (债券 API)")
	fmt.Println("  - tushareClient.ETF        (ETF API)")
	fmt.Println("  - tushareClient.Fund       (基金 API)")
	fmt.Println("  - tushareClient.Futures    (期货 API)")
	fmt.Println("  - tushareClient.Forex      (外汇 API)")
	fmt.Println("  - tushareClient.HKStock    (港股 API)")
	fmt.Println("  - tushareClient.Index      (指数 API)")
	fmt.Println("  - tushareClient.Industry   (行业 API)")
	fmt.Println("  - tushareClient.LLMCorpus  (LLM 语料 API)")
	fmt.Println("  - tushareClient.Options    (期权 API)")
	fmt.Println("  - tushareClient.Spot       (现货 API)")
	fmt.Println("  - tushareClient.USStock    (美股 API)")
	fmt.Println("  - tushareClient.Wealth     (理财 API)")

	fmt.Println("\n========================================")
	fmt.Println("链式调用的优势")
	fmt.Println("========================================")

	fmt.Println("\n📦 统一的 API 接口:")
	fmt.Println("   • 所有 API 都通过一个客户端访问")
	fmt.Println("   • 类型安全，IDE 自动提示")
	fmt.Println("   • 无需手动解析结果")
	fmt.Println("   • 更少的 import 语句")

	fmt.Println("\n🎯 代码对比:")
	fmt.Println("\n【传统方式】需要导入多个包：")
	fmt.Println("  import (")
	fmt.Println("    stockbasic \".../pkg/sdk/api/stock/stock_basic\"")
	fmt.Println("    stockmarket \".../pkg/sdk/api/stock/stock_market\"")
	fmt.Println("    stockboard \".../pkg/sdk/api/stock/stock_board\"")
	fmt.Println("  )")
	fmt.Println("  stocks, _ := stockbasic.StockBasic(ctx, client, req)")
	fmt.Println("  daily, _ := stockmarket.Daily(ctx, client, req)")
	fmt.Println("  topList, _ := stockboard.TopList(ctx, client, req)")

	fmt.Println("\n【链式调用】只需导入一个包：")
	fmt.Println("  import \"tushare-go/pkg/sdk/apis\"")
	fmt.Println("  client := apis.NewTushareClient(sdkClient)")
	fmt.Println("  stocks, _ := client.StockBasic(ctx, req)")
	fmt.Println("  daily, _ := client.Daily(ctx, req)")
	fmt.Println("  topList, _ := client.TopList(ctx, req)")

	fmt.Println("\n💡 提示:")
	fmt.Println("   • 查看具体 API 的请求和响应结构，请参考 pkg/sdk/api 目录下的包")
	fmt.Println("   • 运行其他示例程序了解更多用法：")
	fmt.Println("     - go run cmd/examples/daily/main.go")
	fmt.Println("     - go run cmd/examples/stock_basic/main.go")
	fmt.Println("     - go run cmd/examples/boards/main.go")
	fmt.Println("     - go run cmd/examples/chain_usage/main.go (实际使用示例)")
}
