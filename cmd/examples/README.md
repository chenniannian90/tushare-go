# Tushare Go SDK 示例

本目录包含演示如何使用 Tushare Go SDK 的示例程序。

## 前提条件

将您的 Tushare Pro API token 设置为环境变量：

```bash
export TUSHARE_TOKEN=your_token_here
```

## 📊 可用示例

### 🏢 股票数据示例

#### 1. 股票基本信息 (`stock_basic`)
演示获取基本股票信息：

```bash
go run cmd/examples/stock_basic/main.go
```

功能：
- 获取所有股票
- 通过代码获取特定股票
- 按交易所筛选股票（SSE、SZSE）
- 按上市状态筛选

#### 2. 日度市场数据 (`daily`)
演示获取日度 OHLCV 市场数据：

```bash
go run cmd/examples/daily/main.go
```

功能：
- 获取特定股票的日度数据
- 按日期范围查询
- 获取特定日期的市场数据

#### 3. 日度基本面指标 (`daily_basic`)
演示获取日度基本面指标：

```bash
go run cmd/examples/daily_basic/main.go
```

功能：
- PE、PB、PS 比率
- 市值（总市值、流通市值）
- 换手率、量比

#### 4. 涨跌停和概念板块 (`boards`)
演示获取涨跌停和板块数据：

```bash
go run cmd/examples/boards/main.go
```

功能：
- 涨跌停股票列表
- 龙虎榜数据
- 概念板块分类
- 东方财富热榜

#### 5. 财务数据 (`financial_data`)
演示获取财务报表数��：

```bash
go run cmd/examples/financial_data/main.go
```

功能：
- 利润表数据
- 资产负债表数据
- 财务指标分析
- 分红派息数据

#### 6. 交易日历 (`trade_cal`)
演示获取交易日历信息：

```bash
go run cmd/examples/trade_cal/main.go
```

功能：
- 获取日期范围内的交易日
- 确定市场开市/收市状态
- 统计交易日与非交易日

### 7. 链式调用 (`chain_call`)
演示两种不同的 API 调用方式：

```bash
go run cmd/examples/chain_call/main.go
```

功能：
- 对比直接调用和链式调用两种方式
- 展示如何实现链式调用包装器
- 提供最佳实践建议

**链式调用示例：**
```go
// 创建链式调用客户端
apiClient := NewAPIClient(client)

// 使用链式调用
apiClient.StockBoard().TopList(ctx, req)
apiClient.StockMarket().Daily(ctx, req)
apiClient.StockBasic().TradeCal(ctx, req)
```

### 📈 指数数据示例

#### 7. 指数数据 (`index_data`)
演示获取各类指数数据：

```bash
go run cmd/examples/index_data/main.go
```

功能：
- 指数基本信息
- 指数日线数据
- 按市场筛选（SSE、SZSE）
- 主要市场指数查询

### 💰 期货数据示例

#### 8. 期货数据 (`futures`)
演示获取期货市场数据：

```bash
go run cmd/examples/futures/main.go
```

功能：
- 期货合约基本信息
- 期货日线行情
- 按交易所筛选
- 主力合约数据

### 💼 基金数据示例

#### 9. 公募基金 (`fund`)
演示获取公募基金数据：

```bash
go run cmd/examples/fund/main.go
```

功能：
- 基金基本信息
- 基金净值数据
- 基金分红记录
- 基金经理信息

### 🌏 港股数据示例

#### 10. 港股数据 (`hk_stock`)
演示获取港股市场数据：

```bash
go run cmd/examples/hk_stock/main.go
```

功能：
- 港股基本信息
- 港股日线行情
- 港股交易日历

## 🏗️ 构建示例

您可以将示例构建为可执行文件：

```bash
# 构建 stock_basic 示例
go build -o bin/stock_basic_example cmd/examples/stock_basic/main.go

# 运行它
./bin/stock_basic_example
```

或者使用 Makefile：

```bash
make build-examples
```

## ⚠️ 错误处理

所有示例都包含适当的错误处理：

```go
data, err := api.StockBasic(ctx, client, req)
if err != nil {
    log.Fatalf("获取股票失败: %v", err)
}
```

## ⏱️ 速率限制

请注意 Tushare API 速率限制：
- **免费账户**: 200次请求/分钟
- **付费账户**: 根据级别有不同的限制
- 在生产环境中考虑在请求之间添加延迟

## 📚 API 目录结构

新的 SDK 目录结构按数据类型组织：

```
pkg/sdk/api/
├── stock/          # 股票数据
├── index/          # 指数数据
├── futures/        # 期货数据
├── fund/           # 基金数据
├── hk_stock/       # 港股数据
├── bond/           # 债券数据
├── etf/            # ETF数据
├── options/        # 期权数据
└── ...
```

## 🔍 探索更多 API

### 股票相关
- `weekly`、`monthly` - 不同时间范围数据
- `stock_market` - 市场数据汇总
- `stock_financial` - 财务数据
- `stock_margin` - 融资融券
- `stock_reference` - 参考数据

### 指数相关
- `index_classify` - 指数分类
- `index_member` - 指数成分股
- `index_weekly`、`index_monthly` - 周线、月线

### 其他
- `moneyflow` - 资金流向数据
- `block_trade` - 大宗交易
- `top10_holders` - 十大股东
- 还有更多！

## 📖 获取帮助

- [Tushare Pro 文档](https://tushare.pro/document/2)
- [API 参考](https://tushare.pro/document/2?doc_id=109)
- [GitHub Issues](https://github.com/chenniannian90/tushare-go/issues)

## 💡 使用技巧

### 1. 使用 context 控制超时
```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
data, err := api.StockBasic(ctx, client, req)
```

### 2. 批量查询优化
```go
// 分批查询大量数据
for i := 0; i < len(stockCodes); i += batchSize {
    end := i + batchSize
    if end > len(stockCodes) {
        end = len(stockCodes)
    }
    batch := stockCodes[i:end]
    // 处理批次...
}
```

### 3. 缓存常用数据
```go
// 缓存股票列表等不常变化的数据
cache := make(map[string][]Stock)
if cached, ok := cache["stock_list"]; ok {
    return cached, nil
}
// 从 API 获取并缓存...
```

## 🤝 贡献示例

欢迎提交更多示例！请确保：
1. 代码格式规范（运行 `go fmt`）
2. 添加适当的注释
3. 更新本 README 文件
4. 遵循项目编码规范
