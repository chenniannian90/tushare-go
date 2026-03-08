package sdk

import (
	"encoding/json"
	"testing"
)

func TestAPIResponse_DetectFormat(t *testing.T) {
	tests := []struct {
		name     string
		items    string
		expected ResponseFormat
	}{
		{
			name:     "对象数组格式",
			items:    `[{"ts_code":"000001.SZ","name":"平安银行"},{"ts_code":"000002.SZ","name":"万��A"}]`,
			expected: FormatObjectArray,
		},
		{
			name:     "二维数组格式",
			items:    `[["000001.SZ","平安银行"],["000002.SZ","万科A"]]`,
			expected: FormatArrayArray,
		},
		{
			name:     "空数组",
			items:    `[]`,
			expected: FormatArrayArray,
		},
		{
			name:     "null",
			items:    `null`,
			expected: FormatArrayArray,
		},
		{
			name:     "嵌套对象数组",
			items:    `[{"ts_code":"000001.SZ"},{"ts_code":"000002.SZ"}]`,
			expected: FormatObjectArray,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := APIResponse{
				Items: json.RawMessage(tt.items),
			}
			format := resp.DetectFormat()
			if format != tt.expected {
				t.Errorf("DetectFormat() = %v, want %v", format, tt.expected)
			}
		})
	}
}

func TestAPIResponse_ParseAndConvert(t *testing.T) {
	fields := []string{"ts_code", "name"}

	tests := []struct {
		name    string
		items   string
		wantLen int
		wantErr bool
	}{
		{
			name:    "对象数组格式",
			items:   `[{"ts_code":"000001.SZ","name":"平安银行"},{"ts_code":"000002.SZ","name":"万科A"}]`,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "二维数组格式",
			items:   `[["000001.SZ","平安银行"],["000002.SZ","万科A"]]`,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "空数组",
			items:   `[]`,
			wantLen: 0,
			wantErr: false,
		},
		{
			name:    "无效的JSON",
			items:   `invalid json`,
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := APIResponse{
				Fields: fields,
				Items:  json.RawMessage(tt.items),
			}

			items, err := resp.ParseAndConvert()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAndConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(items) != tt.wantLen {
				t.Errorf("ParseAndConvert() len = %v, want %v", len(items), tt.wantLen)
			}

			// 验证转换后的数据结构正确
			if !tt.wantErr && len(items) > 0 {
				if _, ok := items[0]["ts_code"]; !ok {
					t.Errorf("ParseAndConvert() 结果缺少 ts_code 字段")
				}
				if _, ok := items[0]["name"]; !ok {
					t.Errorf("ParseAndConvert() 结果缺少 name 字段")
				}
			}
		})
	}
}

func TestConvertArrayRowToMap(t *testing.T) {
	fields := []string{"ts_code", "name", "amount"}

	tests := []struct {
		name    string
		row     []interface{}
		wantErr bool
	}{
		{
			name:    "正常数据",
			row:     []interface{}{"000001.SZ", "平安银行", 100.0},
			wantErr: false,
		},
		{
			name:    "字段数量不匹配",
			row:     []interface{}{"000001.SZ", "平安银行"},
			wantErr: true,
		},
		{
			name:    "包含nil值",
			row:     []interface{}{"000001.SZ", nil, 100.0},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertArrayRowToMap(tt.row, fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertArrayRowToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(result) != len(fields) {
					t.Errorf("ConvertArrayRowToMap() len(result) = %v, want %v", len(result), len(fields))
				}
				if result["ts_code"] != tt.row[0] {
					t.Errorf("ConvertArrayRowToMap() ts_code = %v, want %v", result["ts_code"], tt.row[0])
				}
			}
		})
	}
}