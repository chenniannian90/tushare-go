// Package apis 提供类型化的链式调用方法
// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法
package apis

import "github.com/chenniannian90/tushare-go/pkg/sdk"

// TushareClient 提供链式调用客户端，集成所有 API 接口
type TushareClient struct {
	client *sdk.Client
	Fund Fund
	Futures Futures
	MacroBusiness MacroBusiness
	Options Options
	StockBasic StockBasic
	StockBoard StockBoard
	StockFinancial StockFinancial
	StockFundFlow StockFundFlow
	MacroInterestRate MacroInterestRate
	Index Index
	IndustryTmt IndustryTmt
	MacroPrice MacroPrice
	StockFeature StockFeature
	StockMargin StockMargin
	StockReference StockReference
	WealthFundSales WealthFundSales
	HkStock HkStock
	LlmCorpus LlmCorpus
	StockMarket StockMarket
	UsStock UsStock
	Bond Bond
	Etf Etf
	Forex Forex
	MacroEconomy MacroEconomy
	MacroMoneySupply MacroMoneySupply
	MacroSocialFinancing MacroSocialFinancing
	MacroUsRate MacroUsRate
	Spot Spot
}

// NewTushareClient 创建一个新的链式调用客户端
func NewTushareClient(client *sdk.Client) *TushareClient {
	return &TushareClient{
		client: client,
		Fund: newfundImpl(client),
		Futures: newfuturesImpl(client),
		MacroBusiness: newmacroBusinessImpl(client),
		Options: newoptionsImpl(client),
		StockBasic: newstockBasicImpl(client),
		StockBoard: newstockBoardImpl(client),
		StockFinancial: newstockFinancialImpl(client),
		StockFundFlow: newstockFundFlowImpl(client),
		MacroInterestRate: newmacroInterestRateImpl(client),
		Index: newindexImpl(client),
		IndustryTmt: newindustryTmtImpl(client),
		MacroPrice: newmacroPriceImpl(client),
		StockFeature: newstockFeatureImpl(client),
		StockMargin: newstockMarginImpl(client),
		StockReference: newstockReferenceImpl(client),
		WealthFundSales: newwealthFundSalesImpl(client),
		HkStock: newhkStockImpl(client),
		LlmCorpus: newllmCorpusImpl(client),
		StockMarket: newstockMarketImpl(client),
		UsStock: newusStockImpl(client),
		Bond: newbondImpl(client),
		Etf: newetfImpl(client),
		Forex: newforexImpl(client),
		MacroEconomy: newmacroEconomyImpl(client),
		MacroMoneySupply: newmacroMoneySupplyImpl(client),
		MacroSocialFinancing: newmacroSocialFinancingImpl(client),
		MacroUsRate: newmacroUsRateImpl(client),
		Spot: newspotImpl(client),
	}
}
