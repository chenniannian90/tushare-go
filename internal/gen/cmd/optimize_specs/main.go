package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Spec 规范文件结构
type Spec struct {
	APIName        string            `json:"api_name"`
	APICode        string            `json:"api_code"`
	Description    string            `json:"description"`
	Describe       DescribeInfo      `json:"__describe__"`
	RequestParams  []APIParam        `json:"request_params"`
	ResponseFields []APIField        `json:"response_fields"`
}

type DescribeInfo struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type APIParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type APIField struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	DefaultDisplay bool   `json:"default_display"`
	Description    string `json:"description"`
}

func main() {
	specsDir := "/Users/mac-new/go/src/tushare-go/internal/gen/specs"

	err := filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		fmt.Printf("处理文件: %s\n", path)

		// 读取原始文件
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("  ❌ 读取失败: %v\n", err)
			return nil
		}

		var spec Spec
		if err := json.Unmarshal(data, &spec); err != nil {
			fmt.Printf("  ❌ 解析失败: %v\n", err)
			return nil
		}

		// 如果已经有完整的参数，跳过
		if len(spec.RequestParams) > 0 || len(spec.ResponseFields) > 0 {
			fmt.Printf("  ⏭️  已有参数，跳过\n")
			return nil
		}

		// 从文档抓取信息
		docURL := spec.Describe.URL
		if docURL == "" {
			fmt.Printf("  ⏭️  没有 URL，跳过\n")
			return nil
		}

		optimized, err := fetchFromTushare(docURL)
		if err != nil {
			fmt.Printf("  ❌ 抓取失败: %v\n", err)
			return nil
		}

		// 更新 spec
		spec.APIName = optimized.APIName
		spec.APICode = optimized.APICode
		if optimized.Description != "" {
			spec.Description = optimized.Description
		}
		spec.RequestParams = optimized.RequestParams
		spec.ResponseFields = optimized.ResponseFields

		// 写回文件
		newData, err := json.MarshalIndent(spec, "", "  ")
		if err != nil {
			fmt.Printf("  ❌ 序列化失败: %v\n", err)
			return nil
		}

		if err := os.WriteFile(path, newData, 0644); err != nil {
			fmt.Printf("  ❌ 写入失败: %v\n", err)
			return nil
		}

		fmt.Printf("  ✅ 优化完成\n")
		time.Sleep(1 * time.Second) // 避免请求过快

		return nil
	})

	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ 批量优化完成!")
}

func fetchFromTushare(url string) (*Spec, error) {
	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 发送请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 状态码: %d", resp.StatusCode)
	}

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析 HTML 失败: %w", err)
	}

	spec := &Spec{}

	// 提取 API 名称
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "接口：") {
			re := regexp.MustCompile(`接口：\s*(\w+)`)
			matches := re.FindStringSubmatch(text)
			if len(matches) > 1 {
				spec.APIName = matches[1]
				spec.APICode = matches[1]
			}
		}
	})

	// 提取描述
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "描述：") && spec.Description == "" {
			re := regexp.MustCompile(`描述：(.+?)(?:限量：|$)`)
			matches := re.FindStringSubmatch(text)
			if len(matches) > 1 {
				spec.Description = strings.TrimSpace(matches[1])
			}
		}
	})

	// 提取输入参数
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "输入参数") {
			// 找到下一个 table
			table := s.Next().Find("table")
			if table.Length() > 0 {
				spec.RequestParams = extractTableParams(table, "input")
			}
		}
	})

	// 提取输出参数
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "输出参数") {
			// 找到下一个 table
			table := s.Next().Find("table")
			if table.Length() > 0 {
				spec.ResponseFields = extractTableParams(table, "output")
			}
		}
	})

	return spec, nil
}

func extractTableParams(table *goquery.Selection, paramType string) []APIParam {
	var params []APIParam

	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		// 跳过表头
		if i == 0 {
			return
		}

		cells := row.Find("td")
		if cells.Length() < 3 {
			return
		}

		param := APIParam{
			Name:        strings.TrimSpace(cells.Eq(0).Text()),
			Type:        strings.TrimSpace(cells.Eq(1).Text()),
			Description: strings.TrimSpace(cells.Eq(3).Text()),
		}

		if paramType == "input" {
			required := strings.TrimSpace(cells.Eq(2).Text())
			param.Required = required == "Y" || strings.ToLower(required) == "true"
		}

		if param.Name != "" {
			params = append(params, param)
		}
	})

	return params
}