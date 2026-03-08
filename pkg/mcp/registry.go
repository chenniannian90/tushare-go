package mcp

import (
	"context"
	"fmt"

	"tushare-go/pkg/mcp/common"
	"tushare-go/pkg/sdk"
)

// ToolRegistry manages available MCP tools
type ToolRegistry struct {
	client *sdk.Client
	tools  map[string]common.Tool
}

// NewToolRegistry creates a new tool registry
func NewToolRegistry(client *sdk.Client) *ToolRegistry {
	return &ToolRegistry{
		client: client,
		tools:  make(map[string]common.Tool),
	}
}

// RegisterTool registers a tool in the registry
func (r *ToolRegistry) RegisterTool(tool common.Tool) {
	r.tools[tool.Name] = tool
}

// GetTools returns all registered tools
func (r *ToolRegistry) GetTools() []common.Tool {
	tools := make([]common.Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// CallTool executes a tool by name
func (r *ToolRegistry) CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error) {
	_, exists := r.tools[toolName]
	if !exists {
		// Return a message instead of an error for unknown tools
		return &common.ToolResult{
			Content: []common.Content{
				{
					Type: "text",
					Text: fmt.Sprintf("Tool not found: %s", toolName),
				},
			},
		}, nil
	}

	// For now, return a simple result
	// TODO: Implement actual tool execution logic
	return &common.ToolResult{
		Content: []common.Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Tool %s called with args: %v", toolName, args),
			},
		},
	}, nil
}
