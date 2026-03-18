package tushare

// MoneyFlow 个股资金流向
func (api *TuShare) MoneyFlow(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "moneyflow",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
