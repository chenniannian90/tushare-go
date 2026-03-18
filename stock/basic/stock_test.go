package basic

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

func TestStockBasic(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.StockBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestBakBasic(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.BakBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTradeCal(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.TradeCal(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHSConst(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"hs_type": "SH"}
	var fields []string
	resp, err := client.HSConst(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without hs_type should return error
	_, err = client.HSConst(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when hs_type is missing")
	}
}

func TestNameChange(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.NameChange(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStockCompany(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.StockCompany(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestNewShare(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.NewShare(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
