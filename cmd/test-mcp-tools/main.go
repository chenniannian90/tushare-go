package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// 导入所有工具包
	bondapi "tushare-go/pkg/sdk/api/bond"
	etfapi "tushare-go/pkg/sdk/api/etf"
	fundapi "tushare-go/pkg/sdk/api/fund"
	hk_stockapi "tushare-go/pkg/sdk/api/hk_stock"
	indexapi "tushare-go/pkg/sdk/api/index"
	stock_basicapi "tushare-go/pkg/sdk/api/stock_basic"
	stock_boardapi "tushare-go/pkg/sdk/api/stock_board"
	stock_featureapi "tushare-go/pkg/sdk/api/stock_feature"
	stock_financialapi "tushare-go/pkg/sdk/api/stock_financial"
	stock_marketapi "tushare-go/pkg/sdk/api/stock_market"
	us_stockapi "tushare-go/pkg/sdk/api/us_stock"

	"tushare-go/pkg/sdk"
)

type TestResult struct {
	Category string
	APIName  string
	Success  bool
	HasData  bool
	Count    int
	Error    string
	Duration time.Duration
}

func main() {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("❌ 请设置 TUSHARE_TOKEN 环境变量")
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

	results := make([]TestResult, 0)

	fmt.Println("🚀 开始测试 Tushare MCP 工具...")
	fmt.Println("=" + string(make([]byte, 60)))

	// 测试各个类别的工具
	fmt.Println("\n📊 测试 Stock Basic 工具...")
	results = append(results, testStockBasicTools(ctx, client)...)

	fmt.Println("\n📈 测试 Stock Market 工具...")
	results = append(results, testStockMarketTools(ctx, client)...)

	fmt.Println("\n💰 测试 Stock Financial 工具...")
	results = append(results, testStockFinancialTools(ctx, client)...)

	fmt.Println("\n🎯 测试 Stock Board 工具...")
	results = append(results, testStockBoardTools(ctx, client)...)

	fmt.Println("\n✨ 测试 Stock Feature 工具...")
	results = append(results, testStockFeatureTools(ctx, client)...)

	fmt.Println("\n💵 测试 Bond 工具...")
	results = append(results, testBondTools(ctx, client)...)

	fmt.Println("\n🏛️  测试 Fund 工具...")
	results = append(results, testFundTools(ctx, client)...)

	fmt.Println("\n📊 测试 Index 工具...")
	results = append(results, testIndexTools(ctx, client)...)

	fmt.Println("\n🌏 测试 HK Stock 工具...")
	results = append(results, testHKStockTools(ctx, client)...)

	fmt.Println("\n🇺🇸 测试 US Stock 工具...")
	results = append(results, testUSStockTools(ctx, client)...)

	fmt.Println("\n🏦 测试 ETF 工具...")
	results = append(results, testETFTools(ctx, client)...)

	// 生成报告
	generateReport(results)
}

// Stock Basic 工具测试
func testStockBasicTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"stock_basic", testStockBasic},
		{"new_share", testNewShare},
		{"namechange", testNamechange},
		{"trade_cal", testTradeCal},
		{"stock_company", testStockCompany},
		{"stk_managers", testStkManagers},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Stock Basic"
		results = append(results, result)
	}

	return results
}

// Stock Market 工具测试
func testStockMarketTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"daily", testDaily},
		{"daily_basic", testDailyBasic},
		{"weekly", testWeekly},
		{"monthly", testMonthly},
		{"adj_factor", testAdjFactor},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Stock Market"
		results = append(results, result)
	}

	return results
}

// Stock Financial 工具测试
func testStockFinancialTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"income", testIncome},
		{"balancesheet", testBalanceSheet},
		{"cashflow", testCashFlow},
		{"fina_indicator", testFinaIndicator},
		{"dividend", testDividend},
		{"forecast", testForecast},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Stock Financial"
		results = append(results, result)
	}

	return results
}

// Stock Board 工具测试
func testStockBoardTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"stk_auction", testStkAuction},
		{"hm_detail", testHmDetail},
		{"top_list", testTopList},
		{"limit_list_d", testLimitListD},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Stock Board"
		results = append(results, result)
	}

	return results
}

// Stock Feature 工具测试
func testStockFeatureTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"hk_hold", testHKHold},
		{"stk_auction_c", testStkAuctionC},
		{"report_rc", testReportRC},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Stock Feature"
		results = append(results, result)
	}

	return results
}

// Bond 工具测试
func testBondTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"cb_basic", testCbBasic},
		{"cb_daily", testCbDaily},
		{"bond_oc", testBondOC},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Bond"
		results = append(results, result)
	}

	return results
}

// Fund 工具测试
func testFundTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"fund_basic", testFundBasic},
		{"fund_nav", testFundNav},
		{"fund_manager", testFundManager},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Fund"
		results = append(results, result)
	}

	return results
}

// Index 工具测试
func testIndexTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"index_basic", testIndexBasic},
		{"index_daily", testIndexDaily},
		{"index_member_all", testIndexMemberAll},
	}

	for _, tt := range tests {
		fmt.Printf("  测��� %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "Index"
		results = append(results, result)
	}

	return results
}

// HK Stock 工具测试
func testHKStockTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"hk_basic", testHKBasic},
		{"hk_daily", testHKDaily},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "HK Stock"
		results = append(results, result)
	}

	return results
}

// US Stock 工具测试
func testUSStockTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"us_basic", testUSBasic},
		{"us_daily", testUSDaily},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "US Stock"
		results = append(results, result)
	}

	return results
}

// ETF 工具测试
func testETFTools(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	tests := []struct {
		name string
		test func(context.Context, *sdk.Client) TestResult
	}{
		{"etf_basic", testETFBasic},
		{"fund_daily", testFundDaily},
	}

	for _, tt := range tests {
		fmt.Printf("  测试 %s...\n", tt.name)
		result := tt.test(ctx, client)
		result.Category = "ETF"
		results = append(results, result)
	}

	return results
}

// ============ 具体的测试函数 ============

// Stock Basic
func testStockBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.StockBasicRequest{
		TsCode:     "000001.SZ",
		ListStatus: "L",
	}
	data, err := stock_basicapi.StockBasic(ctx, client, req)
	return buildResult("stock_basic", data, err, time.Since(start))
}

func testNewShare(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.NewShareRequest{
		StartDate: "20240101",
		EndDate:   "20240310",
	}
	data, err := stock_basicapi.NewShare(ctx, client, req)
	return buildResult("new_share", data, err, time.Since(start))
}

func testNamechange(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.NamechangeRequest{
		TsCode: "000001.SZ",
	}
	data, err := stock_basicapi.Namechange(ctx, client, req)
	return buildResult("namechange", data, err, time.Since(start))
}

func testTradeCal(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.TradeCalRequest{
		Exchange: "SSE",
		StartDate: "20240301",
		EndDate:   "20240310",
	}
	data, err := stock_basicapi.TradeCal(ctx, client, req)
	return buildResult("trade_cal", data, err, time.Since(start))
}

func testStockCompany(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.StockCompanyRequest{}
	data, err := stock_basicapi.StockCompany(ctx, client, req)
	return buildResult("stock_company", data, err, time.Since(start))
}

func testStkManagers(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_basicapi.StkManagersRequest{
		TsCode: "000001.SZ",
	}
	data, err := stock_basicapi.StkManagers(ctx, client, req)
	return buildResult("stk_managers", data, err, time.Since(start))
}

// Stock Market
func testDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_marketapi.DailyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := stock_marketapi.Daily(ctx, client, req)
	return buildResult("daily", data, err, time.Since(start))
}

func testDailyBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_marketapi.DailyBasicRequest{
		TsCode:    "000001.SZ",
		TradeDate: "20240305",
	}
	data, err := stock_marketapi.DailyBasic(ctx, client, req)
	return buildResult("daily_basic", data, err, time.Since(start))
}

func testWeekly(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_marketapi.WeeklyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240201",
		EndDate:   "20240305",
	}
	data, err := stock_marketapi.Weekly(ctx, client, req)
	return buildResult("weekly", data, err, time.Since(start))
}

func testMonthly(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_marketapi.MonthlyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240101",
		EndDate:   "20240305",
	}
	data, err := stock_marketapi.Monthly(ctx, client, req)
	return buildResult("monthly", data, err, time.Since(start))
}

func testAdjFactor(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_marketapi.AdjFactorRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := stock_marketapi.AdjFactor(ctx, client, req)
	return buildResult("adj_factor", data, err, time.Since(start))
}

// Stock Financial
func testIncome(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.IncomeRequest{
		TsCode: "000001.SZ",
		Period: "20231231",
	}
	data, err := stock_financialapi.Income(ctx, client, req)
	return buildResult("income", data, err, time.Since(start))
}

func testBalanceSheet(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.BalancesheetRequest{
		TsCode: "000001.SZ",
		Period: "20231231",
	}
	data, err := stock_financialapi.Balancesheet(ctx, client, req)
	return buildResult("balancesheet", data, err, time.Since(start))
}

func testCashFlow(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.CashflowRequest{
		TsCode: "000001.SZ",
		Period: "20231231",
	}
	data, err := stock_financialapi.Cashflow(ctx, client, req)
	return buildResult("cashflow", data, err, time.Since(start))
}

func testFinaIndicator(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.FinaIndicatorRequest{
		TsCode: "000001.SZ",
		StartDate: "20230101",
		EndDate:   "20231231",
	}
	data, err := stock_financialapi.FinaIndicator(ctx, client, req)
	return buildResult("fina_indicator", data, err, time.Since(start))
}

func testDividend(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.DividendRequest{
		TsCode: "000001.SZ",
	}
	data, err := stock_financialapi.Dividend(ctx, client, req)
	return buildResult("dividend", data, err, time.Since(start))
}

func testForecast(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_financialapi.ForecastRequest{
		StartDate: "20240101",
		EndDate:   "20240310",
	}
	data, err := stock_financialapi.Forecast(ctx, client, req)
	return buildResult("forecast", data, err, time.Since(start))
}

// Stock Board
func testStkAuction(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_boardapi.StkAuctionRequest{}
	data, err := stock_boardapi.StkAuction(ctx, client, req)
	return buildResult("stk_auction", data, err, time.Since(start))
}

func testHmDetail(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_boardapi.HmDetailRequest{
		TradeDate: "20240308",
	}
	data, err := stock_boardapi.HmDetail(ctx, client, req)
	return buildResult("hm_detail", data, err, time.Since(start))
}

func testTopList(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_boardapi.TopListRequest{
		TradeDate: "20240308",
	}
	data, err := stock_boardapi.TopList(ctx, client, req)
	return buildResult("top_list", data, err, time.Since(start))
}

func testLimitListD(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_boardapi.LimitListDRequest{
		TradeDate: "20240308",
	}
	data, err := stock_boardapi.LimitListD(ctx, client, req)
	return buildResult("limit_list_d", data, err, time.Since(start))
}

// Stock Feature
func testHKHold(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_featureapi.HkHoldRequest{
		TradeDate: "20240308",
	}
	data, err := stock_featureapi.HkHold(ctx, client, req)
	return buildResult("hk_hold", data, err, time.Since(start))
}

func testStkAuctionC(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_featureapi.StkAuctionCRequest{
		TradeDate: "20240308",
	}
	data, err := stock_featureapi.StkAuctionC(ctx, client, req)
	return buildResult("stk_auction_c", data, err, time.Since(start))
}

func testReportRC(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &stock_featureapi.ReportRcRequest{
		TsCode: "000001.SZ",
	}
	data, err := stock_featureapi.ReportRc(ctx, client, req)
	return buildResult("report_rc", data, err, time.Since(start))
}

// Bond
func testCbBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &bondapi.CbBasicRequest{
		TsCode: "113001.SZ",
	}
	data, err := bondapi.CbBasic(ctx, client, req)
	return buildResult("cb_basic", data, err, time.Since(start))
}

func testCbDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &bondapi.CbDailyRequest{
		TsCode:    "113001.SZ",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := bondapi.CbDaily(ctx, client, req)
	return buildResult("cb_daily", data, err, time.Since(start))
}

func testBondOC(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &bondapi.BcOtcqtRequest{
		TradeDate: "20240308",
	}
	data, err := bondapi.BcOtcqt(ctx, client, req)
	return buildResult("bc_otcqt", data, err, time.Since(start))
}

// Fund
func testFundBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &fundapi.FundBasicRequest{
		Market: "E",
	}
	data, err := fundapi.FundBasic(ctx, client, req)
	return buildResult("fund_basic", data, err, time.Since(start))
}

func testFundNav(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &fundapi.FundNavRequest{
		TsCode: "000001.SZ",
	}
	data, err := fundapi.FundNav(ctx, client, req)
	return buildResult("fund_nav", data, err, time.Since(start))
}

func testFundManager(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &fundapi.FundManagerRequest{
		TsCode: "000001.SZ",
	}
	data, err := fundapi.FundManager(ctx, client, req)
	return buildResult("fund_manager", data, err, time.Since(start))
}

// Index
func testIndexBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &indexapi.IndexBasicRequest{
		Market: "SSE",
	}
	data, err := indexapi.IndexBasic(ctx, client, req)
	return buildResult("index_basic", data, err, time.Since(start))
}

func testIndexDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &indexapi.IndexDailyRequest{
		TsCode:    "000001.SH",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := indexapi.IndexDaily(ctx, client, req)
	return buildResult("index_daily", data, err, time.Since(start))
}

func testIndexMemberAll(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &indexapi.IndexMemberAllRequest{
		IsNew: "Y",
	}
	data, err := indexapi.IndexMemberAll(ctx, client, req)
	return buildResult("index_member_all", data, err, time.Since(start))
}

// HK Stock
func testHKBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &hk_stockapi.HkBasicRequest{
		ListStatus: "L",
	}
	data, err := hk_stockapi.HkBasic(ctx, client, req)
	return buildResult("hk_basic", data, err, time.Since(start))
}

func testHKDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &hk_stockapi.HkDailyRequest{
		TsCode:    "00700.HK",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := hk_stockapi.HkDaily(ctx, client, req)
	return buildResult("hk_daily", data, err, time.Since(start))
}

// US Stock
func testUSBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &us_stockapi.UsBasicRequest{
		Limit: "5",
	}
	data, err := us_stockapi.UsBasic(ctx, client, req)
	return buildResult("us_basic", data, err, time.Since(start))
}

func testUSDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &us_stockapi.UsDailyRequest{
		TsCode:    "AAPL",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := us_stockapi.UsDaily(ctx, client, req)
	return buildResult("us_daily", data, err, time.Since(start))
}

// ETF
func testETFBasic(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &etfapi.EtfBasicRequest{
		Exchange: "SSE",
	}
	data, err := etfapi.EtfBasic(ctx, client, req)
	return buildResult("etf_basic", data, err, time.Since(start))
}

func testFundDaily(ctx context.Context, client *sdk.Client) TestResult {
	start := time.Now()
	req := &etfapi.FundDailyRequest{
		TsCode:    "510300.SH",
		StartDate: "20240301",
		EndDate:   "20240305",
	}
	data, err := etfapi.FundDaily(ctx, client, req)
	return buildResult("fund_daily", data, err, time.Since(start))
}

// ============ 辅助函数 ============

func buildResult(apiName string, data interface{}, err error, duration time.Duration) TestResult {
	if err != nil {
		return TestResult{
			APIName:  apiName,
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	count := 0
	hasData := false

	// 使用反射来获取数据长度
	switch v := data.(type) {
	case []stock_basicapi.StockBasicItem:
		count = len(v)
		hasData = count > 0
	case []stock_basicapi.NewShareItem:
		count = len(v)
		hasData = count > 0
	case []stock_basicapi.NamechangeItem:
		count = len(v)
		hasData = count > 0
	case []stock_basicapi.TradeCalItem:
		count = len(v)
		hasData = count > 0
	case []stock_basicapi.StockCompanyItem:
		count = len(v)
		hasData = count > 0
	case []stock_basicapi.StkManagersItem:
		count = len(v)
		hasData = count > 0
	case []stock_marketapi.DailyItem:
		count = len(v)
		hasData = count > 0
	case []stock_marketapi.DailyBasicItem:
		count = len(v)
		hasData = count > 0
	case []stock_marketapi.WeeklyItem:
		count = len(v)
		hasData = count > 0
	case []stock_marketapi.MonthlyItem:
		count = len(v)
		hasData = count > 0
	case []stock_marketapi.AdjFactorItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.IncomeItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.BalancesheetItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.CashflowItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.FinaIndicatorItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.DividendItem:
		count = len(v)
		hasData = count > 0
	case []stock_financialapi.ForecastItem:
		count = len(v)
		hasData = count > 0
	case []stock_boardapi.StkAuctionItem:
		count = len(v)
		hasData = count > 0
	case []stock_boardapi.HmDetailItem:
		count = len(v)
		hasData = count > 0
	case []stock_boardapi.TopListItem:
		count = len(v)
		hasData = count > 0
	case []stock_boardapi.LimitListDItem:
		count = len(v)
		hasData = count > 0
	case []stock_featureapi.HkHoldItem:
		count = len(v)
		hasData = count > 0
	case []stock_featureapi.StkAuctionCItem:
		count = len(v)
		hasData = count > 0
	case []stock_featureapi.ReportRcItem:
		count = len(v)
		hasData = count > 0
	case []bondapi.CbBasicItem:
		count = len(v)
		hasData = count > 0
	case []bondapi.CbDailyItem:
		count = len(v)
		hasData = count > 0
	case []bondapi.BcOtcqtItem:
		count = len(v)
		hasData = count > 0
	case []fundapi.FundBasicItem:
		count = len(v)
		hasData = count > 0
	case []fundapi.FundNavItem:
		count = len(v)
		hasData = count > 0
	case []fundapi.FundManagerItem:
		count = len(v)
		hasData = count > 0
	case []indexapi.IndexBasicItem:
		count = len(v)
		hasData = count > 0
	case []indexapi.IndexDailyItem:
		count = len(v)
		hasData = count > 0
	case []indexapi.IndexMemberAllItem:
		count = len(v)
		hasData = count > 0
	case []hk_stockapi.HkBasicItem:
		count = len(v)
		hasData = count > 0
	case []hk_stockapi.HkDailyItem:
		count = len(v)
		hasData = count > 0
	case []us_stockapi.UsBasicItem:
		count = len(v)
		hasData = count > 0
	case []us_stockapi.UsDailyItem:
		count = len(v)
		hasData = count > 0
	case []etfapi.EtfBasicItem:
		count = len(v)
		hasData = count > 0
	case []etfapi.FundDailyItem:
		count = len(v)
		hasData = count > 0
	}

	return TestResult{
		APIName:  apiName,
		Success:  true,
		HasData:  hasData,
		Count:    count,
		Duration: duration,
	}
}

func generateReport(results []TestResult) {
	fmt.Println("\n" + string(make([]byte, 70)))
	fmt.Println("                    测试报告")
	fmt.Println(string(make([]byte, 70)))

	// 按类别统计
	categoryStats := make(map[string]struct {
		total    int
		success  int
		hasData  int
		failed   int
		duration time.Duration
	})

	for _, r := range results {
		stats := categoryStats[r.Category]
		stats.total++
		stats.duration += r.Duration

		if r.Success {
			stats.success++
			if r.HasData {
				stats.hasData++
			}
		} else {
			stats.failed++
		}

		categoryStats[r.Category] = stats
	}

	// 打印类别汇总
	fmt.Printf("\n%-20s %-6s %-8s %-8s %-8s %-10s\n",
		"类别", "总计", "成功", "有数据", "失败", "耗时")
	fmt.Println(string(make([]byte, 70)))

	categories := []string{"Stock Basic", "Stock Market", "Stock Financial",
		"Stock Board", "Stock Feature", "Bond", "Fund", "Index",
		"HK Stock", "US Stock", "ETF"}

	totalTests := 0
	totalSuccess := 0
	totalHasData := 0
	totalFailed := 0

	for _, cat := range categories {
		if stats, ok := categoryStats[cat]; ok {
			fmt.Printf("%-20s %-6d %-8d %-8d %-8d %-10v\n",
				cat, stats.total, stats.success, stats.hasData,
				stats.failed, stats.duration)

			totalTests += stats.total
			totalSuccess += stats.success
			totalHasData += stats.hasData
			totalFailed += stats.failed
		}
	}

	// 打印总体统计
	fmt.Println("\n" + string(make([]byte, 70)))
	fmt.Println("总体统计")
	fmt.Println(string(make([]byte, 70)))
	fmt.Printf("总测试数: %d\n", totalTests)
	fmt.Printf("成功: %d (%.1f%%)\n", totalSuccess, float64(totalSuccess)*100/float64(totalTests))
	fmt.Printf("有数据: %d (%.1f%%)\n", totalHasData, float64(totalHasData)*100/float64(totalTests))
	fmt.Printf("失败: %d (%.1f%%)\n", totalFailed, float64(totalFailed)*100/float64(totalTests))

	// 显示失败的API
	if totalFailed > 0 || (totalSuccess-totalHasData) > 0 {
		fmt.Println("\n" + string(make([]byte, 70)))
		fmt.Println("需要关注的 API")
		fmt.Println(string(make([]byte, 70)))

		for _, r := range results {
			if !r.Success {
				fmt.Printf("❌ [%s] %s: %s\n", r.Category, r.APIName, r.Error)
			} else if !r.HasData {
				fmt.Printf("⚠️  [%s] %s: 调用成功但无数据返回\n", r.Category, r.APIName)
			}
		}
	}

	// 保存详细报告
	saveDetailedReport(results, categoryStats)

	fmt.Println("\n" + string(make([]byte, 70)))
	fmt.Println("✅ 详细报告已保存到: tests/mcp-tools-test-report.md")
	fmt.Println(string(make([]byte, 70)))
}

func saveDetailedReport(results []TestResult,
	categoryStats map[string]struct {
		total    int
		success  int
		hasData  int
		failed   int
		duration time.Duration
	}) {

	content := "# Tushare MCP 工具测试报告\n\n"
	content += "**测试时间**: " + time.Now().Format("2006-01-02 15:04:05") + "\n"
	content += "**测试目的**: 验证所有 MCP 工具的接口可用性和数据返回情况\n\n"
	content += "---\n\n"

	// 添加汇总表
	content += "## 测试结果汇总\n\n"
	content += "| 类别 | 总计 | 成功 | 有数据 | 失败 | 耗时 |\n"
	content += "|------|------|------|--------|------|------|\n"

	categories := []string{"Stock Basic", "Stock Market", "Stock Financial",
		"Stock Board", "Stock Feature", "Bond", "Fund", "Index",
		"HK Stock", "US Stock", "ETF"}

	for _, cat := range categories {
		if stats, ok := categoryStats[cat]; ok {
			content += fmt.Sprintf("| %s | %d | %d | %d | %d | %v |\n",
				cat, stats.total, stats.success, stats.hasData,
				stats.failed, stats.duration)
		}
	}

	content += "\n---\n\n"
	content += "## 详细测试结果\n\n"

	// 按类别分组显示详细结果
	for _, cat := range categories {
		categoryResults := make([]TestResult, 0)
		for _, r := range results {
			if r.Category == cat {
				categoryResults = append(categoryResults, r)
			}
		}

		if len(categoryResults) > 0 {
			content += fmt.Sprintf("### %s\n\n", cat)

			for _, r := range categoryResults {
				status := "✅ 成功"
				if !r.Success {
					status = "❌ 失败"
				} else if !r.HasData {
					status = "⚠️ 无数据"
				}

				content += fmt.Sprintf("#### %s\n\n", r.APIName)
				content += fmt.Sprintf("- **状态**: %s\n", status)
				content += fmt.Sprintf("- **耗时**: %v\n", r.Duration)

				if r.Success {
					content += fmt.Sprintf("- **返回数据量**: %d 条\n", r.Count)
					content += fmt.Sprintf("- **数据状态**: %s\n",
						map[bool]string{true: "有数据 ✅", false: "无数据 ⚠️"}[r.HasData])
				} else {
					content += fmt.Sprintf("- **错误信息**: `%s`\n", r.Error)
				}
				content += "\n"
			}
		}
	}

	content += `---

## 测试说明

- ✅ **成功**: API调用成功，服务器正常响应
- ⚠️ **无数据**: API调用成功，但未返回数据（可能测试参数无匹配数据）
- ❌ **失败**: API调用失败，出现错误

## 常见测试参数

- **股票代码**: 000001.SZ (平安银行)
- **交易日期**: 2024-03-08
- **日期范围**: 2024-03-01 至 2024-03-10
- **报告期**: 20231231 (2023年年报)
- **可转债代码**: 113001.SZ (东财转债)
- **ETF代码**: 510300.SH (沪深300ETF)
- **港股代码**: 00700.HK (腾讯控股)
- **美股代码**: AAPL (苹果)

---

**报告生成时间**: ` + time.Now().Format("2006-01-02 15:04:05") + "\n"
	content += "**测试工具**: tushare-go MCP Tools Tester\n"

	err := os.MkdirAll("tests", 0755)
	if err != nil {
		log.Printf("创建测试目录失败: %v", err)
		return
	}

	err = os.WriteFile("tests/mcp-tools-test-report.md", []byte(content), 0644)
	if err != nil {
		log.Printf("写入报告失败: %v", err)
	}
}
