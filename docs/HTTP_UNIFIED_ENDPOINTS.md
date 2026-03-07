# HTTP 统一端点设计

本文档描述了 MCP 服务器的 HTTP 路由设计，特别是针对 `hk_stock` 模块实现的统一端点功能。

## 概述

HTTP 路由系统支持两种访问方式：

1. **传统方式**: 每个工具对应一个独立的 HTTP 路径
2. **统一端点方式**: 一个模块下的所有工具共享一个 HTTP 路径，通过参数区分具体工具

## hk_stock 模块统一端点

### 传统方式（向后兼容）

```bash
# 访问基础信息
POST /api/v1/hk_stock/hk_basic

# 访问日线数据
POST /api/v1/hk_stock/hk_daily

# 访问交易日历
POST /api/v1/hk_stock/hk_cal

# 访问分钟数据
POST /api/v1/hk_stock/hk_min

# 访问因子数据
POST /api/v1/hk_stock/hk_factor
```

### 统一端点方式（推荐）

```bash
# 统一端点，通过 tool 参数指定具体功能
POST /api/v1/hk_stock?tool=hk_basic
POST /api/v1/hk_stock?tool=hk_daily
POST /api/v1/hk_stock?tool=hk_cal
POST /api/v1/hk_stock?tool=hk_min
POST /api/v1/hk_stock?tool=hk_factor
```

## API 设计

### 请求格式

#### 使��统一端点

```http
POST /api/v1/hk_stock?tool=hk_basic HTTP/1.1
Content-Type: application/json

{
  "ts_code": "00700.HK",
  "list_date": "20240101"
}
```

#### 使用传统路径

```http
POST /api/v1/hk_stock/hk_basic HTTP/1.1
Content-Type: application/json

{
  "ts_code": "00700.HK",
  "list_date": "20240101"
}
```

### 响应格式

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "fields": ["ts_code", "symbol", "name", "area", "industry"],
    "items": [
      {
        "ts_code": "00700.HK",
        "symbol": "00700",
        "name": "腾讯控股",
        "area": "香港",
        "industry": "科技"
      }
    ]
  }
}
```

## 可用工具

### hk_stock 模块工具列表

| 工具名称 | 描述 | 统一端点调用 |
|---------|------|-------------|
| `hk_basic` | 港股基础信息 | `/api/v1/hk_stock?tool=hk_basic` |
| `hk_daily` | 港股日线数据 | `/api/v1/hk_stock?tool=hk_daily` |
| `hk_cal` | 港股交易日历 | `/api/v1/hk_stock?tool=hk_cal` |
| `hk_min` | 港股分钟数据 | `/api/v1/hk_stock?tool=hk_min` |
| `hk_factor` | 港股因子数据 | `/api/v1/hk_stock?tool=hk_factor` |

## 使用示例

### cURL 示例

```bash
# 使用统一端点获取港股基础信息
curl -X POST "http://localhost:8080/api/v1/hk_stock?tool=hk_basic" \
  -H "Content-Type: application/json" \
  -d '{
    "ts_code": "00700.HK",
    "list_date": "20240101"
  }'

# 使用统一端点获取港股日线数据
curl -X POST "http://localhost:8080/api/v1/hk_stock?tool=hk_daily" \
  -H "Content-Type: application/json" \
  -d '{
    "ts_code": "00700.HK",
    "start_date": "20240101",
    "end_date": "20240131"
  }'
```

### Go 客户端示例

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

func main() {
    baseURI := "http://localhost:8080/api/v1/hk_stock"

    // 使用统一端点
    tool := "hk_basic"
    params := url.QueryEscape("tool=" + tool)
    fullURL := baseURI + "?" + params

    request := map[string]interface{}{
        "ts_code":   "00700.HK",
        "list_date": "20240101",
    }

    body, _ := json.Marshal(request)
    resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Printf("Status: %s\n", resp.Status)
}
```

### Python 客户端示例

```python
import requests
import json

base_url = "http://localhost:8080/api/v1/hk_stock"

# 使用统一端点获取港股基础信息
params = {
    "tool": "hk_basic"
}

data = {
    "ts_code": "00700.HK",
    "list_date": "20240101"
}

response = requests.post(
    base_url,
    params=params,
    json=data
)

print(f"Status: {response.status_code}")
print(f"Response: {response.json()}")
```

## 错误处理

### 无效的工具名称

```json
{
  "error": {
    "code": -32601,
    "message": "tool 'invalid_tool' not found in module 'hk_stock'. Available tools: [hk_basic, hk_daily, hk_cal, hk_min, hk_factor]"
  }
}
```

### 缺少工具参数

```json
{
  "error": {
    "code": -32602,
    "message": "Missing required parameter: tool"
  }
}
```

## 技术实现

### 路由结构

```go
type HTTPRoute struct {
    Path        string  // HTTP 路径
    Module      string  // 模块名称
    Tool        string  // 工具名称（"*" 表示统一端点）
    HTTPMethod  string  // HTTP 方法
    Description string  // 描述
    ParamName   string  // 参数名称（对于统一端点为 "tool"）
    QueryParam  bool    // 是否使用查询参数
}
```

### 核心方法

```go
// 获取统一端点列表
func (r *HTTPRouter) GetUnifiedEndpoints() []HTTPRoute

// 获取模块的统一端点
func (r *HTTPRouter) GetModuleUnifiedEndpoint(module string) (*HTTPRoute, bool)

// 解析统一端点调用
func (r *HTTPRouter) ResolveUnifiedEndpoint(path, tool string) (string, error)

// 获取模块下的所有工具
func (r *HTTPRouter) GetModuleTools(module string) []string
```

## 优势

### 统一端点方式的优势

1. **更简洁的 URL**: 减少路径层级，更容易记忆
2. **更好的可扩展性**: 添加新工具时无需修改路由结构
3. **更灵活的配置**: 可以在运行时动态调整可用工具
4. **统一的管理**: 一个模块的所有工具集中在同一个路径下

### 向后兼容性

- 传统路径方式仍然支持
- 现有客户端无需修改
- 两种方式可以同时使用

## 扩展性

目前只有 `hk_stock` 模块实现了统一端点，但这个设计可以扩展到其他模块：

```go
// 可以轻松为其他模块添加统一端点
r.addUnifiedModuleRoutes("stock_basic", []string{
    "stock_basic", "trade_cal", "stock_company", /* ... */
})

r.addUnifiedModuleRoutes("stock_market", []string{
    "daily", "weekly", "monthly", /* ... */
})
```

## 最佳实践

1. **新项目**: 优先使用统一端点方式
2. **现有项目**: 可以逐步迁移，保持向后兼容
3. **文档**: 在 API 文档中明确标注支持的访问方式
4. **测试**: 确保两种方式都能正常工作

## 相关文档

- [MCP 协议规范](https://modelcontextprotocol.io/)
- [HTTP 路由实现](../pkg/mcp/server/http_routes.go)
- [API Key 认证](MCP_AUTH.md)
