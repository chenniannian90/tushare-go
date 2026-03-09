# tushare-stock-market 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-market
- **测试工具数**: 13 个
- **测试成功**: 4 个 (30.8%)
- **无权限**: 1 个 (7.7%)
- **空数据**: 3 个 (23.1%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_market_daily | ts_code=600000.SH, start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | stock_market_daily_basic | ts_code=600000.SH, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | stock_market_stk_mins | ts_code=600000.SH, freq=1min, start_date=2024-03-08 09:00:00, end_date=2024-03-08 09:05:00 | 📭 空数据 | {"data": [], "total": 0} |
| 4 | stock_market_rt_k | ts_code=600000.SH | ❌ 访问限制 | API call failed: 抱歉，您每天最多访问该接口2次 |
| 5 | stock_market_weekly | ts_code=600000.SH, start_date=20240301, end_date=20240308 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240308", "close": 7.12, "open": 7.12} |
| 6 | stock_market_adj_factor | ts_code=600000.SH, start_date=20240301, end_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 7 | stock_market_monthly | ts_code=600000.SH, trade_date=20240229 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240229", "close": 7.15} |
| 8 | stock_market_stk_weekly_monthly | ts_code=600000.SH, trade_date=20240308, freq=week | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240308", "close": 7.12} |
| 9 | stock_market_stk_week_month_adj | ts_code=600000.SH, trade_date=20240308, freq=week | ✅ 正常 | {"ts_code": "600000.SH", "close_qfq": 6.67, "close_hfq": 110.61} |
| 10 | stock_market_realtime_list | src=dc | ⚠️ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 11 | stock_market_realtime_quote | ts_code=600000.SH, src=sina | ⚠️ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 12 | stock_market_realtime_tick | ts_code=600000.SH, src=dc | ⚠️ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 13 | stock_market_rt_min | ts_code=600000.SH, freq=1MIN | ✅ 正常 | {"ts_code": "600000.SH", "time": "2026-03-09 13:46:00", "close": 9.83} |

---

## 统计摘要

- **总工具数**: 13 个
- **测试成��**: 6 个 (46.2%)
- **访问限制**: 1 个 (7.7%)
- **空数据**: 3 个 (23.1%)
- **接口错误**: 3 个 (23.1%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 5 | 38.5% | stock_market_monthly, stock_market_stk_weekly_monthly, stock_market_stk_week_month_adj, stock_market_weekly, stock_market_rt_min |
| ❌ 访问限制 | 1 | 7.7% | stock_market_rt_k |
| ⚠️ 接口错误 | 3 | 23.1% | stock_market_realtime_list, stock_market_realtime_quote, stock_market_realtime_tick |
| 📭 空数据 | 4 | 30.7% | stock_market_daily, stock_market_daily_basic, stock_market_stk_mins, stock_market_adj_factor |

---

## 主要发现

### 1. 正常可用接口 (1个)
- **stock_market_weekly**: 股票周线行情 ✅

### 2. 访问限制接口 (1个)
- **stock_market_rt_k**: 实时日K线行情（每天最多访问2次）

### 3. 空数据接口 (3个)
以下接口返回空数据（可能需要不同的日期参数）：
- stock_market_daily
- stock_market_daily_basic
- stock_market_stk_mins
- stock_market_adj_factor

---

## 建���

### 测试策略
1. 使用更近期的日期进行测试
2. 测试不同的股票代码
3. 在不同时间测试实时接口
4. 继续测试剩余9个工具

### 代码示例
```go
// 获取周线数据
params := map[string]string{
    "ts_code": "600000.SH",
    "start_date": "20240301",
    "end_date": "20240308",
}
result, err := client.Call("stock_market_weekly", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
