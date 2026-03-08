# 版本发布指南

## 快速发布流程

### 1. 准备发布

```bash
# 确保工作目录干净
git status

# 确保在主分支
git branch

# 拉取最新代码
git pull origin main
```

### 2. 运行测试

```bash
# 运行所有测试
make test

# 检查构建
make build-mcp
```

### 3. 创建版本 Tag

```bash
# 确定版本号（遵循语义化版本）
# 格式: vMAJOR.MINOR.PATCH

# 创建 tag
git tag v1.0.0

# 查看 tag
git tag -l "v1.0.0"

# 推送 tag 到远程
git push origin v1.0.0
```

### 4. 构建发布版本

```bash
# 使用 tag 版本构建
make build-mcp

# 验证版本信息
./bin/tushare-mcp --version

# 预期输出:
# Tushare MCP Server
# Version: v1.0.0 (ca3c630)
# Full Info: v1.0.0, commit: ca3c630, built at: 2026-03-08_09:00:00
```

### 5. 验证发布

```bash
# 检查二进制文件
ls -lh bin/tushare-mcp

# 运行基本功能测试
./bin/tushare-mcp --help
```

## 版本号规范

### 语义化版本 (Semantic Versioning)

```
MAJOR.MINOR.PATCH

示例:
v1.0.0  - 第一个稳定版本
v1.1.0  - 新增功能，向后兼容
v1.1.1  - Bug 修复
v2.0.0  - 破坏性变更
```

### 版本类型说明

| 类型 | 说明 | 示例 |
|------|------|------|
| **MAJOR** | 破坏性变更，不兼容旧版本 | v1.0.0 → v2.0.0 |
| **MINOR** | 新增功能，向后兼容 | v1.0.0 → v1.1.0 |
| **PATCH** | Bug 修复，不影响 API | v1.0.0 → v1.0.1 |

### 预发布版本

```bash
# Alpha 版本（内部测试）
git tag v1.0.0-alpha.1

# Beta 版本（公开测试）
git tag v1.0.0-beta.1

# Release Candidate（候选版本）
git tag v1.0.0-rc.1
```

## 完整发布示例

### 发布 v1.0.0

```bash
# 1. 准备工作
git checkout main
git pull origin main
make test

# 2. 创建 tag
git tag -a v1.0.0 -m "Release v1.0.0: Initial stable release"

# 3. 推送到远程
git push origin main
git push origin v1.0.0

# 4. 构建发布版本
make clean
make build-mcp

# 5. 验证版本
./bin/tushare-mcp --version

# 6. 创建 GitHub Release（在网页上操作）
# 访问: https://github.com/yourusername/tushare-go/releases/new
# 选择 tag: v1.0.0
# 上传二进制文件: bin/tushare-mcp
```

### 热修复版本 v1.0.1

```bash
# 1. 创建 hotfix 分支
git checkout -b hotfix/v1.0.1

# 2. 修复问题...
# ... 编写代码 ...
# ... 测试修复 ...

# 3. 提交修复
git add .
git commit -m "Fix: Critical bug in API handler"

# 4. 创建修复 tag
git tag v1.0.1

# 5. 合并回主分支
git checkout main
git merge hotfix/v1.0.1

# 6. 推送
git push origin main
git push origin v1.0.1

# 7. 构建修复版本
make build-mcp
```

## 多平台构建

### 为不同平台构建

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 make build-mcp
mv bin/tushare-mcp bin/tushare-mcp-linux-amd64

# macOS AMD64
GOOS=darwin GOARCH=amd64 make build-mcp
mv bin/tushare-mcp bin/tushare-mcp-darwin-amd64

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 make build-mcp
mv bin/tushare-mcp bin/tushare-mcp-darwin-arm64

# Windows AMD64
GOOS=windows GOARCH=amd64 make build-mcp
mv bin/tushare-mcp.exe bin/tushare-mcp-windows-amd64.exe
```

### 构建脚本

创建 `scripts/build-release.sh`:

```bash
#!/bin/bash
VERSION=$(git describe --tags --abbrev=0)
echo "Building version: $VERSION"

# Clean
make clean

# Build for multiple platforms
platforms=("linux/amd64" "darwin/amd64" "darwin/arm64" "windows/amd64")

for platform in "${platforms[@]}"; do
    GOOS="${platform%/*}"
    GOARCH="${platform#*/}"

    echo "Building for $GOOS/$GOARCH..."

    if [ "$GOOS" = "windows" ]; then
        OUTPUT="bin/tushare-mcp-$VERSION-$GOOS-$GOARCH.exe"
    else
        OUTPUT="bin/tushare-mcp-$VERSION-$GOOS-$GOARCH"
    fi

    GOOS=$GOOS GOARCH=$GOARCH make build-mcp
    mv "bin/tushare-mcp" "$OUTPUT" || mv "bin/tushare-mcp.exe" "$OUTPUT"
done

echo "Build complete!"
ls -lh bin/
```

## 发布检查清单

### 发布前

- [ ] 所有测试通过 (`make test`)
- [ ] 代码审查完成
- [ ] 文档更新完成
- [ ] CHANGELOG 更新
- [ ] 版本号确定
- [ ] 发布说明准备好

### 发布中

- [ ] 创建 Git tag
- [ ] 推送 tag 到远程
- [ ] 构建二进制文件
- [ ] 验证版本信息
- [ ] 测试基本功能

### 发布后

- [ ] 创建 GitHub Release
- [ ] 上传二进制文件
- [ ] 发布公告
- [ ] 更新文档
- [ ] 监控问题反馈

## 回滚流程

如果发布后发现严重问题：

```bash
# 1. 删除有问题的 tag
git tag -d v1.0.0
git push origin :refs/tags/v1.0.0

# 2. 创建修复版本
git tag v1.0.1
git push origin v1.0.1

# 3. 通知用户使用新版本
```

## 版本查询

### 查看所有版本

```bash
# 本地 tags
git tag

# 远程 tags
git ls-remote --tags origin
```

### 查看版本详情

```bash
# 查看 tag 信息
git show v1.0.0

# 查看 tag 日期
git log -1 --format=%ci v1.0.0
```

### 比较版本差异

```bash
# 比较两个版本
git diff v1.0.0 v1.1.0

# 查看版本间的提交
git log v1.0.0..v1.1.0 --oneline
```

## 故障排除

### 版本显示不正确

```bash
# 清理并重新构建
make clean
make build-mcp

# 确认 tag 存在
git tag -l "v*"

# 手动指定版本
VERSION=v1.0.0 make build-mcp
```

### Git 描述失败

```bash
# 确保有足够的 git 历史
git fetch --unshallow

# 检查是否有 tag
git tag

# 创建初始 tag
git tag v0.1.0
```

## 相关文档

- [版本管理指南](VERSION_MANAGEMENT.md)
- [语义化版本规范](https://semver.org/)
- [Git Tag 文档](https://git-scm.com/book/en/v2/Git-Basics-Tagging)
