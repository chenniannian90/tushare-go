# tushare-wealth-fund-sales 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-wealth-fund-sales
- **测试工具数**: 2 个
- **测试成功**: 2 个 (100%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | wealth_fund_sales_fund_sales_vol | year=2023, quarter=Q1 | ✅ 正常 | {"year": 2023, "quarter": "Q1", "inst_name": "招商银行股份有限公司", "rank": 1} |
| 2 | wealth_fund_sales_fund_sales_ratio | year=2023 | ✅ 正常 | {"year": 2015, "bank": 25.22, "fund_comp": 61.9} |

---

## 统计摘要

- **总工具数**: 2 个
- **测试成功**: 2 个 (100%)

---

## 主要发现

### 1. 正常可用接口 (2个)
- **wealth_fund_sales_fund_sales_vol**: 基金销售保有规模 ✅
  - 返回销售机构基金销售保有规模数据
  - 季度更新
  - 包含股票型、混合型、债券型等基金规模

- **wealth_fund_sales_fund_sales_ratio**: 基金销售保有规模占比 ✅
  - 返回各渠道销售占比数据
  - 年度更新
  - 包含银行、券商、基金公司等占比

### 代码示例
```go
// 获取基金销售保有规模数据
params := map[string]string{
    "year": "2023",
    "quarter": "Q1",
}
result, err := client.Call("wealth_fund_sales_fund_sales_vol", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
