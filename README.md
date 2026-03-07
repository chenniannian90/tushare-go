# Tushare Go SDK + MCP 服务器

一个生产就绪的 Go SDK，用于 [Tushare Pro](https://tushare.pro)（拥有200+ API的中国金融数据平台），以及一个 MCP（模型上下文协议）服务器，将这些API暴露给Claude等AI代理。

## 功能特性

- ✅ 针对20+ Tushare Pro REST API的类型安全Go SDK
- ✅ 代码生成器，可从JSON规范自动生成API封装
- ✅ MCP服务器，用于Claude桌面版集成
- ✅ 全面的测试覆盖率（≥80%）
- ✅ 零外部依赖（除MCP SDK外）
- ✅ 生产就绪，包含7个完整实现的工具

## 安装

```bash
go get github.com/chenniannian90/tushare-go
```

## 快速开始

### SDK使用

```go
package main

import (
    "context"
    "fmt"
    "github.com/chenniannian90/tushare-go/pkg/sdk"
    "github.com/chenniannian90/tushare-go/pkg/sdk/api"
)

func main() {
    config, _ := sdk.NewConfig("your-tushare-token")
    client := sdk.NewClient(config)

    req := &api.StockBasicRequest{
        TsCode: "000001.SZ",
    }

    stocks, err := api.StockBasic(context.Background(), client, req)
    if err != nil {
        panic(err)
    }

    fmt.Println(stocks)
}
```

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

## 开发

```bash
# 运行测试
make test

# 构建MCP服务器
make build-mcp

# 构建代码生成器
make build-gen

# 生成API代码
make gen
```

## 测试

- 单元测试使用 `httptest.Server` 进行HTTP模拟
- 集成测试需要 `TUSHARE_TOKEN` 环境变量
- 运行集成测试：`go test -tags=integration ./...`

## 架构

```
tushare-go/
├── cmd/                    # CLI应用程序
│   ├── mcp-server/         # MCP服务器
│   └── generator/          # 代码生成器
├── internal/               # 内部包
│   └── gen/                # 代码生成器
├── pkg/                    # 公共包
│   ├── sdk/                # 核心SDK
│   └── mcp/                # MCP服务器
└── go.mod
```

## 许可证

MIT许可证
