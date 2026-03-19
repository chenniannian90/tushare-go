package basic

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/types"
)

// Client represents the stock basic API client
type Client struct {
	postData  types.PostFunc
	getToken  types.TokenFunc
}

// New creates a new stock basic client
func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{
		postData: postData,
		getToken: getToken,
	}
}

// StockBasic 获取股票基础信息数据
//
// 获取股票基础信息数据，包括股票代码、名称、上市日期、退市日期等基础信息。
// 输入参数支持股票代码、上市状态、交易所等多个维度的过滤查询。
//
// 参数说明:
//   - ts_code: 股票代码（支持多选或csv格式，如 "600000.SH,000001.SZ"）
//   - list_status: 上市状态（L上市 D退市 P暂停上市，默认L）
//   - exchange: 交易所代码（SSE上交所 SZSE深交所 CFXCC大商所 SHFE上期所 CFFEX郑商所 DCE郑商所 CZCE大商��� INE上能源）
//   - market: 市场类型（主板股票 创业板 科创板 北交所（stock_basic；默认主板股票）
//   - name: 股票名称（支持模糊搜索）
//   - fields: 可选字段（参见官网文档）
//   - limit: 单次返回数据长度（默认3000，最大6000）
//   - offset: 偏移量（从第几条开始返回）
//
// 输出字段:
//   - ts_code: TS代码
//   - symbol: 股票简称
//   - name: 股票名称
//   - area: 所属地域
//   - industry: 所属行业
//   - market: 市场类型（主板/创业板/科创板/北交所）
//   - list_date: 上市日期
//   - delist_date: 退市日期
//   - is_hs: 是否沪深港通标的
//
// 示例:
//
//	params := map[string]string{
//	    "ts_code": "600000.SH",
//	    "list_status": "L",
//	    "limit": "5",
//	}
//	fields := []string{"ts_code", "name", "industry", "list_date"}
//	resp, err := client.StockBasic(params, fields)
func (c *Client) StockBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// BakBasic 获取B股基础信息
//
// 获取B股基础信息数据，包括B股的代码、名称、上市日期等信息。
//
// 参数说明:
//   - ts_code: B股代码（支持多选）
//   - fields: 可选字段
//   - limit: 单次返回数据长度
//
// 输出字段:
//   - ts_code: TS代码
//   - symbol: 股票简称
//   - name: 股票名称
//   - area: 所属地域
//   - industry: 所属行业
//   - list_date: 上市日期
//   - is_hs: 是否沪深港通标的
//
// 示例:
//
//	params := map[string]string{
//	    "limit": "10",
//	}
//	fields := []string{"ts_code", "name", "list_date"}
//	resp, err := client.BakBasic(params, fields)
func (c *Client) BakBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "bak_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// TradeCal 获取交易日历数据
//
// 获取各大交易所交易日历数据，默认提取的是上交所。
// 输出包括交易日历、是否交易、是否开市等信息。
//
// 参数说明:
//   - exchange: 交易所代码（SSE上交所 SZSE深交所 CFFEX郑商所 SHFE上期所 DCE郑商所 CZCE大商所 INE上能源 CFXCC大商所）
//   - start_date: 开始日期 (YYYYMMDD格式)
//   - end_date: 结束日期 (YYYYMMDD格式)
//   - is_open: 是否开市（1开市 0休市）
//   - fields: 可选字段
//   - limit: 单次返回数据长度
//
// 输出字段:
//   - cal_date: 日历日期
//   - exchange: 交易所代码
//   - is_open: 是否开市（1开市 0休市）
//   - pretrade: 是否有盘前交易（1有 0无）
//
// 示例:
//
//	params := map[string]string{
//	    "exchange": "SSE",
//	    "start_date": "20240101",
//	    "end_date": "20240131",
//	}
//	fields := []string{"cal_date", "exchange", "is_open"}
//	resp, err := client.TradeCal(params, fields)
func (c *Client) TradeCal(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// HSConst 获取沪深港通成分数据
//
// 获取沪股通、深股通成分数据，包括港股通标的的基本信息和变动情况。
//
// 参数说明:
//   - hs_type: 沪港通类型（SH沪港通 SZ深港通，必填）
//   - fields: 可选字段
//   - limit: 单次返回数据长度
//
// 输出字段:
//   - ts_code: TS代码
//   - hs_type: 沪港通类型
//   - in_date: 纳入日期
//   | out_date: 剔出日期
//   - - in_out: 纳入/剔除标记（1纳入 2剔除）
//
// 示例:
//
//	params := map[string]string{
//	    "hs_type": "SH",
//	    "limit": "10",
//	}
//	fields := []string{"ts_code", "hs_type", "in_date", "out_date"}
//	resp, err := client.HSConst(params, fields)
func (c *Client) HSConst(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "hs_const",
		"token":    c.getToken(),
		"fields":   fields,
	}
	if _, ok := params["hs_type"]; !ok {
		return nil, fmt.Errorf("hs_type required")
	}
	body["params"] = params
	return c.postData(body)
}

// NameChange 获取股票历史名称变更记录
//
// 获取股票历史名称变更记录，包括股票简称的更名历史。
//
// 参数说明:
//   - ts_code: 股票代码
//   - fields: 可选字段
//   - start_date: 更名公告日期
//   - end_date: 更名公告日期
//   - limit: 单次返回数据长度
//
// 输出字段:
//   - ts_code: TS代码
//   - name: 股票简称
//   - start_date: 更名公告日期
//   - end_date: 更名公告日期
//   - change_reason: 变更原因
//
// 示例:
//
//	params := map[string]string{
//	    "ts_code": "600000.SH",
//	    "start_date": "20200101",
//	    "end_date": "20231231",
//	}
//	fields := []string{"ts_code", "name", "start_date", "end_date"}
//	resp, err := client.NameChange(params, fields)
func (c *Client) NameChange(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "namechange",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// StockCompany 获取上市公司基础信息
//
// 获取上市公司基础信息，包括公司注册地、办公地址、法人代表等信息。
//
// 参数说明:
//   - ts_code: 股票代码
//   - fields: 可选字段
//   - exchange: 交易所
//
// 输出字段:
//   - ts_code: TS代码
//   - name: 股票简称
//   - chairman: 法人代表
//   - manager: 总经理
//   - secretary: 董事会秘书
//   - setup_date: 成立日期
//   - province: 所在省份
//   - city: 所在城市
//   - introduction: 公司介绍
//   - website: 公司网站
//   - email: 电子邮箱
//   - employees: 员工人数
//
// 示例:
//
//	params := map[string]string{
//	    "ts_code": "600000.SH",
//	    "exchange": "SSE",
//	}
//	fields := []string{"ts_code", "name", "chairman", "province", "city"}
//	resp, err := client.StockCompany(params, fields)
func (c *Client) StockCompany(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_company",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// NewShare 获取新股上市列表数据
//
// 获取新股上市列表数据，包括新股发行、上市日期等信息。
//
// 参数说明:
//   - start_date: 开始日期
//   - end_date: 结束日期
//   - fields: 可选字段
//   - limit: 单次返回数据长度
//
// 输出 fields:
//   - ts_code: TS代码
//   - subcode: 申购代码
//   - name: 股票名称
//   - ipo_date: 发行公告日
//   - issue_date: 上市日期
//   - issue_price: 发行价格
//   - issue_amount: 发行数量
//   - market: 上市市场
//   - industry: 所属行业
//
// 示例:
//
//	params := map[string]string{
//	    "start_date": "20240101",
//	    "end_date": "20240131",
//	    "limit": "10",
//	}
//	fields := []string{"ts_code", "name", "ipo_date", "issue_date", "issue_price"}
//	resp, err := client.NewShare(params, fields)
func (c *Client) NewShare(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "new_share",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}
