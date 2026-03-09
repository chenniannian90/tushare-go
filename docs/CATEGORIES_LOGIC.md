# Categories 处理逻辑说明

## 概述

MCP 服务器支持根据配置中的 `Categories` 字段动态创建服务实例。每个服务实例可以包含特定的工具集，并注册到独立的 HTTP 路径。

## 核心逻辑

### 1. 服务创建规则 (server.go:18-59)

```go
if len(svcConfig.Categories) == 0 || (len(svcConfig.Categories) == 1 && svcConfig.Categories[0] == name) {
    // 情况 1 & 2: 创建一个服务实例
} else {
    // 情况 3: 为每个 category 创建独立服务实例
}
```

### 2. 三种情况

#### 情况 1: 空 Categories
```json
{
  "name": "all",
  "path": "/",
  "categories": []
}
```
**结果：**
- 创建 1 个服务实例
- 服务名：`all`
- 路径：`/`
- 工具：**所有工具** (registerAllTools)

#### 情况 2: 单个 Category 匹配 Name
```json
{
  "name": "bond",
  "path": "/bond",
  "categories": ["bond"]
}
```
**结果：**
- 创建 1 个服务实例
- 服务名：`bond`
- 路径：`/bond`
- 工具：只有 `bond` 相关的工具

#### 情况 3: 多个 Categories
```json
{
  "name": "stock",
  "path": "/stock",
  "categories": ["stock_basic", "stock_market", "stock_financial"]
}
```
**结果：**
- 创建 3 个独立服务实例
- 服务名：`stock_stock_basic`, `stock_stock_market`, `stock_stock_financial`
- 路径：`/stock/stock_basic`, `/stock/stock_market`, `/stock/stock_financial`
- 工具：每个服务只包含对应 category 的工具

### 3. HTTP 路由注册 (server.go:118-194)

```go
// 跳过 "all" 服务，为所有其他服务注册 HTTP 处理器
for name, svc := range s.services {
    if name == "all" {
        continue
    }
    mux.Handle(svc.config.Path, corsHandler)
}
```

**注意：** "all" 服务在 HTTP 模式下不会注册，因为已经有细粒度的服务可用。

## 实际配置示例

### config.example.json 分析

**会创建多个子服务的配置：**
1. `stock` (8 categories) → 8 个独立服务

**会创建单个服务的配置：**
1. `all` ([]) → 1 个服务（所有工具）
2. `bond` (["bond"]) → 1 个服务
3. `fund` (["fund"]) → 1 个服务
4. `index` (["index"]) → 1 个服务
5. `hk_stock` (["hk_stock"]) → 1 个服务
6. `us_stock` (["us_stock"]) → 1 个服务
7. `etf` (["etf"]) → 1 个服务

## 工具注册 (tools.go:69-94)

```go
func registerToolsForService(server *mcpsdk.Server, categories []string, client *sdk.Client) error {
    if len(categories) == 0 {
        return registerAllTools(server, client)  // 注册所有工具
    }

    for _, category := range categories {
        registrar := toolRegistry[category]
        registrar(server, client)  // 注册特定 category 的工具
    }
    return nil
}
```

## 使用场景

### STDIO 模式
- 使用 "all" 服务
- 只能有一个服务实例
- 包含所有工具

### HTTP 模式
- 使用细粒度的服务
- 每个 path 对应特定的工具集
- 客户端可以选择访问特定的 path

## 优势

1. **灵活性**：可以根据需要组织工具到不同的服务
2. **隔离性**：每个服务实例有独立的工具集
3. **细粒度访问控制**：可以为不同的 category 设置不同的认证策略
4. **清晰的 URL 结构**：URL 结构反映了工具的组织方式

## 测试

已通过单元测试验证所有三种情况：
- 空 categories
- 单个 category 匹配 name
- 多个 categories

测试文件：`cmd/mcp-server/server_test.go`
