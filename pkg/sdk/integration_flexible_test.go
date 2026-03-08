package sdk

import (
	"encoding/json"
	"testing"
)

// TestResponseParser_RealWorldCase 测试真实世界的new_share响应
func TestResponseParser_RealWorldCase(t *testing.T) {
	// 这是从调试日志中提取的真实响应
	realJSON := `{
		"fields": ["ts_code", "sub_code", "name", "ipo_date", "issue_date", "amount", "market_amount", "price", "pe", "limit_amount", "funds", "ballot"],
		"items": [
			["688781.SH", "787781", "视涯科技", "20260316", null, 10000.0, 0.0, 0.0, 0.0, 1.4, 0.0, 0.0],
			["301682.SZ", "301682", "宏明电子", "20260316", null, 3039.0, 0.0, 0.0, 0.0, 0.85, 0.0, 0.0]
		]
	}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(realJSON), &data); err != nil {
		t.Fatalf("解析JSON失败: %v", err)
	}

	// 提取items和fields
	fieldsJSON, _ := json.Marshal(data["fields"])
	var fields []string
	json.Unmarshal(fieldsJSON, &fields)

	itemsJSON, _ := json.Marshal(data["items"])

	// 使用我们的响应解析器
	resp := APIResponse{
		Fields: fields,
		Items:  itemsJSON,
	}

	// 测试格式检测
	format := resp.DetectFormat()
	if format != FormatArrayArray {
		t.Errorf("格式检测错误，期望 %v，实际 %v", FormatArrayArray, format)
	}

	// 测试解析和转换
	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("ParseAndConvert() error = %v", err)
	}

	// 验证结果
	if len(items) != 2 {
		t.Fatalf("期望2条记录，实际得到%d条", len(items))
	}

	// 验证字段完整性
	expectedFields := []string{"ts_code", "sub_code", "name", "ipo_date", "issue_date", "amount", "market_amount", "price", "pe", "limit_amount", "funds", "ballot"}
	for _, field := range expectedFields {
		if _, ok := items[0][field]; !ok {
			t.Errorf("转换后的数据缺少字段: %s", field)
		}
	}

	// 验证具体值
	if items[0]["ts_code"] != "688781.SH" {
		t.Errorf("ts_code = %v, want 688781.SH", items[0]["ts_code"])
	}
	if items[0]["name"] != "视涯科技" {
		t.Errorf("name = %v, want 视涯科技", items[0]["name"])
	}
	if items[0]["amount"] != 10000.0 {
		t.Errorf("amount = %v, want 10000.0", items[0]["amount"])
	}
	if items[1]["ts_code"] != "301682.SZ" {
		t.Errorf("第二条记录ts_code = %v, want 301682.SZ", items[1]["ts_code"])
	}

	t.Logf("✅ 真实new_share响应测试通过")
	t.Logf("   - 格式检测: %v", format)
	t.Logf("   - 记录数量: %d", len(items))
	t.Logf("   - 第一条: ts_code=%s, name=%s, amount=%.0f", items[0]["ts_code"], items[0]["name"], items[0]["amount"])
	t.Logf("   - 第二条: ts_code=%s, name=%s, amount=%.0f", items[1]["ts_code"], items[1]["name"], items[1]["amount"])
}
