package special

import "github.com/chenniannian90/tushare-go/types"

type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{postData: postData, getToken: getToken}
}

func (c *Client) CyqChips(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "cyq_chips", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) StkSurv(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "stk_surv", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) HmList(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "hm_list", "token": c.getToken(), "params": params, "fields": fields,
	})
}

func (c *Client) HmDetail(params map[string]string, fields []string) (*types.APIResponse, error) {
	return c.postData(map[string]interface{}{
		"api_name": "hm_detail", "token": c.getToken(), "params": params, "fields": fields,
	})
}
