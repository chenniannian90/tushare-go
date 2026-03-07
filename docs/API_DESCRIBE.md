# API 规范 `__describe__` 字段说明

## 概述

在 Tushare API 的 JSON 规范文件中，`__describe__` 字段用于存储接口的元数据信息，包括官方文档链接、中文名称和分类信息。

## 字段结构

```json
{
  "api_name": "daily",
  "description": "获取日线行情数据",
  "__describe__": {
    "url": "https://tushare.pro/document/2?doc_id=109",
    "name": "日线行情",
    "category": "行情数据"
  },
  "request_params": [...],
  "response_fields": [...]
}
```

## 字段说明

### `url`
- **类型**: `string`
- **描述**: Tushare Pro 官方文档的完整 URL
- **用途**: 提供接口文档的快速访问链接
- **示例**: `https://tushare.pro/document/2?doc_id=109`

### `name`
- **类型**: `string`
- **描述**: 接口的中文名称
- **用途**: 用于生成文档、MCP 工具描述等
- **示例**: `日线行情`

### `category`
- **类型**: `string`
- **描述**: 接口所属的分类
- **用途**: 帮助组织和筛选相关接口
- **可选值**:
  - `行情数据` - 市场行情相关数据
  - `股票信息` - 股票基本信息
  - `财务数���` - 财务报表数据
  - `指数数据` - 指数相关数据
  - `交易日历` - 交易日历信息
  - `权益数据` - 分红送股等权益数据

## 使用场景

### 1. 文档生成
在生成 API 文档时，可以自动引用官方文档链接：

```go
// 生成包含文档链接的API说明
func GenerateAPIDoc(spec *APISpec) string {
    return fmt.Sprintf(`
## %s

**描述**: %s
**官方文档**: %s
**分类**: %s
`, spec.__describe__.Name, spec.Description, spec.__describe__.URL, spec.__describe__.Category)
}
```

### 2. MCP 工具描述
在 MCP 服务器中为工具添加详细的描述信息：

```go
tool := mcp.Tool{
    Name:        spec.APIName,
    Description: fmt.Sprintf("%s - %s", spec.__describe__.Name, spec.Description),
    InputSchema: generateInputSchema(spec),
}
```

### 3. 在线帮助系统
为 CLI 工具提供在线帮助链接：

```bash
# 显示API信息并附上文档链接
$ tushare-cli info daily
名称: 日线行情
描述: 获取日线行情数据
分类: 行情数据
文档: https://tushare.pro/document/2?doc_id=109
```

## API 规范文件组织结构

API 规范文件按照 `category` 字段分类存放在不同的子目录中：

```
internal/gen/specs/
├── market_data/          # 行情数据
│   ├── daily.json
│   ├── weekly.json
│   ├── monthly.json
│   ├── pro_bar.json
│   ├── daily_basic.json
│   ├── moneyflow.json
│   └── limit_list.json
├── stock_info/           # 股票信息
│   ├── stock_basic.json
│   ├── top10_holders.json
│   └── holder_number.json
├── financial_data/       # 财务数据
│   ├── income.json
│   ├── balancesheet.json
│   └── fina_indicator.json
├── index_data/           # 指数数据
│   ├── index_basic.json
│   ├── index_daily.json
│   ├── concept.json
│   └── concept_detail.json
├── trading_calendar/     # 交易日历
│   └── trade_cal.json
├── equity_data/          # 权益数据
│   └── dividend.json
└── complex_types_example.json  # 复杂类型示例
```

## 已添加 `__describe__` 的 API

目前以下 API 规范文件已添加 `__describe__` 字段：

### 行情数据
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `daily` | 日线行情 | 行情数据 | doc_id=109 |
| `weekly` | 周线行情 | 行情数据 | doc_id=109 |
| `monthly` | 月线行情 | 行情数据 | doc_id=109 |
| `pro_bar` | 综合行情 | 行情数据 | doc_id=109 |
| `daily_basic` | 每日基本面 | 行情数据 | doc_id=32 |
| `moneyflow` | 个股资金流向 | 行情数据 | doc_id=109 |
| `limit_list` | 涨跌停列表 | 行情数据 | doc_id=109 |

### 股票信息
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `stock_basic` | 股票信息 | 股票信息 | doc_id=25 |
| `top10_holders` | 十大股东数据 | 股票信息 | doc_id=109 |
| `holder_number` | 股东人数数据 | 股票信息 | doc_id=109 |

### 财务数据
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `income` | 利润表 | 财务数据 | doc_id=79 |
| `balancesheet` | 资产负债表 | 财务数据 | doc_id=79 |
| `fina_indicator` | 财务指标 | 财务数据 | doc_id=86 |

### 指数数据
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `index_basic` | 指数基本信息 | 指数数据 | doc_id=126 |
| `index_daily` | 指数日线行情 | 指数数据 | doc_id=126 |
| `concept` | 概念板块分类 | 指数数据 | doc_id=126 |
| `concept_detail` | 概念板块成分股 | 指数数据 | doc_id=126 |

### 交易日历
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `trade_cal` | 交易日历 | 交易日历 | doc_id=109 |

### 权益数据
| API 名称 | 接口名称 | 分类 | 文档 ID |
|---------|---------|------|---------|
| `dividend` | 分红送股 | 权益数据 | doc_id=117 |

## 代码生成器支持

代码生成器模板已更新，可以读取和使用 `__describe__` 字段：

```go
// 在生成的代码中包含文档链接
// 官方文档: https://tushare.pro/document/2?doc_id=109
func Daily(ctx context.Context, client *sdk.Client, req *DailyRequest) ([]DailyItem, error) {
    // ...
}
```

## 添加新的 `__describe__` 字段

要为其他 API 规范添加 `__describe__` 字段：

1. **确定文档信息**：
   - 访问 https://tushare.pro/document/2
   - 找到对应的 API 文档
   - 获取文档 ID（从 URL 参数 `doc_id` 提取）

2. **添加字段到 JSON**：
```json
{
  "api_name": "your_api",
  "description": "API描述",
  "__describe__": {
    "url": "https://tushare.pro/document/2?doc_id=XXX",
    "name": "接口中文名称",
    "category": "接口分类"
  }
}
```

3. **验证格式**：
```bash
# 检查 JSON 格式是否正确
jq '.' internal/gen/specs/your_api.json

# 验证 __describe__ 字段
jq '.__describe__' internal/gen/specs/your_api.json
```

## 常见文档 ID 参考

根据 Tushare Pro 文档结构，一些常见的 doc_id：

- `25` - 股票列表 (stock_basic)
- `32` - 每日基本面 (daily_basic)
- `79` - 财务报表 (income, balancesheet)
- `86` - 财务指标 (fina_indicator)
- `109` - 行情数据 (daily, trade_cal)
- `117` - 分红送股 (dividend)
- `126` - 指数数据 (index_basic)

## 最佳实践

1. **保持一致性**: 所有 API 规范都应包含 `__describe__` 字段
2. **准确链接**: 确保 URL 指向正确的官方文档页面
3. **简洁描述**: name 字段应该简短明了，便于显示
4. **合理分类**: category 应该符合业务逻辑，便于用户查找

## 扩展性

`__describe__` 字段可以扩展包含更多元数据：

```json
"__describe__": {
    "url": "https://tushare.pro/document/2?doc_id=109",
    "name": "日线行情",
    "category": "行情数据",
    "tags": ["A股", "日线", "OHLC"],
    "rate_limit": "200次/分钟",
    "update_time": "实时",
    "data_type": "时间序列"
}
```

这样可以支持更强大的文档生成和工具开发。
