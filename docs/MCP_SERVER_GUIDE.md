# Tushare-Go MCP 服务器启动指南

## 🎯 概述

本项目提供完整的 Model Context Protocol (MCP) 服务器实现，支持访问 Tushare 金融数据 API。服务器采用模块化架构，支持多种传输方式和多服务部署。

## 📋 前置要求

### 必需条件
- **Go 版本**: Go 1.24+ (使用官方 MCP SDK)
- **Tushare Token**: 从 [Tushare Pro](https://tushare.pro) 获取
- **系统**: macOS/Linux/Windows

### 可选条件
- **Docker**: 用于容器化部署
- **Make**: 用于构建自动化

## 🚀 快速启动

### 1. 设置 Token

```bash
# 设置 Tushare Token (必需)
export TUSHARE_TOKEN="your-tushare-token-here"

# 或者创建 .env 文件
echo "TUSHARE_TOKEN=your-tushare-token-here" > .env
```

### 2. 启动方式选择

#### 方式 A: 直接运行 (推荐用于开发)

```bash
# 进入项目目录
cd /path/to/tushare-go

# 运行 MCP 服务器
go run cmd/mcp-server/main.go
```

#### 方式 B: 使用 Makefile (推荐用于生产)

```bash
# 构建
make build-mcp

# 运行
./bin/tushare-mcp
```

#### 方式 C: 构建认证版本

```bash
# 构建带认证的服务器
make build-mcp-auth

# 运行
./bin/mcp_server_with_auth
```

## 🔧 配置选项

### 默认配置

服务器默认使用以下配置：

```json
{
  "host": "0.0.0.0",
  "port": 8080,
  "transport": "both",
  "services": {
    "all": {
      "name": "tushare-all",
      "path": "/mcp",
      "description": "All Tushare APIs",
      "categories": []
    }
  }
}
```

### 自定义配置

创建配置文件 `config.json`:

```json
{
  "host": "localhost",
  "port": 9000,
  "transport": "http",
  "api_tokens": [
    "your-tushare-token-1",
    "your-tushare-token-2"
  ],
  "services": {
    "stock": {
      "name": "tushare-stock",
      "path": "/stock",
      "description": "Stock market data",
      "categories": ["stock_basic", "stock_market", "stock_financial"]
    },
    "bond": {
      "name": "tushare-bond",
      "path": "/bond",
      "description": "Bond market data",
      "categories": ["bond"]
    }
  }
}
```

使用自定义配置:

```bash
# 如果配置了 api_tokens，可以省略 TUSHARE_TOKEN 环境变量
go run cmd/mcp-server/main.go -config config.json

# 或者仍然使用环境变量
export TUSHARE_TOKEN="your-tushare-token"
go run cmd/mcp-server/main.go -config config.json
```

**配置字段说明**：
- `host`: 服务器监听地址
- `port`: 服务器端口
- `transport`: 传输类型 (`stdio` 或 `http`)
- `api_tokens`: 可选的 API token 列表，用于多用户认证
- `services`: 服务端点配置

## 📡 传输模式

### 1. Stdio 模式 (适合 AI 助手)

```bash
# 启动 stdio 传输
export TUSHARE_TOKEN="your-token"
go run cmd/mcp-server/main.go -transport stdio
```

**特点**:
- 通过 stdin/stdout 通信
- 适合与 AI 助手集成
- 单一服务端点

### 2. HTTP 模式 (适合 Web 应用)

```bash
# 启动 HTTP 传输
export TUSHARE_TOKEN="your-token"
go run cmd/mcp-server/main.go -transport http
```

**特点**:
- RESTful API 接口
- 支持多路径服务
- 浏览器可访问
- 健康检查端点

### 3. 双重模式 (同时支持)

```bash
# 同时启动 HTTP 和 stdio
export TUSHARE_TOKEN="your-token"
go run cmd/mcp-server/main.go -transport both
```

## 🔐 API Token 认证

### 概述

MCP 服务器现在支持基于 API Token 的认证功能，允许你：

1. 在配置文件中预设多个合法的 Tushare API token
2. 为不同用户或应用配置不同的 token
3. 用户请求时必须提供其中一个 token 进行认证
4. 认证通过后，用户的 token 将自动用于 Tushare API 调用

### 配置方式

#### 方式 1: 配置文件 (推荐)

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

启动服务器时可以省略 `TUSHARE_TOKEN` 环境变量：

```bash
go run cmd/mcp-server/main.go -config config.json
```

#### 方式 2: 环境变量 (向后兼容)

如果不设置 `api_tokens`，系统将继续使用 `TUSHARE_TOKEN` 环境变量：

```bash
export TUSHARE_TOKEN="your-tushare-token"
go run cmd/mcp-server/main.go
```

### 使用方式

#### HTTP 传输认证

**使用 Authorization Header (推荐)**：

```bash
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "stock_basic.stock_basic",
    "arguments": {}
  }'
```

**使用 X-API-Token Header**：

```bash
curl -X POST http://localhost:8080/stock \
  -H "X-API-Token: your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "stock_basic.stock_basic",
    "arguments": {}
  }'
```

#### 认证失败响应

如果 token 无效或缺失，服务器将返回 401 Unauthorized：

```json
{
  "error": "Invalid authentication token"
}
```

或

```json
{
  "error": "Missing authentication token"
}
```

### 安全建议

1. **配置文件保护**：
   ```bash
   chmod 600 config.json  # 只有所有者可读写
   ```

2. **不要提交敏感信息**：
   ```bash
   echo "config.json" >> .gitignore
   ```

3. **使用 HTTPS**：在生产环境中，建议使用 HTTPS 来保护 token 在传输过程中的安全

4. **定期轮换**：定期更换 API tokens 以提高安全性

## 🧪 测试服务器

### 健康检查

```bash
# HTTP 模式
curl http://localhost:8080/health

# 预期响应
{
  "status": "ok",
  "services": ["all"]
}
```

### 列出可用工具

```bash
# HTTP MCP 端点
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/list"
  }'
```

### 调用工具示例

```bash
# 调用股票基础信息工具
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "stock_basic.stock_basic",
      "arguments": {
        "ts_code": "000001.SZ"
      }
    }
  }'
```

## 🏗️ 架构说明

### 模块化工具结构

```
tushare-go/
├── cmd/mcp-server/          # MCP 服务器主程序
│   └── main.go
├── pkg/mcp/
│   ├── server/              # MCP 服务器实现
│   │   ├── config.go
│   │   └── http_routes.go
│   └── tools/               # API 工具模块
│       ├── bond/           # 债券 API 工具
│       ├── stock_basic/    # 股票基础 API 工具
│       ├── fund/           # 基金 API 工具
│       └── ...             # 其他模块
└── bin/                     # 构建输出
    ├── tushare-mcp
    └── mcp_server_with_auth
```

### 服务分类

支持 25 个 API 模块，共 223 个工具：

| 类别 | 模块数 | 工具数 | 说明 |
|------|--------|--------|------|
| **股票** | 7 | 102 | 股票行情、财务、资金流等 |
| **债券** | 1 | 15 | 债券行情、可转债等 |
| **期货** | 1 | 14 | 期货行情、持仓等 |
| **基金** | 2 | 9 | 公募基金、ETF 等 |
| **指数** | 1 | 19 | 指数行情、成分股等 |
| **宏观** | 4 | 10 | 宏观经济数据 |
| **其他** | 9 | 54 | 外汇、期权、现货等 |

## 🐳 Docker 部署

### 构建 Docker 镜像

```dockerfile
# Dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o tushare-mcp cmd/mcp-server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/tushare-mcp .
EXPOSE 8080
CMD ["./tushare-mcp"]
```

### 运行容器

```bash
# 构建镜像
docker build -t tushare-mcp:latest .

# 运行容器
docker run -d \
  -p 8080:8080 \
  -e TUSHARE_TOKEN="your-token" \
  --name tushare-mcp \
  tushare-mcp:latest

# 查看日志
docker logs -f tushare-mcp

# 健康检查
curl http://localhost:8080/health
```

## 🔒 安全配置

### API Key 认证

```bash
# 启用认证
export MCP_API_KEY="your-secure-api-key"
export MCP_REQUIRE_AUTH="true"

# 启动服务器
go run cmd/mcp-server/main.go
```

### 使用认证

```bash
# 带认证的请求
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-secure-api-key" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list"}'
```

## 📊 监控和日志

### 日志级别

```bash
# 设置日志级别
export LOG_LEVEL=DEBUG  # DEBUG, INFO, WARN, ERROR

# 启动服务器
go run cmd/mcp-server/main.go
```

### 性能监控

```bash
# 查看进程状态
ps aux | grep tushare-mcp

# 查看端口监听
lsof -i :8080

# 查看资源使用
top -p $(pgrep tushare-mcp)
```

## 🛠️ 故障排除

### 常见问题

#### 1. 端口被占用

```bash
# 查找占用进程
lsof -i :8080

# 更换端口
export MCP_PORT=9000
go run cmd/mcp-server/main.go
```

#### 2. Token 无效

```bash
# 验证 Token
curl -X POST https://api.tushare.pro/api/stock/basic \
  -d "token=YOUR_TOKEN&ts_code=000001.SZ"

# 重置 Token
export TUSHARE_TOKEN="new-valid-token"
```

#### 3. 工具调用失败

```bash
# 检查工具名称
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | jq '.result.tools[] | .name'
```

## 📚 高级用法

### 多服务部署

```json
{
  "transport": "http",
  "services": {
    "stock": {
      "name": "stock-service",
      "path": "/stock",
      "categories": ["stock_basic", "stock_market"]
    },
    "bond": {
      "name": "bond-service",
      "path": "/bond",
      "categories": ["bond"]
    }
  }
}
```

访问不同服务:
```bash
# 股票服务
curl http://localhost:8080/stock/health

# 债券服务
curl http://localhost:8080/bond/health
```

### 负载均衡配置

可以使用 nginx 或其他负载均衡器:

```nginx
upstream tushare_mcp {
    server localhost:8080;
    server localhost:8081;
    server localhost:8082;
}

server {
    listen 80;
    location /mcp {
        proxy_pass http://tushare_mcp;
    }
}
```

## 🎯 生产部署建议

### 1. 使用 Systemd (Linux)

创建 `/etc/systemd/system/tushare-mcp.service`:

```ini
[Unit]
Description=Tushare MCP Server
After=network.target

[Service]
Type=simple
User=tushare
WorkingDirectory=/opt/tushare-mcp
Environment="TUSHARE_TOKEN=your-token"
ExecStart=/opt/tushare-mcp/bin/tushare-mcp
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

启动服务:
```bash
sudo systemctl daemon-reload
sudo systemctl start tushare-mcp
sudo systemctl enable tushare-mcp
sudo systemctl status tushare-mcp
```

### 2. 环境变量管理

创建 `/opt/tushare-mcp/.env`:
```bash
TUSHARE_TOKEN=your-token
MCP_API_KEY=your-api-key
LOG_LEVEL=INFO
MCP_PORT=8080
```

### 3. 日志管理

```bash
# 日志轮转配置 /etc/logrotate.d/tushare-mcp
/opt/tushare-mcp/logs/*.log {
    daily
    rotate 30
    compress
    delaycompress
    notifempty
    create 0640 tushare tushare
    sharedscripts
    postrotate
        systemctl reload tushare-mcp
    endscript
}
```

## 📖 相关文档

- [快速开始指南](QUICK_START.md)
- [MCP 认证文档](MCP_AUTH.md)
- [实现总结](MCP_IMPLEMENTATION_SUMMARY.md)
- [多服务架构](MCP_MULTI_SERVICE.md)
- [SDK 兼容性](MCP_SDK_COMPATIBILITY.md)

## 🆘 获取帮助

- **问题报告**: [GitHub Issues](https://github.com/chenniannian90/tushare-go/issues)
- **文档**: 查看项目 `docs/` 目录
- **示例**: 查看 `examples/` 目录

**祝您使用愉快！** 🎉
