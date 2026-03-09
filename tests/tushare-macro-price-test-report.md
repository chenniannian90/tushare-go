# tushare-macro-price 服务测试报告

## 测试概览

- **测���日期**: 2026-03-09
- **服务名称**: tushare-macro-price
- **测试工具数**: 2 个
- **测试成功**: 2 个 (100%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | macro_price_cn_cpi | m=202401 | ✅ 正常 | {"month": "202401", "nt_val": 99.2, "nt_yoy": -0.8} |
| 2 | macro_price_cn_ppi | m=202401 | ✅ 正常 | {"month": "202401", "ppi_yoy": -2.5, "ppi_mom": -0.2} |

---

## 统计摘要

- **总工具数**: 2 个
- **测试成功**: 2 个 (100%)

---

## 主要发现

### 1. 正常可用接口 (2个)
- **macro_price_cn_cpi**: CPI居民消费价格指数 ✅
  - 返回全国、城市和农村的CPI数据
  - 包含当月值、同比、环比等指标

- **macro_price_cn_ppi**: PPI工业生产者出厂价格指数 ✅
  - 返回PPI同比、环比数据
  - 包含生产资料、生活资料分类数据

### 代码示例
```go
// 获取CPI数据
params := map[string]string{
    "m": "202401",
}
result, err := client.Call("macro_price_cn_cpi", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
