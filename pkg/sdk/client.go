package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// logAPIError 记录API调用的详细错误信息
func logAPIError(apiName string, reqBody map[string]interface{}, respBody []byte, errMsg string) {
	reqJSON, _ := json.Marshal(reqBody)
	log.Printf("=== API调用失败 ===")
	log.Printf("API名称: %s", apiName)
	log.Printf("错误信息: %s", errMsg)
	log.Printf("请求体: %s", string(reqJSON))
	log.Printf("响应体: %s", string(respBody))
	log.Printf("===================")
}

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
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("请求序列化失败: %v", err))
		return fmt.Errorf("请求序列化失败: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, "POST", c.config.Endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("创建请求失败: %v", err))
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("HTTP请求失败: %v", err))
		return WrapNetworkError(err, "请求失败")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("读取响应失败: %v", err))
		return &NetworkError{Message: "读取响应失败", Err: err}
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("HTTP状态码错误: %d", resp.StatusCode))
		return WrapAPIErrorWithCode(
			ClassifyAPIError(resp.StatusCode, ""),
			fmt.Sprintf("HTTP 请求失败，状态码 %d", resp.StatusCode),
			resp.StatusCode,
			nil,
		)
	}

	// 解析响应
	var apiResp struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("响应反序列化失败: %v", err))
		return fmt.Errorf("响应反序列化失败: %w", err)
	}

	// 检查 API 错误
	if apiResp.Code != 0 {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("API错误: code=%d, msg=%s", apiResp.Code, apiResp.Msg))
		errorCode := ClassifyAPIError(apiResp.Code, apiResp.Msg)
		return WrapAPIErrorWithCode(errorCode, apiResp.Msg, apiResp.Code, nil)
	}

	// 解析数据
	if err := json.Unmarshal(apiResp.Data, result); err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("数据反序列化失败: %v\n原始数据: %s", err, string(apiResp.Data)))
		return fmt.Errorf("数据反序列化失败: %w", err)
	}

	return nil
}

// CallAPIFlexible 向 Tushare 发起 API 调用（支持灵活的响应格式）
// 这个方法会自动检测API返回的是对象数组还是二维数组，并统一转换为对象数组
// 向后兼容：返回的result会被填充为对象数组格式
func (c *Client) CallAPIFlexible(
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
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("请求序列化失败: %v", err))
		return fmt.Errorf("请求序列化失败: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, "POST", c.config.Endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("创建请求失败: %v", err))
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		logAPIError(apiName, reqBody, nil, fmt.Sprintf("HTTP请求失败: %v", err))
		return WrapNetworkError(err, "请求失败")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("读取响应失败: %v", err))
		return &NetworkError{Message: "读取响应失败", Err: err}
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("HTTP状态码错误: %d", resp.StatusCode))
		return WrapAPIErrorWithCode(
			ClassifyAPIError(resp.StatusCode, ""),
			fmt.Sprintf("HTTP 请求失败，状态码 %d", resp.StatusCode),
			resp.StatusCode,
			nil,
		)
	}

	// 解析响应
	var apiResp struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("响应反序列化失败: %v", err))
		return fmt.Errorf("响应反序列化失败: %w", err)
	}

	// 检查 API 错误
	if apiResp.Code != 0 {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("API错误: code=%d, msg=%s", apiResp.Code, apiResp.Msg))
		errorCode := ClassifyAPIError(apiResp.Code, apiResp.Msg)
		return WrapAPIErrorWithCode(errorCode, apiResp.Msg, apiResp.Code, nil)
	}

	// 使用灵活的响应解析器
	var rawResponse APIResponse
	if err := json.Unmarshal(apiResp.Data, &rawResponse); err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("解析响应结构失败: %v", err))
		return fmt.Errorf("解析响应结构失败: %w", err)
	}

	// 自动检测格式并转换
	items, err := rawResponse.ParseAndConvert()
	if err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("灵活解析数据失败: %v\n原始数据: %s", err, string(apiResp.Data)))
		return fmt.Errorf("灵活解析数据失败: %w", err)
	}

	// 构建统一的响应格式
	normalizedResponse := struct {
		Fields []string                 `json:"fields"`
		Items  []map[string]interface{} `json:"items"`
	}{
		Fields: rawResponse.Fields,
		Items:  items,
	}

	// 将标准化后的响应解析到用户提供的result中
	normalizedData, err := json.Marshal(normalizedResponse)
	if err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("序列化标准化响应失败: %v", err))
		return fmt.Errorf("序列化标准化响应失败: %w", err)
	}

	if err := json.Unmarshal(normalizedData, result); err != nil {
		logAPIError(apiName, reqBody, respBody, fmt.Sprintf("反序列化到目标结构体失败: %v", err))
		return fmt.Errorf("反序列化到目标结构体失败: %w", err)
	}

	return nil
}
