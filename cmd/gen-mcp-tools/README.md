# MCP Tools Generator

自动从 API 定义生成 MCP 工具包装器。

## 使用方法

### 基本用法

在项目根目录运行：

```bash
go run cmd/gen-mcp-tools/main.go
```

或者使用编译后的二进制文件：

```bash
./bin/gen-mcp-tools
```

### 命令行选项

```bash
gen-mcp-tools [选项]
```

**选项：**

- `-api-path string`: API 目录路径（默认: `pkg/sdk/api`）
- `-mcp-tools-path string`: MCP 工具输出目录路径（默认: `pkg/mcp/tools`）
- `-help`: 显示帮助信息

### 示例

使用默认路径：

```bash
gen-mcp-tools
```

使用相对路径：

```bash
gen-mcp-tools -api-path ./api -mcp-tools-path ./tools
```

使用绝对路径：

```bash
gen-mcp-tools -api-path /path/to/api -mcp-tools-path /path/to/tools
```

## 工作原理

1. **扫描 API 模块**: 扫描 `pkg/sdk/api` 目录下的所有 API 模块
2. **生成工具包装器**: 为每个 API 函数生成 MCP 工具包装器
3. **创建注册表**: 自动生成工具注册表代码

生成的代码位于 `pkg/mcp/tools/` 目录，每个 API 模块对应一个子目录。

## 注意事项

- 生成的代码包含 `DO NOT EDIT` 注释，不应手动编辑
- 每次运行生成器会覆盖之前的生成文件
- 确保在项目根目录运行，或者使用正确的相对/绝对路径

## 故障排除

### 找不到 go.mod

如果看到 "go.mod not found" 错误：

1. 确保在项目根目录运行
2. 或者使用绝对路径指定 API 和输出目录

### 路径问题

如果遇到路径问题：

1. 使用 `-help` 查看使用说明
2. 使用绝对路径避免相对路径问题
3. 确保指定的目录存在
