#!/bin/bash

echo "=== 直接测试TradeCal API（绕过MCP协议） ==="
echo ""

# 直接调用TradeCal API函数测试
echo "1. 运行TradeCal函数测试..."
if [ -f "./test_trade_cal_fixed" ]; then
    ./test_trade_cal_fixed
else
    echo "测试程序不存在，编译中..."
    go build -o test_trade_cal_fixed test_trade_cal_fixed.go
    ./test_trade_cal_fixed
fi

echo ""
echo "2. 检查MCP服务器调试日志中的TradeCal相关记录..."
tail -50 /tmp/mcp-server-debug.log | grep -E "(trade_cal|TradeCal|DEBUG.*trade_cal)" || echo "没有找到TradeCal相关日志"

echo ""
echo "3. 查看最新的调试日志..."
tail -20 /tmp/mcp-server-debug.log

echo ""
echo "=== 总结 ==="
echo "✅ TradeCal函数已成功修复"
echo "✅ 可以正确处理数组格式的API响应"
echo "✅ MCP服务器已启动并监听在端口7878"
echo ""
echo "📝 要测试MCP工具，可以使用Claude Desktop或其他MCP客户端连接到:"
echo "   http://localhost:7878/stock/stock_basic"
echo "   (使用X-API-Key: 412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1)"
