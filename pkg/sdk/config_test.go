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
				if got.Token != tt.token {
					t.Errorf("NewConfig().Token = %v, want %v", got.Token, tt.token)
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
