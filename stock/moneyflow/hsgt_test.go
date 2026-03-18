package moneyflow

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

func TestMoneyflowHsgt(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"trade_date": "20240101"}
	var fields []string
	resp, err := client.MoneyflowHsgt(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without trade_date or start_date should return error
	_, err = client.MoneyflowHsgt(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when neither trade_date nor start_date is provided")
	}
}

func TestHsgtTop10(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.HsgtTop10(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code or trade_date should return error
	_, err = client.HsgtTop10(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when neither ts_code nor trade_date is provided")
	}
}

func TestGgtTop10(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.GgtTop10(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code or trade_date should return error
	_, err = client.GgtTop10(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when neither ts_code nor trade_date is provided")
	}
}
