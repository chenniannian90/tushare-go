# tushare-fund 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-fund
- **测试工具数**: 8 个
- **测试成功**: 2 个 (25.0%)
- **参数错误**: 6 个 (75.0%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | fund_fund_basic | market=E | ⚠️ 参数错误 | API call failed: 无效的 duration_year 类型 |
| 2 | fund_fund_company | - | ⚠️ 参数错误 | API call failed: 无效的 employees 类型 |
| 3 | fund_fund_div | ts_code=000001.OF | ⚠️ 参数错误 | API call failed: 无效的 ear_amount 类型 |
| 4 | fund_fund_factor_pro | ts_code=000001.OF, trade_date=20240308 | ❌ 无权限 | ACCESS_DENIED |
| 5 | fund_fund_manager | ts_code=000001.OF | ✅ 正常 | {"ts_code": "000001.OF", "name": "刘睿聪", "gender": "M"} |
| 6 | fund_fund_nav | ts_code=000001.OF, start_date=20240301 | ⚠️ 参数错误 | API call failed: 无效的 accum_div 类型 |
| 7 | fund_fund_portfolio | ts_code=000001.OF, period=20231231 | ✅ 数据量大 | 包含大量基金持仓数据（185条记录） |
| 8 | fund_fund_share | ts_code=000001.OF, trade_date=20240308 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 8 个
- **测试成功**: 2 个 (25.0%)
- **参数错误**: 4 个 (50.0%)
- **无权限**: 1 个 (12.5%)
- **空数据**: 1 个 (12.5%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 25.0% | fund_fund_manager, fund_fund_portfolio |
| ⚠️ 参数错误 | 4 | 50.0% | fund_fund_basic, fund_fund_company, fund_fund_div, fund_fund_nav |
| ❌ 无权限 | 1 | 12.5% | fund_fund_factor_pro |
| 📭 空数据 | 1 | 12.5% | fund_fund_share |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **fund_fund_manager**: 基金经理数据 ✅
  - 返回基金经理详细信息
  - 包含简历、任职时间等

- **fund_fund_portfolio**: 基金持仓数据 ✅
  - 返回基金股票持仓明细
  - 包含市值、占比等信息
  - 数据量较大（185条记录）

### 2. 参数错误问题 (4个)
以下接口存在参数类型验证问题，需要修复API参数转换逻辑：
- fund_fund_basic - "无效的 duration_year 类型"
- fund_fund_company - "无效的 employees 类型"
- fund_fund_div - "无效的 ear_amount 类型"
- fund_fund_nav - "无效的 accum_div 类型"

### 3. 需要权限的接口 (1个)
- fund_fund_factor_pro: 基金技术面因子

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 4. 空数据接口 (1个)
- fund_fund_share: 基金规模数据

---

## 建议

### 高优先级修复
需要修复以下接口的参数类型转换问题：
1. 检查API参数类型定义
2. 修复参数转换逻辑
3. 确保参数类型匹配

### 测试建议
- 使用不同的基金代码进行测试
- 使用有分红数据的基金

### 代码示例
```go
// 获取基金经理数据
params := map[string]string{
    "ts_code": "000001.OF",
}
result, err := client.Call("fund_fund_manager", params)

// 获取基金持仓（使用limit限制数量）
params := map[string]string{
    "ts_code": "000001.OF",
    "period": "20231231",
    "limit": "50",
}
result, err := client.Call("fund_fund_portfolio", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
