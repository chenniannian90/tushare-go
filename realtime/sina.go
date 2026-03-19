package realtime

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// SinaStockQuote 新浪股票实时行情数据结构
// 参考: http://finance.sina.com.cn/realstock/company/sh600000/nc.shtml
type SinaStockQuote struct {
	Name         string  `json:"name"`           // 股票名称
	Open         float64 `json:"open"`           // 开盘价
	PreClose     float64 `json:"pre_close"`      // 前收盘价
	Price        float64 `json:"price"`          // 当前价
	High         float64 `json:"high"`           // 最高价
	Low          float64 `json:"low"`            // 最低价
	Bid          float64 `json:"bid"`            // 买一价
	Ask          float64 `json:"ask"`            // 卖一价
	Volume       int64   `json:"volume"`         // 成交量（手）
	Amount       float64 `json:"amount"`          // 成交额（元）
	BidVol1      int64   `json:"bid_vol_1"`      // 买一量（手）
	BidVol2      int64   `json:"bid_vol_2"`      // 买二量（手）
	BidVol3      int64   `json:"bid_vol_3"`      // 买三量（手）
	BidVol4      int64   `json:"bid_vol_4"`      // 买四量（手）
	BidVol5      int64   `json:"bid_vol_5"`      // 买五量（手）
	BidPrice1    float64 `json:"bid_price_1"`    // 买一价
	BidPrice2    float64 `json:"bid_price_2"`    // 买二价
	BidPrice3    float64 `json:"bid_price_3"`    // 买三价
	BidPrice4    float64 `json:"bid_price_4"`    // 买四价
	BidPrice5    float64 `json:"bid_price_5"`    // 买五价
	AskVol1      int64   `json:"ask_vol_1"`      // 卖一量（手）
	AskVol2      int64   `json:"ask_vol_2"`      // 卖二量（手）
	AskVol3      int64   `json:"ask_vol_3"`      // 卖三量（手）
	AskVol4      int64   `json:"ask_vol_4"`      // 卖四量（手）
	AskVol5      int64   `json:"ask_vol_5"`      // 卖五量（手）
	AskPrice1    float64 `json:"ask_price_1"`    // 卖一价
	AskPrice2    float64 `json:"ask_price_2"`    // 卖二价
	AskPrice3    float64 `json:"ask_price_3"`    // 卖三价
	AskPrice4    float64 `json:"ask_price_4"`    // 卖四价
	AskPrice5    float64 `json:"ask_price_5"`    // 卖五价
	Date         string  `json:"date"`           // 日期 YYYY-MM-DD
	Time         string  `json:"time"`           // 时间 HH:MM:SS
	Timestamp    int64   `json:"timestamp"`      // 时间戳
}

// SinaQuoteResponse 新浪实时行情响应
type SinaQuoteResponse struct {
	Quotes map[string]*SinaStockQuote `json:"quotes"`
	Error  error                       `json:"error,omitempty"`
}

// SinaClient 新浪行情客户端
type SinaClient struct {
	client *http.Client
}

// NewSinaClient 创建新浪行情客户端
func NewSinaClient() *SinaClient {
	return &SinaClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetRealtimeQuote 获取实时行情
// codes: 股票代码列表，支持以下格式：
//   - "sh600000" (上交所)
//   - "sz000001" (深交所)
//   - "000001.SZ" (自动转换)
//   - "600000.SH" (自动转换)
func (c *SinaClient) GetRealtimeQuote(codes ...string) (*SinaQuoteResponse, error) {
	if len(codes) == 0 {
		return nil, fmt.Errorf("股票代码不能为空")
	}

	// 转换股票代码格式
	sinaCodes := make([]string, 0, len(codes))
	for _, code := range codes {
		sinaCode := convertToSinaCode(code)
		if sinaCode != "" {
			sinaCodes = append(sinaCodes, sinaCode)
		}
	}

	if len(sinaCodes) == 0 {
		return nil, fmt.Errorf("无效的股票代码")
	}

	// 构建请求 URL
	url := fmt.Sprintf("https://hq.sinajs.cn/list=%s", strings.Join(sinaCodes, ","))

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置必需的请求头
	req.Header.Set("Referer", "https://finance.sina.com.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	// 发送请求
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析响应
	quotes, err := parseSinaResponse(string(body))
	if err != nil {
		return &SinaQuoteResponse{
			Quotes: quotes,
			Error:  err,
		}, err
	}

	return &SinaQuoteResponse{
		Quotes: quotes,
	}, nil
}

// convertToSinaCode 将各种格式的股票代码转换为新浪格式
func convertToSinaCode(code string) string {
	code = strings.TrimSpace(code)

	if code == "" {
		return ""
	}

	// 已经是新浪格式（小写检查）
	lowerCode := strings.ToLower(code)
	if strings.HasPrefix(lowerCode, "sh") || strings.HasPrefix(lowerCode, "sz") || strings.HasPrefix(lowerCode, "bj") {
		return lowerCode
	}

	// 转换为大写继续处理
	code = strings.ToUpper(code)

	// Tushare 格式: 000001.SZ 或 600000.SH
	if strings.Contains(code, ".") {
		parts := strings.Split(code, ".")
		if len(parts) != 2 {
			return ""
		}

		market := strings.ToUpper(parts[1])
		stockCode := parts[0]

		switch market {
		case "SH":
			return "sh" + stockCode
		case "SZ":
			return "sz" + stockCode
		case "BJ":
			return "bj" + stockCode
		default:
			return ""
		}
	}

	// 6位数字代码，根据首位数字判断市场
	if len(code) == 6 {
		switch code[0] {
		case '6': // 上海证券交易所
			return "sh" + code
		case '0', '3': // 深圳证券交易所
			return "sz" + code
		case '8', '4': // 北京证券交易所
			return "bj" + code
		default:
			return ""
		}
	}

	// 指数代码（6位代码，如000001, 399001等）
	// 上证指数: 000001, 深证成指: 399001, 创业板指: 399006等
	if len(code) == 6 && code[0:1] >= "0" && code[0:1] <= "9" {
		// 根据指数代码判断市场
		switch code[0:3] {
		case "000": // 上证指数、沪深300等
			if code == "000001" || code == "000300" || code == "000016" || code == "000905" {
				return "sh" + code
			}
			// 其他000开头可能是深圳
			return "sz" + code
		case "399": // 深证指数
			return "sz" + code
		default:
			// 默认判断
			if code[0:1] == "0" || code[0:1] == "3" {
				return "sz" + code
			}
			return "sh" + code
		}
	}

	return ""
}

// parseSinaResponse 解析新浪 API 响应
func parseSinaResponse(response string) (map[string]*SinaStockQuote, error) {
	quotes := make(map[string]*SinaStockQuote)

	lines := strings.Split(response, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 解析格式: var hq_str_sh600000="平安银行,11.25,11.20,11.17,11.19,11.20,11.21,117850..."
		if !strings.Contains(line, "hq_str_") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		// 提取股票代码
		codePart := strings.TrimPrefix(parts[0], "var hq_str_")
		if codePart == parts[0] {
			continue
		}
		code := codePart

		// 提取数据部分
		dataStr := strings.Trim(parts[1], `";`)
		if dataStr == "" {
			continue
		}

		// 解析数据
		fields := strings.Split(dataStr, ",")
		if len(fields) < 32 {
			continue // 数据不完整，跳过
		}

		quote := &SinaStockQuote{}

		// 解析各字段
		// 字段说明（以实际返回为准）：
		// 0: 股票名称
		// 1: 开盘价
		// 2: 前收盘价
		// 3: 当前价
		// 4: 最高价
		// 5: 最低价
		// 6: 买一价
		// 7: 卖一价
		// 8: 成交量（手）
		// 9: 成交额（元）
		// 10-19: 买一到买五的量和价格
		// 20-29: 卖一到卖五的量和价格
		// 30: 日期
		// 31: 时间

		quote.Name = fields[0]
		quote.Open = parseFloat(fields[1])
		quote.PreClose = parseFloat(fields[2])
		quote.Price = parseFloat(fields[3])
		quote.High = parseFloat(fields[4])
		quote.Low = parseFloat(fields[5])
		quote.Bid = parseFloat(fields[6])
		quote.Ask = parseFloat(fields[7])
		quote.Volume = parseInt64(fields[8])
		quote.Amount = parseFloat(fields[9])

		// 买盘
		if len(fields) > 10 {
			quote.BidPrice1 = parseFloat(fields[10])
			quote.BidVol1 = parseInt64(fields[11])
		}
		if len(fields) > 12 {
			quote.BidPrice2 = parseFloat(fields[12])
			quote.BidVol2 = parseInt64(fields[13])
		}
		if len(fields) > 14 {
			quote.BidPrice3 = parseFloat(fields[14])
			quote.BidVol3 = parseInt64(fields[15])
		}
		if len(fields) > 16 {
			quote.BidPrice4 = parseFloat(fields[16])
			quote.BidVol4 = parseInt64(fields[17])
		}
		if len(fields) > 18 {
			quote.BidPrice5 = parseFloat(fields[18])
			quote.BidVol5 = parseInt64(fields[19])
		}

		// 卖盘
		if len(fields) > 20 {
			quote.AskPrice1 = parseFloat(fields[20])
			quote.AskVol1 = parseInt64(fields[21])
		}
		if len(fields) > 22 {
			quote.AskPrice2 = parseFloat(fields[22])
			quote.AskVol2 = parseInt64(fields[23])
		}
		if len(fields) > 24 {
			quote.AskPrice3 = parseFloat(fields[24])
			quote.AskVol3 = parseInt64(fields[25])
		}
		if len(fields) > 26 {
			quote.AskPrice4 = parseFloat(fields[26])
			quote.AskVol4 = parseInt64(fields[27])
		}
		if len(fields) > 28 {
			quote.AskPrice5 = parseFloat(fields[28])
			quote.AskVol5 = parseInt64(fields[29])
		}

		// 日期和时间
		if len(fields) > 30 {
			quote.Date = fields[30]
		}
		if len(fields) > 31 {
			quote.Time = fields[31]
		}

		// 生成时间戳
		if quote.Date != "" && quote.Time != "" {
			dateTime := fmt.Sprintf("%s %s", quote.Date, quote.Time)
			t, err := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.Local)
			if err == nil {
				quote.Timestamp = t.Unix()
			}
		}

		quotes[code] = quote
	}

	return quotes, nil
}

// parseFloat 安全解析浮点数
func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" || s == "-" {
		return 0
	}
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		return 0
	}
	return f
}

// parseInt64 安全解析整数
func parseInt64(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" || s == "-" {
		return 0
	}
	var i int64
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		return 0
	}
	return i
}

// ToJSON 转换为 JSON 格式
func (q *SinaStockQuote) ToJSON() string {
	data, _ := json.MarshalIndent(q, "", "  ")
	return string(data)
}

// ToMap 转换为 Map 格式（便于与其他数据格式兼容）
func (q *SinaStockQuote) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":         q.Name,
		"open":         q.Open,
		"pre_close":    q.PreClose,
		"price":        q.Price,
		"high":         q.High,
		"low":          q.Low,
		"bid":          q.Bid,
		"ask":          q.Ask,
		"volume":       q.Volume,
		"amount":       q.Amount,
		"bid_vol_1":    q.BidVol1,
		"bid_price_1":  q.BidPrice1,
		"ask_vol_1":    q.AskVol1,
		"ask_price_1":  q.AskPrice1,
		"date":         q.Date,
		"time":         q.Time,
		"timestamp":    q.Timestamp,
	}
}

// GetChange 获取涨跌额
func (q *SinaStockQuote) GetChange() float64 {
	return q.Price - q.PreClose
}

// GetChangePercent 获取涨跌幅（百分比）
func (q *SinaStockQuote) GetChangePercent() float64 {
	if q.PreClose == 0 {
		return 0
	}
	return (q.Price - q.PreClose) / q.PreClose * 100
}

// IsTrading 判断是否在交易时间
func (q *SinaStockQuote) IsTrading() bool {
	if q.Time == "" {
		return false
	}

	// 简单判断：如果有时间且有价格变化，认为在交易
	now := time.Now()
	hour := now.Hour()
	minute := now.Minute()
	weekday := now.Weekday()

	// 周末不交易
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	// 交易时间段判断（简化版）
	// 上午: 9:30-11:30
	// 下午: 13:00-15:00
	morning := (hour == 9 && minute >= 30) || (hour == 10) || (hour == 11 && minute <= 30)
	afternoon := (hour == 13) || (hour == 14) || (hour == 15 && minute == 0)

	return morning || afternoon
}

// SortByChangePercent 按涨跌幅排序
func SortByChangePercent(quotes map[string]*SinaStockQuote, descending bool) []string {
	type kv struct {
		Key   string
		Value float64
	}

	var ss []kv
	for k, v := range quotes {
		ss = append(ss, kv{k, v.GetChangePercent()})
	}

	sort.Slice(ss, func(i, j int) bool {
		if descending {
			return ss[i].Value > ss[j].Value
		}
		return ss[i].Value < ss[j].Value
	})

	result := make([]string, len(ss))
	for i, kv := range ss {
		result[i] = kv.Key
	}
	return result
}

// GetTopGainers 获取涨幅榜前 N 名
func (r *SinaQuoteResponse) GetTopGainers(n int) []string {
	if r.Quotes == nil {
		return []string{}
	}
	sorted := SortByChangePercent(r.Quotes, true)
	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}

// GetTopLosers 获取跌幅榜前 N 名
func (r *SinaQuoteResponse) GetTopLosers(n int) []string {
	if r.Quotes == nil {
		return []string{}
	}
	sorted := SortByChangePercent(r.Quotes, false)
	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}

// ToJSON 转换整个响应为 JSON
func (r *SinaQuoteResponse) ToJSON() string {
	data, _ := json.MarshalIndent(r, "", "  ")
	return string(data)
}

// GetRealtimeTick 获取实时分笔成交数据
// 注意：新浪财经的分笔数据包含在实时行情中，此函数返回相同的五档行情数据
func (c *SinaClient) GetRealtimeTick(codes ...string) (*SinaQuoteResponse, error) {
	return c.GetRealtimeQuote(codes...)
}

// GetRealtimeList 获取实时行情列表数据
// 新浪财经不提供全市场列表，返回主要指数和热门股票
func (c *SinaClient) GetRealtimeList(limit int) (*SinaQuoteResponse, error) {
	// 新浪财经不提供全市场列表，返回主要指数和热门股票
	defaultCodes := []string{
		// 主要指数
		"sh000001", // 上证指数
		"sz399001", // 深证成指
		"sz399006", // 创业板指
		"sh000688", // 科创50
		"sz399300", // 沪深300
		// 热门蓝筹股
		"sh600519", // 贵州茅台
		"sz000001", // 平安银行
		"sz000002", // 万科A
		"sh600036", // 招商银行
		"sz300750", // 宁德时代
		"sh601318", // 中国平安
		"sh600000", // 浦发银行
		"sz000858", // 五粮液
		"sh600276", // 恒瑞医药
		"sz300059", // 东方财富
		"sh601012", // 隆基绿能
		"sz002594", // 比亚迪
		"sh600900", // 长江电力
		"sz002475", // 立讯精密
		"sh601888", // 中国中免
	}

	// 如果指定了 limit，截取相应的数量
	if limit > 0 && limit < len(defaultCodes) {
		defaultCodes = defaultCodes[:limit]
	}

	return c.GetRealtimeQuote(defaultCodes...)
}

// TestConvertToSinaCode 测试用的代码转换函数（仅用于测试）
func TestConvertToSinaCode(code string) string {
	return convertToSinaCode(code)
}
