// Package apis 提供类型化的链式调用方法
// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法
package apis

import "github.com/chenniannian90/tushare-go/pkg/sdk"

// TushareClient 提供链式调用客户端，集成所有 API 接口
type TushareClient struct {
	client *sdk.Client
	Forex Forex
	Fund Fund
	Futures Futures
	HkStock HkStock
	Index Index
	IndustryTmt IndustryTmt
	LlmCorpus LlmCorpus
	MacroEconomy MacroEconomy
	Bond Bond
	MacroSocialFinancing MacroSocialFinancing
	MacroPrice MacroPrice
	MacroUsRate MacroUsRate
	Options Options
	Spot Spot
	StockFundFlow StockFundFlow
	StockReference StockReference
	MacroBusiness MacroBusiness
	MacroMoneySupply MacroMoneySupply
	MacroInterestRate MacroInterestRate
	StockBasic StockBasic
	StockBoard StockBoard
	StockFinancial StockFinancial
	WealthFundSales WealthFundSales
	Etf Etf
	StockFeature StockFeature
	StockMargin StockMargin
	StockMarket StockMarket
	UsStock UsStock
}

// NewTushareClient 创建一个新的链式调用客户端
func NewTushareClient(client *sdk.Client) *TushareClient {
	return &TushareClient{
		client: client,
		Forex: newforexImpl(client),
		Fund: newfundImpl(client),
		Futures: newfuturesImpl(client),
		HkStock: newhkStockImpl(client),
		Index: newindexImpl(client),
		IndustryTmt: newindustryTmtImpl(client),
		LlmCorpus: newllmCorpusImpl(client),
		MacroEconomy: newmacroEconomyImpl(client),
		Bond: newbondImpl(client),
		MacroSocialFinancing: newmacroSocialFinancingImpl(client),
		MacroPrice: newmacroPriceImpl(client),
		MacroUsRate: newmacroUsRateImpl(client),
		Options: newoptionsImpl(client),
		Spot: newspotImpl(client),
		StockFundFlow: newstockFundFlowImpl(client),
		StockReference: newstockReferenceImpl(client),
		MacroBusiness: newmacroBusinessImpl(client),
		MacroMoneySupply: newmacroMoneySupplyImpl(client),
		MacroInterestRate: newmacroInterestRateImpl(client),
		StockBasic: newstockBasicImpl(client),
		StockBoard: newstockBoardImpl(client),
		StockFinancial: newstockFinancialImpl(client),
		WealthFundSales: newwealthFundSalesImpl(client),
		Etf: newetfImpl(client),
		StockFeature: newstockFeatureImpl(client),
		StockMargin: newstockMarginImpl(client),
		StockMarket: newstockMarketImpl(client),
		UsStock: newusStockImpl(client),
	}
}
