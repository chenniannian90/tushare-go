package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestWeekly(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "trade_date", "open", "high", "low", "close", "vol", "amount"],
				"items": [
					{"ts_code": "000001.SZ", "trade_date": "20240105", "open": 10.5, "high": 11.0, "low": 10.2, "close": 10.8, "vol": 1000000, "amount": 10800000}
				]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &WeeklyRequest{
		TsCode: "000001.SZ",
	}

	items, err := Weekly(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].TsCode != "000001.SZ" {
		t.Errorf("expected ts_code '000001.SZ', got %s", items[0].TsCode)
	}

	if items[0].Close != 10.8 {
		t.Errorf("expected close 10.8, got %f", items[0].Close)
	}
}

func TestWeekly_APIError(t *testing.T) {
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

	req := &WeeklyRequest{}

	_, err := Weekly(context.Background(), client, req)
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

func TestWeekly_WithDateRange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "trade_date", "open", "high", "low", "close", "vol", "amount"],
				"items": []
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &WeeklyRequest{
		TsCode:    "000001.SZ",
		StartDate: "20240101",
		EndDate:   "20240131",
	}

	_, err := Weekly(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
