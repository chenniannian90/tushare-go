// Package apis 提供类型化的链式调用方法
// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法
package apis

import (
	"context"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
	stockmarket "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_market"
	stockbasic "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_basic"
	stockfinancial "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_financial"
)

// ==================== StockBoard 方法 ====================

// TopList 获取龙虎榜每日统���
func TopList(ctx context.Context, client *sdk.Client, req *stockboard.TopListRequest) ([]stockboard.TopListItem, error) {
	return stockboard.TopList(ctx, client, req)
}

// LimitList 获取涨跌停和炸板数据
func LimitList(ctx context.Context, client *sdk.Client, req *stockboard.LimitListRequest) ([]stockboard.LimitListItem, error) {
	return stockboard.LimitList(ctx, client, req)
}

// DragonList 获取游资交易每日明细
func DragonList(ctx context.Context, client *sdk.Client, req *stockboard.DragonListRequest) ([]stockboard.DragonListItem, error) {
	return stockboard.DragonList(ctx, client, req)
}

// TopInst 获取营业部席位买入排名
func TopInst(ctx context.Context, client *sdk.Client, req *stockboard.TopInstRequest) ([]stockboard.TopInstItem, error) {
	return stockboard.TopInst(ctx, client, req)
}

// ThsConcept 获取同花顺概念板块
func ThsConcept(ctx context.Context, client *sdk.Client, req *stockboard.ThsConceptRequest) ([]stockboard.ThsConceptItem, error) {
	return stockboard.ThsConcept(ctx, client, req)
}

// EmHot 获取东方财富App热榜
func EmHot(ctx context.Context, client *sdk.Client, req *stockboard.EmHotRequest) ([]stockboard.EmHotItem, error) {
	return stockboard.EmHot(ctx, client, req)
}

// ==================== StockMarket 方法 ====================

// Daily 获取日线行情
func Daily(ctx context.Context, client *sdk.Client, req *stockmarket.DailyRequest) ([]stockmarket.DailyItem, error) {
	return stockmarket.Daily(ctx, client, req)
}

// DailyBasic 获取每日基本面指标
func DailyBasic(ctx context.Context, client *sdk.Client, req *stockmarket.DailyBasicRequest) ([]stockmarket.DailyBasicItem, error) {
	return stockmarket.DailyBasic(ctx, client, req)
}

// Weekly 获取周线行情
func Weekly(ctx context.Context, client *sdk.Client, req *stockmarket.WeeklyRequest) ([]stockmarket.WeeklyItem, error) {
	return stockmarket.Weekly(ctx, client, req)
}

// Monthly 获取月线行情
func Monthly(ctx context.Context, client *sdk.Client, req *stockmarket.MonthlyRequest) ([]stockmarket.MonthlyItem, error) {
	return stockmarket.Monthly(ctx, client, req)
}

// ==================== StockBasic 方法 ====================

// TradeCal 获取交易日历
func TradeCal(ctx context.Context, client *sdk.Client, req *stockbasic.TradeCalRequest) ([]stockbasic.TradeCalItem, error) {
	return stockbasic.TradeCal(ctx, client, req)
}

// StockBasicInfo 获取股票列表（避免与 StockBasic 结构体冲突）
func StockBasicInfo(ctx context.Context, client *sdk.Client, req *stockbasic.StockBasicRequest) ([]stockbasic.StockBasicItem, error) {
	return stockbasic.StockBasic(ctx, client, req)
}

// ==================== StockFinancial 方法 ====================

// Income 获取利润表数据
func Income(ctx context.Context, client *sdk.Client, req *stockfinancial.IncomeRequest) ([]stockfinancial.IncomeItem, error) {
	return stockfinancial.Income(ctx, client, req)
}

// Balancesheet 获取资产负债表
func Balancesheet(ctx context.Context, client *sdk.Client, req *stockfinancial.BalancesheetRequest) ([]stockfinancial.BalancesheetItem, error) {
	return stockfinancial.Balancesheet(ctx, client, req)
}

// FinaIndicator 获取财务指标
func FinaIndicator(ctx context.Context, client *sdk.Client, req *stockfinancial.FinaIndicatorRequest) ([]stockfinancial.FinaIndicatorItem, error) {
	return stockfinancial.FinaIndicator(ctx, client, req)
}

// Dividend 获取分红数据
func Dividend(ctx context.Context, client *sdk.Client, req *stockfinancial.DividendRequest) ([]stockfinancial.DividendItem, error) {
	return stockfinancial.Dividend(ctx, client, req)
}
