# MCP Multi-Service Architecture

## Overview

The Tushare MCP server now supports running multiple MCP services on different URL paths, following the pattern from the reference CI-MCP implementation.

## Architecture

### Multiple Services

The server can host multiple MCP services, each with its own:
- **URL Path**: Different endpoints for different service categories
- **Tool Registry**: Separate tool registries per service
- **Authentication**: Individual auth configuration per service

### Default Services

The following services are configured by default:

1. **`/` (all)**: All Tushare APIs combined
   - Includes: stock, bond, futures, and all other categories
   - Best for: General use cases

2. **`/stock`**: Stock market data APIs
   - Includes: stock_basic, stock_market, stock_financial, stock_board, etc.
   - Best for: Equity-focused applications

3. **`/bond`**: Bond market data APIs
   - Includes: All bond-related APIs
   - Best for: Fixed-income applications

4. **`/futures`**: Futures market data APIs
   - Includes: All futures-related APIs
   - Best for: Derivatives trading applications

## Usage

### Stdio Transport (Default)

```bash
# Run with stdio transport (all services combined)
TUSHARE_TOKEN=your_token ./mcp-server
```

### HTTP Transport

```bash
# Run with HTTP transport on custom port
TUSHARE_TOKEN=your_token ./mcp-server -transport http -addr :8080
```

Then access services:
- http://localhost:8080/ - All APIs
- http://localhost:8080/stock - Stock APIs
- http://localhost:8080/bond - Bond APIs
- http://localhost:8080/futures - Futures APIs

### Both Transports

```bash
# Run both HTTP and stdio simultaneously
TUSHARE_TOKEN=your_token ./mcp-server -transport both -addr :8080
```

## Configuration

Services are configured in the `createDefaultServerConfig()` function in `cmd/mcp-server/main.go`.

To add a new service, modify the `Services` map:

```go
Services: map[string]ServiceConfig{
    "custom": {
        Name:        "custom",
        Path:        "/custom",
        Description: "Custom service description",
        Categories:  []string{"category1", "category2"},
        Auth: AuthConfig{
            Type:     "none",  // or "apikey"
            Required: false,
        },
    },
}
```

## MCP Client Usage

### Using with Desktop Claude (stdio)

```bash
# In your Claude Desktop config
{
  "mcpServers": {
    "tushare": {
      "command": "/path/to/mcp-server",
      "env": {
        "TUSHARE_TOKEN": "your_token_here"
      }
    }
  }
}
```

### Using with HTTP Clients

```bash
# Access stock service
curl -X POST http://localhost:8080/stock \
  -H "Content-Type: application/json" \
  -d '...'

# Access bond service
curl -X POST http://localhost:8080/bond \
  -H "Content-Type: application/json" \
  -d '...'
```

## Implementation Details

The implementation follows these key patterns from the reference CI-MCP server:

1. **Service Isolation**: Each service has its own MCP server instance
2. **Path-based Routing**: Uses `http.NewServeMux()` for routing
3. **Streamable HTTP Handlers**: Uses `mcpsdk.NewStreamableHTTPHandler()`
4. **Middleware**: CORS middleware applied to all services
5. **Graceful Shutdown**: Proper signal handling and cleanup

## Future Enhancements

- Add authentication support per service
- Add rate limiting per service
- Add service health checks
- Add service-specific configuration files
- Add dynamic service registration

## Migration from Single Service

The new implementation is backward compatible:
- Stdio mode uses the "all" service (combined tools)
- Existing integrations continue to work without changes
- HTTP mode provides new capabilities for advanced use cases
