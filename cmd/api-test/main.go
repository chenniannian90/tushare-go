package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	stock_basicapi "tushare-go/pkg/sdk/api/stock_basic"
	stock_marketapi "tushare-go/pkg/sdk/api/stock_market"
	stock_financialapi "tushare-go/pkg/sdk/api/stock_financial"
	bondapi "tushare-go/pkg/sdk/api/bond"
	fundapi "tushare-go/pkg/sdk/api/fund"
	indexapi "tushare-go/pkg/sdk/api/index"
	hk_stockapi "tushare-go/pkg/sdk/api/hk_stock"
	us_stockapi "tushare-go/pkg/sdk/api/us_stock"
	etfapi "tushare-go/pkg/sdk/api/etf"

	"tushare-go/pkg/sdk"
)

type TestResult struct {
	Category string
	APIName  string
	Success  bool
	HasData  bool
	Count    int
	Error    string
}

func main() {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		token = "412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
	}

	config := &sdk.Config{
		Tokens: []string{token},
		Endpoint: "https://api.tushare.pro",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	client := sdk.NewClient(config)
	ctx := context.Background()

	results := make([]TestResult, 0, 9)

	fmt.Println("🚀 开始测试 Tushare MCP API...")

	// Test Stock APIs
	fmt.Println("📈 测试 Stock APIs...")
	results = append(results, testStockBasic(ctx, client))
	results = append(results, testStockMarket(ctx, client))
	results = append(results, testStockFinancial(ctx, client))

	// Test Bond APIs
	fmt.Println("💰 测试 Bond APIs...")
	results = append(results, testBond(ctx, client))

	// Test Fund APIs
	fmt.Println("💵 测试 Fund APIs...")
	results = append(results, testFund(ctx, client))

	// Test Index APIs
	fmt.Println("📊 测试 Index APIs...")
	results = append(results, testIndex(ctx, client))

	// Test HK Stock APIs
	fmt.Println("🌏 测试 HK Stock APIs...")
	results = append(results, testHKStock(ctx, client))

	// Test US Stock APIs
	fmt.Println("🇺🇸 测试 US Stock APIs...")
	results = append(results, testUSStock(ctx, client))

	// Test ETF APIs
	fmt.Println("🏛️  测试 ETF APIs...")
	results = append(results, testETF(ctx, client))

	// Generate Report
	generateReport(results)
}

func testStockBasic(ctx context.Context, client *sdk.Client) TestResult {
	req := &stock_basicapi.StockBasicRequest{
		TsCode:     "000001.SZ",
		ListStatus: "L",
	}

	data, err := stock_basicapi.StockBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Stock",
			APIName:  "stock_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Stock",
		APIName:  "stock_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testStockMarket(ctx context.Context, client *sdk.Client) TestResult {
	req := &stock_marketapi.DailyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240301",
		EndDate:   "20240305",
	}

	data, err := stock_marketapi.Daily(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Stock",
			APIName:  "stock_market_daily",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Stock",
		APIName:  "stock_market_daily",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testStockFinancial(ctx context.Context, client *sdk.Client) TestResult {
	req := &stock_financialapi.IncomeRequest{
		TsCode: "000001.SZ",
		Period: "20231231",
	}

	data, err := stock_financialapi.Income(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Stock",
			APIName:  "stock_financial_income",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Stock",
		APIName:  "stock_financial_income",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testBond(ctx context.Context, client *sdk.Client) TestResult {
	req := &bondapi.CbBasicRequest{
		TsCode: "113001.SZ",
	}

	data, err := bondapi.CbBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Bond",
			APIName:  "bond_cb_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Bond",
		APIName:  "bond_cb_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testFund(ctx context.Context, client *sdk.Client) TestResult {
	req := &fundapi.FundBasicRequest{
		Market: "E",
	}

	data, err := fundapi.FundBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Fund",
			APIName:  "fund_fund_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Fund",
		APIName:  "fund_fund_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testIndex(ctx context.Context, client *sdk.Client) TestResult {
	req := &indexapi.IndexBasicRequest{
		Market: "SSE",
	}

	data, err := indexapi.IndexBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "Index",
			APIName:  "index_index_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "Index",
		APIName:  "index_index_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testHKStock(ctx context.Context, client *sdk.Client) TestResult {
	req := &hk_stockapi.HkBasicRequest{
		ListStatus: "L",
	}

	data, err := hk_stockapi.HkBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "HK_Stock",
			APIName:  "hk_stock_hk_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "HK_Stock",
		APIName:  "hk_stock_hk_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testUSStock(ctx context.Context, client *sdk.Client) TestResult {
	req := &us_stockapi.UsBasicRequest{
		Limit: "5",
	}

	data, err := us_stockapi.UsBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "US_Stock",
			APIName:  "us_stock_us_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "US_Stock",
		APIName:  "us_stock_us_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func testETF(ctx context.Context, client *sdk.Client) TestResult {
	req := &etfapi.EtfBasicRequest{
		Exchange: "SSE",
	}

	data, err := etfapi.EtfBasic(ctx, client, req)
	if err != nil {
		return TestResult{
			Category: "ETF",
			APIName:  "etf_etf_basic",
			Success:  false,
			Error:    err.Error(),
		}
	}

	return TestResult{
		Category: "ETF",
		APIName:  "etf_etf_basic",
		Success:  true,
		HasData:  len(data) > 0,
		Count:    len(data),
	}
}

func generateReport(results []TestResult) {
	fmt.Println("\n============================================")
	fmt.Println("        TUSHARE MCP API 测试报告")
	fmt.Println("============================================")

	// Summary by category
	categorySummary := make(map[string]int)
	categorySuccess := make(map[string]int)
	categoryNoData := make(map[string]int)

	for _, r := range results {
		categorySummary[r.Category]++
		if r.Success {
			categorySuccess[r.Category]++
			if !r.HasData {
				categoryNoData[r.Category]++
			}
		}
	}

	fmt.Printf("\n%-20s %-10s %-12s %-12s %-12s\n", "API类别", "测试接口", "成功", "有数据", "无数据")
	fmt.Println("--------------------------------------------")

	categories := []string{"Stock", "Bond", "Fund", "Index", "HK_Stock", "US_Stock", "ETF"}
	for _, cat := range categories {
		if count, ok := categorySummary[cat]; ok {
			success := categorySuccess[cat]
			noData := categoryNoData[cat]
			withData := success - noData
			fmt.Printf("%-20s %-10d %-12d %-12d %-12d\n", cat, count, success, withData, noData)
		}
	}

	fmt.Println("\n============================================")
	fmt.Println("详细测试结果")
	fmt.Println("============================================")

	for _, r := range results {
		status := "✅ 成功"
		if !r.Success {
			status = "❌ 失败"
		} else if !r.HasData {
			status = "⚠️  无数据"
		}

		fmt.Printf("\n[%s] %s.%s\n", r.Category, r.Category, r.APIName)
		fmt.Printf("  状态: %s\n", status)
		if r.Success {
			fmt.Printf("  返回数据量: %d 条\n", r.Count)
			if r.Count > 0 {
				fmt.Printf("  数据状态: %s\n", map[bool]string{true: "有数据 ✅", false: "无数据 ⚠️"}[r.HasData])
			}
		} else {
			fmt.Printf("  错误信息: %s\n", r.Error)
		}
	}

	fmt.Println("\n============================================")
	fmt.Println("测试统计")
	fmt.Println("============================================")

	totalTests := len(results)
	successTests := 0
	noDataTests := 0

	for _, r := range results {
		if r.Success {
			successTests++
			if !r.HasData {
				noDataTests++
			}
		}
	}

	fmt.Printf("总测试数: %d\n", totalTests)
	fmt.Printf("成功: %d (%.1f%%)\n", successTests, float64(successTests)*100/float64(totalTests))
	fmt.Printf("有数据: %d (%.1f%%)\n", successTests-noDataTests, float64(successTests-noDataTests)*100/float64(totalTests))
	fmt.Printf("无数据: %d (%.1f%%)\n", noDataTests, float64(noDataTests)*100/float64(totalTests))
	fmt.Printf("失败: %d (%.1f%%)\n", totalTests-successTests, float64(totalTests-successTests)*100/float64(totalTests))

	// Save report to file
	saveDetailedReport(results)

	fmt.Println("\n============================================")
	fmt.Println("✅ 详细报告已保存到: tests/mcp-api-test-report.md")
	fmt.Println("============================================")
}

func saveDetailedReport(results []TestResult) {
	content := `# Tushare MCP API 接口测试报告

**测试时间**: 2026-03-09
**测试目的**: 验证所有7个MCP服务类别的接口可用性和数据返回情况

---

## 测试结果汇总

| API类别 | 测试接口 | 成功 | 有数据 | 无数据 |
|---------|---------|------|--------|--------|
`

	// Add summary table
	categories := []string{"Stock", "Bond", "Fund", "Index", "HK_Stock", "US_Stock", "ETF"}
	categorySummary := make(map[string]int)
	categorySuccess := make(map[string]int)
	categoryNoData := make(map[string]int)

	for _, r := range results {
		categorySummary[r.Category]++
		if r.Success {
			categorySuccess[r.Category]++
			if !r.HasData {
				categoryNoData[r.Category]++
			}
		}
	}

	for _, cat := range categories {
		if count, ok := categorySummary[cat]; ok {
			success := categorySuccess[cat]
			noData := categoryNoData[cat]
			withData := success - noData
			content += fmt.Sprintf("| %s | %d | %d | %d | %d |\n", cat, count, success, withData, noData)
		}
	}

	content += `

---

## 详细测试结果

`

	for _, r := range results {
		status := "✅ 成功"
		if !r.Success {
			status = "❌ 失败"
		} else if !r.HasData {
			status = "⚠️ 无数据"
		}

		content += fmt.Sprintf("### %s.%s\n\n", r.Category, r.APIName)
		content += fmt.Sprintf("- **状态**: %s\n", status)
		if r.Success {
			content += fmt.Sprintf("- **返回数据量**: %d 条\n", r.Count)
			content += fmt.Sprintf("- **数据状态**: %s\n", map[bool]string{true: "有数据 ✅", false: "无数据 ⚠️"}[r.HasData])
		} else {
			content += fmt.Sprintf("- **错误信息**: `%s`\n", r.Error)
		}
		content += "\n"
	}

	content += `---

## 测试说明

- ✅ **成功**: API调用成功，服务器正常响应
- ⚠️ **无数据**: API调用成功，但未返回数据（可能测试参数无匹配数据）
- ❌ **失败**: API调用失败，出现错误

## 测试参数

- **股票基础数据**: ts_code=000001.SZ (平安银行)
- **股票行情**: ts_code=000001.SZ, date_range=2024-03-01 to 2024-03-05
- **股票财务**: ts_code=000001.SZ, period=20231231
- **债券**: ts_code=113001.SZ (东财转债)
- **基金**: market=E (场内基金)
- **指数**: market=SSE (上交所指数)
- **港股**: list_status=L (上市港股)
- **美股**: limit=5 (前5只美股)
- **ETF**: exchange=SSE (上交所ETF)

---

**报告生成时间**: 2026-03-09
**测试工具**: tushare-go SDK
`

	err := os.WriteFile("tests/mcp-api-test-report.md", []byte(content), 0644)
	if err != nil {
		log.Printf("Error writing report: %v", err)
	}
}
