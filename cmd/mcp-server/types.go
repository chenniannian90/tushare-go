package main

import (
	"net/http"

	"tushare-go/cmd/mcp-server/config"
	"tushare-go/pkg/sdk"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// mcpService represents a single MCP service with its server and configuration
type mcpService struct {
	name   string
	config config.ServiceConfig
	server *mcpsdk.Server
}

// Server represents the multi-service MCP server
type Server struct {
	config     *config.ServerConfig
	client     *sdk.Client
	services   map[string]*mcpService
	httpServer *http.Server
}
