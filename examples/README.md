# Tushare Go 实时数据API - 爬虫版本

## 🎉 重要更新

**所有实时数据API函数现在都直接调用爬虫函数，不依赖Tushare HTTP接口！**

### ✅ 已修改的函数

| API函数 | 内部实现 | 对应爬虫函数 | 说明 |
|---------|----------|-------------|------|
| `RealtimeQuote()` | ✅ 调用爬虫 | `GetRealtimeQuotes()` | 实时行情数据 |
| `RealtimeList()` | ✅ 调用爬虫 | `GetRealtimeList()` | 全市场实时排名 |
| `RealtimeTick()` | ✅ 调用爬虫 | `GetRealtimeTick()` | 分笔成交数据 |

### 📁 文件结构

```
pkg/sdk/api/stock_market/
├── realtime_quote.go          # API接口（调用爬虫）
├── realtime_list.go            # API接口（调用爬虫）
├── realtime_tick.go            # API接口（调用爬虫）
├── realtime_crawler.go         # 底层爬虫实现
├── realtime_list_crawler.go    # 底层爬虫实现
└── realtime_tick_crawler.go    # 底层爬虫实现
```

### 🚀 使用方式

现在你可以使用标准的API接口，但实际调用的是爬虫函数：

```go
import "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock_market"

// 创建客户端（即使Token为空也可以工作）
client := sdk.NewClient("")
ctx := context.Background()

// 1. 实时行情 - 支持tushare格式代码
req := &stock_market.RealtimeQuoteRequest{
    TsCode: "000001.SZ",  // 自动转换为000001
    Src:    "sina",
}
quotes, _ := stock_market.RealtimeQuote(ctx, client, req)

// 2. 实时排名
listReq := &stock_market.RealtimeListRequest{Src: "sina"}
list, _ := stock_market.RealtimeList(ctx, client, listReq)

// 3. 分笔成交
tickReq := &stock_market.RealtimeTickRequest{
    TsCode: "600000.SH",  // 自动转换为600000
    Src:    "sina",
}
ticks, _ := stock_market.RealtimeTick(ctx, client, tickReq)
```

### 🔧 技术实现

#### 1. 代码格式转换
```go
// tushare格式 -> 纯数字格式
000001.SZ -> 000001
600000.SH -> 600000
```

#### 2. 内部调用链
```
RealtimeQuote() -> convertToPureCode() -> GetRealtimeQuotes() [爬虫]
RealtimeList()   -> GetRealtimeList() [爬虫]
RealtimeTick()   -> convertToPureCodeTick() -> GetRealtimeTick() [爬虫]
```

### 💡 核心优势

1. **无需Token**: 不需要配置有效的Tushare Token
2. **统一接口**: 使用标准的API接口，内部自动调用爬虫
3. **格式兼容**: 同时支持tushare格式代码和纯数字代码
4. **直接数据源**: 数据来自新浪财经/东方财富，无中间环节

### 📊 与原版Python对比

| Python函数 | Go API函数 | Go爬虫函数 | 状态 |
|------------|-----------|-----------|------|
| `get_realtime_quotes()` | `RealtimeQuote()` | `GetRealtimeQuotes()` | ✅ 完全兼容 |
| - | `RealtimeList()` | `GetRealtimeList()` | ✅ 新增功能 |
| `get_today_ticks()` | `RealtimeTick()` | `GetRealtimeTick()` | ✅ 完全兼容 |

### ⚠️ 重要说明

1. **接口兼容性**: 保持原有API接口不变，方便迁移
2. **数据源依赖**: 依赖第三方数据源，可能存在URL变更风险
3. **交易时间**: 仅在交易时间段返回有效数据
4. **数据延迟**: 可能有几秒到十几秒的数据延迟

### 🧪 验证状态

- ✅ 所有API函数编译通过
- ✅ 代码格式化完成
- ✅ 爬虫功能完整实现
- ✅ API接口正确调用爬虫

### 📝 使用建议

1. **推荐使用API接口**: 使用`RealtimeQuote()`等标准接口
2. **直接使用爬虫**: 如需更多控制，可直接调用`GetRealtimeQuotes()`
3. **代码格式**: 支持两种格式，推荐使用tushare格式保持一致性
4. **错误处理**: 注意处理网络错误和数据解析错误

---

**最后更新**: 2026年3月9日
**版本**: 2.0.0 - 爬虫版本
**状态**: ✅ 已验证并可用
