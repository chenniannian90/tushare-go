# API Token 认证功能

## 概述

MCP Server 现在支持基于 API Token 的认证功能。这个功能允许你：

1. 在配置文件中预设合法的 API Token 列表
2. 用户请求时必须提供其中一个 Token
3. 验证通过后，用户的 Token 将作为 Tushare API 的 Token 使用

## 配置方式

### 1. 配置文件设置

在 `config.json` 中添加 `api_tokens` 字段：

```json
{
  "host": "0.0.0.0",
  "port": 8080,
  "transport": "http",
  "api_tokens": [
    "your-tushare-token-1",
    "your-tushare-token-2",
    "your-tushare-token-3"
  ],
  "services": {
    "stock": {
      "name": "stock",
      "path": "/stock",
      "categories": ["stock_basic", "stock_market"]
    }
  }
}
```

### 2. 环境变量方式（向后兼容）

如果不设置 `api_tokens`，系统将继续使用 `TUSHARE_TOKEN` 环境变量：

```bash
export TUSHARE_TOKEN=your-tushare-token
./bin/mcp-server -config config.json
```

## 使用方式

### HTTP 传输

#### 使用 Authorization Header（推荐）

```bash
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "stock_basic.stock_basic",
    "arguments": {}
  }'
```

#### 使用 X-API-Token Header

```bash
curl -X POST http://localhost:8080/stock \
  -H "X-API-Token: your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "stock_basic.stock_basic",
    "arguments": {}
  }'
```

### Stdio 传输

对于 stdio 传输，建议通过配置文件设置 `api_tokens`，第一个 token 将被用作默认 token。

## 工作原理

1. **Token 验证**：
   - 请求到达时，中间件检查 `Authorization` 或 `X-API-Token` header
   - 验证 token 是否在配置的 `api_tokens` 列表中
   - 如果不在列表中，返回 401 Unauthorized

2. **Token 使用**：
   - 验证通过后，token 被存储在请求上下文中
   - SDK Client 自动从上下文中提取 token
   - 该 token 被用于所有 Tushare API 调用

3. **向后兼容**：
   - 如果未配置 `api_tokens`，则允许所有请求（不进行认证）
   - 仍然支持 `TUSHARE_TOKEN` 环境变量方式

## 安全建议

1. **配置文件保护**：
   - 确保 `config.json` 文件权限设置正确（如 `chmod 600 config.json`）
   - 不要将包含真实 token 的配置文件提交到版本控制系统

2. **Token 管理**：
   - 定期轮换 API tokens
   - 为不同的用户或应用使用不同的 token
   - 在 Tushare 平台上为每个 token 设置适当的使用限制

3. **HTTPS**：
   - 在生产环境中，建议使用 HTTPS 来保护 token 在传输过程中的安全

## 测试

运行相关测试：

```bash
# 测试认证中间件
go test -v ./cmd/mcp-server/ -run TestAuth

# 测试配置加载
go test -v ./cmd/mcp-server/config/ -run TestLoadFile

# 测试动态 token 功能
go test -v ./pkg/sdk/ -run "TestWithContext|TestGetToken|TestWithToken"
```

## 故障排除

### 401 Unauthorized 错误

- 检查 token 是否在配置文件的 `api_tokens` 列表中
- 确认 header 格式正确（`Authorization: Bearer <token>` 或 `X-API-Token: <token>`）

### Token 不生效

- 确认配置文件正确加载
- 检查 `api_tokens` 是否为空数组（空数组表示不进行认证）
- 查看服务器日志获取详细错误信息

### 编译错误

- 确保 Go 版本 >= 1.24
- 运行 `go mod tidy` 更新依赖
- 重新构建：`go build -o bin/mcp-server ./cmd/mcp-server/`