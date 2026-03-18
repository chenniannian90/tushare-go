package market

import (
	"fmt"
	"regexp"
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func IsDateFormat(dates ...string) bool {
	pattern := `^\d{8}$`
	for _, date := range dates {
		if date == "" {
			continue // Skip empty strings
		}
		match, _ := regexp.MatchString(pattern, date)
		if !match {
			return false
		}
	}
	return true
}

func (c *Client) MoneyFlow(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "moneyflow", "token": c.getToken(), "params": params, "fields": fields})
}

func (c *Client) DailyInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "daily_info", "token": c.getToken(), "params": params, "fields": fields})
}

func (c *Client) SzDailyInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "sz_daily_info", "token": c.getToken(), "params": params, "fields": fields})
}

// Daily 获取日线行情数据

// 参数说明:
//   - ts_code: 股票代码（支持多选）
//   - trade_date: 交易日期（YYYYMMDD格式）
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, open, high, low, close, pre_close, change...

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     trade_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "open", "high", ...}
// resp, err := client.Daily(params, fields)
func (c *Client) Daily(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "daily", "token": c.getToken(), "fields": fields, "params": params})
}

// Weekly 获取周线行情数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, open, high, low, close, vol, amount

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "open", "high", ...}
// resp, err := client.Weekly(params, fields)
func (c *Client) Weekly(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "weekly", "token": c.getToken(), "fields": fields, "params": params})
}

// Monthly 获取月线行情数据

// 参数说明:
//   - ts_code: 股票代码
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - limit: 单次返回数据长度

// 输出字段:
//   ts_code, trade_date, open, high, low, close, vol, amount

// 示例:
//
// params := map[string]string{
//     ts_code: "示例值",
//     start_date: "示例值",
// }
// fields := []string{"ts_code", "trade_date", "open", "high", ...}
// resp, err := client.Monthly(params, fields)
func (c *Client) Monthly(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "monthly", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) DailyBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "daily_basic", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) AdjFactor(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "adj_factor", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) Suspend(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["suspend_date"]
	_, hasResumeDate := params["resume_date"]
	argsCount := 0
	if hasTsCode { argsCount++ }
	if hasTradeDate { argsCount++ }
	if hasResumeDate { argsCount++ }
	if argsCount != 1 {
		return nil, fmt.Errorf("need one argument among ts_code, suspend_date, resume_date")
	}
	if !IsDateFormat(params["suspend_date"], params["resume_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "suspend", "token": c.getToken(), "params": params, "fields": fields})
}

// RTK 获取实时K线数据

// 参数说明:
//   - ts_code: 股票代码（可选，默认为所有股票）

// 输出字段:
//   ts_code, trade_date, close, open, high, low, vol, amount

// 示例:
//
// params := map[string]string{}
// fields := []string{"ts_code", "trade_date", "close", "open", ...}
// resp, err := client.RTK(params, fields)
func (c *Client) RTK(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		params["ts_code"] = "3*.SZ,6*.SH,0*.SZ,9*.BJ"
	}
	return c.postData(map[string]interface{}{
		"api_name": "rt_k", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// RealTimeQuote 获取实时行情数据

// 参数说明:
//   - ts_code: 股票代码（必填）

// 输出字段:
//   ts_code, trade_date, close, open, high, low, vol, amount

// 示例:
//
// params := map[string]string{"ts_code": "000001.SZ"}
// fields := []string{"ts_code", "trade_date", "close", "open", ...}
// resp, err := client.RealTimeQuote(params, fields)
func (c *Client) RealTimeQuote(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}
	return c.postData(map[string]interface{}{
		"api_name": "realtime_quote", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// RealTimeTick 获取实时分笔成交数据

// 参数说明:
//   - ts_code: 股票代码（必填）

// 输出字段:
//   ts_code, trade_date, close, open, high, low, vol, amount

// 示例:
//
// params := map[string]string{"ts_code": "000001.SZ"}
// fields := []string{"ts_code", "trade_date", "close", "open", ...}
// resp, err := client.RealTimeTick(params, fields)
func (c *Client) RealTimeTick(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}
	return c.postData(map[string]interface{}{
		"api_name": "realtime_tick", "token": c.getToken(), "params": params, "fields": fields,
	})
}

// RealTimeList 获取实时行情列表数据

// 参数说明:
//   - 无需参数

// 输出字段:
//   ts_code, trade_date, close, open, high, low, vol, amount

// 示例:
//
// params := map[string]string{}
// fields := []string{"ts_code", "trade_date", "close", "open", ...}
// resp, err := client.RealTimeList(params, fields)
func (c *Client) RealTimeList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "realtime_list", "token": c.getToken(), "params": params, "fields": fields,
	})
}
