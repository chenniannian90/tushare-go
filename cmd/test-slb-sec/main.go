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
	fmt.Println("   slb_sec API 测试")
	fmt.Println("   转融通标的查询")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: 查询指定日期的所有转融券标的
	fmt.Println("测试 1: 查询 2024-01-10 的所有转融券标的")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.SlbSecRequest{
		TradeDate: "20240110",
	}

	items1, err := stock_margin.SlbSec(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 只转融券标的\n", len(items1))

		if len(items1) > 0 {
			// 统计总量
			totalOpeInv := 0.0
			totalLentQnt := 0.0
			totalClsInv := 0.0
			totalEndBal := 0.0

			for _, item := range items1 {
				totalOpeInv += item.OpeInv
				totalLentQnt += item.LentQnt
				totalClsInv += item.ClsInv
				totalEndBal += item.EndBal
			}

			fmt.Println("\n全市场汇总:")
			fmt.Printf("  期初余量: %.2f 万股\n", totalOpeInv)
			fmt.Printf("  融出数量: %.2f 万股\n", totalLentQnt)
			fmt.Printf("  期末余量: %.2f 万股\n", totalClsInv)
			fmt.Printf("  期末余额: %.2f 万元\n", totalEndBal)

			// 显示前10个标的
			fmt.Println("\n前10个转融券标的:")
			for i := 0; i < len(items1) && i < 10; i++ {
				item := items1[i]
				fmt.Printf("  [%d] %s (%s): 期初%.0f万股, 融出%.0f万股, 期末%.0f万股\n",
					i+1, item.TsCode, item.Name,
					item.OpeInv, item.LentQnt, item.ClsInv)
			}
		}
	}
	fmt.Println()

	// 测试 2: 查询指定股票的转融券数据
	fmt.Println("测试 2: 查询 600000.SH 的转融券数据")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.SlbSecRequest{
		TsCode:    "600000.SH",
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items2, err := stock_margin.SlbSec(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		if len(items2) > 0 {
			fmt.Printf("✅ 成功: %s 是转融券标的！\n", items2[0].Name)
			fmt.Printf("   查询到 %d 条记录\n\n", len(items2))

			// 显示最近的记录
			fmt.Println("最近的记录:")
			for i := len(items2) - 1; i >= 0 && i >= len(items2)-5; i-- {
				item := items2[i]
				fmt.Printf("  %s: 期初%.0f万股, 融出%.0f万股, 期末%.0f万股, 余额%.0f万元\n",
					item.TradeDate,
					item.OpeInv, item.LentQnt, item.ClsInv, item.EndBal)
			}
		} else {
			fmt.Println("❌ 该股票不是转融券标的或查询期间无数据")
		}
	}
	fmt.Println()

	// 测试 3: 查询日期范围的转融券数据
	fmt.Println("测试 3: 查询日期范围的转融券数据")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.SlbSecRequest{
		StartDate: "20240108",
		EndDate:   "20240110",
	}

	items3, err := stock_margin.SlbSec(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条记录\n", len(items3))

		// 按日期统计
		dateCount := make(map[string]int)
		for _, item := range items3 {
			dateCount[item.TradeDate]++
		}

		fmt.Println("\n按日期统计:")
		for date := range dateCount {
			fmt.Printf("  %s: %d 只标的\n", date, dateCount[date])
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
}
