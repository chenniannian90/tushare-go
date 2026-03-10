# Stock Board 工具使用示例

本文档提供了 stock_board MCP 工具的实际使用示例。

## stk_auction (集合竞价) 工具

### 场景1: 获取最新集合竞价数据（推荐）

**MCP 客户端调用**:

```python
import mcp

# 创建 MCP 客户端
client = mcp.Client()

# 调用 stk_auction 工具（不指定日期）
result = client.call_tool("stock_board.stk_auction", {})

# 处理结果
if result["total"] > 0:
    print(f"获取到 {result['total']} 条集合竞价数据")
    # 查看前5条数据
    for item in result["data"][:5]:
        print(f"{item['ts_code']}: 价格={item['price']}, "
              f"成交量={item['vol']}, 换手率={item['turnover_rate']}")
```

**HTTP API 调用**:

```bash
curl -X POST http://localhost:7878/stock_board \
  -H "Content-Type: application/json" \
  -H "X-API-Key: YOUR_TOKEN" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "stk_auction",
      "arguments": {}
    }
  }'
```

### 场景2: 获取指定日期的集合竞价数据

**MCP 客户端调用**:

```python
# 获取今天的集合竞价数据
from datetime import datetime

today = datetime.now().strftime("%Y%m%d")
result = client.call_tool("stock_board.stk_auction", {
    "trade_date": today
})

print(f"日期 {today} 的集合竞价数据:")
print(f"总数据量: {result['total']} 条")
```

### 场景3: 分析集合竞价活跃股票

```python
# 获取集合竞价数据
result = client.call_tool("stock_board.stk_auction", {})

# 筛选有成交的股票
active_stocks = [
    item for item in result["data"]
    if item["vol"] > 0 and item["price"] > 0
]

# 按成交量排序
active_stocks.sort(key=lambda x: x["vol"], reverse=True)

print("集合竞价最活跃的前10只股票:")
for stock in active_stocks[:10]:
    print(f"{stock['ts_code']}: "
          f"价格={stock['price']}, "
          f"成交量={stock['vol']:,}, "
          f"成交额={stock['amount']:,}, "
          f"换手率={stock['turnover_rate']:.4f}")
```

### 场景4: 监控特定股票

```python
# 获取集合竞价数据
result = client.call_tool("stock_board.stk_auction", {})

# 监控特定股票
target_stocks = ["000001.SZ", "600000.SH", "600519.SH"]

print("目标股票集合竞价情况:")
for stock in result["data"]:
    if stock["ts_code"] in target_stocks:
        change_pct = ((stock["price"] - stock["pre_close"]) /
                     stock["pre_close"] * 100) if stock["pre_close"] > 0 else 0

        print(f"{stock['ts_code']}: "
              f"竞价价={stock['price']}, "
              f"昨收={stock['pre_close']}, "
              f"涨跌={change_pct:.2f}%, "
              f"成交量={stock['vol']:,}")
```

## hm_detail (游资营业部明细) 工具

### 场景: 分析游资动向

```python
# 获取指定日期的游资交易明细
result = client.call_tool("stock_board.hm_detail", {
    "trade_date": "20240308"
})

print(f"游资交易明细（共 {result['total']} 条）:\n")

# 按营业部统计
department_stats = {}
for item in result["data"]:
    dept = item.get("exalter", "未知")
    if dept not in department_stats:
        department_stats[dept] = {
            "buy_count": 0,
            "sell_count": 0,
            "total_amount": 0
        }

    # 统计买入卖出
    if item.get("buy_amount", 0) > 0:
        department_stats[dept]["buy_count"] += 1
        department_stats[dept]["total_amount"] += item["buy_amount"]
    if item.get("sell_amount", 0) > 0:
        department_stats[dept]["sell_count"] += 1

# 输出活跃营业部
print("最活跃的游资营业部:")
sorted_depts = sorted(
    department_stats.items(),
    key=lambda x: x[1]["total_amount"],
    reverse=True
)

for dept, stats in sorted_depts[:5]:
    print(f"{dept}: "
          f"买入{stats['buy_count']}次, "
          f"卖出{stats['sell_count']}次, "
          f"总金额={stats['total_amount']:,.0f}")
```

## 错误处理示例

```python
import mcp

client = mcp.Client()

def safe_call_tool(tool_name, arguments):
    """安全调用 MCP 工具"""
    try:
        result = client.call_tool(tool_name, arguments)

        if "error" in result:
            print(f"API 错误: {result['error']}")
            return None

        if result.get("total", 0) == 0:
            print(f"警告: {tool_name} 返回 0 条数据")
            return None

        return result

    except Exception as e:
        print(f"调用失败: {e}")
        return None

# 使用示例
result = safe_call_tool("stock_board.stk_auction", {})
if result:
    print(f"成功获取 {result['total']} 条数据")
```

## 批量查询示例

```python
import time
from datetime import datetime, timedelta

# 获取最近几天的集合竞价数据
def get_auction_data(days=3):
    client = mcp.Client()
    all_data = []

    for i in range(days):
        date = (datetime.now() - timedelta(days=i)).strftime("%Y%m%d")

        result = safe_call_tool("stock_board.stk_auction", {
            "trade_date": date
        })

        if result and result["total"] > 0:
            all_data.extend(result["data"])
            print(f"获取 {date} 数据: {result['total']} 条")

        time.sleep(0.5)  # 避免请求过快

    return all_data

# 使用
auction_data = get_auction_data(days=3)
print(f"总共获取 {len(auction_data)} 条数据")
```

## 最佳实践

### 1. 参数选择

```python
# ✅ 推荐：不指定日期，获取最新数据
result = client.call_tool("stock_board.stk_auction", {})

# ✅ 推荐：指定当前日期
today = datetime.now().strftime("%Y%m%d")
result = client.call_tool("stock_board.stk_auction", {
    "trade_date": today
})

# ❌ 不推荐：使用历史日期（可能返回类型错误）
result = client.call_tool("stock_board.stk_auction", {
    "trade_date": "20250307"  # 可能失败
})
```

### 2. 数据验证

```python
def validate_auction_data(data):
    """验证集合竞价数据"""
    if not data:
        return False

    required_fields = [
        "ts_code", "trade_date", "vol", "price",
        "amount", "pre_close", "turnover_rate"
    ]

    for item in data:
        for field in required_fields:
            if field not in item:
                print(f"缺少字段: {field}")
                return False

    return True

# 使用
result = client.call_tool("stock_board.stk_auction", {})
if validate_auction_data(result.get("data", [])):
    print("数据验证通过")
```

### 3. 性能优化

```python
# 批量处理数据
def process_auction_data(data):
    """批量处理集合竞价数据"""
    # 使用列表推导式
    active_stocks = [
        item for item in data
        if item.get("vol", 0) > 0 and item.get("price", 0) > 0
    ]

    # 使用 pandas 进行数据分析（如果可用）
    try:
        import pandas as pd
        df = pd.DataFrame(active_stocks)
        return df
    except ImportError:
        return active_stocks

# 使用
result = client.call_tool("stock_board.stk_auction", {})
processed = process_auction_data(result["data"])
```

---

## 相关文档

- [Stock Board MCP 最佳实践](./STOCK_BOARD_MCP_BEST_PRACTICES.md)
- [MCP 工具总览](./MCP_TOOLS.md)
- [测试报告](../tests/stk_auction_improved_report.md)

---

**最后更新**: 2026-03-09
**维护者**: tushare-go 项目组
