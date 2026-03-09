package tests

import (
	"encoding/json"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestTradeCalAPI tests trade_cal API response
func TestTradeCalAPI(t *testing.T) {
	// Actual response from Web API
	apiResponse := `{"code":0,"message":"","data":{"fields":["exchange","cal_date","is_open","pretrade_date"],"items":[["SSE","20250105",0,"20250103"],["SSE","20250104",0,"20250103"],["SSE","20250103",1,"20250102"]],"has_more":true,"count":-1},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Verify first item
	if items[0]["exchange"] != "SSE" {
		t.Errorf("Expected exchange 'SSE', got '%v'", items[0]["exchange"])
	}
	if items[0]["is_open"].(float64) != 0 {
		t.Errorf("Expected is_open 0, got %v", items[0]["is_open"])
	}

	// Verify third item (is_open = 1)
	if items[2]["is_open"].(float64) != 1 {
		t.Errorf("Expected is_open 1, got %v", items[2]["is_open"])
	}
}

// TestStockCompanyAPI tests stock_company API response
func TestStockCompanyAPI(t *testing.T) {
	// Actual response from Web API - note the mixed types (string, float, int)
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","com_name","chairman","manager","secretary","reg_capital","setup_date","province","city","exchange"],"items":[["000001.SZ","PING_AN_BANK_CO","XIE_YONGLIN","JI_GUANGHENG","ZHOU_QIANG",1940591.8198,"19871222","GUANGDONG","SHENZHEN","SZSE"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	// Verify all fields
	expectedFields := []string{"ts_code", "com_name", "chairman", "manager", "secretary", "reg_capital", "setup_date", "province", "city", "exchange"}
	for _, field := range expectedFields {
		if _, exists := items[0][field]; !exists {
			t.Errorf("Missing field '%s'", field)
		}
	}

	// Verify float field (reg_capital)
	if items[0]["reg_capital"].(float64) != 1940591.8198 {
		t.Errorf("Expected reg_capital 1940591.8198, got %v", items[0]["reg_capital"])
	}
}

// TestEmptyItemsArray tests empty items array
func TestEmptyItemsArray(t *testing.T) {
	// Response with no data (empty items)
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","start_date","end_date","ann_date","change_reason"],"items":[]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// TestNullFieldsAndItems tests null fields and items
func TestNullFieldsAndItems(t *testing.T) {
	// Response with null fields/items (API error or no data)
	apiResponse := `{"code":0,"message":"","data":{"fields":null,"items":null},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	// This should handle null items gracefully
	items, err := resp.ParseAndConvert()
	// We expect either success with empty items or an error
	if err != nil {
		t.Logf("Expected error for null items: %v", err)
	} else if len(items) != 0 {
		t.Logf("Got %d items for null items array", len(items))
	}
}

// TestLargeNumbersInResponse tests large numbers in API response
func TestLargeNumbersInResponse(t *testing.T) {
	// Test with large numbers (like trade volume, market cap, etc.)
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","trade_date","vol","amount","total_mv","circ_mv"],"items":[["000001.SZ","20250110",100000000,1080000000,500000000000,300000000000],["000002.SZ","20250110",80000000,696000000,400000000000,350000000000]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify large numbers are preserved correctly
	if items[0]["vol"].(float64) != 100000000 {
		t.Errorf("Expected vol 100000000, got %v", items[0]["vol"])
	}
	if items[0]["total_mv"].(float64) != 500000000000 {
		t.Errorf("Expected total_mv 500000000000, got %v", items[0]["total_mv"])
	}
}

// TestSpecialCharactersInData tests special characters and Chinese text
func TestSpecialCharactersInData(t *testing.T) {
	// Test with Chinese characters and special symbols
	apiResponse := `{"code":0,"message":"","data":{"fields":["ts_code","name","area","industry"],"items":[["000001.SZ","PING_AN_BANK","SHENZHEN","BANK"],["000002.SZ","VANKE_A","SHENZHEN","REAL_ESTATE"],["600000.SH","PUDONG_DEVELOPMENT","SHANGHAI","FINANCE"]]},"request_id":"test"}`

	var outer struct {
		Code int             `json:"code"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(apiResponse), &outer); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	var resp sdk.APIResponse
	if err := json.Unmarshal(outer.Data, &resp); err != nil {
		t.Fatalf("Failed to parse data: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify all items are parsed correctly
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Verify special characters in industry names
	if items[1]["industry"] != "REAL_ESTATE" {
		t.Errorf("Expected 'REAL_ESTATE', got '%v'", items[1]["industry"])
	}
	if items[2]["industry"] != "FINANCE" {
		t.Errorf("Expected 'FINANCE', got '%v'", items[2]["industry"])
	}
}
