package finance

import (
	"fmt"
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) Income(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "income", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) BalanceSheet(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "balancesheet", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) CashFlow(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "cashflow", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaIndicator(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_indicator", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaAudit(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_audit", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) FinaMainbz(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "fina_mainbz", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Forecast(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasAnnDate := params["ann_date"]
	if (!hasTsCode && !hasAnnDate) || (hasTsCode && hasAnnDate) {
		return nil, fmt.Errorf("need one argument ts_code or ann_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "forecast", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Express(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "express", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Dividend(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "dividend", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) DisclosureDate(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "disclosure_date", "token": c.getToken(), "params": params, "fields": fields,
	})
}
