#!/bin/bash

echo "=== 测试MCP TradeCal工具 (使用stdio传输) ==="
echo ""

# 创建一个临时的配置文件，只包含stock_basic服务
cat > /tmp/test_mcp_config.json << 'EOF'
{
  "transport": "stdio",
  "services": {
    "stock_basic": {
      "name": "stock_basic",
      "path": "/stock_basic",
      "description": "股票基础数据API",
      "categories": ["stock_basic"]
    }
  },
  "api_tokens": [
    "412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
  ]
}
EOF

echo "1. 使用stdio模式测试MCP服务器..."
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' | ./bin/tushare-mcp --config=/tmp/test_mcp_config.json 2>&1 | head -10

echo ""
echo "2. 检查TradeCal工具名称..."
echo '{"jsonrpc":"2.0","id":2,"method":"notifications/initialized"}
{"jsonrpc":"2.0","id":3,"method":"tools/list"}' | ./bin/tushare-mcp --config=/tmp/test_mcp_config.json 2>&1 | grep -i "trade_cal" || echo "未找到trade_cal工具"

echo ""
echo "3. 模拟TradeCal工具调用..."
cat << 'EOF' | ./bin/tushare-mcp --config=/tmp/test_mcp_config.json 2>&1 | head -20
{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}
{"jsonrpc":"2.0","method":"notifications/initialized"}
{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"stock_basic.trade_cal","arguments":{"exchange":"SSE","start_date":"20240101","end_date":"20240105"}}}
EOF

echo ""
echo "4. 清理临时文件..."
rm -f /tmp/test_mcp_config.json

echo ""
echo "=== 测试完成 ==="
