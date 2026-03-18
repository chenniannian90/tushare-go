package toplist

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) TopList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "top_list", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) TopInst(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "top_inst", "token": c.getToken(), "params": params, "fields": fields,
	})
}
