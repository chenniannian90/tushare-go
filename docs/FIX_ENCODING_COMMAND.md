# fix-encoding 命令使用说明

## 概述

`fix-encoding` 是一个自动检查和修复文件编码问题的命令行工具，专门用于处理 API 规范文件中的乱码问题。

## 功能

1. **检查文件名编码**: 检测文件名中的 UTF-8 编码问题
2. **修复文件名**: 自动修复包含乱码的文件名
3. **检查文件内容**: 检测 JSON 文件内容的编码问题
4. **修复文件内容**: 自动修复 JSON 文件中的编码问题
5. **清理无效字符**: 移除控制字符和其他无效字符

## 使用方法

### 方式1: 使用 Makefile（推荐）

```bash
# 检查并修复默认目录（internal/gen/specs）
make fix-encoding

# 指定其他目录
make fix-encoding DIR=path/to/directory
```

### 方式2: 直接使用脚本

```bash
# 检查并修复默认目录
.claude/commands/fix-encoding.sh

# 检查并修复指定目录
.claude/commands/fix-encoding.sh /path/to/directory
```

## 输出示例

```
🔍 Scanning directory: internal/gen/specs

📊 Summary:
   Total files found: 233
   Files checked: 233
   Files fixed: 233

🔧 Encoding issues fixed:
   - Fixed encoding in: 游资交易每日明细___dragon_detail.json
   - Fixed encoding in: 东方财富App热榜___em_hot.json
   - Fixed encoding in: 涨停最强板块统计___limit_list_sec.json
   ...

✅ Fixed 233 file(s) with encoding issues
```

## 检测和修复的问题

### 1. 文件名编码问题
- 非 UTF-8 编码的文件名
- 混合编码导致的乱码
- 特殊字符显示异常

### 2. 文件内容编码问题
- UTF-8 替换字符（）
- 双重编码的 UTF-8
- 无效的 UTF-8 序列
- 控制字符
- JSON 格式错误

## 技术细节

### 编码检测
- 使用 `iconv` 检测 UTF-8 编码有效性
- 检查常见的编码混淆模式（Mojibake）
- 验证 JSON 格式正确性

### 编码修复
- 自动检测原始编码（Latin-1, GBK, 等）
- 转换为标准 UTF-8
- 移除无效字符
- 重新格式化 JSON

### 文件处理
- 递归处理指定目录下的所有 JSON 文件
- 保持目录结构不变
- 自动备份（重命名）原文件

## 注意事项

1. **备份重要文件**: 虽然工具会自动处理，但建议先备份重要文件
2. **测试环境**: 先在测试环境运行，确认效果后再应用到生产环境
3. **只处理 JSON 文件**: 当前版本只处理 `.json` 文件
4. **需要 Python3**: 工具依赖 Python3 进行编码检测和修复

## 故障排除

### 问题: 提示 "command not found: python3"

**解决方案**:
```bash
# macOS
brew install python3

# Ubuntu/Debian
sudo apt-get install python3

# CentOS/RHEL
sudo yum install python3
```

### 问题: 提示 "permission denied"

**解决方案**:
```bash
chmod +x .claude/commands/fix-encoding.sh
```

### 问题: 修复后仍有乱码

**解决方案**:
- 检查终端编码设置: `echo $LANG`
- 设置为 UTF-8: `export LANG=en_US.UTF-8`
- 重新运行 fix-encoding 命令

## 定期维护建议

建议在以下情况下运行 fix-encoding：

1. **生成 API 规范后**: 确保新生成的文件编码正确
2. **编辑文件后**: 如果手动编辑了 JSON 文件
3. **从不同系统复制后**: 从 Windows 或其他系统复制文件后
4. **定期检查**: 每周或每月运行一次，确保编码一致

## 相关文件

- **命令脚本**: `.claude/commands/fix-encoding.sh`
- **Makefile 目标**: `make fix-encoding`
- **API 规范目录**: `internal/gen/specs`

## 版本历史

- **v1.0** (2026-03-07): 初始版本
  - 支持文件名编码检查和修复
  - 支持文件内容编码检查和修复
  - 自动处理 JSON 格式错误
