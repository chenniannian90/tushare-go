package sdk

import (
	"errors"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	ErrEmptyToken = errors.New("token 不能为空")
)

// Config 保存 SDK 配置
type Config struct {
	Tokens     []string // token 数组，支持负载均衡
	tokenIndex uint64   // 轮询索引（原子操作）
	Endpoint   string
	HTTPClient *http.Client
}

// GetToken 获取一个可用的 token（支持负载均衡）
func (c *Config) GetToken() string {
	if len(c.Tokens) == 0 {
		return ""
	}

	// 如果有多个 token，使用轮询负载均衡
	if len(c.Tokens) > 1 {
		// 原子操作获取下一个索引
		i := atomic.AddUint64(&c.tokenIndex, 1)
		return c.Tokens[(i-1)%uint64(len(c.Tokens))]
	}

	// 只有一个 token，直接使用
	return c.Tokens[0]
}

// NewConfig 创建一个新的配置（向后兼容：单个 token）
func NewConfig(token string) (*Config, error) {
	if token == "" {
		return nil, ErrEmptyToken
	}

	return &Config{
		Tokens:   []string{token}, // 单个 token 也转为数组
		Endpoint: "https://api.tushare.pro",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     30 * time.Second,
			},
		},
	}, nil
}

// NewConfigWithTokens 创建带有多个 token 的配置（自动轮询负载均衡）
func NewConfigWithTokens(tokens []string) (*Config, error) {
	if len(tokens) == 0 {
		return nil, ErrEmptyToken
	}

	// 检查是否有空 token
	for _, token := range tokens {
		if token == "" {
			return nil, ErrEmptyToken
		}
	}

	return &Config{
		Tokens:   tokens,
		Endpoint: "https://api.tushare.pro",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     30 * time.Second,
			},
		},
	}, nil
}
