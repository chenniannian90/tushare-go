#!/bin/bash

# 验证配置文件的脚本

if [ $# -lt 1 ]; then
    echo "Usage: $0 <config-file>"
    exit 1
fi

CONFIG_FILE=$1

# 检查文件是否存在
if [ ! -f "$CONFIG_FILE" ]; then
    echo "✗ Error: File not found: $CONFIG_FILE"
    exit 1
fi

# 检查JSON格式是否有效
if command -v python &> /dev/null; then
    if ! python -m json.tool "$CONFIG_FILE" > /dev/null 2>&1; then
        echo "✗ Error: Invalid JSON format in $CONFIG_FILE"
        exit 1
    fi
elif command -v jq &> /dev/null; then
    if ! jq empty "$CONFIG_FILE" 2>/dev/null; then
        echo "✗ Error: Invalid JSON format in $CONFIG_FILE"
        exit 1
    fi
else
    echo "⚠ Warning: Cannot validate JSON format (python or jq not found)"
fi

# 显示配置信息
echo "✓ Configuration file: $CONFIG_FILE"
echo ""

# 提取并显示关键配置
if command -v jq &> /dev/null; then
    echo "Configuration summary:"
    jq -r '[
        "  Host: \(.host)",
        "  Port: \(.port)",
        "  Transport: \(.transport)",
        "  Services: \(.services | length)",
        "  Service names: \(.services | join(", "))",
        ""
    ] | join("\n")' "$CONFIG_FILE"
else
    echo "Install jq for better output"
fi

echo "✓ Configuration file is valid!"
