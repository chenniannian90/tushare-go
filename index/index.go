package index

import (
	"github.com/chenniannian90/tushare-go/types"
)

// Client represents the index API client
type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

// New creates a new index client
func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{
		postData: postData,
		getToken: getToken,
	}
}

// IndexDaily 获取指数日线行情

// 参数说明:
//   - ts_code: 指数代码
//   - trade_date: 交易日期
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, close, open, high, low, vol, amount

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     trade_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "close", "open", ...}
// resp, err := client.IndexDaily(params, fields)
func (c *Client) IndexDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_daily",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexDailyBasic 大盘指数每日指标
func (c *Client) IndexDailyBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_dailybasic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexClassify 申万行业分类
func (c *Client) IndexClassify(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_classify",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexGlobal 国际指数
func (c *Client) IndexGlobal(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_global",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexBasic 获取指数基础信息

// 参数说明:
//   - market: 市场代码
//   - publisher: 发布人
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, name, market, publisher, category, base_date, base_point, list_date

// 示例:
//
// params := map[string]string{
//     market: "示例值",
//     publisher: "示例值",
// }
// fields := []string{"ts_code", "name", "market", "publisher", ...}
// resp, err := client.IndexBasic(params, fields)
func (c *Client) IndexBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexWeight 获取指数成分和权重

// 参数说明:
//   - index_code: 指数代码
//   - trade_date: 交易日期
//   - limit: 单次返回数据长度

// 输出字段:
//   index_code, con_code, trade_date, weight, is_new

// 示例:
//
// params := map[string]string{
//     index_code: "示例值",
//     trade_date: "示例值",
// }
// fields := []string{"index_code", "con_code", "trade_date", "weight", ...}
// resp, err := client.IndexWeight(params, fields)
func (c *Client) IndexWeight(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_weight",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexMember 申万行业成分构成
func (c *Client) IndexMember(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_member",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}
