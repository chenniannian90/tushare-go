package realtime

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

func (c *Client) RTK(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		params["ts_code"] = "3*.SZ,6*.SH,0*.SZ,9*.BJ"
	}
	return c.postData(map[string]interface{}{
		"api_name": "rt_k", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) RealTimeQuote(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}
	return c.postData(map[string]interface{}{
		"api_name": "realtime_quote", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) RealTimeTick(params map[string]string, fields []string) (*types.APIResponse, error) {
	if _, ok := params["ts_code"]; !ok {
		return nil, fmt.Errorf("need one argument ts_code")
	}
	return c.postData(map[string]interface{}{
		"api_name": "realtime_tick", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) RealTimeList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "realtime_list", "token": c.getToken(), "params": params, "fields": fields,
	})
}
