# MCP 工具和 HTTP 路由文档

## 概述

本项目提供了自动生成的 MCP 工具和 HTTP 路由，覆盖了所有 Tushare API 模块。

## 架构

### 目录结构

```
pkg/
├── sdk/
│   └── api/              # SDK API 实现
│       ├── bond/         # 债券 API
│       ├── etf/          # ETF API
│       ├── fund/         # 基金 API
│       ├── hk_stock/     # 港股 API
│       ├── index/        # 指数 API
│       ├── stock/        # 股票 API
│       └── us_stock/     # 美股 API
└── mcp/
    ├── tools/            # 生成的 MCP 工具
    │   ├── registry.go   # 工具注册表
    │   ├── bond_tools.go
    │   ├── etf_tools.go
    │   └── ...
    └── http_routes.go    # HTTP 路由映射
```

## MCP 工具

### 工具命名规范

MCP 工具使用以下命名格式：`<module>.<tool_name>`

例如：
- `bond.bond_oc` - 债券柜台交易
- `fund.fund_basic` - 基金基本信息
- `stock.daily` - 股票日线数据

### 可用工具列表

#### 债券工具 (bond)
- `bond.api272` - 大宗交易明细
- `bond.api392` - 债券相关数据
- `bond.block_trade` - 柜台债券交易
- `bond.bond_oc` - 柜台流通式债券报价
- `bond.bond_repurchase` - 债券回购日行情
- `bond.bond_zs` - 债券综述
- `bond.cb_basic` - 可转债基本信息
- `bond.cb_call` - 可转债赎回
- `bond.cb_daily` - 可转债日线
- `bond.cb_interest` - 可转债票面利率
- `bond.cb_issue` - 可转债发行
- `bond.cb_redemption` - 可转债回售
- `bond.global_calendar` - 全球财经事件

#### ETF 工具 (etf)
- `etf.api127` - ETF日线行情
- `etf.api199` - ETF复权因子
- `etf.api385` - ETF基本信息
- `etf.api387` - ETF历史分钟
- `etf.api400` - ETF实时日线
- `etf.api408` - ETF份额规模
- `etf.api416` - ETF实时分钟

#### 基金工具 (fund)
- `fund.fund_basic` - 基金基本信息
- `fund.fund_nav` - 基金净值
- `fund.fund_div` - 基金分红
- `fund.fund_manager` - 基金经理
- `fund.api359` - 基金技术面因子

#### 港股工具 (hk_stock)
- `hk_stock.hk_basic` - 港股基本信息
- `hk_stock.hk_daily` - 港股日线行情
- `hk_stock.hk_cal` - 港股交易日历
- `hk_stock.hk_min` - 港股分钟数据
- `hk_stock.hk_factor` - 港股因子数据

#### 指数工具 (index)
- `index.index_basic` - 指数基本信息
- `index.index_daily` - 指数日线行情
- `index.index_member` - 指数成分股
- `index.index_weight` - 指数权重
- `index.api358` - 指数数据
- `index.index_weekly` - 指数周线

#### 美股工具 (us_stock)
- `us_stock.us_basic` - 美股基本信息
- `us_stock.us_daily` - 美股日线行情
- `us_stock.us_cal` - 美股交易日历
- `us_stock.us_factor` - 美股因子数据

## HTTP 路由

### 路由格式

所有 HTTP API 端点遵循以下格式：

```
POST http://localhost:7878/{service_path}
```

例如：
- `POST http://localhost:7878/bond` - 债券数据服务
- `POST http://localhost:7878/fund` - 基金数据服务
- `POST http://localhost:7878/stock/stock_basic` - 股票基础信息

### 路由列表

每个 MCP 服务都有对应的 HTTP 端点。查看服务列表了解所有可用的路由。

### MCP 协议请求格式

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "tool_name",
    "arguments": {
      "ts_code": "000001.SZ",
    "start_date": "20240101",
    "end_date": "20240131"
  }
}
```

### HTTP 响应格式

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "fields": ["ts_code", "trade_date", "open", "close", "high", "low", "vol"],
    "items": [
      ["000001.SZ", "20240101", 10.5, 10.8, 10.9, 10.4, 1234567.0]
    ]
  }
}
```

## 代码生成

### 生成器位置

```
cmd/gen-mcp-tools/main.go
```

### 重新生成工具

当 SDK API 有更新时，可以重新生成 MCP 工具：

```bash
go run cmd/gen-mcp-tools/main.go
```

### 生成器功能

1. **扫描 API 目录**：自动发现所有 API 模块
2. **解析函数定义**：从 API 文件中提取函数信息
3. **生成工具代码**：为每个 API 生成对应的 MCP 工具
4. **创建路由映射**：自动生成 HTTP 路由配置

## 使用示例

### MCP 客户端使用

```python
import mcp

# 创建 MCP 客户端
client = mcp.Client()

# 调用基金基本信息工具
result = client.call_tool("fund.fund_basic", {
    "ts_code": "000001.SZ"
})

# 调用债券柜台交易工具
result = client.call_tool("bond.bond_oc", {})
```

### HTTP MCP API 使用

```bash
# 调用基金基本信息 API
curl -X POST http://localhost:7878/fund \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_TOKEN" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "fund_basic",
      "arguments": {
        "ts_code": "000001.SZ"
      }
    }
  }'

# 调用债券柜台交易 API
curl -X POST http://localhost:7878/bond \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_TOKEN" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/call",
    "params": {
      "name": "bond_oc",
      "arguments": {}
    }
  }'
```

## 配置

### 环境变量

- `TUSHARE_TOKEN` - Tushare API token
- `LOG_LEVEL` - 日志级别 (DEBUG, INFO, WARN, ERROR)
- `MCP_SERVER_PORT` - MCP 服务器端口 (默认: 8080)

### 启动服务器

```bash
# 设置环境变量
export TUSHARE_TOKEN=your_token_here
export LOG_LEVEL=INFO

# 启动 MCP 服务器
go run cmd/mcp-server/main.go
```

## 错误处理

### 常见错误

1. **无效的工具名称**
   - 错误: `unknown tool: invalid.name`
   - 解决: 使用正确的工具名称格式

2. **缺少必需参数**
   - 错误: `missing required parameter: ts_code`
   - 解决: 提供所有必需的参数

3. **API 调用失败**
   - 错误: `failed to call API: ...`
   - 解决: 检查 Tushare token 和网络连接

## 扩展

### 添加新的 API 模块

1. 在 `pkg/sdk/api/` 下创建新的目录
2. 添加 API 实现文件
3. 运行 MCP 工具生成器
4. 重新编译 MCP 服务器

### 自定义工具行为

编辑 `pkg/mcp/tools/` 下对应的工具文件，修改方法实现。

## 维护

### 更新依赖

```bash
go mod tidy
go get -u github.com/chenniannian90/tushare-go
```

### 运行测试

```bash
# 单元测试
go test ./pkg/mcp/...

# 集成测试
go test -tags=integration ./pkg/mcp/...
```

## 支持

如有问题，请查看：
- 项目 README.md
- Tushare 官方文档: https://tushare.pro
- MCP 协议规范: https://modelcontextprotocol.io
