#!/bin/bash

# 测试 new_share API 修复脚本
# 使用方法: ./test_fix.sh YOUR_TOKEN_HERE

set -e

echo "🔧 new_share API 修复验证测试"
echo "================================"
echo ""

# 检查参数
if [ -z "$1" ]; then
    if [ -z "$TUSHARE_TOKEN" ]; then
        echo "❌ 错误: 请提供 Tushare Token"
        echo ""
        echo "使用方法:"
        echo "  ./test_fix.sh YOUR_TOKEN_HERE"
        echo "  或"
        echo "  TUSHARE_TOKEN=YOUR_TOKEN ./test_fix.sh"
        exit 1
    else
        TOKEN=$TUSHARE_TOKEN
    fi
else
    TOKEN=$1
fi

echo "📝 Token: ${TOKEN:0:10}..."
echo ""

# 运行测试程序
echo "🚀 开始测试..."
echo ""

export TUSHARE_TOKEN="$TOKEN"
go run examples/test_new_share_fix.go

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "✅ 测试完成！"
echo ""
echo "📊 测试结果说明:"
echo "   • 如果看到成功获取数据，说明 CallAPIFlexible 工作正常"
echo "   • 如果显示 '当前时间段无IPO数据'，也说明API调用成功"
echo "   • 不应再出现 'json: cannot unmarshal array' 错误"
echo ""
echo "💡 技术细节:"
echo "   • 使用了新的 CallAPIFlexible 方法"
echo "   • 自动检测API返回的是对象数组还是二维数组"
echo "   • 统一转换为对象数组格式供业务代码使用"
echo "   • 保持向后兼容，不影响其他API"
