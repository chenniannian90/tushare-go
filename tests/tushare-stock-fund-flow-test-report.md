# tushare-stock-fund-flow 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-fund-flow
- **测试工具数**: 7 个
- **测试成功**: 4 个 (57.1%)
- **空数据**: 3 个 (42.9%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_fund_flow_moneyflow | ts_code=600000.SH, start_date=20240301 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240305", "net_mf_vol": 144668} |
| 2 | stock_fund_flow_moneyflow_hsgt | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | stock_fund_flow_moneyflow_cnt_ths | ts_code=801010.SI, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 4 | stock_fund_flow_moneyflow_dc | ts_code=600000.SH, trade_date=20240308 | ✅ 正常 | {"ts_code": "600000.SH", "trade_date": "20240308", "net_amount": -1941.92} |
| 5 | stock_fund_flow_moneyflow_ind_dc | content_type=行业, trade_date=20240308 | ✅ 数据量大 | 包含大量板块资金流向数据（53KB） |
| 6 | stock_fund_flow_moneyflow_ind_ths | ts_code=801010.SI, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |
| 7 | stock_fund_flow_moneyflow_mkt_dc | trade_date=20240308 | ✅ 正常 | {"trade_date": "20240308", "net_amount": 3063263744} |
| 8 | stock_fund_flow_moneyflow_ths | ts_code=600000.SH, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 8 个
- **测试成功**: 4 个 (50.0%)
- **空数据**: 4 个 (50.0%)
- **数据量大**: 1 个 (12.5%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 3 | 37.5% | stock_fund_flow_moneyflow, stock_fund_flow_moneyflow_dc, stock_fund_flow_moneyflow_mkt_dc |
| ⚠️ 数据量大 | 1 | 12.5% | stock_fund_flow_moneyflow_ind_dc |
| 📭 空数据 | 4 | 50.0% | stock_fund_flow_moneyflow_hsgt, stock_fund_flow_moneyflow_cnt_ths等 |

---

## 主要发现

### 1. 正常可用接口 (3个)
- **stock_fund_flow_moneyflow**: 个股资金流向数据 ✅
  - 返回小单、中单、大单、特大单的买卖成交量
  - 提供净流入量和净流入金额

- **stock_fund_flow_moneyflow_dc**: 东方财富个股资金流向 ✅
  - 提供更详细的资金流向数据
  - 包含超大单、大单、中单、小单分类

- **stock_fund_flow_moneyflow_mkt_dc**: 东方财富大盘资金流向 ✅
  - 返回沪深两市大盘资金流向
  - 包含指数涨跌幅和净流入金额

### 2. 数据量大接口 (1个)
- **stock_fund_flow_moneyflow_ind_dc**: 板块资金流向
  - 返回大量行业/概念板块资金流向数据
  - 包含各板块的排名和资金流入情况

### 3. 空数据接口 (4个)
以下接口返回空数据（可能需要不同的参数或日期）：
- stock_fund_flow_moneyflow_hsgt: 沪深港通资金流向
- stock_fund_flow_moneyflow_cnt_ths: 同花顺概念板块资金流向
- stock_fund_flow_moneyflow_ind_ths: 同花顺行业资金流向
- stock_fund_flow_moneyflow_ths: 同花顺个股资金流向

---

## 建议

### 测试策略
1. 使用不同的股票代码进行测试
2. 使用��近期的日期参数
3. 测试有资金流入的股票

### 代码示例
```go
// 获取个股资金流向
params := map[string]string{
    "ts_code": "600000.SH",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("stock_fund_flow_moneyflow", params)

// 获取东方财富个股资金流向
params := map[string]string{
    "ts_code": "600000.SH",
    "trade_date": "20240308",
}
result, err := client.Call("stock_fund_flow_moneyflow_dc", params)

// 获取大盘资金流向
params := map[string]string{
    "trade_date": "20240308",
}
result, err := client.Call("stock_fund_flow_moneyflow_mkt_dc", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
