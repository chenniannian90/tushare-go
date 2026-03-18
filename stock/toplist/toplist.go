package toplist

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

// TopList 获取龙虎榜每日统计数据

// 参数说明:
//   - trade_date: 交易日期（YYYYMMDD格式）
//   - exchange: 交易所代码（SSE上交所 SZSE深交所）
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, exalter, buy_amount, sell_amount, abnormal_reason

// 示例:
//
// params := map[string]string{
//     trade_date: "示例值",
//     exchange: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "exalter", "buy_amount", ...}
// resp, err := client.TopList(params, fields)
func (c *Client) TopList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "top_list", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// TopInst 获取龙虎榜机构交易明细

// 参数说明:
//   - trade_date: 交易日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, exalter, buy_amount, sell_amount, abnormal_reason

// 示例:
//
// params := map[string]string{
//     trade_date: "示例值",
//     limit: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "exalter", "buy_amount", ...}
// resp, err := client.TopInst(params, fields)
func (c *Client) TopInst(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "top_inst", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Concept(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "concept", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ConceptDetail(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "concept_detail", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) LimitList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "limit_list_d", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) STKLimit(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_limit", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) CiDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ci_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) SwDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "sw_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ThsIndex(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_index", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ThsDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ThsMember(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_member", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MoneyflowThs(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "moneyflow_ths", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MoneyflowIndThs(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "moneyflow_ind_ths", "token": c.getToken(), "params": params, "fields": fields,
	})
}
