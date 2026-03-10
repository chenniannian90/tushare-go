#!/bin/bash

# MCP 工具测试脚本

echo "================================"
echo "  MCP 工具测试程序"
echo "================================"
echo ""

# 检查是否设置了 TUSHARE_TOKEN
if [ -z "$TUSHARE_TOKEN" ]; then
    echo "❌ 错误: 请设置 TUSHARE_TOKEN 环境变量"
    echo ""
    echo "使用方法:"
    echo "  export TUSHARE_TOKEN=你的token"
    echo "  ./test.sh"
    echo ""
    exit 1
fi

echo "✅ TUSHARE_TOKEN 已设置"
echo ""
echo "开始测试..."
echo ""

# 运行测试程序
go run cmd/test-mcp-tools/main.go

echo ""
echo "================================"
echo "测试完成！"
echo "================================"
