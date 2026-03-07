package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

// StdioMCPServer implements a simple MCP server using stdio
type StdioMCPServer struct {
	Client         *sdk.Client
	Registry       *ToolRegistry
	Health         *HealthManager
	Lifecycle      *LifecycleManager
}

// NewStdioMCPServer creates a new stdio-based MCP server
func NewStdioMCPServer(client *sdk.Client) *StdioMCPServer {
	return &StdioMCPServer{
		Client:    client,
		Registry:  NewToolRegistry(),
		Health:    NewHealthManager(),
		Lifecycle: NewLifecycleManager(),
	}
}

// Start starts the MCP server and processes messages from stdin
func (s *StdioMCPServer) Start(ctx context.Context) error {
	// Setup logging
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}
	SetupLogging(logLevel)

	log.Printf("Starting Tushare MCP Server...")

	// Log server info
	config := DefaultServerConfig()
	if err := config.Validate(); err != nil {
		log.Printf("Warning: %v", err)
	}

	info := GetServerInfo()
	LogStartup(config, info)

	// Log available tools
	tools := s.Registry.GetTools()
	log.Printf("Registered %d tools:", len(tools))
	for _, tool := range tools {
		log.Printf("  - %s: %s", tool.Name, tool.Description)
	}

	// Process messages from stdin
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		select {
		case <-ctx.Done():
			log.Println("Server shutting down...")
			return nil
		default:
			var request map[string]interface{}
			if err := decoder.Decode(&request); err != nil {
				// EOF is expected when client closes connection
				if err.Error() != "EOF" {
					log.Printf("Error decoding request: %v", err)
				}
				continue
			}

			// Process request
			response := s.handleRequest(ctx, request)
			if err := encoder.Encode(response); err != nil {
				log.Printf("Error encoding response: %v", err)
			}
		}
	}
}

// handleRequest processes an MCP request
func (s *StdioMCPServer) handleRequest(ctx context.Context, request map[string]interface{}) map[string]interface{} {
	// Extract request method and params
	method, _ := request["method"].(string)
	params, _ := request["params"].(map[string]interface{})
	id := request["id"]

	log.Printf("Received request: %s", method)

	switch method {
	case "tools/list":
		return s.listTools(id)
	case "tools/call":
		return s.callTool(ctx, params, id)
	case "initialize":
		return s.initialize(id)
	default:
		return s.errorResponse(id, -32601, fmt.Sprintf("Method not supported: %s", method))
	}
}

// initialize handles the initialize request
func (s *StdioMCPServer) initialize(id interface{}) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result": map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"serverInfo": map[string]interface{}{
				"name":    "tushare-mcp-server",
				"version": "1.0.0",
			},
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
		},
	}
}

// listTools handles the tools/list request
func (s *StdioMCPServer) listTools(id interface{}) map[string]interface{} {
	tools := s.Registry.GetTools()

	mcpTools := make([]map[string]interface{}, 0, len(tools))
	for _, tool := range tools {
		mcpTools = append(mcpTools, map[string]interface{}{
			"name":        tool.Name,
			"description": tool.Description,
		})
	}

	return map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result": map[string]interface{}{
			"tools": mcpTools,
		},
	}
}

// callTool handles the tools/call request
func (s *StdioMCPServer) callTool(ctx context.Context, params map[string]interface{}, id interface{}) map[string]interface{} {
	// Extract tool name and arguments
	name, _ := params["name"].(string)
	arguments, _ := params["arguments"].(map[string]interface{})

	log.Printf("Calling tool: %s with args: %v", name, arguments)

	// Record request
	s.Health.RecordRequest()

	// Call the tool
	result, err := s.Registry.CallTool(ctx, s.Client, name, arguments)
	if err != nil {
		s.Health.RecordError()
		return s.errorResponse(id, -32603, fmt.Sprintf("Tool execution error: %v", err))
	}

	// Convert result to MCP format
	content := make([]map[string]interface{}, 0, len(result.Content))
	for _, c := range result.Content {
		content = append(content, map[string]interface{}{
			"type": c.Type,
			"text": c.Text,
		})
	}

	return map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result": map[string]interface{}{
			"content": content,
		},
	}
}

// errorResponse creates an error response
func (s *StdioMCPServer) errorResponse(id interface{}, code int, message string) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
}
