package tushare

import "fmt"

// Margin 获取融资融券每日交易汇总数据
func (api *TuShare) Margin(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "margin",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// MarginDetail 获取沪深两市每日融资融券明细
func (api *TuShare) MarginDetail(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "margin_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
