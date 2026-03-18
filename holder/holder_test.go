package holder

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

func TestTop10Holders(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Top10Holders(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTop10FloatHolders(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Top10FloatHolders(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStkHolderNumber(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.StkHolderNumber(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
