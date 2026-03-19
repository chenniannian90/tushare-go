package etf

import (
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{
		postData: postData,
		getToken: getToken,
	}
}

// ETFBasic 获取ETF基础信息
// 参数:
//   - ts_code: TS代码
//   - index_code: 指数代码
//   - list_date: 上市日期
//   - list_status: 上市状态 (L上市D退市P暂停)
//   - exchange: 交易所 (SSE上交所SZSE深交所)
//   - mgr: 管理人
//   - limit: 单次返回数据长度
// ETFBasic 获取ETF基础信息

// 参数说明:
//   - ts_code: ETF代码
//   - market: 市场类型
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, name, list_date, list_status, fund_type, manage_type, underlying_index

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     market: "示例值",
// }
// fields := []string{"ts_code", "name", "list_date", "list_status", ...}
// resp, err := client.ETFBasic(params, fields)
func (c *Client) ETFBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "etf_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// FundDaily 获取基金日线行情
// 参数:
//   - trade_date: 交易日期
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - ts_code: TS代码
//   - limit: 单次返回数据长度
// FundDaily 获取基金日线行情

// 参数说明:
//   - ts_code: 基金代码
//   - trade_date: 交易日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, open, high, low, close, vol, amount

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     trade_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "open", "high", ...}
// resp, err := client.FundDaily(params, fields)
func (c *Client) FundDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "fund_daily",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// FundAdj 获取基金复权因子
// 参数:
//   - ts_code: TS代码
//   - trade_date: 交易日期
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - offset: 偏移量
// FundAdj 获取基金复权因子

// 参数说明:
//   - ts_code: 基金代码
//   - trade_date: 交易日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, adj_factor

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     trade_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "adj_factor"}
// resp, err := client.FundAdj(params, fields)
func (c *Client) FundAdj(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "fund_adj",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}
