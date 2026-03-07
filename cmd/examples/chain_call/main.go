package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
	stockmarket "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_market"
	stockbasic "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_basic"
)

// ==================== 链式调用包装器 ====================
// 这个示例展示如何创建链式调用的 API 客户端

// APIClient 提供链式调用方式的 API 客户端
type APIClient struct {
	*sdk.Client
}

// NewAPIClient 创建新的 API 客户端
func NewAPIClient(client *sdk.Client) *APIClient {
	return &APIClient{Client: client}
}

// StockBoard 返回股票板块 API
func (c *APIClient) StockBoard() *StockBoardAPI {
	return &StockBoardAPI{client: c.Client}
}

// StockMarket 返回股票市场数据 API
func (c *APIClient) StockMarket() *StockMarketAPI {
	return &StockMarketAPI{client: c.Client}
}

// StockBasic 返回股票基础信息 API
func (c *APIClient) StockBasic() *StockBasicAPI {
	return &StockBasicAPI{client: c.Client}
}

// ==================== 股票板块 API ====================

// StockBoardAPI 股票板块相关 API
type StockBoardAPI struct {
	client *sdk.Client
}

// TopList 获取龙虎榜每日统计
func (s *StockBoardAPI) TopList(ctx context.Context, req *stockboard.TopListRequest) ([]stockboard.TopListItem, error) {
	return stockboard.TopList(ctx, s.client, req)
}

// LimitList 获取涨跌停和炸板数据
func (s *StockBoardAPI) LimitList(ctx context.Context, req *stockboard.LimitListRequest) ([]stockboard.LimitListItem, error) {
	return stockboard.LimitList(ctx, s.client, req)
}

// DragonList 获取游资交易每日明细
func (s *StockBoardAPI) DragonList(ctx context.Context, req *stockboard.DragonListRequest) ([]stockboard.DragonListItem, error) {
	return stockboard.DragonList(ctx, s.client, req)
}

// ThsConcept 获取同花顺概念板块
func (s *StockBoardAPI) ThsConcept(ctx context.Context, req *stockboard.ThsConceptRequest) ([]stockboard.ThsConceptItem, error) {
	return stockboard.ThsConcept(ctx, s.client, req)
}

// EmHot 获取东方财富App热榜
func (s *StockBoardAPI) EmHot(ctx context.Context, req *stockboard.EmHotRequest) ([]stockboard.EmHotItem, error) {
	return stockboard.EmHot(ctx, s.client, req)
}

// ==================== 股票市场数据 API ====================

// StockMarketAPI 股票市场数据相关 API
type StockMarketAPI struct {
	client *sdk.Client
}

// Daily 获取日线行情
func (s *StockMarketAPI) Daily(ctx context.Context, req *stockmarket.DailyRequest) ([]stockmarket.DailyItem, error) {
	return stockmarket.Daily(ctx, s.client, req)
}

// DailyBasic 获取每日基本面指标
func (s *StockMarketAPI) DailyBasic(ctx context.Context, req *stockmarket.DailyBasicRequest) ([]stockmarket.DailyBasicItem, error) {
	return stockmarket.DailyBasic(ctx, s.client, req)
}

// ==================== 股票基础信息 API ====================

// StockBasicAPI 股票基础信息相关 API
type StockBasicAPI struct {
	client *sdk.Client
}

// TradeCal 获取交易日历
func (s *StockBasicAPI) TradeCal(ctx context.Context, req *stockbasic.TradeCalRequest) ([]stockbasic.TradeCalItem, error) {
	return stockbasic.TradeCal(ctx, s.client, req)
}

// StockBasic 获取股票列表
func (s *StockBasicAPI) StockBasic(ctx context.Context, req *stockbasic.StockBasicRequest) ([]stockbasic.StockBasicItem, error) {
	return stockbasic.StockBasic(ctx, s.client, req)
}

// ==================== 示例代码 ====================

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
	// 创建链式调用客户端
	apiClient := NewAPIClient(client)

	fmt.Println("======================================")
	fmt.Println("方式 1: 直接调用 API 函数（原有方式）")
	fmt.Println("======================================")

	// 方式 1: 直接调用 API 函数
	fmt.Println("\n=== Example 1: 直接调用 ===")
	limitList1, err := stockboard.TopList(context.Background(), client, &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		fmt.Printf("Found %d top list entries (方式1)\n", len(limitList1))
	}

	fmt.Println("\n======================================")
	fmt.Println("方式 2: 使用链式调用（推荐）")
	fmt.Println("======================================")

	// 方式 2: 使用链式调用
	fmt.Println("\n=== Example 2: 链式调用 ===")
	limitList2, err := apiClient.StockBoard().TopList(context.Background(), &stockboard.TopListRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	} else {
		fmt.Printf("Found %d top list entries (方式2)\n", len(limitList2))
	}

	fmt.Println("\n=== Example 3: 链式调用多个 API ===")
	// 可以连续调用多个 API
	_, err = apiClient.StockBoard().LimitList(context.Background(), &stockboard.LimitListRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	_, err = apiClient.StockBasic().TradeCal(context.Background(), &stockbasic.TradeCalRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	_, err = apiClient.StockMarket().Daily(context.Background(), &stockmarket.DailyRequest{})
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	fmt.Println("\n======================================")
	fmt.Println("链式调用的优势")
	fmt.Println("======================================")

	fmt.Println("\n✅ 更好的代码组织:")
	fmt.Println("   apiClient.StockBoard().TopList()")
	fmt.Println("   apiClient.StockBoard().LimitList()")
	fmt.Println("   apiClient.StockBoard().EmHot()")
	fmt.Println("   apiClient.StockMarket().Daily()")
	fmt.Println("   apiClient.StockBasic().TradeCal()")

	fmt.Println("\n✅ 更清晰的 API 分类:")
	fmt.Println("   apiClient.StockBoard() - 股票板块相关API")
	fmt.Println("   apiClient.StockMarket() - 股票市场数据API")
	fmt.Println("   apiClient.StockBasic() - 股票基础信息API")

	fmt.Println("\n✅ 更符合面向对象设计:")
	fmt.Println("   apiClient 作为统一入口")
	fmt.Println("   按业务领域分组相关方法")

	fmt.Println("\n✅ 更好的 IDE 提示:")
	fmt.Println("   输入 apiClient. 后自动提示所有可用的 API 分类")
	fmt.Println("   选择分类后再提示该分类下的所有方法")

	fmt.Println("\n======================================")
	fmt.Println("实现说明")
	fmt.Println("======================================")

	fmt.Println("\n链式调用实现非常简单:")
	fmt.Println("1. 创建一个 APIClient 包装器嵌入 sdk.Client")
	fmt.Println("2. 为每个 API 分类创建一个 API 结构体")
	fmt.Println("3. 添加对应的方法，内部调用原始 API 函数")

	fmt.Println("\n核心代码:")
	fmt.Println(`
	// 1. 创建包装器
	type APIClient struct {
	    *sdk.Client
	}

	// 2. 为分类创建 API 结构体
	type StockBoardAPI struct {
	    client *sdk.Client
	}

	// 3. 添加分类访问方法
	func (c *APIClient) StockBoard() *StockBoardAPI {
	    return &StockBoardAPI{client: c.Client}
	}

	// 4. 包装具体 API 方法
	func (s *StockBoardAPI) TopList(ctx, req) {
	    return stockboard.TopList(ctx, s.client, req)
	}

	// 5. 使用
	apiClient := NewAPIClient(client)
	apiClient.StockBoard().TopList(ctx, req)
	`)

	fmt.Println("\n======================================")
	fmt.Println("扩展更多 API")
	fmt.Println("======================================")

	fmt.Println("\n你可以用同样的方式为其他 API 添加链式调用:")
	fmt.Println(`
	// 指数 API
	type IndexAPI struct { client *sdk.Client }
	func (c *APIClient) Index() *IndexAPI {
	    return &IndexAPI{client: c.Client}
	}
	func (i *IndexAPI) IndexBasic(ctx, req) {
	    return index.IndexBasic(ctx, i.client, req)
	}

	// 基金 API
	type FundAPI struct { client *sdk.Client }
	func (c *APIClient) Fund() *FundAPI {
	    return &FundAPI{client: c.Client}
	}

	// 期货 API
	type FuturesAPI struct { client *sdk.Client }
	func (c *APIClient) Futures() *FuturesAPI {
	    return &FuturesAPI{client: c.Client}
	}
	`)

	fmt.Println("\n======================================")
	fmt.Println("两种调用方式对比")
	fmt.Println("======================================")

	fmt.Println("\n📝 直接调用方式:")
	fmt.Println("   优点: 简单直接，不需要额外代码")
	fmt.Println("   缺点: 需要导入多个包")
	fmt.Println("   适用: 小型项目，调用 API 少")
	fmt.Println("\n   示例:")
	fmt.Println("   stockboard.TopList(ctx, client, req)")
	fmt.Println("   stockmarket.Daily(ctx, client, req)")

	fmt.Println("\n🔗 链式调用方式:")
	fmt.Println("   优点: 代码组织更好，IDE 提示友好")
	fmt.Println("   缺点: 需要额外的包装代码")
	fmt.Println("   适用: 大型项目，调用多种 API")
	fmt.Println("\n   示例:")
	fmt.Println("   apiClient.StockBoard().TopList(ctx, req)")
	fmt.Println("   apiClient.StockMarket().Daily(ctx, req)")

	fmt.Println("\n💡 建议:")
	fmt.Println("   - 可以在项目中创建自己的 api.go 文件")
	fmt.Println("   - 只为需要使用的 API 添加包装器")
	fmt.Println("   - 复用本示例的代码模板")

	fmt.Println("\n注意：当前 API spec 文件的 response_fields 为空，")
	fmt.Println("需要补充 Tushare API 的字段定义以生成完整的数据结构。")
	fmt.Println("请参考：https://tushare.pro/document/2")
}
