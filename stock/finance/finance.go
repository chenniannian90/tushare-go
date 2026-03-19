package finance

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

// Income 获取利润表数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, ann_date, f_ann_date, end_date, report_type, comp_type, basic_eps, diluted_eps

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "ann_date", "f_ann_date", "end_date", ...}
// resp, err := client.Income(params, fields)
func (c *Client) Income(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "income", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// BalanceSheet 获取资产负债表数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, ann_date, end_date, report_type, comp_type, total_assets, total_liab, equities

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "ann_date", "end_date", "report_type", ...}
// resp, err := client.BalanceSheet(params, fields)
func (c *Client) BalanceSheet(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "balancesheet", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// CashFlow 获取现金流量表数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, ann_date, end_date, comp_type, net_cash_flows, n_cash_flows_frm_oa

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "ann_date", "end_date", "comp_type", ...}
// resp, err := client.CashFlow(params, fields)
func (c *Client) CashFlow(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "cashflow", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaIndicator(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_indicator", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaAudit(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_audit", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaMainbz(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_mainbz", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Forecast(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasAnnDate := params["ann_date"]
	if (!hasTsCode && !hasAnnDate) || (hasTsCode && hasAnnDate) {
		return nil, fmt.Errorf("need one argument ts_code or ann_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "forecast", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Express(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "express", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Dividend(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "dividend", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) DisclosureDate(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "disclosure_date", "token": c.getToken(), "params": params, "fields": fields,
	})
}
