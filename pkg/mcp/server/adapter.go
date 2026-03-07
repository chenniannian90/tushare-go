// Package server provides MCP server adapters for integrating with official MCP SDK
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/chenniannian90/tushare-go/pkg/mcp/common"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// ToolRegistryInterface defines the interface we need from ToolRegistry
type ToolRegistryInterface interface {
	ListTools() []common.Tool
	CallTool(ctx context.Context, toolName string, args map[string]interface{}) (*common.ToolResult, error)
}

// ToolAdapter bridges the existing ToolRegistry with the official MCP SDK
type ToolAdapter struct {
	registry ToolRegistryInterface
}

// NewToolAdapter creates a new tool adapter
func NewToolAdapter(registry ToolRegistryInterface) *ToolAdapter {
	return &ToolAdapter{
		registry: registry,
	}
}

// RegisterTools registers all tools from the registry with the MCP server
func (a *ToolAdapter) RegisterTools(srv *mcpsdk.Server) error {
	tools := a.registry.ListTools()

	log.Printf("Registering %d tools with MCP server", len(tools))

	for _, tool := range tools {
		if err := a.registerTool(srv, tool); err != nil {
			return fmt.Errorf("failed to register tool %s: %w", tool.Name, err)
		}
	}

	return nil
}

// registerTool registers a single tool with the MCP server
func (a *ToolAdapter) registerTool(srv *mcpsdk.Server, tool common.Tool) error {
	// Create MCP tool definition with input schema
	// Using "any" type for flexible input parameters
	mcpTool := &mcpsdk.Tool{
		Name:        tool.Name,
		Description: tool.Description,
		InputSchema: map[string]interface{}{
			"type": "object",
		},
	}

	// Create handler function using Server.AddTool directly
	handler := a.createToolHandler(tool)

	// Add tool to server using the lower-level API
	srv.AddTool(mcpTool, handler)

	return nil
}

// createToolHandler creates a handler function for a tool
// Using the mcpsdk.ToolHandler interface (lower-level API)
func (a *ToolAdapter) createToolHandler(tool common.Tool) mcpsdk.ToolHandler {
	return func(ctx context.Context, req *mcpsdk.CallToolRequest) (*mcpsdk.CallToolResult, error) {
		log.Printf("Calling tool: %s", tool.Name)

		// Extract arguments from request
		var args map[string]interface{}
		if req.Params != nil && len(req.Params.Arguments) > 0 {
			// Parse JSON RawMessage
			if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
				log.Printf("Failed to unmarshal arguments: %v", err)
				return &mcpsdk.CallToolResult{
					Content: []mcpsdk.Content{
						&mcpsdk.TextContent{
							Text: fmt.Sprintf("Error: failed to parse arguments: %v", err),
						},
					},
					IsError: true,
				}, nil
			}
		}

		log.Printf("Tool %s args: %v", tool.Name, args)

		// Call the tool using existing registry
		result, err := a.registry.CallTool(ctx, tool.Name, args)
		if err != nil {
			log.Printf("Tool %s failed: %v", tool.Name, err)
			return &mcpsdk.CallToolResult{
				Content: []mcpsdk.Content{
					&mcpsdk.TextContent{
						Text: fmt.Sprintf("Error: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		// Convert result to MCP format
		if result == nil || len(result.Content) == 0 {
			return &mcpsdk.CallToolResult{
				Content: []mcpsdk.Content{
					&mcpsdk.TextContent{
						Text: "No results returned",
					},
				},
			}, nil
		}

		// Convert content items
		content := make([]mcpsdk.Content, 0, len(result.Content))
		for _, c := range result.Content {
			content = append(content, &mcpsdk.TextContent{
				Text: c.Text,
			})
		}

		return &mcpsdk.CallToolResult{
			Content: content,
		}, nil
	}
}