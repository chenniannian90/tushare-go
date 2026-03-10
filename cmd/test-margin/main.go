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
	fmt.Println("   Stock Margin APIs 测试")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: margin - 融资融券交易汇总
	fmt.Println("测试 1: margin - 融资融券交易汇总")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.MarginRequest{
		TradeDate: "20240105",
	}

	items1, err := stock_margin.Margin(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items1))
		if len(items1) > 0 {
			for i, item := range items1 {
				fmt.Printf("  [%d] %s: 融资余额=%.2f亿, 融券余额=%.2f亿, 总计=%.2f亿\n",
					i+1, item.ExchangeId,
					item.Rzye/100000000, item.Rqye/100000000, item.Rzrqye/100000000)
			}
		}
	}
	fmt.Println()

	// 测试 2: margin_detail - 融资融券交易明细
	fmt.Println("测试 2: margin_detail - 融资融券交易明细")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.MarginDetailRequest{
		TsCode:    "600000.SH",
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items2, err := stock_margin.MarginDetail(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items2))
		if len(items2) > 0 {
			jsonData, _ := json.MarshalIndent(items2[0], "", "  ")
			fmt.Printf("数据示例:\n%s\n", string(jsonData))

			// 显示前3条记录的摘要
			fmt.Println("\n前3条记录摘要:")
			for i := 0; i < len(items2) && i < 3; i++ {
				item := items2[i]
				fmt.Printf("  [%d] %s: 融资余额=%.2f亿, 融券余额=%.2f万, 融券余量=%.0f股\n",
					i+1, item.TradeDate,
					item.Rzye/100000000,
					item.Rqye/10000,
					item.Rqyl)
			}
		}
	}
	fmt.Println()

	// 测试 3: margin_detail - 按日期查询多只股票
	fmt.Println("测试 3: margin_detail - 查询指定日期的所有数据")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.MarginDetailRequest{
		TradeDate: "20240105",
	}

	items3, err := stock_margin.MarginDetail(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条数据\n", len(items3))
		if len(items3) > 0 {
			// 统计前10条
			fmt.Println("\n前10只股票融资融券情况:")
			for i := 0; i < len(items3) && i < 10; i++ {
				item := items3[i]
				fmt.Printf("  [%d] %s: 融资=%.2f亿, 融券=%.2f万\n",
					i+1, item.TsCode,
					item.Rzye/100000000,
					item.Rqye/10000)
			}
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
}
