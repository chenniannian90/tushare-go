# 复杂类型支持文档

## 概述

代码生成器现已支持复杂和嵌套数据类型，包括数组、对象和布尔类型。

## 支持的复杂类型

### 1. 数组类型 (Array)

数组类型允许您定义包含多个相同类型元素的字段。

```json
{
  "name": "tags",
  "type": "array",
  "description": "标签数组",
  "items": {
    "name": "tag",
    "type": "string",
    "description": "单个标签"
  }
}
```

生成的 Go 代码：
```go
Tags []string `json:"tags"`
```

### 2. 对象类型 (Object)

对象类型用于嵌套的键值对数据。

```json
{
  "name": "metadata",
  "type": "object",
  "description": "元数据对象"
}
```

生成的 Go 代码：
```go
Metadata map[string]interface{} `json:"metadata"`
```

### 3. 布尔类型 (Boolean)

布尔类型用于true/false值。

```json
{
  "name": "include_nested",
  "type": "bool",
  "description": "是否包含嵌套数据"
}
```

生成的 Go 代码：
```go
IncludeNested bool `json:"include_nested,omitempty"`
```

## 完整示例

查看 `specs/complex_types_example.json` 文件了解完整示例：

```bash
# 生成复杂类型示例API
go run ./cmd/generate-single/main.go ./internal/gen/specs/complex_types_example.json /tmp/complex_example.go
```

## 生成的代码特性

### 智能类型转换

生成的代码包含智能类型转换逻辑，可以处理不同的数组类型：

- `[]interface{}` - 通用接口数组
- `[]string` - 字符串数组
- `[]float64` - 浮点数数组
- `[]int` - 整数数组

### 安全的数据提取

代码生成器为每种类型生成安全的类型断言和错误处理：

```go
// Handle array type for tags
tagsRaw, ok := item["tags"]
if !ok {
    return nil, fmt.Errorf("missing field tags")
}
var tags []string
if tagsRaw != nil {
    switch v := tagsRaw.(type) {
    case []interface{}:
        // 安全转换
    case []string:
        // 直接赋值
    default:
        tags = []string{}
    }
}
```

### 正确的参数处理

布尔类型参数使用正确的Go习惯用法：

```go
if req.IncludeNested {
    params["include_nested"] = req.IncludeNested
}
```

而不是：
```go
if req.IncludeNested != "" {  // 错误！
    params["include_nested"] = req.IncludeNested
}
```

## 类型映射表

| API类型 | Go类型 | 说明 |
|---------|--------|------|
| `string` | `string` | 字符串 |
| `int`, `integer` | `int` | 整数 |
| `float`, `double` | `float64` | 浮点数 |
| `bool`, `boolean` | `bool` | 布尔值 |
| `array<string>` | `[]string` | 字符串数组 |
| `array<float>` | `[]float64` | 浮点数数组 |
| `object` | `map[string]interface{}` | 对象/字典 |

## 测试

运行复杂类型测试：

```bash
# 单元测试
go test ./internal/gen -v -run TestComplexTypesSupport

# 集成测试
go test ./internal/gen -v -tags=integration -run TestGeneratedComplexTypesCode
```

## 迁移现有API

如果您的API规范使用了复杂类型，只需在JSON规范中添加相应的类型定义，代码生成器会自动处理其余部分。

1. 更新API规范JSON文件
2. 重新生成代码：`go run ./cmd/generator/main.go ./pkg/sdk/api`
3. 运行测试验证：`go test ./pkg/sdk/api`

## 限制和注意事项

1. **嵌套对象**：当前对象类型使用 `map[string]interface{}`，如果需要强类型嵌套结构，请在规范中定义完整结构。

2. **空数组**：生成的代码会妥善处理空数组和nil值。

3. **类型转换**：数组类型转换支持多种输入格式，但建议API返回一致的数据类型。

4. **性能考虑**：使用 `interface{}` 类型会有轻微的性能开销，但在大多数情况下可以忽略不计。
