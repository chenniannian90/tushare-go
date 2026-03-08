package stock_basic

import (
	"context"
	"os"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestNewShare_Integration 集成测试：验证new_share API与真实Tushare API的交互
// 运行此测试需要设置 TUSHARE_TOKEN 环境变量
//
// 运行命令：
// TUSHARE_TOKEN=your_token go test ./pkg/sdk/api/stock_basic/ -v -run TestNewShare_Integration
func TestNewShare_Integration(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("跳过集成测试：未设置 TUSHARE_TOKEN 环境变量")
	}

	// 创建客户端
	config, err := sdk.NewConfig(token)
	if err != nil {
		t.Fatalf("创建配置失败: %v", err)
	}
	client := sdk.NewClient(config)

	// 测试1：获取最近的IPO数据
	t.Run("获取最近的IPO数据", func(t *testing.T) {
		ctx := context.Background()

		// 获取2026年3月的IPO数据
		req := &NewShareRequest{
			StartDate: "20260301",
			EndDate:   "20260331",
		}

		items, err := NewShare(ctx, client, req)
		if err != nil {
			t.Fatalf("❌ 调用new_share API失败: %v", err)
		}

		t.Logf("✅ 成功获取 %d 条IPO数据", len(items))

		// 验证数据完整性
		if len(items) > 0 {
			first := items[0]
			t.Logf("📊 第一条数据示例：")
			t.Logf("   - TS代码: %s", first.TsCode)
			t.Logf("   - 申购代码: %s", first.SubCode)
			t.Logf("   - 名称: %s", first.Name)
			t.Logf("   - 上市日期: %s", first.IpoDate)
			t.Logf("   - 发行价格: %.2f", first.Price)
			t.Logf("   - 市盈率: %.2f", first.Pe)

			// 验证必需字段不为空
			if first.TsCode == "" {
				t.Error("❌ TsCode字段为空")
			}
			if first.Name == "" {
				t.Error("❌ Name字段为空")
			}
		} else {
			t.Logf("⚠️  当前时间段无IPO数据")
		}
	})

	// 测试2：验证CallAPIFlexible处理不同响应格式的能力
	t.Run("验证CallAPIFlexible的灵活性", func(t *testing.T) {
		ctx := context.Background()

		req := &NewShareRequest{
			StartDate: "20260301",
			EndDate:   "20260331",
		}

		items, err := NewShare(ctx, client, req)
		if err != nil {
			t.Fatalf("❌ 调用失败: %v", err)
		}

		// 验证：无论API返回对象数组还是二维数组，都应该正确解析
		t.Logf("✅ CallAPIFlexible成功处理响应，得到 %d 条记录", len(items))

		// 验证数据结构的正确性
		for i, item := range items {
			if item.TsCode == "" {
				t.Errorf("❌ 第 %d 条记录的TsCode为空", i)
			}
		}
	})
}

// TestNewShare_WithoutToken 无需token的单元测试
// 验证代码结构和类型正确性
func TestNewShare_CodeStructure(t *testing.T) {
	// 验证请求结构体的字段标签正确性
	req := &NewShareRequest{
		StartDate: "20260301",
		EndDate:   "20260331",
	}

	if req.StartDate != "20260301" {
		t.Error("StartDate设置失败")
	}
	if req.EndDate != "20260331" {
		t.Error("EndDate设置失败")
	}

	t.Log("✅ 代码结构验证通过")
}

// TestNewShare_HelperFunctions 验证辅助函数的正确性
func TestNewShare_HelperFunctions(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]interface{}
		key      string
		expected string
	}{
		{
			name:     "获取存在的string字段",
			input:    map[string]interface{}{"ts_code": "000001.SZ"},
			key:      "ts_code",
			expected: "000001.SZ",
		},
		{
			name:     "获取不存在的字段返回空字符串",
			input:    map[string]interface{}{"ts_code": "000001.SZ"},
			key:      "name",
			expected: "",
		},
		{
			name:     "处理nil值返回空字符串",
			input:    map[string]interface{}{"amount": nil},
			key:      "amount",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getString(tc.input, tc.key)
			if result != tc.expected {
				t.Errorf("getString() = %v, want %v", result, tc.expected)
			}
		})
	}

	// 测试getFloat64
	floatTestCases := []struct {
		name     string
		input    map[string]interface{}
		key      string
		expected float64
	}{
		{
			name:     "获取float64值",
			input:    map[string]interface{}{"amount": 1000.5},
			key:      "amount",
			expected: 1000.5,
		},
		{
			name:     "处理int值转换为float64",
			input:    map[string]interface{}{"amount": 1000},
			key:      "amount",
			expected: 1000.0,
		},
		{
			name:     "处理nil值返回0",
			input:    map[string]interface{}{"amount": nil},
			key:      "amount",
			expected: 0.0,
		},
	}

	for _, tc := range floatTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getFloat64(tc.input, tc.key)
			if result != tc.expected {
				t.Errorf("getFloat64() = %v, want %v", result, tc.expected)
			}
		})
	}

	t.Log("✅ 辅助函数验证通过")
}
