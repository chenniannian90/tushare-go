package reference

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

func TestTop10Holders(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Top10Holders(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.Top10Holders(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestTop10FloatHolders(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.Top10FloatHolders(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.Top10FloatHolders(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestStkHolderNumber(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.StkHolderNumber(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestPledgeStat(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.PledgeStat(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.PledgeStat(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestPledgeDetail(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"ts_code": "000001.SZ"}
	var fields []string
	resp, err := client.PledgeDetail(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}

	// Test without ts_code should return error
	_, err = client.PledgeDetail(map[string]string{}, fields)
	if err == nil {
		t.Errorf("Api should return an error when ts_code is missing")
	}
}

func TestRepurchase(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.Repurchase(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestShareFloat(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.ShareFloat(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestBlockTrade(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.BlockTrade(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestStkAccount(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.StkAccount(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
