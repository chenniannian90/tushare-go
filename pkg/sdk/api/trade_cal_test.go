package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestTradeCal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["exchange", "cal_date", "is_open", "pretrade_date"],
				"items": [
					{"exchange": "SSE", "cal_date": "20240101", "is_open": "0", "pretrade_date": "20231229"}
				]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &TradeCalRequest{
		Exchange: "SSE",
	}

	items, err := TradeCal(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].Exchange != "SSE" {
		t.Errorf("expected exchange 'SSE', got %s", items[0].Exchange)
	}

	if items[0].IsOpen != "0" {
		t.Errorf("expected is_open '0', got %s", items[0].IsOpen)
	}
}

func TestTradeCal_APIError(t *testing.T) {
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

	req := &TradeCalRequest{}

	_, err := TradeCal(context.Background(), client, req)
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
}

func TestTradeCal_WithDateRange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["exchange", "cal_date", "is_open", "pretrade_date"],
				"items": []
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &TradeCalRequest{
		Exchange:  "SSE",
		StartDate: "20240101",
		EndDate:   "20240131",
	}

	_, err := TradeCal(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
