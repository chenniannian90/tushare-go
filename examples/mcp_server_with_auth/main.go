package main

import (
	"fmt"
	"log"
	"os"

	"tushare-go/pkg/mcp/server"
)

func main() {
	// Example: How to use MCP server with API Key authentication

	// Set up environment variables for demonstration
	// In production, these would be set externally
	os.Setenv("TUSHARE_TOKEN", "your-tushare-token-here")
	os.Setenv("MCP_API_KEY", "your-mcp-api-key-here")
	os.Setenv("MCP_REQUIRE_AUTH", "true")

	// Create server configuration
	config := server.DefaultServerConfig()

	// Validate configuration
	if err := config.Validate(); err != nil {
		log.Fatalf("Configuration validation failed: %v", err)
	}

	// Log configuration details
	fmt.Printf("=== MCP Server Configuration ===\n")
	fmt.Printf("Tushare Endpoint: %s\n", config.Endpoint)
	fmt.Printf("Authentication Required: %v\n", config.RequireAuth)
	if config.RequireAuth {
		fmt.Printf("API Key: %s***\n", config.APIKey[:4]) // Show only first 4 chars
	}
	fmt.Printf("==============================\n\n")

	// Create SDK client
	client, err := config.CreateSDKClient()
	if err != nil {
		log.Fatalf("Failed to create SDK client: %v", err)
	}

	// Create MCP server with configuration
	mcpServer := server.NewStdioMCPServer(client, config)

	// Simulate client initialization with correct API key
	initParams := map[string]interface{}{
		"clientInfo": map[string]interface{}{
			"name":    "test-client",
			"version": "1.0.0",
			"apiKey":  config.APIKey, // Use the same API key from config
		},
	}

	// Test initialization
	response := mcpServer.Initialize(1, initParams)

	// Check if initialization succeeded
	if errMsg, ok := response["error"]; ok {
		log.Printf("❌ Initialization failed: %v", errMsg)
		os.Exit(1)
	}

	fmt.Println("✅ MCP Server initialized successfully!")
	fmt.Println("Server is ready to handle tool calls with API key authentication.")

	// In a real scenario, the server would start processing MCP protocol messages
	// from stdin and responding to stdout
	fmt.Println("\nNote: In production, the server would run:")
	fmt.Println("  mcpServer.Start(context.Background())")

	// Example of what happens with wrong API key
	fmt.Println("\n=== Testing Authentication ===")

	wrongAPIKeyParams := map[string]interface{}{
		"clientInfo": map[string]interface{}{
			"name":    "malicious-client",
			"version": "1.0.0",
			"apiKey":  "wrong-api-key",
		},
	}

	response = mcpServer.Initialize(2, wrongAPIKeyParams)
	if errMsg, ok := response["error"]; ok {
		fmt.Printf("✅ Correctly rejected client with wrong API key: %v\n", errMsg)
	} else {
		fmt.Println("❌ Should have rejected wrong API key!")
	}
}
