# Tushare Go SDK 示例

本目录包含演示如何使用 Tushare Go SDK 的示例程序。

## 前提条件

将您的 Tushare Pro API token 设置为环境变量：

```bash
export TUSHARE_TOKEN=your_token_here
```

## 可用示例

### 1. 股票基本信息 (`stock_basic`)

演示获取基本股票信息：

```bash
go run cmd/examples/stock_basic/main.go
```

功能：
- 获取所有股票
- 通过代码获取特定股票
- 按交易所筛选股票
- 按上市状态筛选

### 2. 日度市场数据 (`daily`)

演示获取日度 OHLCV 市场数据：

```bash
go run cmd/examples/daily/main.go
```

功能：
- 获取特定股票的日度数据
- 按日期范围查询
- 获取特定日期的市场数据

### 3. 日度基本面指标 (`daily_basic`)

演示获取日度基本面指标：

```bash
go run cmd/examples/daily_basic/main.go
```

功能：
- PE、PB、PS 比率
- 市值
- 换手率
- 量比

### 4. 交易日历 (`trade_cal`)

演示获取交易日历信息：

```bash
go run cmd/examples/trade_cal/main.go
```

功能：
- 获取日期范围内的交易日
- 确定市场开市/收市状态
- 统计交易日与非交易日

## 构建示例

您可以将示例构建为可执行文件：

```bash
# 构建 stock_basic 示例
go build -o bin/stock_basic_example cmd/examples/stock_basic/main.go

# 运行它
./bin/stock_basic_example
```

## 错误处理

所有示例都包含适当的错误处理：

```go
data, err := api.StockBasic(ctx, client, req)
if err != nil {
    log.Fatalf("获取股票失败: %v", err)
}
```

## 速率限制

请注意 Tushare API 速率限制：
- 免费账户：200次请求/分钟
- 在生产环境中考虑在请求之间添加延迟

## 后续步骤

探索其他可用的 API：
- `weekly`、`monthly` 用于不同时间范围
- `pro_bar` 用于统一市场数据
- `income`、`balancesheet` 用于财务数据
- `fina_indicator` 用于财务比率
- `moneyflow` 用于资金流向数据
- 还有更多！

## 获取帮助

- [Tushare Pro 文档](https://tushare.pro/document/2)
- [API 参考](https://tushare.pro/document/2?doc_id=109)
