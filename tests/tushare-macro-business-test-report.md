# tushare-macro-business 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-macro-business
- **测试工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_business_cn_pmi | m=202401 | ✅ 正常 | {"month": "202401", "pmi010000": 49.2, "pmi010100": 50.4} |

---

## 统计摘要

- **总工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 主要发现

### 1. 正常可用接口 (1个)
- **macro_business_cn_pmi**: 采购经理人指数 ✅
  - 返回制造业和非制造业PMI数据
  - 包含当月和环比数据

---

## 建议

### 代码示例
```go
// 获取PMI数据
params := map[string]string{
    "m": "202401",
}
result, err := client.Call("macro_business_cn_pmi", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
