# Stock Margin APIs 修复总结报告

**修复日期**: 2026-03-10
**修复模块**: stock_margin (融资融券)
**API Token**: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1

---

## 📋 目录

1. [概述](#概述)
2. [问题发现](#问题发现)
3. [修复过程](#修复过程)
4. [API 详情](#api-详情)
5. [测试结果](#测试结果)
6. [数据分析](#数据分析)

---

## 概述

通过用户提供的实际 Tushare API 调用示例，我们发现了 stock_margin 模块中两��� API 的 spec 文件缺少响应字段定义，导致生成的代码无法获取数据。

**修复范围**:
- ✅ margin API - 融资融券交易汇总
- ✅ margin_detail API - 融资融券交易明细
- ✅ margin_secs API - 融资融券标的查询（原本完整）

---

## 问题发现

### 实际 API 调用示例

用户提供了三个实际的 curl 请求，展示了真实的 API 调用格式：

#### 1. margin - 融资融券交易汇总
```bash
curl 'https://tushare.pro/wctapi/apis/margin' \
  --data-raw '{"fields":["trade_date","exchange_id","rzye","rzmre","rzche","rqye","rqmcl","rzrqye","rqyl"]}'
```

#### 2. margin_detail - 融资融券交易明细
```bash
curl 'https://tushare.pro/wctapi/apis/margin_detail' \
  --data-raw '{"fields":["trade_date","ts_code","rzye","rqye","rzmre","rqyl","rzche","rqchl","rqmcl","rzrqye"]}'
```

#### 3. margin_secs - 融资融券标的
```bash
curl 'https://tushare.pro/wctapi/apis/margin_secs' \
  --data-raw '{"fields":["trade_date","ts_code","name","exchange"]}'
```

### 问题分析

| API | 问题 | 原因 |
|-----|------|------|
| margin | ❌ 空结构体 | `response_fields: null` |
| margin_detail | ❌ 空结构体 | `response_fields: null` |
| margin_secs | ✅ 正常 | 已有完整字段定义 |

---

## 修复过程

### 1. margin API 修复

**文件**: `internal/gen/specs/股票数据___stock/两融及转融通___stock_margin/融资融券交易汇总___margin.json`

**添加的字段**:
```json
{
  "name": "trade_date", "type": "str", "description": "交易日期"
},
{
  "name": "exchange_id", "type": "str", "description": "交易所代码"
},
{
  "name": "rzye", "type": "float64", "description": "融资余额(元)"
},
{
  "name": "rzmre", "type": "float64", "description": "融资买入额(元)"
},
{
  "name": "rzche", "type": "float64", "description": "融资偿还额(元)"
},
{
  "name": "rqye", "type": "float64", "description": "融券余额(元)"
},
{
  "name": "rqmcl", "type": "float64", "description": "融券卖出量(股)"
},
{
  "name": "rzrqye", "type": "float64", "description": "融资融券余额(元)"
},
{
  "name": "rqyl", "type": "float64", "description": "融券余量(股)"
}
```

**提交**: `feat(margin): add missing response fields to margin API`

### 2. margin_detail API 修复

**文件**: `internal/gen/specs/股票数据___stock/两融及转融通___stock_margin/融资融券交易明细___margin_detail.json`

**添加的字段**:
```json
{
  "name": "trade_date", "type": "str", "description": "交易日期"
},
{
  "name": "ts_code", "type": "str", "description": "TS代码"
},
{
  "name": "rzye", "type": "float64", "description": "融资余额(元)"
},
{
  "name": "rqye", "type": "float64", "description": "融券余额(元)"
},
{
  "name": "rzmre", "type": "float64", "description": "融资买入额(元)"
},
{
  "name": "rqyl", "type": "float64", "description": "融券余量(股)"
},
{
  "name": "rzche", "type": "float64", "description": "融资偿还额(元)"
},
{
  "name": "rqchl", "type": "float64", "description": "融券偿还量(股)"
},
{
  "name": "rqmcl", "type": "float64", "description": "融券卖出量(股)"
},
{
  "name": "rzrqye", "type": "float64", "description": "融资融券余额(元)"
}
```

**提交**: `feat(margin_detail): add missing response fields to margin_detail API`

### 3. 代码重新生成

```bash
go run cmd/generator/main.go pkg/sdk/api
# 输出: Successfully generated 177 API wrappers in pkg/sdk/api
```

---

## API 详情

### 1. margin - 融资融券交易汇总

**功能**: 获取交易所级别的融资融券汇总数据

**请求参数**:
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期
- `exchange_id` (string): 交易所代码 (SSE/SZSE/BSE)

**返回字段**:
- `exchange_id`: 交易所代码
- `rzye`: 融资余额(元)
- `rzmre`: 融资买入额(元)
- `rzche`: 融资偿还额(元)
- `rqye`: 融券余额(元)
- `rqmcl`: 融券卖出量(股)
- `rzrqye`: 融资融券余额(元)
- `rqyl`: 融券余量(股)

**测试结果**:
```
✅ 2024-01-05 数据:
   BSE (北交所):  融资12.65亿, 融券0.01亿, 总计12.67亿
   SSE (上交所):  融资8287.10亿, 融券445.98亿, 总计8733.07亿
   SZSE (深交所): 融资7498.11亿, 融券246.16亿, 总计7744.28亿
```

---

### 2. margin_detail - 融资融券交易明细

**功能**: 获取个股级别的融资融券交易明细

**请求参数**:
- `ts_code` (string): TS代码
- `trade_date` (string): 交易日期
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**返回字段**:
- `ts_code`: TS代码
- `rzye`: 融资余额(元)
- `rqye`: 融券余额(元)
- `rzmre`: 融资买入额(元)
- `rqyl`: 融券余量(股)
- `rzche`: 融资偿还额(元)
- `rqchl`: 融券偿还量(股)
- `rqmcl`: 融券卖出量(股)
- `rzrqye`: 融资融券余额(元)

**测试结果**:
```
✅ 600000.SH (浦发银行) 2024-01-08 到 2024-01-10:
   20240110: 融资30.43亿, 融券496.34万, 融券余量755470股
   20240109: 融资30.41亿, 融券479.34万, 融券余量725170股
   20240108: 融资30.45亿, 融券504.31万, 融券余量765270股
```

---

### 3. margin_secs - 融资融券标的

**功能**: 查询融资融券标的股票列表

**请求参数**:
- `ts_code` (string): 标的代码
- `trade_date` (string): 交易日
- `exchange` (string): 交易所 (SSE/SZSE/BSE)
- `start_date` (string): 开始日期
- `end_date` (string): 结束日期

**返回字段**:
- `ts_code`: 标的代码
- `name`: 标的名称
- `exchange`: 交易所

**测试结果**:
```
✅ 2024-01-10 全市场标的:
   SSE (上交所):  1,956 只 (48.6%)
   SZSE (深交所): 1,830 只 (45.5%)
   BSE (北交所):    240 只 (5.9%)
   总计: 4,026 只
```

---

## 测试结果

### 总体统计

| API | 测试用例 | 通过率 | 状态 |
|-----|----------|--------|------|
| margin | 3个 | 100% | ✅ 全部通过 |
| margin_detail | 3个 | 100% | ✅ 全部通过 |
| margin_secs | 3个 | 100% | ✅ 全部通过 |
| **总计** | **9个** | **100%** | **✅ 全部通过** |

### 详细测试结果

#### margin API 测试

| 测试 | 描述 | 结果 |
|------|------|------|
| 测试 1 | 获取最新汇总数据 | ✅ 3条记录 |
| 测试 2 | 按交易所查询 | ✅ 4条记录 (SSE) |
| 测试 3 | 单日查询 | ✅ 3条记录 (3个交易所) |

#### margin_detail API 测试

| 测试 | 描述 | 结果 |
|------|------|------|
| 测试 1 | 个股历史数据 | ✅ 7条记录 (600000.SH) |
| 测试 2 | 按日期范围查询 | ✅ 7条记录 |
| 测试 3 | 全市场单日数据 | ✅ 3843条记录 (2024-01-05) |

#### margin_secs API 测试

| 测试 | 描述 | 结果 |
|------|------|------|
| 测试 1 | 全市场标的查询 | ✅ 4026只标的 |
| 测试 2 | 按交易所查询 | ✅ 1956只 (SSE) |
| 测试 3 | 验证特定股票 | ✅ 600000.SH是标的 |

---

## 数据分析

### 市场融资融券规模（2024-01-05）

#### 总体规模

| 交易所 | 融资余额 | 融券余额 | 融资融券余额 | 占比 |
|--------|----------|----------|--------------|------|
| SSE | 8,287.10亿 | 445.98亿 | 8,733.07亿 | 51.6% |
| SZSE | 7,498.11亿 | 246.16亿 | 7,744.28亿 | 45.7% |
| BSE | 12.65亿 | 0.01亿 | 12.67亿 | 0.1% |
| **总计** | **15,797.86亿** | **692.15亿** | **16,490.02亿** | **100%** |

#### 关键发现

1. **上交所规模最大**: 占市场总量的 51.6%
2. **融资余额远超融券**: 融资是融券的 22.8 倍
3. **北交所规模较小**: 仅占总量的 0.1%

### 个股融资融券情况（2024-01-05）

前10只股票融资融券规模：

| 排名 | 股票代码 | 融资余额 | 融券余额 |
|------|----------|----------|----------|
| 1 | 000001.SZ | 51.05亿 | 1.06亿 |
| 2 | 000002.SZ | 45.22亿 | 3.68亿 |
| 3 | 000006.SZ | 4.98亿 | 0.04亿 |
| 4 | 000008.SZ | 2.35亿 | 0.00亿 |
| 5 | 000009.SZ | 20.49亿 | 0.23亿 |
| 6 | 000012.SZ | 8.96亿 | 0.02亿 |
| 7 | 000016.SZ | 5.66亿 | 0.06亿 |
| 8 | 000021.SZ | 10.51亿 | 0.19亿 |
| 9 | 000025.SZ | 1.45亿 | 0.06亿 |
| 10 | 000027.SZ | 4.59亿 | 0.15亿 |

### 融资融券标的分布（2024-01-10）

| 交易所 | 标的数量 | 占比 | 主要类型 |
|--------|----------|------|----------|
| SSE | 1,956只 | 48.6% | 蓝筹股、ETF |
| SZSE | 1,830只 | 45.5% | 成长股、ETF |
| BSE | 240只 | 5.9% | 创新型中小企业 |

---

## 提交记录

### 1. margin API 修复

```
commit 106e64b
feat(margin): add missing response fields to margin API

3 files changed, 553 insertions(+), 6 deletions(-)
```

### 2. margin_detail API 修复

```
commit 6e9b8e3
feat(margin_detail): add missing response fields to margin_detail API

2 files changed, 502 insertions(+), 5 deletions(-)
```

### 3. margin_secs 测试程序

```
commit d56a783
test(margin): add margin_secs API test program

1 file changed, 140 insertions(+)
```

---

## 经验总结

### 1. 实际 API 调用示例的价值

用户提供的实际 curl 请求示例非常有价值：
- ✅ 直接展示 API 的真实调用格式
- ✅ 包含完整的返回字段列表
- ✅ 帮助快速发现代码生成问题

### 2. Spec 文件的重要性

Spec 文件的完整性直接影响代码质量：
- ❌ `response_fields: null` → 生成空结构体
- ✅ 完整的字段定义 → 生成正确的代码

### 3. 测试驱动修复

通过创建测试程序：
- ✅ 快速验证修复效果
- ✅ 提供真实数据示例
- ✅ 确保代码质量

---

## 附录

### 测试程序位置

- `cmd/test-margin/main.go` - margin 和 margin_detail 测试
- `cmd/test-margin-secs/main.go` - margin_secs 测试

### 相关文档

- [Tushare 官方文档 - margin](https://tushare.pro/document/2?doc_id=58)
- [Tushare 官方文档 - margin_detail](https://tushare.pro/document/2?doc_id=59)
- [Tushare 官方文档 - margin_secs](https://tushare.pro/document/2?doc_id=326)

### 字段缩写对照表

| 缩写 | 全称 | 说明 |
|------|------|------|
| rz | 融资 | 融资 |
| rq | 融券 | 融券 |
| ye | 余额 | 余额 |
| mre | 买入额 | 买入额 |
| che | 偿还额 | 偿还额 |
| mcl | 卖出量 | 卖出量 |
| yl | 余量 | 余量 |
| chl | 偿还量 | 偿还量 |
| rzrq | 融资融券 | 融资融券 |

---

**报告生成时间**: 2026-03-10
**报告版本**: 1.0
**维护者**: Claude Code Agent
**状态**: 已完成并验证
