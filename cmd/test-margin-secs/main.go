package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_margin"
)

func main() {
	// 从环境变量获取 Token
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("错误: 请设置 TUSHARE_TOKEN 环境变量")
	}

	// 创建客户端
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("创建配置失败: %v", err)
	}
	client := sdk.NewClient(config)
	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("   margin_secs API 测试")
	fmt.Println("   融资融券标的查询")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: 查询指定日期的所有融资融券标的
	fmt.Println("测试 1: 查询 2024-01-10 的所有融资融券标的")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.MarginSecsRequest{
		TradeDate: "20240110",
	}

	items1, err := stock_margin.MarginSecs(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 只融资融券标的\n", len(items1))

		// 按交易所统计
		exchangeCount := make(map[string]int)
		for _, item := range items1 {
			exchangeCount[item.Exchange]++
		}
		fmt.Println("\n按交易所统计:")
		for ex, count := range exchangeCount {
			exchangeName := map[string]string{
				"SSE":  "上交所",
				"SZSE": "深交所",
				"BSE":  "北交所",
			}[ex]
			fmt.Printf("  %s (%s): %d 只\n", exchangeName, ex, count)
		}

		// 显示前10个标的
		if len(items1) > 0 {
			fmt.Println("\n前10个融资融券标的:")
			for i := 0; i < len(items1) && i < 10; i++ {
				item := items1[i]
				exchangeName := map[string]string{
					"SSE":  "上交所",
					"SZSE": "深交所",
					"BSE":  "北交所",
				}[item.Exchange]
				fmt.Printf("  [%d] %s (%s) - %s\n", i+1, item.TsCode, item.Name, exchangeName)
			}
		}
	}
	fmt.Println()

	// 测试 2: 查询指定交易所的标的
	fmt.Println("测试 2: 查询上交所 (SSE) 的融资融券标的")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.MarginSecsRequest{
		Exchange:  "SSE",
		TradeDate: "20240110",
	}

	items2, err := stock_margin.MarginSecs(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 上交所共有 %d 只融资融券标的\n", len(items2))

		if len(items2) > 0 {
			fmt.Println("\n前10个标的:")
			for i := 0; i < len(items2) && i < 10; i++ {
				item := items2[i]
				fmt.Printf("  [%d] %s - %s\n", i+1, item.TsCode, item.Name)
			}
		}
	}
	fmt.Println()

	// 测试 3: 查询指定股票是否为融资融券标的
	fmt.Println("测试 3: 查询 600000.SH 是否为融资融券标的")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.MarginSecsRequest{
		TsCode:    "600000.SH",
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items3, err := stock_margin.MarginSecs(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		if len(items3) > 0 {
			fmt.Printf("✅ 是融资融券标的！\n")
			fmt.Printf("   名称: %s\n", items3[0].Name)
			fmt.Printf("   交易所: %s\n", items3[0].Exchange)
			fmt.Printf("   查询到 %d 条记录\n", len(items3))

			// 显示最近的记录
			if len(items3) > 0 {
				fmt.Println("\n最近的记录:")
				for i := len(items3) - 1; i >= 0 && i >= len(items3)-5; i-- {
					fmt.Printf("  %s\n", items3[i].TradeDate)
				}
			}
		} else {
			fmt.Printf("❌ 不是融资融券标的或查询期间无数据\n")
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
}
