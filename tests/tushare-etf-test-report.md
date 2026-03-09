# tushare-etf 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-etf
- **测试工具数**: 7 个
- **测试成功**: 2 个 (28.6%)
- **无权限**: 4 个 (57.1%)
- **访问限制**: 1 个 (14.3%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状�� | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | etf_fund_daily | ts_code=159001.SZ, start_date=20240301, end_date=20240305 | ✅ 正常 | {"ts_code": "159001.SZ", "trade_date": "20240305", "close": 100.001} |
| 2 | etf_rt_etf_k | ts_code=15*.SZ | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权��� |
| 3 | etf_etf_share_size | ts_code=159001.SZ, start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |

---

## 统计摘要

- **总工具数**: 7 个
- **测试成功**: 2 个 (28.6%)
- **无权限**: 4 个 (57.1%)
- **访问限制**: 1 个 (14.3%)

---

## 主要发现

### 1. 正常可用接口 (2个)
- **etf_fund_daily**: ETF基金日线行情 ✅
  - 返回开高低收、成交��等数据
  - 支持历史数据查询

- **etf_fund_adj**: ETF复权因子 ✅
  - 返回复权因子数据
  - 用于计算复权行情

### 2. 需要权限的接口 (4个)
- etf_rt_etf_k: 实时行情接口
- etf_etf_share_size: ETF份额和规模数据
- etf_etf_basic: ETF基础信息
- etf_etf_index: ETF指数信息

### 3. 访问限制接口 (1个)
- etf_stk_mins: 分钟数据（每天最多访问2次）

### 代码示例
```go
// 获取ETF日线行情
params := map[string]string{
    "ts_code": "159001.SZ",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("etf_fund_daily", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
