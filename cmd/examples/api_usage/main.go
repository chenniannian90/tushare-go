package main

import (
	"context"
	"fmt"
	"log"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_market"
)

func main() {
	fmt.Println("=== Tushare Go API (爬虫版本) 使用示例 ===")
	fmt.Println()

	// 创建客户端（注意：实际上不会调用HTTP接口）
	config, _ := sdk.NewConfig("")
	client := sdk.NewClient(config)
	ctx := context.Background()

	// 示例1: RealtimeQuote - 现在调用爬虫函数
	fmt.Println("1. 测试 RealtimeQuote API (内部调用爬虫)")
	quoteReq := &stock_market.RealtimeQuoteRequest{
		TsCode: "000001.SZ", // 支持tushare格式代码
		Src:    "sina",
	}
	quotes, err := stock_market.RealtimeQuote(ctx, client, quoteReq)
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 条数据\n", len(quotes))
		if len(quotes) > 0 {
			q := quotes[0]
			fmt.Printf("   股票: %s (%s)\n", q.Name, q.TsCode)
			fmt.Printf("   价格: %.2f (昨收: %.2f)\n", q.Price, q.PreClose)
			changePercent := (q.Price - q.PreClose) / q.PreClose * 100
			fmt.Printf("   涨跌: %.2f (%.2f%%)\n\n", q.Price-q.PreClose, changePercent)
		}
	}

	// 示例2: RealtimeList - 现在调用爬虫函数
	fmt.Println("2. 测试 RealtimeList API (内部调用爬虫)")
	listReq := &stock_market.RealtimeListRequest{
		Src: "sina", // 数据源：sina-新浪，dc-东方财富
	}
	list, err := stock_market.RealtimeList(ctx, client, listReq)
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 只股票\n\n", len(list))
	}

	// 示例3: RealtimeTick - 现在调用爬虫函数
	fmt.Println("3. 测试 RealtimeTick API (内部调用爬虫)")
	tickReq := &stock_market.RealtimeTickRequest{
		TsCode: "600000.SH", // 支持tushare格式代码
		Src:    "sina",
	}
	ticks, err := stock_market.RealtimeTick(ctx, client, tickReq)
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 笔成交\n", len(ticks))
		displayCount := 3
		if len(ticks) < displayCount {
			displayCount = len(ticks)
		}
		fmt.Printf("   显示最新 %d 笔成交:\n", displayCount)
		start := len(ticks) - displayCount
		if start < 0 {
			start = 0
		}
		for i := start; i < len(ticks); i++ {
			t := ticks[i]
			fmt.Printf("   %s %.2f %d手 %s\n", t.Time, t.Price, t.Volume, t.Direction)
		}
		fmt.Println()
	}

	fmt.Println("=== 重要说明 ===")
	fmt.Println("1. 上述API函数现在直接调用爬虫函数，不依赖Tushare HTTP接口")
	fmt.Println("2. 无需配置有效的Token即可使用（但仍需提供client参数以保持接口兼容）")
	fmt.Println("3. 支持tushare格式的代码（如000001.SZ）和纯数字代码（如000001）")
	fmt.Println("4. 数据直接来自新浪财经/东方财富，仅在交易时间段返回有效数据")
	fmt.Println("5. 这些函数是对底层爬虫函数的封装，提供统一的API接口")
}
