#!/usr/bin/env python3
"""
Parse Tushare API documentation HTML and generate YAML schema
"""

from html import unescape
import re
from pathlib import Path

def parse_html_navigation():
    """Parse the HTML navigation structure and extract API information"""

    html_content = '''
    <nav class="sidebar col-md-3 col-sm-4 col-xs-12">
        <div id="jstree" class="jstree jstree-1 jstree-default" role="tree">
            <ul class="jstree-container-ul jstree-children">
                <!-- Stock Data -->
                <li role="treeitem" aria-level="1">
                    <a class="jstree-anchor" href="/document/2?doc_id=14">股票数据</a>
                    <ul>
                        <li role="treeitem" aria-level="2">
                            <a class="jstree-anchor" href="/document/2?doc_id=24">基础数据</a>
                            <ul>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=25">股票列表</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=329">每日股本（盘前）</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=26">交易日历</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=397">ST股票列表</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=423">ST风险警示板股票</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=398">沪深港通股票列表</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=100">股票曾用名</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=112">上市公司基本信息</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=193">上市公司管理层</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=194">管理层薪酬和持股</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=375">北交所新旧代码对照</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=123">IPO新股上市</a>
                                </li>
                                <li role="treeitem" aria-level="3">
                                    <a class="jstree-anchor" href="/document/2?doc_id=262">股票历史列表</a>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </nav>
    '''

    # Parse the provided HTML structure
    categories = {
        "stock": {
            "name": "股票数据",
            "subcategories": {
                "basic": {
                    "name": "基础数据",
                    "apis": [
                        {"doc_id": "25", "name": "股票列表", "api_name": "stock_basic"},
                        {"doc_id": "329", "name": "每日股本（盘前）", "api_name": "daily_basic"},
                        {"doc_id": "26", "name": "交易日历", "api_name": "trade_cal"},
                        {"doc_id": "397", "name": "ST股票列表", "api_name": "st_basic"},
                        {"doc_id": "423", "name": "ST风险警示板股票", "api_name": "st_stock"},
                        {"doc_id": "398", "name": "沪深港通股票列表", "api_name": "hs_const"},
                        {"doc_id": "100", "name": "股票曾用名", "api_name": "namechange"},
                        {"doc_id": "112", "name": "上市公司基本信息", "api_name": "stock_company"},
                        {"doc_id": "193", "name": "上市公司管理层", "api_name": "stk_managers"},
                        {"doc_id": "194", "name": "管理层薪酬和持股", "api_name": "stk_rewards"},
                        {"doc_id": "375", "name": "北交所新旧代码对照", "api_name": "bse_mapping"},
                        {"doc_id": "123", "name": "IPO新股上市", "api_name": "new_share"},
                        {"doc_id": "262", "name": "股票历史列表", "api_name": "stock_basic_hist"},
                    ]
                },
                "market": {
                    "name": "行情数据",
                    "apis": [
                        {"doc_id": "27", "name": "历史日线", "api_name": "daily"},
                        {"doc_id": "372", "name": "实时日线", "api_name": "daily_std"},
                        {"doc_id": "370", "name": "历史分钟", "api_name": "stk_mins"},
                        {"doc_id": "374", "name": "实时分钟", "api_name": "stk_mins_std"},
                        {"doc_id": "144", "name": "周线行情", "api_name": "weekly"},
                        {"doc_id": "145", "name": "月线行情", "api_name": "monthly"},
                        {"doc_id": "146", "name": "复权行情", "api_name": "adj_factor"},
                        {"doc_id": "336", "name": "周/月线行情(每日更新)", "api_name": "weekly_md"},
                        {"doc_id": "365", "name": "周/月线复权行情(每日更新)", "api_name": "adj_factor_wm"},
                        {"doc_id": "28", "name": "复权因子", "api_name": "adj_factor"},
                        {"doc_id": "315", "name": "实时Tick（爬虫）", "api_name": "idxt"},
                        {"doc_id": "316", "name": "实时成交（爬虫）", "api_name": "idxd"},
                        {"doc_id": "317", "name": "实时排名（爬虫）", "api_name": "idxx"},
                        {"doc_id": "32", "name": "每日指标", "api_name": "daily_basic"},
                        {"doc_id": "109", "name": "通用行情接口", "api_name": "universal"},
                        {"doc_id": "183", "name": "每日涨跌停价格", "api_name": "stk_limit"},
                        {"doc_id": "214", "name": "每日停复牌信息", "api_name": "suspend"},
                        {"doc_id": "48", "name": "沪深股通十大成交股", "api_name": "hk_hold"},
                        {"doc_id": "49", "name": "港股通十大成交股", "api_name": "gg_hold"},
                        {"doc_id": "196", "name": "港股通每日成交统计", "api_name": "gg_trade"},
                        {"doc_id": "197", "name": "港股通每月成交统计", "api_name": "gg_trade_month"},
                        {"doc_id": "255", "name": "备用行情", "api_name": "bak_daily"},
                    ]
                },
                "finance": {
                    "name": "财务数据",
                    "apis": [
                        {"doc_id": "33", "name": "利润表", "api_name": "income"},
                        {"doc_id": "36", "name": "资产负债表", "api_name": "balancesheet"},
                        {"doc_id": "44", "name": "现金流量表", "api_name": "cashflow"},
                        {"doc_id": "45", "name": "业绩预告", "api_name": "forecast"},
                        {"doc_id": "46", "name": "业绩快报", "api_name": "express"},
                        {"doc_id": "103", "name": "分红送股数据", "api_name": "dividend"},
                        {"doc_id": "79", "name": "财务指标数据", "api_name": "fina_indicator"},
                        {"doc_id": "80", "name": "财务审计意见", "api_name": "fina_audit"},
                        {"doc_id": "81", "name": "主营业务构成", "api_name": "fina_mainbz"},
                        {"doc_id": "162", "name": "财报披露日期表", "api_name": "disclosure_date"},
                    ]
                },
                "reference": {
                    "name": "参考数据",
                    "apis": [
                        {"doc_id": "451", "name": "个股异常波动", "api_name": "stk_abnormal"},
                        {"doc_id": "452", "name": "个股严重异常波动", "api_name": "stk_serious"},
                        {"doc_id": "453", "name": "交易所重点提示证券", "api_name": "stk_warning"},
                        {"doc_id": "61", "name": "前十大股东", "api_name": "top10holders"},
                        {"doc_id": "62", "name": "前十大流通股东", "api_name": "top10floatholders"},
                        {"doc_id": "110", "name": "股权质押统计数据", "api_name": "pledge_stat"},
                        {"doc_id": "111", "name": "股权质押明细数据", "api_name": "pledge_detail"},
                        {"doc_id": "124", "name": "股票回购", "api_name": "repurchase"},
                        {"doc_id": "160", "name": "限售股解禁", "api_name": "share_float"},
                        {"doc_id": "161", "name": "大宗交易", "api_name": "block_trade"},
                        {"doc_id": "164", "name": "股票开户数据（停）", "api_name": "user_login"},
                        {"doc_id": "165", "name": "股票开户数据（旧）", "api_name": "user_login_old"},
                        {"doc_id": "166", "name": "股东人数", "api_name": "stk_holder_number"},
                        {"doc_id": "175", "name": "股东增减持", "api_name": "stk_holders_change"},
                    ]
                },
                "special": {
                    "name": "特色数据",
                    "apis": [
                        {"doc_id": "292", "name": "券商盈利预测数据", "api_name": "broker_profit"},
                        {"doc_id": "293", "name": "每日筹码及胜率", "api_name": "cyq_chips"},
                        {"doc_id": "294", "name": "每日筹码分布", "api_name": "cyq_dist"},
                        {"doc_id": "328", "name": "股票技术面因子(专业版）", "api_name": "stk_factor"},
                        {"doc_id": "295", "name": "中央结算系统持股统计", "api_name": "ccass_stat"},
                        {"doc_id": "274", "name": "中央结算系统持股明细", "api_name": "ccass_detail"},
                        {"doc_id": "188", "name": "沪深股通持股明细", "api_name": "hk_hold_detail"},
                        {"doc_id": "353", "name": "股票开盘集合竞价数据", "api_name": "auction"},
                        {"doc_id": "354", "name": "股票收盘集合竞价数据", "api_name": "auction_detail"},
                        {"doc_id": "364", "name": "神奇九转指标", "api_name": "tdx"},
                        {"doc_id": "399", "name": "AH股比价", "api_name": "ah_spot"},
                        {"doc_id": "275", "name": "机构调研数据", "api_name": "org_survey"},
                        {"doc_id": "267", "name": "券商月度金股", "api_name": "broker_rec"},
                    ]
                },
                "margin": {
                    "name": "两融及转融通",
                    "apis": [
                        {"doc_id": "58", "name": "融资融券交易汇总", "api_name": "margin"},
                        {"doc_id": "59", "name": "融资融券交易明细", "api_name": "margin_detail"},
                        {"doc_id": "326", "name": "融资融券标的（盘前）", "api_name": "margin_std"},
                        {"doc_id": "332", "name": "转融券交易汇总(停）", "api_name": "ssl_ol"},
                        {"doc_id": "331", "name": "转融资交易汇总", "api_name": "ssl_lend"},
                        {"doc_id": "333", "name": "转融券交易明细(停）", "api_name": "ssl_ol_detail"},
                        {"doc_id": "334", "name": "做市借券交易汇总(停）", "api_name": "ssl_mkt"},
                    ]
                },
                "moneyflow": {
                    "name": "资金流向数据",
                    "apis": [
                        {"doc_id": "170", "name": "个股资金流向", "api_name": "moneyflow"},
                        {"doc_id": "348", "name": "个股资金流向（THS）", "api_name": "moneyflow_ths"},
                        {"doc_id": "349", "name": "个股资金流向（DC）", "api_name": "moneyflow_dc"},
                        {"doc_id": "371", "name": "板块资金流向（THS)", "api_name": "moneyflow_block_ths"},
                        {"doc_id": "343", "name": "行业资金流向（THS）", "api_name": "moneyflow_ind_ths"},
                        {"doc_id": "344", "name": "板块资金流向（DC）", "api_name": "moneyflow_block_dc"},
                        {"doc_id": "345", "name": "大盘资金流向（DC）", "api_name": "moneyflow_index_dc"},
                        {"doc_id": "47", "name": "沪深港通资金流向", "api_name": "moneyflow_hsgt"},
                    ]
                },
                "toplist": {
                    "name": "打板专题数据",
                    "apis": [
                        {"doc_id": "106", "name": "龙虎榜每日统计单", "api_name": "top_list"},
                        {"doc_id": "107", "name": "龙虎榜机构交易单", "api_name": "top_inst"},
                        {"doc_id": "355", "name": "同花顺涨跌停榜单", "api_name": "limit_list_ths"},
                        {"doc_id": "298", "name": "涨跌停和炸板数据", "api_name": "limit_list"},
                        {"doc_id": "356", "name": "涨停股票连板天梯", "api_name": "limit_list_d"},
                        {"doc_id": "357", "name": "涨停最强板块统计", "api_name": "limit_list_sec"},
                        {"doc_id": "259", "name": "同花顺行业概念板块", "api_name": "concept"},
                        {"doc_id": "260", "name": "同花顺概念和行业指数行情", "api_name": "ths_daily"},
                        {"doc_id": "261", "name": "同花顺行业概念成分", "api_name": "ths_member"},
                        {"doc_id": "362", "name": "东方财富概念板块", "api_name": "concept_em"},
                        {"doc_id": "363", "name": "东方财富概念成分", "api_name": "concept_member_em"},
                        {"doc_id": "382", "name": "东财概念和行业指数行情", "api_name": "em_daily"},
                        {"doc_id": "369", "name": "开盘竞价成交（当日）", "api_name": "auction_sip"},
                        {"doc_id": "311", "name": "市场游资最全名录", "api_name": "hot"},
                        {"doc_id": "312", "name": "游资交易每日明细", "api_name": "hot_detail"},
                        {"doc_id": "320", "name": "同花顺App热榜数", "api_name": "ths_hot"},
                        {"doc_id": "321", "name": "东方财富App热榜", "api_name": "em_hot"},
                        {"doc_id": "376", "name": "通达信板块信息", "api_name": "tdx_sector"},
                        {"doc_id": "377", "name": "通达信板块成分", "api_name": "tdx_member"},
                        {"doc_id": "378", "name": "通达信板块行情", "api_name": "tdx_daily"},
                        {"doc_id": "347", "name": "榜单数据（开盘啦）", "api_name": "bk_lists"},
                        {"doc_id": "351", "name": "题材成分（开盘啦）", "api_name": "bk_members"},
                        {"doc_id": "421", "name": "题材数据（东方财富）", "api_name": "em_ctheme"},
                        {"doc_id": "422", "name": "题材成分（东方财富）", "api_name": "em_ctmember"},
                    ]
                }
            }
        },
        "etf": {
            "name": "ETF专题",
            "subcategories": {
                "basic": {
                    "name": "ETF基础",
                    "apis": [
                        {"doc_id": "454", "name": "深交所ETF实时快照", "api_name": "etf_snap"},
                        {"doc_id": "385", "name": "ETF基本信息", "api_name": "etf_basic"},
                        {"doc_id": "386", "name": "ETF基准指数", "api_name": "etf_base"},
                        {"doc_id": "416", "name": "ETF实时分钟", "api_name": "etf_min_std"},
                        {"doc_id": "387", "name": "ETF历史分钟", "api_name": "etf_min"},
                        {"doc_id": "400", "name": "ETF实时日线", "api_name": "etf_daily_std"},
                        {"doc_id": "127", "name": "ETF日线行情", "api_name": "fund_daily"},
                        {"doc_id": "199", "name": "ETF复权因子", "api_name": "fund_adj"},
                        {"doc_id": "408", "name": "ETF份额规模", "api_name": "etf_size"},
                    ]
                }
            }
        },
        "index": {
            "name": "指数专题",
            "subcategories": {
                "basic": {
                    "name": "指数基础",
                    "apis": [
                        {"doc_id": "94", "name": "指数基本信息", "api_name": "index_basic"},
                        {"doc_id": "95", "name": "指数日线行情", "api_name": "index_daily"},
                        {"doc_id": "403", "name": "指数实时日线", "api_name": "index_daily_std"},
                        {"doc_id": "420", "name": "指数实时分钟", "api_name": "index_min_std"},
                        {"doc_id": "171", "name": "指数周线行情", "api_name": "index_weekly"},
                        {"doc_id": "419", "name": "指数历史分钟", "api_name": "index_min"},
                        {"doc_id": "172", "name": "指数月线行情", "api_name": "index_monthly"},
                        {"doc_id": "96", "name": "指数成分和权重", "api_name": "index_weight"},
                        {"doc_id": "128", "name": "大盘指数每日指标", "api_name": "index_basic"},
                        {"doc_id": "181", "name": "申万行业分类", "api_name": "index_classify"},
                        {"doc_id": "335", "name": "申万行业成分（分级）", "api_name": "index_member"},
                        {"doc_id": "327", "name": "申万行业指数日行情", "api_name": "sw_daily"},
                        {"doc_id": "417", "name": "申万实时行情", "api_name": "sw_daily_std"},
                        {"doc_id": "373", "name": "中信行业成分", "api_name": "citic_member"},
                        {"doc_id": "308", "name": "中信行业指数日行情", "api_name": "citic_daily"},
                        {"doc_id": "211", "name": "国际主要指数", "api_name": "index_global"},
                        {"doc_id": "358", "name": "指数技术面因子(专业版)", "api_name": "index_factor"},
                        {"doc_id": "215", "name": "沪深市场每日交易统计", "api_name": "index_market"},
                        {"doc_id": "268", "name": "深圳市场每日交易情况", "api_name": "index_market_sz"},
                    ]
                }
            }
        },
        "fund": {
            "name": "公募基金",
            "subcategories": {
                "basic": {
                    "name": "基金基础",
                    "apis": [
                        {"doc_id": "19", "name": "基金列表", "api_name": "fund_basic"},
                        {"doc_id": "118", "name": "基金管理人", "api_name": "fund_manager"},
                        {"doc_id": "208", "name": "基金经理", "api_name": "fund_manager_info"},
                        {"doc_id": "207", "name": "基金规模", "api_name": "fund_scale"},
                        {"doc_id": "119", "name": "基金净值", "api_name": "fund_nav"},
                        {"doc_id": "120", "name": "基金分红", "api_name": "fund_dividend"},
                        {"doc_id": "121", "name": "基金持仓", "api_name": "fund_portfolio"},
                        {"doc_id": "359", "name": "基金技术面因子(专业版)", "api_name": "fund_factor"},
                    ]
                }
            }
        }
    }

    return categories

def generate_yaml():
    """Generate YAML file with API schema"""

    categories = parse_html_navigation()

    yaml_lines = [
        'version: "1.0.0"',
        'description: "Tushare Pro API Schema"',
        'categories:',
    ]

    for cat_key, cat_data in categories.items():
        yaml_lines.append(f'  - id: {cat_key}')
        yaml_lines.append(f'    name: "{cat_data["name"]}"')
        yaml_lines.append('    subcategories:')

        for sub_key, sub_data in cat_data['subcategories'].items():
            yaml_lines.append(f'      - id: {sub_key}')
            yaml_lines.append(f'        name: "{sub_data["name"]}"')
            yaml_lines.append('        apis:')

            for api in sub_data['apis']:
                yaml_lines.append(f'          - doc_id: {api["doc_id"]}')
                yaml_lines.append(f'            name: "{api["name"]}"')
                yaml_lines.append(f'            api_name: {api["api_name"]}')
                yaml_lines.append(f'            url: "https://tushare.pro/document/2?doc_id={api["doc_id"]}"')

    return '\n'.join(yaml_lines)

def main():
    """Main function to generate YAML file"""

    yaml_content = generate_yaml()

    # Output YAML file
    schema_dir = Path('/Users/mac-new/go/src/github.com/chenniannian90/tushare-go/schema')
    schema_dir.mkdir(parents=True, exist_ok=True)

    yaml_file = schema_dir / 'api_schema.yaml'

    with open(yaml_file, 'w', encoding='utf-8') as f:
        f.write(yaml_content)

    print(f"✅ Generated YAML file: {yaml_file}")

    # Count categories and APIs
    categories = parse_html_navigation()
    print(f"📊 Total categories: {len(categories)}")

    for cat_key, cat_data in categories.items():
        total_apis = sum(len(sub_data['apis']) for sub_data in cat_data['subcategories'].values())
        print(f"   {cat_data['name']}: {total_apis} APIs")

if __name__ == '__main__':
    main()
