package tushare

// IndexDaily 指数日线行情
func (api *TuShare) IndexDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// IndexDailyBasic 大盘指数每日指标
func (api *TuShare) IndexDailyBasic(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_dailybasic",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// IndexClassify 申万行业分类
func (api *TuShare) IndexClassify(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_classify",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// IndexGlobal 国际指数
func (api *TuShare) IndexGlobal(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_global",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
