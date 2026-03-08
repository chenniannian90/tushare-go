package main

import (
	"fmt"
	bondtools "tushare-go/pkg/mcp/tools/bond"
	etftools "tushare-go/pkg/mcp/tools/etf"
	forextools "tushare-go/pkg/mcp/tools/forex"
	fundtools "tushare-go/pkg/mcp/tools/fund"
	futurestools "tushare-go/pkg/mcp/tools/futures"
	hk_stocktools "tushare-go/pkg/mcp/tools/hk_stock"
	indextools "tushare-go/pkg/mcp/tools/index"
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

	mcpsdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"tushare-go/pkg/sdk"

	industry_tmttools "tushare-go/pkg/mcp/tools/industry_tmt"
	llm_corpustools "tushare-go/pkg/mcp/tools/llm_corpus"
	wealth_fund_salestools "tushare-go/pkg/mcp/tools/wealth_fund_sales"
)

// toolRegistrar defines the interface for tool registration
type toolRegistrar func(server *mcpsdk.Server, client *sdk.Client)

// toolRegistry maps category names to their registration functions
var toolRegistry = map[string]toolRegistrar{
	"bond":           func(s *mcpsdk.Server, c *sdk.Client) { bondtools.NewBondTools(s, c).RegisterAll() },
	"etf":            func(s *mcpsdk.Server, c *sdk.Client) { etftools.NewEtfTools(s, c).RegisterAll() },
	"forex":          func(s *mcpsdk.Server, c *sdk.Client) { forextools.NewForexTools(s, c).RegisterAll() },
	"fund":           func(s *mcpsdk.Server, c *sdk.Client) { fundtools.NewFundTools(s, c).RegisterAll() },
	"futures":        func(s *mcpsdk.Server, c *sdk.Client) { futurestools.NewFuturesTools(s, c).RegisterAll() },
	"hk_stock":       func(s *mcpsdk.Server, c *sdk.Client) { hk_stocktools.NewHk_stockTools(s, c).RegisterAll() },
	"index":          func(s *mcpsdk.Server, c *sdk.Client) { indextools.NewIndexTools(s, c).RegisterAll() },
	"industry_tmt":   func(s *mcpsdk.Server, c *sdk.Client) { industry_tmttools.NewIndustry_tmtTools(s, c).RegisterAll() },
	"llm_corpus":     func(s *mcpsdk.Server, c *sdk.Client) { llm_corpustools.NewLlm_corpusTools(s, c).RegisterAll() },
	"macro_business": func(s *mcpsdk.Server, c *sdk.Client) { macro_businesstools.NewMacro_businessTools(s, c).RegisterAll() },
	"macro_economy":  func(s *mcpsdk.Server, c *sdk.Client) { macro_economytools.NewMacro_economyTools(s, c).RegisterAll() },
	"macro_interest_rate": func(s *mcpsdk.Server, c *sdk.Client) {
		macro_interest_ratetools.NewMacro_interest_rateTools(s, c).RegisterAll()
	},
	"macro_price":   func(s *mcpsdk.Server, c *sdk.Client) { macro_pricetools.NewMacro_priceTools(s, c).RegisterAll() },
	"options":       func(s *mcpsdk.Server, c *sdk.Client) { optionstools.NewOptionsTools(s, c).RegisterAll() },
	"spot":          func(s *mcpsdk.Server, c *sdk.Client) { spottools.NewSpotTools(s, c).RegisterAll() },
	"stock_basic":   func(s *mcpsdk.Server, c *sdk.Client) { stock_basictools.NewStock_basicTools(s, c).RegisterAll() },
	"stock_board":   func(s *mcpsdk.Server, c *sdk.Client) { stock_boardtools.NewStock_boardTools(s, c).RegisterAll() },
	"stock_feature": func(s *mcpsdk.Server, c *sdk.Client) { stock_featuretools.NewStock_featureTools(s, c).RegisterAll() },
	"stock_financial": func(s *mcpsdk.Server, c *sdk.Client) {
		stock_financialtools.NewStock_financialTools(s, c).RegisterAll()
	},
	"stock_fund_flow": func(s *mcpsdk.Server, c *sdk.Client) {
		stock_fund_flowtools.NewStock_fund_flowTools(s, c).RegisterAll()
	},
	"stock_margin": func(s *mcpsdk.Server, c *sdk.Client) { stock_margintools.NewStock_marginTools(s, c).RegisterAll() },
	"stock_market": func(s *mcpsdk.Server, c *sdk.Client) { stock_markettools.NewStock_marketTools(s, c).RegisterAll() },
	"stock_reference": func(s *mcpsdk.Server, c *sdk.Client) {
		stock_referencetools.NewStock_referenceTools(s, c).RegisterAll()
	},
	"us_stock": func(s *mcpsdk.Server, c *sdk.Client) { us_stocktools.NewUs_stockTools(s, c).RegisterAll() },
	"wealth_fund_sales": func(s *mcpsdk.Server, c *sdk.Client) {
		wealth_fund_salestools.NewWealth_fund_salesTools(s, c).RegisterAll()
	},
}

// registerToolsForService registers tools for a specific service based on categories
func registerToolsForService(server *mcpsdk.Server, categories []string, client *sdk.Client) error {
	// If no categories specified, register all tools
	if len(categories) == 0 {
		return registerAllTools(server, client)
	}

	// Register tools based on categories
	for _, category := range categories {
		registrar, ok := toolRegistry[category]
		if !ok {
			return fmt.Errorf("unknown tool category: %s", category)
		}
		registrar(server, client)
	}

	return nil
}

// registerAllTools registers all available tools
func registerAllTools(server *mcpsdk.Server, client *sdk.Client) error {
	for _, registrar := range toolRegistry {
		registrar(server, client)
	}
	return nil
}
