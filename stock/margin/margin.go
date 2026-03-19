package margin

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

// Margin 获取融资融券交易汇总数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, exchange_id, rzye, rzmre, rzche, rqyl, rqylchl

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "exchange_id", "rzye", ...}
// resp, err := client.Margin(params, fields)
func (c *Client) Margin(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "margin", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MarginDetail(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "margin_detail", "token": c.getToken(), "params": params, "fields": fields,
	})
}
