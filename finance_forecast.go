package tushare

import "fmt"

// Forecast 获取业绩预告数据
func (api *TuShare) Forecast(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasAnnDate := params["ann_date"]

	if (!hasTsCode && !hasAnnDate) || (hasTsCode && hasAnnDate) {
		return nil, fmt.Errorf("need one argument ts_code or ann_date")
	}

	body := map[string]interface{}{
		"api_name": "forecast",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// Express 获取上市公司业绩快报
func (api *TuShare) Express(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "express",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// Dividend 分红送股数据
func (api *TuShare) Dividend(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "dividend",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// DisclosureDate 获取财报披露计划日期
func (api *TuShare) DisclosureDate(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "disclosure_date",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
