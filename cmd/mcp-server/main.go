package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"tushare-go/cmd/mcp-server/config"
	"tushare-go/pkg/sdk"
	"tushare-go/pkg/sdk/logger"
)

func main() {
	// Parse command-line flags
	showVersion := flag.Bool("version", false, "Show version information")
	configPath := flag.String("config", "config.json", "Path to configuration file (JSON)")
	flag.Parse()

	// Show version information if requested
	if *showVersion {
		fmt.Printf("Tushare MCP Server\n")
		fmt.Printf("Version: %s\n", GetVersion())
		fmt.Printf("Full Info: %s\n", GetFullVersionInfo())
		os.Exit(0)
	}

	// Load server configuration first to check for api_tokens
	var serverConfig *config.ServerConfig
	var err error
	if *configPath != "" {
		// Load from configuration file
		serverConfig, err = config.LoadFile(*configPath)
		if err != nil {
			log.Fatalf("Failed to load config file: %v", err)
		}
		log.Printf("Loaded configuration from: %s", *configPath)
	} else {
		log.Fatal("Configuration file is required. Please specify using --config flag")
	}

	// Initialize logging system
	if serverConfig.Logging != nil {
		logConfig := &logger.LogConfig{
			Level:      serverConfig.Logging.Level,
			Format:     serverConfig.Logging.Format,
			Filename:   serverConfig.Logging.Filename,
			MaxSize:    serverConfig.Logging.MaxSize,
			MaxAge:     serverConfig.Logging.MaxAge,
			MaxBackups: serverConfig.Logging.MaxBackups,
			Compress:   serverConfig.Logging.Compress,
		}
		logger.Init(logConfig)
		logger.Infof("Logging system initialized with config: level=%s, format=%s, file=%s",
			logConfig.Level, logConfig.Format, logConfig.Filename)
	}

	// Determine token for client creation
	// If api_tokens are configured in config file, use the first one as default
	// Otherwise, require TUSHARE_TOKEN environment variable
	var token string
	if len(serverConfig.APITokens) > 0 {
		token = serverConfig.APITokens[0]
		logger.Infof("Using token from configuration file (first of %d tokens)", len(serverConfig.APITokens))
	} else {
		token = os.Getenv("TUSHARE_TOKEN")
		if token == "" {
			logger.Fatal("Either TUSHARE_TOKEN environment variable or api_tokens in config file is required")
		}
		logger.Info("Using token from TUSHARE_TOKEN environment variable")
	}

	// Create SDK client
	sdkConfig, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create SDK config: %v", err)
	}

	client := sdk.NewClient(sdkConfig)

	// Create server
	srv, err := NewServer(serverConfig, client)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start server based on transport type
	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Server shutdown complete")
}
