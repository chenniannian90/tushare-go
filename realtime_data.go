package tushare

import "fmt"

// RTK 获取实时日k线行情，支持按股票代码及股票代码通配符一次性提取全部股票实时日k线行情
func (api *TuShare) RTK(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		params["ts_code"] = "3*.SZ,6*.SH,0*.SZ,9*.BJ"
	}
	body := map[string]interface{}{
		"api_name": "rt_k",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}
	return api.postData(body)
}

func (api *TuShare) RealTimeQuote(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]

	// ts_code & trade_date required
	if !hasTsCode {
		return nil, fmt.Errorf("need one argument ts_code")
	}

	body := map[string]interface{}{
		"api_name": "realtime_quote",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}
	return api.postData(body)
}

func (api *TuShare) RealTimeTick(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]

	// ts_code & trade_date required
	if !hasTsCode {
		return nil, fmt.Errorf("need one argument ts_code")
	}

	body := map[string]interface{}{
		"api_name": "realtime_tick",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}
	return api.postData(body)
}

func (api *TuShare) RealTimeList(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "realtime_list",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}
	return api.postData(body)
}
