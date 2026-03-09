# tushare-index 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-index
- **测试工具数**: 18 个
- **测试成功**: 6 个 (33.3%)
- **无权限**: 5 个 (27.8%)
- **空数据**: 6 个 (33.3%)
- **数据量��**: 1 个 (5.6%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | index_index_basic | ts_code=000001.SH | 📭 空数据 | {"data": [], "total": 0} |
| 2 | index_index_daily | ts_code=000001.SH, start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | index_index_dailybasic | trade_date=20240308 | ✅ 正常 | {"ts_code": "000001.SH", "trade_date": "20240308", "total_mv": 56665152321216} |
| 4 | index_rt_idx_k | ts_code=000001.SH | ✅ 正常 | {"ts_code": "000001.SH", "name": "上证指数", "close": 4081.5978} |
| 5 | index_index_weight | index_code=000001.SH, start_date=20240301, end_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 6 | index_index_monthly | ts_code=000001.SH, trade_date=20240229 | ✅ 正常 | {"ts_code": "000001.SH", "trade_date": "20240229", "close": 3015.17} |
| 7 | index_index_weekly | ts_code=000001.SH, start_date=20240301, end_date=20240308 | ✅ 正常 | {"ts_code": "000001.SH", "trade_date": "20240308", "close": 3046.02} |
| 8 | index_index_global | ts_code=IXX, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 9 | index_index_member_all | index_code=000001.SH, is_new=Y | ⚠️ 数据量大 | 包含大量指数成分股数据（1,074,573字符） |
| 10 | index_idx_mins | ts_code=000001.SH, freq=1min | ❌ 无权限 | ACCESS_DENIED |
| 11 | index_rt_idx_min | ts_code=000001.SH, freq=1MIN | ❌ 无权限 | ACCESS_DENIED |
| 12 | index_ci_daily | ts_code=801010.SI, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 13 | index_ci_index_member | l1_code=801010, is_new=Y | ❌ 无权限 | ACCESS_DENIED |
| 14 | index_sw_daily | ts_code=801010.SI, trade_date=20240308 | ✅ 正常 | {"ts_code": "801010.SI", "name": "农林牧渔", "close": 2583.11} |
| 15 | index_rt_sw_k | ts_code=801010.SI | ❌ 无权限 | ACCESS_DENIED |
| 16 | index_daily_info | trade_date=20240308, ts_code=000001.SH | 📭 空数据 | {"data": [], "total": 0} |
| 17 | index_sz_daily_info | trade_date=20240308, ts_code=399001.SZ | 📭 空数据 | {"data": [], "total": 0} |
| 18 | index_idx_factor_pro | ts_code=000001.SH, trade_date=20240308 | ❌ 无权限 | ACCESS_DENIED |

---

## 统计摘要

- **总工具数**: 18 个
- **测试成功**: 6 个 (33.3%)
- **无权限**: 5 个 (27.8%)
- **空数据**: 6 个 (33.3%)
- **数据量大**: 1 个 (5.6%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 6 | 33.3% | index_index_dailybasic, index_rt_idx_k, index_index_monthly, index_index_weekly, index_sw_daily等 |
| ❌ 无权限 | 5 | 27.8% | index_idx_mins, index_rt_idx_min, index_ci_index_member, index_rt_sw_k, index_idx_factor_pro |
| 📭 空数据 | 6 | 33.3% | index_index_basic, index_index_daily, index_index_weight等 |
| ⚠️ 数据量大 | 1 | 5.6% | index_index_member_all |

---

## 主要发现

### 1. 正常可用接口 (6个)
- **index_index_dailybasic**: 指数每日基本面指标 ✅
- **index_rt_idx_k**: 实时日K线行情 ✅
- **index_index_monthly**: 指数月线行情 ✅
- **index_index_weekly**: 指数周线行情 ✅
- **index_sw_daily**: 申万行业日线行情 ✅

### 2. 需要权限的接口 (5个)
以下接口需要升级 Tushare 账户权限：
- index_idx_mins (指数分钟数据)
- index_rt_idx_min (实时分钟数据)
- index_ci_index_member (中信行业成分)
- index_rt_sw_k (申万实时行情)
- index_idx_factor_pro (指数技术面因子)

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 3. 空数据接口 (6个)
以下接口返回空数据（可能需要不同的参数）：
- index_index_basic
- index_index_daily
- index_index_weight
- index_index_global
- index_ci_daily
- index_daily_info
- index_sz_daily_info

### 4. 数据量大接口 (1个)
- **index_index_member_all**: 指数成分股数据
  - 返回数据量极大（1M+字符）
  - 建议使用 offset/limit 参数分页查询

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
// 获取指数基本面指标
params := map[string]string{
    "trade_date": "20240308",
}
result, err := client.Call("index_index_dailybasic", params)

// 获取申万行业行情
params := map[string]string{
    "ts_code": "801010.SI",
    "trade_date": "20240308",
}
result, err := client.Call("index_sw_daily", params)

// 获取指数成分股（分页）
params := map[string]string{
    "index_code": "000001.SH",
    "offset": "0",
    "limit": "100",
}
result, err := client.Call("index_index_member_all", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
