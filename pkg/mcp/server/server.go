package server

import (
	"context"

	"tushare-go/pkg/sdk"
	"tushare-go/pkg/mcp"
	"tushare-go/pkg/mcp/common"
)

// Server represents the MCP server
type Server struct {
	Client  *sdk.Client
	Registry *mcp.ToolRegistry
}

// NewServer creates a new MCP server
func NewServer(client *sdk.Client) *Server {
	return &Server{
		Client:   client,
		Registry: mcp.NewToolRegistry(client),
	}
}

// NewHTTPMCPServer creates a new HTTP MCP server (alias for NewServer)
func NewHTTPMCPServer(client *sdk.Client, config interface{}) *Server {
	return NewServer(client)
}

// NewStdioMCPServer creates a new stdio MCP server (alias for NewServer)
func NewStdioMCPServer(client *sdk.Client, config interface{}) *Server {
	return NewServer(client)
}

// GetTools returns the list of available tools
func (s *Server) GetTools() []common.Tool {
	return s.Registry.GetTools()
}

// CallTool executes a tool call
func (s *Server) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	return s.Registry.CallTool(ctx, toolName, args)
}

// Initialize initializes the server with the given parameters
func (s *Server) Initialize(params map[string]interface{}) error {
	// TODO: Implement initialization logic
	return nil
}

// Start starts the server (placeholder for HTTP server implementation)
func (s *Server) Start(ctx context.Context, addr string) error {
	// TODO: Implement HTTP server start logic
	return nil
}
