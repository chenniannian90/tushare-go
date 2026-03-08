# Tushare Go MCP 服务器 - 快速开始指南

## 🚀 5 分钟快速启动

### 选择适合你的实现

根据你的 Go 版本和需求选择合适的 MCP 服务器实现：

---

## 🟢 推荐：HTTP 服务器 (适合所有 Go 版本)

**适用场景**: Web 应用、微服务、API 集成

### 1. 启动服务器
```bash
# 设置你的 Tushare Token
export TUSHARE_TOKEN="your-actual-tushare-token"

# 启动 MCP 服务器 (推荐)
go run cmd/mcp-server/main.go

# 或者使用 Makefile
make build-mcp
./bin/tushare-mcp
```

### 2. 测试服务器
```bash
# 健康检查
curl http://localhost:8080/health

# 列出所有工具
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'

# 调用工具 (使用 API Token 认证)
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token" \
  -H "Content-Type: application/json" \
  -d '{"name":"stock_basic.stock_basic","arguments":{"ts_code":"000001.SZ"}}'
```

### 2.1 使用 API Token 认证 (可选)

如果你想在配置文件中预设多个 API token：

```bash
# 1. 创建��置文件
cp config.example.json config.json

# 2. 编辑 config.json，添加 api_tokens 字段
# {
#   "api_tokens": [
#     "your-tushare-token-1",
#     "your-tushare-token-2"
#   ]
# }

# 3. 使用配置文件启动（可以省略 TUSHARE_TOKEN 环境变量）
go run cmd/mcp-server/main.go -config config.json
```

使用认证的 API 调用：

```bash
# 使用 Authorization header
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{"name":"stock_basic.stock_basic","arguments":{}}'

# 或使用 X-API-Token header
curl -X POST http://localhost:8080/stock \
  -H "X-API-Token: your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{"name":"stock_basic.stock_basic","arguments":{}}'
```

### 3. 在浏览器中访问
```
http://localhost:8080/health
```

---

## 🟢 Stdio 服务器 (标准 MCP 客户端)

**适用场景**: AI 助手、命令行工具

### 1. 启动服务器
```bash
export TUSHARE_TOKEN="your-actual-tushare-token"

# Stdio 模式
go run cmd/mcp-server/main.go -transport stdio

# 或者构建后运行
make build-mcp
./bin/tushare-mcp -transport stdio
```

### 2. 通过 stdin/stdout 通信
```bash
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"clientInfo":{"name":"test","version":"1.0"}}}' | \
  go run examples/stdio_mcp_server/main.go
```

---

## 🟡 官方 SDK 服务器 (需要 Go 1.24+)

**适用场景**: 最新 MCP 特性、官方支持

### 1. 检查 Go 版本
```bash
go version  # 需要 go1.24.x
```

### 2. 启动服务器
```bash
export TUSHARE_TOKEN="your-actual-tushare-token"

# 使用官方 MCP SDK
go run cmd/mcp-server/main.go

# 或者构建认证版本
make build-mcp-auth
./bin/mcp_server_with_auth
```

---

## 📊 服务器功能对比

| 功能 | HTTP 服务器 | Stdio 服务器 | SDK 服务器 |
|------|------------|-------------|-----------|
| **Go 版本要求** | 1.24+ ✅ | 1.24+ ✅ | 1.24+ ✅ |
| **Web 界面** | ✅ | ❌ | ✅ |
| **RESTful API** | ✅ | ❌ | ✅ |
| **多服务支持** | ✅ | ❌ | ✅ |
| **统一端点** | ✅ | ❌ | ✅ |
| **健康检查** | ✅ | ❌ | ✅ |
| **CORS 支持** | ✅ | ❌ | ✅ |
| **工具数量** | 223 | 223 | 223 |
| **模块数量** | 25 | 25 | 25 |
| **生产就绪** | ✅ | ✅ | ✅ |

---

## 🔧 常用配置

### 环境变量配置
```bash
# 必需配置
export TUSHARE_TOKEN="your-tushare-token"     # 你的 Tushare Token

# 可选配置
export MCP_API_KEY="your-api-key"              # MCP API Key (用于认证)
export MCP_REQUIRE_AUTH="true"                 # 是否启用认证
export LOG_LEVEL="INFO"                        # 日志级别
```

### 编译部署
```bash
# 使用 Makefile (推荐)
make build-mcp

# 或者手动编译
go build -o bin/tushare-mcp cmd/mcp-server/main.go

# 运行
./bin/tushare-mcp
```

### Docker 部署
```bash
# 构建镜像
docker build -t tushare-mcp-server .

# 运行容器
docker run -p 8080:8080 \
  -e TUSHARE_TOKEN="your-token" \
  tushare-mcp-server

# 或者使用认证版本
docker run -p 8080:8080 \
  -e TUSHARE_TOKEN="your-token" \
  -e MCP_API_KEY="your-api-key" \
  tushare-mcp-server
```

---

## 📝 使用示例

### Python 客户端
```python
import requests

# 初始化
response = requests.post("http://localhost:8080/mcp", json={
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
        "clientInfo": {
            "name": "python-client",
            "version": "1.0.0"
        }
    }
})

# 调用工具
result = requests.post("http://localhost:7878/stock/stock_basic", json={
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/call",
    "params": {
        "name": "stock_basic",
        "arguments": {
            "ts_code": "000001.SZ"
        }
    }
})

print(result.json())
```

### JavaScript/Node.js 客户端
```javascript
// 初始化
const response = await fetch('http://localhost:8080/mcp', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    jsonrpc: '2.0',
    id: 1,
    method: 'initialize',
    params: {
      clientInfo: { name: 'js-client', version: '1.0.0' }
    }
  })
});

// 调用工具
const result = await fetch('http://localhost:7878/stock/stock_basic', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    jsonrpc: '2.0',
    id: 2,
    method: 'tools/call',
    params: {
      name: 'stock_basic',
      arguments: { ts_code: '000001.SZ' }
    }
  })
});

console.log(await result.json());
```

---

## 🔍 故障排除

### 问题 1: 端口被占用
```bash
# 检查端口占用
lsof -i :8080

# 更换端口
# 修改代码中的 addr := ":8080" 为其他端口
```

### 问题 2: 认证失败
```bash
# 检查 API Key 是否正确
curl -H "Authorization: Bearer your-api-key" \
  http://localhost:8080/mcp
```

### 问题 3: 工具调用失败
```bash
# 检查 Tushare Token 是否有效 (使用 MCP 协议)
curl -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_TOKEN" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {
        "name": "test-client",
        "version": "1.0.0"
      }
    }
  }'
```

---

## 📚 更多文档

- [MCP 服务器完整指南](MCP_SERVER_GUIDE.md) - 详细的服务器配置和部署指南
- [兼容性指南](MCP_SDK_COMPATIBILITY.md) - SDK 版本兼容性说明
- [API 认证文档](MCP_AUTH.md) - 认证配置说明
- [实现总结](MCP_IMPLEMENTATION_SUMMARY.md) - 技术实现细节
- [多服务架构](MCP_MULTI_SERVICE.md) - 多服务部署指南

---

## 🎯 推荐选择

### 快速开始
```bash
# 设置 Token
export TUSHARE_TOKEN="your-token"

# 运行服务器
go run cmd/mcp-server/main.go
```

### 生产部署
```bash
# Docker
docker run -p 8080:8080 -e TUSHARE_TOKEN="your-token" tushare-mcp-server

# 直接运行
./tushare-mcp-server
```

### 开发调试
```bash
# 查看详细日志
LOG_LEVEL=DEBUG go run cmd/mcp-server/main.go

# 使用特定传输模式
go run cmd/mcp-server/main.go -transport http
go run cmd/mcp-server/main.go -transport stdio
go run cmd/mcp-server/main.go -transport both
```

---

## ✨ 下一步

1. **获取 Tushare Token**: 访问 [Tushare Pro](https://tushare.pro)
2. **选择实现**: 根据你的需求选择合适的服务器
3. **启动服务器**: 按照上述步骤启动
4. **调用 API**: 使用示例代码进行测试
5. **集成应用**: 将 MCP 服务器集成到你的应用中

**祝你使用愉快！** 🎉