# tushare-options 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-options
- **测试工具数**: 3 个
- **数据量大**: 1 个 (33.3%)
- **空数据**: 2 个 (66.7%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | options_opt_basic | exchange=SSE | ⚠️ 数据量大 | 包含大量期权合约信息（7,232,220字符） |
| 2 | options_opt_daily | ts_code=10007976.SH, start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | options_opt_mins | ts_code=10007976.SH, freq=1min, start_date=2024-03-05 09:00:00 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 3 个
- **数据量大**: 1 个 (33.3%)
- **空数据**: 2 个 (66.7%)

---

## 主要发现

### 1. 数据量大接口
- **options_opt_basic**: 期权合约基础信息 ⚠️
  - 返回数据量极大（7.2M字符）
  - 需要使用分页参数或筛选条件

### 2. 空数据接口 (2个)
- options_opt_daily: 期权日线行情
- options_opt_mins: 期权分钟数据

---

## 建议

### 测试策略
1. 使用有数据的期权代码进行查询
2. 使用交易当日的时间范围测试分钟数据

### 代码示例
```go
// 获取期权合约信息（带筛选）
params := map[string]string{
    "exchange": "SSE",
    "ts_code": "10007976.SH",
}
result, err := client.Call("options_opt_basic", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
