# Tushare MCP 服务配置命令

本文档列出了所有可用的 MCP 服务添加命令，用于配置 Claude Desktop 或其他 MCP 客户端。

## 基础配置

**服务器地址**: `http://localhost:8080`
**认证方式**: API Key (通过 `X-API-Key` header)
**传输协议**: HTTP

---

## 主服务器（完整访问）

```bash
claude mcp add --transport http tushare-main http://localhost:8080/mcp --header "X-API-Key:YOUR_TOKEN" --scope project
```

---

## 分类服务命令

### 📈 股票市场数据

#### 股票基础数据
```bash
claude mcp add --transport http tushare-stock-basic http://localhost:8080/api/v1/stock_basic --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票板块数据
```bash
claude mcp add --transport http tushare-stock-board http://localhost:8080/api/v1/stock_board --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票特征数据
```bash
claude mcp add --transport http tushare-stock-feature http://localhost:8080/api/v1/stock_feature --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票财务数据
```bash
claude mcp add --transport http tushare-stock-financial http://localhost:8080/api/v1/stock_financial --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票资金流
```bash
claude mcp add --transport http tushare-stock-fund-flow http://localhost:8080/api/v1/stock_fund_flow --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 融资融券
```bash
claude mcp add --transport http tushare-stock-margin http://localhost:8080/api/v1/stock_margin --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票行情
```bash
claude mcp add --transport http tushare-stock-market http://localhost:8080/api/v1/stock_market --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 股票参考数据
```bash
claude mcp add --transport http tushare-stock-reference http://localhost:8080/api/v1/stock_reference --header "X-API-Key:YOUR_TOKEN" --scope project
```

---

### 🌏 港股数据（统一端点）

```bash
claude mcp add --transport http tushare-hk-stock http://localhost:8080/api/v1/hk_stock --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `hk_basic` - 基础信息
- `hk_daily` - 日线数据
- `hk_cal` - 交易日历
- `hk_min` - 分钟数据
- `hk_factor` - 因子数据

---

### 🇺🇸 美股数据

```bash
claude mcp add --transport http tushare-us-stock http://localhost:8080/api/v1/us_stock --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `us_basic` - 基础信息
- `us_daily` - 日线数据
- `us_cal` - 交易日历
- `us_factor` - 因子数据

---

### 💰 债券数据

```bash
claude mcp add --transport http tushare-bond http://localhost:8080/api/v1/bond --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `api272`, `api392`, `block_trade`, `bond_oc`, `bond_repurchase`
- `bond_zs`, `cb_basic`, `cb_call`, `cb_daily`, `cb_interest`
- `cb_issue`, `cb_redemption`, `global_calendar`

---

### 📊 指数数据

```bash
claude mcp add --transport http tushare-index http://localhost:8080/api/v1/index --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `index_basic` - 指数基础
- `index_daily` - 指数日线
- `index_member` - 指数成员
- `index_weight` - 指数权重
- `api358`, `index_weekly`

---

### 💵 基金数据

```bash
claude mcp add --transport http tushare-fund http://localhost:8080/api/v1/fund --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `fund_basic` - 基金基础
- `fund_nav` - 基金净值
- `fund_div` - 基金分红
- `fund_manager` - 基金经理
- `api359`

---

### 📈 期货数据

```bash
claude mcp add --transport http tushare-futures http://localhost:8080/api/v1/futures --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `fut_basic` - 期货基础
- `fut_daily` - 期货日线
- `fut_weekly` - 期货周线
- `fut_settlement` - 结算数据
- `trade_cal` - 交易日历

---

### 💱 外汇数据

```bash
claude mcp add --transport http tushare-forex http://localhost:8080/api/v1/forex --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `api178`, `forex_daily`

---

### 🏦 ETF 数据

```bash
claude mcp add --transport http tushare-etf http://localhost:8080/api/v1/etf --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `api127`, `api199`, `api385`, `api387`, `api400`, `api408`, `api416`

---

### 📝 期权数据

```bash
claude mcp add --transport http tushare-options http://localhost:8080/api/v1/options --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `opt_basic` - 期权基础
- `opt_daily` - 期权日线
- `opt_min` - 期权分钟

---

### 🎯 现货数据

```bash
claude mcp add --transport http tushare-spot http://localhost:8080/api/v1/spot --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `spot_basic` - 现货基础
- `spot_daily` - 现货日线

---

### 🤖 LLM 语料数据

```bash
claude mcp add --transport http tushare-llm-corpus http://localhost:8080/api/v1/llm_corpus --header "X-API-Key:YOUR_TOKEN" --scope project
```

**可用工具**:
- `announcement` - 公告
- `api143`, `api195`, `einteraction`
- `news_broadcast` - 新闻播报
- `policy` - 政策
- `research_report` - 研究报告

---

### 🏭 宏观经济数据

#### 商业周期
```bash
claude mcp add --transport http tushare-macro-business http://localhost:8080/api/v1/macro_business --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 经济指标
```bash
claude mcp add --transport http tushare-macro-economy http://localhost:8080/api/v1/macro_economy --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 利率
```bash
claude mcp add --transport http tushare-macro-interest-rate http://localhost:8080/api/v1/macro_interest_rate --header "X-API-Key:YOUR_TOKEN" --scope project
```

#### 价格指数
```bash
claude mcp add --transport http tushare-macro-price http://localhost:8080/api/v1/macro_price --header "X-API-Key:YOUR_TOKEN" --scope project
```

---

### 🏢 行业数据

#### TMT 行业
```bash
claude mcp add --transport http tushare-industry-tmt http://localhost:8080/api/v1/industry_tmt --header "X-API-Key:YOUR_TOKEN" --scope project
```

---

### 💎 财富基金销售

```bash
claude mcp add --transport http tushare-wealth-fund-sales http://localhost:8080/api/v1/wealth_fund_sales --header "X-API-Key:YOUR_TOKEN" --scope project
```

---

## 使用示例

### 1. 仅添加主服务器（推荐）
```bash
# 通过主服务器访问所有功能
claude mcp add --transport http tushare-main http://localhost:8080/mcp --header "X-API-Key:YOUR_TOKEN" --scope project
```

### 2. 添加特定分类服务
```bash
# 如果您只关心股票和基金数据
claude mcp add --transport http tushare-stock-basic http://localhost:8080/api/v1/stock_basic --header "X-API-Key:YOUR_TOKEN" --scope project
claude mcp add --transport http tushare-fund http://localhost:8080/api/v1/fund --header "X-API-Key:YOUR_TOKEN" --scope project
```

### 3. 测试服务连接
```bash
# 测试健康检查
curl http://localhost:8080/health

# 测试特定 API
curl -X POST http://localhost:8080/api/v1/hk_stock?tool=hk_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_TOKEN" \
  -d '{"ts_code": "00700.HK", "list_date": "20240101"}'
```

---

## 注意事项

1. **API Key 替换**: 将所有命令中的 `YOUR_TOKEN` 替换为您的实际 API Key
2. **端口配置**: 默认端口是 `8080`，如需修改请更改 URL 中的端口号
3. **统一端点**: 港股模块支持统一端点，可通过 `?tool=xxx` 参数调用不同工具
4. **认证要求**: 所有端点都需要有效的 API Key
5. **CORS 支持**: 服务器支持 CORS，可从浏览器直接调用

---

## 服务总数统计

- **股票相关**: 8 个服务
- **指数数据**: 1 个服务
- **基金数据**: 1 个服务
- **期货数据**: 1 个服务
- **债券数据**: 1 个服务
- **外汇数据**: 1 个服务
- **ETF 数据**: 1 个服务
- **期权数据**: 1 个服务
- **现货数据**: 1 个服务
- **港股数据**: 1 个服务（统一端点）
- **美股数据**: 1 个服务
- **LLM 语料**: 1 个服务
- **宏观经济**: 4 个服务
- **行业数据**: 1 个服务
- **财富基金**: 1 个服务

**总计**: 26 个 MCP 服务端点

---

## 快速部署脚本

```bash
#!/bin/bash
# 批量添加所有 Tushare MCP 服务

API_KEY="YOUR_TOKEN"
BASE_URL="http://localhost:8080"

# 主服务器
claude mcp add --transport http tushare-main ${BASE_URL}/mcp --header "X-API-Key:${API_KEY}" --scope project

# 股票数据
claude mcp add --transport http tushare-stock-basic ${BASE_URL}/api/v1/stock_basic --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-board ${BASE_URL}/api/v1/stock_board --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-feature ${BASE_URL}/api/v1/stock_feature --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-financial ${BASE_URL}/api/v1/stock_financial --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-fund-flow ${BASE_URL}/api/v1/stock_fund_flow --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-margin ${BASE_URL}/api/v1/stock_margin --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-market ${BASE_URL}/api/v1/stock_market --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-stock-reference ${BASE_URL}/api/v1/stock_reference --header "X-API-Key:${API_KEY}" --scope project

# 港股、美股
claude mcp add --transport http tushare-hk-stock ${BASE_URL}/api/v1/hk_stock --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-us-stock ${BASE_URL}/api/v1/us_stock --header "X-API-Key:${API_KEY}" --scope project

# 其他类别
claude mcp add --transport http tushare-bond ${BASE_URL}/api/v1/bond --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-fund ${BASE_URL}/api/v1/fund --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-futures ${BASE_URL}/api/v1/futures --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-index ${BASE_URL}/api/v1/index --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-forex ${BASE_URL}/api/v1/forex --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-etf ${BASE_URL}/api/v1/etf --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-options ${BASE_URL}/api/v1/options --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-spot ${BASE_URL}/api/v1/spot --header "X-API-Key:${API_KEY}" --scope project

# 宏观经济
claude mcp add --transport http tushare-macro-business ${BASE_URL}/api/v1/macro_business --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-macro-economy ${BASE_URL}/api/v1/macro_economy --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-macro-interest-rate ${BASE_URL}/api/v1/macro_interest_rate --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-macro-price ${BASE_URL}/api/v1/macro_price --header "X-API-Key:${API_KEY}" --scope project

# 其他
claude mcp add --transport http tushare-llm-corpus ${BASE_URL}/api/v1/llm_corpus --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-industry-tmt ${BASE_URL}/api/v1/industry_tmt --header "X-API-Key:${API_KEY}" --scope project
claude mcp add --transport http tushare-wealth-fund-sales ${BASE_URL}/api/v1/wealth_fund_sales --header "X-API-Key:${API_KEY}" --scope project

echo "✅ 所有 Tushare MCP 服务已添加完成！"
```

---

**最后更新**: 2026-03-08
**版本**: 1.0.0