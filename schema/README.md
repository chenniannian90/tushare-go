# Tushare API Schema

本目录包含 Tushare Pro API 的完整结构化定义。

## 📁 文件说明

### `api_schema.yaml`
完整的 API 接口定义文件，包含：
- API 分类（股票、ETF、指数、基金等）
- 子分类（基础数据、行情数据、财务数据等）
- 具体接口信息（文档ID、接口名称、文档链接）

### `parse_api.py`
Python 脚本，用于从 Tushare Pro 官方文档网站解析并生成 `api_schema.yaml`

### `types.go`
Go 语言类型定义，提供 YAML schema 的类型安全��问接口

## 📊 API 统计

| 分类 | 接口数量 |
|------|----------|
| 股票数据 | 111 |
| ETF专题 | 9 |
| 指数专题 | 19 |
| 公募基金 | 8 |
| **总计** | **147** |

## 🚀 使用方法

### 在 Go 代码中使用

```go
package main

import (
    "fmt"
    "github.com/chenniannian90/tushare-go/schema"
)

func main() {
    // 加载 schema
    s, err := schema.LoadSchema()
    if err != nil {
        panic(err)
    }

    // 获取总 API 数量
    fmt.Printf("Total APIs: %d\n", s.TotalAPIs())

    // 列出所有分类
    categories := s.ListCategories()
    for _, cat := range categories {
        fmt.Println(cat)
    }

    // 根据接口名称查找 API
    api, err := s.GetAPIByName("stock_basic")
    if err == nil {
        fmt.Printf("API: %s, URL: %s\n", api.Name, api.URL)
    }

    // 获取特定分类的所有 API
    stockAPIs, err := s.GetAPIsByCategory("stock")
    if err == nil {
        fmt.Printf("Stock APIs: %d\n", len(stockAPIs))
    }
}
```

### 直接读取 YAML

```yaml
# api_schema.yaml 结构示例
version: "1.0.0"
description: "Tushare Pro API Schema"
categories:
  - id: stock
    name: "股票数据"
    subcategories:
      - id: basic
        name: "基础数据"
        apis:
          - doc_id: 25
            name: "股票列表"
            api_name: stock_basic
            url: "https://tushare.pro/document/2?doc_id=25"
```

## 🔧 更新 Schema

当 Tushare Pro 官方文档更新时，可以运行以下命令重新生成 schema：

```bash
python3 schema/parse_api.py
```

这将：
1. 解析最新的 API 文档结构
2. 生成新的 `api_schema.yaml` 文件
3. 打印统计信息

## 📦 依赖

- Go: `gopkg.in/yaml.v3`
- Python: 无外部依赖（仅使用标准库）

## 📝 Schema 结构

```
api_schema.yaml
├── version: 版本号
├── description: 描述
└── categories: 分类列表
    ├── id: 分类ID
    ├── name: 分类名称
    └── subcategories: 子分类列表
        ├── id: 子分类ID
        ├── name: 子分类名称
        └── apis: API列表
            ├── doc_id: 文档ID
            ├── name: API名称
            ├── api_name: 接口名称
            └── url: 文档链接
```

## 💡 使用场景

1. **API 发现**: 快速查找可用的 API 接口
2. **文档生成**: 自动生成 API 文档
3. **代码生成**: 根据 schema 生成客户端代码
4. **接口测试**: 批量测试所有可用接口
5. **版本管理**: 跟踪 API 变更历史
