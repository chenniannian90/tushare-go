package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestStockBasic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "symbol", "name", "area", "industry"],
				"items": [
					{"ts_code": "000001.SZ", "symbol": "000001", "name": "平安银行", "area": "深圳", "industry": "银行"}
				]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &StockBasicRequest{
		TsCode: "000001.SZ",
	}

	items, err := StockBasic(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].Name != "平安银行" {
		t.Errorf("expected name '平安银行', got %s", items[0].Name)
	}

	if items[0].TsCode != "000001.SZ" {
		t.Errorf("expected ts_code '000001.SZ', got %s", items[0].TsCode)
	}

	if items[0].Industry != "银行" {
		t.Errorf("expected industry '银行', got %s", items[0].Industry)
	}
}

func TestStockBasic_EmptyResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "symbol", "name", "area", "industry"],
				"items": []
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &StockBasicRequest{
		TsCode: "000001.SZ",
	}

	items, err := StockBasic(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("expected 0 items, got %d", len(items))
	}
}

func TestStockBasic_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 40203,
			"msg": "insufficient privileges"
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &StockBasicRequest{}

	_, err := StockBasic(context.Background(), client, req)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*sdk.APIError)
	if !ok {
		t.Fatalf("expected APIError, got %T", err)
	}

	if apiErr.Code != sdk.ErrAccessDenied {
		t.Errorf("expected code ACCESS_DENIED, got %s", apiErr.Code)
	}
	if apiErr.APICode != 40203 {
		t.Errorf("expected APICode 40203, got %d", apiErr.APICode)
	}
}

func TestStockBasic_WithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body to verify params
		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("failed to decode request: %v", err)
			return
		}

		params, ok := reqBody["params"].(map[string]interface{})
		if !ok {
			t.Error("params should be a map")
			return
		}

		// Verify optional params are passed
		if params["list_status"] != "L" {
			t.Errorf("expected list_status 'L', got %v", params["list_status"])
		}

		if params["exchange"] != "SSE" {
			t.Errorf("expected exchange 'SSE', got %v", params["exchange"])
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "symbol", "name", "area", "industry"],
				"items": []
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &StockBasicRequest{
		ListStatus: "L",
		Exchange:   "SSE",
	}

	_, err := StockBasic(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
