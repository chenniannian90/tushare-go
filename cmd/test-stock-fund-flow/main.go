package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_fund_flow"
)

// TestResult 测试结果
type TestResult struct {
	Name      string      `json:"name"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Count     int         `json:"count,omitempty"`
}

// TestReport 测试报告
type TestReport struct {
	TotalTests int         `json:"total_tests"`
	Passed     int         `json:"passed"`
	Failed     int         `json:"failed"`
	Skipped    int         `json:"skipped"`
	Results    []TestResult `json:"results"`
}

func main() {
	ctx := context.Background()

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

	// 定义测试用例
	testCases := []struct {
		Name      string
		TestFunc  func(context.Context, *sdk.Client) (interface{}, error)
		Enabled   bool
	}{
		{
			Name: "moneyflow - 个股资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowRequest{
					TsCode:    "600000.SH",
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.Moneyflow(ctx, client, req)
			},
			Enabled: true,
		},
		{
			Name: "moneyflow_hsgt - 沪深股通资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowHsgtRequest{
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowHsgt(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_ths - 个股沪深股通资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowThsRequest{
					TsCode:    "600000.SH",
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowThs(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_cnt_ths - 沪深股通成份股资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowCntThsRequest{
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowCntThs(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_dc - 大单资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowDcRequest{
					TsCode:    "600000.SH",
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowDc(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_ind_dc - 行业资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowIndDcRequest{
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowIndDc(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_ind_ths - 行业沪深股通资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowIndThsRequest{
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowIndThs(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
		{
			Name: "moneyflow_mkt_dc - 市场大单资金流向",
			TestFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_fund_flow.MoneyflowMktDcRequest{
					StartDate: "20240101",
					EndDate:   "20240110",
				}
				return stock_fund_flow.MoneyflowMktDc(ctx, client, req)
			},
			Enabled: false, // 注释掉的工具
		},
	}

	// 运行测试
	report := TestReport{
		TotalTests: len(testCases),
		Results:    make([]TestResult, 0, len(testCases)),
	}

	fmt.Println("========================================")
	fmt.Println("   Stock Fund Flow MCP Tools 测试")
	fmt.Println("========================================")
	fmt.Println()

	enabledCount := 0
	for i, tc := range testCases {
		result := TestResult{Name: tc.Name}

		if !tc.Enabled {
			result.Success = false
			result.Error = "工具已禁用（在 registry.go 中被注释）"
			report.Skipped++
			fmt.Printf("[%d/%d] ⏭️  跳过: %s\n", i+1, len(testCases), tc.Name)
			fmt.Printf("    原因: %s\n", result.Error)
			report.Results = append(report.Results, result)
			fmt.Println()
			continue
		}

		enabledCount++
		fmt.Printf("[%d/%d] 🧪 测试: %s\n", i+1, len(testCases), tc.Name)

		data, err := tc.TestFunc(ctx, client)
		if err != nil {
			result.Success = false
			result.Error = err.Error()
			report.Failed++
			fmt.Printf("    ❌ 失败: %v\n", err)
		} else {
			result.Success = true
			result.Data = data

			// 尝试获取数据数量
			var count int
			if arr, ok := data.([]interface{}); ok {
				count = len(arr)
				result.Count = count

				// 显示第一条数据示例
				if count > 0 {
					if firstItem, ok := arr[0].(map[string]interface{}); ok {
						if jsonBytes, err := json.MarshalIndent(firstItem, "    ", "  "); err == nil {
							fmt.Printf("    数据示例:\n%s\n", string(jsonBytes))
						}
					}
				}
			}
			report.Passed++
			fmt.Printf("    ✅ 成功: 获取到 %d 条数据\n", count)
		}

		report.Results = append(report.Results, result)
		fmt.Println()
	}

	// 打印总结
	fmt.Println("========================================")
	fmt.Printf("测试完成!\n")
	fmt.Printf("已启用工具: %d\n", enabledCount)
	fmt.Printf("通过: %d, 失败: %d, 跳过: %d, 总计: %d\n",
		report.Passed, report.Failed, report.Skipped, report.TotalTests)
	fmt.Println("========================================")

	// 保存报告
	reportJSON, _ := json.MarshalIndent(report, "", "  ")
	reportPath := "stock_fund_flow_mcp_test_report.json"
	if err := os.WriteFile(reportPath, reportJSON, 0644); err != nil {
		log.Printf("无法保存报告文件: %v", err)
	} else {
		fmt.Printf("\n报告已保存到: %s\n", reportPath)
	}

	if report.Failed > 0 {
		os.Exit(1)
	}
}