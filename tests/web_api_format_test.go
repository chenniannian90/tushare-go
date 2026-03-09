package tests

import (
	"encoding/json"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestWebAPIResponseFormat tests the exact response format from Tushare Web API
func TestWebAPIResponseFormat(t *testing.T) {
	// Compact JSON to avoid whitespace issues
	webAPIResponse := `{"code":0,"message":"","data":{"fields":["ts_code","symbol","name","area","industry","market","list_date"],"items":[["000001.SZ","000001","PING_AN_BANK","SHENZHEN","BANK","MAIN_BOARD","19910403"],["000002.SZ","000002","VANKE_A","SHENZHEN","REAL_ESTATE","MAIN_BOARD","19910129"],["000004.SZ","000004","ST_STOCK","SHENZHEN","SOFTWARE","MAIN_BOARD","19901201"]],"has_more":true,"count":-1},"request_id":"test-request-id"}`

	// Parse outer response
	var outerResp struct {
		Code int             `json:"code"`
		Msg  string          `json:"message"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(webAPIResponse), &outerResp); err != nil {
		t.Fatalf("Failed to parse outer response: %v", err)
	}

	// Parse data part using APIResponse
	var apiResp sdk.APIResponse
	if err := json.Unmarshal(outerResp.Data, &apiResp); err != nil {
		t.Fatalf("Failed to parse API response: %v", err)
	}

	// Log raw items for debugging
	t.Logf("Raw items: %s", string(apiResp.Items))

	// Verify format detection
	format := apiResp.DetectFormat()
	t.Logf("Detected format: %d (ArrayArray=%d, ObjectArray=%d, Unknown=%d)", 
		format, sdk.FormatArrayArray, sdk.FormatObjectArray, sdk.FormatUnknown)
	
	if format != sdk.FormatArrayArray {
		t.Errorf("Expected FormatArrayArray (%d), got %d", sdk.FormatArrayArray, format)
	}

	// Parse and convert to object array
	items, err := apiResp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse and convert: %v", err)
	}

	// Verify results
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Verify first item in detail
	firstItem := items[0]
	t.Logf("First item: %+v", firstItem)

	expectedFields := []string{"ts_code", "symbol", "name", "area", "industry", "market", "list_date"}
	for _, field := range expectedFields {
		if _, exists := firstItem[field]; !exists {
			t.Errorf("Missing field '%s' in first item", field)
		}
	}

	// Verify specific values
	if firstItem["ts_code"] != "000001.SZ" {
		t.Errorf("Expected ts_code '000001.SZ', got '%v'", firstItem["ts_code"])
	}
	if firstItem["name"] != "PING_AN_BANK" {
		t.Errorf("Expected name 'PING_AN_BANK', got '%v'", firstItem["name"])
	}
	if firstItem["market"] != "MAIN_BOARD" {
		t.Errorf("Expected market 'MAIN_BOARD', got '%v'", firstItem["market"])
	}
	if firstItem["list_date"] != "19910403" {
		t.Errorf("Expected list_date '19910403', got '%v'", firstItem["list_date"])
	}

	// Verify third item (ST stock)
	thirdItem := items[2]
	if thirdItem["ts_code"] != "000004.SZ" {
		t.Errorf("Expected ts_code '000004.SZ', got '%v'", thirdItem["ts_code"])
	}
	if thirdItem["name"] != "ST_STOCK" {
		t.Errorf("Expected name 'ST_STOCK', got '%v'", thirdItem["name"])
	}
}

// TestOfficialAPIResponseFormat tests the official API response format
func TestOfficialAPIResponseFormat(t *testing.T) {
	// Compact JSON
	officialAPIResponse := `{"code":0,"msg":null,"data":{"fields":["ts_code","trade_date","open","high","low","close","vol","amount"],"items":[{"ts_code":"000001.SZ","trade_date":"20230101","open":10.5,"high":11.0,"low":10.2,"close":10.8,"vol":100000,"amount":1080000},{"ts_code":"000002.SZ","trade_date":"20230101","open":8.5,"high":8.8,"low":8.3,"close":8.7,"vol":80000,"amount":696000}]}}`

	// Parse outer response
	var outerResp struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal([]byte(officialAPIResponse), &outerResp); err != nil {
		t.Fatalf("Failed to parse outer response: %v", err)
	}

	// Parse data part using APIResponse
	var apiResp sdk.APIResponse
	if err := json.Unmarshal(outerResp.Data, &apiResp); err != nil {
		t.Fatalf("Failed to parse API response: %v", err)
	}

	// Log raw items for debugging
	t.Logf("Raw items: %s", string(apiResp.Items))

	// Verify format detection
	format := apiResp.DetectFormat()
	t.Logf("Detected format: %d (ArrayArray=%d, ObjectArray=%d, Unknown=%d)", 
		format, sdk.FormatArrayArray, sdk.FormatObjectArray, sdk.FormatUnknown)
	
	if format != sdk.FormatObjectArray {
		t.Errorf("Expected FormatObjectArray (%d), got %d", sdk.FormatObjectArray, format)
	}

	// Parse
	items, err := apiResp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Verify results
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify first item
	firstItem := items[0]
	if firstItem["ts_code"] != "000001.SZ" {
		t.Errorf("Expected ts_code '000001.SZ', got '%v'", firstItem["ts_code"])
	}
	if firstItem["open"].(float64) != 10.5 {
		t.Errorf("Expected open 10.5, got %v", firstItem["open"])
	}
}
