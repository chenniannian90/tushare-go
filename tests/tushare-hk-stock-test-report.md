# tushare-hk-stock 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-hk-stock
- **测试工具数**: 10 个
- **测试成功**: 2 个 (20.0%)
- **无权限**: 4 个 (40.0%)
- **参数错误**: 1 个 (10.0%)
- **空数据**: 2 个 (20.0%)
- **数据量大**: 1 个 (10.0%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | hk_stock_hk_basic | list_status=L | ⚠️ 数据量大 | 包含大量港股基础信息（1,149,399字符） |
| 2 | hk_stock_hk_daily | ts_code=00700.HK, start_date=20240301 | ✅ 正常 | {"ts_code": "00700.HK", "trade_date": "20240305", "open": 270.2, "close": 268.2} |
| 3 | hk_stock_hk_daily_adj | ts_code=00700.HK, start_date=20240301 | ⚠️ 参数错误 | API call failed: 无效的 vol 类型 |
| 4 | hk_stock_hk_mins | ts_code=00700.HK, freq=1min | 📭 空数据 | {"data": [], "total": 0} |
| 5 | hk_stock_hk_tradecal | start_date=20240301, end_date=20240305 | ✅ 正常 | {"cal_date": "20240305", "is_open": "1", "pretrade_date": "20240304"} |
| 6 | hk_stock_hk_balancesheet | ts_code=00700.HK, period=20231231 | ❌ 无权限 | ACCESS_DENIED |
| 7 | hk_stock_hk_cashflow | ts_code=00700.HK, period=20231231 | ❌ 无权限 | ACCESS_DENIED |
| 8 | hk_stock_hk_fina_indicator | ts_code=00700.HK, period=20231231 | ❌ 无权限 | ACCESS_DENIED |
| 9 | hk_stock_hk_income | ts_code=00700.HK, period=20231231 | ❌ 无权限 | ACCESS_DENIED |

---

## 统计摘要

- **总工具数**: 10 个
- **测试成功**: 2 个 (20.0%)
- **无权限**: 4 个 (40.0%)
- **参数错误**: 1 个 (10.0%)
- **空数据**: 1 个 (10.0%)
- **数据量大**: 1 个 (10.0%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 20.0% | hk_stock_hk_daily, hk_stock_hk_tradecal |
| ⚠️ 数据量大 | 1 | 10.0% | hk_stock_hk_basic |
| ❌ 无权限 | 4 | 40.0% | hk_stock_hk_balancesheet, hk_stock_hk_cashflow等 |
| ⚠️ 参数错误 | 1 | 10.0% | hk_stock_hk_daily_adj |
| 📭 空数据 | 1 | 10.0% | hk_stock_hk_mins |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **hk_stock_hk_daily**: 港股日线行情 ✅
  - 返回开高低收、成交量等数据
  - 支持历史数据查询

- **hk_stock_hk_tradecal**: 港股交易日历 ✅
  - 返回交易日历信息
  - 包含是否交易、前一个交易日等

### 2. 数据量大接口 (1个)
- **hk_stock_hk_basic**: 港股基础信息
  - 返回数据量极大（1.1M字符）
  - 建议使用 list_status, offset, limit 参数分页

### 3. 需要权限的接口 (4个)
以下接口需要升级 Tushare 账户权限：
- hk_stock_hk_balancesheet: 资产负债表
- hk_stock_hk_cashflow: 现金流量表
- hk_stock_hk_fina_indicator: 财务指标
- hk_stock_hk_income: 利润表

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 4. 参数错误问题 (1个)
- hk_stock_hk_daily_adj - "无效的 vol 类型"

---

## 建议

### 分页使用示例
```json
{
  "list_status": "L",
  "offset": "0",
  "limit": "100"
}
```

### 代码示例
```go
// 获取港股日线行情
params := map[string]string{
    "ts_code": "00700.HK",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("hk_stock_hk_daily", params)

// 获取港股交易日历
params := map[string]string{
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("hk_stock_hk_tradecal", params)

// 获取港股基础信息（分页）
params := map[string]string{
    "list_status": "L",
    "offset": "0",
    "limit": "100",
}
result, err := client.Call("hk_stock_hk_basic", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
