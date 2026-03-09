# tushare-macro-social-financing 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-macro-social-financing
- **测试工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_social_financing_sf_month | m=202401 | ✅ 正常 | {"month": "202401", "inc_month": 64734} |

---

## 统计摘要

- **总工具数**: 1 个
- **测试成功**: 1 个 (100%)

---

## 主要发现

### 1. 正常可用接口 (1个)
- **macro_social_financing_sf_month**: 社会融资数据 ✅
  - 返回月度社会融资规模增量
  - 包含当月新增数据

---

## 建议

### 代码示例
```go
// 获取社会融资数据
params := map[string]string{
    "m": "202401",
}
result, err := client.Call("macro_social_financing_sf_month", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
