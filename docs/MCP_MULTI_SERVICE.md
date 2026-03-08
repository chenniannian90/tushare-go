# MCP Multi-Service Architecture

## Overview

The Tushare MCP server supports running multiple MCP services on different URL paths with fine-grained tool isolation. Each service endpoint contains only the tools for its specific category, providing better organization and flexibility.

## Architecture

### Category-Based Service Creation

The server automatically creates service instances based on the `categories` configuration:

1. **Empty categories**: Creates one service with **all tools**
   - Used for the "all" service in stdio mode

2. **Single category matching name**: Creates one service with **that category's tools**
   - Example: `{"name": "bond", "categories": ["bond"]}` → `/bond` with bond tools only

3. **Multiple categories**: Creates **separate service instances** for each category
   - Example: `{"name": "stock", "categories": ["stock_basic", "stock_market"]}`
   - Creates `/stock/stock_basic` (stock_basic tools only)
   - Creates `/stock/stock_market` (stock_market tools only)

### Service Isolation

Each service endpoint has:
- **Unique URL Path**: Different endpoints for different tool categories
- **Isolated Tool Registry**: Each path contains only its specific category's tools
- **Independent Authentication**: Individual auth configuration per service

### Default Services

Based on `config.example.json`, the following services are configured:

#### Single-Category Services (Simple paths)
- **`/bond`**: Bond market data APIs (bond tools only)
- **`/futures`**: Futures market data APIs (futures tools only)
- **`/fund`**: Fund market data APIs (fund tools only)
- **`/index`**: Index data APIs (index tools only)
- **`/options`**: Options data APIs (options tools only)
- **`/forex`**: Forex data APIs (forex tools only)
- **`/hk_stock`**: Hong Kong stock APIs (hk_stock tools only)
- **`/us_stock`**: US stock APIs (us_stock tools only)
- **`/etf`**: ETF data APIs (etf tools only)
- **`/spot`**: Spot data APIs (spot tools only)

#### Multi-Category Services (Hierarchical paths)
- **`/stock/stock_basic`**: Basic stock data APIs (stock_basic tools only)
- **`/stock/stock_market`**: Stock market data APIs (stock_market tools only)
- **`/stock/stock_financial`**: Stock financial data APIs (stock_financial tools only)
- **`/stock/stock_board`**: Stock board data APIs (stock_board tools only)
- **`/stock/stock_feature`**: Stock feature APIs (stock_feature tools only)
- **`/stock/stock_fund_flow`**: Stock fund flow APIs (stock_fund_flow tools only)
- **`/stock/stock_margin`**: Stock margin APIs (stock_margin tools only)
- **`/stock/stock_reference`**: Stock reference APIs (stock_reference tools only)

- **`/macro/macro_economy`**: Macro economy APIs (macro_economy tools only)
- **`/macro/macro_business`**: Macro business APIs (macro_business tools only)
- **`/macro/macro_price`**: Macro price APIs (macro_price tools only)
- **`/macro/macro_interest_rate`**: Macro interest rate APIs (macro_interest_rate tools only)

## Usage

### Stdio Transport (Default)

```bash
# Run with stdio transport (uses "all" service with all tools)
TUSHARE_TOKEN=your_token ./mcp-server
```

### HTTP Transport

```bash
# Run with HTTP transport on custom port
TUSHARE_TOKEN=your_token ./mcp-server -transport http -addr :8080
```

#### Accessing Single-Category Services

```bash
# Access bond service (bond tools only)
curl -X POST http://localhost:8080/bond \
  -H "Content-Type: application/json" \
  -d '...'

# Access futures service (futures tools only)
curl -X POST http://localhost:8080/futures \
  -H "Content-Type: application/json" \
  -d '...'
```

#### Accessing Multi-Category Services

```bash
# Access stock basic data (stock_basic tools only)
curl -X POST http://localhost:8080/stock/stock_basic \
  -H "Content-Type: application/json" \
  -d '...'

# Access stock market data (stock_market tools only)
curl -X POST http://localhost:8080/stock/stock_market \
  -H "Content-Type: application/json" \
  -d '...'

# Access macro economy data (macro_economy tools only)
curl -X POST http://localhost:8080/macro/macro_economy \
  -H "Content-Type: application/json" \
  -d '...'
```

## Configuration

Services are configured via JSON config file (default: `config.json`).

### Configuration File Structure

```json
{
  "host": "0.0.0.0",
  "port": 8080,
  "transport": "http",
  "services": {
    "service_name": {
      "name": "service_name",
      "path": "/base_path",
      "description": "Service description",
      "categories": ["category1", "category2"],
      "auth": {
        "type": "none",
        "required": false
      }
    }
  },
  "global_auth": {
    "type": "none",
    "required": false
  }
}
```

### Category Behavior Rules

#### Rule 1: Empty Categories → All Tools
```json
{
  "name": "all",
  "path": "/",
  "categories": []
}
```
**Result**: Creates 1 service at `/` with **all tools**

#### Rule 2: Single Category Matching Name → Specific Tools
```json
{
  "name": "bond",
  "path": "/bond",
  "categories": ["bond"]
}
```
**Result**: Creates 1 service at `/bond` with **bond tools only**

#### Rule 3: Multiple Categories → Separate Services per Category
```json
{
  "name": "stock",
  "path": "/stock",
  "categories": ["stock_basic", "stock_market"]
}
```
**Result**: Creates 2 services:
- `/stock/stock_basic` (stock_basic tools only)
- `/stock/stock_market` (stock_market tools only)

### Adding a New Service

Edit `config.json` (copy from `config.example.json`):

```json
{
  "services": {
    "my_service": {
      "name": "my_service",
      "path": "/my_service",
      "description": "My custom service",
      "categories": ["category1", "category2"],
      "auth": {
        "type": "none",
        "required": false
      }
    }
  }
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

For HTTP clients, you need to specify the exact category path:

```bash
# Access specific stock category
curl -X POST http://localhost:8080/stock/stock_basic \
  -H "Content-Type: application/json" \
  -d '...'

# Access bond service
curl -X POST http://localhost:8080/bond \
  -H "Content-Type: application/json" \
  -d '...'

# Access specific macro category
curl -X POST http://localhost:8080/macro/macro_economy \
  -H "Content-Type: application/json" \
  -d '...'
```

## Implementation Details

The implementation follows these key patterns:

1. **Dynamic Service Creation**: Service instances are created based on categories configuration
   - See `cmd/mcp-server/server.go:NewServer()`
   - Empty/single categories → one service instance
   - Multiple categories → one service instance per category

2. **Tool Isolation**: Each service registers only its specific category's tools
   - See `cmd/mcp-server/tools.go:registerToolsForService()`
   - Empty categories → all tools registered
   - Specific categories → only matching tools registered

3. **Path-based Routing**: Uses `http.NewServeMux()` for routing
   - Each service instance registered to its unique path
   - Multi-category services use hierarchical paths (`/base/category`)

4. **Streamable HTTP Handlers**: Uses `mcpsdk.NewStreamableHTTPHandler()`
   - Each service gets its own HTTP handler
   - Handlers are independent and isolated

5. **Middleware**: CORS middleware applied to all services
   - See `cmd/mcp-server/middleware.go:corsMiddleware()`

6. **Graceful Shutdown**: Proper signal handling and cleanup
   - Supports both stdio and HTTP transports
   - Clean shutdown on SIGINT/SIGTERM

## Advantages of Category-Based Architecture

1. **Fine-Grained Access Control**: Each endpoint contains only specific tools
   - Clients can choose exactly which tool categories they need
   - Reduced API surface for improved security

2. **Better Organization**: Clear hierarchical structure
   - Related tools grouped under logical paths
   - Easy to understand which tools are available where

3. **Independent Scaling**: Each service can be scaled independently
   - High-traffic categories can be scaled separately
   - Resource allocation based on usage patterns

4. **Flexible Deployment**: Deploy only needed service categories
   - Reduced resource usage for specialized deployments
   - Faster startup times

5. **Clear Tool Boundaries**: No tool overlap between endpoints
   - Predictable tool availability per path
   - Easier debugging and monitoring

## Future Enhancements

- Add authentication support per service/category
- Add rate limiting per service/category
- Add service health checks per endpoint
- Add metrics collection per service/category
- Add dynamic service registration without restart

## Migration from Single Service

The new implementation is backward compatible:

### Stdio Mode (No Changes)
- Uses the "all" service (all tools combined)
- Existing integrations continue to work without changes
- No migration needed

### HTTP Mode (Breaking Changes)
If you were using HTTP paths from the old implementation:

**Old behavior:**
- `/stock` → All stock-related tools (stock_basic, stock_market, etc.)

**New behavior:**
- `/stock/stock_basic` → Only stock_basic tools
- `/stock/stock_market` → Only stock_market tools
- etc.

**Migration steps:**
1. Update your HTTP client URLs to use the new hierarchical paths
2. Or use multiple endpoints if you need tools from different categories
3. See the "Default Services" section above for available paths

### Example Migration

**Before:**
```bash
curl -X POST http://localhost:8080/stock -d '{"tool": "stock_basic"}'
```

**After:**
```bash
curl -X POST http://localhost:8080/stock/stock_basic -d '{"tool": "stock_basic"}'
```
