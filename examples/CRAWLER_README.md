# Tushare Go 实时数据爬虫实现

本文档总结了直接参考 Python Tushare 实现的 Go 版本实时数据爬虫函数。

## 📁 实现的文件

### 1. `realtime_crawler.go` - 实时行情爬虫
**函数**: `GetRealtimeQuotes(symbols interface{})`

**功能**: 获取单个或多个股票的实时行情数据

**数据源**: 新浪财经 (hq.sinajs.cn)

**参数**:
- `symbols`: 股票代码，支持字符串或字符串数组
  - 单个股票: `"000001"`
  - 多个股票: `[]string{"000001", "600000"}`

**返回**: `[]RealtimeQuoteItem`

**实现细节**:
```go
// 转换为���浪格式代码
symbol := "sh" + code  // 上海市场
symbol := "sz" + code  // 深圳市场

// 构建请求URL
url := "http://hq.sinajs.cn/rn=random&list=sh600000,sz000001"

// 解析响应
// 正则提取: `="数据内容";`
// 分割字段: 用逗号分隔的32个字段
```

### 2. `realtime_list_crawler.go` - 实时排名爬虫
**函数**: `GetRealtimeList(src string)`

**功能**: 获取全市场股票实时排名数据

**数据源**:
- 新浪财经 (sina)
- 东方财富 (dc/eastmoney)

**参数**:
- `src`: 数据源选择，`"sina"` 或 `"dc"`

**返回**: `[]RealtimeListItem`

**实现细节**:
```go
// 新浪接口
url := "http://hq.sinajs.cn/list=s_sh000001,s_sz399001"

// 东方财富接口
url := "http://push2.eastmoney.com/api/qt/clist/get?..."
```

### 3. `realtime_tick_crawler.go` - 分笔成交爬虫
**函数**: `GetRealtimeTick(code string, src string)` 和 `GetTodayTicks(code string)`

**功能**: 获取单个股票的分笔成交数据

**数据源**:
- 新浪财经 (sina)
- 东方财富 (dc/eastmoney)

**参数**:
- `code`: 6位股票代码，如 `"000001"`
- `src`: 数据源选择

**返回**: `[]RealtimeTickItem`

**实现细节**:
```go
// 新浪分笔数据URL
url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/..."

// 东方财富分笔数据URL
url := "http://push2.eastmoney.com/api/qt/stock/fflow/daykline/get?..."
```

## 🔧 核心实现特点

### 1. 完全模拟Python版本逻辑
- **URL构建**: 与Python版本相同的数据源URL格式
- **数据解析**: 使用相同的正则表达式和解析逻辑
- **字段映射**: A股32字段，美股28字段
- **错误处理**: 网络请求、数据解析的完整错误处理

### 2. 股票代码转换
```go
func codeToSymbol(code string) string {
    // 5,6,9开头或11,13开头 -> sh (上海)
    // 其他 -> sz (深圳)
    // gb_ 前缀保持不变 (美股)
}
```

### 3. 数据解析
```go
// 新浪响应格式: sh600000="平安银行,11.50,11.45,..."
// 正则提取: `\="(.*?)\";`
// 分割处理: 用逗号分隔各字段
```

### 4. 量字段处理
```go
// 委买委卖量需要去掉末尾两位（手转换为股）
volume := volume[:len(volume)-2]
```

## 🚀 使用示例

### 示例1: 获取单个股票实时行情
```go
import "github.com/chenniannian90/tushare-go/pkg/sdk/api/stock_market"

quotes, err := stock_market.GetRealtimeQuotes("000001")
if err == nil && len(quotes) > 0 {
    q := quotes[0]
    fmt.Printf("股票: %s (%s)\n", q.Name, q.TsCode)
    fmt.Printf("价格: %.2f\n", q.Price)
    fmt.Printf("涨跌: %.2f%%\n", (q.Price-q.PreClose)/q.PreClose*100)
}
```

### 示例2: 获取多个股票实时行情
```go
codes := []string{"000001", "600000", "000002"}
quotes, err := stock_market.GetRealtimeQuotes(codes)
```

### 示例3: 获取实时排名
```go
list, err := stock_market.GetRealtimeList("sina")
for _, item := range list {
    fmt.Printf("%s: %.2f (%.2f%%)\n",
        item.TsCode, item.Price,
        (item.Price-item.PreClose)/item.PreClose*100)
}
```

### 示例4: 获取分笔成交
```go
ticks, err := stock_market.GetRealtimeTick("000001", "sina")
for _, tick := range ticks {
    fmt.Printf("%s %.2f %d手 %s\n",
        tick.Time, tick.Price, tick.Volume, tick.Direction)
}
```

### 示例5: 获取完整分笔数据
```go
fullTicks, err := stock_market.GetTodayTicks("600000")
fmt.Printf("共 %d 笔成交\n", len(fullTicks))
```

## 📊 数据字段说明

### A股实时行情字段 (32个)
- **基础**: name, open, pre_close, price, high, low, bid, ask
- **成交**: volume, amount
- **五档盘口**: b1_v~b5_p, a1_v~a5_p (买卖一到五的量和价)
- **时间**: date, time

### 分笔成交字段
- **ts_code**: 股票代码
- **time**: 成交时间
- **price**: 成交价格
- **volume**: 成交量(手)
- **direction**: 买卖方向(买盘/卖盘/中性盘)
- **amount**: 成交金额

## ⚠️ 重要注意事项

### 1. 不需要 Tushare Token
- 这些函数直接爬取公开数据源
- 不依赖 Tushare 官方 API
- 不受 Tushare 积分限制

### 2. 数据源限制
- **交易时间**: 仅在交易时间段返回有效数据
- **数据延迟**: 可能有几秒到十几秒的延迟
- **数据准确性**: 依赖数据源的准确性

### 3. 法律合规
- **使用条款**: 请遵守相关网站的使用条款
- **robots.txt**: 遵守网站的爬虫规则
- **频率限制**: 避免频繁请求，设置合理的间隔时间

### 4. 技术维护
- **URL变更**: 数据源URL可能会变更
- **格式变化**: 数据格式可能会调整
- **反爬虫**: 可能需要处理反爬虫机制

## 🔄 与 Python 版本的对应关系

| Python 函数 | Go 函数 | 说明 |
|------------|---------|------|
| `get_realtime_quotes()` | `GetRealtimeQuotes()` | 实时行情 |
| - | `GetRealtimeList()` | 实时排名 |
| `get_today_ticks()` | `GetTodayTicks()` | 当日分笔 |
| - | `GetRealtimeTick()` | 实时分笔 |

## 🧪 测试建议

1. **交易时间测试**: 在交易时间段内测试实时数据
2. **多股票测试**: 测试批量获取功能
3. **错误处理**: 测试网络错误、无效代码等异常情况
4. **性能测试**: 测试大量股票获取的性能
5. **数据验证**: 对比多个数据源的数据准确性

## 📈 未来改进方向

1. **更多数据源**: 添加腾讯、网易等数据源
2. **缓存机制**: 添加本地缓存减少请求频率
3. **异步处理**: 支持并发获取多个股票数据
4. **WebSocket**: 实现真正的实时推送
5. **数据监控**: 监控数据源可用性和数据质量

---

**开发者**: 基于Python Tushare项目实现
**最后更新**: 2026年3月
**版本**: 1.0.0