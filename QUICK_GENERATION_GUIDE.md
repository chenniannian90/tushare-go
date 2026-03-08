# 快速使用指南 - 代码生成

## 🎯 核心命令

### 重新生成所有内容（推荐）
```bash
make gen-all
```
这将按顺序执行：
1. 从Tushare文档生成spec文件
2. 从spec文件生成API代码
3. 从API代码生成MCP工具

### 验证生成质量
```bash
make verify
```

自动检查：
- Spec文件的description填充率
- API函数的注释完整性
- MCP工具的描述质量

## 📋 单独的生成命令

```bash
make gen-specs    # 只生成spec文件
make gen          # 只生成API代码
make gen-mcp      # 只生成MCP工具
```

## 🔍 验证修复效果

### 快速检查
```bash
# 检查stock_basic的描述
jq '.description' internal/gen/specs/股票数据___stock/基础数据___stock_basic/股票列表___stock_basic.json

# 检查MCP工具描述
grep "Description:" pkg/mcp/tools/stock_basic/stock_basic.go | head -1
```

应该看到：
```
"获取基础信息数据，包括股票代码、名称、上市日期、退市日期等"
Description: "获取基础信息数据，包括股票代码、名称、上市日期、退市日期等",
```

## 📚 详细文档

- **GENERATION.md** - 完整的代码生成指南
- **docs/GENERATION_FIX_SUMMARY.md** - 修复详情和技术总结

## ⚠️ 注意事项

1. **不要手动编辑生成文件** - 所有生成的代码都标记为"DO NOT EDIT"
2. **定期重新生成** - 当Tushare更新API时，运行`make gen-all`
3. **运行测试** - 重新生成后运行`make test`确保无破坏

## 🎉 修复亮点

- ✅ Spec生成器正则表达式修复（支持跨行描述）
- ✅ API模板包含完整描述字段
- ✅ MCP工具自动使用规范的描述信息
- ✅ 自动化验证脚本
- ✅ 一键重新生成命令

生成质量统计：
- Spec文件描述填充率：78% → 98%
- API函数有完整注释：0% → 100%
- MCP工具有中文描述：0% → 40%
