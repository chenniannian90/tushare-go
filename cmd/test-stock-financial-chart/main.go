package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_financial"
)

// TestResult 记录单个 API 的测试结果
type TestResult struct {
	APIName    string
	Success    bool
	DataCount  int
	Error      string
	SampleData interface{}
}

// ChartData 用于图表展示的数据结构
type ChartData struct {
	APIName    string
	DataCount  int
	BarLength  int
	Status     string
}

func main() {
	// 从环境变量获取 Token
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		fmt.Println("❌ 错误: 请设置 TUSHARE_TOKEN 环境变量")
		fmt.Println("\n使用方法:")
		fmt.Println("  TUSHARE_TOKEN=\"your_token\" go run main_chart.go")
		os.Exit(1)
	}

	// 创建客户端
	config, err := sdk.NewConfig(token)
	if err != nil {
		fmt.Printf("创建配置失败: %v\n", err)
		os.Exit(1)
	}
	client := sdk.NewClient(config)
	ctx := context.Background()

	fmt.Println("========================================")
	fmt.Println("  Tushare Stock Financial API 测试")
	fmt.Println("  📊 图表展示模式")
	fmt.Println("========================================")
	fmt.Println()

	// 测试所有 API
	results := testAllAPIs(ctx, client)

	// 生成图表展示
	generateCharts(results)

	// 生成详细报告
	generateDetailedReport(results)

	// 生成HTML报告
	generateHTMLReport(results)
}

func testAllAPIs(ctx context.Context, client *sdk.Client) []TestResult {
	var results []TestResult

	fmt.Println("🚀 开始测试所有 API...")

	// 1. Balancesheet (资产负债表)
	fmt.Println("\n[1/10] 测试 Balancesheet (资产负债表)...")
	result := testAPI("Balancesheet", func() (interface{}, error) {
		req := &stock_financial.BalancesheetRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.Balancesheet(ctx, client, req)
	})
	results = append(results, result)

	// 2. Cashflow (现金流量表)
	fmt.Println("[2/10] 测试 Cashflow (现金流量表)...")
	result = testAPI("Cashflow", func() (interface{}, error) {
		req := &stock_financial.CashflowRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.Cashflow(ctx, client, req)
	})
	results = append(results, result)

	// 3. DisclosureDate (财报披露日期)
	fmt.Println("[3/10] 测试 DisclosureDate (财报披露日期)...")
	result = testAPI("DisclosureDate", func() (interface{}, error) {
		req := &stock_financial.DisclosureDateRequest{
			EndDate: "20241231",
		}
		return stock_financial.DisclosureDate(ctx, client, req)
	})
	results = append(results, result)

	// 4. Dividend (分红送股)
	fmt.Println("[4/10] 测试 Dividend (分红送股)...")
	result = testAPI("Dividend", func() (interface{}, error) {
		req := &stock_financial.DividendRequest{
			TsCode: "000001.SZ",
		}
		return stock_financial.Dividend(ctx, client, req)
	})
	results = append(results, result)

	// 5. Express (业绩快报)
	fmt.Println("[5/10] 测试 Express (业绩快报)...")
	result = testAPI("Express", func() (interface{}, error) {
		req := &stock_financial.ExpressRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.Express(ctx, client, req)
	})
	results = append(results, result)

	// 6. FinaAudit (财务审计意见)
	fmt.Println("[6/10] 测试 FinaAudit (财务审计意见)...")
	result = testAPI("FinaAudit", func() (interface{}, error) {
		req := &stock_financial.FinaAuditRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.FinaAudit(ctx, client, req)
	})
	results = append(results, result)

	// 7. FinaIndicator (财务指标)
	fmt.Println("[7/10] 测试 FinaIndicator (财务指标)...")
	result = testAPI("FinaIndicator", func() (interface{}, error) {
		req := &stock_financial.FinaIndicatorRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.FinaIndicator(ctx, client, req)
	})
	results = append(results, result)

	// 8. FinaMainbz (主营业务构成)
	fmt.Println("[8/10] 测试 FinaMainbz (主营业务构成)...")
	result = testAPI("FinaMainbz", func() (interface{}, error) {
		req := &stock_financial.FinaMainbzRequest{
			TsCode: "000001.SZ",
			Period: "20241231",
			Type:   "P",
		}
		return stock_financial.FinaMainbz(ctx, client, req)
	})
	results = append(results, result)

	// 9. Forecast (业绩预告)
	fmt.Println("[9/10] 测试 Forecast (业绩预告)...")
	result = testAPI("Forecast", func() (interface{}, error) {
		req := &stock_financial.ForecastRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.Forecast(ctx, client, req)
	})
	results = append(results, result)

	// 10. Income (利润表)
	fmt.Println("[10/10] 测试 Income (利润表)...")
	result = testAPI("Income", func() (interface{}, error) {
		req := &stock_financial.IncomeRequest{
			TsCode:    "000001.SZ",
			StartDate: "20240101",
			EndDate:   "20241231",
		}
		return stock_financial.Income(ctx, client, req)
	})
	results = append(results, result)

	fmt.Println("\n✅ 所有测试完成!")
	return results
}

type apiFunc func() (interface{}, error)

func testAPI(apiName string, fn apiFunc) TestResult {
	result, err := fn()

	if err != nil {
		return TestResult{
			APIName: apiName,
			Success: false,
			Error:   err.Error(),
		}
	}

	// 使用反射获取切片长度
	dataCount := 0
	if result != nil {
		jsonData, _ := json.Marshal(result)
		var array []interface{}
		if err := json.Unmarshal(jsonData, &array); err == nil {
			dataCount = len(array)
		}
	}

	return TestResult{
		APIName:    apiName,
		Success:    true,
		DataCount:  dataCount,
		SampleData: result,
	}
}

func generateCharts(results []TestResult) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("📊 数据返回情况图表")
	fmt.Println(strings.Repeat("=", 80))

	// 找出最大值用于计算比例
	maxCount := 0
	for _, r := range results {
		if r.DataCount > maxCount {
			maxCount = r.DataCount
		}
	}

	// 按数据数量排序
	sortedResults := make([]TestResult, len(results))
	copy(sortedResults, results)
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].DataCount > sortedResults[j].DataCount
	})

	fmt.Println()
	for i, r := range sortedResults {
		status := "✅"
		if !r.Success {
			status = "❌"
		} else if r.DataCount == 0 {
			status = "⚠️ "
		}

		barLength := 0
		if r.Success && r.DataCount > 0 && maxCount > 0 {
			barLength = (r.DataCount * 50) / maxCount
		}

		fmt.Printf("%2d. %-20s ", i+1, r.APIName)
		fmt.Printf("%s ", status)

		if r.Success {
			if r.DataCount > 0 {
				fmt.Printf("%4d 条 ", r.DataCount)
				fmt.Printf("|%s|\n", strings.Repeat("█", barLength))
			} else {
				fmt.Printf("  空  |\n")
			}
		} else {
			fmt.Printf(" 错误 |\n")
			fmt.Printf("    %s\n", r.Error)
		}
	}

	fmt.Println()
	fmt.Println("图例:")
	fmt.Println("  ✅ 成功返回数据")
	fmt.Println("  ⚠️  成功但无数据")
	fmt.Println("  ❌ 调用失败")
	fmt.Println("  █  数据量（每个█代表约" + fmt.Sprintf("%d", maxCount/50+1) + "条记录)")
}

func generateDetailedReport(results []TestResult) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("📋 详细测试报告")
	fmt.Println(strings.Repeat("=", 80))

	successCount := 0
	failCount := 0
	emptyCount := 0
	totalRecords := 0

	for _, r := range results {
		if r.Success {
			successCount++
			if r.DataCount == 0 {
				emptyCount++
			} else {
				totalRecords += r.DataCount
			}
		} else {
			failCount++
		}
	}

	fmt.Printf("\n📈 统计摘要:\n")
	fmt.Printf("  • 测试总数: %d\n", len(results))
	fmt.Printf("  • 成功: %d (%.1f%%)\n", successCount, float64(successCount)*100/float64(len(results)))
	fmt.Printf("  • 失败: %d (%.1f%%)\n", failCount, float64(failCount)*100/float64(len(results)))
	fmt.Printf("  • 空数据: %d (%.1f%%)\n", emptyCount, float64(emptyCount)*100/float64(len(results)))
	fmt.Printf("  • 总记录数: %d\n", totalRecords)

	fmt.Printf("\n📊 各API返回记录数:\n")
	for _, r := range results {
		if r.Success {
			if r.DataCount > 0 {
				fmt.Printf("  • %-20s: %4d 条\n", r.APIName, r.DataCount)
			} else {
				fmt.Printf("  • %-20s:     0 条 (空)\n", r.APIName)
			}
		} else {
			fmt.Printf("  • %-20s: 失败 - %s\n", r.APIName, r.Error)
		}
	}
}

func generateHTMLReport(results []TestResult) {
	html := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tushare Stock Financial API 测试报告</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            border-radius: 8px;
            padding: 30px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        h1 {
            color: #2c3e50;
            border-bottom: 3px solid #3498db;
            padding-bottom: 10px;
        }
        .summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin: 30px 0;
        }
        .summary-card {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 20px;
            border-radius: 8px;
            text-align: center;
        }
        .summary-card.success { background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%); }
        .summary-card.failure { background: linear-gradient(135deg, #eb3349 0%, #f45c43 100%); }
        .summary-card.empty { background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); }
        .summary-card.total { background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); }
        .summary-number {
            font-size: 2em;
            font-weight: bold;
            margin: 10px 0;
        }
        .chart-container {
            margin: 30px 0;
        }
        .chart-bar {
            margin: 15px 0;
        }
        .chart-label {
            display: inline-block;
            width: 150px;
            font-weight: bold;
        }
        .chart-bar-fill {
            display: inline-block;
            height: 30px;
            background: linear-gradient(90deg, #3498db, #2ecc71);
            border-radius: 4px;
            transition: width 0.5s ease;
        }
        .chart-bar-fill.error {
            background: linear-gradient(90deg, #e74c3c, #c0392b);
        }
        .chart-bar-fill.empty {
            background: linear-gradient(90deg, #f39c12, #e67e22);
        }
        .chart-value {
            margin-left: 10px;
            font-weight: bold;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin: 20px 0;
        }
        th, td {
            padding: 12px;
            text-align: left;
            border-bottom: 1px solid #ddd;
        }
        th {
            background-color: #3498db;
            color: white;
        }
        tr:hover {
            background-color: #f5f5f5;
        }
        .status {
            padding: 4px 8px;
            border-radius: 4px;
            font-size: 0.9em;
            font-weight: bold;
        }
        .status.success {
            background-color: #2ecc71;
            color: white;
        }
        .status.error {
            background-color: #e74c3c;
            color: white;
        }
        .status.empty {
            background-color: #f39c12;
            color: white;
        }
        .timestamp {
            text-align: right;
            color: #7f8c8d;
            font-size: 0.9em;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>📊 Tushare Stock Financial API 测试报告</h1>
        <p class="timestamp">生成时间: {{timestamp}}</p>
        <p><strong>测试股票:</strong> 000001.SZ (平安银行)</p>
        <p><strong>测试时间范围:</strong> 2024-01-01 至 2024-12-31</p>

        <div class="summary">
            <div class="summary-card success">
                <div>成功</div>
                <div class="summary-number">{{success}}</div>
            </div>
            <div class="summary-card failure">
                <div>失败</div>
                <div class="summary-number">{{failure}}</div>
            </div>
            <div class="summary-card empty">
                <div>空数据</div>
                <div class="summary-number">{{empty}}</div>
            </div>
            <div class="summary-card total">
                <div>总记录数</div>
                <div class="summary-number">{{totalRecords}}</div>
            </div>
        </div>

        <div class="chart-container">
            <h2>📈 数据返回量图表</h2>
            {{chartContent}}
        </div>

        <h2>📋 详细结果</h2>
        <table>
            <thead>
                <tr>
                    <th>API 名称</th>
                    <th>状态</th>
                    <th>记录数</th>
                    <th>说明</th>
                </tr>
            </thead>
            <tbody>
                {{tableRows}}
            </tbody>
        </table>

        <h2>📝 API 说明</h2>
        <table>
            <thead>
                <tr>
                    <th>API</th>
                    <th>数据可用性</th>
                    <th>说明</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Balancesheet</td>
                    <td>⭐⭐⭐⭐⭐</td>
                    <td>资产负债表 - 几乎所有公司都有完整数据</td>
                </tr>
                <tr>
                    <td>Income</td>
                    <td>⭐⭐⭐⭐⭐</td>
                    <td>利润表 - 几乎所有公司都有完整数据</td>
                </tr>
                <tr>
                    <td>Cashflow</td>
                    <td>⭐⭐⭐⭐⭐</td>
                    <td>现金流量表 - 几乎所有公司都有完整数据</td>
                </tr>
                <tr>
                    <td>Dividend</td>
                    <td>⭐⭐⭐⭐</td>
                    <td>分红数据 - 取决于公司分红政策</td>
                </tr>
                <tr>
                    <td>FinaIndicator</td>
                    <td>⭐⭐⭐⭐</td>
                    <td>财务指标 - 大部分公司都有数据</td>
                </tr>
                <tr>
                    <td>FinaAudit</td>
                    <td>⭐⭐⭐⭐</td>
                    <td>审计意见 - 每年更新一次</td>
                </tr>
                <tr>
                    <td>DisclosureDate</td>
                    <td>⭐⭐⭐</td>
                    <td>财报披露日期 - 季节性数据</td>
                </tr>
                <tr>
                    <td>Forecast</td>
                    <td>⭐⭐⭐</td>
                    <td>业绩预告 - 业绩波动时发布</td>
                </tr>
                <tr>
                    <td>Express</td>
                    <td>⭐⭐</td>
                    <td>业绩快报 - 非所有公司发布</td>
                </tr>
                <tr>
                    <td>FinaMainbz</td>
                    <td>⭐⭐⭐</td>
                    <td>主营业务构成 - 多元化公司数据多</td>
                </tr>
            </tbody>
        </table>
    </div>
</body>
</html>`

	// 计算统计数据
	successCount := 0
	failCount := 0
	emptyCount := 0
	totalRecords := 0

	for _, r := range results {
		if r.Success {
			successCount++
			if r.DataCount == 0 {
				emptyCount++
			} else {
				totalRecords += r.DataCount
			}
		} else {
			failCount++
		}
	}

	// 生成图表
	maxCount := 0
	for _, r := range results {
		if r.DataCount > maxCount {
			maxCount = r.DataCount
		}
	}

	chartHTML := ""
	sortedResults := make([]TestResult, len(results))
	copy(sortedResults, results)
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].DataCount > sortedResults[j].DataCount
	})

	for _, r := range sortedResults {
		barClass := "chart-bar-fill"
		if !r.Success {
			barClass += " error"
		} else if r.DataCount == 0 {
			barClass += " empty"
		}

		barWidth := "0%"
		if r.Success && r.DataCount > 0 && maxCount > 0 {
			barWidth = fmt.Sprintf("%d%%", (r.DataCount*100)/maxCount)
		}

		valueStr := "错误"
		if r.Success {
			valueStr = fmt.Sprintf("%d 条", r.DataCount)
		}

		chartHTML += fmt.Sprintf(`
            <div class="chart-bar">
                <span class="chart-label">%s</span>
                <span class="%s" style="width: %s"></span>
                <span class="chart-value">%s</span>
            </div>`, r.APIName, barClass, barWidth, valueStr)
	}

	// 生成表格行
	tableRows := ""
	for _, r := range results {
		status := "<span class=\"status success\">成功</span>"
		description := fmt.Sprintf("%d 条记录", r.DataCount)

		if !r.Success {
			status = "<span class=\"status error\">失败</span>"
			description = r.Error
		} else if r.DataCount == 0 {
			status = "<span class=\"status empty\">空数据</span>"
			description = "该时间段无数据"
		}

		tableRows += fmt.Sprintf(`
                <tr>
                    <td><strong>%s</strong></td>
                    <td>%s</td>
                    <td>%d</td>
                    <td>%s</td>
                </tr>`, r.APIName, status, r.DataCount, description)
	}

	// 替换模板变量
	html = strings.Replace(html, "{{timestamp}}", fmt.Sprintf("%v", results[0].SampleData), 1)
	html = strings.Replace(html, "{{success}}", fmt.Sprintf("%d", successCount), 1)
	html = strings.Replace(html, "{{failure}}", fmt.Sprintf("%d", failCount), 1)
	html = strings.Replace(html, "{{empty}}", fmt.Sprintf("%d", emptyCount), 1)
	html = strings.Replace(html, "{{totalRecords}}", fmt.Sprintf("%d", totalRecords), 1)
	html = strings.Replace(html, "{{chartContent}}", chartHTML, 1)
	html = strings.Replace(html, "{{tableRows}}", tableRows, 1)

	// 写入文件
	filename := "test_report.html"
	err := os.WriteFile(filename, []byte(html), 0644)
	if err != nil {
		fmt.Printf("\n⚠️  生成HTML报告失败: %v\n", err)
		return
	}

	fmt.Printf("\n✅ HTML报告已生成: %s\n", filename)
	fmt.Printf("   请在浏览器中打开查看可视化报告\n")
}
