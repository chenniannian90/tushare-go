# SDK 链式调用

Tushare Go SDK 现在支持多种 API 调用方式，可以根据项目需求选择最适合的方式。

## 🎯 三种调用方式

### 方式 1: 直接调用（原有方式）

```go
import stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"

// 直接调用 API 函数
data, err := stockboard.TopList(ctx, client, req)
```

**优点:**
- 简单直接
- 无需额外学习

**缺点:**
- 需要导入多个 API 包
- IDE 提示不够友好

**适合:** 快速脚本、API 调用较少的场景

---

### 方式 2: SDK 内置链式调用

```go
import "github.com/chenniannian90/tushare-go/pkg/sdk"

// 使用链式调用 + 通用 CallAPI 方法
var result struct {
    Fields []string                 `json:"fields"`
    Items  []map[string]interface{} `json:"items"`
}

err := client.StockBoard().CallAPI(
    ctx,
    "top_list",
    params,
    fields,
    &result,
)
```

**优点:**
- 无需导入具体 API 包
- 代码组织清晰
- 可以调用任何 API

**缺点:**
- 需要手动定义结果结构
- 需要手动解析字段

**适合:** 已知 API 名称和格式、需要灵活性的场景

---

### 方式 3: apis 包类型化方法（推荐）⭐

```go
import (
    "github.com/chenniannian90/tushare-go/pkg/sdk"
    sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
)

// 使用类型化方法
data, err := sdkapis.TopList(ctx, client, req)
dailyData, err := sdkapis.Daily(ctx, client, req)
```

**优点:**
- 类型安全
- IDE 自动提示完整
- 无需手动解析结果
- 代码更简洁

**缺点:**
- 需要导入 apis 包

**适合:** 大型项目、需要类型安全和良好 IDE 支持的场景

---

## 📚 可用的 API 分类

SDK 提供以下 API 分类：

### 股票相关
- `StockBoard()` - 股票板块（龙虎榜、涨跌停、概念板块）
- `StockMarket()` - 股票市场数据（日线、周线、月线）
- `StockBasic()` - 股票基础信息（交易日历、股票列表）
- `StockFinancial()` - 财务数据（利润表、资产负债表、财务指标）

### 其他市场
- `Index()` - 指数数据
- `Futures()` - 期货数据
- `Fund()` - 公募基金
- `HKStock()` - 港股数据
- `Bond()` - 债券数据
- `ETF()` - ETF 数据
- `Forex()` - 外汇数据
- `Options()` - 期权数据
- `Spot()` - 现货数据
- `USStock()` - 美股数据
- `Wealth()` - 财富管理
- `Industry()` - 行业经济
- `LLMCorpus()` - 大模型语料
- `Macro()` - 宏观经济

## 📦 apis 包提供的方法

### 股票板块 (StockBoard)
- `TopList()` - 龙虎榜每日统计
- `LimitList()` - 涨跌停和炸板数据
- `DragonList()` - 游资交易每日明细
- `TopInst()` - 营业部席位买入排名
- `ThsConcept()` - 同花顺概念板块
- `EmHot()` - 东方财富App热榜

### 股票市场 (StockMarket)
- `Daily()` - 日线行情
- `DailyBasic()` - 每日基本面指标
- `Weekly()` - 周线行情
- `Monthly()` - 月线行情

### 股票基础 (StockBasic)
- `TradeCal()` - 交易日历
- `StockBasicInfo()` - 股票列表

### 财务数据 (StockFinancial)
- `Income()` - 利润表数据
- `Balancesheet()` - 资产负债表
- `FinaIndicator()` - 财务指标
- `Dividend()` - 分红数据

*更多方法持续添加中...*

## 🎓 使用示例

### 基础用法

```go
package main

import (
    "context"
    "github.com/chenniannian90/tushare-go/pkg/sdk"
    sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
)

func main() {
    // 创建客户端
    config, _ := sdk.NewConfig("your_token")
    client := sdk.NewClient(config)
    ctx := context.Background()

    // 使用 apis 包调用 API（推荐）
    data, err := sdkapis.TopList(ctx, client, &sdkapis.TopListRequest{})
    if err != nil {
        panic(err)
    }

    // 处理数据...
}
```

### 链式调用

```go
// 连续调用多个 API
sdkapis.TopList(ctx, client, req1)
sdkapis.Daily(ctx, client, req2)
sdkapis.TradeCal(ctx, client, req3)

// 或使用 SDK 链式调用
client.StockBoard().CallAPI(ctx, "top_list", params, fields, &result)
client.StockMarket().CallAPI(ctx, "daily", params, fields, &result)
```

### 在项目中组织 API 调用

```go
// myapp/api/stock.go
package api

import (
    "context"
    "github.com/chenniannian90/tushare-go/pkg/sdk"
    sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
)

type StockService struct {
    client *sdk.Client
}

func NewStockService(client *sdk.Client) *StockService {
    return &StockService{client: client}
}

func (s *StockService) GetTopList(ctx context.Context) ([]TopListData, error) {
    req := &stockboard.TopListRequest{
        TradeDate: "20240115",
    }
    return sdkapis.TopList(ctx, s.client, req)
}

func (s *StockService) GetDailyData(ctx context.Context, tsCode string) ([]DailyData, error) {
    req := &stockmarket.DailyRequest{
        TsCode: tsCode,
    }
    return sdkapis.Daily(ctx, s.client, req)
}
```

## 💡 最佳实践

### 1. 选择合适的调用方式

- **小型项目/脚本**: 直接调用或 apis 包
- **大型项目**: apis 包 + 自定义服务层
- **特殊需求**: CallAPI 方法

### 2. 统一使用 apis 包

在项目中统一使用 `sdkapis` 包：
```go
import sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
```

这样可以获得：
- 一致的 API 接口
- 类型安全
- 更好的错误处理

### 3. 创建服务层

对于复杂项目，建议创建服务层封装业务逻辑：
```go
type StockService struct {
    client *sdk.Client
}

func (s *StockService) GetTopGainers(ctx) ([]Stock, error) {
    // 业务逻辑...
}
```

### 4. 错误处理

始终检查和处理错误：
```go
data, err := sdkapis.TopList(ctx, client, req)
if err != nil {
    return fmt.Errorf("获取龙虎榜失败: %w", err)
}
```

## 🔄 从旧代码迁移

### 之前:
```go
import stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"

data, err := stockboard.TopList(ctx, client, req)
```

### 现在（推荐）:
```go
import sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"

data, err := sdkapis.TopList(ctx, client, req)
```

只需简单替换包名即可！

## 📖 更多示例

查看 `cmd/examples/` 目录下的完整示例：
- `sdk_usage` - SDK 基础使用
- `chain_call` - 链式调用对比
- `boards` - 板块数据示例
- `financial_data` - 财务数据示例
- 更多...

## ❓ 常见问题

**Q: 为什么要提供多种调用方式？**
A: 不同场景有不同需求。我们提供灵活性，让你选择最适合的方式。

**Q: 哪种方式性能最好？**
A: 性能相同。选择取决于代码组织和可维护性。

**Q: apis 包会包含所有 API 吗？**
A: 我们会持续添加常用的 API。如果需要，可以自己扩展或使用 CallAPI。

**Q: 如何添加新的 API 到 apis 包？**
A: 在 `pkg/sdk/apis/` 中创建新文件，参考现有代码添加方法。
