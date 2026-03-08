#!/bin/bash

# 批量迁移脚本：将 CallAPI 替换为 CallAPIFlexible

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "🔧 Tushare Go SDK - API 迁移工具"
echo "================================"
echo ""

# 检查 Python 3 是否安装
if ! command -v python3 &> /dev/null; then
    echo "❌ 错误: 未找到 python3"
    echo "请先安装 Python 3"
    exit 1
fi

# 检查是否在 dry-run 模式
if [[ "$1" == "--dry-run" || "$1" == "-n" ]]; then
    echo "⚠️  DRY RUN 模式：不会实际修改文件"
    echo ""
    python3 "$SCRIPT_DIR/migrate_to_flexible_api.py" --dry-run
else
    echo "📋 即将批量迁移所有 API 文件"
    echo "   • 将 CallAPI 替换为 CallAPIFlexible"
    echo "   • 大约 260 个文件将被修改"
    echo ""
    echo "⚠️  建议先执行 dry-run: $0 --dry-run"
    echo ""
    read -p "确认继续? (yes/no): " confirm

    if [[ "$confirm" != "yes" && "$confirm" != "y" ]]; then
        echo "❌ 操作已取消"
        exit 0
    fi

    python3 "$SCRIPT_DIR/migrate_to_flexible_api.py"
fi

echo ""
echo "✅ 脚本执行完成"
