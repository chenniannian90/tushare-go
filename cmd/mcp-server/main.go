package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"tushare-go/pkg/sdk"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"

	// Import all tool modules
	bondtools "tushare-go/pkg/mcp/tools/bond"
	etftools "tushare-go/pkg/mcp/tools/etf"
	forextools "tushare-go/pkg/mcp/tools/forex"
	fundtools "tushare-go/pkg/mcp/tools/fund"
	futurestools "tushare-go/pkg/mcp/tools/futures"
	hk_stocktools "tushare-go/pkg/mcp/tools/hk_stock"
	indextools "tushare-go/pkg/mcp/tools/index"
	industry_tmttools "tushare-go/pkg/mcp/tools/industry_tmt"
	llm_corpustools "tushare-go/pkg/mcp/tools/llm_corpus"
	macro_businesstools "tushare-go/pkg/mcp/tools/macro_business"
	macro_economytools "tushare-go/pkg/mcp/tools/macro_economy"
	macro_interest_ratetools "tushare-go/pkg/mcp/tools/macro_interest_rate"
	macro_pricetools "tushare-go/pkg/mcp/tools/macro_price"
	optionstools "tushare-go/pkg/mcp/tools/options"
	spottools "tushare-go/pkg/mcp/tools/spot"
	stock_basictools "tushare-go/pkg/mcp/tools/stock_basic"
	stock_boardtools "tushare-go/pkg/mcp/tools/stock_board"
	stock_featuretools "tushare-go/pkg/mcp/tools/stock_feature"
	stock_financialtools "tushare-go/pkg/mcp/tools/stock_financial"
	stock_fund_flowtools "tushare-go/pkg/mcp/tools/stock_fund_flow"
	stock_margintools "tushare-go/pkg/mcp/tools/stock_margin"
	stock_markettools "tushare-go/pkg/mcp/tools/stock_market"
	stock_referencetools "tushare-go/pkg/mcp/tools/stock_reference"
	us_stocktools "tushare-go/pkg/mcp/tools/us_stock"
	wealth_fund_salestools "tushare-go/pkg/mcp/tools/wealth_fund_sales"
)

// ServiceConfig defines configuration for a single MCP service
type ServiceConfig struct {
	Name        string            `json:"name"`
	Path        string            `json:"path"`
	Description string            `json:"description"`
	Categories  []string          `json:"categories,omitempty"` // e.g., ["stock", "bond", "futures"]
	Auth        AuthConfig        `json:"auth,omitempty"`
}

// AuthConfig defines authentication configuration for a service
type AuthConfig struct {
	Type     string `json:"type"`     // "none", "apikey"
	Required bool   `json:"required"` // whether auth is required
}

// ServerConfig defines the overall server configuration
type ServerConfig struct {
	Host        string                   `json:"host"`
	Port        int                      `json:"port"`
	Transport   string                   `json:"transport"`   // "stdio", "http", "both"
	Services    map[string]ServiceConfig `json:"services"`    // named service configurations
	GlobalAuth  AuthConfig               `json:"global_auth"` // fallback auth config
}

// mcpService represents a single MCP service with its server and configuration
type mcpService struct {
	name   string
	config ServiceConfig
	server *mcpsdk.Server
}

// Server represents the multi-service MCP server
type Server struct {
	config     *ServerConfig
	client     *sdk.Client
	services   map[string]*mcpService
	httpServer *http.Server
}

func main() {
	// Parse command-line flags
	transport := flag.String("transport", "stdio", "Transport type: stdio, http, or both")
	addr := flag.String("addr", ":8080", "HTTP server address (for http/both transports)")
	flag.Parse()

	// Get Tushare token from environment
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		log.Fatal("TUSHARE_TOKEN environment variable is required")
	}

	// Create SDK client
	config, err := sdk.NewConfig(token)
	if err != nil {
		log.Fatalf("Failed to create SDK config: %v", err)
	}

	client := sdk.NewClient(config)

	// Create server configuration
	serverConfig := createDefaultServerConfig(*transport, *addr)

	// Create server
	srv, err := NewServer(serverConfig, client)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start server based on transport type
	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Server shutdown complete")
}

// createDefaultServerConfig creates a default server configuration with multiple services
func createDefaultServerConfig(transport, addr string) *ServerConfig {
	host, port := parseAddr(addr)

	return &ServerConfig{
		Host:      host,
		Port:      port,
		Transport: transport,
		Services: map[string]ServiceConfig{
			"all": {
				Name:        "all",
				Path:        "/",
				Description: "All Tushare APIs (stock, bond, futures, etc.)",
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"stock": {
				Name:        "stock",
				Path:        "/stock",
				Description: "Stock market data APIs",
				Categories:  []string{"stock_basic", "stock_market", "stock_financial", "stock_board", "stock_feature", "stock_fund_flow", "stock_margin", "stock_reference"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"bond": {
				Name:        "bond",
				Path:        "/bond",
				Description: "Bond market data APIs",
				Categories:  []string{"bond"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
			"futures": {
				Name:        "futures",
				Path:        "/futures",
				Description: "Futures market data APIs",
				Categories:  []string{"futures"},
				Auth: AuthConfig{
					Type:     "none",
					Required: false,
				},
			},
		},
		GlobalAuth: AuthConfig{
			Type:     "none",
			Required: false,
		},
	}
}

// parseAddr parses host:port address
func parseAddr(addr string) (string, int) {
	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		return parts[0], parseInt(parts[1], 8080)
	}
	return "0.0.0.0", 8080
}

// parseInt parses string to int with default
func parseInt(s string, defaultVal int) int {
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
		return defaultVal
	}
	return i
}

// NewServer creates a new multi-service MCP server
func NewServer(config *ServerConfig, client *sdk.Client) (*Server, error) {
	srv := &Server{
		config:   config,
		client:   client,
		services: make(map[string]*mcpService),
	}

	// Create services based on configuration
	for name, svcConfig := range config.Services {
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
func createMCPService(name string, svcConfig ServiceConfig, client *sdk.Client) (*mcpService, error) {
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

// registerToolsForService registers tools for a specific service based on categories
func registerToolsForService(server *mcpsdk.Server, categories []string, client *sdk.Client) error {
	// If no categories specified, register all tools
	if len(categories) == 0 {
		return registerAllTools(server, client)
	}

	// Register tools based on categories
	for _, category := range categories {
		if err := registerToolCategory(server, category, client); err != nil {
			return fmt.Errorf("failed to register %s tools: %w", category, err)
		}
	}

	return nil
}

// registerToolCategory registers all tools for a specific category
func registerToolCategory(server *mcpsdk.Server, category string, client *sdk.Client) error {
	switch category {
	case "bond":
		tools := bondtools.NewBondTools(server, client)
		tools.RegisterAll()
	case "etf":
		tools := etftools.NewEtfTools(server, client)
		tools.RegisterAll()
	case "forex":
		tools := forextools.NewForexTools(server, client)
		tools.RegisterAll()
	case "fund":
		tools := fundtools.NewFundTools(server, client)
		tools.RegisterAll()
	case "futures":
		tools := futurestools.NewFuturesTools(server, client)
		tools.RegisterAll()
	case "hk_stock":
		tools := hk_stocktools.NewHk_stockTools(server, client)
		tools.RegisterAll()
	case "index":
		tools := indextools.NewIndexTools(server, client)
		tools.RegisterAll()
	case "industry_tmt":
		tools := industry_tmttools.NewIndustry_tmtTools(server, client)
		tools.RegisterAll()
	case "llm_corpus":
		tools := llm_corpustools.NewLlm_corpusTools(server, client)
		tools.RegisterAll()
	case "macro_business":
		tools := macro_businesstools.NewMacro_businessTools(server, client)
		tools.RegisterAll()
	case "macro_economy":
		tools := macro_economytools.NewMacro_economyTools(server, client)
		tools.RegisterAll()
	case "macro_interest_rate":
		tools := macro_interest_ratetools.NewMacro_interest_rateTools(server, client)
		tools.RegisterAll()
	case "macro_price":
		tools := macro_pricetools.NewMacro_priceTools(server, client)
		tools.RegisterAll()
	case "options":
		tools := optionstools.NewOptionsTools(server, client)
		tools.RegisterAll()
	case "spot":
		tools := spottools.NewSpotTools(server, client)
		tools.RegisterAll()
	case "stock_basic":
		tools := stock_basictools.NewStock_basicTools(server, client)
		tools.RegisterAll()
	case "stock_board":
		tools := stock_boardtools.NewStock_boardTools(server, client)
		tools.RegisterAll()
	case "stock_feature":
		tools := stock_featuretools.NewStock_featureTools(server, client)
		tools.RegisterAll()
	case "stock_financial":
		tools := stock_financialtools.NewStock_financialTools(server, client)
		tools.RegisterAll()
	case "stock_fund_flow":
		tools := stock_fund_flowtools.NewStock_fund_flowTools(server, client)
		tools.RegisterAll()
	case "stock_margin":
		tools := stock_margintools.NewStock_marginTools(server, client)
		tools.RegisterAll()
	case "stock_market":
		tools := stock_markettools.NewStock_marketTools(server, client)
		tools.RegisterAll()
	case "stock_reference":
		tools := stock_referencetools.NewStock_referenceTools(server, client)
		tools.RegisterAll()
	case "us_stock":
		tools := us_stocktools.NewUs_stockTools(server, client)
		tools.RegisterAll()
	case "wealth_fund_sales":
		tools := wealth_fund_salestools.NewWealth_fund_salesTools(server, client)
		tools.RegisterAll()
	default:
		return fmt.Errorf("unknown category: %s", category)
	}

	return nil
}

// registerAllTools registers all available tools
func registerAllTools(server *mcpsdk.Server, client *sdk.Client) error {
	// Register all tool modules
	bondtools.NewBondTools(server, client).RegisterAll()
	etftools.NewEtfTools(server, client).RegisterAll()
	forextools.NewForexTools(server, client).RegisterAll()
	fundtools.NewFundTools(server, client).RegisterAll()
	futurestools.NewFuturesTools(server, client).RegisterAll()
	hk_stocktools.NewHk_stockTools(server, client).RegisterAll()
	indextools.NewIndexTools(server, client).RegisterAll()
	industry_tmttools.NewIndustry_tmtTools(server, client).RegisterAll()
	llm_corpustools.NewLlm_corpusTools(server, client).RegisterAll()
	macro_businesstools.NewMacro_businessTools(server, client).RegisterAll()
	macro_economytools.NewMacro_economyTools(server, client).RegisterAll()
	macro_interest_ratetools.NewMacro_interest_rateTools(server, client).RegisterAll()
	macro_pricetools.NewMacro_priceTools(server, client).RegisterAll()
	optionstools.NewOptionsTools(server, client).RegisterAll()
	spottools.NewSpotTools(server, client).RegisterAll()
	stock_basictools.NewStock_basicTools(server, client).RegisterAll()
	stock_boardtools.NewStock_boardTools(server, client).RegisterAll()
	stock_featuretools.NewStock_featureTools(server, client).RegisterAll()
	stock_financialtools.NewStock_financialTools(server, client).RegisterAll()
	stock_fund_flowtools.NewStock_fund_flowTools(server, client).RegisterAll()
	stock_margintools.NewStock_marginTools(server, client).RegisterAll()
	stock_markettools.NewStock_marketTools(server, client).RegisterAll()
	stock_referencetools.NewStock_referenceTools(server, client).RegisterAll()
	us_stocktools.NewUs_stockTools(server, client).RegisterAll()
	wealth_fund_salestools.NewWealth_fund_salesTools(server, client).RegisterAll()

	return nil
}

// Start starts the server based on transport configuration
func (s *Server) Start() error {
	switch s.config.Transport {
	case "stdio":
		return s.startStdio()
	case "http":
		return s.startHTTP()
	case "both":
		// Start HTTP in background, then run stdio in foreground
		go func() {
			if err := s.startHTTP(); err != nil {
				log.Printf("HTTP server error: %v", err)
			}
		}()
		return s.startStdio()
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

// corsMiddleware adds CORS headers to all responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
