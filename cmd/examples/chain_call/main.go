package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/apis"
	stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
)

func main() {
	// Get Tushare token from environment
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	// Create SDK client
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	client := sdk.NewClient(config)

	fmt.Println("========================================")
	fmt.Println("方式 1: 直接调用（原有方式）")
	fmt.Println("========================================")

	// 方式 1: 直接调用 API 函数
	fmt.Println("\n=== Example 1: 直接调用 ===")
	limitList1, err := stockboard.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		fmt.Printf("Found %d entries (方式1)\n", len(limitList1))
	}

	fmt.Println("\n========================================")
	fmt.Println("方式 2: 使用 SDK 内置的链式调用")
	fmt.Println("========================================")

	// 方式 2: 使用 SDK 内置的链式调用 + 通用 CallAPI 方法
	fmt.Println("\n=== Example 2: 链式调用 + CallAPI ===")
	var result2 struct {
		Fields []string                 `json:"fields"`
		Items  []map[string]interface{} `json:"items"`
	}
	err = client.StockBoard().CallAPI(
		context.Background(),
		"top_list",
		map[string]interface{}{},
		[]string{},
		&result2,
	)
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		fmt.Printf("Found %d entries (方式2)\n", len(result2.Items))
	}

	fmt.Println("\n========================================")
	fmt.Println("方式 3: 使用 apis 包的类型化方法")
	fmt.Println("========================================")

	// 方式 3: 使用 apis 包提供的类型化方法
	fmt.Println("\n=== Example 3: apis 包类型化方法 ===")
	limitList3, err := apis.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		fmt.Printf("Found %d entries (方式3)\n", len(limitList3))
	}

	fmt.Println("\n========================================")
	fmt.Println("三种调用方式对比")
	fmt.Println("========================================")

	fmt.Println("\n📝 方式 1: 直接调用")
	fmt.Println("   优点: 简单直接，无需额外包装")
	fmt.Println("   缺点: 需要导入具体的 API 包")
	fmt.Println("   适合: 快速脚本，API 调用较少")
	fmt.Println("\n   示例:")
	fmt.Println("   import stockboard \".../pkg/sdk/api/stock/stock_board\"")
	fmt.Println("   stockboard.TopList(ctx, client, req)")

	fmt.Println("\n🔗 方式 2: SDK 链式调用 + CallAPI")
	fmt.Println("   优点: 代码组织清晰，无需导入具体包")
	fmt.Println("   缺点: 需要手动处理结果解析")
	fmt.Println("   适合: 已知 API 名称和参数格式")
	fmt.Println("\n   示例:")
	fmt.Println("   client.StockBoard().CallAPI(ctx, \"top_list\", params, fields, &result)")
	fmt.Println("   client.StockMarket().CallAPI(ctx, \"daily\", params, fields, &result)")

	fmt.Println("\n🎯 方式 3: apis 包类型化方法（推荐）")
	fmt.Println("   优点: 类型安全，IDE 提示友好，自动处理结果解析")
	fmt.Println("   缺点: 需要导入 apis 包")
	fmt.Println("   适合: 大型项目，需要类型安全和良好的 IDE 支持")
	fmt.Println("\n   示例:")
	fmt.Println("   import \".../pkg/sdk/apis\"")
	fmt.Println("   apis.TopList(ctx, client, req)")
	fmt.Println("   apis.Daily(ctx, client, req)")

	fmt.Println("\n========================================")
	fmt.Println("实际使用建议")
	fmt.Println("========================================")

	fmt.Println("\n🚀 推荐做法: 混合使用")
	fmt.Println("\n在项目中创建自己的 API 包装层:")
	fmt.Println(`
	// myapp/api/stock.go
	package api

	import (
	    "github.com/chenniannian90/tushare-go/pkg/sdk"
	    "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
	)

	type StockAPI struct {
	    client *sdk.Client
	}

	func NewStockAPI(client *sdk.Client) *StockAPI {
	    return &StockAPI{client: client}
	}

	// 使用 apis 包提供的方法
	func (s *StockAPI) GetTopList(ctx) ([]TopListData, error) {
	    req := &stockboard.TopListRequest{...}
	    return apis.TopList(ctx, s.client, req)
	}

	// 或直接使用 CallAPI
	func (s *StockAPI) GetCustomData(ctx) ([]Data, error) {
	    var result struct {...}
	    err := s.client.StockBoard().CallAPI(ctx, "api_name", params, fields, &result)
	    // 处理 result...
	    return data, err
	}
	`)

	fmt.Println("\n========================================")
	fmt.Println("可用的 API 分类")
	fmt.Println("========================================")

	fmt.Println("\n📊 SDK 内置分类:")
	fmt.Println("   client.StockBoard()   - 股票板块")
	fmt.Println("   client.StockMarket()  - 股票市场")
	fmt.Println("   client.StockBasic()   - 股票基础")
	fmt.Println("   client.StockFinancial() - 财务数据")
	fmt.Println("   client.Index()        - 指数")
	fmt.Println("   client.Futures()      - 期货")
	fmt.Println("   client.Fund()         - 基金")
	fmt.Println("   client.HKStock()      - 港股")
	fmt.Println("   client.Bond()         - 债券")
	fmt.Println("   client.ETF()          - ETF")
	fmt.Println("   client.Forex()        - 外汇")
	fmt.Println("   client.Options()      - 期权")
	fmt.Println("   client.Spot()         - 现货")
	fmt.Println("   client.USStock()      - 美股")
	fmt.Println("   client.Wealth()       - 财富管理")
	fmt.Println("   client.Industry()     - 行业经济")
	fmt.Println("   client.LLMCorpus()    - 大模型语料")
	fmt.Println("   client.Macro()        - 宏观经济")

	fmt.Println("\n📦 apis 包提供的方法:")
	fmt.Println("   apis.TopList()       - 龙虎榜")
	fmt.Println("   apis.LimitList()     - 涨跌停")
	fmt.Println("   apis.Daily()         - 日线行情")
	fmt.Println("   apis.DailyBasic()    - 每日基本面")
	fmt.Println("   apis.TradeCal()      - 交易日历")
	fmt.Println("   apis.Income()        - 利润表")
	fmt.Println("   apis.Balancesheet()  - 资产负债表")
	fmt.Println("   apis.Dividend()      - 分红数据")
	fmt.Println("   ...更多方法持续添加中")

	fmt.Println("\n========================================")
	fmt.Println("完整示例")
	fmt.Println("========================================")

	fmt.Println("\n💡 推荐的使用方式:")
	fmt.Println(`
	package main

	import (
	    "context"
	    "github.com/chenniannian90/tushare-go/pkg/sdk"
	    "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
	)

	func main() {
	    client := sdk.NewClient(config)

	    // 方式 A: 使用 apis 包（推荐）
	    list, _ := apis.TopList(ctx, client, req)

	    // 方式 B: 使用 CallAPI
	    var result struct{...}
	    client.StockBoard().CallAPI(ctx, "top_list", params, fields, &result)
	}
	`)

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
