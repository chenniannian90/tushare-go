// Package apis 提供类型化的链式调用方法
// 这个包导入了 SDK 和具体的 API 包，为 SDK 的链式调用客户端提供类型化方法
package apis

import "github.com/chenniannian90/tushare-go/pkg/sdk"

// TushareClient 提供链式调用客户端，集成所有 API 接口
type TushareClient struct {
	client *sdk.Client
	Stock
	Bond
	ETF
	Fund
	Futures
	Forex
	HKStock
	Index
	Industry
	LLMCorpus
	Options
	Spot
	USStock
	Wealth
}

// NewTushareClient 创建一个新的链式调用客户端
func NewTushareClient(client *sdk.Client) *TushareClient {
	return &TushareClient{
		client:     client,
		Stock:      newStockImpl(client),
		Bond:       newBondImpl(client),
		ETF:        newEtfImpl(client),
		Fund:       newFundImpl(client),
		Futures:    newFuturesImpl(client),
		Forex:      newForexImpl(client),
		HKStock:    newHkStockImpl(client),
		Index:      newIndexImpl(client),
		Industry:   newIndustryImpl(client),
		LLMCorpus:  newLLMCorpusImpl(client),
		Options:    newOptionsImpl(client),
		Spot:       newSpotImpl(client),
		USStock:    newUsStockImpl(client),
		Wealth:     newWealthImpl(client),
	}
}
