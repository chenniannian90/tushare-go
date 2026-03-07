#!/usr/bin/env python3
"""
检查和修复文件编码问题的脚本
"""
import os
import sys
import json
import re
from pathlib import Path

def check_and_fix_file(filepath):
    """检查并修复单个文件的编码问题"""
    issues = []
    fixed = False

    try:
        # 尝试以 UTF-8 读取
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
            data = json.loads(content)
        return False, issues  # 文件正常，无需修复
    except UnicodeDecodeError:
        # UTF-8 解码失败，尝试修复
        try:
            with open(filepath, 'r', encoding='latin-1') as f:
                content = f.read()
            # 尝试检测是否是双重编码
            try:
                # 如果是 UTF-8 被错误地按 latin-1 解码再编码
                fixed_content = content.encode('latin-1').decode('utf-8')
                # 验证修复后的内容是否是有效的 JSON
                json.loads(fixed_content)
                content = fixed_content
            except:
                # 如果双重编码修复失败，保持原样
                pass

            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            issues.append(f"Fixed encoding in: {os.path.basename(filepath)}")
            return True, issues
        except Exception as e:
            issues.append(f"Failed to fix {os.path.basename(filepath)}: {str(e)}")
            return False, issues
    except json.JSONDecodeError as e:
        # JSON 格式错误，尝试清理
        try:
            with open(filepath, 'r', encoding='utf-8') as f:
                content = f.read()

            # 移除控制字符
            content = re.sub(r'[\x00-\x08\x0B\x0C\x0E-\x1F\x7F]', '', content)

            # 验证清理后的内容
            json.loads(content)

            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)
            issues.append(f"Sanitized JSON in: {os.path.basename(filepath)}")
            return True, issues
        except Exception as e:
            issues.append(f"Failed to sanitize {os.path.basename(filepath)}: {str(e)}")
            return False, issues

def scan_directory(directory):
    """扫描目录并修复所有 JSON 文件"""
    target_dir = Path(directory)

    if not target_dir.exists():
        print(f"❌ Error: Directory '{directory}' does not exist")
        sys.exit(1)

    print(f"🔍 Scanning directory: {directory}")
    print()

    total_files = 0
    checked_files = 0
    fixed_files = 0
    content_issues = []

    # 递归查找所有 JSON 文件
    for filepath in target_dir.rglob('*.json'):
        total_files += 1
        checked_files += 1

        is_fixed, issues = check_and_fix_file(filepath)
        if is_fixed:
            fixed_files += 1
            content_issues.extend(issues)

        # 每处理 50 个文件显示一次进度
        if total_files % 50 == 0:
            print(f"   Processed {total_files} files...")

    # 打印摘要
    print()
    print("📊 Summary:")
    print(f"   Total files found: {total_files}")
    print(f"   Files checked: {checked_files}")
    print(f"   Files fixed: {fixed_files}")
    print()

    if content_issues:
        print("🔧 Content issues fixed:")
        for issue in content_issues:
            print(f"   - {issue}")
        print()

    if fixed_files == 0:
        print("✅ No encoding issues found!")
    else:
        print(f"✅ Fixed {fixed_files} file(s) with encoding issues")

if __name__ == '__main__':
    target_dir = sys.argv[1] if len(sys.argv) > 1 else 'internal/gen/specs'
    scan_directory(target_dir)