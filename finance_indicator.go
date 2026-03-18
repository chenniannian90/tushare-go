package tushare

import "fmt"

// FinaIndicator 获取上市公司财务指标数据
func (api *TuShare) FinaIndicator(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_indicator",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// FinaAudit 获取上市公司定期财务审计意见数据
func (api *TuShare) FinaAudit(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_audit",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// FinaMainbz 获得上市公司主营业务构成，分地区和产品两种方式
func (api *TuShare) FinaMainbz(params map[string]string, fields []string) (*APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fina_mainbz",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
