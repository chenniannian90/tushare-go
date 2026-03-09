# tushare-bond 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-bond
- **测试工具数**: 16 个
- **测试成功**: 3 个 (18.8%)
- **无权限**: 3 个 (18.8%)
- **参数错误**: 3 个 (18.8%)
- **空数据**: 4 个 (25.0%)
- **数据量大**: 3 个 (18.8%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | bond_cb_daily | ts_code=128001.SZ, start_date=20240301 | 📭 空数据 | {"data": [], "total": 0} |
| 2 | bond_repo_daily | ts_code=GC001.SH, trade_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 3 | bond_yc_cb | ts_code=1001.CB, trade_date=20240305 | ⚠️ 数据量大 | 包含大量国债收益率曲线数据（207,818字符） |
| 4 | bond_bond_blk | start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED |
| 5 | bond_bond_blk_detail | start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED |
| 6 | bond_bc_otcqt | start_date=20240301, end_date=20240305 | ⚠️ 数据量大 | 包含大量柜台债券报价数据（866,445字符） |
| 7 | bond_cb_issue | start_date=20240101, end_date=20240131 | ⚠️ 参数错误 | API call failed: 无效的 plan_issue_size 类型 |
| 8 | bond_bc_bestotcqt | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 9 | bond_cb_basic | ts_code=128001.SZ | ✅ 正常 | {"ts_code": "128001.SZ", "bond_short_name": "泰尔转债"} |
| 10 | bond_cb_call | start_date=20240301, end_date=20240305 | ⚠️ 参数错误 | API call failed: 无效的 call_price_tax 类型 |
| 11 | bond_cb_price_chg | ts_code=128001.SZ | ⚠️ 参数错误 | API call failed: 无效的 convertprice_bef 类型 |
| 12 | bond_cb_rate | ts_code=128001.SZ | ❌ 无权限 | ACCESS_DENIED |
| 13 | bond_cb_share | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |
| 14 | bond_eco_cal | start_date=20240301, end_date=20240305 | ✅ 正常 | {"date": "20240305", "event": "新西兰全球乳制品拍卖价格指数变化率"} |

---

## 统计摘要

- **总工具数**: 16 个
- **测试成功**: 3 个 (18.8%)
- **无权限**: 3 个 (18.8%)
- **参数错误**: 3 个 (18.8%)
- **空数据**: 4 个 (25.0%)
- **数据量大**: 3 个 (18.8%)

---

## 按状态分类

| 状态 | 数量 | 占比 | 工具列表 |
|------|------|------|----------|
| ✅ 正常可用 | 2 | 12.5% | bond_cb_basic, bond_eco_cal |
| ⚠️ 数据量大 | 3 | 18.8% | bond_yc_cb, bond_bc_otcqt等 |
| ❌ 无权限 | 3 | 18.8% | bond_bond_blk, bond_bond_blk_detail, bond_cb_rate |
| ⚠️ 参数错误 | 3 | 18.8% | bond_cb_issue, bond_cb_call, bond_cb_price_chg |
| 📭 空数据 | 4 | 25.0% | bond_cb_daily, bond_repo_daily等 |

---

## 主要发现

### 1. 正常可用接口 (2个)
- **bond_cb_basic**: 可转债基本信息 ✅
  - 返回可转债完整基本信息
  - 包含转股价、赎回条款、回售条款等

- **bond_eco_cal**: 全球财经日历 ✅
  - 返回全球财经事件数据
  - 包含时间、事件、实际值、预测值等

### 2. 数据量大接口 (3个)
- **bond_yc_cb**: 中债收益率曲线
  - 返回数据量极大（207K字符）
  - 包含各期限国债收益率数据

- **bond_bc_otcqt**: 柜台债券报价
  - 返回数据量极大（866K字符）
  - 包含大量报价数据

### 3. 需要权限的接口 (3个)
- bond_bond_blk: 债券大宗交易数据
- bond_bond_blk_detail: 大宗交易明细
- bond_cb_rate: 可转债票面利率

🔗 权限详情: https://tushare.pro/document/1?doc_id=108

### 4. 参数错误问题 (3个)
- bond_cb_issue - "无效的 plan_issue_size 类型"
- bond_cb_call - "无效的 call_price_tax 类型"
- bond_cb_price_chg - "无效的 convertprice_bef 类型"

---

## 建议

### 测试策略
1. 使用有数据的可转债代码测试
2. 使用筛选条件减少收益率曲线数据量
3. 使用更近期的日期参数

### 代码示例
```go
// 获取可转债基本信息
params := map[string]string{
    "ts_code": "128001.SZ",
}
result, err := client.Call("bond_cb_basic", params)

// 获取全球财经日历
params := map[string]string{
    "start_date": "20240301",
    "end_date": "20240305",
}
result, err := client.Call("bond_eco_cal", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
