package tushare

// CiDaily 中信行业指数行情
func (api *TuShare) CiDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ci_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// SwDaily 万申���数日线行情
func (api *TuShare) SwDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "sw_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
