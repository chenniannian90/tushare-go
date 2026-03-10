# Stock Margin 模块 - 所有 API 验证报告

**验证日期**: 2026-03-10
**模块**: stock_margin (融资融券及转融通)
**API Token**: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1

---

## 📊 总体概览

stock_margin 模块包含 **4 个 API**，全部已验证通过！

| API | 功能 | 状态 | 测试 | 标的数量/记录数 |
|-----|------|------|------|-----------------|
| **margin** | 融资融券交易汇总 | ✅ 已修复 | ✅ 通过 | 3个交易所 |
| **margin_detail** | 融资融券交易明细 | ✅ 已修复 | ✅ 通过 | 3843条记录 |
| **margin_secs** | 融资融券标的 | ✅ 原本完整 | ✅ 通过 | 4026只标的 |
| **slb_sec** | 转融券标的 | ✅ 原本完整 | ✅ 通过 | 2640只标的 |

---

## 🔧 修复记录

### 1. margin - 融资融券交易汇总

**问题**: `response_fields: null`
**修复**: 添加 9 个响应字段
**提交**: `feat(margin): add missing response fields to margin API`

**字段**:
- trade_date, exchange_id, rzye, rzmre, rzche, rqye, rqmcl, rzrqye, rqyl

**数据示例（2024-01-05）**:
```
SSE: 融资8287.10亿, 融券445.98亿, 总��8733.07亿
SZSE: 融资7498.11亿, 融券246.16亿, 总计7744.28亿
BSE: 融资12.65亿, 融券0.01亿, 总计12.67亿
```

---

### 2. margin_detail - 融资融券交易明细

**问题**: `response_fields: null`
**修复**: 添加 10 个响应字段
**提交**: `feat(margin_detail): add missing response fields to margin_detail API`

**字段**:
- trade_date, ts_code, rzye, rqye, rzmre, rqyl, rzche, rqchl, rqmcl, rzrqye

**数据示例（600000.SH - 浦发银行）**:
```
20240110: 融资30.43亿, 融券496.34万, 融券余量755470股
20240109: 融资30.41亿, 融券479.34万, 融券余量725170股
```

---

### 3. margin_secs - 融资融券标的

**状态**: ✅ 原本完整，无需修复
**验证**: 创建测试程序验证功能
**提交**: `test(margin): add margin_secs API test program`

**字段**:
- trade_date, ts_code, name, exchange

**数据统计（2024-01-10）**:
```
总计: 4026只融资融券标的
SSE (上交所):  1956只 (48.6%)
SZSE (深交所): 1830只 (45.5%)
BSE (北交所):   240只 (5.9%)
```

---

### 4. slb_sec - 转融券标的

**状态**: ✅ 原本完整，无需修复
**验证**: 创建测试程序验证功能
**提交**: `test(slb_sec): add slb_sec API test program`

**字段**:
- trade_date, ts_code, name, ope_inv, lent_qnt, cls_inv, end_bal

**数据统计（2024-01-10）**:
```
总计: 2640只转融券标的
期初余量: 666,693.46万股
融出数量: 15,259.92万股
期末余量: 653,556.98万股
期末余额: 96.94亿元
```

---

## 📈 市场数据分析

### 融资融券市场（2024-01-05）

**全市场规模**:
- 总计: **16,490.02亿元**
- 融资余额: 15,797.86亿元 (95.8%)
- 融券余额: 692.15亿元 (4.2%)

**交易所分布**:
| 交易所 | 融资融券余额 | 占比 |
|--------|--------------|------|
| SSE | 8,733.07亿 | 51.6% |
| SZSE | 7,744.28亿 | 45.7% |
| BSE | 12.67亿 | 0.1% |

**融资融券标的分布**:
| 交易所 | 标的数量 | 占比 |
|--------|----------|------|
| SSE | 1,956只 | 48.6% |
| SZSE | 1,830只 | 45.5% |
| BSE | 240只 | 5.9% |

### 转融券市场（2024-01-10）

**市场规模**:
- 期末余额: **96.94亿元**
- 期初余量: 666,693.46万股
- 融出数量: 15,259.92万股
- 期末余量: 653,556.98万股

**标的数量**: 2,640只

**个股示例（600000.SH 浦发银行）**:
```
最近5个交易日转融券情况:
20240108: 融出12万股, 期末420万股
20240105: 融出0万股, 期末415万股
20240104: 融出3万股, 期末415万股
```

---

## 🎯 API 使用场景

### margin - 交易所汇总数据

**适用场景**:
- 查看各交易所融资融券总体规模
- 分析市场整体融资融券趋势
- 对比不同交易所的活跃度

**使用示例**:
```go
req := &stock_margin.MarginRequest{
    TradeDate: "20240110",
}
items, _ := stock_margin.Margin(ctx, client, req)
```

### margin_detail - 个股明细数据

**适用场景**:
- 查询特定股票的融资融券情况
- 分析个股融资买入/偿还趋势
- 跟踪主力资金动向

**使用示例**:
```go
req := &stock_margin.MarginDetailRequest{
    TsCode:    "600000.SH",
    StartDate: "20240101",
    EndDate:   "20240110",
}
items, _ := stock_margin.MarginDetail(ctx, client, req)
```

### margin_secs - 融资融券标的查询

**适用场景**:
- 查询某日所有融资融券标的
- 验证股票是否为融资融券标的
- 按交易所筛选标的

**使用示例**:
```go
req := &stock_margin.MarginSecsRequest{
    Exchange:  "SSE",
    TradeDate: "20240110",
}
items, _ := stock_margin.MarginSecs(ctx, client, req)
```

### slb_sec - 转融券标的查询

**适用场景**:
- 查询转融券标的列表
- 查询个股转融券交易数据
- 分析转融券融出情况

**使用示例**:
```go
req := &stock_margin.SlbSecRequest{
    TsCode:    "600000.SH",
    StartDate: "20240101",
    EndDate:   "20240110",
}
items, _ := stock_margin.SlbSec(ctx, client, req)
```

---

## 📝 字段缩写对照表

### 融资融券相关

| 缩写 | 全称 | 说明 |
|------|------|------|
| rz | 融资 | 融资 |
| rq | 融券 | 融券 |
| ye | 余额 | 余额 |
| mre | 买入额 | 买入额 |
| che | 偿还额 | 偿还额 |
| mcl | 卖出量 | 卖出量 |
| yl | 余量 | 余量 |
| chl | 偿还量 | 偿还量 |
| rzrq | 融资融券 | 融资融券 |

### 转融券相关

| 缩写 | 全称 | 说明 |
|------|------|------|
| slb | securities lending | 转融券 |
| sec | securities | 证券/标的 |
| ope_inv | opening inventory | 期初余量 |
| lent_qnt | lent quantity | 融出数量 |
| cls_inv | closing inventory | 期末余量 |
| end_bal | ending balance | 期末余额 |

---

## 🧪 测试程序

### 已创建的测试程序

1. **cmd/test-margin/main.go**
   - 测试 margin 和 margin_detail API
   - 验证交易所汇总和个股明细

2. **cmd/test-margin-secs/main.go**
   - 测试 margin_secs API
   - 验证融资融券标的查询

3. **cmd/test-slb-sec/main.go**
   - 测试 slb_sec API
   - 验证转融券标的查询

### 运行测试

```bash
# 设置环境变量
export TUSHARE_TOKEN="your_token_here"

# 运行所有测试
go run cmd/test-margin/main.go
go run cmd/test-margin-secs/main.go
go run cmd/test-slb-sec/main.go
```

---

## ✅ 验证结论

### 所有 API 状态

| API | 问题 | 修复状态 | 测试状态 | 最终状态 |
|-----|------|----------|----------|----------|
| margin | 缺少字段定义 | ✅ 已修复 | ✅ 通过 | ✅ 可用 |
| margin_detail | 缺少字段定义 | ✅ 已修复 | ✅ 通过 | ✅ 可用 |
| margin_secs | 无问题 | ✅ 无需修复 | ✅ 通过 | ✅ 可用 |
| slb_sec | 无问题 | ✅ 无需修复 | ✅ 通过 | ✅ 可用 |

### 总体评价

✅ **stock_margin 模块所有 4 个 API 全部可用！**

- **代码质量**: 所有 API 使用健壮的类型转换逻辑
- **测试覆盖**: 100% 测试通过率
- **数据准确性**: 真实市场数据验证
- **文档完整性**: 详细的字段说明和示例

---

## 📚 相关资源

### 官方文档

- [margin - 融资融券交易汇总](https://tushare.pro/document/2?doc_id=58)
- [margin_detail - 融资融券交易明细](https://tushare.pro/document/2?doc_id=59)
- [margin_secs - 融资融券标的](https://tushare.pro/document/2?doc_id=326)
- [slb_sec - 转融券标的](https://tushare.pro/document/2?doc_id=332)

### 项目文档

- `STOCK_MARGIN_FIXES_SUMMARY.md` - 详细修复报告
- 测试程序: `cmd/test-margin/`, `cmd/test-margin-secs/`, `cmd/test-slb-sec/`

---

## 🎉 总结

通过用户提供的实际 API 调用示例，我们成功：

1. ✅ 发现并修复了 2 个 API 的 spec 文件问题
2. ✅ 验证了 2 个原本完整的 API
3. ✅ 创建了完整的测试程序
4. ✅ 获取并分析了真实市场数据
5. ✅ 提供了详细的使用文档

**stock_margin 模块现已完全可用，可以放心使用！**

---

**报告生成时间**: 2026-03-10
**报告版本**: 1.0
**维护者**: Claude Code Agent
**状态**: 全部验证通过 ✅
