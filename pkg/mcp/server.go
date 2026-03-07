package mcp

import (
	"context"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

// Server represents the MCP server
type Server struct {
	Client  *sdk.Client
	Registry *ToolRegistry
}

// NewServer creates a new MCP server
func NewServer(client *sdk.Client) *Server {
	return &Server{
		Client:   client,
		Registry: NewToolRegistry(),
	}
}

// Tool represents an MCP tool definition
type Tool struct {
	Name        string
	Description string
}

// ToolResult represents the result of a tool call
type ToolResult struct {
	Content []Content
}

// Content represents MCP content
type Content struct {
	Type string
	Text string
}

// GetTools returns the list of available tools
func (s *Server) GetTools() []Tool {
	return s.Registry.GetTools()
}

// CallTool executes a tool call
func (s *Server) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*ToolResult, error) {
	return s.Registry.CallTool(ctx, s.Client, toolName, args)
}
