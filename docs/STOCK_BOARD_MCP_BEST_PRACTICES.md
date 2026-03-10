# Stock Board MCP 工具最佳实践

本文档提供了 stock_board 模块中 MCP 工具的使用说明和最佳实践，基于实际测试结果整理。

## 测试概述

**测试时间**: 2026-03-09
**测试工具数量**: 22个
**测试环境**: Tushare API

## 工具分类

### ✅ 可用工具 (3/22)

| 工具名称 | 状态 | 数据量 | 说明 |
|---------|------|--------|------|
| `stock_board.hm_detail` | ✅ 有数据 | 236条 | 游资营业部明细 |
| `stock_board.stk_auction` | ✅ 有数据 | 8000条 | 集合竞价（最佳实践见下文）|
| `stock_board.tdx_member` | ⚠️ 无数据 | 0条 | 同花顺板块成分股 |

### ❌ 权限受限工具 (16/22)

以下工具需要更高权限的 Tushare 账户：

- `stock_board.dc_daily` - 东财板块日线（每天最多访问2次）
- `stock_board.dc_hot` - 东财热点（每天最多访问2次）
- `stock_board.dc_index` - 东财板块指数
- `stock_board.dc_member` - 东财板块成分股
- `stock_board.hm_list` - 游资营业部列表（每天最多访问2次）
- `stock_board.kpl_concept_cons` - 看涨概念成分
- `stock_board.kpl_list` - 看涨榜列表
- `stock_board.limit_cpt_list` - 涨跌停板块统计
- `stock_board.limit_list_d` - 大宗交易（每天最多访问1次）
- `stock_board.limit_list_ths` - 同花顺涨跌停
- `stock_board.limit_step` - 连板步长
- `stock_board.tdx_daily` - 同花顺日线（每天最多访问2次）
- `stock_board.ths_daily` - 同花顺日线
- `stock_board.ths_hot` - 同花顺热点（每天最多访问2次）
- `stock_board.ths_index` - 同花顺指数
- `stock_board.ths_member` - 同花顺成分股

### ⚠️ 数据类型错误 (3/22)

以下工具存在数据类型转换问题，需要修复：

- `stock_board.tdx_index` - 无效的 total_mv 类型
- `stock_board.top_inst` - 无效的 sell 类型
- `stock_board.top_list` - 无效的 turnover_rate 类型

---

## stk_auction 工具最佳实践

### 工具信息

- **工具名称**: `stock_board.stk_auction`
- **功能**: 获取当日个股和ETF的集合竞价成交情况
- **可用时间**: 每天9:25-9:29分之间可获取当日数据
- **权限**: 无特殊权限要求

### 推荐使用方式

#### 方式1: 不指定日期（推荐）⭐

```json
{
  "name": "stock_board.stk_auction",
  "arguments": {}
}
```

**优点**:
- 返回数据量最大（约8000条）
- 自动获取最新可用数据
- 避免日期格式错误

**返回数据示例**:
```json
{
  "data": [
    {
      "ts_code": "920992.BJ",
      "trade_date": "20260309",
      "vol": 0,
      "price": 0,
      "amount": 0,
      "pre_close": 16.42,
      "turnover_rate": 0,
      "volume_ratio": 0,
      "float_share": 4851.8
    }
  ],
  "total": 8000
}
```

#### 方式2: 指定当前日期

```json
{
  "name": "stock_board.stk_auction",
  "arguments": {
    "trade_date": "20260309"
  }
}
```

**返回数据量**: 约5477条

**注意事项**:
- 日期格式必须是 `YYYYMMDD`
- 使用当天或最近日期
- 历史日期可能导致类型错误

### 不推荐的参数

#### ❌ 使用历史日期

```json
{
  "name": "stock_board.stk_auction",
  "arguments": {
    "trade_date": "20250307"
  }
}
```

**问题**: 会返回 "无效的 turnover_rate 类型" 错误

#### ❌ 使用日期范围

```json
{
  "name": "stock_board.stk_auction",
  "arguments": {
    "start_date": "20240301",
    "end_date": "20240310"
  }
}
```

**问题**: 返回 0 条数据（集合竞价数据不保存历史记录）

### 字段说明

| 字段 | 类型 | 说明 | 示例值 |
|------|------|------|--------|
| ts_code | string | 股票代码 | "600000.SH" |
| trade_date | string | 交易日期 | "20260309" |
| vol | number | 成交量 | 891400 |
| price | number | 竞价价格 | 9.0 |
| amount | number | 成交额 | 876242 |
| pre_close | number | 昨收价 | 9.89 |
| turnover_rate | number | 换手率 | 0.00267641 |
| volume_ratio | number | 量比 | 0.824503 |
| float_share | number | 流通股本 | 3330580 |

### 使用场景

1. **盘前分析**: 9:25-9:29之间获取集合竞价数据
2. **市场情绪**: 通过集合竞价情况判断市场开盘情绪
3. **个股监控**: 监控特定股票的集合竞价情况
4. **数据分析**: 研究集合竞价与当日表现的关系

### 限制说明

1. **时间限制**: 仅在交易日9:25-9:29之间有完整数据
2. **数据时效**: 集合竞价数据通常不保存历史记录
3. **数据完整性**: 部分股票可能无集合竞价数据（显示为0）

---

## hm_detail 工具使用说明

### 工具信息

- **工具名称**: `stock_board.hm_detail`
- **功能**: 获取游资营业部交易明细
- **权限**: 无特殊权限要求

### 推荐参数

```json
{
  "name": "stock_board.hm_detail",
  "arguments": {
    "trade_date": "20240308"
  }
}
```

**返回数据量**: 约236条

### 可选参数

- `ts_code`: 股票代码
- `hm_name`: 游资名称
- `start_date`: 开始日期
- `end_date`: 结束日期

---

## 使用建议

### 1. 权限升级

如需使用受限工具，请访问 https://tushare.pro 升级账户权限。

### 2. 数据类型修复

对于存在类型错误的工具，等待 SDK 修复或手动调整字段类型。

### 3. 测试环境

在实际使用前，建议先用小参数测试 API 可用性：

```bash
# 测试 stk_auction
go run cmd/api-test/test_stk_auction_improved.go
```

### 4. 错误处理

常见错误及处理方式：

1. **ACCESS_DENIED**: 升级 Tushare 权限
2. **无效的 xxx 类型**: 等待 SDK 修复或使用其他工具
3. **无数据返回**: 检查参数或尝试其他日期

---

## 更新日志

- **2026-03-09**: 初始版本，添加 stk_auction 最佳实践
- 基于 22 个工具的实际测试结果

---

## 相关文档

- [MCP 工具总览](./MCP_TOOLS.md)
- [API 目录](./api-directory.json)
- [测试报告](../tests/stock-board-test-report.md)
- [stk_auction 详细测试](../tests/stk_auction_improved_report.md)

---

**最后更新**: 2026-03-09
**维护者**: tushare-go 项目组
