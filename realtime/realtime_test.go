package realtime

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

func TestRTK(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.RTK(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRealTimeQuote(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "600000.SH", "src": "dc"}
	var fields []string
	_, err := client.RealTimeQuote(params, fields)
	if err != nil {
		t.Errorf("need one argument ts_code, got: %s", err)
	}

	params["ts_code"] = "600000.SH"
	resp, err := client.RealTimeQuote(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRealTimeTick(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	_, err := client.RealTimeTick(params, fields)
	if err != nil {
		t.Logf("Expected error when ts_code is missing: %s", err)
	}

	params["ts_code"] = "600000.SH"
	resp, err := client.RealTimeTick(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRealTimeList(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.RealTimeList(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
