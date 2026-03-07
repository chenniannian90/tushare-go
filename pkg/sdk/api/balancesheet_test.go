package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func TestBalancesheet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "ann_date", "f_ann_date", "end_date", "report_type", "total_assets", "total_hldr_eqy_exc_min_int", "total_liab"],
				"items": [
					{"ts_code": "000001.SZ", "ann_date": "20240101", "f_ann_date": "20240101", "end_date": "20231231", "report_type": "0", "total_assets": 1000000000, "total_hldr_eqy_exc_min_int": 500000000, "total_liab": 500000000}
				]
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &BalancesheetRequest{
		TsCode: "000001.SZ",
	}

	items, err := Balancesheet(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].TsCode != "000001.SZ" {
		t.Errorf("expected ts_code '000001.SZ', got %s", items[0].TsCode)
	}

	if items[0].TotalAssets != 1000000000 {
		t.Errorf("expected total_assets 1000000000, got %f", items[0].TotalAssets)
	}
}

func TestBalancesheet_APIError(t *testing.T) {
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

	req := &BalancesheetRequest{}

	_, err := Balancesheet(context.Background(), client, req)
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

func TestBalancesheet_WithPeriod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"fields": ["ts_code", "ann_date", "f_ann_date", "end_date", "report_type", "total_assets", "total_liab"],
				"items": []
			}
		}`))
	}))
	defer server.Close()

	config, _ := sdk.NewConfig("test-token")
	config.Endpoint = server.URL
	client := sdk.NewClient(config)

	req := &BalancesheetRequest{
		TsCode:     "000001.SZ",
		Period:     "20231231",
		ReportType: "0",
	}

	_, err := Balancesheet(context.Background(), client, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
