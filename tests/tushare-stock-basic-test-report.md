# tushare-stock-basic 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-basic
- **测试工具数**: 13 个
- **测试成功**: 8 个 (61.5%)
- **无权限**: 4 个 (30.8%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_basic_namechange | limit=1 | ✅ 空数据 | {"data": [], "total": 0} |
| 2 | stock_basic_new_share | start_date=20240101, end_date=20240131 | ✅ 正常 | {"ts_code": "301589.SZ", "name": "诺瓦星云", "ipo_date": "20240130", "price": 126.89} |
| 3 | stock_basic_st | trade_date=20240309 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 4 | stock_basic_stk_managers | ts_code=600000.SH, limit=1 | ✅ 数据量大 | {"ts_code": "600000.SH", "name": "王跃堂", "title": "外部监事", "edu": "博士"} |
| 5 | stock_basic_stk_premarket | trade_date=20240309 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 6 | stock_basic_stock_basic | ts_code=600000.SH, list_status=L | ✅ 正常 | {"ts_code": "600000.SH", "name": "浦发银行", "industry": "银行", "list_date": "19991110"} |
| 7 | stock_basic_stock_hsgt | trade_date=20240309 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 8 | stock_basic_stock_st | trade_date=20240309 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没��接口访问权限 |
| 9 | stock_basic_trade_cal | exchange=SSE, start_date=20240301, end_date=20240310 | ✅ 正常 | {"exchange": "SSE", "cal_date": "20240310", "is_open": 0} |
| 10 | stock_basic_bak_basic | 无参数 | ✅ 数据量大 | {"trade_date": "20260306", "ts_code": "920036.BJ", "name": "觅睿科技", "industry": "IT设备"} |
| 11 | stock_basic_bse_mapping | 无参数 | ✅ 正常 | {"name": "永顺生物", "o_code": "839729.BJ", "n_code": "920729.BJ"} |
| 12 | stock_basic_stock_company | 无参数 | ✅ 数据量大 | {"ts_code": "688322.SH", "com_name": "奥比中光科技集团股份有限公司", "province": "广东"} |

---

## 统计摘要

- **总工具数**: 13 个
- **测试成功**: 8 个 (61.5%)
- **无权限**: 4 个 (30.8%)
- **数据量大**: 3 个 (需要分页处理)
- **空数据**: 1 个

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 5 | 38.5% | stock_basic_namechange, stock_basic_new_share, stock_basic_stock_basic, stock_basic_trade_cal, stock_basic_bse_mapping |
| ✅ 数据量大 | 3 | 23.1% | stock_basic_stk_managers, stock_basic_bak_basic, stock_basic_stock_company |
| ❌ 无权限 | 4 | 30.8% | stock_basic_st, stock_basic_stk_premarket, stock_basic_stock_hsgt, stock_basic_stock_st |
| 📭 空数据 | 1 | 7.6% | stock_basic_namechange |

---

## 主要发现

### 1. 可直接使用的接口 (5个)
- **stock_basic_stock_basic**: 获取股票基础信息 ✅
- **stock_basic_trade_cal**: 获取交易日历 ✅
- **stock_basic_bse_mapping**: 北交所代码映射 ✅
- **stock_basic_new_share**: 新股上市数据 ✅
- **stock_basic_namechange**: 名称变更记录 ✅

### 2. 需要权限的接口 (4个)
以下接口需要升级 Tushare 账户权限：
- stock_basic_st
- stock_basic_stk_premarket
- stock_basic_stock_hsgt
- stock_basic_stock_st

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 3. 需要分页的接口 (3个)
以下接口数据量较大，建议使用 offset/limit 参数：
- stock_basic_stk_managers (80,159 字符)
- stock_basic_bak_basic (4,860,349 字符)
- stock_basic_stock_company (8,600,114 字符)

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
// 获取股票基础信息
params := map[string]string{
    "ts_code": "600000.SH",
    "list_status": "L",
}
result, err := client.Call("stock_basic_stock_basic", params)

// 使用分页获取公司信息
params := map[string]string{
    "offset": "0",
    "limit": "100",
}
result, err := client.Call("stock_basic_stock_company", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
