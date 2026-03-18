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
	APIName   string `json:"api_name"`
	Success   bool   `json:"success"`
	Error     string `json:"error,omitempty"`
	HasData   bool   `json:"has_data"`
	DataCount int    `json:"data_count,omitempty"`
	Duration  string `json:"duration"`
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

func runTest(apiName string, testFunc func() (*types.APIResponse, error)) TestResult {
	start := time.Now()
	result := TestResult{
		APIName: apiName,
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

	// Generate report
	generateReport(t)
}

func testStockAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"StockBasic",
			func() (*types.APIResponse, error) {
				return client.Stock.StockBasic(map[string]string{
					"list_status": "L",
					"exchange":    "SSE",
					"limit":       "5",
				}, nil)
			},
		},
		{
			"TradeCal",
			func() (*types.APIResponse, error) {
				return client.Stock.TradeCal(map[string]string{
					"exchange":  "SSE",
					"start_date": "20240101",
					"end_date":   "20240110",
				}, nil)
			},
		},
		{
			"HSConst",
			func() (*types.APIResponse, error) {
				return client.Stock.HSConst(map[string]string{
					"hs_type": "SH",
				}, nil)
			},
		},
		{
			"StockCompany",
			func() (*types.APIResponse, error) {
				return client.Stock.StockCompany(map[string]string{
					"exchange": "SSE",
					"limit":    "5",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Stock."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testIndexAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"IndexDaily",
			func() (*types.APIResponse, error) {
				return client.Index.IndexDaily(map[string]string{
					"ts_code":    "000001.SH",
					"start_date": "20240101",
					"end_date":   "20240105",
				}, nil)
			},
		},
		{
			"IndexBasic",
			func() (*types.APIResponse, error) {
				return client.Index.IndexBasic(map[string]string{
					"market": "SSE",
					"limit":  "5",
				}, nil)
			},
		},
		{
			"IndexWeight",
			func() (*types.APIResponse, error) {
				return client.Index.IndexWeight(map[string]string{
					"index_code": "000001.SH",
					"start_date": "20240101",
					"end_date":   "20240102",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Index."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testMarketAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"Daily",
			func() (*types.APIResponse, error) {
				return client.Market.Daily(map[string]string{
					"trade_date": "20240101",
				}, nil)
			},
		},
		{
			"DailyBasic",
			func() (*types.APIResponse, error) {
				return client.Market.DailyBasic(map[string]string{
					"trade_date": "20240101",
					"limit":      "5",
				}, nil)
			},
		},
		{
			"MoneyFlow",
			func() (*types.APIResponse, error) {
				return client.Market.MoneyFlow(map[string]string{
					"ts_code":    "000001.SZ",
					"trade_date": "20240101",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Market."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testFinanceAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"Income",
			func() (*types.APIResponse, error) {
				return client.Finance.Income(map[string]string{
					"ts_code":    "000001.SZ",
					"start_date": "20230101",
					"end_date":   "20231231",
					"limit":      "1",
				}, nil)
			},
		},
		{
			"BalanceSheet",
			func() (*types.APIResponse, error) {
				return client.Finance.BalanceSheet(map[string]string{
					"ts_code":    "000001.SZ",
					"start_date": "20230101",
					"end_date":   "20231231",
					"limit":      "1",
				}, nil)
			},
		},
		{
			"FinaIndicator",
			func() (*types.APIResponse, error) {
				return client.Finance.FinaIndicator(map[string]string{
					"ts_code":    "000001.SZ",
					"start_date": "20230101",
					"end_date":   "20231231",
					"limit":      "1",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Finance."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testHsgtAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"MoneyflowHsgt",
			func() (*types.APIResponse, error) {
				return client.Hsgt.MoneyflowHsgt(map[string]string{
					"trade_date": "20240101",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Hsgt."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testMarginAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"Margin",
			func() (*types.APIResponse, error) {
				return client.Margin.Margin(map[string]string{
					"trade_date": "20240101",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Margin."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
			}
		})
	}
}

func testRealtimeAPIs(t *testing.T, client *TuShare) {
	apis := []struct {
		name string
		test func() (*types.APIResponse, error)
	}{
		{
			"RTK",
			func() (*types.APIResponse, error) {
				return client.Realtime.RTK(map[string]string{}, nil)
			},
		},
		{
			"RealTimeQuote",
			func() (*types.APIResponse, error) {
				return client.Realtime.RealTimeQuote(map[string]string{
					"ts_code": "000001.SZ",
					"src":     "sz",
				}, nil)
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			result := runTest("Realtime."+api.name, api.test)
			report.Results = append(report.Results, result)

			if !result.Success {
				t.Errorf("❌ %s: %s", api.name, result.Error)
			} else if !result.HasData {
				t.Logf("⚠️  %s: No data", api.name)
			} else {
				t.Logf("✅ %s: %d rows (%s)", api.name, result.DataCount, result.Duration)
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

	// Save markdown report
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

| API Name | Status | Data Rows | Duration | Notes |
|----------|--------|-----------|----------|-------|
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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[6:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### 📊 Index APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[6:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### 📉 Market APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[7:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### 💰 Finance APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[8:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### 🌊 Hsgt APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[5:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### 💸 Margin APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[7:], status, dataCount, r.Duration, notes)
		}
	}

	md += "\n### ⚡ Realtime APIs\n\n| API Name | Status | Data Rows | Duration | Notes |\n|----------|--------|-----------|----------|-------|\n"

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
			notes := ""
			if r.Error != "" {
				notes = r.Error
			}
			md += fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
				r.APIName[9:], status, dataCount, r.Duration, notes)
		}
	}

	md += fmt.Sprintf("\n---\n\n*Generated on %s*\n\n*Note: This report is automatically generated by integration tests.*\n",
		report.TestDate)

	os.WriteFile("integration_test_report.md", []byte(md), 0644)
}
