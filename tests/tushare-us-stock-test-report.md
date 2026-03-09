# tushare-us-stock 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-us-stock
- **测试工具数**: 6 个
- **测试成功**: 3 个 (50.0%)
- **无权限**: 2 个 (33.3%)
- **工具不可用**: 1 个 (16.7%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | us_stock_us_daily | ts_code=AAPL, start_date=20240301, end_date=20240305 | ✅ 正常 | {"ts_code": "AAPL", "trade_date": "20240305", "close": 170.12, "open": 170.76} |
| 2 | us_stock_us_basic | - | ❌ 工具不可用 | Error: No such tool available |
| 3 | us_stock_us_daily_adj | ts_code=AAPL, start_date=20240301, end_date=20240305 | ✅ 正常 | {"ts_code": "AAPL", "close": 169.89, "total_mv": 2626971000000} |
| 4 | us_stock_us_balancesheet | ts_code=AAPL, period=20231231 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 5 | us_stock_us_cashflow | ts_code=AAPL, period=20231231 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 6 | us_stock_us_fina_indicator | ts_code=AAPL, period=20231231, report_type=Q4 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 7 | us_stock_us_income | ts_code=AAPL, period=20231231 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 8 | us_stock_us_tradecal | start_date=20240301, end_date=20240305, is_open=1 | ✅ 正常 | {"cal_date": "20240305", "is_open": "1"} |

---

## 统计摘要

- **总工具数**: 6 个
- **测试成功**: 3 个 (50.0%)
- **无权限**: 3 个 (50.0%)
- **工具不可用**: 1 个 (16.7%)

---

## 主要发现

### 1. 正常可用接口 (3个)
- **us_stock_us_daily**: 美股日线行情 ✅
  - 返回开高低收、成交量等数据
  - 支持历史数据查询

- **us_stock_us_daily_adj**: 美股复权日线行情 ✅
  - 返回复权后的行情数据
  - 包含市值、流通市值等详细信息

- **us_stock_us_tradecal**: 美股交易日历 ✅
  - 返回交易日历信息
  - 包含是否开市、前一交易日等信息

### 2. 工具问题
- **us_stock_us_basic**: 工具在当前环境中不可用

### 3. 需要权限的接口 (3个)
- us_stock_us_balancesheet: 资产负债表
- us_stock_us_cashflow: 现金流量表
- us_stock_us_fina_indicator: 财务指标
- us_stock_us_income: 利润表

### 代码示例
```go
// 获取美股日线行情
params := map[string]string{
    "ts_code": "AAPL",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("us_stock_us_daily", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
