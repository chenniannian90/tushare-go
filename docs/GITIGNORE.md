# .gitignore 配置说明

## 概述

本项目使用 `.gitignore` 文件来指定不需要纳入版本控制的文件和目录。这样可以保持仓库清洁，避免提交不必要的文件。

## 主要忽略类别

### 1. 构建产物
```
/bin/              # 二进制可执行文件
*.exe             # Windows可执行文件
*.test            # Go测试二进制文件
*.out             # 各种输出文件
coverage.out      # 测试覆盖率文件
coverage.html     # HTML格式的覆盖率报告
```

### 2. 临时和缓存文件
```
*.tmp             # 临时文件
*.log             # 日志文件
*.cache           # 缓存文件
tmp/              # ���时目录
temp/             # 临时目录
```

### 3. IDE和编辑器配置
```
.vscode/          # VSCode配置
.idea/            # JetBrains IDE配置
*.swp             # Vim交换文件
*.swo             # Vim交换文件
.DS_Store         # macOS系统文件
```

### 4. 敏感信息
```
.env              # 环境变量
*.key             # 私钥文件
*.pem             # PEM证书文件
credentials.json  # 凭证文件
tushare_token.txt # Tushare API令牌
```

### 5. Go特定文件
```
go.work           # Go工作区文件
go.work.sum       # Go工作区和校验
vendor/           # 依赖目录
```

## 使用示例

### 验证文件是否被忽略
```bash
# 检查特定文件是否被忽略
git check-ignore <filename>

# 运行验证脚本
bash verify_gitignore.sh
```

### 查看实际状态
```bash
# 查看git状态（排除已忽略的文件）
git status

# 查看所有文件（包括被忽略的）
git status --ignored
```

### 添加例外情况
如果需要在被忽略的目录中添加特定文件到版本控制：

```bash
# 强制添加文件
git add -f <filename>

# 或在.gitignore中使用否定模式
!.env.example
```

## 项目特定配置

### Tushare相关
- `tushare_token.txt` - API令牌文件
- `credentials.json` - 凭证文件
- `*.csv` - 数据导出文件（通常很大）

### MCP服务器相关
- `.mcp/` - MCP缓存目录
- `*mcp-cache*` - MCP缓存文件

## 维护建议

1. **定期检查**：运行 `git status --ignored` 查看是否有重要文件被意外忽略
2. **敏感信息**：永远不要提交API密钥、密码等敏感信息
3. **构建产物**：所有可重新生成的文件都应该被忽略
4. **IDE文件**：个人开发环境配置应该被忽略，团队配置可以提交

## 故障排除

### 文件仍然被跟踪
如果文件已经被git跟踪，修改.gitignore不会自动停止跟踪它：

```bash
# 停止跟踪文件（但保留本地文件）
git rm --cached <filename>

# 停止跟踪目录
git rm -r --cached <directory>
```

### 查看为什么文件被忽略
```bash
# 查看具体的忽略规则
git check-ignore -v <filename>
```

## 相关命令

```bash
# 清理未跟踪的文件
git clean -fd

# 清理所有未跟踪的文件（包括被忽略的）
git clean -fdX

# 查看git状态
git status

# 添加所有更改
git add .

# 提交更改
git commit -m "更新.gitignore配置"
```
