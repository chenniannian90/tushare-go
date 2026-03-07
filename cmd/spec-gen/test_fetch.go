package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://tushare.pro/document/2?doc_id=25"

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %d\n", resp.StatusCode)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return
	}

	// 查找所有 h3 标签
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("h3[%d]: %s\n", i, s.Text())
	})

	// 查找包含"接口："的段落
	foundAPI := false
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "接口：") {
			fmt.Printf("\nFound API name in paragraph %d:\n%s\n\n", i, text)
			foundAPI = true
		}
		if strings.Contains(text, "描述：") {
			fmt.Printf("\nFound description in paragraph %d:\n%s\n\n", i, text)
		}
	})

	if !foundAPI {
		fmt.Println("\nNo API name found!")
	}

	// 查找表格
	tableCount := 0
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		tableCount++
		fmt.Printf("\nTable %d found\n", i)
		// 打印前几行
		s.Find("tr").Each(func(j int, row *goquery.Selection) {
			if j < 3 {
				row.Find("th, td").Each(func(k int, cell *goquery.Selection) {
					fmt.Printf("  Cell[%d][%d]: %s\n", j, k, strings.TrimSpace(cell.Text()))
				})
			}
		})
	})

	fmt.Printf("\nTotal tables found: %d\n", tableCount)
}
