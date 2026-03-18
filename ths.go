package tushare

// ThsIndex 同花顺概念和行业指数
func (api *TuShare) ThsIndex(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_index",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// ThsDaily 同花顺板块指数行情
func (api *TuShare) ThsDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// ThsMember 同花顺概念板块成分
func (api *TuShare) ThsMember(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_member",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
