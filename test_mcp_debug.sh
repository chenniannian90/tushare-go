#!/bin/bash

# Test MCP Server with debug logging
echo "Testing MCP Server Debug Logging..."
echo ""

# Test a simple tool call - let's try the trade calendar API
echo "1. Testing Trade Calendar API..."
curl -s -X POST http://localhost:7878/stock/mcp \
  -H "Content-Type: application/json" \
  -H "X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "stock_basic.trade_cal",
      "arguments": {
        "exchange": "SSE",
        "start_date": "20240101",
        "end_date": "20240105"
      }
    }
  }' 2>&1 | jq '.' || echo "Request failed"

echo ""
echo "2. Checking debug logs..."
tail -50 /tmp/mcp-server-debug.log | grep -E "(DEBUG|trade_cal|Error)" || echo "No debug logs found yet"

echo ""
echo "3. Full recent logs:"
tail -20 /tmp/mcp-server-debug.log
