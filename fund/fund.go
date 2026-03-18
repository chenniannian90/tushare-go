package fund

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) OpenFundDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "fund_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) OpenFundInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "fund_info", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) OpenFundPort(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "fund_portfolio", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MFundDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "mfund_daily", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MFundInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "mfund_info", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) MFundPort(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "mfund_portfolio", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) BailFundInfo(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "bailfund_info", "token": c.getToken(), "params": params, "fields": fields,
	})
}
