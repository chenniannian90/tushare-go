# MCP Server API Key Authentication

本文档说明���何为 Tushare MCP 服务器配置和使用 API Key 认证。

## 概述

MCP 服务器支持可选的 API Key 认证机制，用于保护服务器免受未授权访问。当启用认证时，客户端必须在初始化时提供有效的 API Key。

## 配置

### 环境变量

通过以下环境变量配置 API Key 认证：

| 变量名 | 必需 | 说明 | 示例 |
|--------|------|------|------|
| `TUSHARE_TOKEN` | 是 | Tushare API token | `your-tushare-token` |
| `MCP_API_KEY` | 条件 | MCP 服务器 API Key | `your-mcp-api-key` |
| `MCP_REQUIRE_AUTH` | 否 | 是否启用认证 (`true`/`false`) | `true` |

### 配置示例

```bash
# 启用 API Key 认证
export TUSHARE_TOKEN="your-tushare-token"
export MCP_API_KEY="your-mcp-api-key"
export MCP_REQUIRE_AUTH="true"

# 不启用认证（默认）
export TUSHARE_TOKEN="your-tushare-token"
export MCP_REQUIRE_AUTH="false"  # 或不设置此变量
```

## 客户端连接

### 启用认证的连接

当服务器启用了 API Key 认证时，客户端需要在初始化请求中提供 API Key：

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "clientInfo": {
      "name": "your-client-name",
      "version": "1.0.0",
      "apiKey": "your-mcp-api-key"
    }
  }
}
```

### 未启用认证的连接

如果服务器未启用认证，客户端可以正常连接，无需提供 API Key：

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "clientInfo": {
      "name": "your-client-name",
      "version": "1.0.0"
    }
  }
}
```

## 认证失败处理

### 错误码

| 错误码 | 说明 |
|--------|------|
| -32001 | 认证失败 |
| -32600 | 无效请求 |
| -32601 | 方法不支持 |

### 错误响应示例

**API Key 错误：**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32001,
    "message": "Authentication failed: invalid API key"
  }
}
```

**缺少 API Key：**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32001,
    "message": "Authentication failed: missing apiKey in clientInfo"
  }
}
```

**缺少客户端信息：**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32001,
    "message": "Authentication failed: missing clientInfo in initialization parameters"
  }
}
```

## 安全建议

### API Key 管理

1. **生成强 API Key**：使用足够长度的随机字符串
   ```bash
   # 生成 32 字符随机 API Key
   openssl rand -hex 16
   ```

2. **安全存储**：
   - 使用环境变量或配置文件
   - 不要将 API Key 提交到版本控制系统
   - 使用 `.env` 文件（记得添加到 `.gitignore`）

3. **定期轮换**：定期更换 API Key 以提高安全性

### 服务器部署

```bash
# 使用 systemd 服务
cat > /etc/systemd/system/tushare-mcp.service <<EOF
[Unit]
Description=Tushare MCP Server
After=network.target

[Service]
Type=simple
User=tushare
Environment="TUSHARE_TOKEN=your-token"
Environment="MCP_API_KEY=your-api-key"
Environment="MCP_REQUIRE_AUTH=true"
ExecStart=/usr/local/bin/tushare-mcp-server
Restart=always

[Install]
WantedBy=multi-user.target
EOF
```

### 网络安全

1. **使用反向代理**：通过 nginx 或 Apache 提供额外的安全层
2. **启用 HTTPS**：确保通信加密
3. **限制访问**：配置防火墙规则限制访问源

## 测试

### 单元测试

运行 API Key 认证相关测试：

```bash
go test -tags=integration -v ./pkg/mcp/server/ -run TestAPIKey
```

### 示例程序

运行完整的认证示例：

```bash
go run examples/mcp_server_with_auth/main.go
```

## 故障排除

### 配置验证失败

**问题**：`TUSHARE_TOKEN is required`
- **解决方案**：确保设置了 `TUSHARE_TOKEN` 环境变量

**问题**：`MCP_API_KEY is required when MCP_REQUIRE_AUTH is true`
- **解决方案**：当启用认证时，必须设置 `MCP_API_KEY`

### 认证失败

**问题**：客户端收到认证失败错误
- **检查**：
  1. 服务器是否启用了认证（`MCP_REQUIRE_AUTH=true`）
  2. 客户端是否提供了正确的 API Key
  3. 初始化请求格式是否正确

## 最佳实践

1. **开发环境**：禁用认证以简化开发
2. **生产环境**：始终启用 API Key 认证
3. **监控**：记录认证尝试以检测潜在攻击
4. **备份**：定期备份配置和 API Key

## 相关文档

- [MCP 协议规范](https://modelcontextprotocol.io/)
- [服务器配置](../pkg/mcp/server/config.go)
- [认证测试](../pkg/mcp/server/auth_test.go)
