package main

import (
	"flag"
	"log"
	"os"

	"tushare-go/cmd/mcp-server/config"
	"tushare-go/pkg/sdk"
)

func main() {
	// Parse command-line flags
	configPath := flag.String("config", "", "Path to configuration file (JSON)")
	transport := flag.String("transport", "stdio", "Transport type: stdio, http, or both (overridden by config file)")
	addr := flag.String("addr", ":8080", "HTTP server address (for http/both transports) (overridden by config file)")
	flag.Parse()

	// Get Tushare token from environment
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	// Create SDK client
	sdkConfig, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create SDK config: %v", err)
	}

	client := sdk.NewClient(sdkConfig)

	// Load server configuration
	var serverConfig *config.ServerConfig
	if *configPath != "" {
		// Load from configuration file
		serverConfig, err = config.LoadFile(*configPath)
		if err != nil {
			log.Fatalf("Failed to load config file: %v", err)
		}
		log.Printf("Loaded configuration from: %s", *configPath)
	} else {
		// Use default configuration with command-line overrides
		serverConfig = config.DefaultConfig(*transport, *addr)
	}

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
