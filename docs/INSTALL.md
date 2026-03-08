# tushare-go 安装与配置指南

> **本文档旨在帮助用户安装和配置 tushare-go MCP 服务器。** 您可以直接将此文件提供给 AI 助手进行自动化设置。

## 什么是 tushare-go？

tushare-go 是 [Tushare](https://tushare.pro) 金融数据 API 的 Go 语言实现，支持模型上下文协议（MCP）服务器。它提供了访问中国股票市场数据、财务报表、经济指标等的 MCP 工具。

**GitHub**: https://github.com/chenniannian90/tushare-go

## 前置要求

- **Go**: >= 1.24（从源代码构建时需要）
- **Tushare Token**: 访问 https://tushare.pro 注册获取您的 API token
- **MCP 客户端**: Claude Desktop、Cursor 或任何兼容 MCP 的客户端

## 第��步：安装 tushare-go

### 选项 A：从 GitHub 发布页面下载二进制文件（推荐）

访问 https://github.com/chenniannian90/tushare-go/releases 并下载适合您平台的二进制文件：

- `tushare-mcp-linux-amd64` / `tushare-mcp-linux-arm64`
- `tushare-mcp-darwin-amd64` / `tushare-mcp-darwin-arm64`
- `tushare-mcp-windows-amd64.exe`

```bash
# macOS ARM64（Apple Silicon）示例：
curl -L -o tushare-mcp https://github.com/chenniannian90/tushare-go/releases/latest/download/tushare-mcp-darwin-arm64
chmod +x tushare-mcp
sudo mv tushare-mcp /usr/local/bin/

# Linux amd64 示例：
curl -L -o tushare-mcp https://github.com/chenniannian90/tushare-go/releases/latest/download/tushare-mcp-linux-amd64
chmod +x tushare-mcp
sudo mv tushare-mcp /usr/local/bin/
```

在 macOS 上，您可能需要移除隔离属性：

```bash
xattr -d com.apple.quarantine tushare-mcp
```

### 选项 B：从源代码构建

```bash
git clone https://github.com/chenniannian90/tushare-go.git
cd tushare-go
make build
# 二进制文件将位于 ./bin/tushare-mcp
```

### 选项 C：通过 Go install 安装

```bash
go install github.com/chenniannian90/tushare-go/cmd/mcp-server@latest
# 二进制文件将位于 $GOPATH/bin/tushare-mcp
```

验证安装：

```bash
tushare-mcp --version
```

## 第二步：获取 Tushare API Token

1. 访问 https://tushare.pro 并注册/登��
2. 进入用户中心 → API Token
3. 复制您的 token（格式：`xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`）

**注意**：每个用户有不同的 API 访问级别。免费层级有调用频率限制。

## 第三步：创建配置文件

在工作目录中创建 `config.json` 文件：

```bash
# 复制示例配置
cp config.example.json config.json
```

编辑 `config.json` 配置您的设置：

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

### 配置选项

- **`transport`**: `"stdio"` 或 `"http"`
  - `stdio`: 用于 Claude Desktop 和本地 MCP 客户端
  - `http`: 用于基于 HTTP 的多端点访问

- **`log_level`**: `"debug"`, `"info"`, `"warn"`, `"error"`

- **`api_tokens`**: Tushare API token 数组（支持多个 token 进行负载均衡）

- **`services`**: 服务配置
  - **`name`**: 服务标识符
  - **`path`**: HTTP 路径（在 stdio 模式下忽略）
  - **`categories`**: 包含的 API 类别（空表示全部）
  - **`description`**: 服务描述

## 第四步：配置 MCP 客户端

### Claude Desktop（macOS/Windows）

创建或编辑 `~/Library/Application Support/Claude/claude_desktop_config.json`：

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

重启 Claude Desktop 以加载 MCP 服务器。

### Cursor / 其他 MCP 客户端

请参考您客户端的 MCP 配置文档。关键参数包括：
- **Command**: `tushare-mcp` 二进制文件的完整路径
- **Args**: `["-config", "/path/to/config.json"]`
- **Transport**: `"stdio"`

### HTTP 模式（高级）

对于 HTTP 模式，启动服务器：

```bash
tushare-mcp -config config.json
```

服务器将默认在 `http://localhost:8080` 上启动。访问不同的端点：

- `http://localhost:8080/stock` - 仅股票基础数据
- `http://localhost:8080/all` - 所有 API

## 第五步：验证安装

### 使用 MCP 客户端测试

与 Claude 开始对话并尝试：

```
请使用 tushare API 获取股票基本信息，查询代码 000001.SZ
```

您应该看到 Claude 使用 tushare-go 的 `get_stock_basic` 工具。

### 测试 HTTP 模式（如适用）

```bash
# 健康检查
curl http://localhost:8080/health

# 列出可用工具
curl http://localhost:8080/stock/tools
```

## 第六步：可用的 MCP 工具

连接成功后，tushare-go 提供以下 MCP 工具：

### 股票市场数据
- **`get_stock_basic`**: 获取股票基本信息
- **`get_daily`**: 获取日线行情
- **`get_weekly`**: 获取周线行情
- **`get_monthly`**: 获取月线行情
- **`get_realtime_quote`**: 获取实时行情

### 财务数据
- **`get_income`**: 获取利润表
- **`get_balancesheet`**: 获取资产负债表
- **`get_cashflow`**: 获取现金流量表
- **`get_fina_indicator`**: 获取财务指标

### 指数数据
- **`get_index_basic`**: 获取指数基本信息
- **`get_index_daily`**: 获取指数日线行情
- **`get_index_weight`**: 获取指数成分和权重

### 经济数据
- **`get_gdp`**: 获取国内生产总值
- **`get_cpi`**: 获取居民消费价格指数
- **`get_shibor`**: 获取Shibor利率

以及 200+ 更多工具，覆盖所有 Tushare API 接口。

## 高级配置

### 多个 API Token（负载均衡）

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

可用策略：`"round_robin"`、`"random"`、`"least_used"`

### 服务分类

将 API 组织为逻辑组：

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

### HTTP 服务器配置

```json
{
  "transport": "http",
  "http_host": "0.0.0.0",
  "http_port": 8080,
  "http_cors_enabled": true,
  "http_cors_allowed_origins": ["*"]
}
```

## 故障排除

### 常见问题

**"无法连接到 MCP 服务器"**
- 验证配置中的二进制文件路径是否正确
- 检查 `config.json` 语法是否有效
- 确保 Tushare token 有效且未过期

**"API 调用失败：token 超限"**
- 免费层级有调用频率限制
- 考虑升级 Tushare 账户
- 使用多个 token 进行负载均衡

**"HTTP 模式下找不到服务"**
- 检查服务 `path` 配置
- 验证类别拼写是否正确
- 尝试访问 `/all` 端点进行调试

**"无效 token 错误"**
- 验证 token 格式：32 个字符
- 检查是否有额外的空格或引号
- 如需要，在 Tushare 网站重新生成 token

### 调试模式

启用调试日志：

```json
{
  "log_level": "debug"
}
```

检查日志以获取详细错误信息。

### 测试 Tushare Token

```bash
# 手动测试 token
curl "https://api.tushare.pro/api/v1/" \
  -H "Content-Type: application/json" \
  -d '{"token":"YOUR_TOKEN","api_name":"trade_cal","params":"","fields":""}'
```

## 性能优化

### 缓存

启用响应缓存以减少 API 调用：

```json
{
  "cache_enabled": true,
  "cache_ttl": "300s"
}
```

### 连接池

```json
{
  "max_connections": 10,
  "connection_timeout": "30s"
}
```

## 安全最佳实践

1. **切勿将 `config.json` 提交**到版本控制系统
2. **使用环境变量**处理敏感数据：
   ```json
   {
     "api_tokens": ["${TUSHARE_TOKEN}"]
   }
   ```
3. **定期轮换 token**
4. **监控使用情况**以检测未授权访问
5. **在生产环境中限制 CORS 来源**：
   ```json
   {
     "http_cors_allowed_origins": ["https://yourdomain.com"]
   }
   ```

## 下一步

- **文档**: 查看 `docs/` 目录获取详细指南
- **API 参考**: 查看 `docs/MCP_TOOLS.md` 了解所有可用工具
- **示例**: 访问 `examples/` 目录查看使用示例
- **问题反馈**: 在 https://github.com/chenniannian90/tushare-go/issues 报告问题

## 支持

- **Tushare Pro**: https://tushare.pro
- **GitHub Issues**: https://github.com/chenniannian90/tushare-go/issues
- **文档**: https://github.com/chenniannian90/tushare-go/tree/main/docs

---

**版本**: 1.0.0
**最后更新**: 2026-03-08
**许可证**: MIT