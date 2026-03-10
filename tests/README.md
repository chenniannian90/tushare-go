# Tushare MCP API 测试

## 📊 测试概览

本测试套件验证所有7个MCP服务类别的API接口可用性和数据返回情况。

### 🎯 测试的服务类别

1. **Stock** (股票市场数据) - 包含8个子类别
2. **Bond** (债券市场数据)
3. **Fund** (基金市场数据)
4. **Index** (指数数据)
5. **HK_Stock** (港股数据)
6. **US_Stock** (美股数据)
7. **ETF** (ETF数据)

### 🚀 运行测试

#### 方式1: 使用测试程序
```bash
go run cmd/api-test/main.go
```

#### 方式2: 编译后运行
```bash
make build
./bin/api-test
```

### 📈 最新测试结果

**测试时间**: 2026-03-09
**总测试数**: 9个接口

| 状态 | 数量 | 百分比 |
|------|------|--------|
| ✅ 成功 | 6 | 66.7% |
| 📈 有数据 | 3 | 33.3% |
| ⚠️ 无数据 | 3 | 33.3% |
| ❌ 失败 | 3 | 33.3% |

### 🔍 详细测试报告

完整的测试结果请查看: [mcp-api-test-report.md](./mcp-api-test-report.md)

### 📋 测试的API接口

#### Stock (股票数据)
- ✅ `stock_basic` - 股票基础信息 (有数据)
- ⚠️ `stock_market_daily` - 日线行情 (无数据)
- ❌ `stock_financial_income` - 财务数据 (类型错误)

#### Bond (债券数据)
- ⚠️ `bond_cb_basic` - 可转债基础信息 (无数据)

#### Fund (基金数据)
- ❌ `fund_fund_basic` - 基金基础信息 (类型错误)

#### Index (指数数据)
- ⚠️ `index_index_basic` - 指数基础信息 (无数据)

#### HK_Stock (港股数据)
- ✅ `hk_stock_hk_basic` - 港股基础信息 (2,722条数据)

#### US_Stock (美股数据)
- ✅ `us_stock_us_basic` - 美股基础信息 (5条数据)

#### ETF (ETF数据)
- ❌ `etf_etf_basic` - ETF基础信息 (权限不足)

### 🔧 测试参数说明

- **股票基础数据**: ts_code=000001.SZ (平安银行)
- **股票行情**: ts_code=000001.SZ, date_range=2024-03-01 to 2024-03-05
- **股票财务**: ts_code=000001.SZ, period=20231231
- **债券**: ts_code=113001.SZ (东财转债)
- **基金**: market=E (场内基金)
- **指数**: market=SSE (上交所指数)
- **港股**: list_status=L (上市港股)
- **美股**: limit=5 (前5只美股)
- **ETF**: exchange=SSE (上交所ETF)

### ⚠️ 已知问题

1. **数据类型错误**: 部分API存在数据类型转换问题
   - `stock_financial_income`: prem_earned 字段类型错误
   - `fund_fund_basic`: duration_year 字段类型错误

2. **权限限制**: 部分API需要更高权限的Tushare账户
   - `etf_etf_basic`: 需要升级Tushare账户权限

3. **参数优化**: 部分API测试参数需要调整以获取数据
   - `stock_market_daily`: 日期范围可能需要调整
   - `bond_cb_basic`: 可转债代码可能需要更新
   - `index_index_basic`: 指数参数可能需要优化

### 🔄 更新测试

如需更新测试或添加新的测试用例，请编辑:
- 测试程序: `cmd/api-test/main.go`
- 测试报告: 运行后会自动更新 `tests/mcp-api-test-report.md`

### 📞 支持

如有问题或建议，请查看:
- 项目主文档: https://github.com/chenniannian90/tushare-go
- Tushare官方文档: https://tushare.pro/document
- 问题反馈: https://github.com/chenniannian90/tushare-go/issues

---

**最后更新**: 2026-03-09
**测试工具**: tushare-go SDK
**测试环境**: Go 1.24+, Tushare API
