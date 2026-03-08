# 日志系统使用指南

## 概述

tushare-go 使用基于 [logrus](https://github.com/sirupsen/logrus) 和 [lumberjack](https://github.com/natefinch/lumberjack) 的结构化日志系统，提供强大的日志记录功能。

## 主要特性

- ✅ **结构化日志** - 支持字段形式的日志记录
- ✅ **日志轮转** - 自动按大��和时间轮转日志文件
- ✅ **多种输出** - 控制台、文件、同时输出
- ✅ **多种格式** - JSON 和 Text 两种格式
- ✅ **日志级别** - Debug、Info、Warn、Error、Fatal
- ✅ **自动错误记录** - API 调用失败时自动记录详细信息

## 快速开始

### 基础配置（输出到文件）

```go
import (
    "tushare-go/pkg/sdk"
    "tushare-go/pkg/sdk/logger"
    futures "tushare-go/pkg/sdk/api/futures"
)

func main() {
    // 初始化日志系统
    logger.Init(&logger.LogConfig{
        Filename:   "api.log",       // 日志文件路径
        MaxSize:    10,              // 单个文件最大大小（MB）
        MaxAge:     30,              // 保留天数
        MaxBackups: 3,               // 保留备份文件数量
        Compress:   true,            // 压缩旧日志文件
        Level:      "info",          // 日志级别
        Format:     "text",          // 日志格式
    })

    // 创建客户端
    config, _ := sdk.NewConfig("your-token")
    client := sdk.NewClient(config)

    // 使用 API
    calendar, err := futures.TradeCal(ctx, client, &futures.TradeCalRequest{
        Exchange: "SHFE",
        StartDate: "20250101",
        EndDate:   "20250110",
    })

    if err != nil {
        logger.WithError(err).Error("API调用失败")
    } else {
        logger.WithField("count", len(calendar)).Info("获取成功")
    }
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
    // 创建日志文件
    logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    defer logFile.Close()

    // 同时输出到控制台和文件
    multiWriter := io.MultiWriter(os.Stdout, logFile)
    logger.SetOutput(multiWriter)

    // 初始化日志系统
    logger.Init(&logger.LogConfig{
        Level:  "debug",
        Format: "text",
    })

    // 您的代码...
}
```

### 只输出到控制台（开发调试）

```go
logger.Init(&logger.LogConfig{
    Filename: "", // 留空表示输出到控制台
    Level:    "debug",
    Format:   "text",
})
```

## 日志配置选项

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `Filename` | string | `""` | 日志文件路径，留空则输出到控制台 |
| `MaxSize` | int | `100` | 单个日志文件最大大小（MB） |
| `MaxAge` | int | `30` | 日志文件保留天数 |
| `MaxBackups` | int | `3` | 保留的旧日志文件数量 |
| `Compress` | bool | `true` | 是否压缩旧日志文件（.gz） |
| `Level` | string | `"info"` | 日志级别：`debug`/`info`/`warn`/`error` |
| `Format` | string | `"text"` | 日志格式：`json`/`text` |

## 日志级别

```go
logger.Debug("调试信息")    // 开发调试使用
logger.Info("一般信息")     // 常规信息
logger.Warn("警告信息")     // 警告但不影响运行
logger.Error("错误信息")    // 错误需要关注
logger.Fatal("致命错误")    // 致命错误后退出程序
```

## 结构化日志

### 带字段的日志

```go
// 单个字段
logger.WithField("api_name", "trade_cal").Info("调用API")

// 多个字段
logger.WithFields(logger.Fields{
    "api_name": "trade_cal",
    "exchange": "SHFE",
    "count":    10,
}).Info("获取数据成功")

// 带错误信息
logger.WithError(err).Error("API调用失败")
```

### 格式化日志

```go
logger.Infof("处理了 %d 条数据", count)
logger.Errorf("处理失败: %v", err)
```

## API 自动错误日志

SDK 会在以下情况自动记录详细的错误日志：

### 1. API 调用失败

```
time="2026-03-09 01:00:23" level=error msg="=== API调用失败 ===" \
api_name=trade_cal \
error="API错误: code=40101, msg=您的token不对，请确认。" \
request="{\"api_name\":\"trade_cal\",...}" \
response="{\"code\":40101,...}"
```

### 2. 字段类型转换失败

```
time="2026-03-09 00:51:04" level=error \
msg="=== 字段解析失败 ===" \
API="trade_cal" \
字段="is_open" \
错误="类型转换失败，期望类型 string，支持 string/float64/int" \
字段原始值="null" \
字段实际类型="<nil>" \
当前Item="{\"cal_date\":\"20240101\",...}"
```

### 3. 网络错误

```
time="2026-03-09 00:53:23" level=error \
msg="=== API调用失败 ===" \
API名称="trade_cal" \
错误信息="HTTP请求失败: Post \"...\": context deadline exceeded" \
请求_body="{...}" \
响应_body=""
```

## 日志格式

### Text 格式（默认）

```
time="2026-03-09 01:00:23" level=info msg="获取数据成功" count=10 api_name=trade_cal
```

### JSON 格式

```json
{
  "time": "2026-03-09T01:00:23+08:00",
  "level": "info",
  "msg": "获取数据成功",
  "count": 10,
  "api_name": "trade_cal"
}
```

## 日志轮转

日志文件会自动轮转，避免单个文件过大：

```
api.log           # 当前日志文件
api.log.1         # 第1个备份
api.log.2.gz      # 第2个备份（已压缩）
api.log.3.gz      # 第3个备份（已压缩）
```

轮转规则：
- 当文件大小超过 `MaxSize` 时自动轮转
- 保留 `MaxBackups` 个备份文件
- 超过 `MaxAge` 天的备份文件自动删除
- 可选择压缩旧文件节省空间

## 最佳实践

### 1. 生产环境配置

```go
logger.Init(&logger.LogConfig{
    Filename:   "/var/log/tushare-api/api.log",
    MaxSize:    100,               // 100MB
    MaxAge:     90,                // 保留90天
    MaxBackups: 10,                // 保留10个备份
    Compress:   true,
    Level:      "info",            // 生产环境使用 info 级别
    Format:     "json",            // JSON格式便于日志分析
})
```

### 2. 开发环境配置

```go
logger.Init(&logger.LogConfig{
    Filename:   "",                // 输出到控制台
    Level:      "debug",           // 显示所有级别
    Format:     "text",            // Text格式便于阅读
})
```

### 3. 测试环境配置

```go
logger.Init(&logger.LogConfig{
    Filename:   "test.log",
    MaxSize:    10,
    MaxAge:     7,                 // 只保留7天
    MaxBackups: 3,
    Compress:   false,             // 不压缩便于查看
    Level:      "debug",
    Format:     "text",
})
```

## 调试技巧

### 查看完整 API 请求和响应

当 API 调用失败时，日志会自动包含完整的请求和响应信息：

```bash
# 实时查看日志
tail -f api.log

# 搜索错误日志
grep "level=error" api.log

# 搜索特定 API 的日志
grep "api_name=trade_cal" api.log

# 查看最近的错误
grep "level=error" api.log | tail -20
```

### 结构化日志查询

如果使用 JSON 格式，可以使用 `jq` 进行查询：

```bash
# 查看所有错误日志
jq 'select(.level=="error")' api.log

# 查看特定 API 的日志
jq 'select(.api_name=="trade_cal")' api.log

# 统计错误数量
jq 'select(.level=="error")' api.log | wc -l
```

## 常见问题

### Q: 如何禁用日志？

A: 设置 Level 为 `"fatal"` 或将 Filename 设为空并重定向输出：

```go
logger.Init(&logger.LogConfig{
    Filename: "/dev/null",  // 或 ""
    Level:    "fatal",
})
```

### Q: 如何自定义日志格式？

A: 可以在初始化后设置 formatter：

```go
import logrus "github.com/sirupsen/logrus"

logger.L().SetFormatter(&logrus.TextFormatter{
    FullTimestamp:   true,
    TimestampFormat: "2006-01-02 15:04:05",
    ForceColors:     true,
})
```

### Q: 日志文件太大怎么办？

A: 调整轮转配置：

```go
logger.Init(&logger.LogConfig{
    MaxSize:    10,     // 减小单个文件大小
    MaxAge:     7,      // 减少保留天数
    MaxBackups: 3,      // 减少备份数量
    Compress:   true,   // 启用压缩
})
```

## 相关链接

- [logrus 文档](https://github.com/sirupsen/logrus)
- [lumberjack 文档](https://github.com/natefinch/lumberjack)
- [README.md](README.md) - 项目概述
- [ERROR_HANDLING.md](#) - 错误处理指南
