package tushare

import "fmt"

// BalanceSheet 获取上市公司资产负债表
func (api *TuShare) BalanceSheet(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "balancesheet",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
