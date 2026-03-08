# 灵活的API响应解决方案

如果某个Tushare API返回不一致的数据格式，可以使用以下灵活的处理方式：

## 方案1: 使用 RawMessage 延迟解析

```go
var result struct {
    Fields []string          `json:"fields"`
    Items  json.RawMessage   `json:"items"` // 延迟解析
}

if err := client.CallAPI(ctx, "api_name", params, fields, &result); err != nil {
    return nil, err
}

// 先尝试解析为数组
var arrayItems []map[string]interface{}
if err := json.Unmarshal(result.Items, &arrayItems); err == nil {
    // 成功解析为数组
    return arrayItems, nil
}

// 再尝试解析为对象
var objectItem map[string]interface{}
if err := json.Unmarshal(result.Items, &objectItem); err == nil {
    // 成功解析为对象，转换为单元素数组
    return []map[string]interface{}{objectItem}, nil
}

// 如果都失败，返回错误
return nil, fmt.Errorf("无法解析items字段")
```

## 方案2: 检查实际的API响应

使用添加的调试日志查看实际的API响应格式，然后：
1. 检查Tushare官方文档
2. 确认该API是否有特殊的响应格式
3. 为该特定API创建专用的处理逻辑

## 方案3: 修改生成模板（如果需要）

如果发现某个API类别总是返回对象格式，可以：
1. 在spec文件中添加特殊标记
2. 修改生成模板以支持该标记
3. 重新生成API代码
