# tushare-stock-margin 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-margin
- **测试工具数**: 7 个
- **测试成功**: 3 个 (42.9%)
- **参数错误**: 2 个 (28.6%)
- **空数据**: 2 个 (28.6%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_margin_margin | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | stock_margin_margin_detail | ts_code=600000.SH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | stock_margin_margin_secs | ts_code=600000.SH, trade_date=20240308 | ✅ 正常 | {"trade_date": "20240308", "ts_code": "600000.SH", "name": "浦发银行"} |
| 4 | stock_margin_slb_len | start_date=20240301, end_date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 auc_amount 类型 |
| 5 | stock_margin_slb_len_mm | ts_code=600000.SH, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 6 | stock_margin_slb_sec | ts_code=600000.SH, start_date=20240301 | ⚠️ 参数错误 | API call failed: 无效的 lent_qnt 类型 |
| 7 | stock_margin_slb_sec_detail | ts_code=600000.SH, start_date=20240301 | ✅ 正常 | {"trade_date": "20240304", "ts_code": "600000.SH", "fee_rate": 2.6} |

---

## 统计摘要

- **总工具数**: 7 个
- **测试成功**: 3 个 (42.9%)
- **参数错误**: 2 个 (28.6%)
- **空数据**: 2 个 (28.6%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 28.6% | stock_margin_margin_secs, stock_margin_slb_sec_detail |
| ⚠️ 参数错误 | 2 | 28.6% | stock_margin_slb_len, stock_margin_slb_sec |
| 📭 空数据 | 2 | 28.6% | stock_margin_margin, stock_margin_margin_detail, stock_margin_slb_len_mm |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **stock_margin_margin_secs**: 融资融券标的列表 ✅
  - 返回可融资融券的股票列表
  - 包含交易所和股票名称信息

- **stock_margin_slb_sec_detail**: 转融券明细 ✅
  - 返回转融券交易明细
  - 包含期限、费率、融出数量等信息

### 2. 参数错误问题 (2个)
以下接口存在参数类型验证问题：
- stock_margin_slb_len - "无效的 auc_amount 类型"
- stock_margin_slb_sec - "无效的 lent_qnt 类型"

### 3. 空数据接口 (3个)
以下接口返回空数据（可能需要不同的参数或日期）：
- stock_margin_margin: 融资融券汇总
- stock_margin_margin_detail: 融资融券明细
- stock_margin_slb_len_mm: 做市借券交易汇总

---

## 建议

### 高优先级修复
需要修复以下接口的参数类型转换问题：
1. 检查API参数类型定义
2. 修复参数转换逻辑
3. 确保参数类型匹配

### 测试策略
1. 使用有融资融券交易的股票进行测试
2. 使用不同的日期范围
3. 使用上海深圳市场的数据

### 代码示例
```go
// 获取融资融券标的列表
params := map[string]string{
    "ts_code": "600000.SH",
    "trade_date": "20240308",
}
result, err := client.Call("stock_margin_margin_secs", params)

// 获取转融券明细
params := map[string]string{
    "ts_code": "600000.SH",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("stock_margin_slb_sec_detail", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
