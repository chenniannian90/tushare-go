package index

import (
	"github.com/chenniannian90/tushare-go/types"
)

// Client represents the index API client
type Client struct {
	postData types.PostFunc
	getToken types.TokenFunc
}

// New creates a new index client
func New(postData types.PostFunc, getToken types.TokenFunc) *Client {
	return &Client{
		postData: postData,
		getToken: getToken,
	}
}

// IndexDaily 指数日线行情
func (c *Client) IndexDaily(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_daily",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexDailyBasic 大盘指数每日指标
func (c *Client) IndexDailyBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_dailybasic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexClassify 申万行业分类
func (c *Client) IndexClassify(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_classify",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexGlobal 国际指数
func (c *Client) IndexGlobal(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_global",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexBasic 指数基本信息
func (c *Client) IndexBasic(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_basic",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexWeight 指数成分和权重
func (c *Client) IndexWeight(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_weight",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}

// IndexMember 申万行业成分构成
func (c *Client) IndexMember(params map[string]string, fields []string) (*types.APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_member",
		"token":    c.getToken(),
		"params":   params,
		"fields":   fields,
	}
	return c.postData(body)
}
