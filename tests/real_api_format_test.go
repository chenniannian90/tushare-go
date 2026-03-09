package tests

import (
	"encoding/json"
	"strconv"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestDailyAPI tests historical daily market data API
func TestDailyAPI(t *testing.T) {
	// Actual response from daily API (historical daily data)
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","open","high","low","close","vol","amount"],"items":[["688433.SH","20250110",22.9,23.39,22.0,22.05,29066.57,65814.656]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	// Verify numeric fields (float64)
	if items[0]["open"].(float64) != 22.9 {
		t.Errorf("Expected open 22.9, got %v", items[0]["open"])
	}

	// Verify large numbers (vol and amount)
	if items[0]["vol"].(float64) != 29066.57 {
		t.Errorf("Expected vol 29066.57, got %v", items[0]["vol"])
	}
}

// TestWeeklyAPI tests weekly market data API
func TestWeeklyAPI(t *testing.T) {
	// Actual response from weekly API
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","close","open","high","low","vol","amount"],"items":[["688433.SH","20250127",25.4,26.32,26.69,24.69,4109153.0,103973699.0],["688433.SH","20250124",26.73,25.39,26.96,24.2,13851344.0,353861016.0]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify very large numbers
	vol := items[0]["vol"].(float64)
	if vol < 4000000 || vol > 4200000 {
		t.Errorf("Volume seems wrong: %v", vol)
	}

	amount := items[0]["amount"].(float64)
	if amount < 103000000 || amount > 104000000 {
		t.Errorf("Amount seems wrong: %v", amount)
	}
}

// TestDailyBasicAPI tests daily basic indicators API
func TestDailyBasicAPI(t *testing.T) {
	// Actual response from daily_basic API
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","turnover_rate","volume_ratio","pe","pb","total_mv","circ_mv"],"items":[["601226.SH","20250110",4.9689,0.94,82.5444,1.8859,804954.0,801934.353]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify ratio fields
	if items[0]["turnover_rate"].(float64) != 4.9689 {
		t.Errorf("Expected turnover_rate 4.9689, got %v", items[0]["turnover_rate"])
	}

	// Verify PE ratio (important metric)
	if items[0]["pe"].(float64) != 82.5444 {
		t.Errorf("Expected pe 82.5444, got %v", items[0]["pe"])
	}

	// Verify market cap (large numbers)
	totalMV := items[0]["total_mv"].(float64)
	if totalMV < 804900 || totalMV > 805000 {
		t.Errorf("Total market cap seems wrong: %v", totalMV)
	}
}

// TestAdjFactorAPI tests adjustment factor API
func TestAdjFactorAPI(t *testing.T) {
	// Actual response from adj_factor API
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","adj_factor"],"items":[["601226.SH","20250110",1.6061],["601226.SH","20250109",1.6061],["601226.SH","20250108",1.6061]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Verify adjustment factor (should be float64)
	for i, item := range items {
		adjFactor := item["adj_factor"].(float64)
		if adjFactor != 1.6061 {
			t.Errorf("Item %d: Expected adj_factor 1.6061, got %v", i, adjFactor)
		}
	}
}

// TestHsgtTop10API tests top 10 HK stock connect stocks API
func TestHsgtTop10API(t *testing.T) {
	// Actual response from hsgt_top10 API with Chinese characters
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","close","change","rank","amount"],"items":[["20250110","600519.SH","GUIZHOU_MAOTAI",1436.0,-8.0,2,1150467012.0],["20250110","600584.SH","CHANGDIAN_KEJI",40.08,2.39,4,879917871.0]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify negative change (stock price dropped)
	if items[0]["change"].(float64) != -8.0 {
		t.Errorf("Expected change -8.0, got %v", items[0]["change"])
	}

	// Verify rank (integer-like)
	if items[0]["rank"].(float64) != 2 {
		t.Errorf("Expected rank 2, got %v", items[0]["rank"])
	}

	// Verify very large amounts
	amount1 := items[0]["amount"].(float64)
	if amount1 < 1150000000 || amount1 > 1151000000 {
		t.Errorf("Amount seems wrong: %v", amount1)
	}
}

// TestDecimalPrecision tests API responses with high precision decimals
func TestDecimalPrecision(t *testing.T) {
	// Test APIs that return many decimal places
	apiResponse := `{"fields":["ts_code","trade_date","pe","pb","ps","turnover_rate"],"items":[["000001.SZ","20250110",15.234567,1.234567,3.456789,2.345678],["000002.SZ","20250110",22.987654,0.987654,2.345678,4.567890]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify decimal precision is preserved
	if items[0]["pe"].(float64) != 15.234567 {
		t.Errorf("Expected pe 15.234567, got %v", items[0]["pe"])
	}

	// Verify pb with 6 decimal places
	if items[0]["pb"].(float64) != 1.234567 {
		t.Errorf("Expected pb 1.234567, got %v", items[0]["pb"])
	}
}

// TestLargeVolumeData tests APIs with trading volume data
func TestLargeVolumeData(t *testing.T) {
	// Test APIs that return trading volume and amount
	apiResponse := `{"fields":["ts_code","trade_date","vol","amount","buy","sell"],"items":[["600519.SH","20250110",1000000,1500000000,800000,700000],["600584.SH","20250110",2000000,2500000000,1200000,1300000]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify volume (usually large integers)
	if items[0]["vol"].(float64) != 1000000 {
		t.Errorf("Expected vol 1000000, got %v", items[0]["vol"])
	}

	// Verify amount (very large numbers)
	if items[0]["amount"].(float64) != 1500000000 {
		t.Errorf("Expected amount 1500000000, got %v", items[0]["amount"])
	}
}

// TestMultipleDataTypes tests response with various data types
func TestMultipleDataTypes(t *testing.T) {
	// Test response with integer-like, float, and negative values
	apiResponse := `{"fields":["ts_code","trade_date","close","change","pct_chg","vol","rank"],"items":[["000001.SZ","20250110",10.5,-0.5,-4.5,1000000,1],["000002.SZ","20250110",8.7,0.3,3.5,800000,2]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify negative change
	if items[0]["change"].(float64) != -0.5 {
		t.Errorf("Expected change -0.5, got %v", items[0]["change"])
	}

	// Verify negative percentage
	if items[0]["pct_chg"].(float64) != -4.5 {
		t.Errorf("Expected pct_chg -4.5, got %v", items[0]["pct_chg"])
	}

	// Verify positive change
	if items[1]["change"].(float64) != 0.3 {
		t.Errorf("Expected change 0.3, got %v", items[1]["change"])
	}

	// Verify rank (1-based)
	if items[0]["rank"].(float64) != 1 {
		t.Errorf("Expected rank 1, got %v", items[0]["rank"])
	}
}

// TestIncomeAPI tests financial income statement API
func TestIncomeAPI(t *testing.T) {
	// Income statement data with many financial metrics
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","f_ann_date","end_date","report_type","basic_eps","diluted_eps","total_revenue","revenue","operate_profit","total_profit","n_income","n_income_attr_p"],"items":[["601226.SH","20240130","20240129","20231231","年报",1.23,1.21,1000000000,950000000,200000000,180000000,150000000,145000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify EPS (earnings per share)
	if items[0]["basic_eps"].(float64) != 1.23 {
		t.Errorf("Expected basic_eps 1.23, got %v", items[0]["basic_eps"])
	}

	// Verify very large revenue numbers
	totalRevenue := items[0]["total_revenue"].(float64)
	if totalRevenue != 1000000000 {
		t.Errorf("Expected total_revenue 1000000000, got %v", totalRevenue)
	}

	// Verify profit fields
	if items[0]["n_income_attr_p"].(float64) != 145000000 {
		t.Errorf("Expected n_income_attr_p 145000000, got %v", items[0]["n_income_attr_p"])
	}
}

// TestBalanceSheetAPI tests balance sheet API
func TestBalanceSheetAPI(t *testing.T) {
	// Balance sheet data with asset and liability fields
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","total_share","cap_rese","undistr_porfit","total_assets","total_cur_assets","total_nca","total_liab","total_cur_liab","total_ncl","total_hldr_eqy_exc_min_int"],"items":[["601226.SH","20240430","20240331",5000000000,1000000000,2000000000,50000000000,30000000000,20000000000,25000000000,15000000000,10000000000,25000000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify equity fields
	if items[0]["total_share"].(float64) != 5000000000 {
		t.Errorf("Expected total_share 5000000000, got %v", items[0]["total_share"])
	}

	// Verify total assets
	if items[0]["total_assets"].(float64) != 50000000000 {
		t.Errorf("Expected total_assets 50000000000, got %v", items[0]["total_assets"])
	}

	// Verify balance sheet equation: assets = liabilities + equity
	totalAssets := items[0]["total_assets"].(float64)
	totalLiab := items[0]["total_liab"].(float64)
	totalEquity := items[0]["total_hldr_eqy_exc_min_int"].(float64)

	if totalLiab+totalEquity != totalAssets {
		t.Errorf("Balance sheet equation failed: %v + %v != %v", totalLiab, totalEquity, totalAssets)
	}
}

// TestCashflowAPI tests cash flow statement API
func TestCashflowAPI(t *testing.T) {
	// Cash flow data with operating, investing, and financing activities
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","net_profit","n_cashflow_act","n_cashflow_inv_act","n_cash_flows_fnc_act","c_cash_equ_beg_period","c_cash_equ_end_period"],"items":[["601226.SH","20240430","20240331",1000000000,800000000,500000000,200000000,500000000,2000000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net profit
	if items[0]["net_profit"].(float64) != 1000000000 {
		t.Errorf("Expected net_profit 1000000000, got %v", items[0]["net_profit"])
	}

	// Verify operating cash flow
	if items[0]["n_cashflow_act"].(float64) != 800000000 {
		t.Errorf("Expected n_cashflow_act 800000000, got %v", items[0]["n_cashflow_act"])
	}

	// Verify cash balance change
	begBalance := items[0]["c_cash_equ_beg_period"].(float64)
	endBalance := items[0]["c_cash_equ_end_period"].(float64)
	if endBalance <= begBalance {
		t.Errorf("Cash balance should increase: %v -> %v", begBalance, endBalance)
	}
}

// TestForecastAPI tests performance forecast API
func TestForecastAPI(t *testing.T) {
	// Forecast data with prediction ranges
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","type","p_change_min","p_change_max","net_profit_min","net_profit_max","change_reason"],"items":[["601226.SH","20240115","20231231","预增",50.0,80.0,1000000000,1200000000,"市场需求增加"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify forecast type
	if items[0]["type"].(string) != "预增" {
		t.Errorf("Expected type '预增', got %v", items[0]["type"])
	}

	// Verify percentage change range
	if items[0]["p_change_min"].(float64) != 50.0 {
		t.Errorf("Expected p_change_min 50.0, got %v", items[0]["p_change_min"])
	}

	// Verify profit prediction range
	minProfit := items[0]["net_profit_min"].(float64)
	maxProfit := items[0]["net_profit_max"].(float64)
	if minProfit >= maxProfit {
		t.Errorf("Min profit should be less than max profit: %v >= %v", minProfit, maxProfit)
	}

	// Verify Chinese characters in change reason
	if items[0]["change_reason"].(string) != "市场需求增加" {
		t.Errorf("Expected change_reason '市场需求增加', got %v", items[0]["change_reason"])
	}
}

// TestDividendAPI tests dividend data API
func TestDividendAPI(t *testing.T) {
	// Dividend data with various dividend types
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","end_date","ann_date","div_proc","stk_div","cash_div","record_date","ex_date","pay_date"],"items":[["601226.SH","20231231","20240415","实施",10,5.5,"20240420","20240421","20240425"],["600000.SH","20231231","20240410","实施",0,2.0,"20240415","20240416","20240420"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify stock dividend (first item has both stock and cash dividend)
	if items[0]["stk_div"].(float64) != 10 {
		t.Errorf("Expected stk_div 10, got %v", items[0]["stk_div"])
	}

	// Verify cash dividend
	if items[0]["cash_div"].(float64) != 5.5 {
		t.Errorf("Expected cash_div 5.5, got %v", items[0]["cash_div"])
	}

	// Verify second item only has cash dividend
	if items[1]["stk_div"].(float64) != 0 {
		t.Errorf("Expected stk_div 0, got %v", items[1]["stk_div"])
	}
}

// TestDisclosureDateAPI tests disclosure date API
func TestDisclosureDateAPI(t *testing.T) {
	// Disclosure date scheduling data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","pre_date","actual_date"],"items":[["601226.SH","20240110","20231231","20240331","20240428"],["600000.SH","20240105","20231231","20240331","20240425"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify date fields (all should be strings in YYYYMMDD format)
	if items[0]["end_date"].(string) != "20231231" {
		t.Errorf("Expected end_date '20231231', got %v", items[0]["end_date"])
	}

	// Verify planned disclosure date
	if items[0]["pre_date"].(string) != "20240331" {
		t.Errorf("Expected pre_date '20240331', got %v", items[0]["pre_date"])
	}

	// Verify actual disclosure date
	if items[0]["actual_date"].(string) != "20240428" {
		t.Errorf("Expected actual_date '20240428', got %v", items[0]["actual_date"])
	}
}

// TestFinaIndicatorAPI tests financial indicators API
func TestFinaIndicatorAPI(t *testing.T) {
	// Financial indicators with many ratios
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","eps","dt_eps","roe","roa","current_ratio","quick_ratio","gross_margin","netprofit_margin","debt_to_assets"],"items":[["601226.SH","20240430","20240331",1.5,1.45,15.5,8.2,1.5,1.2,35.5,12.3,45.5]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify EPS
	if items[0]["eps"].(float64) != 1.5 {
		t.Errorf("Expected eps 1.5, got %v", items[0]["eps"])
	}

	// Verify ROE (return on equity)
	if items[0]["roe"].(float64) != 15.5 {
		t.Errorf("Expected roe 15.5, got %v", items[0]["roe"])
	}

	// Verify ROA (return on assets)
	if items[0]["roa"].(float64) != 8.2 {
		t.Errorf("Expected roa 8.2, got %v", items[0]["roa"])
	}

	// Verify liquidity ratios
	if items[0]["current_ratio"].(float64) != 1.5 {
		t.Errorf("Expected current_ratio 1.5, got %v", items[0]["current_ratio"])
	}

	// Verify profitability ratios
	if items[0]["gross_margin"].(float64) != 35.5 {
		t.Errorf("Expected gross_margin 35.5, got %v", items[0]["gross_margin"])
	}
}

// TestTop10FloatholdersAPI tests top 10 floating shareholders API
func TestTop10FloatholdersAPI(t *testing.T) {
	// Top 10 floating shareholders data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","holder_name","hold_amount","hold_ratio","hold_float_ratio","hold_change","holder_type"],"items":[["601226.SH","20240430","20240331","中国工商银行股份有限公司",1000000000,20.5,25.3,5000000,"其他"],["601226.SH","20240430","20240331","香港中央结算有限公司",800000000,16.4,20.2,3000000,"其他"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify holder name (Chinese characters)
	if items[0]["holder_name"].(string) != "中国工商银行股份有限公司" {
		t.Errorf("Expected holder_name '中国工商银行股份有限公司', got %v", items[0]["holder_name"])
	}

	// Verify holding amount (large integer)
	if items[0]["hold_amount"].(float64) != 1000000000 {
		t.Errorf("Expected hold_amount 1000000000, got %v", items[0]["hold_amount"])
	}

	// Verify holding ratio (percentage)
	if items[0]["hold_ratio"].(float64) != 20.5 {
		t.Errorf("Expected hold_ratio 20.5, got %v", items[0]["hold_ratio"])
	}

	// Verify holder type
	if items[0]["holder_type"].(string) != "其他" {
		t.Errorf("Expected holder_type '其他', got %v", items[0]["holder_type"])
	}
}

// TestPledgeStatAPI tests stock pledge statistics API
func TestPledgeStatAPI(t *testing.T) {
	// Stock pledge statistics data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","end_date","pledge_count","unrest_pledge","rest_pledge","total_share","pledge_ratio"],"items":[["601226.SH","20240331",10,500000000,300000000,5000000000,16.0]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify pledge count
	if items[0]["pledge_count"].(float64) != 10 {
		t.Errorf("Expected pledge_count 10, got %v", items[0]["pledge_count"])
	}

	// Verify unrestricted pledge amount
	if items[0]["unrest_pledge"].(float64) != 500000000 {
		t.Errorf("Expected unrest_pledge 500000000, got %v", items[0]["unrest_pledge"])
	}

	// Verify restricted pledge amount
	if items[0]["rest_pledge"].(float64) != 300000000 {
		t.Errorf("Expected rest_pledge 300000000, got %v", items[0]["rest_pledge"])
	}

	// Verify total shares
	if items[0]["total_share"].(float64) != 5000000000 {
		t.Errorf("Expected total_share 5000000000, got %v", items[0]["total_share"])
	}

	// Verify pledge ratio calculation
	pledgeAmount := items[0]["unrest_pledge"].(float64) + items[0]["rest_pledge"].(float64)
	totalShare := items[0]["total_share"].(float64)
	expectedRatio := (pledgeAmount / totalShare) * 100
	actualRatio := items[0]["pledge_ratio"].(float64)
	if actualRatio != expectedRatio {
		t.Errorf("Pledge ratio mismatch: expected %.1f, got %.1f", expectedRatio, actualRatio)
	}
}

// TestPledgeDetailAPI tests stock pledge details API
func TestPledgeDetailAPI(t *testing.T) {
	// Stock pledge details data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","holder_name","pledge_amount","start_date","end_date","is_release","release_date","pledgor","holding_amount","pledged_amount","p_total_ratio","h_total_ratio","is_buyback"],"items":[["601226.SH","20240115","张三",100000000,"20240110","20250110","否","","张三",500000000,300000000,20.0,60.0,"否"],["601226.SH","20240120","李四",50000000,"20240115","20250115","是","20240201","李四",300000000,150000000,16.67,50.0,"否"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify first item (not released)
	if items[0]["is_release"].(string) != "否" {
		t.Errorf("Expected is_release '否', got %v", items[0]["is_release"])
	}

	// Verify pledge amount
	if items[0]["pledge_amount"].(float64) != 100000000 {
		t.Errorf("Expected pledge_amount 100000000, got %v", items[0]["pledge_amount"])
	}

	// Verify second item (released)
	if items[1]["is_release"].(string) != "是" {
		t.Errorf("Expected is_release '是', got %v", items[1]["is_release"])
	}

	// Verify release date for second item
	if items[1]["release_date"].(string) != "20240201" {
		t.Errorf("Expected release_date '20240201', got %v", items[1]["release_date"])
	}
}

// TestRepurchaseAPI tests stock repurchase data API
func TestRepurchaseAPI(t *testing.T) {
	// Stock repurchase data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","proc","exp_date","vol","amount","high_limit","low_limit"],"items":[["601226.SH","20240110","20240109","实施","20240210",10000000,500000000,15.5,12.3]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify procedure status
	if items[0]["proc"].(string) != "实施" {
		t.Errorf("Expected proc '实施', got %v", items[0]["proc"])
	}

	// Verify repurchase volume
	if items[0]["vol"].(float64) != 10000000 {
		t.Errorf("Expected vol 10000000, got %v", items[0]["vol"])
	}

	// Verify repurchase amount
	if items[0]["amount"].(float64) != 500000000 {
		t.Errorf("Expected amount 500000000, got %v", items[0]["amount"])
	}

	// Verify price limits
	highLimit := items[0]["high_limit"].(float64)
	lowLimit := items[0]["low_limit"].(float64)
	if highLimit <= lowLimit {
		t.Errorf("High limit should be greater than low limit: %.1f <= %.1f", highLimit, lowLimit)
	}
}

// TestShareFloatAPI tests restricted stock unlock API
func TestShareFloatAPI(t *testing.T) {
	// Restricted stock unlock data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","float_date","float_share","float_ratio","holder_name","share_type"],"items":[["601226.SH","20240110","20240115",100000000,2.5,"首发原股东限售股份","定向增发"],["601226.SH","20240120","20240125",50000000,1.25,"股权激励限售股份","股权激励"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify float share amount
	if items[0]["float_share"].(float64) != 100000000 {
		t.Errorf("Expected float_share 100000000, got %v", items[0]["float_share"])
	}

	// Verify float ratio (percentage)
	if items[0]["float_ratio"].(float64) != 2.5 {
		t.Errorf("Expected float_ratio 2.5, got %v", items[0]["float_ratio"])
	}

	// Verify holder name
	if items[0]["holder_name"].(string) != "首发原股东限售股份" {
		t.Errorf("Expected holder_name '首发原股东限售股份', got %v", items[0]["holder_name"])
	}

	// Verify share type
	if items[0]["share_type"].(string) != "定向增发" {
		t.Errorf("Expected share_type '定向增发', got %v", items[0]["share_type"])
	}
}

// TestBlockTradeAPI tests block trading API
func TestBlockTradeAPI(t *testing.T) {
	// Block trading data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","price","vol","amount","buyer","seller"],"items":[["601226.SH","20240110",15.5,1000000,15500000,"机构专用","机构专用"],["601226.SH","20240110",15.3,500000,7650000,"中信证券股份有限公司上海淮海中路证券营业部","中国国际金融股份有限公司上海淮海中路证券营业部"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify trading price
	if items[0]["price"].(float64) != 15.5 {
		t.Errorf("Expected price 15.5, got %v", items[0]["price"])
	}

	// Verify trading volume
	if items[0]["vol"].(float64) != 1000000 {
		t.Errorf("Expected vol 1000000, got %v", items[0]["vol"])
	}

	// Verify trading amount (price * volume)
	expectedAmount := items[0]["price"].(float64) * items[0]["vol"].(float64)
	actualAmount := items[0]["amount"].(float64)
	if expectedAmount != actualAmount {
		t.Errorf("Amount mismatch: expected %.0f, got %.0f", expectedAmount, actualAmount)
	}

	// Verify buyer and seller
	if items[0]["buyer"].(string) != "机构专用" {
		t.Errorf("Expected buyer '机构专用', got %v", items[0]["buyer"])
	}
}

// TestStkAccountAPI tests stock account opening data API
func TestStkAccountAPI(t *testing.T) {
	// Stock account opening data
	apiResponse := `{"code":0,"message":"","data":{"fields":["date","weekly_new","total","weekly_hold","weekly_trade"],"items":[["20240101",500000,250000000,2000000000,15000000000],["20240108",600000,250600000,2010000000,15100000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify weekly new accounts
	if items[0]["weekly_new"].(float64) != 500000 {
		t.Errorf("Expected weekly_new 500000, got %v", items[0]["weekly_new"])
	}

	// Verify total accounts
	if items[0]["total"].(float64) != 250000000 {
		t.Errorf("Expected total 250000000, got %v", items[0]["total"])
	}

	// Verify weekly holding accounts (2 billion)
	if items[0]["weekly_hold"].(float64) != 2000000000 {
		t.Errorf("Expected weekly_hold 2000000000, got %v", items[0]["weekly_hold"])
	}

	// Verify weekly trading accounts (15 billion)
	if items[0]["weekly_trade"].(float64) != 15000000000 {
		t.Errorf("Expected weekly_trade 15000000000, got %v", items[0]["weekly_trade"])
	}

	// Verify second week has more new accounts
	if items[1]["weekly_new"].(float64) <= items[0]["weekly_new"].(float64) {
		t.Errorf("Second week should have more new accounts: %v <= %v",
			items[1]["weekly_new"], items[0]["weekly_new"])
	}
}

// TestStkHoldernumberAPI tests shareholder number data API
func TestStkHoldernumberAPI(t *testing.T) {
	// Shareholder number data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","end_date","holder_num"],"items":[["601226.SH","20240115","20231231",150000],["601226.SH","20240415","20240331",148000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify shareholder number (should be large)
	if items[0]["holder_num"].(float64) != 150000 {
		t.Errorf("Expected holder_num 150000, got %v", items[0]["holder_num"])
	}

	// Verify second period has fewer shareholders (consolidation)
	if items[1]["holder_num"].(float64) >= items[0]["holder_num"].(float64) {
		t.Errorf("Shareholder count should decrease: %v >= %v",
			items[1]["holder_num"], items[0]["holder_num"])
	}
}

// TestStkHoldertradeAPI tests shareholder trading data API
func TestStkHoldertradeAPI(t *testing.T) {
	// Shareholder trading (increase/decrease) data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","ann_date","holder_name","holder_type","in_de","change_vol","change_ratio","after_share","after_ratio","avg_price","total_share"],"items":[["601226.SH","20240115","张三","个人","增持",1000000,0.5,6000000,3.0,12.5,200000000],["601226.SH","20240116","李四","公司","减持",500000,0.25,4000000,2.0,12.3,200000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify first item (increase)
	if items[0]["in_de"].(string) != "增持" {
		t.Errorf("Expected in_de '增持', got %v", items[0]["in_de"])
	}

	// Verify holder type
	if items[0]["holder_type"].(string) != "个人" {
		t.Errorf("Expected holder_type '个人', got %v", items[0]["holder_type"])
	}

	// Verify change volume (positive for increase)
	if items[0]["change_vol"].(float64) != 1000000 {
		t.Errorf("Expected change_vol 1000000, got %v", items[0]["change_vol"])
	}

	// Verify after share ratio
	if items[0]["after_ratio"].(float64) != 3.0 {
		t.Errorf("Expected after_ratio 3.0, got %v", items[0]["after_ratio"])
	}

	// Verify second item (decrease)
	if items[1]["in_de"].(string) != "减持" {
		t.Errorf("Expected in_de '减持', got %v", items[1]["in_de"])
	}

	// Verify holder type for second item
	if items[1]["holder_type"].(string) != "公司" {
		t.Errorf("Expected holder_type '公司', got %v", items[1]["holder_type"])
	}
}

// TestReportRCAPI tests research report盈利预测 API
func TestReportRCAPI(t *testing.T) {
	// Research report data with盈利预测
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","report_date","report_title","report_type","classify","org_name","author_name","quarter","op_rt","op_pr","tp","np","eps","pe","rd","roe","ev_ebitda","rating","max_price","min_price"],"items":[["601226.SH","华泰证券","20240110","华泰证券2023年业绩预告","公司报告","医药生物","华泰证券","张三",3,15.5,2.3,18.5,5000000000,1.5,12.3,2.5,18.5,25.5,15.5,20.5,12.5]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify report type
	if items[0]["report_type"].(string) != "公司报告" {
		t.Errorf("Expected report_type '公司报告', got %v", items[0]["report_type"])
	}

	// Verify organization name
	if items[0]["org_name"].(string) != "华泰证券" {
		t.Errorf("Expected org_name '华泰证券', got %v", items[0]["org_name"])
	}

	// Verify target price (tp)
	if items[0]["tp"].(float64) != 18.5 {
		t.Errorf("Expected tp 18.5, got %v", items[0]["tp"])
	}

	// Verify max price > min price
	maxPrice := items[0]["max_price"].(float64)
	minPrice := items[0]["min_price"].(float64)
	if maxPrice <= minPrice {
		t.Errorf("Max price should be greater than min price: %.1f <= %.1f", maxPrice, minPrice)
	}
}

// TestCyqPerfAPI tests chip distribution performance API
func TestCyqPerfAPI(t *testing.T) {
	// Chip distribution performance data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","his_low","his_high","cost_5pct","cost_15pct","cost_50pct","cost_85pct","cost_95pct","weight_avg","winner_rate"],"items":[["601226.SH","20240110",8.5,15.5,9.5,11.2,12.5,14.3,15.0,12.5,65.5]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify historical low and high
	hisLow := items[0]["his_low"].(float64)
	hisHigh := items[0]["his_high"].(float64)
	if hisLow >= hisHigh {
		t.Errorf("Historical low should be less than high: %.1f >= %.1f", hisLow, hisHigh)
	}

	// Verify cost percentiles (should be ascending)
	cost5pct := items[0]["cost_5pct"].(float64)
	cost15pct := items[0]["cost_15pct"].(float64)
	cost50pct := items[0]["cost_50pct"].(float64)
	cost85pct := items[0]["cost_85pct"].(float64)
	cost95pct := items[0]["cost_95pct"].(float64)

	if cost5pct >= cost15pct || cost15pct >= cost50pct || cost50pct >= cost85pct || cost85pct >= cost95pct {
		t.Errorf("Cost percentiles should be ascending: 5%%=%.1f, 15%%=%.1f, 50%%=%.1f, 85%%=%.1f, 95%%=%.1f",
			cost5pct, cost15pct, cost50pct, cost85pct, cost95pct)
	}

	// Verify winner rate (percentage)
	if items[0]["winner_rate"].(float64) != 65.5 {
		t.Errorf("Expected winner_rate 65.5, got %v", items[0]["winner_rate"])
	}
}

// TestCyqChipsAPI tests chip distribution details API
func TestCyqChipsAPI(t *testing.T) {
	// Chip distribution details data
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","price","percent"],"items":[["601226.SH","20240110",8.5,2.5],["601226.SH","20240110",9.0,5.3],["601226.SH","20240110",10.0,8.7],["601226.SH","20240110",11.0,6.2],["601226.SH","20240110",12.0,4.8],["601226.SH","20240110",12.5,3.5]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 6 {
		t.Errorf("Expected 6 items, got %d", len(items))
	}

	// Verify price levels (should be ascending)
	if items[0]["price"].(float64) != 8.5 {
		t.Errorf("Expected price 8.5, got %v", items[0]["price"])
	}

	// Verify percentages (all should be positive and sum to reasonable value)
	totalPercent := 0.0
	for i, item := range items {
		percent := item["percent"].(float64)
		if percent <= 0 {
			t.Errorf("Item %d: percent should be positive, got %.1f", i, percent)
		}
		totalPercent += percent
	}

	// Total percentage should be reasonable (around 30% for main chip areas)
	if totalPercent < 20 || totalPercent > 50 {
		t.Errorf("Total percent seems unreasonable: %.1f", totalPercent)
	}
}

// TestBrokerRecommendAPI tests broker recommendations API
func TestBrokerRecommendAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["month","broker","ts_code","name"],"items":[["201601","中信证券","601226.SH","华泰证券"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if items[0]["broker"].(string) != "中信证券" {
		t.Errorf("Expected broker '中信证券'")
	}
}

// TestMarginAPI tests margin trading summary API
func TestMarginAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["exchange_id","trade_date","rzye","rzmre","rzche","rqye","rqmcl","rqchl"],"items":[["SSE","20240205","985546000000.00","534536000000.00","534536000000.00","76543000000.00","38654000000.00","38654000000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// The parser converts all values from strings, so numeric values may be strings
	// Let's check what type we actually got
	rzyeStr := items[0]["rzye"].(string)
	rzye, err := strconv.ParseFloat(rzyeStr, 64)
	if err != nil {
		t.Fatalf("Failed to parse rzye: %v", err)
	}

	rqyeStr := items[0]["rqye"].(string)
	rqye, err := strconv.ParseFloat(rqyeStr, 64)
	if err != nil {
		t.Fatalf("Failed to parse rqye: %v", err)
	}

	// Verify margin balance (融资余额) is positive
	if rzye <= 0 {
		t.Errorf("融资余额 should be positive, got %.2f", rzye)
	}

	// Verify short selling balance (融券余额) is positive
	if rqye <= 0 {
		t.Errorf("融券余额 should be positive, got %.2f", rqye)
	}
}

// TestMarginDetailAPI tests margin trading detail API
func TestMarginDetailAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","rzye","rqye"],"items":[["20240205","512880.SH","1234567890.12","23456789.01"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify ETF margin trading data
	if items[0]["ts_code"].(string) != "512880.SH" {
		t.Errorf("Expected ts_code '512880.SH'")
	}

	// Margin balance should be much larger than short selling balance
	rzyeStr := items[0]["rzye"].(string)
	rzye, _ := strconv.ParseFloat(rzyeStr, 64)

	rqyeStr := items[0]["rqye"].(string)
	rqye, _ := strconv.ParseFloat(rqyeStr, 64)

	if rzye < rqye {
		t.Errorf("融资余额 should generally be larger than 融券余额")
	}
}

// TestMarginSecsAPI tests margin trading securities list API
func TestMarginSecsAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","exchange_id","trade_date"],"items":[["512880.SH","证券ETF","SSE","20240205"],["510300.SH","300ETF","SSE","20240205"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 margin securities, got %d", len(items))
	}

	// Verify ETF margin securities
	if items[0]["ts_code"].(string) != "512880.SH" {
		t.Errorf("Expected first security to be '512880.SH'")
	}
}

// TestSlbSecAPI tests securities lending (SLB) securities summary API
func TestSlbSecAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","slb_qty","slb_vol"],"items":[["20240205","510050.SH","1000000.00","500000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify SLB quantity is positive
	slbQtyStr := items[0]["slb_qty"].(string)
	slbQty, _ := strconv.ParseFloat(slbQtyStr, 64)
	if slbQty <= 0 {
		t.Errorf("SLB quantity should be positive, got %.2f", slbQty)
	}

	// Volume should be less than or equal to quantity
	slbVolStr := items[0]["slb_vol"].(string)
	slbVol, _ := strconv.ParseFloat(slbVolStr, 64)
	if slbVol > slbQty {
		t.Errorf("Volume should not exceed quantity")
	}
}

// TestSlbLenAPI test securities lending (SLB) lending summary API
func TestSlbLenAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","slb_type","slb_amt","slb_vol"],"items":[["20240205","融券","100000000.00","5000000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify SLB type
	if items[0]["slb_type"].(string) != "融券" {
		t.Errorf("Expected SLB type '融券'")
	}

	// Verify amount and volume are positive
	slbAmtStr := items[0]["slb_amt"].(string)
	slbAmt, _ := strconv.ParseFloat(slbAmtStr, 64)
	if slbAmt <= 0 {
		t.Errorf("SLB amount should be positive, got %.2f", slbAmt)
	}
}

// TestSlbSecDetailAPI tests securities lending detail API
func TestSlbSecDetailAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","slb_qty","slb_vol","rate"],"items":[["20240205","510300.SH","2000000.00","1000000.00","5.5"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify rate is reasonable (typically 2-10%)
	rateStr := items[0]["rate"].(string)
	rate, _ := strconv.ParseFloat(rateStr, 64)
	if rate < 2 || rate > 10 {
		t.Errorf("Rate should be reasonable (2-10%%), got %.1f%%", rate)
	}
}

// TestSlbLenMmAPI tests market maker lending API
func TestSlbLenMmAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","mm_code","slb_qty","slb_vol"],"items":[["20240205","512000.SH","MM001","3000000.00","1500000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify market maker code exists
	mmCode := items[0]["mm_code"].(string)
	if mmCode == "" {
		t.Errorf("Market maker code should not be empty")
	}

	// Verify quantity and volume are positive
	slbQtyStr := items[0]["slb_qty"].(string)
	slbQty, _ := strconv.ParseFloat(slbQtyStr, 64)
	if slbQty <= 0 {
		t.Errorf("SLB quantity should be positive, got %.2f", slbQty)
	}
}

// TestMoneyflowAPI tests individual stock money flow API
func TestMoneyflowAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","buy_sm_vol","buy_sm_amount","sell_sm_vol","sell_sm_amount","buy_md_vol","buy_md_amount","sell_md_vol","sell_md_amount","buy_lg_vol","buy_lg_amount","sell_lg_vol","sell_lg_amount","buy_elg_vol","buy_elg_amount","sell_elg_vol","sell_elg_amount","net_mf_vol","net_mf_amount"],"items":[["000001.SZ","20240110","1000000","5000000.00","800000","4000000.00","500000","3000000.00","300000","2000000.00","200000","2000000.00","100000","1000000.00","50000","500000.00","30000","300000.00","2170000","7500000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net money flow amount
	netMfAmtStr := items[0]["net_mf_amount"].(string)
	netMfAmt, _ := strconv.ParseFloat(netMfAmtStr, 64)
	if netMfAmt <= 0 {
		t.Errorf("Net money flow amount should be positive, got %.2f", netMfAmt)
	}

	// Verify buy and sell amounts relationship
	buySmAmtStr := items[0]["buy_sm_amount"].(string)
	buySmAmt, _ := strconv.ParseFloat(buySmAmtStr, 64)
	sellSmAmtStr := items[0]["sell_sm_amount"].(string)
	sellSmAmt, _ := strconv.ParseFloat(sellSmAmtStr, 64)

	if buySmAmt <= 0 || sellSmAmt <= 0 {
		t.Errorf("Buy and sell amounts should be positive")
	}
}

// TestMoneyflowHsgtAPI tests HK stock connect money flow API
func TestMoneyflowHsgtAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ggt_ss","ggt_sz","hgt","sgt","north_money","south_money"],"items":[["20240110","5000000000","3000000000","4000000000","2000000000","6000000000","2500000000"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify north money (北向资金) is positive
	northMoneyStr := items[0]["north_money"].(string)
	northMoney, _ := strconv.ParseFloat(northMoneyStr, 64)
	if northMoney <= 0 {
		t.Errorf("North money should be positive, got %.2f", northMoney)
	}

	// Verify south money (南向资金) is positive
	southMoneyStr := items[0]["south_money"].(string)
	southMoney, _ := strconv.ParseFloat(southMoneyStr, 64)
	if southMoney <= 0 {
		t.Errorf("South money should be positive, got %.2f", southMoney)
	}
}

// TestMoneyflowThsAPI tests Tonghuashun individual stock money flow API
func TestMoneyflowThsAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","pct_change","latest","net_amount","net_d5_amount","buy_lg_amount","buy_lg_amount_rate","buy_md_amount","buy_md_amount_rate","buy_sm_amount","buy_sm_amount_rate"],"items":[["20240110","000001.SZ","平安银行","2.5","12.5","100000000.00","500000000.00","30000000.00","30.5","40000000.00","40.5","30000000.00","29.0"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify percentage change is reasonable
	pctChangeStr := items[0]["pct_change"].(string)
	pctChange, _ := strconv.ParseFloat(pctChangeStr, 64)
	if pctChange < -10 || pctChange > 10 {
		t.Errorf("Percentage change seems unreasonable: %.2f%%", pctChange)
	}

	// Verify buy amounts sum reasonably
	buyLgAmtStr := items[0]["buy_lg_amount"].(string)
	buyLgAmt, _ := strconv.ParseFloat(buyLgAmtStr, 64)
	if buyLgAmt <= 0 {
		t.Errorf("Buy large amount should be positive")
	}
}

// TestMoneyflowDcAPI tests Eastmoney individual stock money flow API
func TestMoneyflowDcAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","pct_change","close","net_amount","net_amount_rate","buy_elg_amount","buy_elg_amount_rate","buy_lg_amount","buy_lg_amount_rate","buy_md_amount","buy_md_amount_rate","buy_sm_amount","buy_sm_amount_rate"],"items":[["20240110","000001.SZ","平安银行","2.5","12.5","150000000.00","1.5","5000000.00","5.0","40000000.00","30.0","60000000.00","40.0","45000000.00","25.0"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify close price
	closeStr := items[0]["close"].(string)
	close, _ := strconv.ParseFloat(closeStr, 64)
	if close <= 0 {
		t.Errorf("Close price should be positive, got %.2f", close)
	}
}

// TestMoneyflowCntThsAPI tests Tonghuashun concept sector money flow API
func TestMoneyflowCntThsAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","lead_stock","close_price","pct_change","industry_index","company_num","pct_change_stock","net_buy_amount","net_sell_amount","net_amount"],"items":[["20240110","000001.SZ","人工智能","000001.SZ","12.5","3.5","1000.5","150","2.5","800000000.00","600000000.00","200000000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify company count is reasonable
	companyNumStr := items[0]["company_num"].(string)
	companyNum, _ := strconv.ParseFloat(companyNumStr, 64)
	if companyNum < 10 || companyNum > 500 {
		t.Errorf("Company number seems unreasonable: %.0f", companyNum)
	}
}

// TestMoneyflowIndThsAPI tests Tonghuashun industry money flow API
func TestMoneyflowIndThsAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","industry","lead_stock","close","pct_change","company_num","pct_change_stock","close_price","net_buy_amount","net_sell_amount","net_amount"],"items":[["20240110","000001.SZ","银行","000001.SZ","12.5","1.5","200","1.2","12.8","1000000000.00","800000000.00","200000000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net amount
	netAmtStr := items[0]["net_amount"].(string)
	netAmt, _ := strconv.ParseFloat(netAmtStr, 64)
	if netAmt <= 0 {
		t.Errorf("Net amount should be positive, got %.2f", netAmt)
	}
}

// TestMoneyflowIndDcAPI tests Eastmoney sector money flow API
func TestMoneyflowIndDcAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","content_type","ts_code","name","pct_change","close","net_amount","net_amount_rate","buy_elg_amount","buy_elg_amount_rate","buy_lg_amount","buy_lg_amount_rate","buy_md_amount","buy_md_amount_rate","buy_sm_amount","buy_sm_amount_rate","buy_sm_amount_stock","rank"],"items":[["20240110","行业","000001.SZ","银行","2.5","3500.5","5000000000.00","1.5","500000000.00","10.5","1500000000.00","30.5","2000000000.00","40.5","1000000000.00","20.5","000001.SZ","1"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify content type
	contentType := items[0]["content_type"].(string)
	if contentType != "行业" {
		t.Errorf("Expected content_type '行业', got '%s'", contentType)
	}

	// Verify rank is reasonable (top sectors)
	rankStr := items[0]["rank"].(string)
	rank, _ := strconv.ParseFloat(rankStr, 64)
	if rank < 1 || rank > 100 {
		t.Errorf("Rank seems unreasonable: %.0f", rank)
	}
}

// TestMoneyflowMktDcAPI tests Eastmoney market money flow API
func TestMoneyflowMktDcAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","close_sh","pct_change_sh","close_sz","pct_change_sz","net_amount","net_amount_rate","buy_elg_amount","buy_elg_amount_rate","buy_lg_amount","buy_lg_amount_rate","buy_md_amount","buy_md_amount_rate","buy_sm_amount","buy_sm_amount_rate"],"items":[["20240110","3200.5","1.5","10500.8","2.0","80000000000.00","2.5","10000000000.00","12.5","25000000000.00","31.5","30000000000.00","37.5","15000000000.00","18.5"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify SH index close
	closeShStr := items[0]["close_sh"].(string)
	closeSh, _ := strconv.ParseFloat(closeShStr, 64)
	if closeSh <= 0 {
		t.Errorf("SH close should be positive, got %.2f", closeSh)
	}

	// Verify SZ index close
	closeSzStr := items[0]["close_sz"].(string)
	closeSz, _ := strconv.ParseFloat(closeSzStr, 64)
	if closeSz <= 0 {
		t.Errorf("SZ close should be positive, got %.2f", closeSz)
	}

	// Verify net amount is in billions
	netAmtStr := items[0]["net_amount"].(string)
	netAmt, _ := strconv.ParseFloat(netAmtStr, 64)
	if netAmt < 10000000000 { // Less than 10 billion
		t.Errorf("Market net amount seems too low: %.2f", netAmt)
	}
}

// TestTopListAPI tests dragon tiger list (龙虎榜) daily trading API
func TestTopListAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","close","pct_change","turnover_rate","amount","l_sell","l_buy","l_amount","net_amount","net_rate","amount_rate","float_values","reason"],"items":[["20240110","601226.SH","华泰证券","12.5","5.2","15.5","1000000000.00","500000000.00","600000000.00","550000000.00","100000000.00","2.5","5.5","1000000000.00","涨停"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net amount
	netAmtStr := items[0]["net_amount"].(string)
	netAmt, _ := strconv.ParseFloat(netAmtStr, 64)
	if netAmt <= 0 {
		t.Errorf("Net amount should be positive, got %.2f", netAmt)
	}
}

// TestTopInstAPI tests dragon tiger list institution trading API
func TestTopInstAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","exalter","buy","buy_rate","sell","sell_rate","net_buy","side","reason"],"items":[["20240110","601226.SH","机构专用","50000000.00","30.5","40000000.00","25.5","10000000.00","买入","大宗交易"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net buy is positive
	netBuyStr := items[0]["net_buy"].(string)
	netBuy, _ := strconv.ParseFloat(netBuyStr, 64)
	if netBuy <= 0 {
		t.Errorf("Net buy should be positive, got %.2f", netBuy)
	}

	// Verify side is reasonable
	side := items[0]["side"].(string)
	if side != "买入" && side != "卖出" {
		t.Errorf("Side should be '买入' or '卖出', got '%s'", side)
	}
}

// TestLimitListThsAPI tests Tonghuashun limit list (涨跌停榜) API
func TestLimitListThsAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","name","price","pct_chg","open_num","lu_desc","limit_type","tag","status","limit_order","limit_amount","turnover_rate","free_float","lu_limit_order","limit_up_suc_rate","turnover","market_type"],"items":[["20240110","601226.SH","华泰证券","12.5","10.0","9:30:00","首次涨停","U","龙头股","封板","100000","500000000.00","15.5","1000000000","80000","95.5","50000000.00","HS"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify limit up success rate
	luSuccRateStr := items[0]["limit_up_suc_rate"].(string)
	luSuccRate, _ := strconv.ParseFloat(luSuccRateStr, 64)
	if luSuccRate < 80 || luSuccRate > 100 {
		t.Errorf("Limit up success rate seems unreasonable: %.1f%%", luSuccRate)
	}
}

// TestLimitListDAPI tests limit list (涨跌停数据) API
func TestLimitListDAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","industry","name","close","pct_chg","amount","limit_amount","float_mv","total_mv","turnover_ratio","fd_amount","first_time","last_time","open_times","up_stat","limit_times","limit"],"items":[["20240110","601226.SH","证券","华泰证券","12.5","10.0","1000000000.00","500000000.00","10000000000.00","15000000000.00","15.5","300000000.00","9:30:00","9:31:00","1","1","5","U"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify limit times is reasonable
	limitTimesStr := items[0]["limit_times"].(string)
	limitTimes, _ := strconv.ParseFloat(limitTimesStr, 64)
	if limitTimes < 1 || limitTimes > 30 {
		t.Errorf("Limit times seems unreasonable: %.0f", limitTimes)
	}
}

// TestLimitStepAPI tests consecutive limit board (连板个股) API
func TestLimitStepAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","trade_date","nums"],"items":[["601226.SH","华泰证券","20240110","5"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify consecutive limit count
	numsStr := items[0]["nums"].(string)
	nums, _ := strconv.ParseFloat(numsStr, 64)
	if nums < 2 || nums > 30 {
		t.Errorf("Consecutive limit count seems unreasonable: %.0f", nums)
	}
}

// TestLimitCptListAPI tests daily limit sector statistics API
func TestLimitCptListAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","trade_date","days","up_stat","cons_nums","up_nums","pct_chg","rank"],"items":[["000001.SZ","人工智能","20240110","5","1","10","8","5.2","1"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify sector rank
	rankStr := items[0]["rank"].(string)
	rank, _ := strconv.ParseFloat(rankStr, 64)
	if rank < 1 || rank > 100 {
		t.Errorf("Sector rank seems unreasonable: %.0f", rank)
	}

	// Verify up numbers is reasonable
	upNumsStr := items[0]["up_nums"].(string)
	upNums, _ := strconv.ParseFloat(upNumsStr, 64)
	if upNums < 1 || upNums > 100 {
		t.Errorf("Up numbers seems unreasonable: %.0f", upNums)
	}
}

// TestThsIndexAPI tests Tonghuashun index API
func TestThsIndexAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","count","exchange","list_date","type"],"items":[["000001.SH","上证指数","5000","SSE","19901219","I"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify stock count in index
	countStr := items[0]["count"].(string)
	count, _ := strconv.ParseFloat(countStr, 64)
	if count < 10 || count > 10000 {
		t.Errorf("Index component count seems unreasonable: %.0f", count)
	}
}

// TestThsDailyAPI tests Tonghuashun index daily API
func TestThsDailyAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","open","high","low","close","pre_close","avg_price","change","pct_change","vol","turnover_rate"],"items":[["000001.SH","20240110","3200.5","3250.0","3180.0","3240.5","3200.0","3220.0","40.5","1.26","100000000","15.5"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify OHLC relationship
	highStr := items[0]["high"].(string)
	high, _ := strconv.ParseFloat(highStr, 64)
	lowStr := items[0]["low"].(string)
	low, _ := strconv.ParseFloat(lowStr, 64)

	if high <= low {
		t.Errorf("High should be greater than low: high=%.2f, low=%.2f", high, low)
	}
}

// TestDcIndexAPI tests Eastmoney sector index API
func TestDcIndexAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","name","leading","leading_code","pct_change","leading_pct","total_mv","turnover_rate","up_num","down_num","idx_type","level"],"items":[["000001.SZ","20240110","人工智能","000001.SZ","000001","3.5","5.2","1000000000000.00","20.5","150","50","概念","二级行业"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify sector type
	idxType := items[0]["idx_type"].(string)
	if idxType != "概念" && idxType != "行业" {
		t.Errorf("Index type should be '概念' or '行业', got '%s'", idxType)
	}

	// Verify up/down numbers relationship
	upNumStr := items[0]["up_num"].(string)
	upNum, _ := strconv.ParseFloat(upNumStr, 64)
	downNumStr := items[0]["down_num"].(string)
	downNum, _ := strconv.ParseFloat(downNumStr, 64)

	if upNum <= 0 || downNum <= 0 {
		t.Errorf("Up and down numbers should be positive: up=%.0f, down=%.0f", upNum, downNum)
	}
}

// TestStmAuctionAPI tests auction trading API
func TestStmAuctionAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","vol","price","amount","pre_close","turnover_rate","volume_ratio","float_share"],"items":[["601226.SH","20240110","5000000","12.5","62500000.00","12.0","5.5","2.5","1000000000"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify auction volume is positive
	volStr := items[0]["vol"].(string)
	vol, _ := strconv.ParseFloat(volStr, 64)
	if vol <= 0 {
		t.Errorf("Auction volume should be positive, got %.2f", vol)
	}
}

// TestHmListAPI tests hot money list API
func TestHmListAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["name","desc","orgs"],"items":[["赵老哥","著名游资","东方财富股份有限公司绍兴分公司"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify hot money name
	name := items[0]["name"].(string)
	if name == "" {
		t.Errorf("Hot money name should not be empty")
	}
}

// TestHmDetailAPI tests hot money detail API
func TestHmDetailAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","ts_code","ts_name","buy_amount","sell_amount","net_amount","hm_name","hm_orgs"],"items":[["20240110","601226.SH","华泰证券","100000000.00","80000000.00","20000000.00","赵老哥","东方财富股份有限公司绍兴分公司"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify net amount
	netAmtStr := items[0]["net_amount"].(string)
	netAmt, _ := strconv.ParseFloat(netAmtStr, 64)
	if netAmt <= 0 {
		t.Errorf("Net amount should be positive, got %.2f", netAmt)
	}
}

// TestDcHotAPI tests Eastmoney hot list API
func TestDcHotAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["trade_date","data_type","ts_code","ts_name","rank","pct_change","current_price","hot","concept","rank_time","rank_reason"],"items":[["20240110","热门股票","601226.SH","华泰证券","1","5.2","12.5","true","人工智能","09:30:00","龙头涨停"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify rank is reasonable
	rankStr := items[0]["rank"].(string)
	rank, _ := strconv.ParseFloat(rankStr, 64)
	if rank < 1 || rank > 100 {
		t.Errorf("Rank seems unreasonable: %.0f", rank)
	}
}

// TestFundDailyAPI tests fund daily data API
func TestFundDailyAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","pre_close","open","high","low","close","change","pct_chg","vol","amount"],"items":[["563050.SH","20240110","1.250","1.255","1.265","1.250","1.260","0.010","0.80","1000000","1260000.00"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify OHLC relationships
	highStr := items[0]["high"].(string)
	high, _ := strconv.ParseFloat(highStr, 64)
	lowStr := items[0]["low"].(string)
	low, _ := strconv.ParseFloat(lowStr, 64)
	openStr := items[0]["open"].(string)
	open, _ := strconv.ParseFloat(openStr, 64)
	closeStr := items[0]["close"].(string)
	close, _ := strconv.ParseFloat(closeStr, 64)

	if high < low {
		t.Errorf("High should be >= low: high=%.2f, low=%.2f", high, low)
	}
	if high < open || high < close {
		t.Errorf("High should be >= open and close: high=%.2f, open=%.2f, close=%.2f", high, open, close)
	}
	if low > open || low > close {
		t.Errorf("Low should be <= open and close: low=%.2f, open=%.2f, close=%.2f", low, open, close)
	}

	// Verify volume and amount are positive
	volStr := items[0]["vol"].(string)
	vol, _ := strconv.ParseFloat(volStr, 64)
	if vol <= 0 {
		t.Errorf("Volume should be positive, got %.2f", vol)
	}

	amountStr := items[0]["amount"].(string)
	amount, _ := strconv.ParseFloat(amountStr, 64)
	if amount <= 0 {
		t.Errorf("Amount should be positive, got %.2f", amount)
	}

	// Verify percentage change is reasonable
	pctChgStr := items[0]["pct_chg"].(string)
	pctChg, _ := strconv.ParseFloat(pctChgStr, 64)
	if pctChg < -11 || pctChg > 11 {
		t.Errorf("Percentage change seems unreasonable: %.2f%%", pctChg)
	}
}

// TestHkBasicAPI tests Hong Kong stock basic information API
func TestHkBasicAPI(t *testing.T) {
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","fullname","enname","cn_spell","market","list_status","list_date","delist_date","trade_unit","isin","curr_type"],"items":[["00700.HK","腾讯控股","腾讯控股有限公司","Tencent Holdings Limited","TX","MainBoard","L","20040416","","100","HK0000700214","HKD"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify essential fields are present
	tsCode := items[0]["ts_code"].(string)
	if tsCode == "" {
		t.Errorf("TS code should not be empty")
	}

	name := items[0]["name"].(string)
	if name == "" {
		t.Errorf("Name should not be empty")
	}

	// Verify market type is valid
	market := items[0]["market"].(string)
	if market != "MainBoard" && market != "GEM" {
		t.Errorf("Market should be 'MainBoard' or 'GEM', got '%s'", market)
	}

	// Verify list status is valid
	listStatus := items[0]["list_status"].(string)
	if listStatus != "L" && listStatus != "D" && listStatus != "P" {
		t.Errorf("List status should be 'L', 'D', or 'P', got '%s'", listStatus)
	}

	// Verify trade unit is positive
	tradeUnitStr := items[0]["trade_unit"].(string)
	tradeUnit, _ := strconv.ParseFloat(tradeUnitStr, 64)
	if tradeUnit <= 0 {
		t.Errorf("Trade unit should be positive, got %.2f", tradeUnit)
	}

	// Verify currency type
	currType := items[0]["curr_type"].(string)
	if currType != "HKD" && currType != "CNY" && currType != "USD" {
		t.Errorf("Currency type should be 'HKD', 'CNY', or 'USD', got '%s'", currType)
	}
}
