package holder

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

func (c *Client) Top10Holders(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "top10_holders", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) Top10FloatHolders(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "top10_floatholders", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) StkHolderNumber(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_holdernumber", "token": c.getToken(), "params": params, "fields": fields,
	})
}
