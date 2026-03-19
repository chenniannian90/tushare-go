#!/usr/bin/env python3
"""
Reorganize API files according to schema hierarchy
Creates subdirectories under stock/, etf/, index/, fund/ based on subcategories
"""

from pathlib import Path
import os
import re

# Hardcoded schema structure based on api_schema.yaml
SCHEMA_STRUCTURE = {
    'stock': {
        'name': '股票数据',
        'subcategories': {
            'basic': {
                'name': '基础数据',
                'apis': ['stock_basic', 'daily_basic', 'trade_cal', 'st_basic', 'st_stock',
                        'hs_const', 'namechange', 'stock_company', 'stk_managers', 'stk_rewards',
                        'bse_mapping', 'new_share', 'stock_basic_hist']
            },
            'market': {
                'name': '行情数据',
                'apis': ['daily', 'daily_std', 'stk_mins', 'stk_mins_std', 'weekly', 'monthly',
                        'adj_factor', 'weekly_md', 'adj_factor_wm', 'idxt', 'idxd', 'idxx',
                        'daily_basic', 'universal', 'stk_limit', 'suspend', 'hk_hold', 'gg_hold',
                        'gg_trade', 'gg_trade_month', 'bak_daily']
            },
            'finance': {
                'name': '财务数据',
                'apis': ['income', 'balancesheet', 'cashflow', 'forecast', 'express', 'dividend',
                        'fina_indicator', 'fina_audit', 'fina_mainbz', 'disclosure_date']
            },
            'reference': {
                'name': '参考数据',
                'apis': ['stk_abnormal', 'stk_serious', 'stk_warning', 'top10holders',
                        'top10floatholders', 'pledge_stat', 'pledge_detail', 'repurchase',
                        'share_float', 'block_trade', 'user_login', 'user_login_old',
                        'stk_holder_number', 'stk_holders_change']
            },
            'special': {
                'name': '特色数据',
                'apis': ['broker_profit', 'cyq_chips', 'cyq_dist', 'stk_factor', 'ccass_stat',
                        'ccass_detail', 'hk_hold_detail', 'auction', 'auction_detail', 'tdx',
                        'ah_spot', 'org_survey', 'broker_rec']
            },
            'margin': {
                'name': '两融及转融通',
                'apis': ['margin', 'margin_detail', 'margin_std', 'ssl_ol', 'ssl_lend',
                        'ssl_ol_detail', 'ssl_mkt']
            },
            'moneyflow': {
                'name': '资金流向数据',
                'apis': ['moneyflow', 'moneyflow_ths', 'moneyflow_dc', 'moneyflow_block_ths',
                        'moneyflow_ind_ths', 'moneyflow_block_dc', 'moneyflow_index_dc', 'moneyflow_hsgt']
            },
            'toplist': {
                'name': '打板专题数据',
                'apis': ['top_list', 'top_inst', 'limit_list_ths', 'limit_list', 'limit_list_d',
                        'limit_list_sec', 'concept', 'ths_daily', 'ths_member', 'concept_em',
                        'concept_member_em', 'em_daily', 'auction_sip', 'hot', 'hot_detail',
                        'ths_hot', 'em_hot', 'tdx_sector', 'tdx_member', 'tdx_daily',
                        'bk_lists', 'bk_members', 'em_ctheme', 'em_ctmember']
            }
        }
    },
    'etf': {
        'name': 'ETF专题',
        'subcategories': {
            'basic': {
                'name': 'ETF基础',
                'apis': ['etf_snap', 'etf_basic', 'etf_base', 'etf_min_std', 'etf_min',
                        'etf_daily_std', 'fund_daily', 'fund_adj', 'etf_size']
            }
        }
    },
    'index': {
        'name': '指数专题',
        'subcategories': {
            'basic': {
                'name': '指数基础',
                'apis': ['index_basic', 'index_daily', 'index_daily_std', 'index_min_std',
                        'index_weekly', 'index_min', 'index_monthly', 'index_weight',
                        'index_classify', 'index_member', 'sw_daily', 'sw_daily_std',
                        'citic_member', 'citic_daily', 'index_global', 'index_factor',
                        'index_market', 'index_market_sz']
            }
        }
    },
    'fund': {
        'name': '公募基金',
        'subcategories': {
            'basic': {
                'name': '基金基础',
                'apis': ['fund_basic', 'fund_manager', 'fund_manager_info', 'fund_scale',
                        'fund_nav', 'fund_dividend', 'fund_portfolio', 'fund_factor']
            }
        }
    }
}

def load_schema():
    """Return the hardcoded schema structure"""
    return SCHEMA_STRUCTURE

def create_directory_structure():
    """Create directory structure based on schema"""
    schema = load_schema()

    base_path = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go')

    # Category to directory mapping
    category_dirs = {
        'stock': base_path / 'stock',
        'etf': base_path / 'etf',
        'index': base_path / 'index',
        'fund': base_path / 'fund',
    }

    created_dirs = []

    for cat_id, category_data in schema.items():
        cat_name = category_data['name']

        if cat_id not in category_dirs:
            print(f"⚠️  Skipping unknown category: {cat_id}")
            continue

        cat_dir = category_dirs[cat_id]

        # Create subdirectories for each subcategory
        for sub_id, subcategory in category_data['subcategories'].items():
            sub_name = subcategory['name']

            # Create subdirectory
            sub_dir = cat_dir / sub_id
            sub_dir.mkdir(parents=True, exist_ok=True)
            created_dirs.append(str(sub_dir.relative_to(base_path)))

            print(f"✅ Created: {sub_dir.relative_to(base_path)}/")

            # Generate README for each subdirectory
            readme_content = f"""# {cat_name} - {sub_name}

## 接口列表

"""
            for api_name in subcategory['apis']:
                readme_content += f"- **{api_name}**\\n"

            readme_file = sub_dir / 'README.md'
            with open(readme_file, 'w', encoding='utf-8') as f:
                f.write(readme_content)

    print(f"\\n📊 Created {len(created_dirs)} directories")

    return created_dirs

def generate_api_mapping():
    """Generate a mapping file showing API to directory relationship"""
    schema = load_schema()

    mapping_file = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go/schema/API_MAPPING.md')
    base_path = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go')

    content = "# API 目录映射\\n\\n本文档显示各 API 接口对应的目录位置。\\n\\n"

    for cat_id, category_data in schema.items():
        cat_name = category_data['name']

        content += f"## {cat_name} (`{cat_id}/` )\\n\\n"

        for sub_id, subcategory in category_data['subcategories'].items():
            sub_name = subcategory['name']

            content += f"### {sub_name} (`{cat_id}/{sub_id}/` )\\n\\n"
            content += "| 接口名称 |\\n"
            content += "|----------|\\n"

            for api_name in subcategory['apis']:
                content += f"| `{api_name}` |\\n"

            content += "\\n"

    with open(mapping_file, 'w', encoding='utf-8') as f:
        f.write(content)

    print(f"✅ Generated API mapping: {mapping_file.relative_to(base_path)}")

def print_structure():
    """Print the directory structure"""
    schema = load_schema()

    print("\\n📁 目录结构预览:")
    print("=" * 60)

    for cat_id, category_data in schema.items():
        cat_name = category_data['name']

        print(f"\\n{cat_id}/  # {cat_name}")

        for sub_id, subcategory in category_data['subcategories'].items():
            sub_name = subcategory['name']
            api_count = len(subcategory['apis'])

            print(f"  {sub_id}/  # {sub_name} ({api_count} APIs)")

            # Show first 3 APIs as examples
            for i, api_name in enumerate(subcategory['apis'][:3]):
                print(f"    - {api_name}")

            if api_count > 3:
                print(f"    ... and {api_count - 3} more")

    print("=" * 60)

def main():
    """Main function"""
    print("🚀 开始重新组织 API 目录结构...\\n")

    # Create directory structure
    created_dirs = create_directory_structure()

    # Generate API mapping
    generate_api_mapping()

    # Print structure preview
    print_structure()

    print(f"\\n✅ 完成！创建了 {len(created_dirs)} 个子目录")
    print("\\n💡 下一步:")
    print("  1. 将各分类下的 API 方法移动到对应的子目录")
    print(" 2. 更新导入路径")
    print(" 3. 运行测试验证")

if __name__ == '__main__':
    main()
