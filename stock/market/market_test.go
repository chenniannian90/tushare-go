package market

import (
	"testing"

	"github.com/chenniannian90/tushare-go/types"
)

var testClient *Client

func setupTestClient() *Client {
	if testClient == nil {
		postFunc := func(body map[string]interface{}) (*types.APIResponse, error) {
			return &types.APIResponse{Code: 0}, nil
		}
		tokenFunc := func() string { return "" }
		testClient = New(postFunc, tokenFunc)
	}
	return testClient
}

func TestMoneyFlow(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.MoneyFlow(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyInfo(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.DailyInfo(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestSzDailyInfo(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.SzDailyInfo(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDaily(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Daily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyInvalidParams(t *testing.T) {
	client := setupTestClient()
	var fields []string

	// Test without ts_code or trade_date
	_, err := client.Daily(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when neither ts_code nor trade_date is provided")
	}

	// Test with both ts_code and trade_date
	_, err = client.Daily(map[string]string{"ts_code": "000001.SZ", "trade_date": "20240101"}, fields)
	if err == nil {
		t.Errorf("Api should return an error when both ts_code and trade_date are provided")
	}
}

func TestWeekly(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Weekly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMonthly(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Monthly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyBasic(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.DailyBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestAdjFactor(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.AdjFactor(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestSuspend(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Suspend(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIsDateFormat(t *testing.T) {
	tests := []struct {
		name string
		date string
		want bool
	}{
		{"Valid date", "20240101", true},
		{"Valid date", "20231231", true},
		{"Invalid date - short", "202301", false},
		{"Invalid date - long", "202401011", false},
		{"Invalid date - letters", "2024010a", false},
		{"Empty string", "", true}, // Empty string is valid for optional parameters
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDateFormat(tt.date)
			if got != tt.want {
				t.Errorf("IsDateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
