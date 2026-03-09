package tests

import (
	"encoding/json"
	"testing"

	"tushare-go/pkg/sdk"
)

// TestArrayArrayFormat tests parsing array array format from Web API
func TestArrayArrayFormat(t *testing.T) {
	// Simulate Web API response in array array format
	jsonData := `{"fields":["ts_code","symbol","name","area"],"items":[["000001.SZ","000001","STOCK1","AREA1"],["000002.SZ","000002","STOCK2","AREA2"]]}`

	var response sdk.APIResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Verify format detection
	format := response.DetectFormat()
	if format != sdk.FormatArrayArray {
		t.Errorf("Format detection error, expected FormatArrayArray(%d), got %d", 
			sdk.FormatArrayArray, format)
	}

	// Verify parsing and conversion
	items, err := response.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse and convert: %v", err)
	}

	// Verify results
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	// Verify first item
	if items[0]["ts_code"] != "000001.SZ" {
		t.Errorf("First item ts_code error, expected '000001.SZ', got '%v'", items[0]["ts_code"])
	}
}

// TestObjectArrayFormat tests parsing object array format from official API
func TestObjectArrayFormat(t *testing.T) {
	// Simulate official API response in object array format
	jsonData := `{"fields":["ts_code","symbol","name"],"items":[{"ts_code":"000001.SZ","symbol":"000001","name":"STOCK1"},{"ts_code":"000002.SZ","symbol":"000002","name":"STOCK2"}]}`

	var response sdk.APIResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Verify format detection
	format := response.DetectFormat()
	if format != sdk.FormatObjectArray {
		t.Errorf("Format detection error, expected FormatObjectArray(%d), got %d",
			sdk.FormatObjectArray, format)
	}

	// Verify parsing
	items, err := response.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Verify results
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0]["ts_code"] != "000001.SZ" {
		t.Errorf("First item error, expected '000001.SZ', got '%v'", items[0]["ts_code"])
	}
}

// TestEmptyArray tests empty array
func TestEmptyArray(t *testing.T) {
	jsonData := `{"fields":["ts_code","name"],"items":[]}`

	var response sdk.APIResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	items, err := response.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// TestMixedTypes tests mixed types (string, number, null)
func TestMixedTypes(t *testing.T) {
	// Second item: ["000002.SZ", null, -15.3, null]
	// This means: ts_code="000002.SZ", name=null, price=-15.3, change=null
	jsonData := `{"fields":["ts_code","name","price","change"],"items":[["000001.SZ","STOCK1",12.5,1.2],["000002.SZ",null,-15.3,null]]}`

	var response sdk.APIResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	items, err := response.ParseAndConvert()
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}

	// Verify first item types
	if items[0]["price"].(float64) != 12.5 {
		t.Errorf("First item price error, expected 12.5, got %v", items[0]["price"])
	}

	// Verify second item null values
	if items[1]["name"] != nil {
		t.Errorf("Second item name should be null, got %v", items[1]["name"])
	}
	
	if items[1]["change"] != nil {
		t.Errorf("Second item change should be null, got %v", items[1]["change"])
	}
	
	// Verify second item negative number (price field, not change!)
	if price, ok := items[1]["price"].(float64); !ok {
		t.Errorf("Second item price should be float64, got %T", items[1]["price"])
	} else if price != -15.3 {
		t.Errorf("Second item price should be -15.3, got %v", price)
	}
}
