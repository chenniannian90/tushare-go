package sdk

import (
	"errors"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "valid config",
			token:   "test-token-12345",
			wantErr: false,
		},
		{
			name:    "empty token",
			token:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Tokens[0] != tt.token {
					t.Errorf("NewConfig().Tokens[0] = %v, want %v", got.Tokens[0], tt.token)
				}
				if got.Endpoint != "https://api.tushare.pro" {
					t.Errorf("NewConfig().Endpoint = %v, want https://api.tushare.pro", got.Endpoint)
				}
			}
		})
	}
}

func TestConfig_DefaultValues(t *testing.T) {
	config, err := NewConfig("test-token")
	if err != nil {
		t.Fatalf("NewConfig() failed: %v", err)
	}

	if config.Endpoint != "https://api.tushare.pro" {
		t.Errorf("default Endpoint = %v, want https://api.tushare.pro", config.Endpoint)
	}

	if config.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}

	if config.HTTPClient.Timeout != 30*time.Second {
		t.Errorf("default Timeout = %v, want 30s", config.HTTPClient.Timeout)
	}
}

func TestConfig_EmptyToken(t *testing.T) {
	config, err := NewConfig("")
	if config != nil {
		t.Error("NewConfig() should return nil config for empty token")
	}

	if !errors.Is(err, ErrEmptyToken) {
		t.Errorf("NewConfig() error = %v, want ErrEmptyToken", err)
	}
}

func TestNewConfigWithTokens(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []string
		wantErr bool
	}{
		{
			name:    "valid tokens with multiple tokens",
			tokens:  []string{"token1", "token2", "token3"},
			wantErr: false,
		},
		{
			name:    "valid tokens with two tokens",
			tokens:  []string{"token1", "token2"},
			wantErr: false,
		},
		{
			name:    "single token",
			tokens:  []string{"token1"},
			wantErr: false,
		},
		{
			name:    "empty tokens",
			tokens:  []string{},
			wantErr: true,
		},
		{
			name:    "token with empty string",
			tokens:  []string{"token1", "", "token3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfigWithTokens(tt.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfigWithTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got.Tokens) != len(tt.tokens) {
					t.Errorf("NewConfigWithTokens().Tokens length = %v, want %v", len(got.Tokens), len(tt.tokens))
				}
				// 验证所有 token 都正确保存
				for i, token := range tt.tokens {
					if got.Tokens[i] != token {
						t.Errorf("NewConfigWithTokens().Tokens[%d] = %v, want %v", i, got.Tokens[i], token)
					}
				}
			}
		})
	}
}

func TestRoundRobinBalancer(t *testing.T) {
	// 测试内置的轮询负载均衡
	tokens := []string{"token1", "token2", "token3"}
	config, _ := NewConfigWithTokens(tokens)

	// 测试轮询顺序
	results := make([]string, 6)
	for i := 0; i < 6; i++ {
		results[i] = config.GetToken()
	}

	expected := []string{"token1", "token2", "token3", "token1", "token2", "token3"}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("GetToken() iteration %d = %v, want %v", i+1, got, expected[i])
		}
	}
}

func TestConfig_GetToken(t *testing.T) {
	tests := []struct {
		name        string
		setupConfig func() *Config
		testFunc    func(*Config)
	}{
		{
			name: "use single token",
			setupConfig: func() *Config {
				config, _ := NewConfig("single-token")
				return config
			},
			testFunc: func(c *Config) {
				token := c.GetToken()
				if token != "single-token" {
					t.Errorf("GetToken() = %v, want single-token", token)
				}
			},
		},
		{
			name: "use multiple tokens with round robin",
			setupConfig: func() *Config {
				config, _ := NewConfigWithTokens([]string{"token1", "token2", "token3"})
				return config
			},
			testFunc: func(c *Config) {
				// 测试多次调用，验证轮询
				token1 := c.GetToken()
				token2 := c.GetToken()
				token3 := c.GetToken()
				token4 := c.GetToken()

				if token1 != "token1" || token2 != "token2" || token3 != "token3" || token4 != "token1" {
					t.Errorf("Round robin failed: %v, %v, %v, %v", token1, token2, token3, token4)
				}
			},
		},
		{
			name: "use single token from array",
			setupConfig: func() *Config {
				config, _ := NewConfigWithTokens([]string{"token1"})
				return config
			},
			testFunc: func(c *Config) {
				token := c.GetToken()
				if token != "token1" {
					t.Errorf("GetToken() = %v, want token1", token)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.setupConfig()
			tt.testFunc(config)
		})
	}
}
