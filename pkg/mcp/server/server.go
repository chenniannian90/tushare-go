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

// GetTools returns the list of available tools
func (s *Server) GetTools() []common.Tool {
	return s.Registry.GetTools()
}

// CallTool executes a tool call
func (s *Server) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	return s.Registry.CallTool(ctx, toolName, args)
}
