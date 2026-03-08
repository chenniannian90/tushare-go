package config

import (
	"fmt"
	"strings"
)

// DefaultConfig creates a default server configuration with multiple services
func DefaultConfig(transport, addr string) *ServerConfig {
	host, port := parseAddr(addr)

	return &ServerConfig{
		Host:      host,
		Port:      port,
		Transport: transport,
		Services: map[string]ServiceConfig{
			"all": {
				Name:        "all",
				Path:        "/",
				Description: "All Tushare APIs (stock, bond, futures, etc.)",
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"stock": {
				Name:        "stock",
				Path:        "/stock",
				Description: "Stock market data APIs",
				Categories:  []string{"stock_basic", "stock_market", "stock_financial", "stock_board", "stock_feature", "stock_fund_flow", "stock_margin", "stock_reference"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"bond": {
				Name:        "bond",
				Path:        "/bond",
				Description: "Bond market data APIs",
				Categories:  []string{"bond"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"futures": {
				Name:        "futures",
				Path:        "/futures",
				Description: "Futures market data APIs",
				Categories:  []string{"futures"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
		},
	}
}

// parseAddr parses host:port address
func parseAddr(addr string) (string, int) {
	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		return parts[0], parseInt(parts[1], 8080)
	}
	return "0.0.0.0", 8080
}

// parseInt parses string to int with default
func parseInt(s string, defaultVal int) int {
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
		return defaultVal
	}
	return i
}
