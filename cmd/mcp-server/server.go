package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"tushare-go/cmd/mcp-server/config"
	"tushare-go/pkg/sdk"
)

// NewServer creates a new multi-service MCP server
func NewServer(cfg *config.ServerConfig, client *sdk.Client) (*Server, error) {
	srv := &Server{
		config:   cfg,
		client:   client,
		services: make(map[string]*mcpService),
	}

	// Create services based on configuration
	for name, svcConfig := range cfg.Services {
		mcpService, err := createMCPService(name, svcConfig, client)
		if err != nil {
			return nil, fmt.Errorf("failed to create service %q: %w", name, err)
		}
		srv.services[name] = mcpService
		log.Printf("Created service '%s' on path '%s'", name, svcConfig.Path)
	}

	return srv, nil
}

// createMCPService creates a single MCP service
func createMCPService(name string, svcConfig config.ServiceConfig, client *sdk.Client) (*mcpService, error) {
	// Create MCP server
	impl := &mcpsdk.Implementation{
		Name:    "tushare-mcp-" + name,
		Version: "1.0.0",
	}
	mcpServer := mcpsdk.NewServer(impl, nil)

	// Register tools based on service categories
	if err := registerToolsForService(mcpServer, svcConfig.Categories, client); err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	return &mcpService{
		name:   name,
		config: svcConfig,
		server: mcpServer,
	}, nil
}

// Start starts the server based on transport configuration
func (s *Server) Start() error {
	switch s.config.Transport {
	case "stdio":
		return s.startStdio()
	case "http":
		return s.startHTTP()
	default:
		return fmt.Errorf("unknown transport type: %s", s.config.Transport)
	}
}

// startStdio starts the server with stdio transport (all services combined)
func (s *Server) startStdio() error {
	// Use the "all" service for stdio
	svc, ok := s.services["all"]
	if !ok {
		return fmt.Errorf("no 'all' service configured for stdio transport")
	}

	// Setup context with cancellation for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("Starting Tushare MCP Server on stdio transport...")
	log.Printf("Service: %s", svc.name)

	// Start server using stdio transport
	if err := svc.server.Run(ctx, &mcpsdk.StdioTransport{}); err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	log.Println("Server shutdown complete")
	return nil
}

// startHTTP starts the server with HTTP transport (multiple services on different paths)
func (s *Server) startHTTP() error {
	// Create HTTP router
	mux := http.NewServeMux()

	// Register each service on its path
	for name, svc := range s.services {
		if name == "all" {
			// Skip "all" service for HTTP, register individual services
			continue
		}

		// Create streamable HTTP handler for this service
		httpHandler := mcpsdk.NewStreamableHTTPHandler(func(*http.Request) *mcpsdk.Server {
			return svc.server
		}, &mcpsdk.StreamableHTTPOptions{})

		// Wrap with CORS middleware
		corsHandler := corsMiddleware(httpHandler)

		// Register the handler on the service's path
		mux.Handle(svc.config.Path, corsHandler)

		log.Printf("Registered HTTP service '%s' on path '%s'", name, svc.config.Path)
	}

	// Also register "all" service on root path if configured
	if allSvc, ok := s.services["all"]; ok {
		httpHandler := mcpsdk.NewStreamableHTTPHandler(func(*http.Request) *mcpsdk.Server {
			return allSvc.server
		}, &mcpsdk.StreamableHTTPOptions{})
		corsHandler := corsMiddleware(httpHandler)
		mux.Handle("/", corsHandler)
		log.Printf("Registered HTTP service 'all' on path '/'")
	}

	// Create HTTP server
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.httpServer = &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server in goroutine
	errChan := make(chan error, 1)
	go func() {
		log.Printf("Tushare MCP HTTP Server listening on %s", addr)
		errChan <- s.httpServer.ListenAndServe()
	}()

	// Wait for either server error or shutdown signal
	select {
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
		return nil

	case sig := <-sigChan:
		log.Printf("Received signal %v, shutting down...", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown failed: %w", err)
		}
		log.Println("HTTP server shutdown complete")
		return nil
	}
}
