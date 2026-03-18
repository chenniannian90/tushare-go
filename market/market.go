package market

import (
	"fmt"
	"regexp"
	"github.com/chenniannian90/tushare-go/types"
)

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func IsDateFormat(dates ...string) bool {
	pattern := `^\d{8}$`
	match, _ := regexp.MatchString(pattern, dates[0])
	return match
}

func (c *Client) MoneyFlow(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "moneyflow", "token": c.getToken(), "params": params, "fields": fields})
}

func (c *Client) DailyInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "daily_info", "token": c.getToken(), "params": params, "fields": fields})
}

func (c *Client) SzDailyInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{"api_name": "sz_daily_info", "token": c.getToken(), "params": params, "fields": fields})
}

func (c *Client) Daily(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "daily", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) Weekly(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "weekly", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) Monthly(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "monthly", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) DailyBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "daily_basic", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) AdjFactor(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	if !IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "adj_factor", "token": c.getToken(), "fields": fields, "params": params})
}

func (c *Client) Suspend(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["suspend_date"]
	_, hasResumeDate := params["resume_date"]
	argsCount := 0
	if hasTsCode { argsCount++ }
	if hasTradeDate { argsCount++ }
	if hasResumeDate { argsCount++ }
	if argsCount != 1 {
		return nil, fmt.Errorf("need one argument among ts_code, suspend_date, resume_date")
	}
	if !IsDateFormat(params["suspend_date"], params["resume_date"]) {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}
	return c.postData(map[string]interface{}{"api_name": "suspend", "token": c.getToken(), "params": params, "fields": fields})
}
