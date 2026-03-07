# MCP 工具 Spec 文件集成

## 概述

成功实现了从 Tushare API spec 文件中提取元数据，生成更准确和完整的 MCP 工具。

## 实现的功能

### 1. 描述信息提取 ✅

**之前**: 使用简单的英文描述生成
```go
Description: "Retrieve bcbestotcqt data from Tushare bond API"
```

**现在**: 从 spec 文件中提取真实的中文描述
```go
Description: "柜台流通式债券最优报价"
```

### 2. 字段信息提取 ✅

**之前**: 使用占位符字段
```go
type BcBestotcqtInput struct {
    TsCode string `json:"ts_code,omitempty" jsonschema:"Security code"`
}
```

**现在**: 从 spec 文件中提取真实字段
```go
type BcBestotcqtInput struct {
    TradeDate string `json:"trade_date,omitempty" jsonschema:"报价日期(YYYYMMDD格式，下同)"`
    StartDate string `json:"start_date,omitempty" jsonschema:"开始日期"`
    EndDate   string `json:"end_date,omitempty" jsonschema:"结束日期"`
    TsCode    string `json:"ts_code,omitempty" jsonschema:"TS代码"`
}
```

### 3. 字段类型映射 ✅

自动将 spec 文件中的类型映射到 Go 类型：

| Spec 类型 | Go 类型 |
|-----------|---------|
| str, string | string |
| int, integer | int |
| float, float64, double | float64 |
| bool, boolean | bool |

### 4. 字段名转换 ✅

自动将 snake_case 转换为 PascalCase：
- `trade_date` → `TradeDate`
- `ts_code` → `TsCode`
- `start_date` → `StartDate`

## Spec 文件结构

```json
{
  "api_name": "柜台流通式债券最优报价",
  "api_code": "bc_bestotcqt",
  "description": "柜台流通式债券最优报价",
  "__describe__": {
    "url": "https://tushare.pro/document/2?doc_id=323",
    "name": "柜台流通式债券最优报价",
    "category": "债券专题___bond"
  },
  "request_params": [
    {
      "name": "trade_date",
      "type": "str",
      "description": "报价日期(YYYYMMDD格式，下同)",
      "required": false
    }
  ],
  "response_fields": [...]
}
```

## 生成的代码示例

### 输入结构体
```go
// BcBestotcqtInput defines the input schema
type BcBestotcqtInput struct {
    TradeDate string `json:"trade_date,omitempty" jsonschema:"报价日期(YYYYMMDD格式，下同)"`
    StartDate string `json:"start_date,omitempty" jsonschema:"开始日期"`
    EndDate   string `json:"end_date,omitempty" jsonschema:"结束日期"`
    TsCode    string `json:"ts_code,omitempty" jsonschema:"TS代码"`
}
```

### 工具注册
```go
tool := &mcp.Tool{
    Name:        "bond.bc_bestotcqt",
    Description: "柜台流通式债券最优报价",  // 从 spec 提取
    InputSchema: inputSchema,              // 自动生成 JSON Schema
}
```

### Handler 函数
```go
handler := func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    var input BcBestotcqtInput
    if err := json.Unmarshal(req.Params.Arguments, &input); err != nil {
        return errorResult(err), nil
    }

    apiReq := &bond.BcBestotcqtRequest{
        TradeDate: input.TradeDate,  // 真实字段名
        StartDate: input.StartDate,
        EndDate:   input.EndDate,
        TsCode:    input.TsCode,
    }

    items, err := bond.BcBestotcqt(ctx, r.client, apiReq)
    // ...
}
```

## 使用方法

### 标准模式
```bash
# 生成标准工具（包含 spec 描述）
go run cmd/gen-mcp-tools/main.go
```

### 优化模式
```bash
# 生成优化工具（包含 spec 描述 + 真实字段 + JSON Schema）
go run cmd/gen-mcp-tools/main.go -optimized
```

### 指定 spec 路径
```bash
# 使用默认路径: internal/gen/specs
go run cmd/gen-mcp-tools/main.go

# 指定自定义路径
go run cmd/gen-mcp-tools/main.go -spec-path /path/to/specs
```

## 智能匹配算法

生成器使用多种模式匹配 spec 文件：

1. **精确匹配**: `module___module/*___api_code.json`
2. **通配符匹配**: `*module*/*___api_code.json`
3. **简单匹配**: `module/*___api_code.json`

## 回退机制

如果找不到 spec 文件或字段信息：

1. **描述**: 回退到自动生成的英文描述
2. **字段**: 使用默认的 `TsCode` 字段
3. **类型**: 默认使用 `string` 类型

## 示例对比

### bc_bestotcqt 工具

| 项目 | 之前 | 现在 |
|------|------|------|
| 描述 | "Retrieve bcbestotcqt data from Tushare bond API" | "柜台流通式债券最优报价" |
| 字段数 | 1 (TsCode) | 4 (TradeDate, StartDate, EndDate, TsCode) |
| 字段描述 | "Security code" | "报价日期(YYYYMMDD格式，下同)" |
| 类型安全性 | ❌ 弱 | ✅ 强 |

### cb_basic 工具

| 项目 | 之前 | 现在 |
|------|------|------|
| 描述 | "Retrieve cbbasic data from Tushare bond API" | "可转债基础信息" (如果有 spec) |
| 字段数 | 1 (TsCode) | 3 (TsCode, ListDate, Exchange) |
| 字段描述 | "Security code" | "转债代码", "上市日期", "上市交易所" |

## 优势

### 1. 准确性 ✅
- 使用官方 API 文档的描述
- 真实的字段名称和类型
- 完整的参数列表

### 2. 开发效率 ✅
- 自动生成准确的输入结构体
- 减少手动编写和调试时间
- 保持与 API 同步

### 3. 用户体验 ✅
- 中文描述更易理解
- 完整的参数说明
- JSON Schema 提供验证

### 4. 维护性 ✅
- 单一数据源 (spec 文件)
- 自动化生成
- 易于更新

## 技术实现

### 核心函数

```go
// 从 spec 文件加载 API 定义
func (g *MCPGenerator) loadAPISpec(module, apiCode string) *APISpec

// 从 spec 提取字段信息
func (g *MCPGenerator) extractFieldsFromSpec(module, apiCode string) []FieldWithGoType

// 生成工具描述（优先使用 spec）
func (g *MCPGenerator) toToolDescription(module, function, apiName string) string
```

### 类型转换

```go
func (g *MCPGenerator) specTypeToGoType(specType string) string {
    switch strings.ToLower(specType) {
    case "str", "string":
        return "string"
    case "int", "integer":
        return "int"
    case "float", "float64", "double":
        return "float64"
    case "bool", "boolean":
        return "bool"
    default:
        return "string"
    }
}
```

### 名称转换

```go
func (g *MCPGenerator) toPascalCase(s string) string {
    parts := strings.Split(s, "_")
    var result strings.Builder
    for _, part := range parts {
        if len(part) > 0 {
            result.WriteString(strings.ToUpper(part[:1]))
            result.WriteString(part[1:])
        }
    }
    return result.String()
}
```

## 未来改进

1. **响应字段**: 生成 Output 结构体的完整字段定义
2. **验证**: 添加字段验证逻辑
3. **默认值**: 从 spec 中提取默认值信息
4. **示例**: 生成使用示例代码
5. **测试**: 自动生成单元测试

## 总结

这次增强成功实现了：
- ✅ 从 spec 文件提取真实的 API 描述
- ✅ 从 spec 文件提取完整的字段信息
- ✅ 自动类型映射和名称转换
- ✅ 生成准确的 JSON Schema
- ✅ 保持回退机制确保兼容性

生成的 MCP 工具现在更加准确、完整和易用，大大提升了开发效率和用户体验。
