package margin

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

func TestMargin(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Margin(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code or trade_date should return error
	_, err = client.Margin(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when neither ts_code nor trade_date is provided")
	}

	// Test with both ts_code and trade_date
	_, err = client.Margin(map[string]string{"ts_code": "000001.SZ", "trade_date": "20240101"}, fields)
	if err == nil {
		t.Errorf("Api should return an error when both ts_code and trade_date are provided")
	}
}

func TestMarginDetail(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.MarginDetail(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
