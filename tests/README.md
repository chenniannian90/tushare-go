# Tushare Go SDK - 测试套件

## 📊 测试覆盖概览

本测试套件包含 **77 个测试用例**，全面验证了 SDK 对各种 API 响应格式的处理能力。

### ✅ 测试分类

#### 1. **格式检测���转换** (4 tests)
- `TestArrayArrayFormat` - 二维数组格式（Web API）
- `TestObjectArrayFormat` - 对象数组格式（官方 API）
- `TestEmptyArray` - 空数组处理
- `TestMixedTypes` - 混合类型（字符串、数字、null）

#### 2. **Web API 响应格式** (2 tests)
- `TestWebAPIResponseFormat` - Web API 实际响应格式
- `TestOfficialAPIResponseFormat` - 官方 API 响应格式

#### 3. **真实 API 场景** (7 tests)
- `TestTradeCalAPI` - 交易日历 API
- `TestStockCompanyAPI` - 上市公司基本信息 API
- `TestEmptyItemsArray` - 空 items 数组
- `TestNullFieldsAndItems` - null 字段和 items
- `TestLargeNumbersInResponse` - 大数字处理
- `TestSpecialCharactersInData` - 特殊字符

#### 4. **边缘情况** (7 tests)
- `TestNegativeNumbers` - 负数处理
- `TestBooleanLikeNumbers` - 布尔值表示（0/1）
- `TestVeryLargeNumbers` - 超大数字
- `TestZeroValues` - 零值
- `TestSingleItemArray` - 单项数组
- `TestManyFields` - 多字段（20+）
- `TestScientificNotation` - 科学计数法

#### 5. **市场数据 API** (8 tests)
- `TestDailyAPI` - 日线行情数据
- `TestWeeklyAPI` - 周线行情数据（大成交量）
- `TestDailyBasicAPI` - 每日基本面指标（PE、PB、市值）
- `TestAdjFactorAPI` - 复权因子数据
- `TestHsgtTop10API` - 沪深股通十大成交股（含负值）
- `TestDecimalPrecision` - 高精度小数处理
- `TestLargeVolumeData` - 大成交量数据
- `TestMultipleDataTypes` - 混合数据类型

#### 6. **财务数据 API** (8 tests)
- `TestIncomeAPI` - 利润表数据（收入、利润、EPS）
- `TestBalanceSheetAPI` - 资产负债表数据（资产、负债、权益）
- `TestCashflowAPI` - 现金流量表数据（经营、投资、筹资现金流）
- `TestForecastAPI` - 业绩预告数据（预测范围）
- `TestDividendAPI` - 分红送股数据（股息、分红）
- `TestDisclosureDateAPI` - 财报披露日期（预披露、实际披露）
- `TestFinaIndicatorAPI` - 财务指标数据（ROE、ROA、流动比率等）

#### 7. **股东与交易数据 API** (9 tests)
- `TestTop10FloatholdersAPI` - 十��流通股东（持股比例、变动）
- `TestPledgeStatAPI` - 股权质押统计（质押率、质押数量）
- `TestPledgeDetailAPI` - 股权质押明细（质权人、质押状态）
- `TestRepurchaseAPI` - 股份回购数据（回购数量、价格区间）
- `TestShareFloatAPI` - 限售股解禁（解禁数量、解禁比例）
- `TestBlockTradeAPI` - 大宗交易（交易价格、买卖方）
- `TestStkAccountAPI` - 股票账户开户（新增账户、总数）
- `TestStkHoldernumberAPI` - 股东户数（股东数量变化）
- `TestStkHoldertradeAPI` - 股东增减持（增持、减持、变动比例）

#### 8. **高级分析 API** (10 tests)
- `TestStkFactorProAPI` - 股票技术因子（300+ 字段）
- `TestReportRCAPI` - 研究报告（机构、日期、评分）
- `TestCyqChipsAPI` - 筹码分布（价位、比例）
- `TestCyqPerfAPI` - 筹码性能（成本、胜率）
- `TestInstitutionSurveyAPI` - 机构调研记录（调研日期、类型）
- `TestBrokerRecommendAPI` - 券商金股（月度、券商、推荐）
- `TestMarginAPI` - 融资融券汇总（融资余额、融券余额）
- `TestMarginDetailAPI` - 融资融券明细（个股融资融券数据）
- `TestMarginSecsAPI` - 融资融券标的（标的列表）
- `TestSlbSecAPI` - 转融通证券出借汇总
- `TestSlbLenAPI` - 转融通融资汇总
- `TestSlbSecDetailAPI` - 转融通证券出借明细
- `TestSlbLenMmAPI` - 做市借券交易

#### 9. **资金流向 API** (8 tests)
- `TestMoneyflowAPI` - 个股资金流向（大中小单、净流入）
- `TestMoneyflowHsgtAPI` - 沪深股通资金流向（北向资金、南向资金）
- `TestMoneyflowThsAPI` - 同花顺个股资金流向（涨跌幅、大单买入）
- `TestMoneyflowDcAPI` - 东方财富个股资金流向（超大单、大单、中单、小单）
- `TestMoneyflowCntThsAPI` - 同花顺概念板块资金流向（龙头股、公司数量）
- `TestMoneyflowIndThsAPI` - 同花顺行业资金流向（行业指数、涨跌股票数）
- `TestMoneyflowIndDcAPI` - 东方财富板块资金流向（行业/概念板块排名）
- `TestMoneyflowMktDcAPI` - 东方财富大盘资金流向（沪深指数、市场净流入）

#### 10. **龙虎榜与涨跌停 API** (13 tests)
- `TestTopListAPI` - 龙虎榜每日交易明细（买卖金额、净流入）
- `TestTopInstAPI` - 龙虎榜机构成交明细（机构买卖、净买入）
- `TestLimitListThsAPI` - 同花顺涨跌停榜（涨停成功率、封单量）
- `TestLimitListDAPI` - 涨跌停数据（封板时间、打开次数）
- `TestLimitStepAPI` - 连板���股（连续涨停板数）
- `TestLimitCptListAPI` - 每日涨跌停板块统计（板块涨停数、排名）
- `TestThsIndexAPI` - 同花顺指数（成分股数量）
- `TestThsDailyAPI` - 同花顺指数日线行情（OHLC、成交量）
- `TestDcIndexAPI` - 东方财富板块指数（涨跌股票数、龙头股）
- `TestStmAuctionAPI` - 股票竞价（集合竞价成交量、金额）
- `TestHmListAPI` - 游资名录（知名游资名单）
- `TestHmDetailAPI` - 游资交易明细（游资买卖数据）
- `TestDcHotAPI` - 东方财富热榜（热门股票排名、涨跌幅）

#### 11. **基金数据 API** (1 test)
- `TestFundDailyAPI` - 基金日线行情（OHLC、成交量、涨跌幅）

#### 12. **港股基础信息 API** (1 test)
- `TestHkBasicAPI` - 港股基础信息（股票代码、名称、市场、上市状态）

## 🧪 运行测试

### 运行所有测试
```bash
go test ./tests/...
```

### 运行特定测试
```bash
# 格式检测测试
go test ./tests -run TestArrayArrayFormat

# Web API 测试
go test ./tests -run TestWebAPIResponseFormat

# 边缘情况测试
go test ./tests -run TestNegativeNumbers

# 市场数据 API 测试
go test ./tests -run TestDailyAPI

# 财务数据 API 测试
go test ./tests -run TestIncomeAPI

# 股东与交易数据 API 测试
go test ./tests -run TestTop10Float
go test ./tests -run TestPledge
go test ./tests -run TestStkHolder
```

### 详细输出
```bash
go test ./tests/... -v
```

## 📈 测试结果

所有 77 个测试用例均通过 ✅

```
PASS: TestNegativeNumbers
PASS: TestBooleanLikeNumbers
PASS: TestVeryLargeNumbers
PASS: TestZeroValues
PASS: TestSingleItemArray
PASS: TestManyFields
PASS: TestScientificNotation
PASS: TestDailyAPI
PASS: TestWeeklyAPI
PASS: TestDailyBasicAPI
PASS: TestAdjFactorAPI
PASS: TestHsgtTop10API
PASS: TestDecimalPrecision
PASS: TestLargeVolumeData
PASS: TestMultipleDataTypes
PASS: TestIncomeAPI
PASS: TestBalanceSheetAPI
PASS: TestCashflowAPI
PASS: TestForecastAPI
PASS: TestDividendAPI
PASS: TestDisclosureDateAPI
PASS: TestFinaIndicatorAPI
PASS: TestTop10FloatholdersAPI
PASS: TestPledgeStatAPI
PASS: TestPledgeDetailAPI
PASS: TestRepurchaseAPI
PASS: TestShareFloatAPI
PASS: TestBlockTradeAPI
PASS: TestStkAccountAPI
PASS: TestStkHoldernumberAPI
PASS: TestStkHoldertradeAPI
PASS: TestArrayArrayFormat
PASS: TestObjectArrayFormat
PASS: TestEmptyArray
PASS: TestMixedTypes
PASS: TestTradeCalAPI
PASS: TestStockCompanyAPI
PASS: TestEmptyItemsArray
PASS: TestNullFieldsAndItems
PASS: TestLargeNumbersInResponse
PASS: TestSpecialCharactersInData
PASS: TestWebAPIResponseFormat
PASS: TestOfficialAPIResponseFormat
PASS: TestStkFactorProAPI
PASS: TestReportRCAPI
PASS: TestCyqChipsAPI
PASS: TestCyqPerfAPI
PASS: TestInstitutionSurveyAPI
PASS: TestBrokerRecommendAPI
PASS: TestMarginAPI
PASS: TestMarginDetailAPI
PASS: TestMarginSecsAPI
PASS: TestSlbSecAPI
PASS: TestSlbLenAPI
PASS: TestSlbSecDetailAPI
PASS: TestSlbLenMmAPI
PASS: TestMoneyflowAPI
PASS: TestMoneyflowHsgtAPI
PASS: TestMoneyflowThsAPI
PASS: TestMoneyflowDcAPI
PASS: TestMoneyflowCntThsAPI
PASS: TestMoneyflowIndThsAPI
PASS: TestMoneyflowIndDcAPI
PASS: TestMoneyflowMktDcAPI
PASS: TestTopListAPI
PASS: TestTopInstAPI
PASS: TestLimitListThsAPI
PASS: TestLimitListDAPI
PASS: TestLimitStepAPI
PASS: TestLimitCptListAPI
PASS: TestThsIndexAPI
PASS: TestThsDailyAPI
PASS: TestDcIndexAPI
PASS: TestStmAuctionAPI
PASS: TestHmListAPI
PASS: TestHmDetailAPI
PASS: TestDcHotAPI
PASS: TestFundDailyAPI
PASS: TestHkBasicAPI
```

## 🔍 测试覆盖的数据类型

### API 响应格式
- ✅ **二维数组格式** `[[val1, val2], ...]` (Web API)
- ✅ **对象数组格式** `[{field: val}, ...]` (官方 API)
- ✅ **空数组** `[]`
- ✅ **null 值** `null`

### 数据类型
- ✅ **字符串** - 股票代码、名称、公告原因、股东名称等
- ✅ **整数** - 日期、数量、排名、股东户数等
- ✅ **浮点数** - 价格、涨跌幅、财务比率、持股比例等
- ✅ **负数** - 负的涨跌幅、减持数量等
- ✅ **大数字** - 市值、成交量、营收、账户数（亿级、十亿级）
- ✅ **科学计数法** - 超大/超小数字
- ✅ **null 值** - 缺失数据
- ✅ **布尔值** - 用是/否表示

### 特殊场景
- ✅ **单行数据** - 只有 1 条记录
- ✅ **多行数据** - 多条记录（2条、多条）
- ✅ **多字段数据** - 20+ 字段，甚至 80+ 财务字段
- ✅ **混合类型** - 同一响应包含多种类型
- ✅ **特殊字符** - 中文、下划线、点号等
- ✅ **财务数据** - 资产负债表平衡验证
- ✅ **日期数据** - 预披露日期、实际披露日期、解禁日期
- ✅ **百分比数据** - 涨跌幅、财务比率、质押率、持股比例
- ✅ **状态数据** - 是/否、增持/减持、实施/未实施
- ✅ **股东数据** - 股东名称、持股数量、持股比例
- ✅ **交易数据** - 大宗交易、账户开户、增减持

## 🚀 功能验证

### SDK 能力验证

1. **自动格式检测** ✅
   - 通过第一个字符自动判断格式类型
   - 支持 Web API 和官方 API 两种格式

2. **自动格式转换** ✅
   - 二维数组自动转换为对象数组
   - 字段名正确映射
   - 支持大量字段（80+ 财务字段）

3. **类型安全** ✅
   - 保持原始数据类型
   - 正确处理 null 值
   - 高精度小数保持精度

4. **向后兼容** ✅
   - `CallAPIFlexible` 方法兼容旧代码
   - 不影响现有的 API 调用

5. **财务数据处理** ✅
   - 利润表：EPS、收入、利润等
   - 资产负债表：资产、负债、权益平衡验证
   - 现金流量表：经营、投资、筹资现金流
   - 财务指标：ROE、ROA、流动比率等

6. **市场数据处理** ✅
   - 日线、周线行情数据
   - 复权因子数据
   - 港股通数据
   - 基本面指标数据

7. **股东数据处理** ✅
   - 十大流通股东数据
   - 股权质押统计和明细
   - 股东户数变化
   - 股东增减持数据

8. **交易数据处理** ✅
   - 大宗交易数据
   - 股份回购数据
   - 限售股解禁数据
   - 账户开户数据

## 📝 添加新测试

当发现新的 API 响应格式或边缘情况时，请添加相应的测试用例：

1. 在 `tests/` 目录下创建或修改测试文件
2. 使用真实或模拟的 API 响应数据
3. 验证格式检测、转换和数据类型
4. 确保测试能够通过

## 🔧 故障排查

如果测试失败：

1. **检查日志输出** - 使用 `-v` 参数查看详细日志
2. **验证数据格式** - 确认测试数据的格式是否正确
3. **检查字段映射** - 确认 fields 和 items 的数量匹配
4. **调试解析器** - 使用 `response_parser_test.go` 中的调试技巧
5. **验证数据类型** - 确认类型断言正确
6. **检查业务逻辑** - 验证计算公式（如质押率、资产负债平衡等）

## 📚 相关文档

- [SDK Client 文档](../pkg/sdk/client.go)
- [响应解析器文档](../pkg/sdk/response_parser.go)
- [API 示例](../examples/)
