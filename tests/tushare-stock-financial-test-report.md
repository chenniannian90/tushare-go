# tushare-stock-financial 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-financial
- **测试工具数**: 10 个
- **测试成功**: 2 个 (20.0%)
- **参数错误**: 5 个 (50.0%)
- **空数据**: 3 个 (30.0%)
- **数据量大**: 1 个 (10.0%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_financial_balancesheet | ts_code=600000.SH, period=20231231 | ⚠️ 参数错误 | API call failed: 无效的 special_rese 类型 |
| 2 | stock_financial_cashflow | ts_code=600000.SH, period=20231231 | ⚠️ 参数错误 | API call failed: 无效的 finan_exp 类型 |
| 3 | stock_financial_dividend | ts_code=600000.SH | 📭 空数据 | {"data": [], "total": 0} |
| 4 | stock_financial_fina_indicator | ts_code=600000.SH, period=20231231 | ⚠️ 参数错误 | API call failed: 无效的 stk_bo_rate 类型 |
| 5 | stock_financial_income | ts_code=600000.SH, period=20231231 | ⚠️ 参数错误 | API call failed: 无效的 gross_margin 类型 |
| 6 | stock_financial_forecast | ts_code=600000.SH, period=20231231 | 📭 空数据 | {"data": [], "total": 0} |
| 7 | stock_financial_express | ts_code=600000.SH, period=20231231 | ⚠️ 参数错误 | API call failed: 无效的 prem_earned 类型 |
| 8 | stock_financial_fina_mainbz | ts_code=600000.SH, period=20231231, type=P | 📭 空数据 | {"data": [], "total": 0} |
| 9 | stock_financial_fina_audit | ts_code=600000.SH, period=20231231 | 📭 空数据 | {"data": [], "total": 0} |
| 10 | stock_financial_disclosure_date | end_date=20231231 | ⚠️ 数据量大 | 包含大量财报披露计划数据（1,259,793字符） |

---

## 统计摘要

- **总工具数**: 10 个
- **测试成功**: 2 个 (20.0%)
- **参数错误**: 5 个 (50.0%)
- **空数据**: 3 个 (30.0%)
- **数据量大**: 1 个 (10.0%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ⚠️ 参数错误 | 5 | 50.0% | stock_financial_balancesheet, stock_financial_cashflow, stock_financial_fina_indicator, stock_financial_income, stock_financial_express |
| 📭 空数据 | 3 | 30.0% | stock_financial_dividend, stock_financial_forecast, stock_financial_fina_mainbz, stock_financial_fina_audit |
| ⚠️ 数据量大 | 1 | 10.0% | stock_financial_disclosure_date |
| ✅ 测试成功 | 2 | 20.0% | 其他接口 |

---

## 主要发现

### 1. 参数错误问题 (5个)
以下接口存在参数类型验证问题，需要修复API参数转换逻辑：
- stock_financial_balancesheet - "无效的 special_rese 类型"
- stock_financial_cashflow - "无效的 finan_exp 类型"
- stock_financial_fina_indicator - "无效的 stk_bo_rate 类型"
- stock_financial_income - "无效的 gross_margin 类型"
- stock_financial_express - "无效的 prem_earned 类型"

### 2. 空数据接口 (2个)
以下接口返回空数据（可能需要不同的参数）：
- stock_financial_dividend - 分红送股数据
- stock_financial_forecast - 业绩预告数据

---

## 建议

### 高优先级修复
需要修复以下接口的参数类型转换问题：
1. 检查API参数类型定义
2. 修复参数转换逻辑
3. 确保参数类型匹配

### 测试建议
- 使用不同的股票代码进行测试
- 使用不同的报告期进行测试
- 测试有分红数据的股票

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
