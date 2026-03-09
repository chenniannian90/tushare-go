# tushare-macro-interest-rate 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-macro-interest-rate
- **测试工具数**: 7 个
- **测试成功**: 1 个 (14.3%)
- **工具不可用**: 1 个 (14.3%)
- **空数据**: 5 个 (71.4%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_interest_rate_shibor | date=20240305 | ✅ 正常 | {"date": "20240305", "on": 1.719, "1w": 1.855} |
| 2 | macro_interest_rate_shibor_quote | date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | macro_interest_rate_lpr | date=20240305 | ❌ 工具不可用 | Error: No such tool available |
| 4 | macro_interest_rate_hibor | date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 5 | macro_interest_rate_libor | date=20240305, curr_type=USD | 📭 空数据 | {"data": [], "total": 0} |
| 6 | macro_interest_rate_wz_index | date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 7 | macro_interest_rate_gz_index | date=20240305 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 7 个
- **测试成功**: 1 个 (14.3%)
- **工具不可用**: 1 个 (14.3%)
- **空数据**: 5 个 (71.4%)

---

## 主要发现

### 1. 正常可用接口 (1个)
- **macro_interest_rate_shibor**: Shibor利率 ✅
  - 返回各期限Shibor利率数据
  - 包含隔夜到1年期利率

### 2. 工具问题
- **macro_interest_rate_lpr**: 工具在当前环境中不可用

### 3. 空数据接口
- macro_interest_rate_wz_index: 温州民间借贷利率

---

## 建议

### 测试策略
1. 继续测试剩余4个工具

### 代码示例
```go
// 获取Shibor利率数据
params := map[string]string{
    "date": "20240305",
}
result, err := client.Call("macro_interest_rate_shibor", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
