package hsgt

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

func (c *Client) MoneyflowHsgt(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTradeDate := params["trade_date"]
	_, hasStartDate := params["start_date"]
	if (!hasTradeDate && !hasStartDate) || (hasTradeDate && hasStartDate) {
		return nil, fmt.Errorf("need one argument trade_date or start_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "moneyflow_hsgt", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) HsgtTop10(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "hsgt_top10", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) GgtTop10(params map[string]string, fields []string) (*types.APIResponse, error) {
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("need one argument ts_code or trade_date")
	}
	return c.postData(map[string]interface{}{
		"api_name": "ggt_top10", "token": c.getToken(), "params": params, "fields": fields,
	})
}
