package main

import (
	"fmt"
	"strings"
)

// 演示程序：展示所有可用的 MCP 工具

type MCPTool struct {
	Category    string
	Name        string
	Description string
	Enabled     bool
}

func main() {
	fmt.Println("==============================================")
	fmt.Println("      Tushare MCP 工具测试演示程序")
	fmt.Println("==============================================")
	fmt.Println()

	// 所有可用的 MCP 工具
	tools := []MCPTool{
		// Stock Basic
		{Category: "Stock Basic", Name: "stock_basic", Description: "股票基本信息", Enabled: true},
		{Category: "Stock Basic", Name: "new_share", Description: "新股数据", Enabled: true},
		{Category: "Stock Basic", Name: "namechange", Description: "股票名称变更", Enabled: true},
		{Category: "Stock Basic", Name: "trade_cal", Description: "交易日历", Enabled: true},
		{Category: "Stock Basic", Name: "stock_company", Description: "上市公司信息", Enabled: true},
		{Category: "Stock Basic", Name: "stk_managers", Description: "高级管理人员", Enabled: true},

		// Stock Market
		{Category: "Stock Market", Name: "daily", Description: "日线行情", Enabled: true},
		{Category: "Stock Market", Name: "daily_basic", Description: "日线基本面", Enabled: true},
		{Category: "Stock Market", Name: "weekly", Description: "周线行情", Enabled: true},
		{Category: "Stock Market", Name: "monthly", Description: "月线行情", Enabled: true},
		{Category: "Stock Market", Name: "adj_factor", Description: "复权因子", Enabled: true},

		// Stock Financial
		{Category: "Stock Financial", Name: "income", Description: "利润表", Enabled: true},
		{Category: "Stock Financial", Name: "balancesheet", Description: "资产负债表", Enabled: true},
		{Category: "Stock Financial", Name: "cashflow", Description: "现金流量表", Enabled: true},
		{Category: "Stock Financial", Name: "fina_indicator", Description: "财务指标", Enabled: true},
		{Category: "Stock Financial", Name: "dividend", Description: "分红送股", Enabled: true},
		{Category: "Stock Financial", Name: "forecast", Description: "业绩预告", Enabled: true},

		// Stock Board
		{Category: "Stock Board", Name: "stk_auction", Description: "集合竞价", Enabled: true},
		{Category: "Stock Board", Name: "hm_detail", Description: "游资营业部明细", Enabled: true},
		{Category: "Stock Board", Name: "top_list", Description: "龙虎榜", Enabled: true},
		{Category: "Stock Board", Name: "limit_list_d", Description: "大宗交易", Enabled: true},

		// Stock Feature
		{Category: "Stock Feature", Name: "hk_hold", Description: "港股持股", Enabled: true},
		{Category: "Stock Feature", Name: "stk_auction_c", Description: "竞价详情", Enabled: false},
		{Category: "Stock Feature", Name: "report_rc", Description: "研报", Enabled: false},

		// Bond
		{Category: "Bond", Name: "cb_basic", Description: "可转债基本信息", Enabled: true},
		{Category: "Bond", Name: "cb_daily", Description: "可转债日线", Enabled: true},
		{Category: "Bond", Name: "bc_otcqt", Description: "柜台债券", Enabled: true},

		// Fund
		{Category: "Fund", Name: "fund_basic", Description: "基金基本信息", Enabled: true},
		{Category: "Fund", Name: "fund_nav", Description: "基金净值", Enabled: true},
		{Category: "Fund", Name: "fund_manager", Description: "基金经理", Enabled: true},

		// Index
		{Category: "Index", Name: "index_basic", Description: "指数基本信息", Enabled: true},
		{Category: "Index", Name: "index_daily", Description: "指数日线", Enabled: true},
		{Category: "Index", Name: "index_member_all", Description: "指数成分股", Enabled: true},

		// HK Stock
		{Category: "HK Stock", Name: "hk_basic", Description: "港股基本信息", Enabled: true},
		{Category: "HK Stock", Name: "hk_daily", Description: "港股日线", Enabled: true},

		// US Stock
		{Category: "US Stock", Name: "us_basic", Description: "美股基本信息", Enabled: true},
		{Category: "US Stock", Name: "us_daily", Description: "美股日线", Enabled: true},

		// ETF
		{Category: "ETF", Name: "etf_basic", Description: "ETF 基本信息", Enabled: true},
		{Category: "ETF", Name: "fund_daily", Description: "ETF 日线", Enabled: true},
	}

	// 按类别统计
	categoryStats := make(map[string]int)
	categoryEnabled := make(map[string]int)

	for _, tool := range tools {
		categoryStats[tool.Category]++
		if tool.Enabled {
			categoryEnabled[tool.Category]++
		}
	}

	// 显示统计信息
	fmt.Println("📊 MCP 工具统计")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-20s %-10s %-10s %-10s\n", "类别", "总工具", "已启用", "未启用")
	fmt.Println(strings.Repeat("-", 60))

	categories := []string{"Stock Basic", "Stock Market", "Stock Financial",
		"Stock Board", "Stock Feature", "Bond", "Fund", "Index",
		"HK Stock", "US Stock", "ETF"}

	totalTools := 0
	totalEnabled := 0

	for _, cat := range categories {
		if count, ok := categoryStats[cat]; ok {
			enabled := categoryEnabled[cat]
			disabled := count - enabled
			fmt.Printf("%-20s %-10d %-10d %-10d\n", cat, count, enabled, disabled)
			totalTools += count
			totalEnabled += enabled
		}
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-20s %-10d %-10d %-10d\n", "总计", totalTools, totalEnabled, totalTools-totalEnabled)
	fmt.Println()

	// 显示所有工具
	fmt.Println("📋 所有可用的 MCP 工具")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	currentCategory := ""
	for _, tool := range tools {
		if tool.Category != currentCategory {
			currentCategory = tool.Category
			fmt.Printf("\n### %s ###\n\n", currentCategory)
		}

		status := "✅ 已启用"
		if !tool.Enabled {
			status = "❌ 未启用"
		}

		fmt.Printf("  %-25s %-15s %s\n",
			fmt.Sprintf("%s.%s", getCategoryPrefix(tool.Category), tool.Name),
			status,
			tool.Description)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()
	fmt.Println("💡 使用提示:")
	fmt.Println()
	fmt.Println("1. 运行实际测试:")
	fmt.Println("   export TUSHARE_TOKEN=你的token")
	fmt.Println("   make -C cmd/test-mcp-tools test")
	fmt.Println()
	fmt.Println("2. 查看详细文档:")
	fmt.Println("   cat cmd/test-mcp-tools/README.md")
	fmt.Println()
	fmt.Println("3. 查看快速开始:")
	fmt.Println("   cat cmd/test-mcp-tools/QUICKSTART.md")
	fmt.Println()
}

func getCategoryPrefix(category string) string {
	prefixes := map[string]string{
		"Stock Basic":     "stock_basic",
		"Stock Market":    "stock_market",
		"Stock Financial": "stock_financial",
		"Stock Board":     "stock_board",
		"Stock Feature":   "stock_feature",
		"Bond":            "bond",
		"Fund":            "fund",
		"Index":           "index",
		"HK Stock":        "hk_stock",
		"US Stock":        "us_stock",
		"ETF":             "etf",
	}

	if prefix, ok := prefixes[category]; ok {
		return prefix
	}
	return "unknown"
}
