package repurchase

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) Repurchase(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "repurchase", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ShareFloat(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "share_float", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) BlockTrade(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "block_trade", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) StkAccount(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_account", "token": c.getToken(), "params": params, "fields": fields,
	})
}
