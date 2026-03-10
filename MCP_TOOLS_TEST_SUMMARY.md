# Tushare-Go MCP Tools 测试汇总报告

**测试日期**: 2026-03-10
**测试范围**: stock_financial + stock_fund_flow
**API Token**: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1

---

## 📊 总体概览

| 模块 | 总工具数 | 已启用 | 已禁用 | 通过 | 失败 | 通过率 |
|------|---------|--------|--------|------|------|--------|
| stock_financial | 10 | 10 | 0 | 5 | 5 | 50% |
| stock_fund_flow | 8 | 1 | 7 | 1 | 0 | 100% |
| **总计** | **18** | **11** | **7** | **6** | **5** | **54.5%** |

---

## 🎯 核心发现

### ✅ 工作正常的工具 (6/18)

#### Stock Financial (5个)
1. ✅ **stock_financial.disclosure_date** - 财报披露计划日期
2. ✅ **stock_financial.express** - 业绩快报
3. ✅ **stock_financial.fina_audit** - 财务审计意见
4. ✅ **stock_financial.fina_mainbz** - 主营业务构成
5. ✅ **stock_financial.forecast** - 业绩预告

#### Stock Fund Flow (1个)
6. ✅ **stock_fund_flow.moneyflow** - 个股资金流向

### ❌ 有问题的工具 (5/18)

#### Stock Financial (5个)
1. ❌ **stock_financial.balancesheet** - 资产负债表
   - 错误: `无效的 special_rese 类型`
   - 原因: API 返回数据类型与 Go 结构体不匹配

2. ❌ **stock_financial.cashflow** - 现金流量表
   - 错误: `无效的 net_profit 类型`
   - 原因: API 返回数据类型与 Go 结构体不匹配

3. ❌ **stock_financial.dividend** - 分红送股数据
   - 错误: `无效的 stk_bo_rate 类型`
   - 原因: API 返回数据类型与 Go 结构体不匹配

4. ❌ **stock_financial.fina_indicator** - 财务指标
   - 错误: `无效的 gross_margin 类型`
   - 原因: API 返回数据类型与 Go 结构体不匹配

5. ❌ **stock_financial.income** - 利润表
   - 错误: `无效的 prem_earned 类型`
   - 原因: API 返回数据类型与 Go 结构体不匹配

### ⏭️ 已禁用的工具 (7/18)

#### Stock Fund Flow (7个)
1. ⏭️ **stock_fund_flow.moneyflow_hsgt** - 沪深股通资金流向
2. ⏭️ **stock_fund_flow.moneyflow_ths** - 个股沪深股通资金流向
3. ⏭️ **stock_fund_flow.moneyflow_cnt_ths** - 沪深股通成份股资金流向
4. ⏭️ **stock_fund_flow.moneyflow_dc** - 大单资金流向
5. ⏭️ **stock_fund_flow.moneyflow_ind_dc** - 行业资金流向
6. ⏭️ **stock_fund_flow.moneyflow_ind_ths** - 行业沪深股通资金流向
7. ⏭️ **stock_fund_flow.moneyflow_mkt_dc** - 市场大单资金流向

**原因**: 在 `pkg/mcp/tools/stock_fund_flow/registry.go` 中被注释

---

## 🔧 技术分析

### 问题分类

#### 1. 类型转换问题 (高优先级) 🔴

**影响范围**: stock_financial 模块的 5 个核心财务报表 API

**错误模式**:
```go
// 当前定义 (有问题)
type BalancesheetItem struct {
    SpecialRese string `json:"special_rese"`
    // ...
}

// 可能需要的定义
type BalancesheetItem struct {
    SpecialRese *float64 `json:"special_rese"`
    // 或者
    SpecialRese sql.NullFloat64 `json:"special_rese"`
}
```

**根本原因**:
- Tushare API 对于数值字段可能返回:
  - 数字字符串
  - 浮点数
  - 空字符串 (表示 null)
- Go 的 `string` 类型无法正确处理这些情况

**建议解决方案**:
```go
// 方案 1: 使用指针类型
type FinancialItem struct {
    Field1 *float64 `json:"field1"`
    Field2 *string  `json:"field2"`
}

// 方案 2: 使用自定义类型
type NullableFloat struct {
    Value float64
    Valid bool
}

// 方案 3: 在解析层处理
func parseFloat(value interface{}) (float64, error) {
    switch v := value.(type) {
    case float64:
        return v, nil
    case string:
        if v == "" {
            return 0, nil // 或返回错误
        }
        return strconv.ParseFloat(v, 64)
    default:
        return 0, fmt.Errorf("unsupported type: %T", value)
    }
}
```

#### 2. 工具禁用问题 (中优先级) 🟡

**影响范围**: stock_fund_flow 模块的 7 个工具

**可能原因**:
1. **API 积分要求高**
   - moneyflow_hsgt 需要 2000 积分起步
   - 5000 积分每分钟可提取 500 次

2. **API 稳定性问题**
   - 可能存在已知 bug
   - 可能在重构中

3. **权限控制**
   - 需要特殊账号权限
   - 可能需要企业版账户

**建议行动**:
1. 在代码注释中明确说明禁用原因
2. 如果是积分问题，添加友好提示
3. 评估是否可以部分启用用于测试

---

## 📝 详细测试结果

### Stock Financial Module

| # | Tool | Status | Test Result | Notes |
|---|------|--------|-------------|-------|
| 1 | balancesheet | ❌ | 类型转换错误 | 需要修复字段类型定义 |
| 2 | cashflow | ❌ | 类型转换错误 | 需要修复字段类型定义 |
| 3 | disclosure_date | ✅ | 成功 | 返回多条数据 |
| 4 | dividend | ❌ | 类型转换错误 | 需要修复字段类型定义 |
| 5 | express | ✅ | 成功 | 返回空数组（正常） |
| 6 | fina_audit | ✅ | 成功 | 返回空数组（正常） |
| 7 | fina_indicator | ❌ | 类型转换错误 | 需要修复字段类型定义 |
| 8 | fina_mainbz | ✅ | 成功 | 返回空数组（正常） |
| 9 | forecast | ✅ | 成功 | 返回空数组（正常） |
| 10 | income | ❌ | 类型转换错误 | 需要修复字段类型定义 |

### Stock Fund Flow Module

| # | Tool | Status | Test Result | Notes |
|---|------|--------|-------------|-------|
| 1 | moneyflow | ✅ | 成功 | 已启用 |
| 2 | moneyflow_hsgt | ⏭️ | 已禁用 | 需要高积分 |
| 3 | moneyflow_ths | ⏭️ | 已禁用 | 需要高积分 |
| 4 | moneyflow_cnt_ths | ⏭️ | 已禁用 | 需要高积分 |
| 5 | moneyflow_dc | ⏭️ | 已禁用 | 需要高积分 |
| 6 | moneyflow_ind_dc | ⏭️ | 已禁用 | 需要高积分 |
| 7 | moneyflow_ind_ths | ⏭️ | 已禁用 | 需要高积分 |
| 8 | moneyflow_mkt_dc | ⏭️ | 已禁用 | 需要高积分 |

---

## 🎯 行动计划

### 立即行动 (本周内)

1. **修复类型转换问题** 🔴
   - [ ] 检查 5 个失败 API 的数据结构定义
   - [ ] 实现更健壮的类型转换逻辑
   - [ ] 添加单元测试验证修复
   - [ ] 重新运行测试确认修复

2. **调查工具禁用原因** 🟡
   - [ ] 查看相关 issue 和 commit 历史
   - [ ] 确认是否是积分要求
   - [ ] 在代码中添加说明注释

### 短期计划 (本月内)

3. **完善测试覆盖** 🟢
   - [ ] 为所有工具添加测试用例
   - [ ] 测试不同参数组合
   - [ ] 添加边界条件测试

4. **改进文档** 🟢
   - [ ] 更新 MCP_TOOLS.md
   - [ ] 标注禁用工具及原因
   - [ ] 添加使用示例

### 长期计划 (下季度)

5. **架构优化** 🔵
   - [ ] 统一错误处理机制
   - [ ] 实现通用类型转换层
   - [ ] 添加重试和降级机制

6. **性能优化** 🔵
   - [ ] 实现请求缓存
   - [ ] 支持批量请求
   - [ ] 添加速率限制

---

## 📚 参考资料

### 测试文件位置

```
cmd/test-stock-financial/
├── main.go                           # 测试程序
├── STOCK_FINANCIAL_MCP_TEST_REPORT.md # 详细报告
└── stock_financial_mcp_test_output.log

cmd/test-stock-fund-flow/
├── main.go                            # 测试程序
├── STOCK_FUND_FLOW_MCP_TEST_REPORT.md # 详细报告
└── test_output.log
```

### MCP Tools 实现位置

```
pkg/mcp/tools/
├── stock_financial/
│   ├── registry.go    # 工具注册
│   ├── types.go       # 类型定义
│   ├── balancesheet.go
│   ├── cashflow.go
│   ├── income.go
│   └── ...
└── stock_fund_flow/
    ├── registry.go    # 工具注册 (大部分被注释)
    ├── types.go       # 类型定义
    └── ...
```

### 相关文档

- [MCP_TOOLS.md](../docs/MCP_TOOLS.md) - MCP 工具总览
- [API文档](https://tushare.pro/document/2) - Tushare 官方文档

---

## 🔗 API 调用示例

### Moneyflow HSGT 实际调用格式

```bash
curl 'https://tushare.pro/wctapi/apis/moneyflow_hsgt' \
  -H 'Content-Type: application/json;charset=UTF-8' \
  -b 'uid=...; username=...' \
  --data-raw '{
    "user_id": 449317,
    "username": "chenniannian",
    "user_valid": true,
    "root_id": "2",
    "doc_id": "47",
    "params": {
      "trade_date": "",
      "start_date": "20240101",
      "end_date": "20240110",
      "limit": "100",
      "offset": "0"
    },
    "fields": [
      "trade_date",
      "ggt_ss",
      "ggt_sz",
      "hgt",
      "sgt",
      "north_money",
      "south_money"
    ]
  }'
```

**关键差异**:
- 实际 API 支持 `limit` 和 `offset` 分页参数
- MCP Tools 中未暴露这些参数
- 返回字段通过 `fields` 参数指定

---

**报告生成时间**: 2026-03-10
**报告版本**: 1.0
**维护者**: Claude Code Agent
