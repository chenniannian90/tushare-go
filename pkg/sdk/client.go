package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Context key type for storing token in context
type contextKey string

const tokenKey contextKey = "tushare_token"

// WithToken adds a token to the context for API calls
func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

// GetTokenFromContext extracts the token from the context
func GetTokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
}

// Client 表示 Tushare API 客户端
type Client struct {
	config *Config
}

// NewClient 创建一个新的 API 客户端
func NewClient(config *Config) *Client {
	return &Client{config: config}
}

// CallAPI 向 Tushare 发起 API 调用
func (c *Client) CallAPI(
	ctx context.Context,
	apiName string,
	params map[string]interface{},
	fields []string,
	result interface{},
) error {
	// Determine which token to use
	token := c.config.GetToken() // 使用负载均衡器获取 token
	if ctxToken, ok := GetTokenFromContext(ctx); ok {
		token = ctxToken // context token 优先级最高
	}

	// 构建请求体
	reqBody := map[string]interface{}{
		"api_name": apiName,
		"token":    token,
		"params":   params,
		"fields":   strings.Join(fields, ","),
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("请求序列化失败: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, "POST", c.config.Endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return WrapNetworkError(err, "请求失败")
	}
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return WrapAPIErrorWithCode(
			ClassifyAPIError(resp.StatusCode, ""),
			fmt.Sprintf("HTTP 请求失败，状态码 %d", resp.StatusCode),
			resp.StatusCode,
			nil,
		)
	}

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &NetworkError{Message: "读取响应失败", Err: err}
	}

	// 解析响应
	var apiResp struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return fmt.Errorf("响应反序列化失败: %w", err)
	}

	// 检查 API 错误
	if apiResp.Code != 0 {
		errorCode := ClassifyAPIError(apiResp.Code, apiResp.Msg)
		return WrapAPIErrorWithCode(errorCode, apiResp.Msg, apiResp.Code, nil)
	}

	// 解析数据
	if err := json.Unmarshal(apiResp.Data, result); err != nil {
		return fmt.Errorf("数据反序列化失败: %w", err)
	}

	return nil
}
