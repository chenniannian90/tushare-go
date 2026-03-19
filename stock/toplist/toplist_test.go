package toplist

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

func TestTopList(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"trade_date": "20240101"}
	var fields []string
	resp, err := client.TopList(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestTopInst(t *testing.T) {
	client := setupTestClient()
	params := map[string]string{"trade_date": "20240101"}
	var fields []string
	resp, err := client.TopInst(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestConcept(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.Concept(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestConceptDetail(t *testing.T) {
	client := setupTestClient()
	params := make(map[string]string)
	var fields []string
	resp, err := client.ConceptDetail(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
