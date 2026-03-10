# Tushare Stock Financial API 测试 - 图表展示版本

这是一个增强版的财务 API 测试程序，提供**可视化图表展示**功能。

## ✨ 特性

- 📊 **终端图表** - 在终端直接显示美观的柱状图
- 🌐 **HTML报告** - 自动生成专业的 HTML 可视化报告
- 📈 **统计分析** - 详细的统计数据和成功率分析
- 🎨 **多维度展示** - 支持按数据量、成功率等多维度展示

## 🚀 快速开始

### 1. 设置 Token

```bash
export TUSHARE_TOKEN="your_token_here"
```

### 2. 运行测试

```bash
# 方式1: 使用脚本（推荐）
TUSHARE_TOKEN="your_token" ./run.sh

# 方式2: 直接运行
TUSHARE_TOKEN="your_token" go run main.go
```

## 📊 输出示例

### 终端输出

程序会在终端显示：

```
========================================
  Tushare Stock Financial API 测试
  📊 图表展示模式
========================================

🚀 开始测试所有 API...

[1/10] 测试 Balancesheet (资产负债表)...
[2/10] 测试 Cashflow (现金流量表)...
...
✅ 所有测试完成!

================================================================================
📊 数据返回情况图表
================================================================================

 1. Balancesheet        ✅    4 条 |████████████████████
 2. Income              ✅    4 条 |████████████████████
 3. Cashflow            ✅    4 条 |████████████████████
 4. Dividend            ✅   10 条 |████████████████████████████████
 5. FinaIndicator       ✅    4 条 |████████████████████
 6. FinaAudit           ✅    2 条 |█████████████
 7. FinaMainbz          ✅    3 条 |████████████████
 8. DisclosureDate      ⚠️     空  |
 9. Forecast            ⚠️     空  |
10. Express             ⚠️     空  |

图例:
  ✅ 成功返回数据
  ⚠️  成功但无数据
  ❌ 调用失败
  █  数据量

================================================================================
📋 详细测试报告
================================================================================

📈 统计摘要:
  • 测试总数: 10
  • 成功: 7 (70.0%)
  • 失败: 0 (0.0%)
  • 空数据: 3 (30.0%)
  • 总记录数: 31

📊 各API返回记录数:
  • Balancesheet        :    4 条
  • Income              :    4 条
  • Cashflow            :    4 条
  • Dividend            :   10 条
  • FinaIndicator       :    4 条
  • FinaAudit           :    2 条
  • FinaMainbz          :    3 条
  • DisclosureDate      :    0 条 (空)
  • Forecast            :    0 条 (空)
  • Express             :    0 条 (空)

✅ HTML报告已生成: test_report.html
   请在浏览器中打开查看可视化报告
```

### HTML 报告

程序会自动生成 `test_report.html`，包含：

- 📊 统计摘要卡片（成功、失败、空数据、总记录数）
- 📈 可视化柱状图（彩色渐变效果）
- 📋 详细结果表格（带状态标签）
- 📝 API 数据可用性说明

在浏览器中打开查看效果：

```bash
# macOS
open test_report.html

# Linux
xdg-open test_report.html

# Windows
start test_report.html
```

## 🎨 图表功能说明

### 1. 终端柱状图

```
 1. Balancesheet        ✅    4 条 |████████████████████
    ↑                      ↑      ↑         ↑
  API名称                状态   记录数    可视化柱
```

- **状态图标**：
  - ✅ 成功返回数据
  - ⚠️  成功但无数据
  - ❌ 调用失败

- **柱状图**：
  - 每个 `█` 代表一定数量的记录
  - 长度与最大值成比例
  - 便于直观比较各 API 的数据量

### 2. HTML 报告图表

- **渐变色彩**：不同状态使用不同渐变色
  - 成功：绿色渐变 (#11998e → #38ef7d)
  - 失败：红色渐变 (#eb3349 → #f45c43)
  - 空数据：橙色渐变 (#f093fb → #f5576c)

- **交互效果**：
  - 鼠标悬停高亮
  - 平滑过渡动画
  - 响应式布局

## 📈 数据可用性分析

根据测试结果，各 API 的数据可用性：

### ⭐⭐⭐⭐⭐ 高可用性（几乎总是有数据）
- **Balancesheet** - 资产负债表
- **Income** - 利润表
- **Cashflow** - 现金流量表

### ⭐⭐⭐⭐ 中等可用性（取决于公司行为）
- **Dividend** - 分红数据
- **FinaIndicator** - 财务指标
- **FinaAudit** - 审计意见

### ⭐⭐⭐ 低可用性（季节性或特定条件）
- **DisclosureDate** - 财报披露日期
- **FinaMainbz** - 主营业务构成
- **Forecast** - 业绩预告

### ⭐⭐ 极低可用性（较少发布）
- **Express** - 业绩快报

详细说明请参考：[DATA_AVAILABILITY.md](../test-stock-financial/DATA_AVAILABILITY.md)

## 🔧 自定义测试

### 修改测试股票

编辑 `main.go`，找到测试函数并修改 `ts_code`：

```go
// 测试茅台
req := &stock_financial.BalancesheetRequest{
    TsCode:    "600519.SH",  // 贵州茅台
    StartDate: "20240101",
    EndDate:   "20241231",
}
```

### 修改时间范围

```go
// 测试近三年数据
req := &stock_financial.BalancesheetRequest{
    TsCode:    "000001.SZ",
    StartDate: "20220101",  // 从2022年开始
    EndDate:   "20241231",
}
```

### 推荐的测试股票

| 股票代码 | 股票名称 | 特点 | 推荐理由 |
|---------|---------|------|---------|
| 000001.SZ | 平安银行 | 银行股 | 数据完整、分红稳定 |
| 600519.SH | 贵州茅台 | 蓝筹股 | 业绩稳定、数据丰富 |
| 600000.SH | 浦发银行 | 银行股 | 历史数据长 |
| 601398.SH | 工商银行 | 大盘银行 | 数据最完整 |

## 🛠️ 故障排除

### 问题1: 编译错误

```
error: undefined: stock_financial
```

**解决方案**：确保在正确的目录

```bash
cd cmd/test-stock-financial-chart
go run main.go
```

### 问题2: HTML 报告无法打开

**解决方案**：
1. 检查文件是否生成：`ls -la test_report.html`
2. 使用绝对路径打开
3. 尝试其他浏览器

### 问题3: 图表显示异常

**解决方案**：
1. 确保终端支持 UTF-8 字符
2. 尝试不同的终端应用
3. 使用 HTML 报告查看

## 📊 输出文件

运行后会生成以下文件：

```
test-stock-financial-chart/
├── main.go              # 主程序
├── run.sh              # 运行脚本
├── README.md           # 本文档
└── test_report.html    # 生成的 HTML 报告
```

## 🎯 使用场景

1. **API 测试** - 验证所有财务 API 是否正常工作
2. **数据评估** - 评估不同股票的数据可用性
3. **性能监控** - 监控 API 返回数据量变化
4. **报告生成** - 生成专业的测试报告
5. **数据对比** - 对比不同股票或时间段的数据

## 💡 最佳实践

1. **定期测试** - 每月运行一次，监控 API 状态
2. **多股票测试** - 测试多只股票获得全面评估
3. **保存报告** - 保存 HTML 报告用于历史对比
4. **关注空数据** - 空数据可能表示业务变化

## 📚 相关文档

- [Tushare 官方文档](https://tushare.pro/document/2)
- [简单测试版本](../test-stock-financial/)
- [数据可用性说明](../test-stock-financial/DATA_AVAILABILITY.md)
- [项目主文档](../../../README.md)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！
