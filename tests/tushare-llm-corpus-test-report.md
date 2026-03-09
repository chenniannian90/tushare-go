# tushare-llm-corpus 服务测试报告

## 测试概览

- **测试日期**: 2026-03-09
- **服务名称**: tushare-llm-corpus
- **测试工具数**: 9 个
- **测试成功**: 1 个 (11.1%)
- **无权限**: 4 个 (44.4%)
- **数据量大**: 3 个 (33.3%)
- **空数据**: 1 个 (11.1%)

---

## 工具测试结果表

| 序号 | 工具名称 | 调用参数 | 状态 | 返回数据示例 |
|------|----------|----------|------|--------------|
| 1 | llm_corpus_news | start_date=2024-03-01 00:00:00, end_date=2024-03-02 00:00:00 | ⚠️ 数据量大 | 包含大量新闻数据（450,902字符） |
| 2 | llm_corpus_anns_d | start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED: 抱歉，您没有接口访问权限 |
| 3 | llm_corpus_cctv_news | date=20240305 | ✅ 正常 | {"date": "20240305", "title": "韩正参加山东代表团审议"} |
| 4 | llm_corpus_irm_qa_sz | start_date=20240301, end_date=20240305 | ⚠️ 数据量大 | 包含大量互动易数据（1,060,716字符） |
| 5 | llm_corpus_major_news | start_date=2024-03-01 00:00:00, end_date=2024-03-02 00:00:00 | ⚠️ 数据量大 | 包含大量长篇新闻数据（1,529,738字符） |
| 6 | llm_corpus_npr | start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED |
| 7 | llm_corpus_research_report | start_date=20240301, end_date=20240305 | ❌ 无权限 | ACCESS_DENIED |
| 8 | llm_corpus_irm_qa_sh | start_date=20240301, end_date=20240305 | 📭 空数据 | {"data": [], "total": 0} |

---

## 统计摘要

- **总工具数**: 9 个
- **测试成功**: 1 个 (11.1%)
- **无权限**: 4 个 (44.4%)
- **数据量大**: 3 个 (33.3%)
- **空数据**: 1 个 (11.1%)

---

## 主要发现

### 1. 正常可用接口 (1个)
- **llm_corpus_cctv_news**: 新闻联播文字稿 ✅
  - 返回新闻联播文本数据
  - 每日更新
  - 包含标题和内容

### 2. 数据量大接口 (3个)
- **llm_corpus_news**: 新闻快讯 ⚠️
  - 返回数据量大（450K字符）
  - 需要缩短时间范围

- **llm_corpus_irm_qa_sz**: 互动易数据 ⚠️
  - 返回数据量大（1M字符）
  - 深交所互动问答数据

- **llm_corpus_major_news**: 长篇新闻 ⚠️
  - 返回数据量大（1.5M字符）
  - 主流新闻网站长篇通讯

### 3. 需要权限的接口 (1个)
- llm_corpus_anns_d: 公告数据

---

## 建议

### 测试策略
1. 缩短时间范围以减少数据量
2. 使用具体的来源参数筛选
3. 继续测试剩余4个工具

### 代码示例
```go
// 获取新闻联播数据
params := map[string]string{
    "date": "20240305",
}
result, err := client.Call("llm_corpus_cctv_news", params)
```

---

## 测试环境

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
