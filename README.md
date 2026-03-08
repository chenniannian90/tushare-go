# Tushare Go SDK + MCP 服务器

一个生产就绪的 Go SDK，用于 [Tushare Pro](https://tushare.pro)（拥有200+ API的中国金融数据平台），以及一个 MCP（模型上下文协议）服务器，将这些API暴露给Claude等AI代理。

## 功能特性

- ✅ 针对200+ Tushare Pro REST API的类型安全Go SDK
- ✅ 代码生成器，可从JSON规范自动生成API封装
- ✅ MCP服务器，用于Claude桌面版集成
- ✅ 全面的测试覆盖率（≥80%）
- ✅ 零外部依赖（除MCP SDK外）
- ✅ 生产就绪，包含26个完整实现的工具模块
- ✅ **链式调用客户端，统一的API访问入口**
- ✅ **优化后的代码结构，简洁高效**
- ✅ **API Token 认证，支持多用户访问控制**
- ✅ **Token 负载均衡，支持多 Token 自动分配**

## 安装

```bash
# 本地安装
git clone https://github.com/chenniannian90/tushare-go.git
cd tushare-go

# 或作为模块使用 (需要 Go 1.24+)
go mod init your-module
echo "require tushare-go v1.0.0" >> go.mod
echo "replace tushare-go => /path/to/tushare-go" >> go.mod
```

## 快速开始

### SDK使用

#### 方式1：直接调用（基础方式）

```go
package main

import (
    "context"
    "fmt"
    "tushare-go/pkg/sdk"
    stockbasic "tushare-go/pkg/sdk/api/stock/stock_basic"
)

func main() {
    config, _ := sdk.NewConfig("your-tushare-token")
    client := sdk.NewClient(config)

    req := &stockbasic.StockBasicRequest{
        TsCode: "000001.SZ",
    }

    stocks, err := stockbasic.StockBasic(context.Background(), client, req)
    if err != nil {
        panic(err)
    }

    fmt.Println(stocks)
}
```

#### 方式2：链式调用（推荐）⭐

```go
package main

import (
    "context"
    "fmt"
    "tushare-go/pkg/sdk"
    "tushare-go/pkg/sdk/apis"
)

func main() {
    config, _ := sdk.NewConfig("your-tushare-token")
    client := sdk.NewClient(config)

    // 创建链式调用客户端
    tushareClient := apis.NewTushareClient(client)
    ctx := context.Background()

    // 使用链式调用访问不同的 API
    // 获取股票基本信息
    stocks, err := tushareClient.StockBasic(ctx, &stockbasic.StockBasicRequest{
        TsCode: "000001.SZ",
    })
    if err != nil {
        panic(err)
    }

    // 获取日线数据
    daily, err := tushareClient.Daily(ctx, &stockmarket.DailyRequest{
        TsCode: "000001.SZ",
    })
    if err != nil {
        panic(err)
    }

    // 获取板块数据
    topList, err := tushareClient.TopList(ctx, &stockboard.TopListRequest{})
    if err != nil {
        panic(err)
    }

    fmt.Println(stocks, daily, topList)
}
```

**链式调用的优势：**
- ✅ 统一的入口点，所有 API 通过一个客户端访问
- ✅ 类型安全，完整的 IDE 自动提示
- ✅ 无需手动导入多个 API 包
- ✅ 更好的代码组织和可维护性

**可用的 API 接口：**
- `Stock` - 股票相关 API（包含 Basic、Board、Market、Financial 等子接口）
- `Bond` - 债券相关 API
- `ETF` - ETF 相关 API
- `Fund` - 基金相关 API
- `Futures` - 期货相关 API
- `Forex` - 外汇相关 API
- `HKStock` - 港股相关 API
- `Index` - 指数相关 API
- `Industry` - 行业相关 API
- `LLMCorpus` - LLM 语料相关 API
- `Options` - 期权相关 API
- `Spot` - 现货相关 API
- `USStock` - 美股相关 API
- `Wealth` - 理财相关 API

### MCP服务器与Claude桌面版

```bash
# 构建MCP服务器
make build-mcp

# 设置您的token
export TUSHARE_TOKEN=your_token_here

# 启动服务器
./bin/tushare-mcp
```

**完整的Claude桌面版集成说明请参见 [CLAUDE_DESKTOP.md](CLAUDE_DESKTOP.md)。**

## 示例代码

项目提供了丰富的示例代码，展示不同的使用方式：

- `cmd/examples/sdk_usage/` - SDK 使用方式对比
- `cmd/examples/chain_client/` - 链式调用客户端示例
- `cmd/examples/daily/` - 获取日线数据
- `cmd/examples/stock_basic/` - 获取股票基本信息
- `cmd/examples/boards/` - 获取板块数据
- `cmd/examples/fund/` - 获取基金数据
- `cmd/examples/futures/` - 获取期货数据

运行示例：
```bash
# 设置 token
export TUSHARE_TOKEN=your_token_here

# 运行链式调用示例
go run cmd/examples/chain_client/main.go

# 运行其他示例
go run cmd/examples/daily/main.go
```

## 开发

```bash
# 运行测试
make test

# 构建 MCP 服务器（自动获取 git tag 版本）
make build-mcp

# 查看构建版本信息
make version

# 构建代码生成器
make build-gen

# 生成 API 代码
make gen

# 生成 MCP 工具
go run cmd/gen-mcp-tools/main.go -optimized
```

**版本管理**：
- 项目使用 Git tag 作为版本号来源
- 构建时自动注入版本信息到二进制文件
- 使用 `--version` 参数查看详细版本信息
- 支持 `VERSION` 环境变量自定义版本

📖 **详细文档**: [版本管理指南](docs/VERSION_MANAGEMENT.md)

## MCP 服务器使用

### 快速启动

#### 方式1：使用命令行参数

```bash
# 设置 Token
export TUSHARE_TOKEN="your-tushare-token"

# 启动 MCP 服务器（stdio模式，用于Claude桌面版）
go run cmd/mcp-server/main.go

# 启动 HTTP 服务器（用于HTTP客户端）
go run cmd/mcp-server/main.go -transport http -addr :8080

# 或者使用构建好的版本
make build-mcp
./bin/tushare-mcp

# 查看版本信息
./bin/tushare-mcp --version
```

#### 方式2：使用配置文件

```bash
# 1. 复制示例配置文件
cp config.example.json config.json

# 2. 根据需要修改配置文件
# 编辑 config.json，设置 host、port、transport 等参数

# 3. (可选) 在配置文件中设置 API tokens
# 在 config.json 中添加 "api_tokens" 字段：
# {
#   "api_tokens": [
#     "your-tushare-token-1",
#     "your-tushare-token-2"
#   ]
# }

# 4. 使用配置文件启动
# 如果配置了 api_tokens，可以省略 TUSHARE_TOKEN 环境变量
go run cmd/mcp-server/main.go -config config.json
```

**配置文件说明**：
- `host`: HTTP服务器监听地址（默认：0.0.0.0）
- `port`: HTTP服务器端口（默认：8080）
- `transport`: 传输类型，可选 "stdio" 或 "http"
- `api_tokens`: 合法的API token列表（可选，用于认证）
- `services`: 服务配置，可以定义多个服务端点
  - `all`: 所有API集合（推荐用于stdio模式）
  - `stock`: 股票市场数据API
  - `bond`: 债券市场数据API
  - `futures`: 期货市场数据API
  - 等等...
- `global_auth`: 全局认证配置

### 传输模式

```bash
# Stdio 模式 (用于 AI 助手，如Claude桌面版)
go run cmd/mcp-server/main.go -transport stdio

# HTTP 模式 (用于HTTP客户端)
go run cmd/mcp-server/main.go -transport http -addr :8080
```

### 工具调用

```bash
# 健康检查
curl http://localhost:8080/health

# 列出所有工具
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list"}'

# 调用工具（使用 API Token 认证）
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc":"2.0",
    "id":1,
    "method":"tools/call",
    "params":{
      "name":"stock_basic.stock_basic",
      "arguments":{"ts_code":"000001.SZ"}
    }
  }'

# 或者使用 X-API-Token header
curl -X POST http://localhost:8080/stock \
  -H "X-API-Token: your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc":"2.0",
    "id":1,
    "method":"tools/call",
    "params":{
      "name":"stock_basic.stock_basic",
      "arguments":{"ts_code":"000001.SZ"}
    }
  }'
```

**📖 详细文档**: 查看 [MCP 服务器完整指南](docs/MCP_SERVER_GUIDE.md) 或 [快速开始指南](docs/QUICK_START.md)

**📚 文档中心**: 访问 [docs/README.md](docs/README.md) 查看所有文档

## 测试

- 单元测试使用 `httptest.Server` 进行HTTP模拟
- 集成测试需要 `TUSHARE_TOKEN` 环境变量
- 运行集成测试：`go test -tags=integration ./...`

## 架构

```
tushare-go/
├── cmd/                    # CLI应用程序
│   ├── mcp-server/         # MCP服务器
│   │   ├── config/         # 配置包
│   │   ├── main.go         # 主函数入口
│   │   ├── server.go       # 服务器核心逻辑
│   │   ├── tools.go        # 工具注册
│   │   ├── middleware.go   # HTTP中间件
│   │   └── types.go        # 核心数据结构
│   ├── generator/          # 代码生成器
│   └── examples/           # 示例代码
│       ├── chain_client/   # 链式调用示例
│       ├── daily/          # 日线数据示例
│       └── ...             # 其他示例
├── internal/               # 内部包
│   └── gen/                # 代码生成器
├── pkg/                    # 公共包
│   ├── sdk/                # 核心SDK
│   │   ├── apis/           # 链式调用接口
│   │   └── api/            # API 实现
│   └── mcp/                # MCP服务器
│       └── tools/          # MCP工具模块（26个）
├── config.example.json     # 配置文件示例
└── go.mod
```

## 最新优化

### Token 负载均衡功能 (2026-03-08)

- ✅ **多 Token 支持**: 配置多个 API token 自动负载均衡
- ✅ **轮询算法**: 按顺序依次使用 token，请求分布均匀
- ✅ **随机算法**: 随机选择 token，更适合高并发场景
- ✅ **并发安全**: 使用原子操作保证线程安全
- ✅ **向后兼容**: 完全兼容原有单 token 使用方式

**使用示例**：
```go
// 配置多个 token，轮询负载均衡
tokens := []string{"token1", "token2", "token3"}
config, _ := sdk.NewConfigWithTokens(tokens, "roundrobin")
client := sdk.NewClient(config)

// 使用方式与原来完全相同
stocks, err := stockbasic.StockBasic(ctx, client, req)
```

📖 **详细文档**: [Token 负载均衡指南](docs/TOKEN_LOAD_BALANCING.md)

### 版本管理自动化 (2026-03-08)

- ✅ **自动版本获取**: 从 Git tag 自动获取版本号
- ✅ **构建信息注入**: 自动注入版本号、commit hash、构建时间
- ✅ **版本显示**: 支持 `--version` 参数查看详细版本信息
- ✅ **发布流程**: 简化版本发布流程，一个命令完成
- ✅ **CI/CD 友好**: 支持环境变量自定义版本信息

**使用示例**：
```bash
# 查看版本信息
./bin/tushare-mcp --version

# 输出示例:
# Tushare MCP Server
# Version: v1.0.0 (ca3c630)
# Full Info: v1.0.0, commit: ca3c630, built at: 2026-03-08_09:00:00
```

📖 **详细文档**: [版本管理指南](docs/VERSION_MANAGEMENT.md)

### API Token 认证功能 (2026-03-08)

- ✅ **多用户支持**: 支持在配置文件中设置多个合法的 API token
- ✅ **灵活认证**: 支持 `Authorization: Bearer` 和 `X-API-Token` 两种认证方式
- ✅ **动态 token**: 用户提供的 token 自动用于 Tushare API 调用
- ✅ **向后兼容**: 仍支持 `TUSHARE_TOKEN` 环境变量方式
- ✅ **安全增强**: 可为不同用户配置不同的 token，实现访问控制

**配置示例**：
```json
{
  "api_tokens": [
    "your-tushare-token-1",
    "your-tushare-token-2"
  ]
}
```

**使用示例**：
```bash
# 使用配置文件中的 token
curl -X POST http://localhost:8080/stock \
  -H "Authorization: Bearer your-tushare-token-1" \
  -H "Content-Type: application/json" \
  -d '{"name":"stock_basic.stock_basic","arguments":{}}'
```

📖 **详细文档**: [API Token 认证指南](docs/API_TOKEN_AUTH.md)

### MCP服务器重构 (2026-03-08)

- ✅ **代码优化**: tools.go 从170行减少到94行 (-45%)
- ✅ **结构优化**: server.go 从254行减少到179行 (-30%)
- ✅ **注册表模式**: 使用工具注册表消除重复代码
- ✅ **配置分离**: 创建config子目录，职责分离更清晰
- ✅ **传输简化**: 移除"both"传输类型，只保留stdio和http
- ✅ **配置统一**: 只保留一个config.example.json配置文件示例
- ✅ **.gitignore修复**: 修复过度宽泛的忽略规则

### 代码质量提升

- 减少总代码行数约5%，同时提高可维护性
- 使用注册表模式统一工具注册接口
- 配置相关逻辑独立成包，提高内聚性
- 删除不必要的并发启动复杂性

## 许可证

MIT许可证

## 📝 模块名称说明

本项目已从 `github.com/chenniannian90/tushare-go` 重构为 `tushare-go`，以提供更简洁的导入路径：

```go
// 旧的导入路径 (仍可使用)
import "github.com/chenniannian90/tushare-go/pkg/sdk"

// 新的导入路径 (推荐)
import "tushare-go/pkg/sdk"
```

### 重构内容

- ✅ 模块名简化为 `tushare-go`
- ✅ 所有 565 个 Go 文件的导入路径已更新
- ✅ 26 个 API 模块，223 个工具，全部编译通过
- ✅ 二进制文件统一放在 `bin/` 目录
- ✅ 完整的文档更新
- ✅ 优化的MCP服务器代码结构

### 迁移指南

如果您正在使用旧版本，需要更新导入路径：

```bash
# 查找需要更新的文件
grep -r "github.com/chenniannian90/tushare-go" --include="*.go"

# 批量替换
find . -name "*.go" -exec sed -i '' 's|github.com/chenniannian90/tushare-go|tushare-go|g' {} +
```
