# tushare-futures 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-futures
- **测试工具数**: 12 个
- **测试成功**: 2 个 (16.7%)
- **参数错误**: 2 个 (16.7%)
- **无权限**: 1 个 (8.3%)
- **空数据**: 4 个 (33.3%)
- **数据量大**: 2 个 (16.7%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | futures_fut_basic | exchange=CFFEX | ⚠️ 参数错误 | API call failed: 无效的 multiplier 类型 |
| 2 | futures_fut_daily | ts_code=CU2310.SHF, start_date=20231001 | ⚠️ 参数错误 | API call failed: 无效的 delv_settle 类型 |
| 3 | futures_ft_mins | ts_code=CU2310.SHF, freq=1min | ✅ 正常 | {"ts_code": "CU2310.SHF", "trade_time": "2023-10-09 09:05:00", "close": 67480} |
| 4 | futures_fut_holding | trade_date=20240308, ts_code=CU2310.SHF | ⚠️ 数据量大 | 包含大量成交持仓数据（1,161,145字符） |
| 5 | futures_fut_mapping | ts_code=CU0.SHF, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 6 | futures_fut_settle | trade_date=20240308, ts_code=CU2310.SHF | 📭 空数据 | {"data": [], "total": 0} |
| 7 | futures_fut_weekly_detail | exchange=SHFE, week=202401, prd=CU | 📭 空数据 | {"data": [], "total": 0} |
| 8 | futures_fut_weekly_monthly | ts_code=CU0.SHF, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 9 | futures_fut_wsr | exchange=SHFE, symbol=CU | ⚠️ 数据量大 | 包含大量仓单数据（91,885字符） |
| 10 | futures_index_daily | ts_code=NH.NH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 11 | futures_rt_fut_min | ts_code=CU2310.SHF, freq=1MIN | ❌ 无权限 | ACCESS_DENIED |
| 12 | futures_trade_cal | exchange=SHFE, start_date=20240301 | ✅ 正常 | {"exchange": "SHFE", "cal_date": "20240308", "is_open": "1"} |

---

## 统计摘要

- **总工具数**: 12 个
- **测试成功**: 2 个 (16.7%)
- **参数错误**: 2 个 (16.7%)
- **无权限**: 1 个 (8.3%)
- **空数据**: 4 个 (33.3%)
- **数据量大**: 2 个 (16.7%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 16.7% | futures_ft_mins, futures_trade_cal |
| ⚠️ 参数错误 | 2 | 16.7% | futures_fut_basic, futures_fut_daily |
| ⚠️ 数据量大 | 2 | 16.7% | futures_fut_holding, futures_fut_wsr |
| 📭 空数据 | 4 | 33.3% | futures_fut_mapping, futures_fut_settle等 |
| ❌ 无权限 | 1 | 8.3% | futures_rt_fut_min |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **futures_ft_mins**: 期货分钟数据 ✅
  - 返回期货合约分钟级别行情
  - 包含开高低收、成交量、持仓量等

- **futures_trade_cal**: 交易日历 ✅
  - 返回交易所交易日历
  - 包含是否交易、前一个交易日等信息

### 2. 数据量大接口 (2个)
- **futures_fut_holding**: 成交持仓排名
  - 返回数据量极大（1.1M字符）
  - 建议使用具体合约代码筛选

- **futures_fut_wsr**: 仓单日报
  - 返回大量仓单数据（91K字符）
  - 建议使用日期范围限制

### 3. 参数错误问题 (2个)
- futures_fut_basic - "无效的 multiplier 类型"
- futures_fut_daily - "无效的 delv_settle 类型"

### 4. 需要权限的接口 (1个)
- futures_rt_fut_min: 实时分钟数据

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 5. 空数据接口 (4个)
- futures_fut_mapping: 主力合约映射
- futures_fut_settle: 结算参数
- futures_fut_weekly_detail: 每周交易统计
- futures_fut_weekly_monthly: 周/月线行情
- futures_index_daily: 南华指数行情

---

## 建议

### 分页使用示例
```json
{
  "offset": "0",    // 起始位置
  "limit": "100"    // 每页数量
}
```

### 代码示例
```go
// 获取期货分钟数据
params := map[string]string{
    "ts_code": "CU2310.SHF",
    "freq": "1min",
    "start_date": "2023-10-09 09:00:00",
    "end_date": "2023-10-09 09:05:00",
}
result, err := client.Call("futures_ft_mins", params)

// 获取交易日历
params := map[string]string{
    "exchange": "SHFE",
    "start_date": "20240301",
    "end_date": "20240308",
}
result, err := client.Call("futures_trade_cal", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
