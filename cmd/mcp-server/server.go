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
		// Determine if we need to create separate services for each category
		if len(svcConfig.Categories) == 0 || (len(svcConfig.Categories) == 1 && svcConfig.Categories[0] == name) {
			// Single category matching name - create one service with all tools
			mcpService, err := createMCPService(name, svcConfig, client, svcConfig.Categories)
			if err != nil {
				return nil, fmt.Errorf("failed to create service %q: %w", name, err)
			}
			srv.services[name] = mcpService
			log.Printf("Created service '%s' on path '%s'", name, svcConfig.Path)
		} else {
			// Multiple categories - create separate service for each category
			for _, category := range svcConfig.Categories {
				// Create unique service name by combining service name and category
				serviceName := fmt.Sprintf("%s_%s", name, category)
				servicePath := fmt.Sprintf("%s/%s", svcConfig.Path, category)

				// Create service config for this specific category
				categoryConfig := svcConfig
				categoryConfig.Path = servicePath

				mcpService, err := createMCPService(serviceName, categoryConfig, client, []string{category})
				if err != nil {
					return nil, fmt.Errorf("failed to create service %q: %w", serviceName, err)
				}
				srv.services[serviceName] = mcpService
				log.Printf("Created service '%s' on path '%s' (category: %s)", serviceName, servicePath, category)
			}
		}
	}

	return srv, nil
}

// createMCPService creates a single MCP service
func createMCPService(name string, svcConfig config.ServiceConfig, client *sdk.Client, categories []string) (*mcpService, error) {
	// Create MCP server with version information
	impl := &mcpsdk.Implementation{
		Name:    "tushare-mcp-" + name,
		Version: Version,
	}
	mcpServer := mcpsdk.NewServer(impl, nil)

	// Register tools based on provided categories
	if err := registerToolsForService(mcpServer, categories, client); err != nil {
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
		httpHandler := mcpsdk.NewStreamableHTTPHandler(func(r *http.Request) *mcpsdk.Server {
			// Extract validated token from request headers and set in context
			if token := ExtractTokenFromRequest(r); token != "" {
				// Create a context with the token for this request
				_ = sdk.WithToken(r.Context(), token)
			}
			return svc.server
		}, &mcpsdk.StreamableHTTPOptions{})

		// Wrap with authentication middleware if tokens are configured
		var finalHandler http.Handler = httpHandler
		if len(s.config.APITokens) > 0 {
			finalHandler = AuthMiddleware(s.config.APITokens, httpHandler)
		}

		// Wrap with CORS middleware
		corsHandler := corsMiddleware(finalHandler)

		// Register the handler on the service's path
		mux.Handle(svc.config.Path, corsHandler)

		log.Printf("Registered HTTP service '%s' on path '%s'", name, svc.config.Path)
	}

	// Also register "all" service on root path if configured
	if allSvc, ok := s.services["all"]; ok {
		httpHandler := mcpsdk.NewStreamableHTTPHandler(func(r *http.Request) *mcpsdk.Server {
			// Extract validated token from request headers and set in context
			if token := ExtractTokenFromRequest(r); token != "" {
				// Create a context with the token for this request
				_ = sdk.WithToken(r.Context(), token)
			}
			return allSvc.server
		}, &mcpsdk.StreamableHTTPOptions{})

		// Wrap with authentication middleware if tokens are configured
		var finalHandler http.Handler = httpHandler
		if len(s.config.APITokens) > 0 {
			finalHandler = AuthMiddleware(s.config.APITokens, httpHandler)
		}

		corsHandler := corsMiddleware(finalHandler)
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
