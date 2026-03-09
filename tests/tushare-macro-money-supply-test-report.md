# tushare-macro-money-supply 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-macro-money-supply
- **测试工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_money_supply_cn_m | m=202401 | ✅ 正常 | {"month": "202401", "m0": 121398.54, "m1": 694197.88, "m2": 2976250.2} |

---

## 统计摘要

- **总工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 主要发现

### 1. 正常可用接口 (1个)
- **macro_money_supply_cn_m**: 货币供应量数据 ✅
  - 返回M0、M1、M2数据
  - 月度更新

---

## 建议

### 代码示例
```go
// 获取货币供应量数据
params := map[string]string{
    "m": "202401",
}
result, err := client.Call("macro_money_supply_cn_m", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
