package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	stock_boardapi "tushare-go/pkg/sdk/api/stock_board"
	"tushare-go/pkg/sdk"
)

type BoardTestResult struct {
	APIName string
	Success bool
	HasData bool
	Count   int
	Error   string
}

func main() {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		token = "412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
	}

	config := &sdk.Config{
		Tokens: []string{token},
		Endpoint: "https://api.tushare.pro",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	client := sdk.NewClient(config)
	ctx := context.Background()

	results := make([]BoardTestResult, 0, 22)

	fmt.Println("🚀 开始测试 Tushare Stock Board MCP API...")

	// 测试所有 stock_board 工具
	results = append(results, testDcDaily(ctx, client))
	results = append(results, testDcHot(ctx, client))
	results = append(results, testDcIndex(ctx, client))
	results = append(results, testDcMember(ctx, client))
	results = append(results, testHmDetail(ctx, client))
	results = append(results, testHmList(ctx, client))
	results = append(results, testKplConceptCons(ctx, client))
	results = append(results, testKplList(ctx, client))
	results = append(results, testLimitCptList(ctx, client))
	results = append(results, testLimitListD(ctx, client))
	results = append(results, testLimitListThs(ctx, client))
	results = append(results, testLimitStep(ctx, client))
	results = append(results, testStkAuction(ctx, client))
	results = append(results, testTdxDaily(ctx, client))
	results = append(results, testTdxIndex(ctx, client))
	results = append(results, testTdxMember(ctx, client))
	results = append(results, testThsDaily(ctx, client))
	results = append(results, testThsHot(ctx, client))
	results = append(results, testThsIndex(ctx, client))
	results = append(results, testThsMember(ctx, client))
	results = append(results, testTopInst(ctx, client))
	results = append(results, testTopList(ctx, client))

	// 生成报告
	generateBoardReport(results)
}

func testDcDaily(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 dc_daily...")
	req := &stock_boardapi.DcDailyRequest{
		TradeDate: "20240308",
		IdxType:   "概念板块",
	}

	data, err := stock_boardapi.DcDaily(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "dc_daily",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "dc_daily",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testDcHot(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 dc_hot...")
	req := &stock_boardapi.DcHotRequest{
		TradeDate: "20240308",
		HotType:   "人气榜",
		IsNew:     "Y",
	}

	data, err := stock_boardapi.DcHot(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "dc_hot",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "dc_hot",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testDcIndex(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 dc_index...")
	req := &stock_boardapi.DcIndexRequest{
		IdxType: "概念板块",
	}

	data, err := stock_boardapi.DcIndex(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "dc_index",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "dc_index",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testDcMember(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 dc_member...")
	req := &stock_boardapi.DcMemberRequest{
		TsCode:    "801005.DC",
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.DcMember(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "dc_member",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "dc_member",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testHmDetail(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 hm_detail...")
	req := &stock_boardapi.HmDetailRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.HmDetail(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "hm_detail",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "hm_detail",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testHmList(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 hm_list...")
	req := &stock_boardapi.HmListRequest{}

	data, err := stock_boardapi.HmList(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "hm_list",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "hm_list",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testKplConceptCons(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 kpl_concept_cons...")
	req := &stock_boardapi.KplConceptConsRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.KplConceptCons(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "kpl_concept_cons",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "kpl_concept_cons",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testKplList(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 kpl_list...")
	req := &stock_boardapi.KplListRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.KplList(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "kpl_list",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "kpl_list",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testLimitCptList(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 limit_cpt_list...")
	req := &stock_boardapi.LimitCptListRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.LimitCptList(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "limit_cpt_list",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "limit_cpt_list",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testLimitListD(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 limit_list_d...")
	req := &stock_boardapi.LimitListDRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.LimitListD(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "limit_list_d",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "limit_list_d",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testLimitListThs(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 limit_list_ths...")
	req := &stock_boardapi.LimitListThsRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.LimitListThs(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "limit_list_ths",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "limit_list_ths",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testLimitStep(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 limit_step...")
	req := &stock_boardapi.LimitStepRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.LimitStep(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "limit_step",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "limit_step",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testStkAuction(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 stk_auction...")
	req := &stock_boardapi.StkAuctionRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.StkAuction(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "stk_auction",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "stk_auction",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testTdxDaily(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 tdx_daily...")
	req := &stock_boardapi.TdxDailyRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.TdxDaily(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "tdx_daily",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "tdx_daily",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testTdxIndex(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 tdx_index...")
	req := &stock_boardapi.TdxIndexRequest{
		IdxType: "概念板块",
	}

	data, err := stock_boardapi.TdxIndex(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "tdx_index",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "tdx_index",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testTdxMember(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 tdx_member...")
	req := &stock_boardapi.TdxMemberRequest{
		TsCode:    "1.TDX",
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.TdxMember(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "tdx_member",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "tdx_member",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testThsDaily(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 ths_daily...")
	req := &stock_boardapi.ThsDailyRequest{
		TsCode:    "851001.THS",
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.ThsDaily(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "ths_daily",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "ths_daily",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testThsHot(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 ths_hot...")
	req := &stock_boardapi.ThsHotRequest{
		TradeDate: "20240308",
		Market:    "热股",
		IsNew:     "Y",
	}

	data, err := stock_boardapi.ThsHot(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "ths_hot",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "ths_hot",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testThsIndex(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 ths_index...")
	req := &stock_boardapi.ThsIndexRequest{
		Exchange: "A",
		Type:     "N",
	}

	data, err := stock_boardapi.ThsIndex(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "ths_index",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "ths_index",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testThsMember(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 ths_member...")
	req := &stock_boardapi.ThsMemberRequest{
		TsCode: "851001.THS",
	}

	data, err := stock_boardapi.ThsMember(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "ths_member",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "ths_member",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testTopInst(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 top_inst...")
	req := &stock_boardapi.TopInstRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.TopInst(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "top_inst",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "top_inst",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func testTopList(ctx context.Context, client *sdk.Client) BoardTestResult {
	fmt.Println("测试 top_list...")
	req := &stock_boardapi.TopListRequest{
		TradeDate: "20240308",
	}

	data, err := stock_boardapi.TopList(ctx, client, req)
	if err != nil {
		return BoardTestResult{
			APIName: "top_list",
			Success: false,
			Error:   err.Error(),
		}
	}

	return BoardTestResult{
		APIName: "top_list",
		Success: true,
		HasData: len(data) > 0,
		Count:   len(data),
	}
}

func generateBoardReport(results []BoardTestResult) {
	fmt.Println("\n============================================")
	fmt.Println("     TUSHARE STOCK BOARD 测试报告")
	fmt.Println("============================================")

	successCount := 0
	hasDataCount := 0
	failCount := 0

	fmt.Printf("\n%-25s %-10s %-10s %-10s\n", "API名称", "状态", "数据量", "结果")
	fmt.Println("--------------------------------------------")

	for _, r := range results {
		var status, resultStr, dataStr string

		switch {
		case !r.Success:
			status = "❌"
			resultStr = "失败"
			dataStr = "-"
			failCount++
		case !r.HasData:
			status = "⚠️"
			resultStr = "无数据"
			dataStr = fmt.Sprintf("%d", r.Count)
			successCount++
		default:
			status = "✅"
			resultStr = "成功"
			dataStr = fmt.Sprintf("%d", r.Count)
			successCount++
			hasDataCount++
		}

		fmt.Printf("%-25s %-10s %-10s %-10s\n", r.APIName, status, dataStr, resultStr)
	}

	fmt.Println("\n============================================")
	fmt.Println("测试统计")
	fmt.Println("============================================")

	totalTests := len(results)
	fmt.Printf("总测试数: %d\n", totalTests)
	fmt.Printf("成功: %d (%.1f%%)\n", successCount, float64(successCount)*100/float64(totalTests))
	fmt.Printf("有数据: %d (%.1f%%)\n", hasDataCount, float64(hasDataCount)*100/float64(totalTests))
	fmt.Printf("无数据: %d (%.1f%%)\n", successCount-hasDataCount, float64(successCount-hasDataCount)*100/float64(totalTests))
	fmt.Printf("失败: %d (%.1f%%)\n", failCount, float64(failCount)*100/float64(totalTests))

	// 显示失败的API
	if failCount > 0 || (successCount-hasDataCount) > 0 {
		fmt.Println("\n============================================")
		fmt.Println("需要关注的API")
		fmt.Println("============================================")

		for _, r := range results {
			if !r.Success {
				fmt.Printf("❌ %s: %s\n", r.APIName, r.Error)
			} else if !r.HasData {
				fmt.Printf("⚠️  %s: 调用成功但无数据返回\n", r.APIName)
			}
		}
	}

	// 保存详细报告
	saveBoardReport(results)

	fmt.Println("\n============================================")
	fmt.Println("✅ 详细报告已保存到: tests/stock-board-test-report.md")
	fmt.Println("============================================")
}

func saveBoardReport(results []BoardTestResult) {
	content := `# Tushare Stock Board MCP API 测试报告

**测试时间**: 2026-03-09
**测试目的**: 验证所有22个Stock Board MCP工具的接口可用性和数据返回情况

---

## 测试结果汇总

| 状态 | 数量 | 百分比 |
|------|------|--------|
`

	successCount := 0
	hasDataCount := 0
	failCount := 0

	for _, r := range results {
		if r.Success {
			successCount++
			if r.HasData {
				hasDataCount++
			}
		} else {
			failCount++
		}
	}

	totalTests := len(results)

	content += fmt.Sprintf("| ✅ 成功 | %d | %.1f%% |\n", successCount, float64(successCount)*100/float64(totalTests))
	content += fmt.Sprintf("| 📈 有数据 | %d | %.1f%% |\n", hasDataCount, float64(hasDataCount)*100/float64(totalTests))
	content += fmt.Sprintf("| ⚠️ 无数据 | %d | %.1f%% |\n", successCount-hasDataCount, float64(successCount-hasDataCount)*100/float64(totalTests))
	content += fmt.Sprintf("| ❌ 失败 | %d | %.1f%% |\n", failCount, float64(failCount)*100/float64(totalTests))

	content += `

---

## 详细测试结果

`

	for _, r := range results {
		status := "✅ 成功"
		if !r.Success {
			status = "❌ 失败"
		} else if !r.HasData {
			status = "⚠️ 无数据"
		}

		content += fmt.Sprintf("### %s\n\n", r.APIName)
		content += fmt.Sprintf("- **状态**: %s\n", status)
		if r.Success {
			content += fmt.Sprintf("- **返回数据量**: %d 条\n", r.Count)
			content += fmt.Sprintf("- **数据状态**: %s\n", map[bool]string{true: "有数据 ✅", false: "无数据 ⚠️"}[r.HasData])
		} else {
			content += fmt.Sprintf("- **错误信息**: `%s`\n", r.Error)
		}
		content += "\n"
	}

	content += `---

## 测试说明

- ✅ **成功**: API调用成功，服务器正常响应
- ⚠️ **无数据**: API调用成功，但未返回数据（可能测试参数无匹配数据）
- ❌ **失败**: API调用失败，出现错误

## 测试参数

- **交易日期**: 2024-03-08 (大部分接口)
- **板块类型**: 概念板块
- **热点类型**: 人气榜
- **市场类型**: A股市场/热股
- **特殊代码**:
  - dc_member: 801005.DC
  - tdx_member: 1.TDX
  - ths_daily/member: 851001.THS

---

**报告生成时间**: 2026-03-09
**测试工具**: tushare-go SDK
**测试工具数量**: 22个
`

	err := os.WriteFile("tests/stock-board-test-report.md", []byte(content), 0644)
	if err != nil {
		log.Printf("Error writing report: %v", err)
	}
}