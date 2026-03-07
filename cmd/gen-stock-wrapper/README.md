# Stock Wrapper Generator

为股票 API 生成包装器函数，提供统一的链式调用接口。

## 使用方法

### 基本用法

在项目根目录运行：

```bash
go run cmd/gen-stock-wrapper/main.go
```

或者使用编译后的二进制文件：

```bash
./bin/gen-stock-wrapper
```

### 命令行选项

```bash
gen-stock-wrapper [选项]
```

**选项：**

- `-api-dir string`: 股票 API 目录路径（默认: `pkg/sdk/api/stock`）
- `-output string`: 输出包装器文件路径（默认: `pkg/sdk/apis/stock.go`）
- `-help`: 显示帮助信息

### 示例

使用默认路径：

```bash
gen-stock-wrapper
```

使用相对路径：

```bash
gen-stock-wrapper -api-dir ./api/stock -output ./apis/stock.go
```

使用绝对路径：

```bash
gen-stock-wrapper -api-dir /path/to/api/stock -output /path/to/stock.go
```

## 工作原理

1. **扫描 API 文件**: 扫描指定目录下的所有 Go 文件
2. **提取函数签名**: 使用 AST 解析提取导出的函数
3. **生成包装器**: 为每个函数生成包装器，提供统一的调用接口

生成的包装器位于 `pkg/sdk/apis/stock.go`。

## 输出示例

生成的包装器函数类似：

```go
// StockBasicDaily 包装函数
func StockBasicDaily(ctx context.Context, client *sdk.Client, req *stock_basic.StockBasicDailyRequest) ([]stock_basic.StockBasicDailyItem, error) {
	return stock_basic.StockBasicDaily(ctx, client, req)
}
```

## 注意事项

- 生成的代码包含包装函数，不应手动编辑
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
