package config

import (
	"encoding/json"
	"os"
)

// AuthConfig defines authentication configuration for a service
type AuthConfig struct {
	Type     string `json:"type"`     // "none", "apikey"
	Required bool   `json:"required"` // whether auth is required
}

// LoggingConfig defines the configuration for the logger
type LoggingConfig struct {
	Level      string `json:"level"`       // Log level: debug, info, warn, error
	Format     string `json:"format"`      // Log format: json, text
	Filename   string `json:"filename"`    // Log file path, empty means stdout
	MaxSize    int    `json:"max_size"`    // Maximum size in megabytes
	MaxAge     int    `json:"max_age"`     // Maximum number of days to retain old log files
	MaxBackups int    `json:"max_backups"` // Maximum number of old log files to retain
	Compress   bool   `json:"compress"`    // Compress rotated files
}

// ServiceConfig defines configuration for a single MCP service
type ServiceConfig struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Description string      `json:"description"`
	Categories  []string    `json:"categories,omitempty"` // e.g., ["stock", "bond", "futures"]
	Auth        AuthConfig  `json:"auth,omitempty"`
}

// ServerConfig defines the overall server configuration
type ServerConfig struct {
	Host       string                   `json:"host"`
	Port       int                      `json:"port"`
	Transport  string                   `json:"transport"`   // "stdio" or "http"
	Services   map[string]ServiceConfig `json:"services"`    // named service configurations
	GlobalAuth AuthConfig               `json:"global_auth"` // fallback auth config
	APITokens  []string                 `json:"api_tokens"`  // list of valid api tokens for authentication
	Logging    *LoggingConfig           `json:"logging"`     // logging configuration
}

// LoadFile loads server configuration from a JSON file
func LoadFile(path string) (*ServerConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
