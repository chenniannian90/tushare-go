package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/apis"
	stockbasic "tushare-go/pkg/sdk/api/stock/stock_basic"
	stockmarket "tushare-go/pkg/sdk/api/stock/stock_market"
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
	fmt.Println("链式调用实际使用示例")
	fmt.Println("========================================")

	ctx := context.Background()

	// 示例1：获取股票基本信息
	fmt.Println("\n📊 示例1: 获取股票基本信息")
	fmt.Println("----------------------------------------")
	stocks, err := tushareClient.StockBasic(ctx, &stockbasic.StockBasicRequest{
		TsCode: "000001.SZ",
	})
	if err != nil {
		log.Printf("获取股票信息失败: %v", err)
	} else if len(stocks) > 0 {
		stock := stocks[0]
		fmt.Printf("股票代码: %s\n", stock.TsCode)
		fmt.Printf("股票名称: %s\n", stock.Name)
		fmt.Printf("所属行业: %s\n", stock.Industry)
		fmt.Printf("上市日期: %s\n", stock.ListDate)
		fmt.Println("✅ 成功获取股票基本信息")
	}

	// 示例2：获取日线数据
	fmt.Println("\n📈 示例2: 获取日线数据")
	fmt.Println("----------------------------------------")
	daily, err := tushareClient.Daily(ctx, &stockmarket.DailyRequest{
		TsCode:   "000001.SZ",
			StartDate: "20240101",
		EndDate:   "20240110",
	})
	if err != nil {
		log.Printf("获取日线数据失败: %v", err)
	} else if len(daily) > 0 {
		fmt.Printf("获取到 %d 条日线数据\n", len(daily))
		for i, item := range daily {
			if i >= 3 { // 只显示前3条
				break
			}
			fmt.Printf("  %s: 开盘=%.2f, 收盘=%.2f, 成交量=%.0f\n",
				item.TradeDate, item.Open, item.Close, item.Vol)
		}
		if len(daily) > 3 {
			fmt.Printf("  ... (还有 %d 条数据)\n", len(daily)-3)
		}
		fmt.Println("✅ 成功获取日线数据")
	}

	// 示例3：获取交易日历
	fmt.Println("\n📅 示例3: 获取交易日历")
	fmt.Println("----------------------------------------")
	calendar, err := tushareClient.TradeCal(ctx, &stockbasic.TradeCalRequest{
		Exchange: "SSE",
		StartDate: "20240101",
		EndDate:   "20240110",
	})
	if err != nil {
		log.Printf("获取交易日历失败: %v", err)
	} else if len(calendar) > 0 {
		fmt.Printf("获取到 %d 条交易日历数据\n", len(calendar))
		for i, item := range calendar {
			if i >= 5 { // 只显示前5条
				break
			}
			status := "休市"
			if item.IsOpen == 1 {
				status = "交易"
			}
			fmt.Printf("  %s: %s\n", item.CalendarDate, status)
		}
		if len(calendar) > 5 {
			fmt.Printf("  ... (还有 %d 条数据)\n", len(calendar)-5)
		}
		fmt.Println("✅ 成功获取交易日历")
	}

	fmt.Println("\n========================================")
	fmt.Println("链式调用客户端使用总结")
	fmt.Println("========================================")

	fmt.Println("\n✨ 主要优势:")
	fmt.Println("  1. 统一入口 - 一个客户端访问所有 API")
	fmt.Println("  2. 类型安全 - 完整的类型检查和 IDE 提示")
	fmt.Println("  3. 简洁易用 - 无需手动导入多个 API 包")
	fmt.Println("  4. 易于维护 - 统一的代码风格和结构")

	fmt.Println("\n📝 常用方法:")
	fmt.Println("  股票基础信息:")
	fmt.Println("    - StockBasic()      股票基本信息")
	fmt.Println("    - TradeCal()        交易日历")
	fmt.Println("    - StockCompany()    公司信息")
	fmt.Println("\n  行情数据:")
	fmt.Println("    - Daily()           日线行情")
	fmt.Println("    - Weekly()          周线行情")
	fmt.Println("    - Monthly()         月线行情")
	fmt.Println("    - StkMins()         分钟行情")
	fmt.Println("\n  其他数据:")
	fmt.Println("    - FundBasic()       基金信息")
	fmt.Println("    - IndexBasic()      指数信息")
	fmt.Println("    - TopList()         板块数据")

	fmt.Println("\n💡 更多示例:")
	fmt.Println("  查看其他示例程序了解更多用法:")
	fmt.Println("  - go run cmd/examples/daily/main.go")
	fmt.Println("  - go run cmd/examples/boards/main.go")
	fmt.Println("  - go run cmd/examples/fund/main.go")

	fmt.Println("\n📚 文档:")
	fmt.Println("  完整的 API 文档请参考项目 README.md")
}
