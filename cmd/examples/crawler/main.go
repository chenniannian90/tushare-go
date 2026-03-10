package main

import (
	"fmt"
	"log"

	"tushare-go/pkg/sdk/api/stock_market"
)

func main() {
	fmt.Println("=== Tushare Go 实时数据爬虫示例 ===")
	fmt.Println()

	// 示例1: 获取单个股票实时行情
	fmt.Println("1. 获取单个股票实时行情 (平安银行 000001)")
	quotes, err := stock_market.GetRealtimeQuotes("000001")
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 条数据\n", len(quotes))
		if len(quotes) > 0 {
			q := quotes[0]
			fmt.Printf("   股票: %s (%s)\n", q.Name, q.TsCode)
			fmt.Printf("   价格: %.2f (昨收: %.2f)\n", q.Price, q.PreClose)
			fmt.Printf("   涨跌: %.2f (%.2f%%)\n", q.Price-q.PreClose, (q.Price-q.PreClose)/q.PreClose*100)
			fmt.Printf("   买一: %d手 @ %.2f\n", q.B1V, q.B1P)
			fmt.Printf("   卖一: %d手 @ %.2f\n", q.A1V, q.A1P)
			fmt.Printf("   时间: %s %s\n\n", q.Date, q.Time)
		}
	}

	// 示例2: ��取多个股票实时行情
	fmt.Println("2. 获取多个股票实时行情")
	codes := []string{"000001", "600000", "000002"}
	multiQuotes, err := stock_market.GetRealtimeQuotes(codes)
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 条数据\n", len(multiQuotes))
		for i, q := range multiQuotes {
			fmt.Printf("   %d. %s (%.2f, %.2f%%)\n", i+1, q.TsCode, q.Price,
				(q.Price-q.PreClose)/q.PreClose*100)
		}
		fmt.Println()
	}

	// 示例3: 获取实时排名列表
	fmt.Println("3. 获取实时排名列表 (新浪数据源)")
	list, err := stock_market.GetRealtimeList("sina")
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 只股票\n", len(list))
		maxDisplay := 5
		if len(list) < maxDisplay {
			maxDisplay = len(list)
		}
		for i := 0; i < maxDisplay; i++ {
			item := list[i]
			changePercent := (item.Price - item.PreClose) / item.PreClose * 100
			fmt.Printf("   %d. %s (%.2f, %.2f%%)\n", i+1, item.TsCode, item.Price, changePercent)
		}
		fmt.Println()
	}

	// 示例4: 获取分笔成交数据
	fmt.Println("4. 获取分笔成交数据 (平安银行 000001)")
	ticks, err := stock_market.GetRealtimeTick("000001", "sina")
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 笔成交\n", len(ticks))
		displayCount := 5
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

	// 示例5: 获取当日完整分笔数据
	fmt.Println("5. 获取当日完整分笔数据 (浦发银行 600000)")
	fullTicks, err := stock_market.GetTodayTicks("600000")
	if err != nil {
		log.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 获取成功，共 %d 笔成交\n", len(fullTicks))

		// 统计买卖盘
		buyVolume := 0
		sellVolume := 0
		for _, t := range fullTicks {
			switch t.Direction {
			case "买盘":
				buyVolume += t.Volume
			case "卖盘":
				sellVolume += t.Volume
			}
		}
		totalVolume := buyVolume + sellVolume
		if totalVolume > 0 {
			fmt.Printf("   买盘: %d手 (%.1f%%), 卖盘: %d手 (%.1f%%)\n",
				buyVolume, float64(buyVolume)/float64(totalVolume)*100,
				sellVolume, float64(sellVolume)/float64(totalVolume)*100)
		}
		fmt.Println()
	}

	fmt.Println("=== 爬虫示例完成 ===")
	fmt.Println("\n注意事项:")
	fmt.Println("1. 这些函数直接爬取新浪财经/东方财富数据，不需要Tushare Token")
	fmt.Println("2. 仅在交易时间段内返回有效数据")
	fmt.Println("3. 数据源可能会变更，需要定期维护")
	fmt.Println("4. 请遵守相关网站的使用条款和robots.txt规定")
}