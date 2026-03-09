# tushare-industry-tmt 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-industry-tmt
- **测试工具数**: 7 个
- **接口错误**: 4 个 (57.1%)
- **空数据**: 2 个 (28.6%)
- **参数错误**: 1 个 (14.3%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | industry_tmt_bo_daily | date=20240305 | ❌ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 2 | industry_tmt_bo_monthly | date=20240301 | ❌ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 3 | industry_tmt_bo_weekly | date=20240304 | ❌ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 4 | industry_tmt_film_record | ann_date=20240305 | ❌ 接口错误 | INVALID_TOKEN: 请指定正确的接口名 |
| 5 | industry_tmt_teleplay_record | report_date=202401 | 📭 空数据 | {"data": [], "total": 0} |
| 6 | industry_tmt_tmt_twincome | date=202401 | ⚠️ 参数错误 | INTERNAL_ERROR: 必填参数, item |
| 7 | industry_tmt_tmt_twincomedetail | date=202401, item=A01 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 7 个
- **接口错误**: 4 个 (57.1%)
- **空数据**: 2 个 (28.6%)
- **参数错误**: 1 个 (14.3%)

---

## 主要发现

### 1. 接口错误
- **industry_tmt_bo_daily**: 接口名称无效

### 2. 参数错误
- **industry_tmt_tmt_twincome**: 缺少必填参数item

### 3. 空数据
- industry_tmt_teleplay_record: 电视剧备案数据

---

## 建议

### 测试策略
1. 使用正确的item参数
2. 使用有数据的日期
3. 继续测试剩余3个工具

### 代码示例
```go
// 获取台湾TMT电子产业营收数据
params := map[string]string{
    "date": "202401",
    "item": "产品代码",
}
result, err := client.Call("industry_tmt_tmt_twincome", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
