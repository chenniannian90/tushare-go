# MCP 工具测试程序

## 概述

这是一个全面的 MCP 工具测试程序，用于测试 tushare-go 项目中所有可用的 MCP 工具。

## 支持的工具类别

本测试程序支持测试以下 11 个类别的 MCP 工具：

1. **Stock Basic** (股票基础) - 6 个工具
   - stock_basic - 股票基本信息
   - new_share - 新股数据
   - namechange - 股票名称变更
   - trade_cal - 交易日历
   - stock_company - 上市公司信息
   - stk_managers - 高级管理人员

2. **Stock Market** (股票行情) - 5 个工具
   - daily - 日线行情
   - daily_basic - 日线基本面
   - weekly - 周线行情
   - monthly - 月线行情
   - adj_factor - 复权因子

3. **Stock Financial** (股票财务) - 6 个工具
   - income - 利润表
   - balancesheet - 资产负债表
   - cashflow - 现金流量表
   - fina_indicator - 财务指标
   - dividend - 分红送股
   - forecast - 业绩预告

4. **Stock Board** (股票板块) - 4 个工具
   - stk_auction - 集合竞价
   - hm_detail - 游资营业部明细
   - top_list - 龙虎榜
   - limit_list_d - 大宗交易

5. **Stock Feature** (股票特征) - 3 个工具
   - hk_hold - 港股持股
   - stk_auction_c - 竞价详情
   - report_rc - 研报

6. **Bond** (债券) - 3 个工具
   - cb_basic - 可转债基本信息
   - cb_daily - 可转债日线
   - bond_oc - 柜台债券

7. **Fund** (基金) - 3 个工具
   - fund_basic - 基金基本信息
   - fund_nav - 基金净值
   - fund_manager - 基金经理

8. **Index** (指数) - 3 个工具
   - index_basic - 指数基本信息
   - index_daily - 指数日线
   - index_member - 指数成分股

9. **HK Stock** (港股) - 2 个工具
   - hk_basic - 港股基本信息
   - hk_daily - 港股日线

10. **US Stock** (美股) - 2 个工具
    - us_basic - 美股基本信息
    - us_daily - 美股日线

11. **ETF** - 2 个工具
    - etf_basic - ETF 基本信息
    - fund_daily - ETF 日线

**总计**: 39 个 MCP 工具

## 使用方法

### 1. 设置环境变量

```bash
export TUSHARE_TOKEN=你的token
```

### 2. 运行测试程序

**方法 1: 使用 Makefile（推荐）**

```bash
# 编译并运行测试
make test

# 或者直接运行（不编译）
make run

# 快速测试
make quick

# 查看所有可用命令
make help
```

**方法 2: 直接运行**

```bash
# 方法 1: 直接运行
go run cmd/test-mcp-tools/main.go

# 方法 2: 先编译后运行
go build -o bin/test-mcp-tools cmd/test-mcp-tools/main.go
./bin/test-mcp-tools

# 方法 3: 使用测试脚本
chmod +x cmd/test-mcp-tools/test.sh
./cmd/test-mcp-tools/test.sh
```

### 3. 查看测试结果

测试程序会：

1. 在终端显示实时测试进度
2. 显示每个工具的测试状态（成功/失败/无数据）
3. 生成详细的测试报告并保存到 `tests/mcp-tools-test-report.md`

查看报��：

```bash
# 使用 Makefile
make report

# 或者直接查看
cat tests/mcp-tools-test-report.md
```

## 测试输出

### 控制台输出

```
🚀 开始测试 Tushare MCP 工具...
============================================================

📊 测试 Stock Basic 工具...
  测试 stock_basic...
  测试 new_share...
  测试 namechange...
  ...

📈 测试 Stock Market 工具...
  测试 daily...
  ...

======================================================================
                          测试报告
======================================================================

类别                  总计   成功     有数据    失败     耗时
----------------------------------------------------------------------
Stock Basic           6      6        6        0       1.2s
Stock Market          5      5        5        0       800ms
...
```

### 测试报告

测试报告包含：

1. **测试结果汇总表** - 按类别统计测试结果
2. **详细测试结果** - 每个工具的详细测试信息
3. **测试说明** - 状态说明和测试参数
4. **常见测试参数** - 使用的测试数据说明

## 测试参数说明

程序使用以下默认参数进行测试：

- **股票代码**: 000001.SZ (平安银行)
- **交易日期**: 2024-03-08
- **日期范围**: 2024-03-01 至 2024-03-10
- **报告期**: 20231231 (2023年年报)
- **可转债代码**: 113001.SZ (东财转债)
- **ETF代码**: 510300.SH (沪深300ETF)
- **港股代码**: 00700.HK (腾讯控股)
- **美股代码**: AAPL (苹果)

## 扩展测试

### 添加新的测试工具

1. 在对应的测试类别中添加测试函数
2. 在测试列表中注册新的测试
3. 实现具体的测试函数

示例：

```go
// 在 testStockBasicTools 中添加
{"stock_company", testStockCompany},

// 实现测试函数
func testStockCompany(ctx context.Context, client *sdk.Client) TestResult {
    start := time.Now()
    req := &stock_basicapi.StockCompanyRequest{
        TsCode: "000001.SZ",
    }
    data, err := stock_basicapi.StockCompany(ctx, client, req)
    return buildResult("stock_company", data, err, time.Since(start))
}
```

### 自定义测试参数

可以修改测试函数中的请求参数来测试不同的数据场景。

## 注意事项

1. **API 限制**: 某些 API 可能需要升级 Tushare 账户权限才能访问
2. **数据时效**: 使用的历史日期数据可能已过期
3. **网络连接**: 需要稳定的网络连接访问 Tushare API
4. **Token 配额**: 每次测试会消耗一定的 API 调用配额

## 相关文档

- [MCP 工具文档](../../docs/MCP_TOOLS.md)
- [Stock Board 最佳实践](../../docs/STOCK_BOARD_MCP_BEST_PRACTICES.md)
- [Tushare 官方文档](https://tushare.pro)

## 维护

- **创建时间**: 2026-03-10
- **维护者**: tushare-go 项目组
- **版本**: 1.0.0
