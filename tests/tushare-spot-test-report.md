# tushare-spot 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-spot
- **测试工具数**: 2 个
- **参数错误**: 1 个 (50%)
- **工具不可用**: 1 个 (50%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | spot_sge_daily | ts_code=AU99.99, trade_date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 settle_vol 类型 |
| 2 | spot_sge_basic | - | ❌ 工具不可用 | Error: No such tool available |

---

## 统计摘要

- **总工具数**: 2 个
- **参数错误**: 1 个 (50%)
- **工具不可用**: 1 个 (50%)

---

## 主要发现

### 1. 参数错误问题
- **spot_sge_daily**: 存在参数类型转换错误（settle_vol）

### 2. 工具问题
- **spot_sge_basic**: 工具在当前环境中不可用

---

## 建议

### 测试策略
1. 修复参数转换问题后重新测试
2. 使用不同的贵金属代码

### 代码示例
```go
// 获取上海黄金交易所现货日线行情
params := map[string]string{
    "ts_code": "AU99.99",
    "trade_date": "20240305",
}
result, err := client.Call("spot_sge_daily", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
