#!/usr/bin/env python3
"""
根据 specs 目录重新生成 pkg/sdk/api 的内容
保持目录结构和文件名一致，但目录名和文件名只取 ___ 后面的部分
"""
import os
import sys
import json
import re
from pathlib import Path

# 添加项目根目录到 Python 路径
project_root = Path(__file__).parent.parent.parent
sys.path.insert(0, str(project_root / "internal/gen"))

try:
    from generator import Generate
except ImportError:
    print("错误: 无法导入 generator 模块")
    print("请确保项目结构正确，并已安装必要的依赖")
    sys.exit(1)


def extract_name_from_path(path_component):
    """从路径组件中提取 ___ 后面的部分"""
    if '___' in path_component:
        return path_component.split('___')[-1]
    return path_component


def get_relative_specs_path(spec_file, specs_root):
    """获取 spec 文件相对于 specs 根目录的路径"""
    spec_path = Path(spec_file)
    relative_path = spec_path.relative_to(specs_root)
    return relative_path


def generate_output_path(spec_file, specs_root, output_root):
    """根据 spec 文件路径生成输出文件路径"""
    relative_path = get_relative_specs_path(spec_file, specs_root)

    # 处理路径中的每个组件
    new_components = []
    for component in relative_path.parts:
        # 去掉 .json 后缀，提取 ___ 后面的部分
        if component.endswith('.json'):
            component = component[:-5]
        new_component = extract_name_from_path(component)
        new_components.append(new_component)

    # 构建输出路径（最后一个组件是文件名，添加 .go 后缀）
    output_dir = output_root / Path(*new_components[:-1])
    output_file = output_dir / f"{new_components[-1]}.go"

    return output_file


def load_spec(spec_file):
    """加载 spec 文件"""
    try:
        with open(spec_file, 'r', encoding='utf-8') as f:
            data = json.load(f)

        # 转换为 APISpec 结构
        from spec import APISpec

        spec = APISpec(
            api_name=data.get('api_name', ''),
            api_code=data.get('api_code', ''),
            description=data.get('description', ''),
            __describe__=data.get('__describe__', {}),
            request_params=data.get('request_params', []),
            response_fields=data.get('response_fields', []),
        )

        return spec
    except Exception as e:
        print(f"警告: 无法加载 spec 文件 {spec_file}: {e}")
        return None


def main():
    """主函数"""
    # 路径配置
    project_root = Path(__file__).parent.parent.parent
    specs_root = project_root / "internal/gen/specs"
    output_root = project_root / "pkg/sdk/api"

    if not specs_root.exists():
        print(f"错误: specs 目录不存在: {specs_root}")
        sys.exit(1)

    print(f"🔍 扫描目录: {specs_root}")
    print(f"📁 输出目录: {output_root}")
    print()

    # 统计信息
    total_files = 0
    success_count = 0
    error_count = 0

    # 遍历所有 JSON 文件
    for spec_file in specs_root.rglob("*.json"):
        total_files += 1

        # 生成输出路径
        output_file = generate_output_path(spec_file, specs_root, output_root)

        # 加载 spec
        spec = load_spec(spec_file)
        if spec is None:
            error_count += 1
            continue

        # 生成代码
        try:
            # 确保输出目录存在
            output_file.parent.mkdir(parents=True, exist_ok=True)

            # 使用现有的生成器
            from generator import Generate
            Generate(spec, str(output_file))

            success_count += 1

            # 显示进度
            if total_files % 50 == 0:
                print(f"   已处理 {total_files} 个文件...")

        except Exception as e:
            print(f"错误: 生成 {spec_file} 失败: {e}")
            error_count += 1

    # 打印摘要
    print()
    print("📊 生成摘要:")
    print(f"   总文件数: {total_files}")
    print(f"   成功: {success_count}")
    print(f"   失败: {error_count}")
    print()

    if error_count == 0:
        print("✅ 所有文件生成成功！")
    else:
        print(f"⚠️  有 {error_count} 个文件生成失败")


if __name__ == "__main__":
    main()