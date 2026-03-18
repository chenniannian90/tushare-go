package tushare

import "fmt"

// MoneyflowHsgt 获取沪股通、深股通、港股通每日资金流向数据
func (api *TuShare) MoneyflowHsgt(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTradeDate := params["trade_date"]
	_, hasStartDate := params["start_date"]

	if (!hasTradeDate && !hasStartDate) || (hasTradeDate && hasStartDate) {
		return nil, fmt.Errorf("need one argument trade_date or start_date")
	}

	body := map[string]interface{}{
		"api_name": "moneyflow_hsgt",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// HsgtTop10 获取沪股通、深股通每日前十大成交详细数据
func (api *TuShare) HsgtTop10(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "hsgt_top10",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// GgtTop10 获取港股通每日成交数据，其中包括沪市、深市详细数据
func (api *TuShare) GgtTop10(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}

	body := map[string]interface{}{
		"api_name": "ggt_top10",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
