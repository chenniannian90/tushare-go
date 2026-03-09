#!/bin/bash

# Tushare Go SDK - 测试运行脚本

set -e

echo "================================"
echo "Tushare Go SDK - 测试套件"
echo "================================"
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 统计变量
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# 运行测试并统计结果
echo -e "${YELLOW}运行所有测试...${NC}"
echo ""

# 运行测试并捕获输出
TEST_OUTPUT=$(go test ./tests/... -v 2>&1)
echo "$TEST_OUTPUT"

# 统计测试结果
TOTAL_TESTS=$(echo "$TEST_OUTPUT" | grep -c "^=== RUN" || true)
PASSED_TESTS=$(echo "$TEST_OUTPUT" | grep -c "^--- PASS:" || true)
FAILED_TESTS=$(echo "$TEST_OUTPUT" | grep -c "^--- FAIL:" || true)

echo ""
echo "================================"
echo "测试总结"
echo "================================"
echo "总测试数: $TOTAL_TESTS"
echo -e "${GREEN}通过: $PASSED_TESTS${NC}"
if [ $FAILED_TESTS -gt 0 ]; then
	echo -e "${YELLOW}失败: $FAILED_TESTS${NC}"
fi
echo ""

# 显示失败测试详情
if [ $FAILED_TESTS -gt 0 ]; then
	echo "失败的测试:"
	echo "$TEST_OUTPUT" | grep "^--- FAIL:" -A 1
	echo ""
fi

# 返回适当的退出码
if [ $FAILED_TESTS -gt 0 ]; then
	exit 1
else
	echo -e "${GREEN}✅ 所有测试通过！${NC}"
	exit 0
fi
