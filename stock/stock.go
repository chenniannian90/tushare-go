package stock

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/types"
)

// Client represents the stock API client
type Client struct {
	postData  types.PostFunc
	getToken  types.TokenFunc
}

// New creates a new stock client
func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{
		postData: postData,
		getToken: getToken,
	}
}

// StockBasic 获取基础信息数据，包括股票代码、名称、上市日期、退市日期等
func (c *Client) StockBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// BakBasic 获取B股基础信息
func (c *Client) BakBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "bak_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// TradeCal 获取各大交易所��易日历数据,默认提取的是上交所
func (c *Client) TradeCal(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// HSConst 获取沪股通、深股通成分数据
func (c *Client) HSConst(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "hs_const",
		"token":    c.getToken(),
		"fields":   fields,
	}
	if _, ok := params["hs_type"]; !ok {
		return nil, fmt.Errorf("hs_type required")
	}
	body["params"] = params
	return c.postData(body)
}

// NameChange 历史名称变更记录
func (c *Client) NameChange(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "namechange",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// StockCompany 获取上市公司基础信息
func (c *Client) StockCompany(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_company",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// NewShare 获取新股上市列表数据
func (c *Client) NewShare(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "new_share",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}
