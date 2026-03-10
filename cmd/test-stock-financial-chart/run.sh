#!/bin/bash

# Tushare Stock Financial API 测试脚本 - 图表展示模式
#
# 使用方法:
#   ./run.sh

set -e  # 遇到错误立即退出

echo "========================================="
echo "  Tushare Stock Financial API 测试"
echo "  📊 图表展示模式"
echo "========================================="
echo ""

# 检查是否设置了 TUSHARE_TOKEN
if [ -z "$TUSHARE_TOKEN" ]; then
    echo "❌ 错误: 未设置 TUSHARE_TOKEN 环境变量"
    echo ""
    echo "请使用以下方式之一设置你的 Token:"
    echo ""
    echo "1. 导出环境变量:"
    echo "   export TUSHARE_TOKEN=\"your_token_here\""
    echo "   ./run.sh"
    echo ""
    echo "2. 在同一命令中设置:"
    echo "   TUSHARE_TOKEN=\"your_token_here\" ./run.sh"
    echo ""
    echo "3. 直接运行 Go 程序:"
    echo "   TUSHARE_TOKEN=\"your_token_here\" go run main.go"
    echo ""
    echo "获取 Token: https://tushare.pro"
    echo "========================================="
    exit 1
fi

echo "✅ Token 已设置: ${TUSHARE_TOKEN:0:10}..."
echo ""

# 运行测试程序
echo "🚀 开始运行图表展示测试..."
echo ""

go run main.go

echo ""
echo "========================================="
echo "✅ 测试脚本执行完成"
echo ""
echo "📄 HTML 报告已生成: test_report.html"
echo "   请在浏览器中打开查看可视化报告"
echo "========================================="
