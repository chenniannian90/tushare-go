package main

import (
	"fmt"

	"tushare-go/pkg/sdk"
	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"

	// Import all tool modules
	bondtools "tushare-go/pkg/mcp/tools/bond"
	etftools "tushare-go/pkg/mcp/tools/etf"
	forextools "tushare-go/pkg/mcp/tools/forex"
	fundtools "tushare-go/pkg/mcp/tools/fund"
	futurestools "tushare-go/pkg/mcp/tools/futures"
	hk_stocktools "tushare-go/pkg/mcp/tools/hk_stock"
	indextools "tushare-go/pkg/mcp/tools/index"
	industry_tmttools "tushare-go/pkg/mcp/tools/industry_tmt"
	llm_corpustools "tushare-go/pkg/mcp/tools/llm_corpus"
	macro_businesstools "tushare-go/pkg/mcp/tools/macro_business"
	macro_economytools "tushare-go/pkg/mcp/tools/macro_economy"
	macro_interest_ratetools "tushare-go/pkg/mcp/tools/macro_interest_rate"
	macro_pricetools "tushare-go/pkg/mcp/tools/macro_price"
	optionstools "tushare-go/pkg/mcp/tools/options"
	spottools "tushare-go/pkg/mcp/tools/spot"
	stock_basictools "tushare-go/pkg/mcp/tools/stock_basic"
	stock_boardtools "tushare-go/pkg/mcp/tools/stock_board"
	stock_featuretools "tushare-go/pkg/mcp/tools/stock_feature"
	stock_financialtools "tushare-go/pkg/mcp/tools/stock_financial"
	stock_fund_flowtools "tushare-go/pkg/mcp/tools/stock_fund_flow"
	stock_margintools "tushare-go/pkg/mcp/tools/stock_margin"
	stock_markettools "tushare-go/pkg/mcp/tools/stock_market"
	stock_referencetools "tushare-go/pkg/mcp/tools/stock_reference"
	us_stocktools "tushare-go/pkg/mcp/tools/us_stock"
	wealth_fund_salestools "tushare-go/pkg/mcp/tools/wealth_fund_sales"
)

// registerToolsForService registers tools for a specific service based on categories
func registerToolsForService(server *mcpsdk.Server, categories []string, client *sdk.Client) error {
	// If no categories specified, register all tools
	if len(categories) == 0 {
		return registerAllTools(server, client)
	}

	// Register tools based on categories
	for _, category := range categories {
		if err := registerToolCategory(server, category, client); err != nil {
			return fmt.Errorf("failed to register %s tools: %w", category, err)
		}
	}

	return nil
}

// registerToolCategory registers all tools for a specific category
func registerToolCategory(server *mcpsdk.Server, category string, client *sdk.Client) error {
	switch category {
	case "bond":
		tools := bondtools.NewBondTools(server, client)
		tools.RegisterAll()
	case "etf":
		tools := etftools.NewEtfTools(server, client)
		tools.RegisterAll()
	case "forex":
		tools := forextools.NewForexTools(server, client)
		tools.RegisterAll()
	case "fund":
		tools := fundtools.NewFundTools(server, client)
		tools.RegisterAll()
	case "futures":
		tools := futurestools.NewFuturesTools(server, client)
		tools.RegisterAll()
	case "hk_stock":
		tools := hk_stocktools.NewHk_stockTools(server, client)
		tools.RegisterAll()
	case "index":
		tools := indextools.NewIndexTools(server, client)
		tools.RegisterAll()
	case "industry_tmt":
		tools := industry_tmttools.NewIndustry_tmtTools(server, client)
		tools.RegisterAll()
	case "llm_corpus":
		tools := llm_corpustools.NewLlm_corpusTools(server, client)
		tools.RegisterAll()
	case "macro_business":
		tools := macro_businesstools.NewMacro_businessTools(server, client)
		tools.RegisterAll()
	case "macro_economy":
		tools := macro_economytools.NewMacro_economyTools(server, client)
		tools.RegisterAll()
	case "macro_interest_rate":
		tools := macro_interest_ratetools.NewMacro_interest_rateTools(server, client)
		tools.RegisterAll()
	case "macro_price":
		tools := macro_pricetools.NewMacro_priceTools(server, client)
		tools.RegisterAll()
	case "options":
		tools := optionstools.NewOptionsTools(server, client)
		tools.RegisterAll()
	case "spot":
		tools := spottools.NewSpotTools(server, client)
		tools.RegisterAll()
	case "stock_basic":
		tools := stock_basictools.NewStock_basicTools(server, client)
		tools.RegisterAll()
	case "stock_board":
		tools := stock_boardtools.NewStock_boardTools(server, client)
		tools.RegisterAll()
	case "stock_feature":
		tools := stock_featuretools.NewStock_featureTools(server, client)
		tools.RegisterAll()
	case "stock_financial":
		tools := stock_financialtools.NewStock_financialTools(server, client)
		tools.RegisterAll()
	case "stock_fund_flow":
		tools := stock_fund_flowtools.NewStock_fund_flowTools(server, client)
		tools.RegisterAll()
	case "stock_margin":
		tools := stock_margintools.NewStock_marginTools(server, client)
		tools.RegisterAll()
	case "stock_market":
		tools := stock_markettools.NewStock_marketTools(server, client)
		tools.RegisterAll()
	case "stock_reference":
		tools := stock_referencetools.NewStock_referenceTools(server, client)
		tools.RegisterAll()
	case "us_stock":
		tools := us_stocktools.NewUs_stockTools(server, client)
		tools.RegisterAll()
	case "wealth_fund_sales":
		tools := wealth_fund_salestools.NewWealth_fund_salesTools(server, client)
		tools.RegisterAll()
	default:
		return fmt.Errorf("unknown category: %s", category)
	}

	return nil
}

// registerAllTools registers all available tools
func registerAllTools(server *mcpsdk.Server, client *sdk.Client) error {
	// Register all tool modules
	bondtools.NewBondTools(server, client).RegisterAll()
	etftools.NewEtfTools(server, client).RegisterAll()
	forextools.NewForexTools(server, client).RegisterAll()
	fundtools.NewFundTools(server, client).RegisterAll()
	futurestools.NewFuturesTools(server, client).RegisterAll()
	hk_stocktools.NewHk_stockTools(server, client).RegisterAll()
	indextools.NewIndexTools(server, client).RegisterAll()
	industry_tmttools.NewIndustry_tmtTools(server, client).RegisterAll()
	llm_corpustools.NewLlm_corpusTools(server, client).RegisterAll()
	macro_businesstools.NewMacro_businessTools(server, client).RegisterAll()
	macro_economytools.NewMacro_economyTools(server, client).RegisterAll()
	macro_interest_ratetools.NewMacro_interest_rateTools(server, client).RegisterAll()
	macro_pricetools.NewMacro_priceTools(server, client).RegisterAll()
	optionstools.NewOptionsTools(server, client).RegisterAll()
	spottools.NewSpotTools(server, client).RegisterAll()
	stock_basictools.NewStock_basicTools(server, client).RegisterAll()
	stock_boardtools.NewStock_boardTools(server, client).RegisterAll()
	stock_featuretools.NewStock_featureTools(server, client).RegisterAll()
	stock_financialtools.NewStock_financialTools(server, client).RegisterAll()
	stock_fund_flowtools.NewStock_fund_flowTools(server, client).RegisterAll()
	stock_margintools.NewStock_marginTools(server, client).RegisterAll()
	stock_markettools.NewStock_marketTools(server, client).RegisterAll()
	stock_referencetools.NewStock_referenceTools(server, client).RegisterAll()
	us_stocktools.NewUs_stockTools(server, client).RegisterAll()
	wealth_fund_salestools.NewWealth_fund_salesTools(server, client).RegisterAll()

	return nil
}
