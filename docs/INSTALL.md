# tushare-go Installation & Configuration Guide

> **This document is designed to help users install and configure tushare-go MCP server.** Feed this file directly to your AI agent for automated setup.

## What is tushare-go?

tushare-go is a Go implementation of the [Tushare](https://tushare.pro) financial data API with Model Context Protocol (MCP) server support. It provides MCP tools for accessing Chinese stock market data, financial statements, economic indicators, and more.

**GitHub**: https://github.com/chenniannian90/tushare-go

## Prerequisites

- **Go**: >= 1.24 (for building from source)
- **Tushare Token**: Register at https://tushare.pro to get your API token
- **MCP Client**: Claude Desktop, Cursor, or any MCP-compatible client

## Step 1: Install tushare-go

### Option A: Download binary from GitHub Releases (recommended)

Go to https://github.com/chenniannian90/tushare-go/releases and download the binary for your platform:

- `tushare-mcp-linux-amd64` / `tushare-mcp-linux-arm64`
- `tushare-mcp-darwin-amd64` / `tushare-mcp-darwin-arm64`
- `tushare-mcp-windows-amd64.exe`

```bash
# Example for macOS ARM64 (Apple Silicon):
curl -L -o tushare-mcp https://github.com/chenniannian90/tushare-go/releases/latest/download/tushare-mcp-darwin-arm64
chmod +x tushare-mcp
sudo mv tushare-mcp /usr/local/bin/

# Example for Linux amd64:
curl -L -o tushare-mcp https://github.com/chenniannian90/tushare-go/releases/latest/download/tushare-mcp-linux-amd64
chmod +x tushare-mcp
sudo mv tushare-mcp /usr/local/bin/
```

On macOS, you may need to remove the quarantine attribute:

```bash
xattr -d com.apple.quarantine tushare-mcp
```

### Option B: Build from source

```bash
git clone https://github.com/chenniannian90/tushare-go.git
cd tushare-go
make build
# Binary will be at ./bin/tushare-mcp
```

### Option C: Install via Go install

```bash
go install github.com/chenniannian90/tushare-go/cmd/mcp-server@latest
# Binary will be at $GOPATH/bin/tushare-mcp
```

Verify installation:

```bash
tushare-mcp --version
```

## Step 2: Get Tushare API Token

1. Visit https://tushare.pro and register/login
2. Go to user center → API Token
3. Copy your token (format: `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`)

**Note**: Each user has different API access levels. Free tier has limited call frequency.

## Step 3: Create Configuration File

Create a `config.json` file in your working directory:

```bash
# Copy example config
cp config.example.json config.json
```

Edit `config.json` with your settings:

```json
{
  "_comment": "Tushare MCP Server 配置文件",
  "_comment_usage": "复制此文件为 config.json 并根据需要修改",

  "transport": "stdio",
  "log_level": "info",

  "api_tokens": [
    "your-tushare-token-here"
  ],

  "services": [
    {
      "name": "stock",
      "path": "/stock",
      "categories": ["stock_basic"],
      "description": "股票基础数据API"
    },
    {
      "name": "all",
      "path": "/",
      "categories": [],
      "description": "所有API接口"
    }
  ]
}
```

### Configuration Options

- **`transport`**: `"stdio"` or `"http"`
  - `stdio`: For Claude Desktop and local MCP clients
  - `http`: For HTTP-based access with multiple endpoints

- **`log_level`**: `"debug"`, `"info"`, `"warn"`, `"error"`

- **`api_tokens`**: Array of Tushare API tokens (supports multiple tokens for load balancing)

- **`services`**: Service configurations
  - **`name`**: Service identifier
  - **`path`**: HTTP path (ignored in stdio mode)
  - **`categories`**: API categories to include (empty = all)
  - **`description`**: Service description

## Step 4: Configure MCP Client

### Claude Desktop (macOS/Windows)

Create or edit `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "tushare": {
      "command": "/usr/local/bin/tushare-mcp",
      "args": ["-config", "/path/to/config.json"],
      "transport": "stdio"
    }
  }
}
```

Restart Claude Desktop to load the MCP server.

### Cursor / Other MCP Clients

Refer to your client's MCP configuration documentation. The key parameters are:
- **Command**: Full path to `tushare-mcp` binary
- **Args**: `["-config", "/path/to/config.json"]`
- **Transport**: `"stdio"`

### HTTP Mode (advanced)

For HTTP mode, start the server:

```bash
tushare-mcp -config config.json
```

The server will start on `http://localhost:8080` by default. Access different endpoints:

- `http://localhost:8080/stock` - Stock basic data only
- `http://localhost:8080/all` - All APIs

## Step 5: Verify Installation

### Test with MCP Client

Start a conversation with Claude and try:

```
请使用 tushare API 获取股票基本信息，查询代码 000001.SZ
```

You should see Claude using the `get_stock_basic` tool from tushare-go.

### Test HTTP Mode (if applicable)

```bash
# Health check
curl http://localhost:8080/health

# List available tools
curl http://localhost:8080/stock/tools
```

## Step 6: Available MCP Tools

Once connected, tushare-go provides these MCP tools:

### Stock Market Data (股票市场数据)
- **`get_stock_basic`**: 获取股票基本信息
- **`get_daily`**: 获取日线行情
- **`get_weekly`**: 获取周线行情
- **`get_monthly`**: 获取月线行情
- **`get_realtime_quote`**: 获取实时行情

### Financial Data (财务数据)
- **`get_income`**: 获取利润表
- **`get_balancesheet`**: 获取资产负债表
- **`get_cashflow`**: 获取现金流量表
- **`get_fina_indicator`**: 获取财务指标

### Index Data (指数数据)
- **`get_index_basic`**: 获取指数基本信息
- **`get_index_daily`**: 获取指数日线行情
- **`get_index_weight`**: 获取指数成分和权重

### Economic Data (经济数据)
- **`get_gdp`**: 获取国内生产总值
- **`get_cpi`**: 获取居民消费价格指数
- **`get_shibor`**: 获取Shibor利率

And 200+ more tools covering all Tushare API interfaces.

## Advanced Configuration

### Multiple API Tokens (Load Balancing)

```json
{
  "api_tokens": [
    "token1",
    "token2",
    "token3"
  ],
  "token_strategy": "round_robin"
}
```

Available strategies: `"round_robin"`, `"random"`, `"least_used"`

### Service Categories

Organize APIs into logical groups:

```json
{
  "services": [
    {
      "name": "stock_market",
      "path": "/stock/market",
      "categories": ["stock_basic", "stock_market"],
      "description": "股票市场行情"
    },
    {
      "name": "financial",
      "path": "/financial",
      "categories": ["stock_financial", "stock_fund"],
      "description": "财务数据"
    },
    {
      "name": "index",
      "path": "/index",
      "categories": ["index"],
      "description": "指数数据"
    }
  ]
}
```

### HTTP Server Configuration

```json
{
  "transport": "http",
  "http_host": "0.0.0.0",
  "http_port": 8080,
  "http_cors_enabled": true,
  "http_cors_allowed_origins": ["*"]
}
```

## Troubleshooting

### Common Issues

**"Failed to connect to MCP server"**
- Verify binary path in config is correct
- Check `config.json` syntax is valid
- Ensure Tushare token is valid and not expired

**"API call failed: token limit exceeded"**
- Free tier has call frequency limits
- Consider upgrading Tushare account
- Use multiple tokens with load balancing

**"Service not found in HTTP mode"**
- Check service `path` configuration
- Verify categories are spelled correctly
- Try accessing `/all` endpoint for debugging

**"Invalid token error"**
- Verify token format: 32 characters
- Check for extra spaces or quotes
- Regenerate token at Tushare website if needed

### Debug Mode

Enable debug logging:

```json
{
  "log_level": "debug"
}
```

Check logs for detailed error messages.

### Test Tushare Token

```bash
# Test token manually
curl "https://api.tushare.pro/api/v1/" \
  -H "Content-Type: application/json" \
  -d '{"token":"YOUR_TOKEN","api_name":"trade_cal","params":"","fields":""}'
```

## Performance Optimization

### Caching

Enable response caching to reduce API calls:

```json
{
  "cache_enabled": true,
  "cache_ttl": "300s"
}
```

### Connection Pooling

```json
{
  "max_connections": 10,
  "connection_timeout": "30s"
}
```

## Security Best Practices

1. **Never commit `config.json`** to version control
2. **Use environment variables** for sensitive data:
   ```json
   {
     "api_tokens": ["${TUSHARE_TOKEN}"]
   }
   ```
3. **Rotate tokens regularly**
4. **Monitor usage** to detect unauthorized access
5. **Restrict CORS origins** in production:
   ```json
   {
     "http_cors_allowed_origins": ["https://yourdomain.com"]
   }
   ```

## Next Steps

- **Documentation**: See `docs/` directory for detailed guides
- **API Reference**: Check `docs/MCP_TOOLS.md` for all available tools
- **Examples**: Visit `examples/` directory for usage examples
- **Issues**: Report bugs at https://github.com/chenniannian90/tushare-go/issues

## Support

- **Tushare Pro**: https://tushare.pro
- **GitHub Issues**: https://github.com/chenniannian90/tushare-go/issues
- **Documentation**: https://github.com/chenniannian90/tushare-go/tree/main/docs

---

**Version**: 1.0.0
**Last Updated**: 2026-03-08
**License**: MIT