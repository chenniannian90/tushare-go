package finance

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

func TestIncome(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Income(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.Income(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestBalanceSheet(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.BalanceSheet(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.BalanceSheet(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestCashFlow(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.CashFlow(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.CashFlow(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestFinaIndicator(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.FinaIndicator(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFinaAudit(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.FinaAudit(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFinaMainbz(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.FinaMainbz(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestForecast(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Forecast(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestExpress(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Express(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDividend(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.Dividend(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDisclosureDate(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.DisclosureDate(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
