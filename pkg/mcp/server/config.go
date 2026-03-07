package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"tushare-go/pkg/sdk"
)

// ServerConfig holds MCP server configuration
type ServerConfig struct {
	Token           string
	APIKey          string
	Endpoint        string
	Timeout         time.Duration
	LogLevel        string
	ShutdownTimeout time.Duration
	RequireAuth     bool
}

// DefaultServerConfig returns default server configuration
func DefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Token:           os.Getenv("TUSHARE_TOKEN"),
		APIKey:          os.Getenv("MCP_API_KEY"),
		Endpoint:        os.Getenv("TUSHARE_ENDPOINT"),
		Timeout:         30 * time.Second,
		LogLevel:        os.Getenv("LOG_LEVEL"),
		ShutdownTimeout: 10 * time.Second,
		RequireAuth:     os.Getenv("MCP_REQUIRE_AUTH") == "true",
	}
}

// Validate validates the server configuration
func (c *ServerConfig) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("TUSHARE_TOKEN is required")
	}

	// Validate API Key if authentication is required
	if c.RequireAuth && c.APIKey == "" {
		return fmt.Errorf("MCP_API_KEY is required when MCP_REQUIRE_AUTH is true")
	}

	if c.Endpoint == "" {
		c.Endpoint = "https://api.tushare.pro"
	}

	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}

	if c.ShutdownTimeout == 0 {
		c.ShutdownTimeout = 10 * time.Second
	}

	if c.LogLevel == "" {
		c.LogLevel = "INFO"
	}

	return nil
}

// CreateSDKClient creates an SDK client from the configuration
func (c *ServerConfig) CreateSDKClient() (*sdk.Client, error) {
	config, err := sdk.NewConfig(c.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create SDK config: %w", err)
	}

	// Override endpoint if specified
	if c.Endpoint != "" && c.Endpoint != "https://api.tushare.pro" {
		config.Endpoint = c.Endpoint
	}

	// Override timeout if specified
	if c.Timeout != 30*time.Second {
		config.HTTPClient.Timeout = c.Timeout
	}

	return sdk.NewClient(config), nil
}

// ServerInfo holds server metadata
type ServerInfo struct {
	Name        string
	Version     string
	Author      string
	Description string
	Homepage    string
	License     string
}

// GetServerInfo returns server metadata
func GetServerInfo() ServerInfo {
	return ServerInfo{
		Name:        "tushare-mcp-server",
		Version:     "1.0.0",
		Author:      "Tushare Go SDK Contributors",
		Description: "MCP server for Tushare Pro - Chinese financial data platform",
		Homepage:    "https://tushare-go",
		License:     "MIT",
	}
}

// SetupLogging configures logging based on the configuration
func SetupLogging(logLevel string) {
	// Set up logging to stderr (MCP uses stdout for messages)
	log.SetOutput(os.Stderr)

	// Configure log level based on environment
	switch logLevel {
	case "DEBUG":
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	case "INFO":
		log.SetFlags(log.Ldate | log.Ltime)
	case "WARN", "ERROR":
		log.SetFlags(log.Ldate | log.Ltime)
	default:
		log.SetFlags(log.Ldate | log.Ltime)
	}
}

// LogStartup logs server startup information
func LogStartup(config *ServerConfig, info ServerInfo) {
	log.Printf("=== %s v%s ===", info.Name, info.Version)
	log.Printf("Author: %s", info.Author)
	log.Printf("Description: %s", info.Description)
	log.Printf("Endpoint: %s", config.Endpoint)
	log.Printf("Timeout: %v", config.Timeout)
	log.Printf("Log Level: %s", config.LogLevel)
	log.Printf("========================")
}

// GetCapabilities returns server capabilities
func GetCapabilities() map[string]interface{} {
	return map[string]interface{}{
		"tools": map[string]interface{}{
			"listChanged": false,
		},
		"resources": map[string]interface{}{
			"subscribe":   false,
			"listChanged": false,
		},
		"prompts": map[string]interface{}{
			"listChanged": false,
		},
	}
}
