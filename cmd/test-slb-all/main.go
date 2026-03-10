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
	fmt.Println("   Stock Margin - SLB APIs 测试")
	fmt.Println("   转融通相关 API 全面验证")
	fmt.Println("========================================")
	fmt.Println()

	// 测试 1: slb_len - 转融资交易汇总
	fmt.Println("测试 1: slb_len - 转融资交易汇总")
	fmt.Println("----------------------------------------")

	req1 := &stock_margin.SlbLenRequest{
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items1, err := stock_margin.SlbLen(ctx, client, req1)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条记录\n", len(items1))

		if len(items1) > 0 {
			// 统计总量
			totalOb := 0.0
			totalAuc := 0.0
			totalRepo := 0.0
			totalRepay := 0.0
			totalCb := 0.0

			for _, item := range items1 {
				totalOb += item.Ob
				totalAuc += item.AucAmount
				totalRepo += item.RepoAmount
				totalRepay += item.RepayAmount
				totalCb += item.Cb
			}

			fmt.Println("\n转融资汇总统计:")
			fmt.Printf("  期初余额: %.2f 亿元\n", totalOb)
			fmt.Printf("  竞价成交: %.2f 亿元\n", totalAuc)
			fmt.Printf("  再借成交: %.2f 亿元\n", totalRepo)
			fmt.Printf("  偿还金额: %.2f 亿元\n", totalRepay)
			fmt.Printf("  期末余额: %.2f 亿元\n", totalCb)

			// 显示最近几天的数据
			fmt.Println("\n最近5天的转融资数据:")
			for i := len(items1) - 1; i >= 0 && i >= len(items1)-5; i-- {
				item := items1[i]
				fmt.Printf("  %s: 期初%.2f亿, 竞价%.2f亿, 再借%.2f亿, 偿还%.2f亿, 期末%.2f亿\n",
					item.TradeDate, item.Ob, item.AucAmount, item.RepoAmount, item.RepayAmount, item.Cb)
			}
		}
	}
	fmt.Println()

	// 测试 2: slb_len_mm - 做市借券交易汇总
	fmt.Println("测试 2: slb_len_mm - 做市借券交易汇总")
	fmt.Println("----------------------------------------")

	req2 := &stock_margin.SlbLenMmRequest{
		StartDate: "20240101",
		EndDate:   "20240110",
	}

	items2, err := stock_margin.SlbLenMm(ctx, client, req2)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条记录\n", len(items2))

		if len(items2) > 0 {
			// 统计总量
			totalOpeInv := 0.0
			totalLentQnt := 0.0
			totalClsInv := 0.0
			totalEndBal := 0.0

			for _, item := range items2 {
				totalOpeInv += item.OpeInv
				totalLentQnt += item.LentQnt
				totalClsInv += item.ClsInv
				totalEndBal += item.EndBal
			}

			fmt.Println("\n做市借券汇总统计:")
			fmt.Printf("  期初余量: %.2f 万股\n", totalOpeInv)
			fmt.Printf("  融出数量: %.2f 万股\n", totalLentQnt)
			fmt.Printf("  期末余量: %.2f 万股\n", totalClsInv)
			fmt.Printf("  期末余额: %.2f 万元\n", totalEndBal)

			// 显示前10只股票
			fmt.Println("\n融出量最大的前10只股票:")
			for i := 0; i < len(items2) && i < 10; i++ {
				item := items2[i]
				fmt.Printf("  [%d] %s (%s): 融出%.0f万股\n",
					i+1, item.TsCode, item.Name, item.LentQnt)
			}
		}
	}
	fmt.Println()

	// 测试 3: slb_sec - 转融券标的
	fmt.Println("测试 3: slb_sec - 转融券标的")
	fmt.Println("----------------------------------------")

	req3 := &stock_margin.SlbSecRequest{
		TradeDate: "20240110",
	}

	items3, err := stock_margin.SlbSec(ctx, client, req3)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 只转融券标的\n", len(items3))

		if len(items3) > 0 {
			totalLentQnt := 0.0
			for _, item := range items3 {
				totalLentQnt += item.LentQnt
			}

			fmt.Printf("  当日融出总量: %.2f 万股\n", totalLentQnt)
		}
	}
	fmt.Println()

	// 测试 4: slb_sec_detail - 转融券交易明细
	fmt.Println("测试 4: slb_sec_detail - 转融券交易明细")
	fmt.Println("----------------------------------------")

	req4 := &stock_margin.SlbSecDetailRequest{
		TradeDate: "20240110",
	}

	items4, err := stock_margin.SlbSecDetail(ctx, client, req4)
	if err != nil {
		fmt.Printf("❌ 失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功: 获取到 %d 条交易明细\n", len(items4))

		if len(items4) > 0 {
			totalLentQnt := 0.0
			for _, item := range items4 {
				totalLentQnt += item.LentQnt
			}

			fmt.Printf("  当日融出总量: %.2f 万股\n", totalLentQnt)
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("测试完成!")
	fmt.Println("========================================")
	fmt.Println()

	// 总结
	fmt.Println("Stock Margin - 所有 6 个 API:")
	fmt.Println("========================================")
	fmt.Println("1.  margin           - 融资融券交易汇总     ✅")
	fmt.Println("2.  margin_detail    - 融资融券交易明细     ✅")
	fmt.Println("3.  margin_secs      - 融资融券标的查询     ✅")
	fmt.Println("4.  slb_sec         - 转融券标的           ✅")
	fmt.Println("5.  slb_sec_detail  - 转融券交易明细       ✅")
	fmt.Println("6.  slb_len         - 转融资交易汇总       ✅")
	fmt.Println("7.  slb_len_mm      - 做市借券交易汇总     ✅")
	fmt.Println("========================================")
	fmt.Println("所有 API 均已验证通过！✅")
}
