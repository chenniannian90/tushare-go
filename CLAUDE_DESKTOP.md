# Claude 桌面版集成指南

本指南说明如何将 Tushare MCP 服务器与 Claude 桌面版集成。

## 前提条件

1. **Claude 桌面版**：从 Anthropic 安装 Claude 桌面版
2. **Tushare Pro 账户**：从 [tushare.pro](https://tushare.pro) 获取您的 API token
3. **已构建的 MCP 服务器**：使用 `make build-mcp` 构建服务器

## 构建 MCP 服务器

```bash
# 构建 MCP 服务器
make build-mcp

# 这将创建 bin/tushare-mcp
```

## 配置

### 步骤 1：设置您的 Tushare Token

将您的 Tushare API token 设置为环境变量：

```bash
export TUSHARE_TOKEN=your_token_here
```

为了持久化配置，将其添加到您的 shell 配置文件（`~/.zshrc` 或 `~/.bashrc`）：

```bash
echo 'export TUSHARE_TOKEN=your_token_here' >> ~/.zshrc
source ~/.zshrc
```

### 步骤 2：配置 Claude 桌面版

Claude 桌面版使用配置文件来发现和连接到 MCP 服务器。

#### macOS 配置

创建或编辑 Claude 桌面版配置文件：

**位置**：`~/Library/Application Support/Claude/claude_desktop_config.json`

**配置**：

```json
{
  "mcpServers": {
    "tushare": {
      "command": "/Users/mac-new/go/src/github.com/chenniannian90/tushare-go/bin/tushare-mcp",
      "env": {
        "TUSHARE_TOKEN": "${TUSHARE_TOKEN}"
      }
    }
  }
}
```

#### Windows 配置

**位置**：`%APPDATA%\Claude\claude_desktop_config.json`

**配置**：

```json
{
  "mcpServers": {
    "tushare": {
      "command": "C:\\path\\to\\tushare-go\\bin\\tushare-mcp.exe",
      "env": {
        "TUSHARE_TOKEN": "%TUSHARE_TOKEN%"
      }
    }
  }
}
```

#### Linux 配置

**位置**：`~/.config/Claude/claude_desktop_config.json`

**配置**：

```json
{
  "mcpServers": {
    "tushare": {
      "command": "/path/to/tushare-go/bin/tushare-mcp",
      "env": {
        "TUSHARE_TOKEN": "$TUSHARE_TOKEN"
      }
    }
  }
}
```

### 步骤 3：重启 Claude 桌面版

更新配置文件后，重启 Claude 桌面版以加载 MCP 服务器。

## 验证安装

### 检查 MCP 服务器状态

1. 打开 Claude 桌面版
2. 在连接状态中查找 Tushare 服务器
3. 检查 Claude 桌面版日志中的任何错误消息

### 测试 MCP 服务器

您可以独立测试 MCP 服务器：

```bash
# 启动服务器
./bin/tushare-mcp

# 服务器将记录可用工具并等待 JSON-RPC 消息
```

发送测试消息（通过 stdin）：

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/list",
  "params": {}
}
```

预期响应：

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "tools": [
      {
        "name": "stock_basic",
        "description": "获取股票基本信息..."
      },
      ...
    ]
  }
}
```

## 可用工具

Tushare MCP 服务器向 Claude 暴露以下工具：

### 市场数据工具

- **stock_basic**：获取基本股票信息（代码、名称、行业等）
- **daily**：获取日度 OHLCV 市场数据
- **weekly**：获取周度市场数据
- **monthly**：获取月度市场数据
- **pro_bar**：综合市场数据接口
- **daily_basic**：获取日度基本面指标（PE、PB、PS、市值）
- **trade_cal**：获取交易日历信息

### 指数工具

- **index_basic**：获取指数基本信息
- **index_daily**：获取指数日度市场数据

### 财务数据工具

- **income**：获取利润表数据
- **balancesheet**：获取资产负债表数据
- **fina_indicator**：获取财务指标

### 其他工具

- **moneyflow**：获取资金流向数据
- **dividend**：获取分红信息
- **top10_holders**：获取前十大股东
- **holder_number**：获取股东统计数据
- **concept**：获取概念板块分类
- **limit_list**：获取涨跌停股票列表

## 使用示例

### 示例 1：获取股票信息

**用户消息**："告诉我关于平安银行的情况"

**Claude 工具调用**：
```json
{
  "name": "stock_basic",
  "arguments": {
    "ts_code": "000001.SZ"
  }
}
```

**Claude 响应**："找到1只股票：000001.SZ (000001)：平安银行 - 银行"

### 示例 2：获取日度市场数据

**用户消息**："平安银行的最新交易数据是什么？"

**Claude 工具调用**：
```json
{
  "name": "daily",
  "arguments": {
    "ts_code": "000001.SZ",
    "limit": "5"
  }
}
```

### 示例 3：获取交易日历

**用户消息**："今天股票市场开盘吗？"

**Claude 工具调用**：
```json
{
  "name": "trade_cal",
  "arguments": {
    "exchange": "SSE",
    "start_date": "20240101",
    "end_date": "20240131"
  }
}
```

## 故障排除

### 服务器无法启动

**症状**：MCP 服务器未出现在 Claude 桌面版中

**解决方案**：
1. 检查配置文件路径是否正确
2. 验证 `command` 路径指向构建的二进制文件
3. 检查 Claude 桌面版日志中的错误消息
4. 确保 `TUSHARE_TOKEN` 环境变量已设置

### API 错误

**症状**：工具返回 API 错误，如"权限不足"

**解决方案**：
1. 验证您的 Tushare API token 是否有效
2. 检查您的账户是否有足够权限
3. 确保您未超过速率限制
4. 尝试直接访问 API

### 连接问题

**症状**：服务器已连接但工具失败

**解决方案**：
1. 独立测试 MCP 服务器（见上文）
2. 检查到 api.tushare.pro 的网络连接
3. 验证防火墙设置
4. 检查 Claude 桌面版日志中的详细错误消息

## 高级配置

### 自定义服务器端点

如果需要使用自定义 API 端点：

```bash
export TUSHARE_ENDPOINT=https://custom-endpoint.com
```

### 超时配置

通过修改 SDK 配置调整超时设置：

```go
config.HTTPClient.Timeout = 60 * time.Second
```

### 调试模式

通过以下设置启用调试日志：

```bash
export DEBUG=1
./bin/tushare-mcp
```

## 速率限制

请注意 Tushare API 速率限制：
- **免费账户**：200次请求/分钟
- **专业账户**：根据订阅更高的限制

Claude 桌面版可能会在对话中进行多次工具调用。监控您的使用情况以避免速率限制。

## 安全考虑

1. **Token 安全**：永远不要将您的 Tushare token 提交到版本控制
2. **环境变量**：使用环境变量存储敏感数据
3. **文件权限**：确保配置文件具有适当权限（chmod 600）
4. **网络**：服务器向 api.tushare.pro 发起 HTTPS 请求

## 支持

如有问题或疑问：
- [Tushare Pro 文档](https://tushare.pro/document/2)
- [GitHub Issues](https://github.com/chenniannian90/tushare-go/issues)
- [MCP 规范](https://modelcontextprotocol.io)

## 后续步骤

1. ✅ 构建 MCP 服务器：`make build-mcp`
2. ✅ 配置 Claude 桌面版
3. ✅ 重启 Claude 桌面版
4. ✅ 开始在 Claude 对话中使用 Tushare 工具！

享受直接在 Claude 桌面版中无缝访问中国金融数据！🚀
