package stock

import (
	"fmt"
	"os"
	"testing"

	"github.com/chenniannian90/tushare-go/types"
	"github.com/stretchr/testify/assert"
)

var token = ""
var testClient *Client

func init() {
	envToken := os.Getenv("TUSHARE_TOKEN")
	if envToken != "" {
		token = envToken
	}
}

func setupTestClient() *Client {
	if testClient == nil {
		// Create a mock postData function for testing
		postFunc := func(body map[string]interface{}) (*types.APIResponse, error) {
			// In real tests, this would make an actual API call
			return &types.APIResponse{Code: 0}, nil
		}
		tokenFunc := func() string { return token }
		testClient = New(postFunc, tokenFunc)
	}
	return testClient
}

func TestStockBasic(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{
		"is_hs":       "N",
		"list_status": "L",
		"exchange":    "SSE",
	}
	var fields []string
	resp, err := client.StockBasic(params, fields)

	if err != nil && token != "" {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestInvalidField(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	fields := []string{"invalid_field"}
	resp, err := client.StockBasic(params, fields)

	if err != nil {
		if resp.Code == -2001 {
			ast.Equal(err.Error(), fmt.Sprintf("argument error: %s", resp.Msg))
		}
	}
}

func TestTradeCal(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{
		"exchange":   "SSE",
		"start_date": "2017-01-01",
		"end_date":   "2019-01-01",
		"is_open":    "1",
	}
	var fields []string
	resp, err := client.TradeCal(params, fields)

	if err != nil && token != "" {
		t.Errorf("Api should not return an error, got: %s", err)
	}

	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHSConst(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{
		"hs_type": "SH",
		"is_new":  "1",
	}
	var fields []string
	resp, err := client.HSConst(params, fields)

	if err != nil && token != "" {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestHSConstParamsRequired(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.HSConst(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "hs_type required")
	}
}

func TestNameChange(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{
		"ts_code":    "000001.SZ",
		"start_date": "2017-01-01",
		"end_date":   "2019-01-01",
	}
	var fields []string
	resp, err := client.NameChange(params, fields)

	if err != nil && token != "" {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStockCompany(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := map[string]string{
		"exchange": "SSE",
	}
	var fields []string
	resp, err := client.StockCompany(params, fields)

	if err != nil {
		if resp != nil && resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestNewShare(t *testing.T) {
	ast := assert.New(t)
	client := setupTestClient()
	params := map[string]string{
		"start_date": "2017-01-01",
		"end_date":   "2019-01-01",
	}
	var fields []string
	resp, err := client.NewShare(params, fields)

	if err != nil {
		if resp != nil && resp.Code == -2002 {
			ast.Equal(err.Error(), "your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
