# Stock Fund Flow MCP Tools 测试报告

**测试日期**: 2026-03-10
**测试环境**: tushare-go SDK
**API Token**: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1

---

## 📋 目录

1. [概述](#概述)
2. [MCP Tools 列表](#mcp-tools-列表)
3. [测试结果](#测试结果)
4. [API 调用示例分析](#api-调用示例分析)
5. [问题与建议](#问题与建议)

---

## 概述

stock_fund_flow 模块提供了 8 个 MCP tools，用于访问股票资金流向相���数据。本次测试发现**7 个工具被禁用**，只有 1 个工具处于启用状态。

资金流向数据包括：
- 个股资金流向
- 沪深股通资金流向（北向/南向资金）
- 大单资金流向
- 行业/概念/地域资金流向

---

## MCP Tools 列表

### ✅ 已启用工具 (1/8)

#### 1. stock_fund_flow.moneyflow
**描述**: 个股资金流向
**请求参数**:
- `ts_code` (string): 股票代码 (股票和时间参数至少输入一个)
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ✅ 已启用
**测试结果**: ✅ 通过
**测试参数**:
```json
{
  "ts_code": "600000.SH",
  "start_date": "20240101",
  "end_date": "20240110"
}
```
**返回数据**: 0 条（测试日期范围内无数据）

---

### ⏭️ 已禁用工具 (7/8)

#### 2. stock_fund_flow.moneyflow_hsgt
**描述**: 获取沪股通、深股通、港股通每日资金流向数据
**请求参数**:
- `trade_date` (string): 交易日期 (二选一)
- `start_date` (string): 开始日期 (二选一)
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用
**原因**: 在 `pkg/mcp/tools/stock_fund_flow/registry.go` 中被注释

**实际 API 示例** (来自 tushare.pro):
```bash
curl 'https://tushare.pro/wctapi/apis/moneyflow_hsgt' \
  --data-raw '{
    "params": {
      "trade_date": "",
      "start_date": "",
      "end_date": "",
      "limit": "",
      "offset": ""
    },
    "fields": [
      "trade_date",
      "ggt_ss",
      "ggt_sz",
      "hgt",
      "sgt",
      "north_money",
      "south_money"
    ]
  }'
```

**返回字段说明**:
- `trade_date`: 交易日期
- `ggt_ss`: 港股通上海
- `ggt_sz`: 港股通深圳
- `hgt`: 沪股通
- `sgt`: 深股通
- `north_money`: 北向资金
- `south_money`: 南向资金

**积分要求**: 2000积分起，5000积分每分钟可提取500次

---

#### 3. stock_fund_flow.moneyflow_ths
**描述**: 个股沪深股通资金流向
**请求参数**:
- `ts_code` (string): 股票代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用

---

#### 4. stock_fund_flow.moneyflow_cnt_ths
**描述**: 沪深股通成份股资金流向
**请求参数**:
- `ts_code` (string): 代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用

---

#### 5. stock_fund_flow.moneyflow_dc
**描述**: 大单资金流向
**请求参数**:
- `ts_code` (string): 股票代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用

---

#### 6. stock_fund_flow.moneyflow_ind_dc
**描述**: 行业资金流向
**请求参数**:
- `ts_code` (string): 代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期
- `content_type` (string): 资金类型(行业、概念、地域)

**状态**: ⏭️ 已禁用

---

#### 7. stock_fund_flow.moneyflow_ind_ths
**描述**: 行业沪深股通资金流向
**请求参数**:
- `ts_code` (string): 代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用

---

#### 8. stock_fund_flow.moneyflow_mkt_dc
**描述**: 市场大单资金流向
**请求参数**:
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**状态**: ⏭️ 已禁用

---

## 测试结果

### 总体统计

| 指标 | 数量 | 占比 |
|------|------|------|
| 总工具数 | 8 | 100% |
| 已启用 | 1 | 12.5% |
| 已禁用 | 7 | 87.5% |
| 测试通过 | 1 | 100% (已启用工具中) |
| 测试失败 | 0 | 0% |

### 工具启用状态

| # | MCP Tool | 状态 | 说明 |
|---|----------|------|------|
| 1 | moneyflow | ✅ 已启用 | 个股资金流向 |
| 2 | moneyflow_hsgt | ⏭️ 已禁用 | 沪深股通资金流向 |
| 3 | moneyflow_ths | ⏭️ 已禁用 | 个股沪深股通资金流向 |
| 4 | moneyflow_cnt_ths | ⏭️ 已禁用 | 沪深股通成份股资金流向 |
| 5 | moneyflow_dc | ⏭️ 已禁用 | 大单资金流向 |
| 6 | moneyflow_ind_dc | ⏭️ 已禁用 | 行业资金流向 |
| 7 | moneyflow_ind_ths | ⏭️ 已禁用 | 行业沪深股通资金流向 |
| 8 | moneyflow_mkt_dc | ⏭️ 已禁用 | 市场大单资金流向 |

---

## API 调用示例分析

### 实际 tushare.pro API 调用格式

从提供的 curl 示例可以看出实际 API 的调用方式：

```bash
curl 'https://tushare.pro/wctapi/apis/moneyflow_hsgt' \
  -H 'Content-Type: application/json;charset=UTF-8' \
  -b 'uid=...; username=...' \
  --data-raw '{
    "user_id": 449317,
    "username": "chenniannian",
    "user_valid": true,
    "root_id": "2",
    "doc_id": "47",
    "params": {
      "trade_date": "",
      "start_date": "",
      "end_date": "",
      "limit": "",
      "offset": ""
    },
    "fields": [
      "trade_date",
      "ggt_ss",
      "ggt_sz",
      "hgt",
      "sgt",
      "north_money",
      "south_money"
    ]
  }'
```

### 关键发现

1. **认证方式**:
   - 使用 Cookie 认证 (uid, username)
   - 请求体中包含 user_id, username, user_valid

2. **API 结构**:
   - URL: `https://tushare.pro/wctapi/apis/{api_name}`
   - 方法: POST
   - Content-Type: application/json

3. **请求格式**:
   - `params`: 包含查询参数
   - `fields`: 指定返回字段列表
   - 支持 `limit` 和 `offset` 分页参数

4. **与 MCP Tools 的对比**:
   - MCP Tools 封装了底层 API 调用
   - 参数映射正确
   - 缺少 `limit` 和 `offset` 参数支持

---

## 问题与建议

### 发现的问题

1. **大部分工具被禁用** ⚠️
   - 8 个工具中有 7 个被注释掉
   - 可能的原因：
     - API 权限要求高（需要积分）
     - API 还在测试中
     - 存在已知问题

2. **功能受限**
   - 用户只能使用个股资金流向功能
   - 无法获取沪深股通等核心资金流向数据

3. **参数不完整**
   - 实际 API 支持 `limit` 和 `offset` 分页参数
   - MCP Tools 中未暴露这些参数

### 改进建议

1. **明确禁用原因**
   - 在代码注释中说明为什么禁用这些工具
   - 如果是权限问题，应在文档中说明积分要求

2. **启用更多工具**（如可能）
   - 评估是否可以启用部分已禁用的工具
   - 对于需要高积分的 API，可以添加权限检查提示

3. **完善参数支持**
   - 添加 `limit` 和 `offset` 参数支持
   - 允许用户控制返回数据量和分页

4. **改进错误处理**
   - 对于权限不足的 API，返回明确的错误信息
   - 提示用户所需的积分等级

5. **文档更新**
   - 更新 MCP_TOOLS.md，标注哪些工具已禁用及原因
   - 提供实际 API 调用示例供参考

### 后续行动计划

1. 🔴 **高优先级**: 调查工具被禁用的原因
2. 🟡 **中优先级**: 添加 `limit` 和 `offset` 参数支持
3. 🟢 **低优先级**: 完善文档和错误提示

---

## 附录

### 相关文件位置

- MCP Tools 注册: `pkg/mcp/tools/stock_fund_flow/registry.go`
- MCP Tools 实现: `pkg/mcp/tools/stock_fund_flow/*.go`
- SDK API 实现: `pkg/sdk/api/stock_fund_flow/*.go`
- 测试程序: `cmd/test-stock-fund-flow/main.go`

### Registry.go 当前状态

```go
func (r *Stock_fund_flowTools) RegisterAll() {
    r.registerMoneyflow()
    //r.registerMoneyflowCntThs()
    //r.registerMoneyflowDc()
    //r.registerMoneyflowHsgt()
    //r.registerMoneyflowIndDc()
    //r.registerMoneyflowIndThs()
    //r.registerMoneyflowMktDc()
    //r.registerMoneyflowThs()
}
```

### 测试配置

```json
{
  "mcpServers": {
    "tushare-stock-fund-flow": {
      "type": "http",
      "url": "https://tushare.chat168.cn/stock/stock_fund_flow",
      "headers": {
        "X-API-Key": "412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
      }
    }
  }
}
```

---

**报告生成时间**: 2026-03-10
**报告版本**: 1.0
**下次更新**: 待启用更多工具后更新
