#!/usr/bin/env python3
"""
Add detailed documentation comments to all API functions
"""

import re
from pathlib import Path

# API文档映射数据
API_DOCUMENTATION = {
    'stock/market': {
        'file': 'market.go',
        'package': 'market',
        'apis': {
            'Daily': {
                'description': '获取日线行情数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码（支持多选）'},
                    {'name': 'trade_date', 'desc': '交易日期（YYYYMMDD格式）'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'open', 'high', 'low', 'close', 'pre_close', 'change', 'pct_chg', 'vol', 'amount']
            },
            'Weekly': {
                'description': '获取周线行情数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'open', 'high', 'low', 'close', 'vol', 'amount']
            },
            'Monthly': {
                'description': '获取月线行情数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'open', 'high', 'low', 'close', 'vol', 'amount']
            },
        }
    },
    'stock/finance': {
        'file': 'finance.go',
        'package': 'finance',
        'apis': {
            'Income': {
                'description': '获取利润表数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'ann_date', 'f_ann_date', 'end_date', 'report_type', 'comp_type', 'basic_eps', 'diluted_eps']
            },
            'BalanceSheet': {
                'description': '获取资产负债表数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'ann_date', 'end_date', 'report_type', 'comp_type', 'total_assets', 'total_liab', 'equities']
            },
            'CashFlow': {
                'description': '获取现金流量表数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'ann_date', 'end_date', 'comp_type', 'net_cash_flows', 'n_cash_flows_frm_oa']
            },
        }
    },
    'stock/margin': {
        'file': 'margin.go',
        'package': 'margin',
        'apis': {
            'Margin': {
                'description': '获取融资融券交易汇总数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'exchange_id', 'rzye', 'rzmre', 'rzche', 'rqyl', 'rqylchl']
            },
        }
    },
    'stock/reference': {
        'file': 'holder.go',
        'package': 'reference',
        'apis': {
            'Top10Holders': {
                'description': '获取前十大股东数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'period', 'desc': '报告期（如20231231）'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'ann_date', 'end_date', 'holder_name', 'hold_amount', 'hold_ratio']
            },
            'Top10FloatHolders': {
                'description': '获取前十大流通股东数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'period', 'desc': '报告期（如20231231）'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'ann_date', 'end_date', 'holder_name', 'hold_amount', 'hold_ratio']
            },
            'StkHolderNumber': {
                'description': '获取股东人数数据',
                'params': [
                    {'name': 'ts_code', 'desc': '股票代码'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'end_date', 'holder_num', 'holder_num_chg']
            },
        }
    },
    'stock/toplist': {
        'files': ['toplist.go', 'concept.go', 'ths.go', 'limit.go'],
        'package': 'toplist',
        'apis': {
            'TopList': {
                'description': '获取龙虎榜每日统计数据',
                'params': [
                    {'name': 'trade_date', 'desc': '交易日期（YYYYMMDD格式）'},
                    {'name': 'exchange', 'desc': '交易所代码（SSE上交所 SZSE深交所）'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'exalter', 'buy_amount', 'sell_amount', 'abnormal_reason']
            },
            'TopInst': {
                'description': '获取龙虎榜机构交易明细',
                'params': [
                    {'name': 'trade_date', 'desc': '交易日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': 'ts_code trade_date exalter buy_amount sell_amount abnormal_reason'.split()
            },
        }
    },
    'etf': {
        'file': 'etf.go',
        'package': 'etf',
        'apis': {
            'ETFBasic': {
                'description': '获取ETF基础信息',
                'params': [
                    {'name': 'ts_code', 'desc': 'ETF代码'},
                    {'name': 'market', 'desc': '市场类型'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'name', 'list_date', 'list_status', 'fund_type', 'manage_type', 'underlying_index']
            },
            'FundDaily': {
                'description': '获取基金日线行情',
                'params': [
                    {'name': 'ts_code', 'desc': '基金代码'},
                    {'name': 'trade_date', 'desc': '交易日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'open', 'high', 'low', 'close', 'vol', 'amount']
            },
            'FundAdj': {
                'description': '获取基金复权因子',
                'params': [
                    {'name': 'ts_code', 'desc': '基金代码'},
                    {'name': 'trade_date', 'desc': '交易日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'adj_factor']
            },
        }
    },
    'index': {
        'file': 'index.go',
        'package': 'index',
        'apis': {
            'IndexDaily': {
                'description': '获取指数日线行情',
                'params': [
                    {'name': 'ts_code', 'desc': '指数代码'},
                    {'name': 'trade_date', 'desc': '交易日期'},
                    {'name': 'start_date', 'desc': '开始日期'},
                    {'name': 'end_date', 'desc': '结束日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'trade_date', 'close', 'open', 'high', 'low', 'vol', 'amount']
            },
            'IndexBasic': {
                'description': '获取指数基础信息',
                'params': [
                    {'name': 'market', 'desc': '市场代码'},
                    {'name': 'publisher', 'desc': '发布人'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['ts_code', 'name', 'market', 'publisher', 'category', 'base_date', 'base_point', 'list_date']
            },
            'IndexWeight': {
                'description': '获取指数成分和权重',
                'params': [
                    {'name': 'index_code', 'desc': '指数代码'},
                    {'name': 'trade_date', 'desc': '交易日期'},
                    {'name': 'limit', 'desc': '单次返回数据长度'},
                ],
                'fields': ['index_code', 'con_code', 'trade_date', 'weight', 'is_new']
            },
        }
    },
}

def generate_function_doc(api_name, api_info):
    """Generate documentation for a function"""
    description = api_info['description']
    params = api_info.get('params', [])
    fields = api_info.get('fields', [])

    doc_lines = [
        f"// {api_name} {description}",
        "",
    ]

    if params:
        doc_lines.append("// 参数说明:")
        for param in params:
            doc_lines.append(f"//   - {param['name']}: {param['desc']}")
        doc_lines.append("")

    if fields:
        doc_lines.append("// 输出字段:")
        if isinstance(fields, list):
            if len(fields) > 8:
                doc_lines.append(f"//   {', '.join(fields[:8])}...")
            else:
                doc_lines.append(f"//   {', '.join(fields)}")
        else:
            doc_lines.append(f"//   {fields}")
        doc_lines.append("")

    doc_lines.append("// 示例:")
    doc_lines.append("//")
    doc_lines.append(f"// params := map[string]string{{")
    doc_lines.append(f"//     \"{params[0]['name']}\": \"示例值\",")

    if len(params) > 1:
        doc_lines.append(f"//     \"{params[1]['name']}\": \"示例值\",")
    doc_lines.append("// }")
    if fields:
        fields_str = ', '.join([f'"{f}"' for f in fields[:4]])
        if len(fields) > 4:
            fields_str += ", ..."
        doc_lines.append(f"// fields := []string{{{fields_str}}}")
    doc_lines.append(f"// resp, err := client.{api_name}(params, fields)")
    doc_lines.append("")

    return "\n".join(doc_lines)

def process_file(file_path, package_name, apis_info):
    """Process a Go file and add documentation"""
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Find all function definitions
    pattern = r'(// [^\n]*\n)?func \(c \*Client\) (\w+)\('
    matches = list(re.finditer(pattern, content))

    if not matches:
        return False

    modified = False
    for match in matches:
        func_name = match.group(2)
        if func_name not in apis_info:
            continue

        # Get current comment
        current_comment = match.group(1) or ""

        # Generate new documentation
        new_doc = generate_function_doc(func_name, apis_info[func_name])

        # Replace old comment with new documentation
        old_text = match.group(0)
        new_text = new_doc + "func (c *Client) " + func_name + "("

        if old_text != new_text:
            content = content.replace(old_text, new_text, 1)
            modified = True
            print(f"  ✅ {func_name}: 添加文档注释")

    if modified:
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        return True

    return False

def add_documentation_to_files():
    """Add documentation to all API files"""
    base_path = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go')

    processed_count = 0
    total_count = 0

    print("🚀 开始为 API 函数添加文档注释...\\n")

    for category, data in API_DOCUMENTATION.items():
        if 'file' in data:
            file_path = base_path / category / data['file']
            if file_path.exists():
                total_count += 1
                print(f"📄 处理: {category}/{data['file']}")
                if process_file(file_path, data['package'], data['apis']):
                    processed_count += 1
                print()

        elif 'files' in data:
            for filename in data['files']:
                file_path = base_path / category / filename
                if file_path.exists():
                    total_count += 1
                    print(f"📄 处理: {category}/{filename}")
                    if process_file(file_path, data['package'], data['apis']):
                        processed_count += 1
                    print()

    print(f"✅ 完成！")
    print(f"   处理文件数: {processed_count}/{total_count}")

if __name__ == '__main__':
    add_documentation_to_files()
