# tushare-stock-board 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-stock-board
- **测试工具数**: 6 个
- **测试成功**: 2 个 (33.3%)
- **无权限**: 2 个 (33.3%)
- **空数据**: 2 个 (33.3%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | stock_board_dc_daily | ts_code=885812.DC, trade_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | stock_board_dc_hot | trade_date=20240305, market=A股市场 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | stock_board_dc_index | name=人形机器人, trade_date=20240305 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 4 | stock_board_limit_list_d | trade_date=20240305, limit_type=U | ❌ 访问限制 | ACCESS_DENIED: 您每小时最多访问该接口1次 |
| 5 | stock_board_top_list | trade_date=20240305 | ✅ 正常 | {"ts_code": "000005.SZ", "name": "ST星源", "close": 0.83} |
| 6 | stock_board_ths_hot | trade_date=20240305, market=热股 | ✅ 正常 | {"ts_code": "000628.SZ", "ts_name": "高新发展", "rank": 1} |

---

## 统计摘要

- **总工具数**: 6 个
- **测试成功**: 2 个 (33.3%)
- **无权限**: 1 个 (16.7%)
- **访问限制**: 1 个 (16.7%)
- **空数据**: 2 个 (33.3%)

---

## 主要发现

### 1. 正常可用接口 (2个)
- **stock_board_top_list**: 龙虎榜每日交易明细 ✅
  - 返回详细的龙虎榜交易数据
  - 包含营业部买卖金额

- **stock_board_ths_hot**: 同花顺热榜数据 ✅
  - 返回当日热股排行
  - 包含概念、热度值等信息

### 2. 访问限制
- **stock_board_limit_list_d**: 每小时最多访问1次
- **stock_board_dc_index**: 需要升级权限

### 3. 空数据接口 (2个)
- stock_board_dc_daily: 板块日线行情
- stock_board_dc_hot: 东财热榜

---

## 建议

### 测试策略
1. 使用不同的板块代码参数测试
2. 使用最新的交易日期
3. 注意接口访问频率限制

### 代码示例
```go
// 获取龙虎榜数据
params := map[string]string{
    "trade_date": "20240305",
}
result, err := client.Call("stock_board_top_list", params)

// 获取同花顺热榜
params := map[string]string{
    "trade_date": "20240305",
    "market": "热股",
}
result, err := client.Call("stock_board_ths_hot", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
