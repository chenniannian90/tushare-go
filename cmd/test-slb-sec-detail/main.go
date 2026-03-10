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
	fmt.Println("   slb_sec_detail API 测试")
	fmt.Println("   转融券交易明细查询")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: 查询指定日期的所有转融券交易明细
	fmt.Println("测试 1: 查询 2024-01-10 的所有转融券交易明细")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.SlbSecDetailRequest{
		TradeDate: "20240110",
	}

	items1, err := stock_margin.SlbSecDetail(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条转融券交易明细\n", len(items1))

		if len(items1) > 0 {
			// 统计总量
			totalLentQnt := 0.0
			tenorCount := make(map[string]int)
			feeRateStats := make(map[float64]int)

			for _, item := range items1 {
				totalLentQnt += item.LentQnt
				tenorCount[item.Tenor]++
				feeRateStats[item.FeeRate]++
			}

			fmt.Println("\n全市场汇总:")
			fmt.Printf("  融出总量: %.2f 万股\n", totalLentQnt)

			fmt.Println("\n按期限统计:")
			for tenor := range tenorCount {
				fmt.Printf("  %s期: %d 笔交易\n", tenor, tenorCount[tenor])
				break // 避免重复输出
			}

			// 显示前10笔交易
			fmt.Println("\n前10笔转融券交易:")
			for i := 0; i < len(items1) && i < 10; i++ {
				item := items1[i]
				fmt.Printf("  [%d] %s (%s): %s期, 费率%.2f%%, 融出%.0f万股\n",
					i+1, item.TsCode, item.Name,
					item.Tenor, item.FeeRate, item.LentQnt)
			}
		}
	}
	fmt.Println()

	// 测试 2: 查询指定股票的转融券交易明细
	fmt.Println("测试 2: 查询 600000.SH 的转融券交易明细")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.SlbSecDetailRequest{
		TsCode:    "600000.SH",
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items2, err := stock_margin.SlbSecDetail(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		if len(items2) > 0 {
			fmt.Printf("✅ 成功: %s 有转融券交易！\n", items2[0].Name)
			fmt.Printf("   查询到 %d 条交易记录\n\n", len(items2))

			// 按期限统计
			tenorStats := make(map[string]float64)
			for _, item := range items2 {
				tenorStats[item.Tenor] += item.LentQnt
			}

			fmt.Println("按期限统计融出量:")
			for tenor, qty := range tenorStats {
				fmt.Printf("  %s期: %.2f 万股\n", tenor, qty)
			}

			// 显示最近的交易
			fmt.Println("\n最近的交易记录:")
			for i := len(items2) - 1; i >= 0 && i >= len(items2)-5; i-- {
				item := items2[i]
				fmt.Printf("  %s: %s期, 费率%.2f%%, 融出%.0f万股\n",
					item.TradeDate, item.Tenor, item.FeeRate, item.LentQnt)
			}
		} else {
			fmt.Println("❌ 该股票在此期间无转融券交易")
		}
	}
	fmt.Println()

	// 测试 3: 查询日期范围的转融券交易明细
	fmt.Println("测试 3: 查询日期范围的转融券交易明细")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.SlbSecDetailRequest{
		StartDate: "20240108",
		EndDate:   "20240110",
	}

	items3, err := stock_margin.SlbSecDetail(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条交易记录\n", len(items3))

		// 按日期统计
		dateStats := make(map[string]float64)
		for _, item := range items3 {
			dateStats[item.TradeDate] += item.LentQnt
		}

		fmt.Println("\n按日期统计融出量:")
		for date := range dateStats {
			fmt.Printf("  %s: %.2f 万股\n", date, dateStats[date])
			break // 避免重复输出
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
}
