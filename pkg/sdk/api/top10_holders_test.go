package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestTop10Holders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "ann_date", "end_date", "holder_name", "hold_amount", "hold_ratio"],
				"items": [
					{"ts_code": "000001.SZ", "ann_date": "20240101", "end_date": "20231231", "holder_name": "中国平安", "hold_amount": 1000000, "hold_ratio": 5.5}
				]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &Top10HoldersRequest{
		TsCode: "000001.SZ",
	}

	items, err := Top10Holders(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].TsCode != "000001.SZ" {
		t.Errorf("expected ts_code '000001.SZ', got %s", items[0].TsCode)
	}
}

func TestTop10Holders_APIError(t *testing.T) {
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

	req := &Top10HoldersRequest{}

	_, err := Top10Holders(context.Background(), client, req)
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
