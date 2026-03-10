# 爬虫实现 API 说明

## 概述

本项目中有部分 API 使用爬虫实现，这些 API **不应该被代码生成工具自动生成或覆盖**。

## 爬虫实现 API 列表

### 股票行情数据 (stock_market)

1. **RealtimeQuote** - 实时 Tick（爬虫）
   - 文件: `realtime_quote.go`
   - 实现: `realtime_crawler.go`
   - 功能: 获取单个股票的实时行情数据
   - 数据源: 新浪财经、东方财富

2. **RealtimeList** - 实时排名（爬虫）
   - 文件: `realtime_list.go`
   - 实现: `realtime_list_crawler.go`
   - 功能: 获取全市场股票实时排名列表
   - 数据源: 新浪财经、东方财富

3. **RealtimeTick** - 实时成交（爬虫）
   - 文件: `realtime_tick.go`
   - 实现: `realtime_tick_crawler.go`
   - 功能: 获取股票的分笔成交数据
   - 数据源: 新浪财经、东方财富

## 标识方式

所有爬虫实现的文件都在文件头部包含以下标记：

```go
// CRAWLER IMPLEMENTATION - DO NOT AUTO GENERATE
// 本文件包含爬虫实现的 API，需要手动维护，禁止代码生成工具覆盖
```

## 重要特性

1. **不依赖 Tushare HTTP 接口**: 这些 API 直接爬取第三方网站数据
2. **无需 Token**: 可以在没有 Tushare Token 的情况下使用
3. **实时性**: 数据直接来自实时行情源，延迟较低
4. **维护性**: 爬虫实现需要定期维护，因为数据源可能会变更
5. **合规性**: 使用爬虫数据时需要遵守相关网站的使用条款和 robots.txt 规定

## 使用注意事项

1. **交易时间**: 仅在交易时间段内返回有效数据
2. **数据格式**: 支持标准 tushare 格式代码（如 000001.SZ）和纯数字代码
3. **错误处理**: 爬虫失败时会返回错误，需要适当的错误处理
4. **频率限制**: 避免过于频繁的请求，以免被数据源封禁

## 代码生成工具配置

如果您使用代码生成工具，请确保配置以下排除规则：

### tushare-gen 工具

在生成代码时，排除以下文件：
- `realtime_quote.go`
- `realtime_list.go`
- `realtime_tick.go`
- `realtime_crawler.go`
- `realtime_list_crawler.go`
- `realtime_tick_crawler.go`

### 示例配置

```yaml
# 代码生成工具配置
exclude:
  - pkg/sdk/api/stock_market/realtime_quote.go
  - pkg/sdk/api/stock_market/realtime_list.go
  - pkg/sdk/api/stock_market/realtime_tick.go
  - pkg/sdk/api/stock_market/realtime_crawler.go
  - pkg/sdk/api/stock_market/realtime_list_crawler.go
  - pkg/sdk/api/stock_market/realtime_tick_crawler.go
```

## 维护说明

当数据源网站发生变化时，需要更新对应的爬虫实现文件：

1. **URL 变更**: 更新爬虫函数中的数据源 URL
2. **HTML 结构变更**: 更新正则表达式或解析逻辑
3. **新增字段**: 更新数据结构体和解析代码
4. **反爬虫机制**: 可能需要添加请求头、延迟等策略

## 相关文档

- API 目录: `docs/api-directory.yaml`
- API 目录 JSON: `docs/api-directory.json`
- 示例代码: `cmd/examples/crawler/main.go`

---

**最后更新**: 2026-03-09
**维护者**: 开发团队
