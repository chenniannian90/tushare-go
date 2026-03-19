package market

import (
	"testing"

	"github.com/chenniannian90/tushare-go/realtime"
)

func TestRealTimeQuoteSina(t *testing.T) {
	// 创建测试客户端
	client := &Client{}

	// 测试真实的新浪 API（需要网络连接）
	t.Run("RealAPI", func(t *testing.T) {
		if testing.Short() {
			t.Skip("跳过真实 API 测试（使用 -short 标志）")
		}

		resp, err := client.RealTimeQuoteSina("000001.SZ", "600000.SH")
		if err != nil {
			t.Fatalf("RealTimeQuoteSina() error = %v", err)
		}

		if resp == nil {
			t.Fatal("RealTimeQuoteSina() returned nil response")
		}

		if len(resp.Quotes) == 0 {
			t.Error("RealTimeQuoteSina() returned no quotes")
		}

		// 验证返回的数据
		for code, quote := range resp.Quotes {
			if quote.Name == "" {
				t.Errorf("股票 %s 的名称为空", code)
			}
			if quote.Price <= 0 {
				t.Errorf("股票 %s 的价格无效: %f", code, quote.Price)
			}
		}
	})
}

func TestSinaStockQuote_GetChange(t *testing.T) {
	quote := &realtime.SinaStockQuote{
		Price:    110.0,
		PreClose: 100.0,
	}

	expected := 10.0
	result := quote.GetChange()

	if result != expected {
		t.Errorf("GetChange() = %f; want %f", result, expected)
	}
}

func TestSinaStockQuote_GetChangePercent(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		preClose float64
		expected float64
	}{
		{"上涨", 110.0, 100.0, 10.0},
		{"下跌", 90.0, 100.0, -10.0},
		{"持平", 100.0, 100.0, 0.0},
		{"零基准", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote := &realtime.SinaStockQuote{
				Price:    tt.price,
				PreClose: tt.preClose,
			}

			result := quote.GetChangePercent()

			if result != tt.expected {
				t.Errorf("GetChangePercent() = %f; want %f", result, tt.expected)
			}
		})
	}
}

func TestSinaStockQuote_ToMap(t *testing.T) {
	quote := &realtime.SinaStockQuote{
		Name:     "测试股票",
		Price:    100.0,
		PreClose: 90.0,
		Volume:   1000000,
		Amount:   100000000,
	}

	result := quote.ToMap()

	if result["name"] != "测试股票" {
		t.Errorf("ToMap()[name] = %v; want 测试股票", result["name"])
	}

	if result["price"] != 100.0 {
		t.Errorf("ToMap()[price] = %v; want 100.0", result["price"])
	}

	if result["volume"] != int64(1000000) {
		t.Errorf("ToMap()[volume] = %v; want 1000000", result["volume"])
	}
}

func TestSinaQuoteResponse_GetTopGainers(t *testing.T) {
	quotes := map[string]*realtime.SinaStockQuote{
		"sh600000": {
			Name:     "股票A",
			Price:    110.0,
			PreClose: 100.0,
		}, // +10%
		"sz000001": {
			Name:     "股票B",
			Price:    90.0,
			PreClose: 100.0,
		}, // -10%
		"sh600519": {
			Name:     "股票C",
			Price:    105.0,
			PreClose: 100.0,
		}, // +5%
		"sz000002": {
			Name:     "股票D",
			Price:    95.0,
			PreClose: 100.0,
		}, // -5%
	}

	resp := &realtime.SinaQuoteResponse{
		Quotes: quotes,
	}

	// 获取前 2 名涨幅榜
	topGainers := resp.GetTopGainers(2)
	if len(topGainers) != 2 {
		t.Errorf("GetTopGainers(2) returned %d items; want 2", len(topGainers))
	}

	// 验证顺序
	if topGainers[0] != "sh600000" {
		t.Errorf("GetTopGainers(2)[0] = %s; want sh600000", topGainers[0])
	}
	if topGainers[1] != "sh600519" {
		t.Errorf("GetTopGainers(2)[1] = %s; want sh600519", topGainers[1])
	}
}

func TestSinaQuoteResponse_GetTopLosers(t *testing.T) {
	quotes := map[string]*realtime.SinaStockQuote{
		"sh600000": {
			Name:     "股票A",
			Price:    110.0,
			PreClose: 100.0,
		}, // +10%
		"sz000001": {
			Name:     "股票B",
			Price:    90.0,
			PreClose: 100.0,
		}, // -10%
		"sh600519": {
			Name:     "股票C",
			Price:    105.0,
			PreClose: 100.0,
		}, // +5%
		"sz000002": {
			Name:     "股票D",
			Price:    95.0,
			PreClose: 100.0,
		}, // -5%
	}

	resp := &realtime.SinaQuoteResponse{
		Quotes: quotes,
	}

	// 获取前 2 名跌幅榜
	topLosers := resp.GetTopLosers(2)
	if len(topLosers) != 2 {
		t.Errorf("GetTopLosers(2) returned %d items; want 2", len(topLosers))
	}

	// 验证顺序
	if topLosers[0] != "sz000001" {
		t.Errorf("GetTopLosers(2)[0] = %s; want sz000001", topLosers[0])
	}
	if topLosers[1] != "sz000002" {
		t.Errorf("GetTopLosers(2)[1] = %s; want sz000002", topLosers[1])
	}
}

func TestGetBatchRealtimeQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过真实 API 测试（使用 -short 标志）")
	}

	client := &Client{}

	// 创建测试代码列表（使用真实的股票代码）
	codes := []string{
		"000001.SZ",
		"000002.SZ",
		"600000.SH",
		"600519.SH",
	}

	// 测试批量查询
	resp, err := client.GetBatchRealtimeQuote(codes, 2)
	if err != nil {
		t.Fatalf("GetBatchRealtimeQuote() error = %v", err)
	}

	if resp == nil {
		t.Fatal("GetBatchRealtimeQuote() returned nil response")
	}

	if len(resp.Quotes) == 0 {
		t.Error("GetBatchRealtimeQuote() returned no quotes")
	}
}
