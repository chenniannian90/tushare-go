package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// APIResponse 通用API响应结构，支持自动检测数据格式
type APIResponse struct {
	Fields []string        `json:"fields"`
	Items  json.RawMessage `json:"items"` // 延迟解析，自动检测格式
}

// ResponseFormat 表示响应数据的格式
type ResponseFormat int

const (
	FormatUnknown ResponseFormat = iota
	FormatObjectArray // 对象数组格式: [{"field": "val"}, ...]
	FormatArrayArray // 二维数组格式: [["val1", "val2"], ...]
)

// DetectFormat 自动检测items的数据格式
func (r *APIResponse) DetectFormat() ResponseFormat {
	// 检查是否为空
	if len(r.Items) == 0 || string(r.Items) == "null" || string(r.Items) == "[]" {
		return FormatArrayArray // 空数组默认为数组格式
	}

	// 查看第一个字符来判断格式
	trimmed := bytes.TrimSpace(r.Items)
	if len(trimmed) == 0 {
		return FormatArrayArray
	}

	firstChar := trimmed[0]
	if firstChar == '{' {
		return FormatObjectArray
	}
	if firstChar == '[' {
		// 需要进一步判断是 [[...]] 还是 [{...}]
		// 跳过第一个 [
		if len(trimmed) > 1 {
			secondChar := trimmed[1]
			if secondChar == '{' {
				return FormatObjectArray // [{...}]
			}
			if secondChar == '[' || secondChar == ']' || secondChar == '"' || secondChar == '-' || secondChar == '0' || secondChar == '1' || secondChar == '2' || secondChar == '3' || secondChar == '4' || secondChar == '5' || secondChar == '6' || secondChar == '7' || secondChar == '8' || secondChar == '9' {
				return FormatArrayArray // [[...]] 或 [value, ...]
			}
		}
	}

	return FormatUnknown
}

// ParseItemsAsMaps 将items解析为对象数组（旧格式，向后兼容）
func (r *APIResponse) ParseItemsAsMaps() ([]map[string]interface{}, error) {
	var items []map[string]interface{}
	if err := json.Unmarshal(r.Items, &items); err != nil {
		return nil, fmt.Errorf("解析为对象数组失败: %w", err)
	}
	return items, nil
}

// ParseItemsAsArrays 将items解析为二���数组（新格式）
func (r *APIResponse) ParseItemsAsArrays() ([][]interface{}, error) {
	var items [][]interface{}
	if err := json.Unmarshal(r.Items, &items); err != nil {
		return nil, fmt.Errorf("解析为数组失败: %w", err)
	}
	return items, nil
}

// ParseItems 自动检测格式并解析为通用类型
func (r *APIResponse) ParseItems() (interface{}, error) {
	format := r.DetectFormat()

	switch format {
	case FormatObjectArray:
		return r.ParseItemsAsMaps()
	case FormatArrayArray:
		return r.ParseItemsAsArrays()
	default:
		// 尝试两种格式
		items, err := r.ParseItemsAsMaps()
		if err == nil {
			return items, nil
		}

		arrays, err := r.ParseItemsAsArrays()
		if err == nil {
			return arrays, nil
		}

		return nil, fmt.Errorf("无法识别数据格式，尝试对象数组失败: %v, 尝试二维数组也失败: %v", err, err)
	}
}

// ConvertArrayRowToMap 将数组行转换为map（根据fields）
// 辅助函数：将二维数组的某一行转为 map[string]interface{}
func ConvertArrayRowToMap(row []interface{}, fields []string) (map[string]interface{}, error) {
	if len(fields) != len(row) {
		return nil, fmt.Errorf("字段数量(%d)与数据数量(%d)不匹配", len(fields), len(row))
	}

	result := make(map[string]interface{}, len(fields))
	for i, field := range fields {
		result[field] = row[i]
	}
	return result, nil
}

// ConvertArrayRowToStruct 将数组行转换为结构体（通过反射）
// 这是一个泛型方法，可以将数组行转换为任意结构体
func ConvertArrayRowToStruct(row []interface{}, fields []string, target interface{}) error {
	// 创建临时map
	tempMap, err := ConvertArrayRowToMap(row, fields)
	if err != nil {
		return err
	}

	// 将map转换为JSON，再反序列化到目标结构体
	jsonData, err := json.Marshal(tempMap)
	if err != nil {
		return fmt.Errorf("序列化临时map失败: %w", err)
	}

	if err := json.Unmarshal(jsonData, target); err != nil {
		return fmt.Errorf("反序列化到目标结构体失败: %w", err)
	}

	return nil
}

// ParseAndConvert 通用的解析和转换方法
// 自动检测格式，如果是数组格式，则转换为对象数组
func (r *APIResponse) ParseAndConvert() ([]map[string]interface{}, error) {
	format := r.DetectFormat()

	switch format {
	case FormatObjectArray:
		// 已经是对象数组，直接返回
		return r.ParseItemsAsMaps()

	case FormatArrayArray:
		// 是二维数组，需要转换为对象数组
		arrays, err := r.ParseItemsAsArrays()
		if err != nil {
			return nil, err
		}

		// 转换每一行
		result := make([]map[string]interface{}, len(arrays))
		for i, row := range arrays {
			itemMap, err := ConvertArrayRowToMap(row, r.Fields)
			if err != nil {
				return nil, fmt.Errorf("转换第%d行失败: %w", i, err)
			}
			result[i] = itemMap
		}
		return result, nil

	default:
		// 最后尝试：先尝试解析为数组，如果成功就转换
		arrays, err := r.ParseItemsAsArrays()
		if err == nil && len(arrays) > 0 {
			// 成功解析为数组，现在转换为对象数组
			result := make([]map[string]interface{}, len(arrays))
			for i, row := range arrays {
				itemMap, err := ConvertArrayRowToMap(row, r.Fields)
				if err != nil {
					return nil, fmt.Errorf("转换第%d行失败: %w", i, err)
				}
				result[i] = itemMap
			}
			return result, nil
		}

		// 如果数组解析也失败，尝试对象数组
		items, err := r.ParseItemsAsMaps()
		return items, err
	}
}
