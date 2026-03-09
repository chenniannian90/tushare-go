#!/bin/bash
# Tushare MCP 服务添加命令列表
# 使用方法:
#   执行命令: ./mcp-services-list.sh
#   只打印不执行: ./mcp-services-list.sh --dry-run 或 ./mcp-services-list.sh -d

API_KEY="412bca00819ea94f31287f3ab54a676d90861306f81c0405275991d1"
BASE_URL="https://tushare.chat168.cn"
DRY_RUN=false

# 解析命令行参数
while [[ $# -gt 0 ]]; do
  case $1 in
    --dry-run|-d)
      DRY_RUN=true
      shift
      ;;
    *)
      echo "未知参数: $1"
      echo "使用方法: $0 [--dry-run|-d]"
      exit 1
      ;;
  esac
done

# 执行或打印命令的函数
run_cmd() {
  if [ "$DRY_RUN" = true ]; then
    echo "$1"
  else
    eval "$1"
  fi
}

# 打印注释的函数
print_comment() {
  if [ "$DRY_RUN" = true ]; then
    echo "$1"
  fi
}

# ==================== MCP 服务器配置 ====================
print_comment "# ==================== MCP 服务器配置 ===================="
print_comment "# 本地服务器: http://127.0.0.1:7878"
print_comment "# 认证方式: X-API-Key header"
print_comment "#"

# ==================== 股票数据 (8个) ====================
print_comment "# ==================== 股票数据 (8个) ===================="
run_cmd "claude mcp add --transport http tushare-stock-basic ${BASE_URL}/stock/stock_basic --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-board ${BASE_URL}/stock/stock_board --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-feature ${BASE_URL}/stock/stock_feature --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-financial ${BASE_URL}/stock/stock_financial --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-fund-flow ${BASE_URL}/stock/stock_fund_flow --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-margin ${BASE_URL}/stock/stock_margin --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-market ${BASE_URL}/stock/stock_market --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-reference ${BASE_URL}/stock/stock_reference --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 港股数据 ====================
print_comment "# ==================== 港股数据 ===================="
run_cmd "claude mcp add --transport http tushare-hk-stock ${BASE_URL}/hk_stock --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 美股数据 ====================
print_comment "# ==================== 美股数据 ===================="
run_cmd "claude mcp add --transport http tushare-us-stock ${BASE_URL}/us_stock --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 指数数据 ====================
print_comment "# ==================== 指数数据 ===================="
run_cmd "claude mcp add --transport http tushare-index ${BASE_URL}/index --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 基金数据 ====================
print_comment "# ==================== 基金数据 ===================="
run_cmd "claude mcp add --transport http tushare-fund ${BASE_URL}/fund --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 债券数据 ====================
print_comment "# ==================== 债券数据 ===================="
run_cmd "claude mcp add --transport http tushare-bond ${BASE_URL}/bond --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== ETF数据 ====================
print_comment "# ==================== ETF数据 ===================="
run_cmd "claude mcp add --transport http tushare-etf ${BASE_URL}/etf --header \"X-API-Key:${API_KEY}\" --scope project"

if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "# 以上为所有 14 个 Tushare MCP 服务的添加命令（仅预览，未执行）"
else
  echo "✅ 所有 14 个 Tushare MCP 服务已添加完成！"
fi