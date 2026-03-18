package tushare

// DailyInfo 市场交易统计
func (api *TuShare) DailyInfo(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "daily_info",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// SzDailyInfo 深圳市场每日交易概况
func (api *TuShare) SzDailyInfo(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "sz_daily_info",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
