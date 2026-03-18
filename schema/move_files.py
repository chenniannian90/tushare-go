#!/usr/bin/env python3
"""
Move API files to new subdirectory structure based on schema mapping
"""

import shutil
from pathlib import Path
import re

# API method to subdirectory mapping
API_TO_SUBDIR = {
    # Stock basic APIs
    'StockBasic': 'stock/basic',
    'BakBasic': 'stock/basic',
    'TradeCal': 'stock/basic',
    'HSConst': 'stock/basic',
    'NameChange': 'stock/basic',
    'StockCompany': 'stock/basic',
    'NewShare': 'stock/basic',

    # Market APIs
    'Daily': 'stock/market',
    'Weekly': 'stock/market',
    'Monthly': 'stock/market',
    'DailyBasic': 'stock/market',
    'AdjFactor': 'stock/market',
    'Suspend': 'stock/market',
    'RTK': 'stock/market',
    'RealTimeQuote': 'stock/market',
    'RealTimeTick': 'stock/market',
    'RealTimeList': 'stock/market',
    'MoneyFlow': 'stock/market',

    # Finance APIs
    'InCome': 'stock/finance',
    'BalanceSheet': 'stock/finance',
    'CashFlow': 'stock/finance',
    'Forecast': 'stock/finance',
    'Dividend': 'stock/finance',
    'Express': 'stock/finance',
    'FinaIndicator': 'stock/finance',
    'FinaAudit': 'stock/finance',
    'FinaMainbz': 'stock/finance',
    'DisclosureDate': 'stock/finance',

    # Holder APIs
    'Top10Holders': 'stock/reference',
    'Top10FloatHolders': 'stock/reference',
    'StkHolderNumber': 'stock/reference',

    # Pledge APIs
    'PledgeStat': 'stock/reference',
    'PledgeDetail': 'stock/reference',

    # Repurchase APIs
    'Repurchase': 'stock/reference',
    'ShareFloat': 'stock/reference',

    # Hsgt APIs
    'MoneyflowHsgt': 'stock/moneyflow',
    'HsgtTop10': 'stock/moneyflow',
    'GgtTop10': 'stock/moneyflow',

    # Margin APIs
    'Margin': 'stock/margin',
    'MarginDetail': 'stock/margin',

    # Toplist APIs
    'TopList': 'stock/toplist',
    'TopInst': 'stock/toplist',

    # Concept APIs
    'Concept': 'stock/toplist',
    'ConceptDetail': 'stock/toplist',

    # Ths APIs
    'ThsIndex': 'stock/toplist',
    'ThsDaily': 'stock/toplist',
    'ThsMember': 'stock/toplist',
    'MoneyflowThs': 'stock/toplist',
    'MoneyflowIndThs': 'stock/toplist',

    # Sw APIs
    'CiDaily': 'stock/toplist',
    'SwDaily': 'stock/toplist',

    # Limit APIs
    'LimitList': 'stock/toplist',
    'STKLimit': 'stock/toplist',

    # Research APIs
    'CyqChips': 'stock/special',
    'StkSurv': 'stock/special',

    # ETF APIs
    'ETFBasic': 'etf/basic',
    'FundDaily': 'etf/basic',
    'FundAdj': 'etf/basic',

    # Index APIs
    'IndexDaily': 'index/basic',
    'IndexBasic': 'index/basic',
    'IndexWeight': 'index/basic',
}

def extract_methods_from_file(file_path):
    """Extract method names from a Go file"""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()

        # Find all method definitions
        pattern = r'func \(c \*Client\) (\w+)\('
        methods = re.findall(pattern, content)
        return methods
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        return []

def determine_target_subdir(file_path, methods):
    """Determine the target subdirectory based on methods in the file"""
    if not methods:
        return None

    # Count occurrences of each subdirectory
    subdir_counts = {}
    for method in methods:
        if method in API_TO_SUBDIR:
            subdir = API_TO_SUBDIR[method]
            subdir_counts[subdir] = subdir_counts.get(subdir, 0) + 1

    if not subdir_counts:
        return None

    # Return the most common subdirectory
    return max(subdir_counts.items(), key=lambda x: x[1])[0]

def move_file_to_subdir(source_file, target_subdir):
    """Move a file to a subdirectory"""
    target_path = Path(target_subdir) / source_file.name

    # Create target directory if it doesn't exist
    target_path.parent.mkdir(parents=True, exist_ok=True)

    # Check if target file already exists
    if target_path.exists():
        print(f"  ⚠️  Target already exists: {target_path}")
        return False

    # Copy the file
    shutil.copy2(source_file, target_path)
    print(f"  ✅ Copied: {source_file} -> {target_path}")
    return True

def process_reorganization():
    """Process the reorganization of files"""
    base_path = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go')

    # Files to process
    files_to_process = [
        'stock/stock.go',
        'market/market.go',
        'finance/finance.go',
        'holder/holder.go',
        'pledge/pledge.go',
        'margin/margin.go',
        'hsgt/hsgt.go',
        'toplist/toplist.go',
        'concept/concept.go',
        'ths/ths.go',
        'sw/sw.go',
        'limit/limit.go',
        'research/research.go',
        'repurchase/repurchase.go',
        'realtime/realtime.go',
        'etf/etf.go',
        'index/index.go',
    ]

    moved_count = 0
    skipped_count = 0

    print("🚀 开始重新组织 API 文件...\n")

    for file_rel in files_to_process:
        source_file = base_path / file_rel

        if not source_file.exists():
            print(f"⚠️  File not found: {file_rel}")
            skipped_count += 1
            continue

        print(f"📄 Processing: {file_rel}")

        # Extract methods from the file
        methods = extract_methods_from_file(source_file)
        print(f"   Found {len(methods)} methods: {', '.join(methods[:3])}{'...' if len(methods) > 3 else ''}")

        # Determine target subdirectory
        target_subdir = determine_target_subdir(source_file, methods)

        if target_subdir:
            # Create full target path
            target_dir = base_path / target_subdir
            target_path = target_dir / source_file.name

            # Create target directory if it doesn't exist
            target_dir.mkdir(parents=True, exist_ok=True)

            # Copy the file
            shutil.copy2(source_file, target_path)
            print(f"   ✅ Moved to: {target_subdir}/")
            moved_count += 1
        else:
            print(f"   ⚠️  No matching subdirectory found, keeping in place")
            skipped_count += 1

        print()

    print(f"✅ 重组完成!")
    print(f"   移动文件: {moved_count}")
    print(f"   跳过文件: {skipped_count}")

    return moved_count, skipped_count

def update_package_declarations():
    """Update package declarations in moved files"""
    base_path = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go')

    # Package name mappings
    package_mappings = {
        'stock/basic': 'basic',
        'stock/market': 'market',
        'stock/finance': 'finance',
        'stock/reference': 'reference',
        'stock/special': 'special',
        'stock/margin': 'margin',
        'stock/moneyflow': 'moneyflow',
        'stock/toplist': 'toplist',
        'etf/basic': 'basic',
        'index/basic': 'basic',
    }

    print("\\n🔧 更新 package 声明...")

    updated_count = 0

    for subdir, package_name in package_mappings.items():
        subdir_path = base_path / subdir
        if not subdir_path.exists():
            continue

        for go_file in subdir_path.glob("*.go"):
            try:
                with open(go_file, 'r', encoding='utf-8') as f:
                    content = f.read()

                # Update package declaration
                updated_content = re.sub(
                    r'^package \\w+',
                    f'package {package_name}',
                    content,
                    flags=re.MULTILINE
                )

                if content != updated_content:
                    with open(go_file, 'w', encoding='utf-8') as f:
                        f.write(updated_content)
                    print(f"   ✅ Updated: {go_file.relative_to(base_path)}")
                    updated_count += 1

            except Exception as e:
                print(f"   ⚠️  Error updating {go_file}: {e}")

    print(f"\\n   更新了 {updated_count} 个文件的 package 声明")

def main():
    """Main function"""
    moved, skipped = process_reorganization()
    update_package_declarations()

    print("\\n💡 下一步:")
    print("  1. 检查移动的文件内容")
    print("  2. 更新导入路径")
    print("  3. 运行 go mod tidy")
    print("  4. 运行测试验证功能")

if __name__ == '__main__':
    main()
