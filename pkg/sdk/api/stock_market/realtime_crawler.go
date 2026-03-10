// CRAWLER IMPLEMENTATION - DO NOT AUTO GENERATE
// 本文件包含爬虫实现的具体逻辑，需要手动维护，禁止代码生成工具覆盖

package stock_market

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// GetRealtimeQuotes 获取实时行情数据（爬虫实现）
// symbols: 股票代码列表，如 ["000001", "600000"] 或单个 "000001"
func GetRealtimeQuotes(symbols interface{}) ([]RealtimeQuoteItem, error) {
	// 构建股票代码列表
	symbolsList := buildSymbolsList(symbols)
	if symbolsList == "" {
		return nil, fmt.Errorf("无效的股票代码")
	}

	// 构建请求URL
	randomStr := generateRandomString(13)
	url := fmt.Sprintf("http://hq.sinajs.cn/rn=%s&list=%s", randomStr, symbolsList)

	// 发送HTTP请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应数据
	reader := bufio.NewReader(resp.Body)
	var content strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("读取响应失败: %v", err)
		}
		content.WriteString(line)
	}

	responseText := content.String()

	// 解析响应数据
	items, err := parseSinaResponse(responseText)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return items, nil
}

// buildSymbolsList 构建股票代码列表字符串
func buildSymbolsList(symbols interface{}) string {
	var codes []string

	switch v := symbols.(type) {
	case string:
		codes = append(codes, v)
	case []string:
		codes = append(codes, v...)
	default:
		return ""
	}

	var symbolsList strings.Builder
	for _, code := range codes {
		symbol := codeToSymbol(code)
		symbolsList.WriteString(symbol)
		symbolsList.WriteString(",")
	}

	result := symbolsList.String()
	// 去掉末尾的逗号
	if len(result) > 0 && result[len(result)-1] == ',' {
		result = result[:len(result)-1]
	}

	return result
}

// codeToSymbol 将股票代码转换为新浪格式
func codeToSymbol(code string) string {
	// 如果已经是新浪格式，直接返回
	if strings.HasPrefix(code, "sh") || strings.HasPrefix(code, "sz") || strings.HasPrefix(code, "gb_") {
		return code
	}

	// 6位代码，转换
	if len(code) == 6 {
		firstChar := string(code[0])
		if firstChar == "5" || firstChar == "6" || firstChar == "9" ||
			strings.HasPrefix(code, "11") || strings.HasPrefix(code, "13") {
			return "sh" + code
		}
		return "sz" + code
	}

	return code
}

// generateRandomString 生成随机字符串（用于避免缓存）
func generateRandomString(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// parseSinaResponse 解析新浪响应数据
func parseSinaResponse(text string) ([]RealtimeQuoteItem, error) {
	// 正则提取数据
	dataReg := regexp.MustCompile(`="(.*?)";`)
	dataMatches := dataReg.FindAllStringSubmatch(text, -1)

	// 正则提取股票代码
	symReg := regexp.MustCompile(`(?:sh|sz|gb_)(.*?)=`)
	symMatches := symReg.FindAllStringSubmatch(text, -1)

	if len(symMatches) == 0 {
		return nil, fmt.Errorf("没有找到股票数据")
	}

	var items []RealtimeQuoteItem

	for i, dataMatch := range dataMatches {
		if len(dataMatch) < 2 {
			continue
		}

		dataStr := dataMatch[1]
		if len(dataStr) == 0 {
			continue
		}

		// 分割数据
		fields := strings.Split(dataStr, ",")

		// 提取股票代码
		var code string
		if i < len(symMatches) && len(symMatches[i]) > 1 {
			code = symMatches[i][1]
		}

		// 根据字段数量判断是A股还是美股
		var item RealtimeQuoteItem
		var err error

		if len(fields) == 28 {
			// 美股数据
			item, err = parseUSStockData(fields, code)
		} else {
			// A股数据
			item, err = parseAStockData(fields, code)
		}

		if err != nil {
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

// parseAStockData 解析A股数据
func parseAStockData(fields []string, code string) (RealtimeQuoteItem, error) {
	if len(fields) < 32 {
		return RealtimeQuoteItem{}, fmt.Errorf("A股数据字段不足")
	}

	item := RealtimeQuoteItem{}
	item.TsCode = code

	// 字段映射
	fieldMap := map[string]string{
		"name":      fields[0],
		"open":      fields[1],
		"pre_close": fields[2],
		"price":     fields[3],
		"high":      fields[4],
		"low":       fields[5],
		"bid":       fields[6],
		"ask":       fields[7],
		"volume":    fields[8],
		"amount":    fields[9],
		"date":      fields[30],
		"time":      fields[31],
	}

	// 解析基础字段
	if val, ok := fieldMap["name"]; ok && val != "" {
		item.Name = val
	}
	item.Open = parseFloat(fieldMap["open"])
	item.PreClose = parseFloat(fieldMap["pre_close"])
	item.Price = parseFloat(fieldMap["price"])
	item.High = parseFloat(fieldMap["high"])
	item.Low = parseFloat(fieldMap["low"])
	item.Bid = parseFloat(fieldMap["bid"])
	item.Ask = parseFloat(fieldMap["ask"])

	// 处理成交量（去掉末尾两位，手转换为股）
	volumeStr := fieldMap["volume"]
	if len(volumeStr) >= 2 {
		volumeStr = volumeStr[:len(volumeStr)-2]
	}
	item.Volume = parseInt(volumeStr)

	item.Amount = parseFloat(fieldMap["amount"])
	item.Date = fieldMap["date"]
	item.Time = fieldMap["time"]

	// 解析五档盘口数据 (索引10-29)
	if len(fields) >= 30 {
		// 买一到买五
		item.B1V = parseVolumeField(fields[10])
		item.B1P = parseFloat(fields[11])
		item.B2V = parseVolumeField(fields[12])
		item.B2P = parseFloat(fields[13])
		item.B3V = parseVolumeField(fields[14])
		item.B3P = parseFloat(fields[15])
		item.B4V = parseVolumeField(fields[16])
		item.B4P = parseFloat(fields[17])
		item.B5V = parseVolumeField(fields[18])
		item.B5P = parseFloat(fields[19])

		// 卖一到卖五
		item.A1V = parseVolumeField(fields[20])
		item.A1P = parseFloat(fields[21])
		item.A2V = parseVolumeField(fields[22])
		item.A2P = parseFloat(fields[23])
		item.A3V = parseVolumeField(fields[24])
		item.A3P = parseFloat(fields[25])
		item.A4V = parseVolumeField(fields[26])
		item.A4P = parseFloat(fields[27])
		item.A5V = parseVolumeField(fields[28])
		item.A5P = parseFloat(fields[29])
	}

	return item, nil
}

// parseUSStockData 解析美股数据
func parseUSStockData(fields []string, code string) (RealtimeQuoteItem, error) {
	if len(fields) < 28 {
		return RealtimeQuoteItem{}, fmt.Errorf("美股数据字段不足")
	}

	item := RealtimeQuoteItem{}
	item.TsCode = code
	item.Name = fields[0]
	item.Price = parseFloat(fields[1])

	// 美股数据字段映射与A股不同，这里简化处理
	// 实际使用时需要根据usLiveDataCols进行完整映射

	return item, nil
}

// parseFloat 解析浮点数
func parseFloat(s string) float64 {
	var f float64
	_, _ = fmt.Sscanf(s, "%f", &f)
	return f
}

// parseInt 解析整数
func parseInt(s string) int {
	var i int
	_, _ = fmt.Sscanf(s, "%d", &i)
	return i
}

// parseVolumeField 解析量字段（去掉末尾两位）
func parseVolumeField(s string) int {
	if len(s) >= 2 {
		s = s[:len(s)-2]
	}
	return parseInt(s)
}

// RealtimeCrawlerRequest defines the request for realtime crawler
type RealtimeCrawlerRequest struct {
	Symbols interface{} `json:"symbols,omitempty"` // Stock codes (string or []string)
}

// RealtimeCrawler is a wrapper function that matches the standard API pattern
func RealtimeCrawler(ctx context.Context, client interface{}, req *RealtimeCrawlerRequest) ([]RealtimeQuoteItem, error) {
	return GetRealtimeQuotes(req.Symbols)
}
