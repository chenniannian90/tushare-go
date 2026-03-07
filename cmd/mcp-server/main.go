package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chenniannian90/tushare-go/pkg/mcp"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

func main() {
	// Get Tushare token from environment
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	// Create SDK client
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create SDK config: %v", err)
	}

	client := sdk.NewClient(config)

	// Create MCP server
	server := mcp.NewStdioMCPServer(client)

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Start(ctx)
	}()

	// Wait for either error or signal
	select {
	case err := <-errChan:
		if err != nil {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	case sig := <-sigChan:
		log.Printf("Received signal %v, shutting down...", sig)
		cancel()
		os.Exit(0)
	}
}
