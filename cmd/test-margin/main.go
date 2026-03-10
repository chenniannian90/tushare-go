package main

import (
	"context"
	"encoding/json"
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
	fmt.Println("   Margin API 测试")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: 获取最新一天的融资融券数据
	fmt.Println("测试 1: 获取最新融资融券数据")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.MarginRequest{
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items1, err := stock_margin.Margin(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items1))
		if len(items1) > 0 {
			jsonData, _ := json.MarshalIndent(items1[0], "", "  ")
			fmt.Printf("数据示例:\n%s\n", string(jsonData))
		}
	}
	fmt.Println()

	// 测试 2: 按交易所查询
	fmt.Println("测试 2: 查询上交所(SSE)数据")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.MarginRequest{
		ExchangeId: "SSE",
		StartDate:  "20240101",
		EndDate:    "20240105",
	}

	items2, err := stock_margin.Margin(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items2))
		if len(items2) > 0 {
			jsonData, _ := json.MarshalIndent(items2[0], "", "  ")
			fmt.Printf("数据示例:\n%s\n", string(jsonData))
		}
	}
	fmt.Println()

	// 测试 3: 单日查询
	fmt.Println("测试 3: 查询单日数据")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.MarginRequest{
		TradeDate: "20240105",
	}

	items3, err := stock_margin.Margin(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items3))
		if len(items3) > 0 {
			for i, item := range items3 {
				fmt.Printf("[%d] %s: 融资余额=%.2f亿, 融券余额=%.2f亿, 融资融券余额=%.2f亿\n",
					i+1, item.ExchangeId,
					item.Rzye/100000000, item.Rqye/100000000, item.Rzrqye/100000000)
			}
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
}
