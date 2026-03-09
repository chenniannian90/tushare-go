# tushare-forex 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-forex
- **测试工具数**: 2 个
- **空数据**: 1 个 (50%)
- **参数错误**: 1 个 (50%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | forex_fx_daily | ts_code=USD/CNY, start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | forex_fx_obasic | exchange=FXCM | ⚠️ 参数错误 | API call failed: 无效的 max_unit 类型 |

---

## 统计摘要

- **总工具数**: 2 个
- **空数据**: 1 个 (50%)
- **参数错误**: 1 个 (50%)

---

## 主要发现

### 1. 参数错误问题
- **forex_fx_obasic**: 存在参数类型转换错误（max_unit）

### 2. 空数据接口
- forex_fx_daily: 外汇日线行情
- 可能需要使用正确的货币对代码格式

---

## 建议

### 测试策略
1. 使用正确的货币对代码格式
2. 测试不同的交易商
3. 修复参数转换问题后重新测试

### 代码示例
```go
// 获取外汇日线行情
params := map[string]string{
    "ts_code": "USD/CNY",
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("forex_fx_daily", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
