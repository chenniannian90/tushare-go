package tushare

import "fmt"

// PledgeStat 获取股权质押统计数据
func (api *TuShare) PledgeStat(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "pledge_stat",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// PledgeDetail 获取股权质押明细数据
func (api *TuShare) PledgeDetail(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "pledge_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
