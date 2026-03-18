package tushare

import "fmt"

// Top10Holders 获取上市公司前十大股东数据，包括持有数量和比例等信息
func (api *TuShare) Top10Holders(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top10_holders",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// Top10FloatHolders 获取上市公司前十大流通股东数据
func (api *TuShare) Top10FloatHolders(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "top10_floatholders",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// StkHolderNumber 获取上市公司股东户数数据，数据不定期公布
func (api *TuShare) StkHolderNumber(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stk_holdernumber",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
