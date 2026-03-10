# Python Tushare vs Go Tushare 爬虫实现对比

## 🎯 实现概述

本文档详细对比了 Python Tushare 和 Go Tushare 的实时数据爬虫实现，展示了从 Python 到 Go 的完整移植过程。

## 📊 核心函数对比

| 功能 | Python 函数 | Go 函数 | 文件 |
|------|------------|---------|------|
| 实时行情 | `get_realtime_quotes()` | `GetRealtimeQuotes()` | `realtime_crawler.go` |
| 实时排名 | - | `GetRealtimeList()` | `realtime_list_crawler.go` |
| 分笔成交 | `get_today_ticks()` | `GetTodayTicks()` | `realtime_tick_crawler.go` |
| 实时分笔 | - | `GetRealtimeTick()` | `realtime_tick_crawler.go` |

## 🔧 详细实现对比

### 1. 实时行情数据

#### Python 版本 (tushare/stock/trading.py:324)
```python
def get_realtime_quotes(symbols=None):
    # 构建股票代码列表
    symbols_list = ''
    if isinstance(symbols, list):
        for code in symbols:
            symbols_list += ct._code_to_symbol(code) + ','
    else:
        symbols_list = ct._code_to_symbol(symbols)

    # 构建请求URL
    request = Request(ct.LIVE_DATA_URL%(ct.P_TYPE['http'],
                                        ct.DOMAINS['sinahq'],
                                        _random(), symbols_list))

    # 发送请求并解析响应
    text = urlopen(request,timeout=10).read()
    text = text.decode('GBK')
    reg = re.compile(r'\="(.*?)\";')
    data = reg.findall(text)

    # 解析股票代码
    regSym = re.compile(r'(?:sh|sz|gb_)(.*?)\=')
    syms = regSym.findall(text)

    # 构建DataFrame
    df = pd.DataFrame(data_list, columns=ct.LIVE_DATA_COLS)
    return df
```

#### Go 版本 (realtime_crawler.go:34)
```go
func GetRealtimeQuotes(symbols interface{}) ([]RealtimeQuoteItem, error) {
    // 构建股票代码列表
    symbolsList := buildSymbolsList(symbols)

    // 构建请求URL
    randomStr := generateRandomString(13)
    url := fmt.Sprintf("http://hq.sinajs.cn/rn=%s&list=%s",
                       randomStr, symbolsList)

    // 发送HTTP请求
    resp, _ := client.Get(url)
    defer resp.Body.Close()

    // 解析响应数据
    items, _ := parseSinaResponse(responseText)
    return items, nil
}
```

### 2. 数据解析对比

#### Python 版本 - 正则表达式
```python
# 提取数据内容
reg = re.compile(r'\="(.*?)\";')
data = reg.findall(text)

# 提取股票代码
regSym = re.compile(r'(?:sh|sz|gb_)(.*?)\=')
syms = regSym.findall(text)

# 分割字段
for index, row in enumerate(data):
    data_list.append([astr for astr in row.split(',')])
```

#### Go 版本 - 正则表达式
```go
// 正则提取数据
dataReg := regexp.MustCompile(`="(.*?)";`)
dataMatches := dataReg.FindAllStringSubmatch(text, -1)

// 正则提取股票代码
symReg := regexp.MustCompile(`(?:sh|sz|gb_)(.*?)=` )
symMatches := symReg.FindAllStringSubmatch(text, -1)

// 分割数据
fields := strings.Split(dataStr, ",")
```

### 3. 股票代码转换对比

#### Python 版本 (tushare/stock/cons.py:409)
```python
def _code_to_symbol(code):
    if code in INDEX_LABELS:
        return INDEX_LIST[code]
    elif code[:3] == 'gb_':  # 美股
        return code
    else:
        if len(code) != 6:
            return code
        else:
            # 5,6,9开头或11,13开头 -> 上海
            return 'sh%s'%code if code[:1] in ['5', '6', '9'] or code[:2] in ['11', '13'] else 'sz%s'%code
```

#### Go 版本 (realtime_crawler.go:78)
```go
func codeToSymbol(code string) string {
    // 已经是新浪格式，直接返回
    if strings.HasPrefix(code, "sh") || strings.HasPrefix(code, "sz") || strings.HasPrefix(code, "gb_") {
        return code
    }

    // 6位代码转换
    if len(code) == 6 {
        firstChar := string(code[0])
        if firstChar == "5" || firstChar == "6" || firstChar == "9" ||
           strings.HasPrefix(code, "11") || strings.HasPrefix(code, "13") {
            return "sh" + code
        }
        return "sz" + code
    }
    return code
}
```

### 4. 数据字段处理对比

#### Python 版本 - A股字段
```python
LIVE_DATA_COLS = ['name', 'open', 'pre_close', 'price', 'high', 'low',
                  'bid', 'ask', 'volume', 'amount',
                  'b1_v', 'b1_p', 'b2_v', 'b2_p', 'b3_v', 'b3_p',
                  'b4_v', 'b4_p', 'b5_v', 'b5_p',
                  'a1_v', 'a1_p', 'a2_v', 'a2_p', 'a3_v', 'a3_p',
                  'a4_v', 'a4_p', 'a5_v', 'a5_p', 'date', 'time', 's']

# 处理量字段（去掉末尾两位）
ls = [cls for cls in df.columns if '_v' in cls]
for txt in ls:
    df[txt] = df[txt].map(lambda x : x[:-2])
```

#### Go 版本 - A股字段解析
```go
var liveDataCols = []string{
    "name", "open", "pre_close", "price", "high", "low", "bid", "ask",
    "volume", "amount",
    "b1_v", "b1_p", "b2_v", "b2_p", "b3_v", "b3_p", "b4_v", "b4_p",
    "b5_v", "b5_p",
    "a1_v", "a1_p", "a2_v", "a2_p", "a3_v", "a3_p", "a4_v", "a4_p",
    "a5_v", "a5_p", "date", "time", "s",
}

// 处理量字段（去掉末尾两位）
func parseVolumeField(s string) int {
    if len(s) >= 2 {
        s = s[:len(s)-2]
    }
    return parseInt(s)
}
```

### 5. 分笔成交对比

#### Python 版本 (tushare/stock/trading.py:232)
```python
def get_today_ticks(code=None, retry_count=3, pause=0.001):
    symbol = ct._code_to_symbol(code)
    date = du.today()

    # 获取分页信息
    request = Request(ct.TODAY_TICKS_PAGE_URL % (...))
    data_str = urlopen(request, timeout=10).read()
    data_str = eval(data_str, ...)
    pages = len(data_str['detailPages'])

    # 逐页获取数据
    for pNo in range(1, pages+1):
        data = data.append(_today_ticks(symbol, date, pNo, ...))
    return data
```

#### Go 版本 (realtime_tick_crawler.go:241)
```go
func GetTodayTicks(code string) ([]RealtimeTickItem, error) {
    symbol := codeToSymbol(code)
    date := time.Now().Format("2006-01-02")

    // 获取分页信息
    pages, _ := getTickPages(symbol, date)

    // 逐页获取数据
    for pageNo := 1; pageNo <= pages; pageNo++ {
        ticks, _ := getTodayTicksByPage(symbol, date, pageNo)
        allTicks = append(allTicks, ticks...)
    }
    return allTicks, nil
}
```

## 🌐 数据源对比

### 新浪财经数据源

| 数据类型 | Python URL | Go URL |
|---------|-----------|--------|
| 实时行情 | `http://hq.sinajs.cn/...` | `http://hq.sinajs.cn/...` |
| 分笔数据 | `http://vip.stock.finance.sina.com.cn/...` | `http://vip.stock.finance.sina.com.cn/...` |

### 东方财富数据源

| 数据类型 | Python URL | Go URL |
|---------|-----------|--------|
| 实时排名 | `http://push2.eastmoney.com/...` | `http://push2.eastmoney.com/...` |
| 分笔数据 | `http://push2.eastmoney.com/...` | `http://push2.eastmoney.com/...` |

## 📈 数据结构对比

### Python 版本 - DataFrame
```python
# 返回 pandas DataFrame
df = pd.DataFrame(data_list, columns=ct.LIVE_DATA_COLS)
df['code'] = syms_list

# 访问数据
print(df['price'])      # 价格列
print(df.iloc[0])       # 第一行数据
```

### Go 版本 - Struct Array
```go
// 返回结构体数组
type RealtimeQuoteItem struct {
    TsCode   string  `json:"ts_code"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    // ... 其他字段
}

// 访问数据
for _, item := range items {
    fmt.Println(item.Price)
}
```

## 🔍 关键差异点

### 1. 编码处理
- **Python**: `text.decode('GBK')`
- **Go**: 使用 `bufio.Reader` 和默认编码

### 2. 数据结构
- **Python**: `pandas.DataFrame`
- **Go**: 结构体数组 `[]RealtimeQuoteItem`

### 3. 错误处理
- **Python**: 异常机制 `try-except`
- **Go**: 显式错误返回 `error`

### 4. 并发处理
- **Python**: 单线程同步
- **Go**: 支持原生并发 `goroutine`

### 5. 类型安全
- **Python**: 动态类型
- **Go**: 静态类型，编译时检查

## 🚀 Go 版本的优势

### 1. 性能优势
- **编译型语言**: 执行效率更高
- **原生并发**: 支持高性能并发获取
- **内存管理**: 更高效的垃圾回收

### 2. 部署优势
- **单一可执行文件**: 无需运行时环境
- **交叉编译**: 轻松跨平台部署
- **容器化**: 适合微服务架构

### 3. 维护优势
- **类型安全**: 编译时发现错误
- **静态分析**: 更好的代码检查工具
- **标准化**: 统一的代码风格

## 📋 使用建议

### 适合使用 Python 版本的场景
- 数据分析和处理
- 机器学习模型训练
- 快速原型开发
- Jupyter Notebook 环境

### 适合使用 Go 版本的场景
- 高性能实时数据采集
- 生产环境部署
- 微服务架构
- 需要高并发的场景

## 🔄 迁移指南

### 从 Python 迁移到 Go

1. **函数映射**
   ```python
   # Python
   quotes = ts.get_realtime_quotes('000001')

   # Go
   quotes, _ := stock_market.GetRealtimeQuotes("000001")
   ```

2. **数据访问**
   ```python
   # Python
   price = quotes['price'][0]

   # Go
   price := quotes[0].Price
   ```

3. **错误处理**
   ```python
   # Python
   try:
       quotes = ts.get_realtime_quotes('000001')
   except Exception as e:
       print(e)

   # Go
   quotes, err := stock_market.GetRealtimeQuotes("000001")
   if err != nil {
       log.Fatal(err)
   }
   ```

## 🎓 学习建议

1. **从简单开始**: 先熟悉实时行情接口
2. **理解数据格式**: 学习新浪数据响应格式
3. **错误处理**: 妥善处理网络和数据解析错误
4. **性能优化**: 利用 Go 的并发特性
5. **持续维护**: 数据源可能会变更，需要定期更新

---

**最后更新**: 2026年3月
**版本**: 1.0.0
**作者**: 基于 Python Tushare 项目移植