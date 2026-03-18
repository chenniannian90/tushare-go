package tushare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRTK(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.RTK(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code")
	}
	resp, err := client.RTK(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRealTimeQuote(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["ts_code"] = "600000.SH"
	params["src"] = "dc"
	var fields []string
	// Check params
	_, err := client.RealTimeQuote(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code")
	}

	params["trade_date"] = "20181101"
	resp, err := client.RealTimeQuote(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestRealTimeTick(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.RealTimeTick(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code")
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

func TestRealTimeList(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.RealTimeList(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "need one argument ts_code")
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
