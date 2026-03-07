# SDK API 调用方式

Tushare Go SDK 提供两种主要的 API 调用方式，可以根据项目需求选择最适合的方式。

## 🎯 两种调用方式

### 方式 1: 直接调用（原有方式）

```go
import stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"

// 直接调用 API 函数
data, err := stockboard.TopList(ctx, client, req)
```

**优点:**
- 简单直接
- 无需额外学习
- 无需额外依赖

**缺点:**
- 需要导入多个 API 包
- IDE 提示不够友好

**适合:** 快速脚本、API 调用较少的场景

---

### 方式 2: apis 包类型化方法（推荐）⭐

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
- 统一的导入路径

**缺点:**
- 需要导入 apis 包

**适合:** 大型项目、需要类型安全和良好 IDE 支持的场景

---

## 📦 apis 包提供的方法

### 股票板块相关
- `TopList()` - 龙虎榜每日统计
- `LimitList()` - 涨跌停和炸板数据
- `DragonList()` - 游资交易每日明细
- `TopInst()` - 营业部席位买入排名
- `ThsConcept()` - 同花顺概念板块
- `EmHot()` - 东方财富App热榜

### 股票市场数据
- `Daily()` - 日线行情
- `DailyBasic()` - 每日基本面指标
- `Weekly()` - 周线行情
- `Monthly()` - 月线行情

### 股票基础信息
- `TradeCal()` - 交易日历
- `StockBasicInfo()` - 股票列表

### 财务数据
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

### 对比两种方式

#### 使用直接调用
```go
import (
    stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
    stockmarket "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_market"
)

// 需要导入多个包
data1, _ := stockboard.TopList(ctx, client, req1)
data2, _ := stockmarket.Daily(ctx, client, req2)
```

#### 使用 apis 包
```go
import sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"

// 只需导入一个包
data1, _ := sdkapis.TopList(ctx, client, req1)
data2, _ := sdkapis.Daily(ctx, client, req2)
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

### 1. 推荐使用 apis 包

在项目中统一使用 `sdkapis` 包：
```go
import sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"
```

这样可以获得：
- 一致的 API 接口
- 类型安全
- 更好的错误处理
- 更少的 import 语句

### 2. 创建服务层

对于复杂项目，建议创建服务层封装业务逻辑：
```go
type StockService struct {
    client *sdk.Client
}

func (s *StockService) GetTopGainers(ctx) ([]Stock, error) {
    // 业务逻辑...
}
```

### 3. 错误处理

始终检查和处理错误：
```go
data, err := sdkapis.TopList(ctx, client, req)
if err != nil {
    return fmt.Errorf("获取龙虎榜失败: %w", err)
}
```

## 🔄 从旧代码迁移

### 之前（直接调用）
```go
import stockboard "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"

data, err := stockboard.TopList(ctx, client, req)
```

### 现在（apis 包，推荐）
```go
import sdkapis "github.com/chenniannian90/tushare-go/pkg/sdk/apis"

data, err := sdkapis.TopList(ctx, client, req)
```

只需简单替换包名即可！

## 📖 更多示例

查看 `cmd/examples/` 目录下的完整示例：
- `sdk_usage` - SDK 基础使用
- `boards` - 板块数据示例
- `financial_data` - 财务数据示例
- 更多...

## ❓ 常见问题

**Q: 为什么要提供两种调用方式？**
A: 不同场景有不同需求。直接调用适合简单场景，apis 包适合复杂项目。

**Q: 哪种方式性能更好？**
A: 性能相同。选择取决于代码组织和可维护性。

**Q: apis 包会包含所有 API 吗？**
A: 我们会持续添加常用的 API。如果需要特殊 API，可以继续使用直接调用方式。

**Q: 如何添加新的 API 到 apis 包？**
A: 在 `pkg/sdk/apis/` 中创建新文件，参考 `stock.go` 添加方法。
