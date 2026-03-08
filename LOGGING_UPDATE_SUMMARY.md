# 日志系统和错误处理更新总结

## 更新日期
2026-03-09

## 概述

本次更新为 tushare-go 项��添加了完整的结构化日志系统和详细的错误记录功能，参考了 ci-mcp 项目的日志实现。

## 主要改进

### 1. 结构化日志系统

**新增文件：**
- `pkg/sdk/logger/logger.go` - 基于 logrus 和 lumberjack 的日志系统

**依赖添加：**
- `github.com/sirupsen/logrus v1.9.4` - 结构化日志库
- `gopkg.in/natefinch/lumberjack.v2 v2.2.1` - 日志轮转库

**功能特性：**
- ✅ 结构化日志（字段形式）
- ✅ 日志轮转（按大小和时间）
- ✅ 多种输出（控制台、文件、同时输出）
- ✅ 多种格式（JSON、Text）
- ✅ 日志级别（Debug、Info、Warn、Error、Fatal）
- ✅ 便捷的包装函数

**配置选项：**
```go
type LogConfig struct {
    Filename   string // 日志文件路径，空表示控制台
    MaxSize    int    // 单个文件最大大小（MB）
    MaxAge     int    // 保留天数
    MaxBackups int    // 保留备份数量
    Compress   bool   // 压缩旧日志
    Level      string // 日志级别
    Format     string // 格式：json/text
}
```

### 2. 自动错误记录

**更新文件：**
- `pkg/sdk/client.go` - 更新日志记录逻辑

**功能特性：**
- ✅ API 调用失败时自动记录详细信息
- ✅ 记录完整的请求体（包括参数和 token）
- ✅ 记录完整的响应体（包括错误信息）
- ✅ 字段类型转换失败时的详细日志
- ✅ 网络错误时的详细信息

**日志内容示例：**
```
=== API调用失败 ===
API名称: trade_cal
错误信息: API错误: code=40101, msg=您的token不对，请确认。
请求体: {"api_name":"trade_cal","fields":"exchange,cal_date,is_open,pretrade_date","params":{...},"token":"***"}
响应体: {"code":40101,"data":null,"msg":"您的token不对，请确认。"}
===================
```

### 3. API 类型处理改进

**更新文件：**
- `internal/gen/templates/api.go.tmpl` - 更新代码生成模板
- `internal/gen/generator.go` - 添加 hasStringFields 函数

**改进内容：**
- ✅ string 类型字段支持自动类型转换（string/float64/int）
- ✅ 使用 `CallAPIFlexible` 替代 `CallAPI`
- ✅ 智能导入优化（只在需要时导入 log 和 json）

**转换逻辑：**
```go
var isOpen string
if v, ok := item["is_open"].(string); ok {
    isOpen = v
} else if v, ok := item["is_open"].(float64); ok {
    isOpen = fmt.Sprintf("%.0f", v)
} else if v, ok := item["is_open"].(int); ok {
    isOpen = fmt.Sprintf("%d", v)
} else {
    // 记录详细的错误日志
    logger.WithFields(...).Error("类型转换失败")
}
```

### 4. 文档更新

**新增文档：**
- `LOGGING.md` - 完整的日志系统使用指南

**更新文档：**
- `README.md` - 添加日志系统说明和配置章节

**新增示例：**
- `cmd/examples/logging_usage/main.go` - 日志使用示例

## 问题修复

### 1. 类型不一致问题

**修复内容：**
- 期货交易日历 `is_open`: `int` → `str`
- 港股交易日历 `is_open`: `int` → `str`
- 美股交易日历 `is_open`: `int` → `str`

**修复文件：**
- `internal/gen/specs/期货数据___futures/交易日历___trade_cal.json`
- `internal/gen/specs/港股数据___hk_stock/港股交易日历___hk_tradecal.json`
- `internal/gen/specs/美股数据___us_stock/美股交易日历___us_tradecal.json`

### 2. 代码生成器模板 Bug

**问题：** 数组类型转换时硬编码了 `case []string` 和 `case []float64`，导致类型不匹配

**修复：** 使用动态类型 `case []{{arrayElemType $f}}`

### 3. Lint 错误

**问题：** 模板生成的代码中存在未使用的导入

**修复：** 添加 `hasStringFields` 函数，只在有 string 字段时导入必要包

### 4. 响应解析器 Exhaustive Switch

**问题：** switch 语句缺少 `FormatUnknown` case

**修复：** 添加所有枚举值的 case 和 panic("unreachable")

## 代码生成

**重新生成文件：**
- 所有 234 个 API 文件
- 所有 MCP 工具

**生成命令：**
```bash
go run cmd/regenerate-api/main.go
go run cmd/gen-mcp-tools/main.go
```

## 测试

**测试场景：**
1. ✅ 正常 API 调用（验证数据正确转换）
2. ✅ 无效 token（验证错误日志）
3. ✅ 网络超时（验证网络错误日志）
4. ✅ 类型转换（验证 int/float64 → string 自动转换）

**测试结果：**
- ✅ 所有 lint 检查通过
- ✅ 所有测试通过
- ✅ 日志功能正常工作
- ✅ 错误日志详细信息完整

## 使用示例

### 基础配置

```go
import "tushare-go/pkg/sdk/logger"

func main() {
    logger.Init(&logger.LogConfig{
        Filename:   "api.log",
        MaxSize:    10,
        MaxAge:     30,
        MaxBackups: 3,
        Compress:   true,
        Level:      "info",
        Format:     "text",
    })
}
```

### 高级配置（同时输出到控制台和文件）

```go
import (
    "io"
    "os"
    "tushare-go/pkg/sdk/logger"
)

func main() {
    logFile, _ := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    multiWriter := io.MultiWriter(os.Stdout, logFile)
    logger.SetOutput(multiWriter)

    logger.Init(&logger.LogConfig{
        Level:  "debug",
        Format: "text",
    })
}
```

### 使用日志

```go
// 简单日志
logger.Info("API调用成功")
logger.Error("API调用失败")

// 结构化日志
logger.WithField("api_name", "trade_cal").Info("调用API")

// 多字段
logger.WithFields(logger.Fields{
    "api_name": "trade_cal",
    "count":    10,
}).Info("获取数据成功")

// 带错误
logger.WithError(err).Error("处理失败")
```

## 提交记录

1. `f4f9e9d` - fix(gen): 修复类型不一致和代码生成器模板bug
2. `d2d99f4` - fix(gen): 改进API类型处理逻辑以兼容实际返回数据
3. `6c43105` - feat(sdk): 添加详细的API错误日志记录
4. `986ce04` - feat(sdk): 添加结构化日志系统

## 后续改进建议

1. **添加日志过滤** - 支持按 API 名称、用户 ID 等过滤日志
2. **性能监控** - 添加 API 调用耗时统计
3. **告警集成** - 集成告警系统，特定错误触发告警
4. **日志分析** - 添加日志分析工具和仪表板
5. **分布式追踪** - 集成 OpenTelemetry 进行分布式追踪

## 相关文档

- [LOGGING.md](LOGGING.md) - 日志系统完整指南
- [README.md](README.md) - 项目概述
- [GENERATION.md](GENERATION.md) - 代码生成说明
- [MCP_SERVICES.md](MCP_SERVICES.md) - MCP 服务说明

## 验证步骤

1. 运行示例代码：
```bash
export TUSHARE_TOKEN=your_token
go run cmd/examples/logging_usage/main.go
```

2. 检查日志输出：
```bash
# 查看日志文件
cat api.log

# 实时查看日志
tail -f api.log

# 搜索错误日志
grep "level=error" api.log
```

3. 验证代码生成：
```bash
go run cmd/regenerate-api/main.go
go run cmd/gen-mcp-tools/main.go
make lint
```

## 总结

本次更新显著提升了 tushare-go 项目的可维护性和可调试性：

- ✅ **开发体验提升** - 详细的错误日志让问题定位更快速
- ✅ **生产就绪** - 完善的日志轮转和压缩
- ✅ **灵活配置** - 支持多种输出方式和格式
- ✅ **结构化数据** - 便于日志分析和查询
- ✅ **自动化** - 无需手动记录日志，自动记录所有错误

所有更改已提交到 git 并推送到远程仓库。
