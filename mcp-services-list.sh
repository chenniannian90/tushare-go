#!/bin/bash
# Tushare MCP 服务添加命令列表
# 使用方法:
#   执行命令: ./mcp-services-list.sh
#   只打印不执行: ./mcp-services-list.sh --dry-run 或 ./mcp-services-list.sh -d

API_KEY="YOUR_TOKEN"
BASE_URL="http://localhost:8080"
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

# ==================== 主服务器 ====================
print_comment "# ==================== 主服务器 ===================="
run_cmd "claude mcp add --transport http tushare-main ${BASE_URL}/mcp --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 股票数据 (8个) ====================
print_comment "# ==================== 股票数据 (8个) ===================="
run_cmd "claude mcp add --transport http tushare-stock-basic ${BASE_URL}/api/v1/stock_basic --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-board ${BASE_URL}/api/v1/stock_board --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-feature ${BASE_URL}/api/v1/stock_feature --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-financial ${BASE_URL}/api/v1/stock_financial --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-fund-flow ${BASE_URL}/api/v1/stock_fund_flow --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-margin ${BASE_URL}/api/v1/stock_margin --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-market ${BASE_URL}/api/v1/stock_market --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-stock-reference ${BASE_URL}/api/v1/stock_reference --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 港股数据 ====================
print_comment "# ==================== 港股数据 ===================="
run_cmd "claude mcp add --transport http tushare-hk-stock ${BASE_URL}/api/v1/hk_stock --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 美股数据 ====================
print_comment "# ==================== 美股数据 ===================="
run_cmd "claude mcp add --transport http tushare-us-stock ${BASE_URL}/api/v1/us_stock --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 指数数据 ====================
print_comment "# ==================== 指数数据 ===================="
run_cmd "claude mcp add --transport http tushare-index ${BASE_URL}/api/v1/index --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 基金数据 ====================
print_comment "# ==================== 基金数据 ===================="
run_cmd "claude mcp add --transport http tushare-fund ${BASE_URL}/api/v1/fund --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 期货数据 ====================
print_comment "# ==================== 期货数据 ===================="
run_cmd "claude mcp add --transport http tushare-futures ${BASE_URL}/api/v1/futures --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 债券数据 ====================
print_comment "# ==================== 债券数据 ===================="
run_cmd "claude mcp add --transport http tushare-bond ${BASE_URL}/api/v1/bond --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 外汇数据 ====================
print_comment "# ==================== 外汇数据 ===================="
run_cmd "claude mcp add --transport http tushare-forex ${BASE_URL}/api/v1/forex --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== ETF数据 ====================
print_comment "# ==================== ETF数据 ===================="
run_cmd "claude mcp add --transport http tushare-etf ${BASE_URL}/api/v1/etf --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 期权数据 ====================
print_comment "# ==================== 期权数据 ===================="
run_cmd "claude mcp add --transport http tushare-options ${BASE_URL}/api/v1/options --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 现货数据 ====================
print_comment "# ==================== 现货数据 ===================="
run_cmd "claude mcp add --transport http tushare-spot ${BASE_URL}/api/v1/spot --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== LLM语料数据 ====================
print_comment "# ==================== LLM语料数据 ===================="
run_cmd "claude mcp add --transport http tushare-llm-corpus ${BASE_URL}/api/v1/llm_corpus --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 宏观经济数据 (4个) ====================
print_comment "# ==================== 宏观经济数据 (4个) ===================="
run_cmd "claude mcp add --transport http tushare-macro-business ${BASE_URL}/api/v1/macro_business --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-macro-economy ${BASE_URL}/api/v1/macro_economy --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-macro-interest-rate ${BASE_URL}/api/v1/macro_interest_rate --header \"X-API-Key:${API_KEY}\" --scope project"
run_cmd "claude mcp add --transport http tushare-macro-price ${BASE_URL}/api/v1/macro_price --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 行业数据 ====================
print_comment "# ==================== 行业数据 ===================="
run_cmd "claude mcp add --transport http tushare-industry-tmt ${BASE_URL}/api/v1/industry_tmt --header \"X-API-Key:${API_KEY}\" --scope project"

# ==================== 财富基金销售 ====================
print_comment "# ==================== 财富基金销售 ===================="
run_cmd "claude mcp add --transport http tushare-wealth-fund-sales ${BASE_URL}/api/v1/wealth_fund_sales --header \"X-API-Key:${API_KEY}\" --scope project"

if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "# 以上为所有 26 个 Tushare MCP 服务的添加命令（仅预览，未执行）"
else
  echo "✅ 所有 26 个 Tushare MCP 服务已添加完成！"
fi