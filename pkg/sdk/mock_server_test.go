//go:build integration

package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// MockTushareServer creates a mock Tushare API server for testing
type MockTushareServer struct {
	server *httptest.Server
}

// NewMockTushareServer creates a new mock server with predefined responses
func NewMockTushareServer() *MockTushareServer {
	return &MockTushareServer{}
}

// Start starts the mock server
func (m *MockTushareServer) Start() {
	m.server = httptest.NewServer(http.HandlerFunc(m.handler))
}

// Close stops the mock server
func (m *MockTushareServer) Close() {
	if m.server != nil {
		m.server.Close()
	}
}

// URL returns the mock server URL
func (m *MockTushareServer) URL() string {
	return m.server.URL
}

// handler handles mock API requests
func (m *MockTushareServer) handler(w http.ResponseWriter, r *http.Request) {
	// Verify request method
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse request to determine API name
	var reqBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code": -2001, "msg": "invalid request format"}`))
		return
	}

	apiName, _ := reqBody["api_name"].(string)

	// Return appropriate mock response based on API name
	switch apiName {
	case "stock_basic":
		m.handleStockBasic(w, reqBody)
	case "trade_cal":
		m.handleTradeCal(w, reqBody)
	case "daily":
		m.handleDaily(w, reqBody)
	case "index_basic":
		m.handleIndexBasic(w, reqBody)
	case "income":
		m.handleIncome(w, reqBody)
	case "balancesheet":
		m.handleBalancesheet(w, reqBody)
	default:
		// Default error response for unknown APIs
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"code": -2001, "msg": "unknown api"}`))
	}
}

// handleStockBasic handles mock stock_basic API requests
func (m *MockTushareServer) handleStockBasic(w http.ResponseWriter, reqBody map[string]interface{}) {
	params, _ := reqBody["params"].(map[string]interface{})

	// Check for specific test conditions
	if tsCode, ok := params["ts_code"].(string); ok && tsCode == "ERROR_40203" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"code": 40203, "msg": "insufficient privileges"}`))
		return
	}

	if tsCode, ok := params["ts_code"].(string); ok && tsCode == "EMPTY_RESULT" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"code": 0, "msg": "success", "data": {"fields": ["ts_code", "name"], "items": []}}`))
		return
	}

	// Default successful response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["ts_code", "symbol", "name", "area", "industry", "market"],
			"items": [
				{"ts_code": "000001.SZ", "symbol": "000001", "name": "平安银行", "area": "深圳", "industry": "银行", "market": "主板"}
			]
		}
	}`))
}

// handleTradeCal handles mock trade_cal API requests
func (m *MockTushareServer) handleTradeCal(w http.ResponseWriter, reqBody map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["exchange", "cal_date", "is_open"],
			"items": [
				{"exchange": "SSE", "cal_date": "20240101", "is_open": "0"},
				{"exchange": "SSE", "cal_date": "20240102", "is_open": "1"}
			]
		}
	}`))
}

// handleDaily handles mock daily API requests
func (m *MockTushareServer) handleDaily(w http.ResponseWriter, reqBody map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["ts_code", "trade_date", "open", "high", "low", "close", "vol"],
			"items": [
				{"ts_code": "000001.SZ", "trade_date": "20240101", "open": 10.5, "high": 11.0, "low": 10.2, "close": 10.8, "vol": 1000000}
			]
		}
	}`))
}

// handleIndexBasic handles mock index_basic API requests
func (m *MockTushareServer) handleIndexBasic(w http.ResponseWriter, reqBody map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["ts_code", "name", "market", "publisher"],
			"items": [
				{"ts_code": "000001.SH", "name": "上证指数", "market": "SSE", "publisher": "中证指数"}
			]
		}
	}`))
}

// handleIncome handles mock income API requests
func (m *MockTushareServer) handleIncome(w http.ResponseWriter, reqBody map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["ts_code", "ann_date", "end_date", "oper_rev", "net_profit"],
			"items": [
				{"ts_code": "000001.SZ", "ann_date": "20240101", "end_date": "20231231", "oper_rev": 1000000000, "net_profit": 300000000}
			]
		}
	}`))
}

// handleBalancesheet handles mock balancesheet API requests
func (m *MockTushareServer) handleBalancesheet(w http.ResponseWriter, reqBody map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"code": 0,
		"msg": "success",
		"data": {
			"fields": ["ts_code", "ann_date", "end_date", "total_assets", "total_liab"],
			"items": [
				{"ts_code": "000001.SZ", "ann_date": "20240101", "end_date": "20231231", "total_assets": 1000000000, "total_liab": 500000000}
			]
		}
	}`))
}

// TestMockServer_BasicFunctionality tests basic mock server functionality
func TestMockServer_BasicFunctionality(t *testing.T) {
	mock := NewMockTushareServer()
	mock.Start()
	defer mock.Close()

	config, _ := NewConfig("test_token")
	config.Endpoint = mock.URL()
	client := NewClient(config)

	ctx := context.Background()

	t.Run("stock_basic_success", func(t *testing.T) {
		req := map[string]interface{}{"ts_code": "000001.SZ"}
		fields := []string{"ts_code", "name"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected items in result")
		}

		t.Logf("✓ Mock server stock_basic call successful")
	})

	t.Run("trade_cal_success", func(t *testing.T) {
		req := map[string]interface{}{
			"exchange":  "SSE",
			"start_date": "20240101",
			"end_date":   "20240105",
		}
		fields := []string{"exchange", "cal_date", "is_open"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "trade_cal", req, fields, &result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected items in result")
		}

		t.Logf("✓ Mock server trade_cal call successful")
	})

	t.Run("api_error_handling", func(t *testing.T) {
		req := map[string]interface{}{"ts_code": "ERROR_40203"}
		fields := []string{"ts_code"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err == nil {
			t.Error("expected error for ERROR_40203")
		} else {
			apiErr, ok := err.(*APIError)
			if !ok {
				t.Errorf("expected APIError, got %T", err)
			} else if apiErr.APICode != 40203 {
				t.Errorf("expected API code 40203, got %d", apiErr.APICode)
			}
		}

		t.Logf("✓ Mock server error handling works correctly")
	})

	t.Run("empty_result_handling", func(t *testing.T) {
		req := map[string]interface{}{"ts_code": "EMPTY_RESULT"}
		fields := []string{"ts_code"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(result.Items) != 0 {
			t.Errorf("expected 0 items, got %d", len(result.Items))
		}

		t.Logf("✓ Mock server empty result handling works correctly")
	})
}

// TestMockServer_PerformanceTests tests mock server performance
func TestMockServer_PerformanceTests(t *testing.T) {
	mock := NewMockTushareServer()
	mock.Start()
	defer mock.Close()

	config, _ := NewConfig("test_token")
	config.Endpoint = mock.URL()
	client := NewClient(config)

	ctx := context.Background()

	t.Run("concurrent_requests", func(t *testing.T) {
		const numRequests = 10
		results := make(chan error, numRequests)

		for i := 0; i < numRequests; i++ {
			go func() {
				req := map[string]interface{}{"ts_code": "000001.SZ"}
				fields := []string{"ts_code", "name"}

				type Result struct {
					Fields []string        `json:"fields"`
					Items  []interface{}   `json:"items"`
				}

				var result Result
				results <- client.CallAPI(ctx, "stock_basic", req, fields, &result)
			}()
		}

		successCount := 0
		for i := 0; i < numRequests; i++ {
			if err := <-results; err != nil {
				t.Errorf("request %d failed: %v", i+1, err)
			} else {
				successCount++
			}
		}

		if successCount != numRequests {
			t.Errorf("expected all %d requests to succeed, got %d", numRequests, successCount)
		}

		t.Logf("✓ Concurrent requests: %d/%d succeeded", successCount, numRequests)
	})

	t.Run("request_timeout", func(t *testing.T) {
		// Test with very short timeout
		shortCtx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		req := map[string]interface{}{"ts_code": "000001.SZ"}
		fields := []string{"ts_code"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(shortCtx, "stock_basic", req, fields, &result)

		// Context should be canceled
		if err != nil {
			t.Logf("✓ Request timeout handled correctly: %v", err)

			// Verify it's a temporary error
			if IsTemporaryError(err) {
				t.Log("✓ Timeout error correctly classified as temporary")
			}
		}
	})
}

// TestMockServer_MultipleAPIs tests multiple different APIs
func TestMockServer_MultipleAPIs(t *testing.T) {
	mock := NewMockTushareServer()
	mock.Start()
	defer mock.Close()

	config, _ := NewConfig("test_token")
	config.Endpoint = mock.URL()
	client := NewClient(config)

	ctx := context.Background()

	apis := []struct {
		name     string
		req      map[string]interface{}
		fields   []string
		minItems int
	}{
		{
			name:     "stock_basic",
			req:      map[string]interface{}{"ts_code": "000001.SZ"},
			fields:   []string{"ts_code", "name"},
			minItems: 1,
		},
		{
			name:     "trade_cal",
			req:      map[string]interface{}{"exchange": "SSE"},
			fields:   []string{"exchange", "cal_date"},
			minItems: 1,
		},
		{
			name:     "daily",
			req:      map[string]interface{}{"ts_code": "000001.SZ"},
			fields:   []string{"ts_code", "trade_date"},
			minItems: 1,
		},
		{
			name:     "index_basic",
			req:      map[string]interface{}{"market": "SSE"},
			fields:   []string{"ts_code", "name"},
			minItems: 1,
		},
		{
			name:     "income",
			req:      map[string]interface{}{"ts_code": "000001.SZ"},
			fields:   []string{"ts_code", "ann_date"},
			minItems: 1,
		},
		{
			name:     "balancesheet",
			req:      map[string]interface{}{"ts_code": "000001.SZ"},
			fields:   []string{"ts_code", "ann_date"},
			minItems: 1,
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			type Result struct {
				Fields []string        `json:"fields"`
				Items  []interface{}   `json:"items"`
			}

			var result Result
			err := client.CallAPI(ctx, api.name, api.req, api.fields, &result)
			if err != nil {
				t.Fatalf("failed to call %s: %v", api.name, err)
			}

			if len(result.Items) < api.minItems {
				t.Errorf("expected at least %d items, got %d", api.minItems, len(result.Items))
			}

			t.Logf("✓ %s API: %d items returned", api.name, len(result.Items))
		})
	}
}

// TestMockServer_ErrorScenarios tests various error scenarios
func TestMockServer_ErrorScenarios(t *testing.T) {
	mock := NewMockTushareServer()
	mock.Start()
	defer mock.Close()

	config, _ := NewConfig("test_token")
	config.Endpoint = mock.URL()
	client := NewClient(config)

	ctx := context.Background()

	t.Run("invalid_api_name", func(t *testing.T) {
		req := map[string]interface{}{}
		fields := []string{"ts_code"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "unknown_api", req, fields, &result)
		if err == nil {
			t.Error("expected error for unknown API")
		} else {
			t.Logf("✓ Unknown API error: %v", err)
		}
	})

	t.Run("network_error_simulation", func(t *testing.T) {
		// Close server to simulate network error
		mock.Close()

		req := map[string]interface{}{"ts_code": "000001.SZ"}
		fields := []string{"ts_code"}

		type Result struct {
			Fields []string        `json:"fields"`
			Items  []interface{}   `json:"items"`
		}

		var result Result
		err := client.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err == nil {
			t.Error("expected network error")
		} else {
			// Check if it's a network error
			if IsNetworkError(err) {
				t.Log("✓ Network error correctly identified")
			}

			if IsTemporaryError(err) {
				t.Log("✓ Network error correctly classified as temporary")
			}
		}

		// Restart server for cleanup
		mock.Start()
	})
}
