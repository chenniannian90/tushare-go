#!/bin/bash

echo "=== Testing MCP Server TradeCal Tool ==="
echo ""

# 使用正确的MCP协议格式
echo "1. 初始化MCP会话..."
curl -s -X POST http://localhost:7878/stock/stock_basic \
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
  }' 2>&1 | head -20

echo ""
echo "2. 调用TradeCal工具..."
curl -s -X POST http://localhost:7878/stock/stock_basic \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
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

echo ""
echo "3. 检查服务器调试日志..."
tail -20 /tmp/mcp-server-debug.log | grep -E "(DEBUG|trade_cal|Error|成功)" || echo "无相关日志"
