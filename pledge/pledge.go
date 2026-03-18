package pledge

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

func (c *Client) PledgeStat(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "pledge_stat", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) PledgeDetail(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	return c.postData(map[string]interface{}{
		"api_name": "pledge_detail", "token": c.getToken(), "params": params, "fields": fields,
	})
}
