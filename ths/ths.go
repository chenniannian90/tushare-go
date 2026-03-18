package ths

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) ThsIndex(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_index", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ThsDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) ThsMember(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "ths_member", "token": c.getToken(), "params": params, "fields": fields,
	})
}
