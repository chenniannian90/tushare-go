package server

import (
	"context"
	"testing"

	"tushare-go/pkg/sdk"
)

func TestNewServer(t *testing.T) {
	config, _ := sdk.NewConfig("test-token")
	client := sdk.NewClient(config)

	server := NewServer(client)
	if server == nil {
		t.Error("NewServer() should not return nil")
	}

	if server.Client != client {
		t.Error("NewServer() should store the client")
	}
}

func TestServer_GetTools(t *testing.T) {
	config, _ := sdk.NewConfig("test-token")
	client := sdk.NewClient(config)
	server := NewServer(client)

	tools := server.GetTools()
	if tools == nil {
		t.Error("GetTools() should not return nil")
	}

	// Note: Server starts with no tools registered. Tools need to be registered manually.
	// The test just verifies that GetTools() works and returns an empty list initially.
	if len(tools) != 0 {
		t.Errorf("GetTools() should return no tools initially, got %d", len(tools))
	}
}

func TestServer_CallTool(t *testing.T) {
	config, _ := sdk.NewConfig("test-token")
	client := sdk.NewClient(config)
	server := NewServer(client)

	tests := []struct {
		name     string
		toolName string
		args     map[string]interface{}
		wantErr  bool
	}{
		{
			name:     "call stock_basic tool",
			toolName: "stock_basic.stock_basic",
			args: map[string]interface{}{
				"ts_code": "000001.SZ",
			},
			wantErr: false, // Returns mock result, not an error
		},
		{
			name:     "unknown tool",
			toolName: "unknown_tool",
			args:     map[string]interface{}{},
			wantErr:  false, // Server now returns a message instead of error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := server.CallTool(context.Background(), tt.toolName, tt.args)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result == nil {
					t.Error("result should not be nil")
				}
			}
		})
	}
}
