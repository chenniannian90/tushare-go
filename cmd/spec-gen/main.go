package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Category represents a top-level category in the API directory
type Category struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Icon         string       `json:"icon"`
	DocID        int          `json:"doc_id"`
	Subcategories []Subcategory `json:"subcategories,omitempty"`
	APIs         []API        `json:"apis,omitempty"`
}

// Subcategory represents a subcategory within a category
type Subcategory struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	DocID        int          `json:"doc_id"`
	Subcategories []Subcategory `json:"subcategories,omitempty"`
	APIs         []API        `json:"apis,omitempty"`
}

// API represents an individual API endpoint
type API struct {
	Name  string `json:"name"`
	DocID int    `json:"doc_id"`
}

// APIDirectory represents the top-level structure of the API directory
type APIDirectory struct {
	Title       string     `json:"title"`
	Version     string     `json:"version"`
	LastUpdated string     `json:"last_updated"`
	BaseURL     string     `json:"doc_id_base_url"`
	Categories  []Category `json:"categories"`
}

// APISpec represents the API specification structure
type APISpec struct {
	APIName        string       `json:"api_name"`
	APICode        string       `json:"api_code,omitempty"` // Actual API code name
	Description    string       `json:"description"`
	Describe       *DescribeInfo `json:"__describe__"`
	RequestParams  []ParamField `json:"request_params"`
	ResponseFields []ParamField `json:"response_fields"`
}

// DescribeInfo contains metadata about the API
type DescribeInfo struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// ParamField represents a parameter or field definition
type ParamField struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Description string       `json:"description"`
	Properties  []ParamField `json:"properties,omitempty"`
	Items       *ParamField  `json:"items,omitempty"`
	Enum        []string     `json:"enum,omitempty"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: spec-gen <json-file> <output-dir>")
		fmt.Println("Example: spec-gen docs/api-directory.json internal/gen/specs")
		os.Exit(1)
	}

	jsonFile := os.Args[1]
	outputDir := os.Args[2]

	// Load JSON file
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading JSON file: %v\n", err)
		os.Exit(1)
	}

	var apiDir APIDirectory
	if err := json.Unmarshal(data, &apiDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Set base URL if not provided
	if apiDir.BaseURL == "" {
		apiDir.BaseURL = "https://tushare.pro/document/2?doc_id="
	}

	fmt.Printf("📁 Generating API specs from %s\n", jsonFile)
	fmt.Printf("📂 Output directory: %s\n\n", outputDir)

	// Generate specs
	count := 0
	for _, category := range apiDir.Categories {
		count += generateCategorySpecs(category, "", apiDir.BaseURL, outputDir)
	}

	fmt.Printf("\n✅ Generated %d API spec files in %s\n", count, outputDir)
}

// generateCategorySpecs generates specs for a category and its subcategories
func generateCategorySpecs(category Category, parentPath string, baseURL string, outputDir string) int {
	count := 0

	// Build current path - use format: 中文名___id
	dirName := sanitizeFileName(category.Name) + "___" + category.ID
	currentPath := dirName
	if parentPath != "" {
		currentPath = parentPath + "/" + currentPath
	}

	// Generate specs for APIs at this level
	for _, api := range category.APIs {
		if err := generateAPISpec(api, currentPath, baseURL, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating spec for %s: %v\n", api.Name, err)
		} else {
			count++
			fmt.Printf("✓ Generated spec for: %s (doc_id: %d)\n", api.Name, api.DocID)
		}
	}

	// Process subcategories
	for _, subcat := range category.Subcategories {
		count += generateSubcategorySpecs(subcat, currentPath, baseURL, outputDir)
	}

	return count
}

// generateSubcategorySpecs generates specs for a subcategory and its nested subcategories
func generateSubcategorySpecs(subcat Subcategory, parentPath string, baseURL string, outputDir string) int {
	count := 0

	// Build current path - use format: 中文名___id
	dirName := sanitizeFileName(subcat.Name) + "___" + subcat.ID
	currentPath := parentPath + "/" + dirName

	// Generate specs for APIs at this level
	for _, api := range subcat.APIs {
		if err := generateAPISpec(api, currentPath, baseURL, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating spec for %s: %v\n", api.Name, err)
		} else {
			count++
			fmt.Printf("✓ Generated spec for: %s (doc_id: %d)\n", api.Name, api.DocID)
		}
	}

	// Process nested subcategories
	for _, nestedSubcat := range subcat.Subcategories {
		count += generateSubcategorySpecs(nestedSubcat, currentPath, baseURL, outputDir)
	}

	return count
}

// generateAPISpec generates a single API spec file
func generateAPISpec(api API, categoryPath string, baseURL string, outputDir string) error {
	// Use original Chinese name as filename (cleaned)
	fileName := sanitizeFileName(api.Name)

	// Get actual API code name
	apiCode := getAPICode(api.Name, api.DocID)

	// Create output subdirectory based on category path (preserve hierarchy)
	outputSubdir := filepath.Join(outputDir, categoryPath)
	if err := os.MkdirAll(outputSubdir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create spec - use API code as api_name, Chinese name as description
	spec := APISpec{
		APIName:     apiCode,
		APICode:     apiCode, // Add api_code field for reference
		Description: api.Name,
		Describe: &DescribeInfo{
			URL:      fmt.Sprintf("%s%d", baseURL, api.DocID),
			Name:     api.Name,
			Category: getCategoryName(categoryPath),
		},
		RequestParams:  []ParamField{}, // Empty for now - will be filled manually
		ResponseFields: []ParamField{}, // Empty for now - will be filled manually
	}

	// Write spec to file with format: 中文名___api_code.json
	outputFile := filepath.Join(outputSubdir, fileName+"___"+apiCode+".json")
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal spec: %w", err)
	}

	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write spec file: %w", err)
	}

	return nil
}

// getAPICode returns the actual API code name for a given Chinese API name
func getAPICode(name string, docID int) string {
	// Mapping of Chinese API names to actual API codes
	apiCodeMap := map[string]string{
		// 基础数据
		"股票列表":             "stock_basic",
		"每日股本（盘前）":          "daily_basic",
		"交易日历":             "trade_cal",
		"ST股票列表":           "st_stock_list",
		"ST风险警示板股票":       "st_risk_warning",
		"沪深港通股票列表":         "hk_hold_stock",
		"股票曾用名":            "namechange",
		"上市公司基本信息":          "stock_company",
		"上市公司管理层":           "stk_managers",
		"管理层薪酬和持股":          "stk_rewards",
		"北交所新旧代码对照":         "bse_code_translate",
		"IPO新股上市":          "new_share",
		"股票历史列表":           "stock_history",

		// 行情数据
		"历史日线":             "daily",
		"实时日线":             "daily_now",
		"历史分钟":             "stk_mins",
		"实时分钟":             "stk_mins_now",
		"周线行情":             "weekly",
		"月线行情":             "monthly",
		"复权行情":             "adj_factor",
		"周_月线行情_每_更新":      "weekly_monthly",
		"周_月线复权行情_每_更新":    "weekly_monthly_factor",
		"复权因子":             "adj_factor",
		"实时Tick":            "stk_mins_tick",
		"实时成交":             "stk_mins_now",
		"实时排名":             "top_list",
		"每日指标":             "daily_basic",
		"通用行情接口":           "universal",
		"每日涨跌停价格":          "limit_list",
		"每日停复牌信息":          "suspend",
		"沪深股通十大成交股":        "hg_top10",
		"港股通十大成交股":        "gg_top10",
		"港股通每日成交统计":        "gg_daily",
		"港股通每月成交统计":        "gg_monthly",
		"备用行情":             "daily_basic",

		// 财务数据
		"利润表":              "income",
		"资产负债表":            "balancesheet",
		"现金流量表":            "cashflow",
		"业绩预告":             "forecast",
		"业绩快报":             "express",
		"分红送股数据":           "dividend",
		"财务指标数据":           "fina_indicator",
		"财务审计意见":           "fina_audit",
		"主营业务构成":           "fina_mainbz",
		"财报披露日期表":          "fina_disclosure",

		// 参考数据
		"前十大股东":            "top10_holders",
		"前十大流通股东":          "top10_floatholders",
		"股权质押统计数据":         "pledge_stat",
		"股权质押明细数据":         "pledge_detail",
		"股票回购":             "repurchase",
		"限售股解禁":            "unlock_share",
		"大宗交易":             "block_trade",
		"股票开户数据_停_":        "stk_accounts",
		"股票开户数据_旧_":        "stk_accounts",
		"股东人数":             "stk_holdernumber",
		"股东增减持":            "stk_change",

		// 特色数据
		"券商盈利预测数据":         "forecast_profit",
		"每日筹码及胜率":         "chip_distribution",
		"每日筹码分布":           "chip_distribution",
		"股票技术面因子_专业版〕":     "stk_factor",
		"中央结算系统持股统计":       "ccass_stat",
		"中央结算系统持股明细":       "ccass_detail",
		"沪深股通持股明细":         "hk_hold",
		"股票开盘集合竞价数据":       "auction",
		"股票收盘集合竞价数据":       "auction",
		"神奇九转指标":           "magic_9",
		"AH股比价":             "ah_lot_size",
		"机构调研数据":           "org_research",
		"券商月度金股":           "monthly_stock",

		// 两融及转融通
		"融资融券交易汇总":         "margin",
		"融资融券交易明细":         "margin_detail",
		"融资融券标的_盘前_":        "margin_target",
		"转融券交易汇总_停_":       "margin_sec",
		"转融资交易汇总":         "margin_lend",
		"转融券交易明细_停_":       "margin_sec_detail",
		"做市借券交易汇总_停_":      "margin_lend_sec",

		// 资金流向数据
		"个股资金流向":           "moneyflow",
		"板块资金流向":          "moneyflow",
		"行业资金流向":          "moneyflow",
		"大盘资金流向":           "moneyflow",
		"沪深港通资金流向":         "moneyflow_hsgt",

		// 打板专题数据
		"龙虎榜每日统计单":         "top_list",
		"龙虎榜机构交易单":         "top_inst",
		"同花顺涨跌停榜单":        "limit_list",
		"涨跌停和炸板数据":         "limit_list",
		"涨停股票连板天梯":         "limit_list_d",
		"涨停最强板块统计":         "limit_list_sec",
		"同花顺行业概念板块":        "ths_concept",
		"同花顺概念和���业指数行情":     "ths_index",
		"同花顺行业概念成分":        "ths_member",
		"东方财富概念板块":         "em_concept",
		"东方财富概念成分":         "em_concept",
		"东财概念和行业指数行情":      "em_index",
		"开盘竞价成交_当日_":      "auction_detail",
		"市场游资最全名录":         "dragon_list",
		"游资交易每日明细":         "dragon_detail",
		"同花顺App热榜数":         "ths_hot",
		"东方财富App热榜":         "em_hot",
		"通达信板块信息":           "tdx_sector",
		"通达信板块成分":           "tdx_member",
		"通达信板块行情":           "tdx_index",
		"榜单数据_开盘啦_":        "dragon_board",
		"题材成分_开盘啦_":         "dragon_theme",

		// 指数专题
		"指数基本信息":            "index_basic",
		"指数日线行情":            "index_daily",
		"指数实时日线":            "index_daily",
		"指数实时分钟":            "index_min",
		"指数周线行情":            "index_weekly",
		"指数历史分钟":            "index_min",
		"指数月线行情":            "index_monthly",
		"指数成分和权重":          "index_weight",
		"大盘指数每日指标":          "index_basic",
		"申万行业分类":            "index_classify",
		"申万行业成分_分级_":       "index_member",
		"申万行业指数日行情":        "index_daily",
		"申万实时行情":           "index_daily",
		"中信行业成分":           "index_member",
		"中信行业指数日行情":        "index_daily",
		"国际主要指数":            "index_global",
		"指数技术面因子_专业版":      "index_factor",
		"沪深市场每日交易统计":      "market",
		"深圳市场每日交易情况":      "market",

		// 公募基金
		"基金列表":             "fund_basic",
		"基金管理人":           "fund_manager",
		"基金经理":             "fund_manager",
		"基金规模":             "fund_basic",
		"基金净值":             "fund_nav",
		"基金分红":             "fund_div",
		"基金持仓":             "fund_portfolio",
		"基金技术面因子_专业版":     "fund_factor",

		// 期货数据
		"合约信息":             "fut_basic",
		"日线行情":             "fut_daily",
		"期货周_月线行情_每_更新":    "fut_weekly_monthly",
		"历史分钟行情":           "fut_min",
		"实时分钟行情":           "fut_min",
		"仓单日报":             "fut_holdings",
		"每日结算参数":           "fut_settlement",
		"历史Tick行情":          "fut_tick",
		"每日持仓排名":           "fut_holding_rank",
		"南华期货指数行情":         "fut_index",
		"期货主力与连续合约":       "fut_continuous",
		"期货主要品种交易周报":      "fut_weekly",
		"期货合约涨跌停价格":       "fut_limit",

		// 现货数据
		"上海黄金基础信息":         "spot_basic",
		"上海黄金现货日行情":         "spot_daily",

		// 期权数据
		"期权合约信息":           "opt_basic",
		"期权日线行情":           "opt_daily",
		"期权分钟行情":           "opt_min",

		// 债券专题
		"可转债基础信息":          "cb_basic",
		"可转债发行":            "cb_issue",
		"可转债赎回信息":          "cb_redemption",
		"可转债票面利率":          "cb_interest",
		"可转债行情":            "cb_daily",
		"可转债技术面因子_专业版":    "cb_factor",
		"可转债转股价变动":        "cb_call",
		"可转债转股结果":          "cb_call",
		"债券回购日行情":          "bond_repurchase",
		"柜台流通式债券报价":        "bond_oc",
		"柜台流通式债券最优报价":      "bond_oc",
		"国债收益率曲线":          "bond_zs",
		"全球财经事件":           "global_calendar",

		// 外汇数据
		"外汇基础信息_海外_":       "forex_basic",
		"外汇日线行情":           "forex_daily",

		// 港股数据
		"港股基础信息":           "hk_basic",
		"港股交易日历":           "hk_cal",
		"港股日线行情":           "hk_daily",
		"港股复权行情":           "hk_daily",
		"港股复权因子":           "hk_factor",
		"港股分钟行情":           "hk_min",
		"港股实时日线":           "hk_daily",
		"港股利润表":             "hk_income",
		"港股资产负债表":           "hk_balancesheet",
		"港股现金流量表":           "hk_cashflow",
		"港股财务指标数据":         "hk_fina_indicator",

		// 美股数据
		"美股基础信息":           "us_basic",
		"美股交易日历":           "us_cal",
		"美股日线行情":           "us_daily",
		"美股复权行情":           "us_daily",
		"美股复权因子":           "us_factor",
		"美股利润表":             "us_income",
		"美股资产负债表":           "us_balancesheet",
		"美股现金流量表":           "us_cashflow",
		"美股财务指标数据":         "us_fina_indicator",

		// 行业经济
		"台湾电子产业月营收":        "tmt_revenue",
		"台湾电子产业月营收明细":      "tmt_revenue_detail",
		"电影月度票房":           "movie_boxoffice",
		"电影周度票房":           "movie_boxoffice",
		"电影日度票房":           "movie_boxoffice",
		"影院日度票房":           "movie_boxoffice",
		"全国电影剧本备案数据":       "movie_script",
		"全国电视剧备案公示数据":     "movie_script",

		// 宏观经济
		"Shibor利率":          "shibor",
		"Shibor报价数据":        "shibor",
		"LPR贷款基础利率":        "lpr",
		"Libor利率":          "libor",
		"Hibor利率":          "hibor",
		"温州民间借贷利率":        "shibor",
		"广州民间借贷利率":        "shibor",
		"国内生产总值_GDP_":      "gdp",
		"居民消费价格指数_CPI_":     "cpi",
		"工业生产者出厂价格指数_PPI_":   "ppi",
		"货币供应量_月_":         "m2",
		"社融增量_月度_":         "shibor",
		"采购经理指数_PMI_":       "pmi",
		"国债收益率曲线利率":       "tsy",
		"国债实际收益率曲线利率":     "tsy",
		"短期国债利率":           "tsy",
		"国债长期利率":           "tsy",
		"国债长期利率平均值":       "tsy",

		// 大模型语料专题数据
		"国家政策库":            "policy",
		"券商研究报告":           "research_report",
		"新闻快讯_短讯_":         "news",
		"新闻通讯_长篇_":         "news",
		"新闻联播文字稿":          "news_broadcast",
		"上市公司公告":           "announcement",
		"上证e互动问答":          "einteraction",
		"深证易互动问答":          "einteraction",

		// 财富管理
		"各渠道公募基金销售保有规模占比": "fund_sales",
		"销售机构公募基金销售保有规模":  "fund_sales",
	}

	// Check if API code exists in mapping
	if code, ok := apiCodeMap[name]; ok {
		return code
	}

	// Fallback: generate a simple code from doc_id
	return fmt.Sprintf("api_%d", docID)
}

// sanitizeFileName cleans Chinese API name to be used as filename
func sanitizeFileName(name string) string {
	// Remove common suffixes
	name = strings.TrimSuffix(name, "（爬虫）")
	name = strings.TrimSuffix(name, "(专业版)")
	name = strings.TrimSuffix(name, "（专业版）")
	name = strings.TrimSuffix(name, "(停)")
	name = strings.TrimSuffix(name, "（停）")
	name = strings.TrimSuffix(name, "(THS)")
	name = strings.TrimSuffix(name, "（THS)")
	name = strings.TrimSuffix(name, "(DC)")
	name = strings.TrimSuffix(name, "（DC）")
	name = strings.TrimSuffix(name, "(盘前)")
	name = strings.TrimSuffix(name, "（盘前）")
	name = strings.TrimSuffix(name, "(当日)")
	name = strings.TrimSuffix(name, "（当日）")
	name = strings.TrimSuffix(name, "(每日更新)")
	name = strings.TrimSuffix(name, "（每日更新）")
	name = strings.TrimSuffix(name, "(历史)")
	name = strings.TrimSuffix(name, "（历史）")

	// Replace problematic characters with safe alternatives
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		"（", "_",
		"）", "",
		"(", "_",
		")", "",
	)
	name = replacer.Replace(name)

	return name
}

// getCategoryName extracts the category name from a category path
func getCategoryName(categoryPath string) string {
	parts := strings.Split(categoryPath, "/")
	return parts[len(parts)-1]
}
