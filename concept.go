package tushare

// Concept 获取概念股分类，目前只有ts一个来源，未来将逐步增加来源
func (api *TuShare) Concept(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "concept",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}

// ConceptDetail 获取概念股分类明细数据
func (api *TuShare) ConceptDetail(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "concept_detail",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}
	return api.postData(body)
}
