#!/bin/bash

echo "=== 完整的MCP TradeCal工具测试 ==="
echo ""

# 使用临时文件来模拟完整的MCP会话
TEMP_FILE=$(mktemp)

# 开始一个curl会话，保持连接打开
echo "1. 初始化MCP会话并保持连接..."
curl -s -N -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {
        "name": "test-client",
        "version": "1.0.0"
      }
    }
  }' > "$TEMP_FILE" 2>&1 &

CURL_PID=$!
sleep 2

echo "2. 发送initialized通知..."
curl -s -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "method": "notifications/initialized"
  }' 2>&1 | head -10

echo ""
echo "3. 获取工具列表..."
curl -s -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }' 2>&1 | head -30

echo ""
echo "4. 调用TradeCal工具..."
curl -s -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "tools/call",
    "params": {
      "name": "stock_basic.trade_cal",
      "arguments": {
        "exchange": "SSE",
        "start_date": "20240101",
        "end_date": "20240105"
      }
    }
  }' 2>&1 | head -30

# 清理
kill $CURL_PID 2>/dev/null
rm -f "$TEMP_FILE"

echo ""
echo "5. 检查服务器调试日志..."
echo "=== 最近的调试日志 ==="
tail -30 /tmp/mcp-server-debug.log | grep -E "(DEBUG|trade_cal|Error|成功|TradeCal|items)" || tail -10 /tmp/mcp-server-debug.log
