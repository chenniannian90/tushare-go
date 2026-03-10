// CRAWLER IMPLEMENTATION - DO NOT AUTO GENERATE
// 本文件包含爬虫实现的具体逻辑，需要手动维护，禁止代码生成工具覆盖

package stock_market

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	defaultDataSource = "sina" // 默认数据源
)

// GetRealtimeList 获取实时排名列表（爬虫实现）
// src: 数据源，"sina" 或 "dc"（东方财富）
func GetRealtimeList(src string) ([]RealtimeListItem, error) {
	if src == "" {
		src = defaultDataSource
	}

	switch src {
	case defaultDataSource, "sina.com.cn":
		return getSinaRealtimeList()
	case "dc", "eastmoney", "eastmoney.com":
		return getEastmoneyRealtimeList()
	default:
		return getSinaRealtimeList()
	}
}

// getSinaRealtimeList 从新浪获取实时排名
func getSinaRealtimeList() ([]RealtimeListItem, error) {
	// 新浪的所有股票列表接口
	url := "http://hq.sinajs.cn/list=s_sh000001,s_sz399001"

	client := &http.Client{
		Timeout: 15 * time.Second,
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

	// 读取响应
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

	// 解析响应 - 这里简化处理，实际应该获取所有股票列表
	// 实际使用时需要调用大盘指数接口，然后获取所有股票的实时数据
	items, err := parseSinaResponse(responseText)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 转换为 RealtimeListItem
	var listItems []RealtimeListItem
	for _, item := range items {
		listItem := convertQuoteItemToListItem(item)
		listItems = append(listItems, listItem)
	}

	return listItems, nil
}

// getEastmoneyRealtimeList 从东方财富获取实时排名
func getEastmoneyRealtimeList() ([]RealtimeListItem, error) {
	// 东方财富的实时行情接口
	url := "http://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=20&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152"

	client := &http.Client{
		Timeout: 15 * time.Second,
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

	// 读取响应
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

	// 解析东方财富JSON响应
	return parseEastmoneyResponse(responseText)
}

// parseEastmoneyResponse 解析东方财富响应
func parseEastmoneyResponse(text string) ([]RealtimeListItem, error) {
	// 东方财富返回的是JSON格式，需要解析
	// 这里简化处理，实际需要完整的JSON解析
	// 由于JSON解析比较复杂，这里只做示例

	var items []RealtimeListItem

	// 简化示例：创建一个空列表
	// 实际实现需要解析JSON字段：
	// f3-涨跌幅, f4-涨跌额, f5-成交量, f6-成交额, f15-最高价, f16-最低价
	// f17-今开价, f18-昨收价, f20-总市值, f21-流通市值
	// f2-最新价, f12-代码, f14-名称

	return items, nil
}

// GetAllStockList 获取所有A股代码列表（用于实时排名）
func GetAllStockList() ([]string, error) {
	// 这个函数可以从东方财富获取所有A股列表
	url := "http://80.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=5000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f2&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23&fields=f12"

	client := &http.Client{
		Timeout: 15 * time.Second,
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

	// 读取响应
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

	// 解析股票代码列表
	return parseStockCodes(responseText)
}

// parseStockCodes 从响应中解析股票代码
func parseStockCodes(text string) ([]string, error) {
	// 简化实现，提取股票代码
	var codes []string

	// 使用正则表达式提取6位股票代码
	codeReg := regexp.MustCompile(`"f12":"(\d{6})"`)
	matches := codeReg.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) > 1 {
			codes = append(codes, match[1])
		}
	}

	return codes, nil
}

// convertQuoteItemToListItem 将 RealtimeQuoteItem 转换为 RealtimeListItem
//nolint:staticcheck // 这两种类型不能直接转换，需要手动复制字段
func convertQuoteItemToListItem(item RealtimeQuoteItem) RealtimeListItem {
	return RealtimeListItem{
		TsCode:   item.TsCode,
		Name:     item.Name,
		PreClose: item.PreClose,
		Open:     item.Open,
		Price:    item.Price,
		High:     item.High,
		Low:      item.Low,
		Bid:      item.Bid,
		Ask:      item.Ask,
		Volume:   item.Volume,
		Amount:   item.Amount,
		B1V:      item.B1V,
		B1P:      item.B1P,
		B2V:      item.B2V,
		B2P:      item.B2P,
		B3V:      item.B3V,
		B3P:      item.B3P,
		B4V:      item.B4V,
		B4P:      item.B4P,
		B5V:      item.B5V,
		B5P:      item.B5P,
		A1V:      item.A1V,
		A1P:      item.A1P,
		A2V:      item.A2V,
		A2P:      item.A2P,
		A3V:      item.A3V,
		A3P:      item.A3P,
		A4V:      item.A4V,
		A4P:      item.A4P,
		A5V:      item.A5V,
		A5P:      item.A5P,
		Date:     item.Date,
		Time:     item.Time,
	}
}

// RealtimeListCrawlerRequest defines the request for realtime list crawler
type RealtimeListCrawlerRequest struct {
	Src string `json:"src,omitempty"` // Data source ("sina" or "dc")
}

// RealtimeListCrawler is a wrapper function that matches the standard API pattern
func RealtimeListCrawler(ctx context.Context, client interface{}, req *RealtimeListCrawlerRequest) ([]RealtimeListItem, error) {
	return GetRealtimeList(req.Src)
}
