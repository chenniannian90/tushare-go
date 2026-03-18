package reference

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

// Top10Holders 获取前十大股东数据

// 参数说明:
//   - ts_code: 股票代码
//   - period: 报告期（如20231231）
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, ann_date, end_date, holder_name, hold_amount, hold_ratio

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     period: "示例值",
// }
// fields := []string{"ts_code", "ann_date", "end_date", "holder_name", ...}
// resp, err := client.Top10Holders(params, fields)
func (c *Client) Top10Holders(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "top10_holders", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// Top10FloatHolders 获取前十大流通股东数据

// 参数说明:
//   - ts_code: 股票代码
//   - period: 报告期（如20231231）
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, ann_date, end_date, holder_name, hold_amount, hold_ratio

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     period: "示例值",
// }
// fields := []string{"ts_code", "ann_date", "end_date", "holder_name", ...}
// resp, err := client.Top10FloatHolders(params, fields)
func (c *Client) Top10FloatHolders(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "top10_floatholders", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// StkHolderNumber 获取股东人数数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, end_date, holder_num, holder_num_chg

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "end_date", "holder_num", "holder_num_chg"}
// resp, err := client.StkHolderNumber(params, fields)
func (c *Client) StkHolderNumber(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_holdernumber", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) PledgeStat(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "pledge_stat", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) PledgeDetail(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "pledge_detail", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Repurchase(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "repurchase", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ShareFloat(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "share_float", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) BlockTrade(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "block_trade", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) StkAccount(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_account", "token": c.getToken(), "params": params, "fields": fields,
	})
}
