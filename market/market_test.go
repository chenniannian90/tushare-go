package market

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

func TestDaily(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Daily(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Daily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := map[string]string{"trade_date": "2018-11-01"}
	var fields []string
	_, err := client.Daily(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "please input right date format YYYYMMDD")
	}
}

func TestWeekly(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Weekly(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Weekly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMonthly(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Monthly(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Monthly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyBasic(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.DailyBasic(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.DailyBasic(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestAdjFactor(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.AdjFactor(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.AdjFactor(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestSuspend(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	paramsTest := map[string]string{"ts_code": "000001.SZ", "resume_date": "20181102"}
	var fields []string
	_, err := client.Suspend(paramsTest, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument among ts_code, suspend_date, resume_date")
	}

	params["suspend_date"] = "20181101"
	resp, err := client.Suspend(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
