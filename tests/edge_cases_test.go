package tests

import (
	"encoding/json"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestNegativeNumbers tests negative numbers in API response
func TestNegativeNumbers(t *testing.T) {
	// Stock prices can go down (negative change)
	apiResponse := `{"fields":["ts_code","change","change_pct"],"items":[["000001.SZ",-1.5,-0.12],["000002.SZ",-2.3,-0.25],["000003.SZ",0.5,0.08]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify negative numbers
	if items[0]["change"].(float64) != -1.5 {
		t.Errorf("Expected -1.5, got %v", items[0]["change"])
	}
	if items[0]["change_pct"].(float64) != -0.12 {
		t.Errorf("Expected -0.12, got %v", items[0]["change_pct"])
	}

	// Verify zero and positive
	if items[2]["change"].(float64) != 0.5 {
		t.Errorf("Expected 0.5, got %v", items[2]["change"])
	}
}

// TestBooleanLikeNumbers tests numbers that represent boolean values
func TestBooleanLikeNumbers(t *testing.T) {
	// Some APIs use 0/1 for boolean values
	apiResponse := `{"fields":["ts_code","is_open","is_hs","list_status"],"items":[["000001.SZ",1,1,"L"],["000002.SZ",0,0,"D"],["000003.SZ",1,0,"P"]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify boolean-like numbers
	if items[0]["is_open"].(float64) != 1 {
		t.Errorf("Expected 1, got %v", items[0]["is_open"])
	}
	if items[1]["is_open"].(float64) != 0 {
		t.Errorf("Expected 0, got %v", items[1]["is_open"])
	}
}

// TestVeryLargeNumbers tests very large numbers
func TestVeryLargeNumbers(t *testing.T) {
	// Market cap, trading volume can be very large
	apiResponse := `{"fields":["ts_code","total_mv","circ_mv","vol"],"items":[["000001.SZ",1.23456789e+12,9.87654321e+11,1234567890],["600000.SH",2.34567890e+12,1.98765432e+12,2345678901]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify large numbers
	totalMV := items[0]["total_mv"].(float64)
	if totalMV < 1.2e+12 || totalMV > 1.3e+12 {
		t.Errorf("Total market cap seems wrong: %v", totalMV)
	}
}

// TestZeroValues tests various zero values
func TestZeroValues(t *testing.T) {
	// Test zero values for different field types
	apiResponse := `{"fields":["ts_code","price","change","vol","amount"],"items":[["000001.SZ",0,0,0,0],["000002.SZ",10.5,0,100000,0]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify zero values
	if items[0]["price"].(float64) != 0 {
		t.Errorf("Expected 0, got %v", items[0]["price"])
	}
	if items[1]["amount"].(float64) != 0 {
		t.Errorf("Expected 0, got %v", items[1]["amount"])
	}
}

// TestSingleItemArray tests array with only one item
func TestSingleItemArray(t *testing.T) {
	apiResponse := `{"fields":["ts_code","name"],"items":[["000001.SZ","STOCK_1"]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
}

// TestManyFields tests response with many fields
func TestManyFields(t *testing.T) {
	// Some APIs return 20+ fields
	apiResponse := `{"fields":["ts_code","trade_date","open","high","low","close","vol","amount","pct_chg","pe","pb","ps","ps_ttm","dv_ratio","dv_ttm","total_share","float_share","free_share","turnover_rate","rate","volume_ratio"],"items":[["000001.SZ","20250110",10.5,11.0,10.2,10.8,100000,1080000,2.85,15.5,1.2,3.5,4.2,0.5,1.8,5000000,3000000,2000000,8.5,12.5,1.5]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	// Verify we have all 21 fields
	if len(items[0]) != 21 {
		t.Errorf("Expected 21 fields, got %d", len(items[0]))
	}
}

// TestScientificNotation tests numbers in scientific notation
func TestScientificNotation(t *testing.T) {
	// Very large or very small numbers may be in scientific notation
	apiResponse := `{"fields":["ts_code","very_large","very_small"],"items":[["000001.SZ",1.23e+10,4.56e-8],["000002.SZ",7.89e+9,1.23e-7]]}`

	var resp sdk.APIResponse
	if err := json.Unmarshal([]byte(apiResponse), &resp); err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	items, err := resp.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to convert: %v", err)
	}

	// Verify scientific notation is preserved
	veryLarge := items[0]["very_large"].(float64)
	if veryLarge < 1.2e+10 || veryLarge > 1.3e+10 {
		t.Errorf("Very large number seems wrong: %v", veryLarge)
	}

	verySmall := items[0]["very_small"].(float64)
	if verySmall < 4.5e-8 || verySmall > 4.7e-8 {
		t.Errorf("Very small number seems wrong: %v", verySmall)
	}
}
