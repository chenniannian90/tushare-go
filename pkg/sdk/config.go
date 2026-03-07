package sdk

import (
	"errors"
	"net/http"
	"time"
)

var (
	ErrEmptyToken = errors.New("token 不能为空")
)

// Config 保存 SDK 配置
type Config struct {
	Token      string
	Endpoint   string
	HTTPClient *http.Client
}

// NewConfig 创建一个带有默认值的新配置
func NewConfig(token string) (*Config, error) {
	if token == "" {
		return nil, ErrEmptyToken
	}

	return &Config{
		Token:    token,
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
