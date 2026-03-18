package index

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

func TestIndexBasic(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexDaily(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexDaily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexDailyBasic(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexDailyBasic(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexClassify(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexClassify(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexGlobal(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexGlobal(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexWeight(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexWeight(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestIndexMember(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.IndexMember(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
