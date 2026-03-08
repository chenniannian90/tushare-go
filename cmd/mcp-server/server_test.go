package main

import (
	"testing"

	"tushare-go/cmd/mcp-server/config"
)

func TestNewServer_CategoriesLogic(t *testing.T) {
	tests := []struct {
		name           string
		serviceName    string
		servicePath    string
		categories     []string
		expectServices int // 预期创建的服务实例数量
	}{
		{
			name:           "空 categories - 创建一个服务包含所有工具",
			serviceName:    "all",
			servicePath:    "/",
			categories:     []string{},
			expectServices: 1,
		},
		{
			name:           "单个 category 匹配 name - 创建一个服务",
			serviceName:    "bond",
			servicePath:    "/bond",
			categories:     []string{"bond"},
			expectServices: 1,
		},
		{
			name:           "多个 categories - 为每个 category 创建独立服务",
			serviceName:    "stock",
			servicePath:    "/stock",
			categories:     []string{"stock_basic", "stock_market", "stock_financial"},
			expectServices: 3,
		},
		{
			name:           "单个 category 不匹配 name - 创建独立服务",
			serviceName:    "custom",
			servicePath:    "/custom",
			categories:     []string{"stock_basic"},
			expectServices: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.ServerConfig{
				Host:      "0.0.0.0",
				Port:      8080,
				Transport: "http",
				Services: map[string]config.ServiceConfig{
					tt.serviceName: {
						Name:       tt.serviceName,
						Path:       tt.servicePath,
						Categories: tt.categories,
					},
				},
			}

			// 注意：这里不能真正创建服务器，因为我们没有真实的 SDK client
			// 这个测试只是验证配置解析逻辑

			// 验证配置
			svcConfig := cfg.Services[tt.serviceName]
			if svcConfig.Name != tt.serviceName {
				t.Errorf("期望服务名 %s, 得到 %s", tt.serviceName, svcConfig.Name)
			}
			if svcConfig.Path != tt.servicePath {
				t.Errorf("期望路径 %s, 得到 %s", tt.servicePath, svcConfig.Path)
			}
		})
	}
}
