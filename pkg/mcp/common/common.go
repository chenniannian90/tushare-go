// Package common provides shared types and functions for MCP
package common

import (
	"encoding/json"
	"fmt"
)

// Tool represents an MCP tool definition
type Tool struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	APINames    []string `json:"api_names,omitempty"`
}

// ToolResult represents the result of a tool execution
type ToolResult struct {
	Content []Content `json:"content"`
}

// Content represents a piece of content in a tool result
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// APIResult formats API results into a ToolResult
func APIResult(items interface{}, apiName, resultType string) (*ToolResult, error) {
	// Convert items to JSON for better formatting
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return &ToolResult{
			Content: []Content{
				{
					Type: "text",
					Text: fmt.Sprintf("Error formatting results: %v", err),
				},
			},
		}, nil
	}

	return &ToolResult{
		Content: []Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Successfully called %s API:\n\n%s", apiName, string(jsonData)),
			},
		},
	}, nil
}

// SimpleResult formats a simple result message
func SimpleResult(message string) *ToolResult {
	return &ToolResult{
		Content: []Content{
			{
				Type: "text",
				Text: message,
			},
		},
	}
}

// ErrorResult formats an error result
func ErrorResult(err error) *ToolResult {
	return &ToolResult{
		Content: []Content{
			{
				Type: "text",
				Text: fmt.Sprintf("Error: %v", err),
			},
		},
	}
}

// ParseInput parses input arguments from map[string]interface{} to target struct
func ParseInput(args map[string]interface{}, target interface{}) error {
	// Convert args to JSON and then unmarshal into target
	jsonData, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to marshal args: %w", err)
	}

	err = json.Unmarshal(jsonData, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal args: %w", err)
	}

	return nil
}
