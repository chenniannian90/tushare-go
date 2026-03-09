# tushare-stock-feature 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-feature
- **测试工具数**: 14 个
- **测试成功**: 7 个 (50.0%)
- **无权限**: 4 个 (28.6%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_feature_broker_recommend | month=202401 | ✅ 正常 | {"month": "202401", "broker": "东兴证券", "ts_code": "000786.SZ", "name": "北新建材"} |
| 2 | stock_feature_ccass_hold | ts_code=605009.SH, trade_date=20240308 | ✅ 正常 | {"trade_date": "20240308", "ts_code": "605009.SH", "name": "豪悅護理", "shareholding": "3054820"} |
| 3 | stock_feature_ccass_hold_detail | ts_code=605009.SH, trade_date=20240308 | ✅ 正常 | {"trade_date": "20240308", "ts_code": "605009.SH", "col_participant_name": "渣打银行(香港)有限公司"} |
| 4 | stock_feature_cyq_chips | ts_code=600000.SH, trade_date=20240308 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240308", "price": 0.6, "percent": 0.02} |
| 5 | stock_feature_cyq_perf | ts_code=600000.SH, trade_date=20240308 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240308", "his_low": 0.6, "his_high": 10.4} |
| 6 | stock_feature_hk_hold | ts_code=600000.SH, trade_date=20240308 | ✅ 正常 | {"trade_date": "20240308", "ts_code": "600000.SH", "name": "浦發銀行", "vol": 658636726} |
| 7 | stock_feature_report_rc | ts_code=600000.SH, start_date=20240301, end_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 8 | stock_feature_stk_ah_comparison | ts_code=600000.SH, trade_date=20240308 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 9 | stock_feature_stk_auction_c | ts_code=600000.SH, trade_date=20240308 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 10 | stock_feature_stk_auction_o | ts_code=600000.SH, trade_date=20240308 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 11 | stock_feature_stk_factor_pro | ts_code=600000.SH, trade_date=20240308 | ✅ 数据量大 | {"ts_code": "600000.SH", "trade_date": "20240308", "open": 7.12, "close": 7.12} |
| 12 | stock_feature_stk_nineturn | ts_code=600000.SH, trade_date=20240308, freq=daily | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 13 | stock_feature_stk_surv | ts_code=600000.SH, start_date=20240301, end_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 14 | stock_feature_stk_rewards | ts_code=600000.SH, end_date=20231231 | 🔧 工具不可用 | Error: No such tool available |

---

## 统计摘要

- **总工具数**: 14 个
- **测试成功**: 7 个 (50.0%)
- **无权限**: 4 个 (28.6%)
- **数据量大**: 1 个 (7.1%)
- **空数据**: 2 个 (14.3%)
- **工具不可用**: 1 个 (7.1%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 6 | 42.9% | stock_feature_broker_recommend, stock_feature_ccass_hold, stock_feature_ccass_hold_detail, stock_feature_cyq_chips, stock_feature_cyq_perf, stock_feature_hk_hold |
| ✅ 数据量大 | 1 | 7.1% | stock_feature_stk_factor_pro |
| ❌ 无权限 | 4 | 28.6% | stock_feature_stk_ah_comparison, stock_feature_stk_auction_c, stock_feature_stk_auction_o, stock_feature_stk_nineturn |
| 📭 空数据 | 2 | 14.3% | stock_feature_report_rc, stock_feature_stk_surv |

---

## 主要发现

### 1. 可直接使用的接口 (6个)
- **stock_feature_broker_recommend**: 券商月度金股 ✅
- **stock_feature_ccass_hold**: 中央结算系统持股汇总 ✅
- **stock_feature_ccass_hold_detail**: 中央结算系统机构持股明细 ✅
- **stock_feature_cyq_chips**: A股筹码分布 ✅
- **stock_feature_cyq_perf**: A股筹码平均成本 ✅
- **stock_feature_hk_hold**: 沪深港股通持股明细 ✅

### 2. 需要权限的接口 (4个)
以下接口需要升级 Tushare 账户权限：
- stock_feature_stk_ah_comparison
- stock_feature_stk_auction_c
- stock_feature_stk_auction_o
- stock_feature_stk_nineturn

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 3. 需要分页的接口 (1个)
以下接口数据量较大，建议使用 offset/limit 参数：
- stock_feature_stk_factor_pro (包含大量技术指标)

### 4. 空数据接口 (2个)
以下接口返回空数据（可能需要不同的日期或参数）：
- stock_feature_report_rc - 券商研报盈利预测
- stock_feature_stk_surv - 机构调研记录

---

## 建议

### 分页使用示例
```json
{
  "offset": "0",    // 起始位置
  "limit": "100"    // 每页数量
}
```

### 代码示例
```go
// 获取券商金股推荐
params := map[string]string{
    "month": "202401",
}
result, err := client.Call("stock_feature_broker_recommend", params)

// 获取筹码分布
params := map[string]string{
    "ts_code": "600000.SH",
    "trade_date": "20240308",
}
result, err := client.Call("stock_feature_cyq_chips", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
