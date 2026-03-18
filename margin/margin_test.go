package margin

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

func TestMargin(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.Margin(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}
	params["trade_date"] = "20181101"
	resp, err := client.Margin(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestMarginDetail(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.MarginDetail(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code or trade_date")
	}
	params["trade_date"] = "20181101"
	resp, err := client.MarginDetail(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
