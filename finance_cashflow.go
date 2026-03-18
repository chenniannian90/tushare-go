package tushare

import "fmt"

// CashFlow 获取上市��司现金流量表
func (api *TuShare) CashFlow(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "cashflow",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
