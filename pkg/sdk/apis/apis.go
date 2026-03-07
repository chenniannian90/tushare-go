// Package apis 提供类型化的链式调用方法
// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法
package apis

import "tushare-go/pkg/sdk"

// TushareClient 提供链式调用客户端，集成所有 API 接口
type TushareClient struct {
	client *sdk.Client
	StockMarket StockMarket
	StockReference StockReference
	UsStock UsStock
	Fund Fund
	Index Index
	MacroUsRate MacroUsRate
	StockBasic StockBasic
	StockBoard StockBoard
	WealthFundSales WealthFundSales
	Forex Forex
	MacroSocialFinancing MacroSocialFinancing
	MacroInterestRate MacroInterestRate
	MacroPrice MacroPrice
	Options Options
	Spot Spot
	StockFeature StockFeature
	StockFundFlow StockFundFlow
	MacroMoneySupply MacroMoneySupply
	Bond Bond
	Etf Etf
	Futures Futures
	IndustryTmt IndustryTmt
	MacroBusiness MacroBusiness
	MacroEconomy MacroEconomy
	StockMargin StockMargin
	HkStock HkStock
	LlmCorpus LlmCorpus
	StockFinancial StockFinancial
}

// NewTushareClient 创建一个新的链式调用客户端
func NewTushareClient(client *sdk.Client) *TushareClient {
	return &TushareClient{
		client: client,
		StockMarket: newstockMarketImpl(client),
		StockReference: newstockReferenceImpl(client),
		UsStock: newusStockImpl(client),
		Fund: newfundImpl(client),
		Index: newindexImpl(client),
		MacroUsRate: newmacroUsRateImpl(client),
		StockBasic: newstockBasicImpl(client),
		StockBoard: newstockBoardImpl(client),
		WealthFundSales: newwealthFundSalesImpl(client),
		Forex: newforexImpl(client),
		MacroSocialFinancing: newmacroSocialFinancingImpl(client),
		MacroInterestRate: newmacroInterestRateImpl(client),
		MacroPrice: newmacroPriceImpl(client),
		Options: newoptionsImpl(client),
		Spot: newspotImpl(client),
		StockFeature: newstockFeatureImpl(client),
		StockFundFlow: newstockFundFlowImpl(client),
		MacroMoneySupply: newmacroMoneySupplyImpl(client),
		Bond: newbondImpl(client),
		Etf: newetfImpl(client),
		Futures: newfuturesImpl(client),
		IndustryTmt: newindustryTmtImpl(client),
		MacroBusiness: newmacroBusinessImpl(client),
		MacroEconomy: newmacroEconomyImpl(client),
		StockMargin: newstockMarginImpl(client),
		HkStock: newhkStockImpl(client),
		LlmCorpus: newllmCorpusImpl(client),
		StockFinancial: newstockFinancialImpl(client),
	}
}
