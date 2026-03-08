#!/bin/bash
# 验证代码生成结果

set -e

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

echo "=== 代码生成验证脚本 ==="
echo ""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 统计变量
spec_count=0
spec_with_desc=0
api_count=0
api_with_desc=0
mcp_count=0
mcp_with_desc=0

# 1. 验证Spec文件
echo "1️⃣  验证Spec文件..."
spec_files=$(find internal/gen/specs -name "*.json" | wc -l)
echo "   找到 $spec_files 个spec文件"

spec_with_desc=$(find internal/gen/specs -name "*.json" -exec jq -r '.description' {} \; | grep -v "^$" | wc -l)
echo "   ${GREEN}✓${NC} 有描述的spec文件: $spec_with_desc / $spec_files"

spec_empty_desc=$(find internal/gen/specs -name "*.json" -exec jq -r '.description' {} \; | grep -c "^$" || true)
if [ $spec_empty_desc -gt 0 ]; then
    echo "   ${YELLOW}⚠${NC}  空描述的spec文件: $spec_empty_desc"
fi
echo ""

# 2. 验证API代码
echo "2️⃣  验证API代码..."
api_files=$(find pkg/sdk/api -name "*.go" ! -name "*_test.go" | wc -l)
echo "   找到 $api_files 个API文件"

# 检查是否有函数注释包含描述
api_with_desc=$(grep -r "^// .* 调用 .* API$" pkg/sdk/api | wc -l)
echo "   ${GREEN}✓${NC} 有注释的API函数: $api_with_desc"
echo ""

# 3. 验证MCP工具
echo "3️⃣  验证MCP工具..."
mcp_files=$(find pkg/mcp/tools -name "*.go" ! -name "types.go" ! -name "registry.go" | wc -l)
echo "   找到 $mcp_files 个MCP工具文件"

mcp_with_chinese_desc=$(find pkg/mcp/tools -name "*.go" -exec grep -l "Description: \"获取" {} \; | wc -l)
echo "   ${GREEN}✓${NC} 有中文描述的MCP工具: $mcp_with_chinese_desc"

mcp_with_fallback=$(find pkg/mcp/tools -name "*.go" -exec grep -l "Description: \"Retrieve" {} \; | wc -l)
if [ $mcp_with_fallback -gt 0 ]; then
    echo "   ${YELLOW}⚠${NC}  使用fallback描述的MCP工具: $mcp_with_fallback"
fi
echo ""

# 4. 抽样检查
echo "4️⃣  抽样检查关键API..."

check_stock_basic=$(jq -r '.description' internal/gen/specs/股票数据___stock/基础数据___stock_basic/股票列表___stock_basic.json 2>/dev/null || echo "")
if [ "$check_stock_basic" = "获取基础信息数据，包括股票代码、名称、上市日期、退市日期等" ]; then
    echo "   ${GREEN}✓${NC} stock_basic spec描述正确"
else
    echo "   ${RED}✗${NC} stock_basic spec描述不正确: $check_stock_basic"
fi

check_mcp_stock_basic=$(grep "Description:" pkg/mcp/tools/stock_basic/stock_basic.go 2>/dev/null | head -1 || echo "")
if echo "$check_mcp_stock_basic" | grep -q "获取基础信息数据"; then
    echo "   ${GREEN}✓${NC} stock_basic MCP工具描述正确"
else
    echo "   ${RED}✗${NC} stock_basic MCP工具描述不正确"
fi
echo ""

# 5. 总结
echo "=== 验证总结 ==="
echo "Spec文件: $spec_with_desc/$spec_files 有描述"
echo "API函数: $api_with_desc 个有注释"
echo "MCP工具: $mcp_with_chinese_desc/$mcp_files 有中文描述"
echo ""

# 判断是否成功
if [ $spec_with_desc -gt $((spec_files * 90 / 100)) ] && [ $mcp_with_chinese_desc -gt 100 ]; then
    echo "🎉 ${GREEN}验证通过！${NC} 代码生成质量良好。"
    exit 0
else
    echo "⚠️  ${YELLOW}验证警告：${NC} 某些指标低于预期，请检查生成过程。"
    exit 1
fi
