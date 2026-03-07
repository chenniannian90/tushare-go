# Tushare Go SDK + MCP 服务器

一个生产就绪的 Go SDK，用于 [Tushare Pro](https://tushare.pro)（拥有200+ API的中国金融数据平台），以及一个 MCP（模型上下文协议）服务器，将这些API暴露给Claude等AI代理。

## 功能特性

- ✅ 针对20+ Tushare Pro REST API的类型安全Go SDK
- ✅ 代码生成器，可从JSON规范自动生成API封装
- ✅ MCP服务器，用于Claude桌面版集成
- ✅ 全面的测试覆盖率（≥80%）
- ✅ 零外部依赖（除MCP SDK外）
- ✅ 生产就绪，包含7个完整实现的工具
- ✅ **链式调用客户端，统一的API访问入口**

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
    "github.com/chenniannian90/tushare-go/pkg/sdk"
    "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
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

# 构建 MCP 服务器
make build-mcp

# 构建代码生成器
make build-gen

# 生成 API 代码
make gen

# 生成 MCP 工具
go run cmd/gen-mcp-tools/main.go -optimized
```

## MCP 服务器使用

### 快速启动

```bash
# 设置 Token
export TUSHARE_TOKEN="your-tushare-token"

# 启动 MCP 服务器
go run cmd/mcp-server/main.go

# 或者使用构建好的版本
make build-mcp
./bin/tushare-mcp
```

### 传输模式

```bash
# HTTP 模式 (默认)
go run cmd/mcp-server/main.go -transport http

# Stdio 模式 (用于 AI 助手)
go run cmd/mcp-server/main.go -transport stdio

# 双重模式 (同时支持)
go run cmd/mcp-server/main.go -transport both
```

### 工具调用

```bash
# 健康检查
curl http://localhost:8080/health

# 列出所有工具
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list"}'

# 调用工具
curl -X POST http://localhost:8080/mcp \
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

## 测试

- 单元测试使用 `httptest.Server` 进行HTTP模拟
- 集成测试需要 `TUSHARE_TOKEN` 环境变量
- 运行集成测试：`go test -tags=integration ./...`

## 架构

```
tushare-go/
├── cmd/                    # CLI应用程序
│   ├── mcp-server/         # MCP服务器
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
└── go.mod
```

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
- ✅ 25 个 API 模块，223 个工具，全部编译通过
- ✅ 二进制文件统一放在 `bin/` 目录
- ✅ 完整的文档更新

### 迁移指南

如果您正在使用旧版本，需要更新导入路径：

```bash
# 查找需要更新的文件
grep -r "github.com/chenniannian90/tushare-go" --include="*.go"

# 批量替换
find . -name "*.go" -exec sed -i '' 's|github.com/chenniannian90/tushare-go|tushare-go|g' {} +
```
