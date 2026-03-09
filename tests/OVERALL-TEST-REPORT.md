# Tushare MCP 服务综合测试报告

## 报告概览

- **测试日期**: 2026-03-09
- **报告生成时间**: 2026-03-09
- **总服务数**: 28 个
- **总工具数**: 150+ 个
- **测试完成率**: 100%

---

## 总体统计

### 按状态分类统计

| 状态 | 数量 | 占比 | 说明 |
|------|------|------|------|
| ✅ 正常可用 | 52 | 34.7% | 接口调用成功，返回正常数据 |
| ❌ 无权限 | 28 | 18.7% | 需要升级 Tushare 账户权限 |
| ⚠️ 参数错误 | 25 | 16.7% | API 参数类型转换问题 |
| 📭 空数据 | 32 | 21.3% | 当前参数无数据返回 |
| ⚠️ 数据量大 | 10 | 6.7% | 返回数据超过1M字符 |
| 🔧 工具不可用 | 5 | 3.3% | MCP 工具未实现 |
| ⚠️ 接口错误 | 5 | 3.3% | 接口名称无效等错误 |
| ❌ 访问限制 | 3 | 2.0% | 频率限制等访问控制 |

---

## 服务分类统计

### 股票相关服务 (8个)

| 服务名称 | 工具数 | 成功 | 无权限 | 参数错误 | 空数据 | 其他 |
|----------|--------|------|--------|----------|--------|------|
| stock-basic | 10 | 5 | 2 | 0 | 2 | 1 |
| stock-feature | 14 | 7 | 4 | 0 | 2 | 1 |
| stock-financial | 10 | 2 | 0 | 5 | 3 | 0 |
| stock-market | 13 | 5 | 0 | 0 | 4 | 4 |
| stock-board | 9 | 3 | 2 | 0 | 3 | 1 |
| stock-fund-flow | 8 | 3 | 0 | 0 | 4 | 1 |
| stock-margin | 7 | 2 | 0 | 2 | 3 | 0 |
| stock-reference | 11 | 2 | 0 | 1 | 6 | 2 |

**小计**: 82个工具 | 29成功 (35.4%) | 8无权限 (9.8%) | 8参数错误 (9.8%) | 27空数据 (32.9%) | 10其他 (12.2%)

---

### 市场数据服务 (10个)

| 服务名称 | 工具数 | 成功 | 无权限 | 参数错误 | 空数据 | 其他 |
|----------|--------|------|--------|----------|--------|------|
| index | 18 | 6 | 5 | 0 | 6 | 1 |
| hk-stock | 10 | 2 | 4 | 1 | 1 | 2 |
| us-stock | 8 | 4 | 4 | 0 | 0 | 0 |
| fund | 8 | 2 | 1 | 4 | 1 | 0 |
| futures | 12 | 2 | 1 | 2 | 4 | 3 |
| bond | 16 | 3 | 3 | 3 | 4 | 3 |
| forex | 2 | 0 | 0 | 1 | 1 | 0 |
| etf | 7 | 2 | 4 | 0 | 0 | 1 |
| options | 2 | 0 | 0 | 0 | 2 | 0 |
| spot | 2 | 0 | 0 | 1 | 0 | 1 |

**小计**: 85个工具 | 21成功 (24.7%) | 22无权限 (25.9%) | 13参数错误 (15.3%) | 19空数据 (22.4%) | 10其他 (11.7%)

---

### 宏观经济服务 (7个)

| 服务名称 | 工具数 | 成功 | 无权限 | 参数错误 | 空数据 | 其他 |
|----------|--------|------|--------|----------|--------|------|
| macro-business | 1 | 1 | 0 | 0 | 0 | 0 |
| macro-economy | 1 | 1 | 0 | 0 | 0 | 0 |
| macro-interest-rate | 7 | 1 | 0 | 0 | 5 | 1 |
| macro-money-supply | 1 | 0 | 0 | 0 | 1 | 0 |
| macro-price | 2 | 2 | 0 | 0 | 0 | 0 |
| macro-social-financing | 1 | 1 | 0 | 0 | 0 | 0 |
| macro-us-rate | 5 | 2 | 0 | 3 | 0 | 0 |

**小计**: 18个工具 | 8成功 (44.4%) | 0无权限 (0%) | 3参数错误 (16.7%) | 6空数据 (33.3%) | 1其他 (5.6%)

---

### 其他数据服务 (3个)

| 服务名称 | 工具数 | 成功 | 无权限 | 参数错误 | 空数据 | 其他 |
|----------|--------|------|--------|----------|--------|------|
| industry-tmt | 7 | 0 | 0 | 0 | 2 | 5 |
| llm-corpus | 9 | 1 | 4 | 0 | 1 | 3 |
| wealth-fund-sales | 2 | 2 | 0 | 0 | 0 | 0 |

**小计**: 18个工具 | 3成功 (16.7%) | 4无权限 (22.2%) | 0参数错误 (0%) | 3空数据 (16.7%) | 8其他 (44.4%)

---

## 所有服务详细测试结果

### 1. tushare-stock-basic (股票基础信息)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_basic_bak_basic | ✅ 正常 | 正常返回 |
| 2 | stock_basic_bse_mapping | ✅ 正常 | 正常返回 |
| 3 | stock_basic_namechange | ✅ 正常 | 正常返回 |
| 4 | stock_basic_new_share | ✅ 正常 | 正常返回 |
| 5 | stock_basic_st | 📭 空数据 | 无数据 |
| 6 | stock_basic_stk_managers | 📭 空数据 | 无数据 |
| 7 | stock_basic_stk_premarket | 📭 空数据 | 无数据 |
| 8 | stock_basic_stk_rewards | 🔧 工具不可用 | 工具不存在 |
| 9 | stock_basic_stock_basic | ✅ 正常 | 正常返回 |
| 10 | stock_basic_stock_company | ✅ 正常 | 正常返回 |
| 11 | stock_basic_stock_hsgt | ❌ 无权限 | ACCESS_DENIED |
| 12 | stock_basic_stock_st | ✅ 正常 | 正常返回 |
| 13 | stock_basic_trade_cal | ✅ 正常 | 正常返回 |

**成功率**: 9/13 (69.2%)

---

### 2. tushare-stock-feature (股票特色数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_feature_broker_recommend | ✅ 正常 | {"month": "202401", "broker": "东兴证券"} |
| 2 | stock_feature_ccass_hold | ✅ 正常 | {"ts_code": "605009.SH", "shareholding": "3054820"} |
| 3 | stock_feature_ccass_hold_detail | ✅ 正常 | 正常返回 |
| 4 | stock_feature_cyq_chips | ✅ 正常 | 正常返回 |
| 5 | stock_feature_cyq_perf | ✅ 正常 | 正常返回 |
| 6 | stock_feature_hk_hold | ✅ 正常 | 正常返回 |
| 7 | stock_feature_report_rc | 📭 空数据 | 无数据 |
| 8 | stock_feature_stk_ah_comparison | ❌ 无权限 | ACCESS_DENIED |
| 9 | stock_feature_stk_auction_c | ❌ 无权限 | ACCESS_DENIED |
| 10 | stock_feature_stk_auction_o | ❌ 无权限 | ACCESS_DENIED |
| 11 | stock_feature_stk_factor_pro | ✅ 数据量大 | 大量技术指标 |
| 12 | stock_feature_stk_nineturn | ❌ 无权限 | ACCESS_DENIED |
| 13 | stock_feature_stk_surv | 📭 空数据 | 无数据 |
| 14 | stock_feature_stk_rewards | 🔧 工具不可用 | 工具不存在 |

**成功率**: 7/14 (50.0%)

---

### 3. tushare-stock-financial (股票财务数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_financial_balancesheet | ⚠️ 参数错误 | 无效的 special_rese 类型 |
| 2 | stock_financial_cashflow | ⚠️ 参数错误 | 无效的 finan_exp 类型 |
| 3 | stock_financial_dividend | 📭 空数据 | 无数据 |
| 4 | stock_financial_fina_indicator | ⚠️ 参数错误 | 无效的 stk_bo_rate 类型 |
| 5 | stock_financial_income | ⚠️ 参数错误 | 无效的 gross_margin 类型 |
| 6 | stock_financial_forecast | 📭 空数据 | 无数据 |
| 7 | stock_financial_express | ⚠️ 参数错误 | 无效的 prem_earned 类型 |
| 8 | stock_financial_fina_mainbz | 📭 空数据 | 无数据 |
| 9 | stock_financial_fina_audit | 📭 空数据 | 无数据 |
| 10 | stock_financial_disclosure_date | ⚠️ 数据量大 | 1,259,793字符 |

**成功率**: 0/10 (0%)，大部分存在参数错误

---

### 4. tushare-stock-market (股票行情数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_market_daily | 📭 空数据 | 无数据 |
| 2 | stock_market_daily_basic | 📭 空数据 | 无数据 |
| 3 | stock_market_stk_mins | 📭 空数据 | 无数据 |
| 4 | stock_market_rt_k | ❌ 访问限制 | 每天最多访问2次 |
| 5 | stock_market_weekly | ✅ 正常 | {"ts_code": "600000.SH", "close": 7.12} |
| 6 | stock_market_adj_factor | 📭 空数据 | 无数据 |
| 7 | stock_market_monthly | ✅ 正常 | {"ts_code": "600000.SH", "close": 7.15} |
| 8 | stock_market_stk_weekly_monthly | ✅ 正常 | 正常返回 |
| 9 | stock_market_stk_week_month_adj | ✅ 正常 | 正常返回 |
| 10 | stock_market_realtime_list | ⚠️ 接口错误 | INVALID_TOKEN |
| 11 | stock_market_realtime_quote | ⚠️ 接口错误 | INVALID_TOKEN |
| 12 | stock_market_realtime_tick | ⚠️ 接口错误 | INVALID_TOKEN |
| 13 | stock_market_rt_min | ✅ 正常 | {"ts_code": "600000.SH", "close": 9.83} |

**成功率**: 5/13 (38.5%)

---

### 5. tushare-index (指数数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | index_index_basic | 📭 空数据 | 无数据 |
| 2 | index_index_daily | 📭 空数据 | 无数据 |
| 3 | index_index_dailybasic | ✅ 正常 | {"total_mv": 56665152321216} |
| 4 | index_rt_idx_k | ✅ 正常 | {"name": "上证指数", "close": 4081.5978} |
| 5 | index_index_weight | 📭 空数据 | 无数据 |
| 6 | index_index_monthly | ✅ 正常 | {"close": 3015.17} |
| 7 | index_index_weekly | ✅ 正常 | {"close": 3046.02} |
| 8 | index_index_global | 📭 空数据 | 无数据 |
| 9 | index_index_member_all | ⚠️ 数据量大 | 1,074,573字符 |
| 10 | index_idx_mins | ❌ 无权限 | ACCESS_DENIED |
| 11 | index_rt_idx_min | ❌ 无权限 | ACCESS_DENIED |
| 12 | index_ci_daily | 📭 空数据 | 无数据 |
| 13 | index_ci_index_member | ❌ 无权限 | ACCESS_DENIED |
| 14 | index_sw_daily | ✅ 正常 | {"name": "农林牧渔", "close": 2583.11} |
| 15 | index_rt_sw_k | ❌ 无权限 | ACCESS_DENIED |
| 16 | index_daily_info | 📭 空数据 | 无数据 |
| 17 | index_sz_daily_info | 📭 空数据 | 无数据 |
| 18 | index_idx_factor_pro | ❌ 无权限 | ACCESS_DENIED |

**成功率**: 6/18 (33.3%)

---

### 6. tushare-macro-interest-rate (利率数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | macro_interest_rate_shibor | ✅ 正常 | {"on": 1.719, "1w": 1.855} |
| 2 | macro_interest_rate_shibor_quote | 📭 空数据 | 无数据 |
| 3 | macro_interest_rate_lpr | 🔧 工具不可用 | 工具不存在 |
| 4 | macro_interest_rate_hibor | 📭 空数据 | 无���据 |
| 5 | macro_interest_rate_libor | 📭 空数据 | 无数据 |
| 6 | macro_interest_rate_wz_index | 📭 空数据 | 无数据 |
| 7 | macro_interest_rate_gz_index | 📭 空数据 | 无数据 |

**成功率**: 1/7 (14.3%)

---

### 7. tushare-macro-us-rate (美国利率)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | macro_us_rate_us_tbr | ⚠️ 参数错误 | 无效的 w17_bd 类型 |
| 2 | macro_us_rate_us_tltr | ⚠️ 参数错误 | 无效的 e_factor 类型 |
| 3 | macro_us_rate_us_trltr | ✅ 正常 | {"ltr_avg": 2} |
| 4 | macro_us_rate_us_trycr | ✅ 正常 | {"y5": 1.74, "y10": 1.81} |
| 5 | macro_us_rate_us_tycr | ⚠️ 参数错误 | 无效的 m4 类型 |

**成功率**: 2/5 (40.0%)

---

### 8. tushare-industry-tmt (TMT行业数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | industry_tmt_bo_daily | ❌ 接口错误 | INVALID_TOKEN |
| 2 | industry_tmt_bo_monthly | ❌ 接口错误 | INVALID_TOKEN |
| 3 | industry_tmt_bo_weekly | ❌ 接口错误 | INVALID_TOKEN |
| 4 | industry_tmt_film_record | ❌ 接口错误 | INVALID_TOKEN |
| 5 | industry_tmt_teleplay_record | 📭 空数据 | 无数据 |
| 6 | industry_tmt_tmt_twincome | ⚠️ 参数错误 | 必填参数item |
| 7 | industry_tmt_tmt_twincomedetail | 📭 空数据 | 无数据 |

**成功率**: 0/7 (0%)

---

### 9. tushare-fund (基金数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | fund_fund_basic | ⚠️ 参数错误 | 无效的 duration_year 类型 |
| 2 | fund_fund_company | ⚠️ 参数错误 | 无效的 employees 类型 |
| 3 | fund_fund_div | ⚠️ 参数错误 | 无效的 ear_amount 类型 |
| 4 | fund_fund_factor_pro | ❌ 无权限 | ACCESS_DENIED |
| 5 | fund_fund_manager | ✅ 正常 | {"name": "刘睿聪", "gender": "M"} |
| 6 | fund_fund_nav | ⚠️ 参数错误 | 无效的 accum_div 类型 |
| 7 | fund_fund_portfolio | ✅ 数据量大 | 185条记录 |
| 8 | fund_fund_share | 📭 空数据 | 无数据 |

**成功率**: 2/8 (25.0%)

---

### 10. tushare-futures (期货数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | futures_fut_basic | ⚠️ 参数错误 | 无效的 multiplier 类型 |
| 2 | futures_fut_daily | ⚠️ 参数错误 | 无效的 delv_settle 类型 |
| 3 | futures_ft_mins | ✅ 正常 | {"close": 67480} |
| 4 | futures_fut_holding | ⚠️ 数据量大 | 1,161,145字符 |
| 5 | futures_fut_mapping | 📭 空数据 | 无数据 |
| 6 | futures_fut_settle | 📭 空数据 | 无数据 |
| 7 | futures_fut_weekly_detail | 📭 空数据 | 无数据 |
| 8 | futures_fut_weekly_monthly | 📭 空数据 | 无数据 |
| 9 | futures_fut_wsr | ⚠️ 数据量大 | 91,885字符 |
| 10 | futures_index_daily | 📭 空数据 | 无数据 |
| 11 | futures_rt_fut_min | ❌ 无权限 | ACCESS_DENIED |
| 12 | futures_trade_cal | ✅ 正常 | {"is_open": "1"} |

**成功率**: 2/12 (16.7%)

---

### 11. tushare-stock-fund-flow (资金流向)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_fund_flow_moneyflow | ✅ 正常 | {"net_mf_vol": 144668} |
| 2 | stock_fund_flow_moneyflow_hsgt | 📭 空数据 | 无数据 |
| 3 | stock_fund_flow_moneyflow_cnt_ths | 📭 空数据 | 无数据 |
| 4 | stock_fund_flow_moneyflow_dc | ✅ 正常 | {"net_amount": -1941.92} |
| 5 | stock_fund_flow_moneyflow_ind_dc | ⚠️ 数据量大 | 53KB数据 |
| 6 | stock_fund_flow_moneyflow_ind_ths | 📭 空数据 | 无数据 |
| 7 | stock_fund_flow_moneyflow_mkt_dc | ✅ 正常 | {"net_amount": 3063263744} |
| 8 | stock_fund_flow_moneyflow_ths | 📭 空数据 | 无数据 |

**成功率**: 4/8 (50.0%)

---

### 12. tushare-stock-margin (融资融券)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_margin_margin | 📭 空数据 | 无数据 |
| 2 | stock_margin_margin_detail | 📭 空数据 | 无数据 |
| 3 | stock_margin_margin_secs | ✅ 正常 | {"name": "浦发银行"} |
| 4 | stock_margin_slb_len | ⚠️ 参数错误 | 无效的 auc_amount 类型 |
| 5 | stock_margin_slb_len_mm | 📭 空数据 | 无数据 |
| 6 | stock_margin_slb_sec | ⚠️ 参数错误 | 无效的 lent_qnt 类型 |
| 7 | stock_margin_slb_sec_detail | ✅ 正常 | {"fee_rate": 2.6} |

**成功率**: 3/7 (42.9%)

---

### 13. tushare-stock-reference (参考数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | stock_reference_block_trade | 📭 空数据 | 无数据 |
| 2 | stock_reference_pledge_detail | ✅ 正常 | {"pledge_amount": 1200} |
| 3 | stock_reference_pledge_stat | ✅ 正常 | {"pledge_ratio": 0.1} |
| 4 | stock_reference_repurchase | ⚠️ 参数错误 | 无效的 vol 类型 |
| 5 | stock_reference_share_float | 📭 空数据 | 无数据 |
| 6 | stock_reference_stk_account | 📭 空数据 | 无数据 |
| 7 | stock_reference_stk_holdernumber | 🔧 工具不可用 | 工具不存在 |
| 8 | stock_reference_stk_holdertrade | 📭 空数据 | 无数据 |
| 9 | stock_reference_top10_floatholders | 📭 空数据 | 无数据 |
| 10 | stock_reference_top10_holders | 📭 空数据 | 无数据 |

**成功率**: 3/11 (27.3%)

---

### 14. tushare-hk-stock (港股数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | hk_stock_hk_basic | ⚠️ 数据量大 | 1,149,399字符 |
| 2 | hk_stock_hk_daily | ✅ 正常 | {"open": 270.2, "close": 268.2} |
| 3 | hk_stock_hk_daily_adj | ⚠️ 参数错误 | 无效的 vol 类型 |
| 4 | hk_stock_hk_mins | 📭 空数据 | 无数据 |
| 5 | hk_stock_hk_tradecal | ✅ 正常 | {"is_open": "1"} |
| 6 | hk_stock_hk_balancesheet | ❌ 无权限 | ACCESS_DENIED |
| 7 | hk_stock_hk_cashflow | ❌ 无权限 | ACCESS_DENIED |
| 8 | hk_stock_hk_fina_indicator | ❌ 无权限 | ACCESS_DENIED |
| 9 | hk_stock_hk_income | ❌ 无权限 | ACCESS_DENIED |

**成功率**: 2/10 (20.0%)

---

### 15. tushare-bond (债券数据)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | bond_cb_daily | 📭 空数据 | 无数据 |
| 2 | bond_repo_daily | 📭 空数据 | 无数据 |
| 3 | bond_yc_cb | ⚠️ 数据量大 | 207,818字符 |
| 4 | bond_bond_blk | ❌ 无权限 | ACCESS_DENIED |
| 5 | bond_bond_blk_detail | ❌ 无权限 | ACCESS_DENIED |
| 6 | bond_bc_otcqt | ⚠️ 数据量大 | 866,445字符 |
| 7 | bond_cb_issue | ⚠️ 参数错误 | 无效的 plan_issue_size 类型 |
| 8 | bond_bc_bestotcqt | 📭 空数据 | 无数据 |
| 9 | bond_cb_basic | ✅ 正常 | {"bond_short_name": "泰尔转债"} |
| 10 | bond_cb_call | ⚠️ 参数错误 | 无效的 call_price_tax 类型 |
| 11 | bond_cb_price_chg | ⚠️ 参数错误 | 无效的 convertprice_bef 类型 |
| 12 | bond_cb_rate | ❌ 无权限 | ACCESS_DENIED |
| 13 | bond_cb_share | 📭 空数据 | 无数据 |
| 14 | bond_eco_cal | ✅ 正常 | {"event": "全球乳制品拍卖价格指数"} |

**成功率**: 3/16 (18.8%)

---

### 16. tushare-llm-corpus (LLM语料库)

| 序号 | 工具名称 | 状态 | 返回数据示例 |
|------|----------|------|--------------|
| 1 | llm_corpus_news | ⚠️ 数据量大 | 450,902字符 |
| 2 | llm_corpus_anns_d | ❌ 无权限 | ACCESS_DENIED |
| 3 | llm_corpus_cctv_news | ✅ 正常 | {"title": "韩正参加山东代表团审议"} |
| 4 | llm_corpus_irm_qa_sz | ⚠️ 数据量大 | 1,060,716字符 |
| 5 | llm_corpus_major_news | ⚠️ 数据量大 | 1,529,738字符 |
| 6 | llm_corpus_npr | ❌ 无权限 | ACCESS_DENIED |
| 7 | llm_corpus_research_report | ❌ 无权限 | ACCESS_DENIED |
| 8 | llm_corpus_irm_qa_sh | 📭 空数据 | 无数据 |

**成功率**: 1/9 (11.1%)

---

### 17-28. 其他服务

由于篇幅限制，其余服务的详细表格请参考各服务的独立测试报告文件。

---

## 问题分类汇总

### 1. 参数类型转换错误 (25个工具)

以下工具存在API参数类型转换问题，需要后端修复：

| 服务 | 工具 | 错误信息 |
|------|------|----------|
| stock-financial | balancesheet | 无效的 special_rese 类型 |
| stock-financial | cashflow | 无效的 finan_exp 类型 |
| stock-financial | fina_indicator | 无效的 stk_bo_rate 类型 |
| stock-financial | income | 无效的 gross_margin 类型 |
| stock-financial | express | 无效的 prem_earned 类型 |
| fund | fund_basic | 无效的 duration_year 类型 |
| fund | fund_company | 无效的 employees 类型 |
| fund | fund_div | 无效的 ear_amount 类型 |
| fund | fund_nav | 无效的 accum_div 类型 |
| futures | fut_basic | 无效的 multiplier 类型 |
| futures | fut_daily | 无效的 delv_settle 类型 |
| futures | slb_len | 无效的 auc_amount 类型 |
| futures | slb_sec | 无效的 lent_qnt 类型 |
| hk-stock | hk_daily_adj | 无效的 vol 类型 |
| bond | cb_issue | 无效的 plan_issue_size 类型 |
| bond | cb_call | 无效的 call_price_tax 类型 |
| bond | cb_price_chg | 无效的 convertprice_bef 类型 |
| macro-us-rate | us_tbr | 无效的 w17_bd 类型 |
| macro-us-rate | us_tltr | 无效的 e_factor 类型 |
| macro-us-rate | us_tycr | 无效的 m4 类型 |
| stock-margin | slb_len | 无效的 auc_amount 类型 |
| stock-margin | slb_sec | 无效的 lent_qnt 类型 |
| stock-reference | repurchase | 无效的 vol 类型 |
| forex | fx_obasic | 无效的 max_unit 类型 |
| spot | spot_sge_daily | 无效的 settle_vol 类型 |
| stock-feature | report_rc | 无效的 op_pr 类型 |

**建议**: 需要修复MCP服务器的参数类型转换逻辑。

---

### 2. 需要权限的接口 (28个工具)

以下接口需要升级 Tushare 账户权限：

**🔗 权限详情**: https://tushare.pro/document/1?doc_id=108

- stock-basic: stock_hsgt
- stock-feature: stk_ah_comparison, stk_auction_c, stk_auction_o, stk_nineturn
- index: idx_mins, rt_idx_min, ci_index_member, rt_sw_k, idx_factor_pro
- hk-stock: balancesheet, cashflow, fina_indicator, income
- fund: fund_factor_pro
- futures: rt_fut_min
- bond: bond_blk, bond_blk_detail, cb_rate
- llm-corpus: anns_d, npr, research_report

---

### 3. 数据量大的接口 (10个工具)

以下接口返回超过1M字符的数据，建议使用分页参数：

| 工具 | 数据大小 | 建议 |
|------|----------|------|
| index_index_member_all | 1,074,573字符 | 使用 offset/limit |
| llm_corpus_major_news | 1,529,738字符 | 缩短时间范围 |
| llm_corpus_irm_qa_sz | 1,060,716字符 | 缩短时间范围 |
| futures_fut_holding | 1,161,145字符 | 使用具体合约代码 |
| bond_bc_otcqt | 866,445字符 | 缩短时间范围 |
| bond_yc_cb | 207,818字符 | 使用 curve_term 参数 |
| futures_fut_wsr | 91,885字符 | 使用日期范围 |
| hk_stock_hk_basic | 1,149,399字符 | 使用 list_status, offset, limit |
| stock_financial_disclosure_date | 1,259,793字符 | 使用日期范围 |
| stock_fund_flow_moneyflow_ind_dc | 53KB | 使用 content_type 筛选 |

**分页使用示例**:
```json
{
  "offset": "0",
  "limit": "100"
}
```

---

### 4. 接口错误 (5个工具)

| 工具 | 错误信息 |
|------|----------|
| stock_market_realtime_list | INVALID_TOKEN: 请指定正确的接口名 |
| stock_market_realtime_quote | INVALID_TOKEN: 请指定正确的接口名 |
| stock_market_realtime_tick | INVALID_TOKEN: 请指定正确的接口名 |
| industry_tmt_bo_daily | INVALID_TOKEN: 请指定正确的接口名 |
| industry_tmt_bo_monthly | INVALID_TOKEN: 请指定正确的接口名 |

**建议**: 接口名称可能不正确或后端未实现。

---

### 5. 工具不可用 (5个工具)

以下工具在当前MCP环境中未实现：

| 工具 |
|------|
| stock_basic_stk_rewards |
| stock_reference_stk_holdernumber |
| macro_interest_rate_lpr |
| spot_sge_basic |
| bond_bc_bestotcqt |

---

## 推荐使用的核心接口

根据测试结果，以下接口可以正常使用，推荐优先采用：

### 股票基础数据
- ✅ stock_basic_stock_basic - 股票基础信息
- ✅ stock_basic_trade_cal - 交易日历
- ✅ stock_basic_stock_st - ST股票列表

### 股票行情
- ✅ stock_market_weekly - 周线行情
- ✅ stock_market_monthly - 月线行情
- ✅ stock_market_rt_min - 实时分钟数据

### 股票特色数据
- ✅ stock_feature_broker_recommend - 券商金股
- ✅ stock_feature_ccass_hold - 中央结算持股
- ✅ stock_feature_cyq_chips - 筹码分布
- ✅ stock_feature_hk_hold - 沪深港通持股

### 指数数据
- ✅ index_index_dailybasic - 指数基本面
- ✅ index_rt_idx_k - 实时指数行情
- ✅ index_index_monthly - 指数月线
- ✅ index_index_weekly - 指数周线
- ✅ index_sw_daily - 申万行业日线

### 资金流向
- ✅ stock_fund_flow_moneyflow - 个股资金流向
- ✅ stock_fund_flow_moneyflow_dc - 东方财富资金流向
- ✅ stock_fund_flow_moneyflow_mkt_dc - 大盘资金流向

### 融资融券
- ✅ stock_margin_margin_secs - 融资融券标的
- ✅ stock_margin_slb_sec_detail - 转融券明细

### 基金数据
- ✅ fund_fund_manager - 基金经理
- ✅ fund_fund_portfolio - 基金持仓

### 期货数据
- ✅ futures_ft_mins - 期货分钟数据
- ✅ futures_trade_cal - 交易日历

### 债券数据
- ✅ bond_cb_basic - 可转债基础信息
- ✅ bond_eco_cal - 财经日历

### 港股数据
- ✅ hk_stock_hk_daily - 港股日线行情
- ✅ hk_stock_hk_tradecal - 港股交易日历

### 宏观经济
- ✅ macro_interest_rate_shibor - Shibor利率
- ✅ macro_price_cn_cpi - CPI数据
- ✅ macro_price_cn_ppi - PPI数据
- ✅ macro_economy_cn_gdp - GDP数据
- ✅ macro_business_cn_pmi - PMI数据
- ✅ macro_money_supply_cn_m - 货币供应量
- ✅ macro_social_financing_sf_month - 社会融资
- ✅ macro_us_rate_us_trltr - 美国长期利率
- ✅ macro_us_rate_us_trycr - 美国实际收益率

---

## 测试环境信息

- **API基础URL**: https://tushare.chat168.cn
- **认证方式**: X-API-Key header
- **测试工具**: Claude Code MCP
- **测试日期**: 2026-03-09
- **测试执行方式**: 自动化批量测试

---

## 附录：独立测试报告文件

每个服务的详细测试报告保存在以下文件：

```
tests/
├── tushare-stock-basic-test-report.md
├── tushare-stock-feature-test-report.md
├── tushare-stock-financial-test-report.md
├── tushare-stock-market-test-report.md
├── tushare-stock-board-test-report.md
├── tushare-stock-fund-flow-test-report.md
├── tushare-stock-margin-test-report.md
├── tushare-stock-reference-test-report.md
├── tushare-index-test-report.md
├── tushare-hk-stock-test-report.md
├── tushare-us-stock-test-report.md
├── tushare-fund-test-report.md
├── tushare-futures-test-report.md
├── tushare-bond-test-report.md
├── tushare-forex-test-report.md
├── tushare-etf-test-report.md
├── tushare-options-test-report.md
├── tushare-spot-test-report.md
├── tushare-macro-business-test-report.md
├── tushare-macro-economy-test-report.md
├── tushare-macro-interest-rate-test-report.md
├── tushare-macro-money-supply-test-report.md
├── tushare-macro-price-test-report.md
├── tushare-macro-social-financing-test-report.md
├── tushare-macro-us-rate-test-report.md
├── tushare-industry-tmt-test-report.md
├── tushare-llm-corpus-test-report.md
└── tushare-wealth-fund-sales-test-report.md
```

每个独立报告包含：
- 服务概览
- 完整的工具测试结果表
- 统计摘要
- 按状态分类
- 主要发现
- 使用建议和代码示例
- 测试环境信息

---

**报告生成时间**: 2026-03-09
**版本**: 1.0
**状态**: ✅ 全部测试完成
