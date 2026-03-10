# MCP 工具展示程序

这是一个用���展示 tushare-go 项目中所有可用的 MCP 工具的演示程序。

## 功能特性

✅ **工具统计**: 显示每个类别的工具数量
✅ **状态显示**: 标注哪些工具已启用，哪些未启用
✅ **分类展示**: 按类别组织所有工具
✅ **详细信息**: 显示每个工具的名称和描述

## 使用方法

### 运行展示程序

```bash
go run cmd/show-mcp-tools/main.go
```

### 输出示例

程序会显示：
1. 📊 MCP 工具统计 - 按类别统计工具数量
2. 📋 所有可用的 MCP 工具 - 详细的工具列表
3. 💡 使用提示 - 如何运行实际测试

## MCP 工具概览

当前支持的 11 个工具类别：

| 类别 | 工具数量 | 说明 |
|------|---------|------|
| Stock Basic | 6 | 股票基础信息 |
| Stock Market | 5 | 股票行情数据 |
| Stock Financial | 6 | 股票财务数据 |
| Stock Board | 4 | 股票板块数据 |
| Stock Feature | 3 | 股票特征数据 |
| Bond | 3 | 债券数据 |
| Fund | 3 | 基金数据 |
| Index | 3 | 指数数据 |
| HK Stock | 2 | 港股数据 |
| US Stock | 2 | 美股数据 |
| ETF | 2 | ETF 数据 |

**总计**: 39 个工具

## 工具状态

- ✅ 已启用 (37个) - 工具已注���到 MCP 服务器
- ❌ 未启用 (2个) - 工具存在但未注册（可能需要额外权限）

## 运行实际测试

查看这些工具的实际测试结果：

```bash
# 1. 设置 Token
export TUSHARE_TOKEN=你的token

# 2. 运行测试
make -C cmd/test-mcp-tools test

# 3. 查看报告
make -C cmd/test-mcp-tools report
```

## 相关文档

- [测试程序说明](../test-mcp-tools/README.md)
- [快速开始指南](../test-mcp-tools/QUICKSTART.md)
- [MCP 工具文档](../../../docs/MCP_TOOLS.md)

## 维护

- **创建时间**: 2026-03-10
- **维护者**: tushare-go 项目组
