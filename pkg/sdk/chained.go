package sdk

import (
	"context"
)

// ==================== 链式调用客户端 ====================
// 以下结构体提供链式调用的 API 访问方式

// StockBoard 股票板块相关 API
type StockBoard struct {
	client *Client
}

// StockMarket 股票市场数据相关 API
type StockMarket struct {
	client *Client
}

// StockBasic 股票基础信息相关 API
type StockBasic struct {
	client *Client
}

// StockFinancial 股票财务数据相关 API
type StockFinancial struct {
	client *Client
}

// Index 指数相关 API
type Index struct {
	client *Client
}

// Futures 期货相关 API
type Futures struct {
	client *Client
}

// Fund 基金相关 API
type Fund struct {
	client *Client
}

// HKStock 港股相关 API
type HKStock struct {
	client *Client
}

// Bond 债券相关 API
type Bond struct {
	client *Client
}

// ETF ETF相关 API
type ETF struct {
	client *Client
}

// Forex 外汇相关 API
type Forex struct {
	client *Client
}

// Options 期权相关 API
type Options struct {
	client *Client
}

// Spot 现货相关 API
type Spot struct {
	client *Client
}

// USStock 美股相关 API
type USStock struct {
	client *Client
}

// Wealth 财富管理相关 API
type Wealth struct {
	client *Client
}

// Industry 行业经济相关 API
type Industry struct {
	client *Client
}

// LLMCorpus 大模型语料相关 API
type LLMCorpus struct {
	client *Client
}

// Macro 宏观经济相关 API
type Macro struct {
	client *Client
}

// ==================== Client 方法 ====================

// StockBoard 返回股票板块 API 客户端
func (c *Client) StockBoard() *StockBoard {
	return &StockBoard{client: c}
}

// StockMarket 返回股票市场数据 API 客户端
func (c *Client) StockMarket() *StockMarket {
	return &StockMarket{client: c}
}

// StockBasic 返回股票基础信息 API 客户端
func (c *Client) StockBasic() *StockBasic {
	return &StockBasic{client: c}
}

// StockFinancial 返回股票财务数据 API 客户端
func (c *Client) StockFinancial() *StockFinancial {
	return &StockFinancial{client: c}
}

// Index 返回指数 API 客户端
func (c *Client) Index() *Index {
	return &Index{client: c}
}

// Futures 返回期货 API 客户端
func (c *Client) Futures() *Futures {
	return &Futures{client: c}
}

// Fund 返回基金 API 客户端
func (c *Client) Fund() *Fund {
	return &Fund{client: c}
}

// HKStock 返回港股 API 客户端
func (c *Client) HKStock() *HKStock {
	return &HKStock{client: c}
}

// Bond 返回债券 API 客户端
func (c *Client) Bond() *Bond {
	return &Bond{client: c}
}

// ETF 返回ETF API 客户端
func (c *Client) ETF() *ETF {
	return &ETF{client: c}
}

// Forex 返回外汇 API 客户端
func (c *Client) Forex() *Forex {
	return &Forex{client: c}
}

// Options 返回期权 API 客户端
func (c *Client) Options() *Options {
	return &Options{client: c}
}

// Spot 返回现货 API 客户端
func (c *Client) Spot() *Spot {
	return &Spot{client: c}
}

// USStock 返回美股 API 客户端
func (c *Client) USStock() *USStock {
	return &USStock{client: c}
}

// Wealth 返回财富管理 API 客户端
func (c *Client) Wealth() *Wealth {
	return &Wealth{client: c}
}

// Industry 返回行业经济 API 客户端
func (c *Client) Industry() *Industry {
	return &Industry{client: c}
}

// LLMCorpus 返回大模型语料 API 客户端
func (c *Client) LLMCorpus() *LLMCorpus {
	return &LLMCorpus{client: c}
}

// Macro 返回宏观经济 API 客户端
func (c *Client) Macro() *Macro {
	return &Macro{client: c}
}

// ==================== 通用 API 调用方法 ====================
// 这些方法使用反射或通用接口来调用具体的 API 函数

// CallAPI 通用 API 调用方法，可以调用任何 API
// 使用示例: client.StockBoard().CallAPI(ctx, "top_list", params, fields, result)
func (s *StockBoard) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return s.client.CallAPI(ctx, apiName, params, fields, result)
}

func (s *StockMarket) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return s.client.CallAPI(ctx, apiName, params, fields, result)
}

func (s *StockBasic) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return s.client.CallAPI(ctx, apiName, params, fields, result)
}

func (s *StockFinancial) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return s.client.CallAPI(ctx, apiName, params, fields, result)
}

func (i *Index) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return i.client.CallAPI(ctx, apiName, params, fields, result)
}

func (f *Futures) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return f.client.CallAPI(ctx, apiName, params, fields, result)
}

func (f *Fund) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return f.client.CallAPI(ctx, apiName, params, fields, result)
}

func (h *HKStock) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	return h.client.CallAPI(ctx, apiName, params, fields, result)
}
