package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/api/stock_financial"
)

// TestCase 定义测试用例结构
type TestCase struct {
	Name      string
	ToolFunc  func(ctx context.Context, client *sdk.Client) (interface{}, error)
	WantError bool
}

func main() {
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
	ctx := context.Background()

	// 定义测试用例
	testCases := []TestCase{
		{
			Name: "测试资产负债表 (Balancesheet)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.BalancesheetRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.Balancesheet(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试现金流量表 (Cashflow)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.CashflowRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.Cashflow(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试财报披露日期 (DisclosureDate)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.DisclosureDateRequest{
					EndDate: "20241231",
				}
				return stock_financial.DisclosureDate(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试分红送股数据 (Dividend)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.DividendRequest{
					TsCode: "000001.SZ",
				}
				return stock_financial.Dividend(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试业绩快报 (Express)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.ExpressRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.Express(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试财务审计意见 (FinaAudit)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.FinaAuditRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.FinaAudit(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试财务指标 (FinaIndicator)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.FinaIndicatorRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.FinaIndicator(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试主营业务构成 (FinaMainbz)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.FinaMainbzRequest{
					TsCode: "000001.SZ",
					Period: "20241231",
					Type:   "P",
				}
				return stock_financial.FinaMainbz(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试业绩预告 (Forecast)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.ForecastRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.Forecast(ctx, client, req)
			},
			WantError: false,
		},
		{
			Name: "测试利润表 (Income)",
			ToolFunc: func(ctx context.Context, client *sdk.Client) (interface{}, error) {
				req := &stock_financial.IncomeRequest{
					TsCode:    "000001.SZ",
					StartDate: "20240101",
					EndDate:   "20241231",
				}
				return stock_financial.Income(ctx, client, req)
			},
			WantError: false,
		},
	}

	// 运行测试
	passed := 0
	failed := 0

	fmt.Println("========================================")
	fmt.Println("   Tushare Stock Financial API 测试")
	fmt.Println("========================================")
	fmt.Println()

	for i, tc := range testCases {
		fmt.Printf("[%d/%d] 运行测试: %s\n", i+1, len(testCases), tc.Name)

		result, err := tc.ToolFunc(ctx, client)

		if err != nil {
			if tc.WantError {
				fmt.Printf("  ✅ 预期错误: %v\n", err)
				passed++
			} else {
				fmt.Printf("  ❌ 意外错误: %v\n", err)
				failed++
			}
		} else {
			// 格式化输出结果
			jsonData, _ := json.MarshalIndent(result, "", "  ")

			// 只显示前 500 字符
			output := string(jsonData)
			if len(output) > 500 {
				output = output[:500] + "\n  ... (截断)"
			}

			fmt.Printf("  ✅ 成功\n")
			fmt.Printf("  返回数据预览:\n  %s\n", output)
			passed++
		}
		fmt.Println()
	}

	// 打印总结
	fmt.Println("========================================")
	fmt.Printf("测试完成! 通过: %d, 失败: %d, 总计: %d\n", passed, failed, len(testCases))
	fmt.Println("========================================")

	if failed > 0 {
		os.Exit(1)
	}
}