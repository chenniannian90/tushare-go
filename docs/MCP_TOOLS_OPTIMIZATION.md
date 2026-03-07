# MCP Tools 优化实现指南

## 概述

本文��描述了对 `pkg/mcp/tools` 模块的优化，遵循参考实现 (`ci-mcp/internal/tools/bug`) 的模式。

## 优化模式 vs 原始模式

### 原始模式的问题

1. **缺少 JSON Schema**: 没有输入参数的类型定义和验证
2. **类型安全**: 使用 `map[string]interface{}` 缺乏类型检查
3. **文档**: 工具描述简单，缺少参数说明
4. **维护**: 需要适配器层，增加复杂度

### 优化模式的优势

1. **JSON Schema 自动生成**: 使用 `jsonschema.For[T](nil)` 自动生成
2. **类型安全**: 明确的 Input/Output 结构体定义
3. **MCP SDK 原生类型**: 直接使用 `mcp.Tool`, `mcp.CallToolResult`
4. **更好的文档**: 通过 `jsonschema` 标签提供字段描述
5. **直接注册**: 无需适配器层，直接注册到 MCP Server

## 优化模式实现

### 1. Input/Output 结构体定义

```go
// CbBasicInput defines the input schema for cb_basic tool
type CbBasicInput struct {
    TsCode      string `json:"ts_code,omitempty" jsonschema:"Convertible bond code"`
    AnnDate     string `json:"ann_date,omitempty" jsonschema:"Announcement date (YYYYMMDD)"`
    StartDate   string `json:"start_date,omitempty" jsonschema:"Start date (YYYYMMDD)"`
    EndDate     string `json:"end_date,omitempty" jsonschema:"End date (YYYYMMDD)"`
}

// CbBasicOutput defines the output schema
type CbBasicOutput struct {
    Data  []CbBasicItem `json:"data" jsonschema:"Convertible bond data list"`
    Total int           `json:"total" jsonschema:"Total count"`
}
```

### 2. 工具注册函数

```go
func (r *OptimizedBondTools) registerCbBasic() {
    // 自动生成输入 Schema
    inputSchema, _ := jsonschema.For[CbBasicInput](nil)

    tool := &mcp.Tool{
        Name:        "bond.cb_basic",
        Description: "Retrieve convertible bond basic information from Tushare API",
        InputSchema: inputSchema,
    }

    handler := func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
        // 类型安全的输入解析
        var input CbBasicInput
        if err := json.Unmarshal(req.Params.Arguments, &input); err != nil {
            return &mcp.CallToolResult{
                IsError: true,
                Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf(`{"error":"Invalid input: %v"}`, err)}},
            }, nil
        }

        // 调用 API
        items, err := bond.CbBasic(ctx, r.client, apiReq)
        if err != nil {
            return &mcp.CallToolResult{
                IsError: true,
                Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf(`{"error":"API call failed: %v"}`, err)}},
            }, nil
        }

        // 格式化输出
        output := CbBasicOutput{
            Data:  convertCbBasicItems(items),
            Total: len(items),
        }

        outputJSON, _ := json.MarshalIndent(output, "", "  ")
        return &mcp.CallToolResult{
            Content: []mcp.Content{&mcp.TextContent{Text: string(outputJSON)}},
        }, nil
    }

    r.server.AddTool(tool, handler)
}
```

## 使用方法

### 生成标准模式工具

```bash
# 生成基础模式的 MCP 工具
go run cmd/gen-mcp-tools/main.go
```

### 生成优化模式工具

```bash
# 生成优化模式的 MCP 工具（推荐）
go run cmd/gen-mcp-tools/main.go -optimized
```

### 在 MCP Server 中使用

#### 标准模式

```go
// 创建标准模式的工具注册表
registry := mcp.NewToolRegistry(client)

// 通过适配器注册
adapter := server.NewToolAdapter(registry)
adapter.RegisterTools(srv)
```

#### 优化模式

```go
// 创建优化模式的工具
bondTools := bondtools.NewOptimizedBondTools(srv, client)
bondTools.RegisterAll() // 注册所有 bond 工具

stockTools := stocktools.NewOptimizedStockTools(srv, client)
stockTools.RegisterAll() // 注册所有 stock 工具
```

## 迁移指南

### 从标准模式迁移到优化模式

1. **生成优化代码**:
   ```bash
   go run cmd/gen-mcp-tools/main.go -optimized
   ```

2. **更新 MCP Server**:
   ```go
   // 旧代码
   registry := mcp.NewToolRegistry(client)
   adapter := server.NewToolAdapter(registry)
   adapter.RegisterTools(srv)

   // 新代码
   bondTools := bondtools.NewOptimizedBondTools(srv, client)
   bondTools.RegisterAll()
   ```

3. **测试集成**: 确保所有工具正常工作

## 模块结构

### 优化模式生成的文件

```
pkg/mcp/tools/bond/
├── registry.go              # 标准模式注册表
├── cb_basic.go              # 标准模式实现
├── optimized_cb_basic.go    # 优化模式实现（新增）
├── optimized_bc_bestotcqt.go # 优化模式实现（新增）
└── ...                      # 其他优化工具
```

### 使用优化模式的优势

1. **更快的开发**: 自动生成 Input/Output 结构体
2. **更好的错误处理**: 编译时类型检查
3. **更清晰的 API**: 明确的输入输出定义
4. **更好的文档**: JSON Schema 提供完整参数说明
5. **更容易维护**: 减少适配器层

## 实现细节

### JSON Schema 生成

使用 `github.com/google/jsonschema-go` 库：

```go
inputSchema, err := jsonschema.For[CbBasicInput](nil)
```

### 字段描述

通过 `jsonschema` 标签提供描述：

```go
TsCode string `json:"ts_code,omitempty" jsonschema:"Convertible bond code (e.g. 113030.SZ)"`
```

### MCP SDK 集成

直接使用 MCP SDK 的原生类型：

```go
tool := &mcp.Tool{
    Name:        "tool.name",
    Description: "Tool description",
    InputSchema: inputSchema,
}
```

## 最佳实践

1. **使用优化模式**: 新项目应使用 `-optimized` 标志生成工具
2. **完整字段定义**: 在模板中添加完整的 API 字段定义
3. **错误处理**: 使用 `fmt.Sprintf` 格式化错误消息
4. **类型转换**: 实现完整的 `convertXxxItems` 函数
5. **文档化**: 为所有字段添加 `jsonschema` 标签描述

## 后续改进

1. **自动字段提取**: 从 API 定义自动提取字段信息
2. **类型映射**: 改进 Go 类型到 JSON Schema 类型的映射
3. **验证**: 添加输入参数验证逻辑
4. **测试**: 生成单元测试代码
5. **文档**: 生成 API 文档

## 参考实现

基于 `/Users/mac-new/go/src/ext-gitlab.denglin.com/ci-tool/ci-mcp/internal/tools/bug` 的实现模式。

## 总结

优化模式提供了：
- ✅ JSON Schema 自动生成
- ✅ 类型安全的输入输出
- ✅ MCP SDK 原生支持
- ✅ 更好的错误处理
- ✅ 更清晰的代码结构
- ✅ 更容易的维护

建议所有新项目使用优化模式，现有项目逐步迁移。
