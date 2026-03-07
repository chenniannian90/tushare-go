#!/usr/bin/env python3
"""
批量优化 Tushare API 规范文件
从 Tushare 文档自动抓取完整的 API 参数和字段信息
"""

import json
import os
import re
import sys
import time
from pathlib import Path
from urllib.parse import urlparse

try:
    import requests
    from bs4 import BeautifulSoup
except ImportError:
    print("请安装依赖: pip install requests beautifulsoup4")
    sys.exit(1)


def extract_api_info_from_url(url):
    """从 Tushare 文档 URL 提取 API 信息"""
    try:
        headers = {
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36'
        }
        response = requests.get(url, headers=headers, timeout=30)
        response.raise_for_status()
        response.encoding = 'utf-8'

        soup = BeautifulSoup(response.text, 'html.parser')

        api_info = {
            'api_name': '',
            'api_code': '',
            'description': '',
            'request_params': [],
            'response_fields': []
        }

        # 提取 API 名称
        for p in soup.find_all('p'):
            text = p.get_text()
            if '接口：' in text:
                match = re.search(r'接口：\s*(\w+)', text)
                if match:
                    api_info['api_name'] = match.group(1)
                    api_info['api_code'] = match.group(1)
                    break

        # 提取描述
        for p in soup.find_all('p'):
            text = p.get_text()
            if '描述：' in text and not api_info['description']:
                match = re.search(r'描述：(.+?)(?:权限：|限量：|$)', text, re.DOTALL)
                if match:
                    api_info['description'] = match.group(1).strip()

        # 提取输入参数
        for h3 in soup.find_all('h3'):
            if '输入参数' in h3.get_text():
                table = h3.find_next('table')
                if table:
                    api_info['request_params'] = extract_table_params(table, 'input')

        # 提取输出参数
        for h3 in soup.find_all('h3'):
            if '输出参数' in h3.get_text():
                table = h3.find_next('table')
                if table:
                    api_info['response_fields'] = extract_table_params(table, 'output')

        return api_info

    except Exception as e:
        print(f"  ❌ 抓取失败: {e}")
        return None


def extract_table_params(table, param_type):
    """从表格中提取参数"""
    params = []
    rows = table.find_all('tr')

    for i, row in enumerate(rows):
        # 跳过表头
        if i == 0:
            continue

        cells = row.find_all('td')
        if len(cells) < 3:
            continue

        param = {
            'name': cells[0].get_text().strip(),
            'type': cells[1].get_text().strip(),
            'description': cells[3].get_text().strip() if len(cells) > 3 else ''
        }

        if param_type == 'input':
            required = cells[2].get_text().strip()
            param['required'] = required in ['Y', 'y', 'true', 'True', '是']
        else:
            # 输出参数有默认显示列
            if len(cells) > 2:
                default_display = cells[2].get_text().strip()
                param['default_display'] = default_display in ['Y', 'y', 'true', 'True', '是']

        if param['name']:
            params.append(param)

    return params


def optimize_spec_file(file_path):
    """优化单个规范文件"""
    print(f"处理文件: {file_path}")

    with open(file_path, 'r', encoding='utf-8') as f:
        spec = json.load(f)

    # 如果已经有完整的参数，跳过
    if len(spec.get('request_params', [])) > 0 or len(spec.get('response_fields', [])) > 0:
        print("  ⏭️  已有参数，跳过")
        return False

    # 获取文档 URL
    url = spec.get('__describe__', {}).get('url', '')
    if not url:
        print("  ⏭️  没有 URL，跳过")
        return False

    # 从文档抓取信息
    api_info = extract_api_info_from_url(url)
    if not api_info:
        return False

    # 更新 spec
    if api_info['api_name']:
        spec['api_name'] = api_info['api_name']
        spec['api_code'] = api_info['api_code']

    if api_info['description']:
        spec['description'] = api_info['description']

    spec['request_params'] = api_info['request_params']
    spec['response_fields'] = api_info['response_fields']

    # 写回文件
    with open(file_path, 'w', encoding='utf-8') as f:
        json.dump(spec, f, ensure_ascii=False, indent=2)

    print("  ✅ 优化完成")
    return True


def main():
    specs_dir = Path("/Users/mac-new/go/src/github.com/chenniannian90/tushare-go/internal/gen/specs")

    if not specs_dir.exists():
        print(f"❌ 目录不存在: {specs_dir}")
        sys.exit(1)

    # 查找所有 JSON 文件
    json_files = list(specs_dir.rglob("*.json"))
    print(f"找到 {len(json_files)} 个规范文件\n")

    success_count = 0
    skip_count = 0
    error_count = 0

    for file_path in json_files:
        try:
            result = optimize_spec_file(file_path)
            if result:
                success_count += 1
            else:
                skip_count += 1

            # 避免请求过快
            time.sleep(1)

        except Exception as e:
            print(f"  ❌ 处理失败: {e}")
            error_count += 1

    print(f"\n{'='*50}")
    print(f"✅ 批量优化完成!")
    print(f"  成功: {success_count}")
    print(f"  跳过: {skip_count}")
    print(f"  失败: {error_count}")


if __name__ == '__main__':
    main()