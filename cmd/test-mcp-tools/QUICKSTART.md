# MCP 工具测试程序 - 快速开始

## 30 秒快速开始

```bash
# 1. 设置 Token
export TUSHARE_TOKEN=你的token

# 2. 运行测试
make test

# 3. 查看报告
make report
```

## 详细说明

### 功能特性

✅ **全面测试**: 覆盖 11 个类别的 39 个 MCP 工具
✅ **实时反馈**: 显示测试进度和状态
✅ **详细报告**: 生成 Markdown 格式的测试报告
✅ **错误处理**: 捕获并显示所有错误信息
✅ **性能统计**: 显示每个 API 调用的耗时

### 支持的工具类别

| 类别 | 工具数量 | 说明 |
|------|---------|------|
| Stock Basic | 6 | 股票基础信息 |
| Stock Market | 5 | 股票行情数据 |
| Stock Financial | 6 | 股票财务数据 |
| Stock Board | 4 | 股票板块数据 |
| Stock Feature | 3 | 股票特征数据 |
| Bond | 3 | 债券数据 |
| Fund | 3 | 基金数据 |
| Index | 3 | 指���数据 |
| HK Stock | 2 | 港股数据 |
| US Stock | 2 | 美股数据 |
| ETF | 2 | ETF 数据 |

### 测试输出示例

```
🚀 开始测试 Tushare MCP 工具...
============================================================

📊 测试 Stock Basic 工具...
  测试 stock_basic...
  测试 new_share...
  ...

📈 测试 Stock Market 工具...
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

### 常见问题

**Q: 测试失败怎么办？**

A: 检查以下几点：
1. Token 是否正确设置
2. 网络连接是否正常
3. Token 是否有足够的权限
4. 查看 `tests/mcp-tools-test-report.md` 了解详细错误

**Q: 某些工具返回无数据？**

A: 可能的原因：
1. 测试使用的历史日期数据已过期
2. 该 API 需要更高的账户权限
3. 测试参数不匹配

**Q: 如何测试特定的工具类别？**

A: 可以修改 `main.go` 中的测试流程，注释掉不需要测试的类别。

### 下一步

- 查看 [README.md](./README.md) 了解详细使用方法
- 查看 [MCP_TOOLS.md](../../docs/MCP_TOOLS.md) 了解所有可用的 MCP 工具
- 查看 [STOCK_BOARD_MCP_BEST_PRACTICES.md](../../docs/STOCK_BOARD_MCP_BEST_PRACTICES.md) 了解最佳实践

## 技术支持

- 项目主页: https://github.com/chenniannian90/tushare-go
- 问题反馈: https://github.com/chenniannian90/tushare-go/issues
