package tushare

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chenniannian90/tushare-go/types"
)

// TestResult represents the result of testing a single API
type TestResult struct {
	APIName   string            `json:"api_name"`
	Success   bool              `json:"success"`
	Error     string            `json:"error,omitempty"`
	HasData   bool              `json:"has_data"`
	DataCount int               `json:"data_count,omitempty"`
	Duration  string            `json:"duration"`
	Params    map[string]string `json:"params"`
	Fields    []string          `json:"fields,omitempty"`
}

// TestReport represents the full test report
type TestReport struct {
	Token    string       `json:"token"`
	TestDate string       `json:"test_date"`
	Total    int          `json:"total"`
	Passed   int          `json:"passed"`
	Failed   int          `json:"failed"`
	Results  []TestResult `json:"results"`
}

var report = TestReport{
	Results: []TestResult{},
}

func getToken() string {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		data, err := os.ReadFile(".env")
		if err == nil {
			content := string(data)
			if len(content) > 13 {
				token = content[13:]
				if len(token) > 0 && token[len(token)-1] == '\n' {
					token = token[:len(token)-1]
				}
			}
		}
	}
	return token
}

func runTest(apiName string, params map[string]string, fields []string, testFunc func() (*types.APIResponse, error)) TestResult {
	start := time.Now()
	result := TestResult{
		APIName: apiName,
		Params:  params,
		Fields:  fields,
	}

	resp, err := testFunc()
	result.Duration = time.Since(start).String()

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		return result
	}

	result.Success = true

	if resp != nil && len(resp.Data.Items) > 0 {
		result.HasData = true
		result.DataCount = len(resp.Data.Items)
	}

	return result
}

func TestIntegrationAPIs(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("No token found in .env file")
		return
	}

	client := New(token)
	report.Token = maskToken(token)
	report.TestDate = time.Now().Format("2006-01-02 15:04:05")

	t.Log("🚀 Starting API Integration Tests")
	t.Log("Token:", report.Token)
	t.Log("")

	// Stock APIs
	t.Run("Stock", func(t *testing.T) {
		testStockAPIs(t, client)
	})

	// Index APIs
	t.Run("Index", func(t *testing.T) {
		testIndexAPIs(t, client)
	})

	// Market APIs
	t.Run("Market", func(t *testing.T) {
		testMarketAPIs(t, client)
	})

	// Finance APIs
	t.Run("Finance", func(t *testing.T) {
		testFinanceAPIs(t, client)
	})

	// Hsgt APIs
	t.Run("Hsgt", func(t *testing.T) {
		testHsgtAPIs(t, client)
	})

	// Margin APIs
	t.Run("Margin", func(t *testing.T) {
		testMarginAPIs(t, client)
	})

	// Realtime APIs
	t.Run("Realtime", func(t *testing.T) {
		testRealtimeAPIs(t, client)
	})

	// ETF APIs
	t.Run("ETF", func(t *testing.T) {
		testETFAPIs(t, client)
	})

	// Pledge APIs
	t.Run("Pledge", func(t *testing.T) {
		testPledgeAPIs(t, client)
	})

	// Toplist APIs
	t.Run("Toplist", func(t *testing.T) {
		testToplistAPIs(t, client)
	})

	// Holder APIs
	t.Run("Holder", func(t *testing.T) {
		testHolderAPIs(t, client)
	})

	// Concept APIs
	t.Run("Concept", func(t *testing.T) {
		testConceptAPIs(t, client)
	})

	// Ths APIs
	t.Run("Ths", func(t *testing.T) {
		testThsAPIs(t, client)
	})

	// Sw APIs
	t.Run("Sw", func(t *testing.T) {
		testSwAPIs(t, client)
	})

	// Limit APIs
	t.Run("Limit", func(t *testing.T) {
		testLimitAPIs(t, client)
	})

	// Research APIs
	t.Run("Research", func(t *testing.T) {
		testResearchAPIs(t, client)
	})

	// Repurchase APIs
	t.Run("Repurchase", func(t *testing.T) {
		testRepurchaseAPIs(t, client)
	})

	// Generate report
	generateReport(t)
}

func testStockAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"StockBasic",
			map[string]string{
				"list_status": "L",
				"exchange":    "SSE",
				"limit":       "5",
			},
			nil,
		},
		{
			"TradeCal",
			map[string]string{
				"exchange":  "SSE",
				"start_date": "20240101",
				"end_date":   "20240110",
			},
			nil,
		},
		{
			"HSConst",
			map[string]string{
				"hs_type": "SH",
			},
			nil,
		},
		{
			"StockCompany",
			map[string]string{
				"exchange": "SSE",
				"limit":    "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runTest("Stock."+test.name, test.params, test.fields, func() (*types.APIResponse, error) {
				return client.Basic.StockBasic(test.params, test.fields)
			})
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testIndexAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"IndexDaily",
			map[string]string{
				"ts_code":    "000001.SH",
				"start_date": "20240101",
				"end_date":   "20240105",
			},
			nil,
		},
		{
			"IndexBasic",
			map[string]string{
				"ts_code":  "",
				"market":   "",
				"publisher": "",
				"category": "",
				"name":     "",
			},
			[]string{"ts_code", "name", "market", "publisher", "category", "base_date", "base_point", "list_date"},
		},
		{
			"IndexWeight",
			map[string]string{
				"index_code": "",
				"trade_date": "20240105",
			},
			[]string{"index_code", "con_code", "trade_date", "weight"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "IndexDaily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Index.IndexDaily(test.params, test.fields)
				}
			case "IndexBasic":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Index.IndexBasic(test.params, test.fields)
				}
			case "IndexWeight":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Index.IndexWeight(test.params, test.fields)
				}
			}

			result := runTest("Index."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testMarketAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Daily",
			map[string]string{
				"trade_date": "20240102",
			},
			nil,
		},
		{
			"DailyBasic",
			map[string]string{
				"trade_date": "20240102",
			},
			nil,
		},
		{
			"MoneyFlow",
			map[string]string{
				"ts_code": "000001.SZ",
				"limit":   "10",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "Daily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Market.Daily(test.params, test.fields)
				}
			case "DailyBasic":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Market.DailyBasic(test.params, test.fields)
				}
			case "MoneyFlow":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Market.MoneyFlow(test.params, test.fields)
				}
			}

			result := runTest("Market."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testFinanceAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Income",
			map[string]string{
				"ts_code":    "000001.SZ",
				"start_date": "20230101",
				"end_date":   "20231231",
				"limit":      "1",
			},
			nil,
		},
		{
			"BalanceSheet",
			map[string]string{
				"ts_code":    "000001.SZ",
				"start_date": "20230101",
				"end_date":   "20231231",
				"limit":      "1",
			},
			nil,
		},
		{
			"FinaIndicator",
			map[string]string{
				"ts_code":    "000001.SZ",
				"start_date": "20230101",
				"end_date":   "20231231",
				"limit":      "1",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runTest("Finance."+test.name, test.params, test.fields, func() (*types.APIResponse, error) {
				return client.Finance.Income(test.params, test.fields)
			})
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testHsgtAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"MoneyflowHsgt",
			map[string]string{
				"start_date": "20240102",
				"end_date":   "20240105",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runTest("Hsgt."+test.name, test.params, test.fields, func() (*types.APIResponse, error) {
				return client.Moneyflow.MoneyflowHsgt(test.params, test.fields)
			})
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testMarginAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Margin",
			map[string]string{
				"ts_code":    "000001.SZ",
				"start_date": "20240102",
				"end_date":   "20240105",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runTest("Margin."+test.name, test.params, test.fields, func() (*types.APIResponse, error) {
				return client.Margin.Margin(test.params, test.fields)
			})
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testRealtimeAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"RTK",
			map[string]string{},
			nil,
		},
		{
			"RealTimeQuote",
			map[string]string{
				"ts_code": "000001.SZ",
				"src":     "sz",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runTest("Realtime."+test.name, test.params, test.fields, func() (*types.APIResponse, error) {
				return client.Market.RTK(test.params, test.fields)
			})
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testETFAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"ETFBasic",
			map[string]string{
				"limit": "10",
			},
			nil,
		},
		{
			"FundDaily",
			map[string]string{
				"ts_code": "159149.SZ",
				"limit":   "5",
			},
			nil,
		},
		{
			"FundAdj",
			map[string]string{
				"ts_code": "159149.SZ",
				"limit":   "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "ETFBasic":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Etf.ETFBasic(test.params, test.fields)
				}
			case "FundDaily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Etf.FundDaily(test.params, test.fields)
				}
			case "FundAdj":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Etf.FundAdj(test.params, test.fields)
				}
			}

			result := runTest("Etf."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testPledgeAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"PledgeStat",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "5",
			},
			nil,
		},
		{
			"PledgeDetail",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "PledgeStat":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.PledgeStat(test.params, test.fields)
				}
			case "PledgeDetail":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.PledgeDetail(test.params, test.fields)
				}
			}

			result := runTest("Pledge."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testToplistAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"TopList",
			map[string]string{
				"trade_date": "20240105",
				"limit":      "10",
			},
			nil,
		},
		{
			"TopInst",
			map[string]string{
				"trade_date": "20240105",
				"limit":      "10",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "TopList":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.TopList(test.params, test.fields)
				}
			case "TopInst":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.TopInst(test.params, test.fields)
				}
			}

			result := runTest("Toplist."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testHolderAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Top10Holders",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "5",
			},
			nil,
		},
		{
			"Top10FloatHolders",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "5",
			},
			nil,
		},
		{
			"StkHolderNumber",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "Top10Holders":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.Top10Holders(test.params, test.fields)
				}
			case "Top10FloatHolders":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.Top10FloatHolders(test.params, test.fields)
				}
			case "StkHolderNumber":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.StkHolderNumber(test.params, test.fields)
				}
			}

			result := runTest("Holder."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testConceptAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Concept",
			map[string]string{
				"ts_code": "000001.SZ",
				"limit":   "5",
			},
			nil,
		},
		{
			"ConceptDetail",
			map[string]string{
				"ts_code": "000001.SZ",
				"limit":   "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "Concept":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.Concept(test.params, test.fields)
				}
			case "ConceptDetail":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.ConceptDetail(test.params, test.fields)
				}
			}

			result := runTest("Concept."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testThsAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"ThsDaily",
			map[string]string{
				"limit": "10",
			},
			[]string{"ts_code", "trade_date", "open", "high", "low", "close", "pre_close", "avg_price", "change", "pct_change", "vol", "turnover_rate"},
		},
		{
			"ThsMember",
			map[string]string{
				"index_code": "880001",
				"limit":      "5",
			},
			nil,
		},
		{
			"MoneyflowThs",
			map[string]string{
				"limit": "10",
			},
			[]string{"trade_date", "ts_code", "name", "pct_change", "latest", "net_amount", "net_d5_amount", "buy_lg_amount", "buy_lg_amount_rate", "buy_md_amount", "buy_md_amount_rate", "buy_sm_amount", "buy_sm_amount_rate"},
		},
		{
			"MoneyflowIndThs",
			map[string]string{
				"limit": "10",
			},
			[]string{"trade_date", "ts_code", "industry", "lead_stock", "close", "pct_change", "company_num", "pct_change_stock", "close_price", "net_buy_amount", "net_sell_amount", "net_amount"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "ThsDaily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.ThsDaily(test.params, test.fields)
				}
			case "ThsMember":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.ThsMember(test.params, test.fields)
				}
			case "MoneyflowThs":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.MoneyflowThs(test.params, test.fields)
				}
			case "MoneyflowIndThs":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.MoneyflowIndThs(test.params, test.fields)
				}
			}

			result := runTest("Ths."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testSwAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"CiDaily",
			map[string]string{
				"limit": "10",
			},
			[]string{"ts_code", "trade_date", "name", "open", "low", "high", "close", "change", "pct_change", "vol", "amount", "pe", "pb", "float_mv", "total_mv"},
		},
		{
			"SwDaily",
			map[string]string{
				"index_code": "801010",
				"limit":      "5",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "CiDaily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.CiDaily(test.params, test.fields)
				}
			case "SwDaily":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.SwDaily(test.params, test.fields)
				}
			}

			result := runTest("Sw."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testLimitAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"LimitList",
			map[string]string{
				"trade_date": "20240105",
				"limit":      "10",
			},
			nil,
		},
		{
			"STKLimit",
			map[string]string{
				"trade_date": "20240105",
				"limit":      "10",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "LimitList":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.LimitList(test.params, test.fields)
				}
			case "STKLimit":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Toplist.STKLimit(test.params, test.fields)
				}
			}

			result := runTest("Limit."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testResearchAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"CyqChips",
			map[string]string{
				"ts_code": "600000.SH",
				"limit":   "10",
			},
			nil,
		},
		{
			"StkSurv",
			map[string]string{
				"limit": "10",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "CyqChips":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Special.CyqChips(test.params, test.fields)
				}
			case "StkSurv":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Special.StkSurv(test.params, test.fields)
				}
			}

			result := runTest("Research."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testRepurchaseAPIs(t *testing.T, client *TuShare) {
	tests := []struct {
		name   string
		params map[string]string
		fields []string
	}{
		{
			"Repurchase",
			map[string]string{
				"limit": "10",
			},
			nil,
		},
		{
			"ShareFloat",
			map[string]string{
				"limit": "10",
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var apiFunc func() (*types.APIResponse, error)

			switch test.name {
			case "Repurchase":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.Repurchase(test.params, test.fields)
				}
			case "ShareFloat":
				apiFunc = func() (*types.APIResponse, error) {
					return client.Reference.ShareFloat(test.params, test.fields)
				}
			}

			result := runTest("Repurchase."+test.name, test.params, test.fields, apiFunc)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", test.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", test.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", test.name, result.DataCount, result.Duration)
			}
		})
	}
}

func maskToken(token string) string {
	if len(token) <= 8 {
		return "***"
	}
	return token[:4] + "****" + token[len(token)-4:]
}

func generateReport(t *testing.T) {
	report.Total = len(report.Results)
	for _, r := range report.Results {
		if r.Success {
			report.Passed++
		} else {
			report.Failed++
		}
	}

	// Save JSON report
	jsonData, _ := json.MarshalIndent(report, "", "  ")
	os.WriteFile("integration_test_report.json", jsonData, 0644)

	// Save markdown report with parameters
	saveMarkdownReport()

	// Print summary
	t.Log("\n" + strings.Repeat("=", 50))
	t.Log("           Integration Test Summary")
	t.Log(strings.Repeat("=", 50))
	t.Logf("📊 Total APIs Tested: %d", report.Total)
	t.Logf("✅ Passed: %d", report.Passed)
	t.Logf("❌ Failed: %d", report.Failed)
	t.Logf("📈 Success Rate: %.1f%%", float64(report.Passed)/float64(report.Total)*100)
	t.Log(strings.Repeat("=", 50))
}

func saveMarkdownReport() {
	successRate := float64(report.Passed) / float64(report.Total) * 100

	md := fmt.Sprintf(`# Tushare API Integration Test Report

**Test Date:** %s
**Token:** %s
**Total APIs:** %d

## Summary

| Metric | Count |
|--------|-------|
| ✅ Passed | %d |
| ❌ Failed | %d |
| 📊 Success Rate | %.1f%% |

---

## Detailed Results by Category

### 📈 Stock APIs

| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |
|----------|--------|-----------|----------|------------|--------|-------|
`,
		report.TestDate,
		report.Token,
		report.Total,
		report.Passed,
		report.Failed,
		successRate,
	)

	for _, r := range report.Results {
		if len(r.APIName) > 5 && r.APIName[:5] == "Stock" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[6:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### �� Index APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 5 && r.APIName[:5] == "Index" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[6:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### 📉 Market APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 6 && r.APIName[:6] == "Market" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[7:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### 💰 Finance APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 7 && r.APIName[:7] == "Finance" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[8:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### 🌊 Hsgt APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 4 && r.APIName[:4] == "Hsgt" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[5:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### 💸 Margin APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 6 && r.APIName[:6] == "Margin" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[7:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### ⚡ Realtime APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 8 && r.APIName[:8] == "Realtime" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[9:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += "\n### 💱 ETF APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if len(r.APIName) > 3 && r.APIName[:3] == "Etf" {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				r.APIName[4:], status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Pledge APIs
	md += "\n### 📋 Pledge APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Pledge.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Toplist APIs
	md += "\n### 🏆 Toplist APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Toplist.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Holder APIs
	md += "\n### 👥 Holder APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Holder.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Concept APIs
	md += "\n### 💡 Concept APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Concept.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Ths APIs
	md += "\n### 📊 Ths APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Ths.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Sw APIs
	md += "\n### 📈 Sw APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Sw.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Limit APIs
	md += "\n### ⛔ Limit APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Limit.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Research APIs
	md += "\n### 🔬 Research APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Research.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	// Repurchase APIs
	md += "\n### 🔄 Repurchase APIs\n\n| API Name | Status | Data Rows | Duration | Parameters | Fields | Notes |\n|----------|--------|-----------|----------|------------|--------|-------|\n"

	for _, r := range report.Results {
		if strings.HasPrefix(r.APIName, "Repurchase.") {
			status := "✅ PASS"
			if !r.Success {
				status = "❌ FAIL"
			} else if !r.HasData {
				status = "⚠️ NO DATA"
			}
			dataCount := fmt.Sprintf("%d", r.DataCount)
			if !r.HasData {
				dataCount = "-"
			}
			paramsStr := formatMap(r.Params)
			fieldsStr := "nil"
			if len(r.Fields) > 0 {
				fieldsStr = fmt.Sprintf("[%s]", strings.Join(r.Fields, ", "))
			}
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			parts := strings.Split(r.APIName, ".")
			apiName := parts[len(parts)-1]
			md += fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				apiName, status, dataCount, r.Duration, paramsStr, fieldsStr, notes)
		}
	}

	md += fmt.Sprintf("\n---\n\n*Generated on %s*\n\n*Note: This report is automatically generated by integration tests.*",
		report.TestDate)

	os.WriteFile("integration_test_report.md", []byte(md), 0644)
}

func formatMap(m map[string]string) string {
	if len(m) == 0 {
		return "{}"
	}

	pairs := make([]string, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, fmt.Sprintf("%s:%s", k, v))
	}
	return "{" + strings.Join(pairs, ", ") + "}"
}
