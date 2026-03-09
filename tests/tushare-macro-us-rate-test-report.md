# tushare-macro-us-rate 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-macro-us-rate
- **测试工具数**: 4 个
- **测试成功**: 2 个 (50.0%)
- **参数错误**: 2 个 (50.0%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_us_rate_us_tbr | date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 w17_bd 类型 |
| 2 | macro_us_rate_us_tltr | date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 e_factor 类型 |
| 3 | macro_us_rate_us_trltr | date=20240305 | ✅ 正常 | {"date": "20240305", "ltr_avg": 2} |
| 4 | macro_us_rate_us_trycr | date=20240305 | ✅ 正常 | {"date": "20240305", "y5": 1.74, "y10": 1.81} |
| 5 | macro_us_rate_us_tycr | date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 m4 类型 |

---

## 统计摘要

- **总工具数**: 5 个
- **测试成功**: 2 个 (40.0%)
- **参数错误**: 3 个 (60.0%)

---

## 主要发现

### 1. 参数错误问题
- **macro_us_rate_us_tbr**: 存在参数类型转换错误

---

## 建议

### 测试策略
1. 修复参数转换问题后重新测试
2. 继续测试剩余3个工具

### 代码示例
```go
// 获取美国短期国债利率数据
params := map[string]string{
    "date": "20240305",
}
result, err := client.Call("macro_us_rate_us_tbr", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
