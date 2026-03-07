# .gitignore 文件使用说明

## 快速开始

项目根目录的 `.gitignore` 文件已经配置好，会自动忽略以下类型的文件：

### 自动忽略的内容
- ✅ **构建产物**: `bin/`, `*.exe`, `*.out`, `coverage.out`
- ✅ **测试文件**: `*.test`, `coverage.html`, `*.prof`
- ✅ **临时文件**: `*.tmp`, `*.log`, `*.cache`
- ✅ **IDE配置**: `.vscode/`, `.idea/`, `*.swp`
- ✅ **系统文件**: `.DS_Store`, `Thumbs.db`
- ✅ **敏感信息**: `.env`, `*.key`, `credentials.json`

## 验证配置

### 检查文件是否被忽略
```bash
# 检查特定文件
git check-ignore <filename>

# 创建测试文件验证
touch test.tmp
git status  # test.tmp 不会出现在未跟踪文件中
rm test.tmp
```

### 查看实际状态
```bash
# 只显示未被忽略的未跟踪文件
git status

# 显示所有文件（包括被忽略的）
git status --ignored
```

## 常见操作

### 停止跟踪已提交的文件
```bash
# 停止跟踪但保留本地文件
git rm --cached <file>

# 停止跟踪整个目录
git rm -r --cached <directory>

# 提交更改
git commit -m "更新.gitignore，停止跟踪临时文件"
```

### 强制添加被忽略的文件
```bash
# 强制添加特定文件
git add -f <important-file>

# 或在.gitignore中添加例外
!important-file.env
```

## 项目特定规则

### Tushare API 相关
- `tushare_token.txt` - API令牌（敏感）
- `credentials.json` - 凭证文件（敏感）
- `*.csv`, `*.xlsx` - 数据导出文件（通常很大）

### MCP 服务器相关
- `.mcp/` - MCP 缓存目录
- `*mcp-cache*` - MCP 缓存文件

## 维护建议

1. **敏感信息优先**: 确保 API 密钥、令牌等永远不会被提交
2. **构建产物**: 所有可以重新生成的文件都应该被忽略
3. **IDE配置**: 个人开发环境配置应该被忽略
4. **定期检查**: 使用 `git status --ignored` 查看是否有重要文件被意外忽略

## 故障排除

### 文件仍被跟踪
修改 `.gitignore` 后，已存在的文件仍会被跟踪。需要手动移除：

```bash
# 移除跟踪但保留文件
git rm --cached <file>

# 移除跟踪并删除文件
git rm <file>
```

### 查看详细状态
```bash
# 查看完整的git状态
git status

# 查看未跟踪的文件
git status --short | grep "^??"

# 查看被忽略的文件
git status --ignored --short
```

## 相关资源

- [Git官方文档 - gitignore](https://git-scm.com/docs/gitignore)
- [GitHub - gitignore模板](https://github.com/github/gitignore)
- [Go项目标准.gitignore](https://github.com/golang/go/wiki/GitIgnore)

## 更新日志

- **2025-03-07**: 创建初始 `.gitignore` 配置
  - 添加 Go 项目常见忽略规则
  - 添加 Tushare 和 MCP 特定规则
  - 包含构建产物、测试文件、IDE配置等