package finance

import (
	"testing"

	"github.com/chenniannian90/tushare-go/types"
	"github.com/stretchr/testify/assert"
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

func TestInCome(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Income(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Income(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestBalanceSheet(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.BalanceSheet(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.BalanceSheet(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestCashFlow(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.CashFlow(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.CashFlow(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestForecast(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Forecast(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or ann_date")
	}
	params["ts_code"] = "000001.SZ"
	resp, err := client.Forecast(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDividend(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.Dividend(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestExpress(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Express(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
}

func TestFinaIndicator(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.FinaIndicator(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
}

func TestFinaAudit(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.FinaAudit(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
}

func TestFinaMainbz(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.FinaMainbz(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "ts_code is a required argument")
	}
}

func TestDisclosureDate(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.DisclosureDate(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
