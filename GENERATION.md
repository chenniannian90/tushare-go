# 代码生成指南

本文档说明如何使用项目的代码生成工具来重新生成API代码、spec文件和MCP工具。

## 生成流程概述

项目使用三层生成架构：

```
Tushare文档 → Spec文件 → API代码 → MCP工具
   ↓              ↓           ↓          ↓
spec-gen     generator   gen-mcp-tools
```

## 快速开始

### 一键重新生成所有内容

```bash
make gen-all
```

这个命令会按顺序执行：
1. `make gen-specs` - 从Tushare文档重新生成spec文件
2. `make gen` - 从spec文件重新生成API代码
3. `make gen-mcp` - 从API代码重新生成MCP工具

## 单独的生成命令

### 1. 生成Spec文件

```bash
make gen-specs
```

**说明**：
- 从Tushare文档（`docs/api-directory.json`）生成spec文件到`internal/gen/specs/`
- Spec文件包含API的元数据、参数定义、响应字段和描述信息
- **关键修复**：正则表达式`描述：([\s\S]+?)(?:权限：|限量：|积分：|接口：|$)`能够正确提取跨行的描述文本

**依赖**：
- `docs/api-directory.json` - Tushare API目录结构文件

**输出**：
- `internal/gen/specs/**/*.json` - API规范文件

### 2. 生成API代码

```bash
make gen
```

**说明**：
- 从spec文件生成Go API包装代码到`pkg/sdk/api/`
- 每个API函数都会包含完整的注释和类型定义
- 模板包含description字段，确保函数注释包含完整描述

**依赖**：
- `internal/gen/specs/` - 规范文件目录
- `internal/gen/templates/api.go.tmpl` - 代码生成模板

**输出**：
- `pkg/sdk/api/**/*.go` - API包装代码

### 3. 生成MCP工具

```bash
make gen-mcp
```

**说明**：
- 从API代码生成MCP工具包装到`pkg/mcp/tools/`
- 自动从spec文件读取description作为工具描述
- 使用JSON schema提供类型安全的输入验证

**依赖**：
- `pkg/sdk/api/` - API代码目录
- `internal/gen/specs/` - 规范文件目录（用于读取description）

**输出**：
- `pkg/mcp/tools/**/*.go` - MCP工具代码

## 关键修复说明

### Spec生成器修复 (cmd/spec-gen/main.go:285)

**问题**：原始正则表达式无法匹配包含换行符的描述文本

**修复**：
```go
// 修复前
re := regexp.MustCompile(`描述：(.+?)(?:权限：|限量：|积分：|接口：|$)`)

// 修复后
re := regexp.MustCompile(`描述：([\s\S]+?)(?:权限：|限量：|积分：|接口：|$)`)
```

**说明**：使用`[\s\S]+?`代替`.+?`来匹配包括换行符在内的所有字符

### API生成模板修复 (internal/gen/templates/api.go.tmpl:32-33)

**修复**：
```go
// {{Title .APICode}} 调用 {{.APIName}} API
// {{.Description}}
func {{Title .APICode}}(ctx context.Context, client *sdk.Client, req *{{Title .APICode}}Request) ([]{{Title .APICode}}Item, error) {
```

**说明**：在函数注释中添加`{{.Description}}`字段，确保生成的函数包含完整描述

### MCP工具生成器 (cmd/gen-mcp-tools/main.go:384-404)

**逻辑**：
1. 优先从spec文件读取`Description`字段
2. 如果Description为空，fallback到`APIName`字段
3. 最后才fallback到自动生成的描述

```go
func (g *MCPGenerator) toToolDescription(module, function, apiName string) string {
    spec := g.loadAPISpec(module, apiName)
    if spec != nil {
        if spec.Description != "" {
            return spec.Description
        }
        if spec.APIName != "" {
            return spec.APIName
        }
    }
    // Fallback...
}
```

## 手动生成命令

如果你想直接运行生成器而不使用Makefile：

```bash
# 生成spec文件
go run cmd/spec-gen/main.go docs/api-directory.json internal/gen/specs

# 生成API代码
go run cmd/generator/main.go pkg/sdk/api

# 生成MCP工具
go run cmd/gen-mcp-tools/main.go -optimized
```

## 验证生成的代码

### 检查Spec文件

```bash
# 检查spec文件中是否有description
jq '.description' internal/gen/specs/股票数据___stock/基础数据___stock_basic/股票列表___stock_basic.json
```

应该输出：`"获取基础信息数据，包括股票代码、名称、上市日期、退市日期等"`

### 检查API代码

```bash
# 检查API函数注释
grep -A 2 "^func StockBasic" pkg/sdk/api/other/股票列表.go
```

应该输出：
```go
// StockBasic 调用 股票列表 API
// 获取基础信息数据，包括股票代码、名称、上市日期、退市日期等
func StockBasic(ctx context.Context, client *sdk.Client, req *StockBasicRequest) ([]StockBasicItem, error) {
```

### 检查MCP工具

```bash
# 检查MCP工具描述
grep "Description:" pkg/mcp/tools/stock_basic/stock_basic.go | head -1
```

应该输出：
```go
Description: "获取基础信息数据，包括股票代码、名称、上市日期、退市日期等",
```

## 故障排除

### 描述信息仍然为空

1. **检查Tushare文档格式**：确保文档页面包含正确的"描述："字段
2. **清除缓存**：删除`bin/`目录，重新构建生成器
3. **检查网络连接**：spec-gen需要从Tushare网站抓取数据

### MCP工具使用fallback描述

这通常意味着：
1. Spec文件中没有找到对应的API
2. Spec文件的`description`字段为空
3. API名称映射不正确

解决方案：检查`cmd/gen-mcp-tools/main.go`中的`loadAPISpec`函数逻辑

### 生成的代码编译错误

1. **清理旧代码**：`rm -rf pkg/sdk/api pkg/mcp/tools`
2. **重新生成**：`make gen-all`
3. **检查依赖**：确保所有依赖包都已安装

## 统计信息

当前生成状态（2026-03-08）：

- **Spec文件**：233个，description填充率约99%
- **API函数**：233个，全部包含完整描述注释
- **MCP工具**：233个，全部包含完整中文描述
- **工具模块**：28个，覆盖所有Tushare API分类

## 最佳实践

1. **定期更新**：当Tushare更新API时，运行`make gen-all`同步更新
2. **版本控制**：生成的内容不应手动编辑，以便重新生成
3. **测试验证**：重新生成后运行`make test`确保没有破坏现有功能
4. **增量生成**：如果只修改了部分API，可以单独运行对应的生成命令

## 相关文件

- `cmd/spec-gen/main.go` - Spec文件生成器
- `cmd/generator/main.go` - API代码生成器
- `cmd/gen-mcp-tools/main.go` - MCP工具生成器
- `internal/gen/templates/api.go.tmpl` - API代码模板
- `Makefile` - 构建和生成命令
