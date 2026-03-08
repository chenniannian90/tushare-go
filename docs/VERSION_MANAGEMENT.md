# 版本管理指南

## 概述

Tushare MCP 服务器现在支持自动版本管理，版本信息会自动从 Git tag 中获取并注入到二进制文件中。

## 版本信息获取

### 自动版本获取

项目使用 Git tag 作为版本号来源：

```bash
# 获取最新 tag
git describe --tags --abbrev=0
```

如果没有 tag，则使用默认版本 `v0.0.0-dev`。

### 版本信息组成

版本信息包含三个部分：

- **Version**: 从 Git tag 获取的版本号
- **GitCommit**: Git commit hash (短格式)
- **BuildDate**: 构建时间戳

## 使用方法

### 1. 查看版本信息

```bash
# 构建后查看版本
./bin/tushare-mcp --version

# 输出示例:
# Tushare MCP Server
# Version: v1.0.0 (ca3c630)
# Full Info: v1.0.0, commit: ca3c630, built at: 2026-03-08_09:00:00
```

### 2. Makefile 版本管理

```bash
# 查看当前版本信息
make version

# 使用当前 git tag 构建
make build-mcp

# 指定版本构建
VERSION=v1.2.3 make build-mcp
```

### 3. 发布新版本

```bash
# 1. 创建版本 tag
git tag v1.0.0

# 2. 推送 tag 到远程
git push origin v1.0.0

# 3. 构建带有版本信息的二进制文件
make build-mcp

# 4. 验证版本信息
./bin/tushare-mcp --version
```

## 版本号规范

### 语义化版本

遵循 [Semantic Versioning 2.0.0](https://semver.org/) 规范：

```
MAJOR.MINOR.PATCH

示例:
- v1.0.0 - 第一个稳定版本
- v1.1.0 - 添加新功能，向后兼容
- v1.1.1 - Bug 修复
- v2.0.0 - 破坏性变更
```

### Tag 命名规则

```bash
# 正确的 tag 格式
v1.0.0
v2.1.3

# 错误的 tag 格式
1.0.0       # 缺少 v 前缀
version-1.0.0 # 不符合规范
```

## 构建流程

### 标准构建

```bash
# 1. 确保工作目录干净
git status

# 2. 创建版本 tag (如果还没有)
git tag v1.0.0

# 3. 构建
make build-mcp

# 4. 验证版本
./bin/tushare-mcp --version
```

### 自定义版本构建

```bash
# 使用环境变量指定版本
VERSION=v2.0.0-beta make build-mcp

# 或者同时指定所有版本信息
VERSION=v2.0.0-beta \
GIT_COMMIT=custom-commit \
BUILD_DATE=2024-01-01 \
make build-mcp
```

## 版本信息在代码中的使用

### 获取版本信息

```go
package main

import "fmt"

func main() {
    // 获取简短版本
    version := GetVersion()
    fmt.Println(version) // v1.0.0 (ca3c630)

    // 获取完整版本信息
    fullInfo := GetFullVersionInfo()
    fmt.Println(fullInfo) // v1.0.0, commit: ca3c630, built at: 2024-01-01_12:00:00
}
```

### MCP 服务器版本

MCP 服务器的版本信息会自动在握手时发送给客户端：

```go
impl := &mcpsdk.Implementation{
    Name:    "tushare-mcp-stock",
    Version: Version, // 自动注入的版本信息
}
```

## Makefile 变量说明

```makefile
# 版本相关变量
VERSION       # Git tag 或指定的版本号
GIT_COMMIT    # Git commit hash (短格式)
BUILD_DATE    # 构建时间戳
LDFLAGS       # Go 编译器 flags，用于版本注入
```

## CI/CD 集成

### GitHub Actions 示例

```yaml
name: Build and Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
        with:
          fetch-depth: 0  # 获取所有历史用于 git describe

      - name: Build
        run: |
          VERSION=${{ github.ref_name }} make build-mcp

      - name: Show Version
        run: ./bin/tushare-mcp --version

      - name: Upload Binary
        uses: actions/upload-artifact@v3
        with:
          name: tushare-mcp
          path: bin/tushare-mcp
```

## 版本管理最佳实践

### 1. 发布前检查清单

- [ ] 所有测试通过
- [ ] 更新 CHANGELOG
- [ ] 创建正确的版本 tag
- [ ] 构建并验证版本信息
- [ ] 推送 tag 到远程仓库

### 2. 开发版本

```bash
# 开发期间不需要创建 tag
# 版本会自动显示为 v0.0.0-dev
make build-mcp
```

### 3. 预发布版本

```bash
# 使用预发布标识符
git tag v1.0.0-beta.1
git tag v1.0.0-rc.1
```

### 4. 热修复版本

```bash
# 从发布分支创建 hotfix
git checkout -b hotfix/v1.0.1
# ... 修复问题 ...
git tag v1.0.1
```

## 常见问题

### Q: 为什么显示 v0.0.0-dev?

A: 这表示没有找到 Git tag。解决方案：

```bash
# 创建一个 tag
git tag v1.0.0

# 或者指定版本
VERSION=v1.0.0 make build-mcp
```

### Q: 如何查看所有版本 tags?

```bash
git tag
```

### Q: 如何删除错误的 tag?

```bash
# 删除本地 tag
git tag -d v1.0.0

# 删除远程 tag
git push origin :refs/tags/v1.0.0
```

### Q: 版本信息如何在运行时获取?

```go
// 在代码中使用
version := GetVersion()
fmt.Printf("Server version: %s\n", version)
```

## 相关文件

- `cmd/mcp-server/version.go` - 版本变量定义
- `Makefile` - 构建和版本注入逻辑
- `cmd/mcp-server/server.go` - MCP 服务器版本使用
- `cmd/mcp-server/main.go` --version 命令实现

## 更多信息

- [Semantic Versioning](https://semver.org/)
- [Git Tag 文档](https://git-scm.com/book/en/v2/Git-Basics-Tagging)
- [Go Linker Flags](https://pkg.go.dev/cmd/link#hdr-Linker_Options)
