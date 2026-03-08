#!/usr/bin/env python3
"""
批量迁移脚本：将 CallAPI 替换为 CallAPIFlexible

这个脚本会自动处理所有使用 CallAPI 的 Go 文件，将其替换为 CallAPIFlexible。

使用���法:
    python3 scripts/migrate_to_flexible_api.py

功能:
    1. 查找所有使用 CallAPI 的文件
    2. 将 CallAPI 替换为 CallAPIFlexible
    3. 保持原有功能不变
    4. 生成迁移报告
"""

import os
import re
import sys
from pathlib import Path
from typing import List, Tuple

# 项目根目录
PROJECT_ROOT = Path(__file__).parent.parent
SDK_API_DIR = PROJECT_ROOT / "pkg" / "sdk" / "api"


def find_go_files_with_call_api() -> List[Path]:
    """查找所有使用 CallAPI 的 Go 文件"""
    go_files = []
    for go_file in SDK_API_DIR.rglob("*.go"):
        # 跳过测试文件
        if go_file.name.endswith("_test.go"):
            continue

        try:
            content = go_file.read_text(encoding='utf-8')
            if '.CallAPI(' in content:
                go_files.append(go_file)
        except Exception as e:
            print(f"⚠️  读取文件失败 {go_file}: {e}")

    return go_files


def migrate_file(file_path: Path) -> Tuple[bool, str]:
    """
    迁移单个文件

    返回: (是否成功, 消息)
    """
    try:
        content = file_path.read_text(encoding='utf-8')

        # 检查是否已经迁移过
        if 'CallAPIFlexible' in content:
            return False, "已经使用 CallAPIFlexible"

        original_content = content

        # 替换 .CallAPI( 为 .CallAPIFlexible(
        content = content.replace('.CallAPI(', '.CallAPIFlexible(')

        # 如果内容没有变化，说明没有需要替换的
        if content == original_content:
            return False, "没有找到需要替换的 CallAPI"

        # 写回文件
        file_path.write_text(content, encoding='utf-8')

        return True, "迁移成功"

    except Exception as e:
        return False, f"迁移失败: {e}"


def migrate_files(files: List[Path], dry_run: bool = False) -> dict:
    """
    批量迁移文件

    返回统计信息
    """
    stats = {
        'total': len(files),
        'success': 0,
        'skipped': 0,
        'failed': 0,
        'errors': []
    }

    print(f"\n🚀 开始迁移 {stats['total']} 个文件...")
    print("=" * 80)

    for i, file_path in enumerate(files, 1):
        # 显示相对路径
        rel_path = file_path.relative_to(PROJECT_ROOT)

        if dry_run:
            print(f"[{i}/{stats['total']}] 🔍 {rel_path} (DRY RUN)")
            stats['skipped'] += 1
            continue

        success, message = migrate_file(file_path)

        if success:
            print(f"[{i}/{stats['total']}] ✅ {rel_path}")
            stats['success'] += 1
        elif "已经使用" in message:
            print(f"[{i}/{stats['total']}] ⏭️  {rel_path} - {message}")
            stats['skipped'] += 1
        else:
            print(f"[{i}/{stats['total']}] ❌ {rel_path} - {message}")
            stats['failed'] += 1
            stats['errors'].append((str(rel_path), message))

    return stats


def main():
    """主函数"""
    print("🔧 Tushare Go SDK - CallAPI 到 CallAPIFlexible 批量迁移工具")
    print("=" * 80)

    # 解析命令行参数
    dry_run = '--dry-run' in sys.argv or '-n' in sys.argv

    if dry_run:
        print("⚠️  DRY RUN 模式：不会实际修改文件\n")

    # 查找需要迁移的文件
    print("🔍 正在查找需要迁移的文件...")
    files = find_go_files_with_call_api()

    if not files:
        print("✅ 没有找到需要迁移的文件")
        return

    print(f"📋 找到 {len(files)} 个文件需要迁移\n")

    # 显示一些示例文件
    print("示例文件:")
    for file_path in files[:5]:
        rel_path = file_path.relative_to(PROJECT_ROOT)
        print(f"  • {rel_path}")

    if len(files) > 5:
        print(f"  ... 还有 {len(files) - 5} 个文件")

    # 确认是否继续
    if not dry_run:
        print("\n⚠️  即将修改这些文件，建议先提交代码")
        response = input("是否继续? (yes/no): ").strip().lower()
        if response not in ['yes', 'y']:
            print("❌ 操作已取消")
            return

    # 执行迁移
    stats = migrate_files(files, dry_run)

    # 显示统计信息
    print("\n" + "=" * 80)
    print("📊 迁移统计:")
    print(f"  • 总文件数: {stats['total']}")
    print(f"  • 成功: {stats['success']} ✅")
    print(f"  • 跳过: {stats['skipped']} ⏭️")
    print(f"  • 失败: {stats['failed']} ❌")

    if stats['errors']:
        print("\n❌ 失败详情:")
        for file_path, error in stats['errors']:
            print(f"  • {file_path}: {error}")

    print("\n✅ 迁移完成!")
    print("\n📝 下一步:")
    print("  1. 运行测试验证: go test ./pkg/sdk/... -v")
    print("  2. 检查编译: go build ./pkg/sdk/...")
    print("  3. 提交代码: git add . && git commit -m 'feat(sdk): migrate all APIs to CallAPIFlexible'")


if __name__ == "__main__":
    main()
