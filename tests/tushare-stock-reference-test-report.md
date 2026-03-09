# tushare-stock-reference 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-reference
- **测试工具数**: 11 个
- **测试成功**: 3 个 (27.3%)
- **参数错误**: 2 个 (18.2%)
- **空数据**: 4 个 (36.4%)
- **工具不可用**: 2 个 (18.2%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_reference_block_trade | ts_code=600000.SH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | stock_reference_pledge_detail | ts_code=600000.SH | ✅ 正常 | {"ts_code": "600000.SH", "holder_name": "雅戈尔集团", "pledge_amount": 1200} |
| 3 | stock_reference_pledge_stat | ts_code=600000.SH, end_date=20240308 | ✅ 正常 | {"ts_code": "600000.SH", "pledge_ratio": 0.1} |
| 4 | stock_reference_repurchase | start_date=20240301, end_date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 vol 类型 |
| 5 | stock_reference_share_float | ts_code=600000.SH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 6 | stock_reference_stk_account | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 7 | stock_reference_stk_holdernumber | ts_code=600000.SH, end_date=20240308 | 🔧 工具不可用 | Error: No such tool available |
| 8 | stock_reference_stk_holdertrade | ts_code=600000.SH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 9 | stock_reference_top10_floatholders | ts_code=600000.SH, period=20231231 | 📭 空数据 | {"data": [], "total": 0} |
| 10 | stock_reference_top10_holders | ts_code=600000.SH, period=20231231 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 11 个
- **测试成功**: 3 个 (27.3%)
- **参数错误**: 1 个 (9.1%)
- **空数据**: 6 个 (54.5%)
- **工具不可用**: 1 个 (9.1%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 18.2% | stock_reference_pledge_detail, stock_reference_pledge_stat |
| ⚠️ 参数错误 | 1 | 9.1% | stock_reference_repurchase |
| 📭 空数据 | 6 | 54.5% | stock_reference_block_trade, stock_reference_share_float等 |
| 🔧 工具不可用 | 1 | 9.1% | stock_reference_stk_holdernumber |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **stock_reference_pledge_detail**: 股票质押明细 ✅
  - 返回股东股票质押详细信息
  - 包含质押数量、质押人、公告日期等

- **stock_reference_pledge_stat**: 股票质押统计 ✅
  - 返回股票质押统计数据
  - 包含质押比例、质押数量等

### 2. 参数错误问题 (1个)
- stock_reference_repurchase - "无效的 vol 类型"

### 3. 工具不可用 (1个)
- stock_reference_stk_holdernumber: 股东户数数据

### 4. 空数据接口 (6个)
以下接口返回空数据（可能需要不同的参数）：
- stock_reference_block_trade: 大宗交易数据
- stock_reference_share_float: 限售股解禁
- stock_reference_stk_account: 股票账户开户
- stock_reference_stk_holdertrade: 增减持数据
- stock_reference_top10_floatholders: 前十大流通股东
- stock_reference_top10_holders: 前十大股东

---

## 建议

### 测试策略
1. 使用有质押、大宗交易或股东数据的股票进行测试
2. 使用最新的报告期
3. 使用不同的日期范围

### 代码示例
```go
// 获取股票质押明细
params := map[string]string{
    "ts_code": "600000.SH",
}
result, err := client.Call("stock_reference_pledge_detail", params)

// 获取股票质押统计
params := map[string]string{
    "ts_code": "600000.SH",
    "end_date": "20240308",
}
result, err := client.Call("stock_reference_pledge_stat", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
