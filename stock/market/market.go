package market

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/realtime"
	"github.com/chenniannian90/tushare-go/types"
	"regexp"
	"strings"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func IsDateFormat(dates ...string) bool {
	pattern := regexp.MustCompile(`^\d{8}$`)
	for _, date := range dates {
		if date == "" {
			continue // Skip empty strings
		}
		if !pattern.MatchString(date) {
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
//	params := map[string]string{
//	    ts_code: "示例值",
//	    trade_date: "示例值",
//	}
//
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
//	params := map[string]string{
//	    ts_code: "示例值",
//	    start_date: "示例值",
//	}
//
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
//	params := map[string]string{
//	    ts_code: "示例值",
//	    start_date: "示例值",
//	}
//
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
	if hasTsCode {
		argsCount++
	}
	if hasTradeDate {
		argsCount++
	}
	if hasResumeDate {
		argsCount++
	}
	if argsCount != 1 {
		return nil, fmt.Errorf("need one argument among ts_code, suspend_date, resume_date")
	}
	if !IsDateFormat(params["suspend_date"], params["resume_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "suspend", "token": c.getToken(), "params": params, "fields": fields})
}

// RealTimeQuote 获取实时行情数据（直接调用新浪财经接口，免费无需 token）

// 参数说明:
//   - ts_code: 股票代码（必填），支持多个代码用逗号分隔，如 "000001.SZ,600000.SH"

// 输出字段:
//   返回 SinaQuoteResponse，包含完整的实时行情数据

// 特性:
//   - 无需 Tushare token，完全免费
//   - 数据来源：新浪财经
//   - 支持批量查询多个股票
//   - 包含五档行情数据

// 示例:

// // 查询单个股��
// resp, err := client.RealTimeQuote(map[string]string{"ts_code": "000001.SZ"}, nil)
//
// // 查询多个股票
// resp, err := client.RealTimeQuote(map[string]string{"ts_code": "000001.SZ,600000.SH"}, nil)
//
// // 遍历结果
// for code, quote := range resp.Quotes {
//     fmt.Printf("%s (%s): %.2f (%.2f%%)\n",
//         quote.Name, code, quote.Price, quote.GetChangePercent())
// }
func (c *Client) RealTimeQuote(params map[string]string, fields []string) (*realtime.SinaQuoteResponse, error) {
	tsCode, ok := params["ts_code"]
	if !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}

	// 解析股票代码（支持逗号分隔的多个代码）
	codes := strings.Split(tsCode, ",")

	// 去除空白
	cleanCodes := make([]string, 0, len(codes))
	for _, code := range codes {
		code = strings.TrimSpace(code)
		if code != "" {
			cleanCodes = append(cleanCodes, code)
		}
	}

	if len(cleanCodes) == 0 {
		return nil, fmt.Errorf("无效的股票代码")
	}

	// 调用新浪财经接口
	client := realtime.NewSinaClient()
	return client.GetRealtimeQuote(cleanCodes...)
}

// RealTimeTick 获取实时分笔成交数据（直接调用新浪财经，免费无需 token）

// 参数说明:
//   - ts_code: 股票代码（必填）

// 输出字段:
//   返回 SinaQuoteResponse，包含该股票的实时五档行情数据

// 注意:
//   - 新浪财经的分笔数据包含在实时行情中
//   - 返回数据包含：买一到买五、卖一到卖五的价格和量

// 示例:
//
// // 查询平安银行的分笔数据
// resp, err := client.RealTimeTick(map[string]string{"ts_code": "000001.SZ"}, nil)
//
// // 查看五档行情
// for code, quote := range resp.Quotes {
//     fmt.Printf("买一: %.2f × %d 手\n", quote.BidPrice1, quote.BidVol1)
//     fmt.Printf("卖一: %.2f × %d 手\n", quote.AskPrice1, quote.AskVol1)
// }
func (c *Client) RealTimeTick(params map[string]string, fields []string) (*realtime.SinaQuoteResponse, error) {
	tsCode, ok := params["ts_code"]
	if !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}

	// 解析股票代码（支持逗号分隔的多个代码）
	codes := strings.Split(tsCode, ",")
	cleanCodes := make([]string, 0, len(codes))
	for _, code := range codes {
		code = strings.TrimSpace(code)
		if code != "" {
			cleanCodes = append(cleanCodes, code)
		}
	}

	if len(cleanCodes) == 0 {
		return nil, fmt.Errorf("无效的股票代码")
	}

	// 调用新浪财经接口
	client := realtime.NewSinaClient()
	return client.GetRealtimeTick(cleanCodes...)
}

// RealTimeList 获取实时行情列表数据（直接调用新浪财经，免费无需 token）

// 参数说明:
//   - 无需参数，或可指定：
//     - limit: 返回股票数量（可选，默认返回主要指数和热门股）

// 输出字段:
//   返回 SinaQuoteResponse，包含多个股票/指数的实时行情

// 注意:
//   - 新浪财经不提供全市场列表接口
//   - 此函数返回主要市场指数和热门股票
//   - 包含：上证指数、深证成指、创业板指、科创50等

// 示例:
//
// // 获取市场主要指数和热门股
// resp, err := client.RealTimeList(map[string]string{"limit": "10"}, nil)
//
// // 遍历结果
// for code, quote := range resp.Quotes {
//     fmt.Printf("%s (%s): %.2f (%.2f%%)\n",
//         quote.Name, code, quote.Price, quote.GetChangePercent())
// }
//
// // 获取涨幅榜
// topGainers := resp.GetTopGainers(5)
// for _, code := range topGainers {
//     quote := resp.Quotes[code]
//     fmt.Printf("%s: %.2f%%\n", quote.Name, quote.GetChangePercent())
// }
func (c *Client) RealTimeList(params map[string]string, fields []string) (*realtime.SinaQuoteResponse, error) {
	// 获取 limit 参数
	limit := 0
	if l, ok := params["limit"]; ok {
		// 尝试解析 limit 参数
		var parsedLimit int
		n, err := fmt.Sscanf(l, "%d", &parsedLimit)
		if err == nil && n == 1 {
			limit = parsedLimit
		}
		// 如果解析失败，limit 保持为 0
	}

	// 调用新浪财经接口
	client := realtime.NewSinaClient()
	return client.GetRealtimeList(limit)
}

// RealTimeQuoteSina 使用新浪数据源获取实时行情数据（免费，无需 token）
//
// 数据来源: 新浪财经 (https://finance.sina.com.cn)
// 接口地址: https://hq.sinajs.cn/list=sh600000
//
// 参数说明:
//   - codes: 股票代码列表，支持以下格式：
//     * "sh600000" (上交所，新浪原生格式)
//     * "sz000001" (深交所，新浪原生格式)
//     * "000001.SZ" (Tushare 格式，自动转换)
//     * "600000.SH" (Tushare 格式，自动转换)
//     * "000001" 或 "600000" (6位代码，自动判断市场)
//
// 输出字段:
//   返回 SinaQuoteResponse，包含以下信息：
//   - name: 股票名称
//   - open: 开盘价
//   - pre_close: 前收盘价
//   - price: 当前价
//   - high: 最高价
//   - low: 最低价
//   - bid/ask: 买一/卖一价
//   - volume: 成交量（手）
//   - amount: 成交额（元）
//   - 买一到买五、卖一到卖五的价格和量
//   - date/time: 日期和时间
//
// 特性:
//   - 无需 Tushare token，完全免费
//   - 支持批量查询多个股票
//   - 包含五档行情数据
//   - 支持涨跌幅排序
//
// 示例:
//
//	// 查询单个股票
//	resp, err := client.RealTimeQuoteSina("000001.SZ")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for code, quote := range resp.Quotes {
//	    fmt.Printf("%s (%s): %.2f (%.2f%%)\n",
//	        quote.Name, code, quote.Price, quote.GetChangePercent())
//	}
//
//	// 批量查询
//	resp, err := client.RealTimeQuoteSina("000001.SZ", "600000.SH", "sz000002")
//
//	// 获取涨幅榜
//	topGainers := resp.GetTopGainers(5)
//	for _, code := range topGainers {
//	    quote := resp.Quotes[code]
//	    fmt.Printf("%s: %.2f%%\n", quote.Name, quote.GetChangePercent())
//	}
func (c *Client) RealTimeQuoteSina(codes ...string) (*realtime.SinaQuoteResponse, error) {
	if len(codes) == 0 {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	client := realtime.NewSinaClient()
	return client.GetRealtimeQuote(codes...)
}

// GetBatchRealtimeQuote 批量获取实时行情（自动分批处理）
//
// 由于新浪 API 对单次请求的股票数量有限制，此方法会自动将大量股票分批请求
//
// 参数说明:
//   - codes: 股票代码列表
//   - batchSize: 每批次的股票数量（建议 100-200）
//
// 示例:
//
//	codes := []string{"000001.SZ", "000002.SZ", "600000.SH", ...}
//	resp, err := client.GetBatchRealtimeQuote(codes, 100)
func (c *Client) GetBatchRealtimeQuote(codes []string, batchSize int) (*realtime.SinaQuoteResponse, error) {
	if len(codes) == 0 {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	if batchSize <= 0 {
		batchSize = 100 // 默认每批 100 个
	}

	client := realtime.NewSinaClient()
	result := &realtime.SinaQuoteResponse{
		Quotes: make(map[string]*realtime.SinaStockQuote),
	}

	// 分批处理
	for i := 0; i < len(codes); i += batchSize {
		end := i + batchSize
		if end > len(codes) {
			end = len(codes)
		}

		batch := codes[i:end]
		resp, err := client.GetRealtimeQuote(batch...)
		if err != nil {
			return nil, fmt.Errorf("批次 %d-%d 请求失败: %w", i, end, err)
		}

		// 合并结果
		for code, quote := range resp.Quotes {
			result.Quotes[code] = quote
		}
	}

	return result, nil
}
