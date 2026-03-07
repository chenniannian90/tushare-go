package mcp

import (
	"context"
	"fmt"
	"reflect"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api"
)

// ToolRegistry manages all available tools
type ToolRegistry struct {
	tools map[string]Tool
}

// NewToolRegistry creates a new tool registry
func NewToolRegistry() *ToolRegistry {
	registry := &ToolRegistry{
		tools: make(map[string]Tool),
	}

	// Register all available API tools
	registry.registerAPITools()

	return registry
}

// Initialize initializes the registry (no-op for now, kept for compatibility)
func (r *ToolRegistry) Initialize(client *sdk.Client) {
	// Tools are pre-registered, but we could use client for validation
}

// registerAPITools registers all generated API functions as tools
func (r *ToolRegistry) registerAPITools() {
	// Market data APIs
	r.registerTool("stock_basic", "Get basic information about stocks including symbol, name, area, and industry")
	r.registerTool("daily", "Get daily market data including open, high, low, close, volume, and amount")
	r.registerTool("weekly", "Get weekly market data")
	r.registerTool("monthly", "Get monthly market data")
	r.registerTool("pro_bar", "Comprehensive market data interface supporting multiple timeframes and asset types")
	r.registerTool("trade_cal", "Get trading calendar data to determine market trading days")
	r.registerTool("daily_basic", "Get daily fundamental metrics including PE, PB, PS ratios and market caps")
	r.registerTool("index_basic", "Get index basic information")
	r.registerTool("index_daily", "Get index daily market data")

	// Financial data APIs
	r.registerTool("income", "Get income statement data including revenue, costs, and net profit")
	r.registerTool("balancesheet", "Get balance sheet data including total assets and liabilities")
	r.registerTool("fina_indicator", "Get financial indicators including ROE, ROA, and various ratios")

	// Other APIs
	r.registerTool("moneyflow", "Get money flow data including large and super-large order flows")
	r.registerTool("dividend", "Get dividend data including dividend per share and ex-dividend dates")
	r.registerTool("top10_holders", "Get top 10 shareholders data")
	r.registerTool("holder_number", "Get shareholder number statistics")
	r.registerTool("concept", "Get concept sector classifications")
	r.registerTool("concept_detail", "Get detailed concept sector constituent stocks")
	r.registerTool("limit_list", "Get limit up/down stock list")
}

// registerTool registers a single tool
func (r *ToolRegistry) registerTool(name, description string) {
	r.tools[name] = Tool{
		Name:        name,
		Description: description,
	}
}

// GetTools returns all registered tools
func (r *ToolRegistry) GetTools() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// CallTool executes a tool call by dispatching to the appropriate API function
func (r *ToolRegistry) CallTool(ctx context.Context, client *sdk.Client, toolName string, args map[string]interface{}) (*ToolResult, error) {
	switch toolName {
	case "stock_basic":
		return r.callStockBasic(ctx, client, args)
	case "daily":
		return r.callDaily(ctx, client, args)
	case "daily_basic":
		return r.callDailyBasic(ctx, client, args)
	case "trade_cal":
		return r.callTradeCal(ctx, client, args)
	case "index_basic":
		return r.callIndexBasic(ctx, client, args)
	case "index_daily":
		return r.callIndexDaily(ctx, client, args)
	case "concept":
		return r.callConcept(ctx, client, args)
	case "weekly":
		return r.callWeekly(ctx, client, args)
	case "monthly":
		return r.callMonthly(ctx, client, args)
	case "pro_bar":
		return r.callProBar(ctx, client, args)
	case "income":
		return r.callIncome(ctx, client, args)
	case "balancesheet":
		return r.callBalancesheet(ctx, client, args)
	case "fina_indicator":
		return r.callFinaIndicator(ctx, client, args)
	case "moneyflow":
		return r.callMoneyflow(ctx, client, args)
	case "dividend":
		return r.callDividend(ctx, client, args)
	case "top10_holders":
		return r.callTop10Holders(ctx, client, args)
	case "holder_number":
		return r.callHolderNumber(ctx, client, args)
	case "limit_list":
		return r.callLimitList(ctx, client, args)
	default:
		// Return a helpful message for unimplemented APIs
		return &ToolResult{
			Content: []Content{
				{
					Type: "text",
					Text: fmt.Sprintf("API '%s' is registered but not yet fully implemented in MCP server. Available APIs: stock_basic, daily, daily_basic, trade_cal, index_basic, index_daily, concept, weekly, monthly, pro_bar, income, balancesheet, fina_indicator, moneyflow, dividend, top10_holders, holder_number, limit_list", toolName),
				},
			},
		}, nil
	}
}

// Implementation methods for each API

// callStockBasic handles stock_basic tool calls
func (r *ToolRegistry) callStockBasic(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.StockBasicRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if listStatus, ok := args["list_status"].(string); ok {
		req.ListStatus = listStatus
	}
	if exchange, ok := args["exchange"].(string); ok {
		req.Exchange = exchange
	}

	items, err := api.StockBasic(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call stock_basic API: %w", err)
	}

	result, err := r.formatItems(items, "stock")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callDaily handles daily tool calls
func (r *ToolRegistry) callDaily(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.DailyRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if tradeDate, ok := args["trade_date"].(string); ok {
		req.TradeDate = tradeDate
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Daily(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call daily API: %w", err)
	}

	result, err := r.formatItems(items, "daily data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callDailyBasic handles daily_basic tool calls
func (r *ToolRegistry) callDailyBasic(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.DailyBasicRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if tradeDate, ok := args["trade_date"].(string); ok {
		req.TradeDate = tradeDate
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.DailyBasic(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call daily_basic API: %w", err)
	}

	result, err := r.formatItems(items, "daily basic data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callTradeCal handles trade_cal tool calls
func (r *ToolRegistry) callTradeCal(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.TradeCalRequest{}
	if exchange, ok := args["exchange"].(string); ok {
		req.Exchange = exchange
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.TradeCal(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call trade_cal API: %w", err)
	}

	result, err := r.formatItems(items, "trading calendar")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callIndexBasic handles index_basic tool calls
func (r *ToolRegistry) callIndexBasic(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.IndexBasicRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if name, ok := args["name"].(string); ok {
		req.Name = name
	}
	if market, ok := args["market"].(string); ok {
		req.Market = market
	}

	items, err := api.IndexBasic(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call index_basic API: %w", err)
	}

	result, err := r.formatItems(items, "index")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callIndexDaily handles index_daily tool calls
func (r *ToolRegistry) callIndexDaily(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.IndexDailyRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if tradeDate, ok := args["trade_date"].(string); ok {
		req.TradeDate = tradeDate
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.IndexDaily(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call index_daily API: %w", err)
	}

	result, err := r.formatItems(items, "index daily data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callConcept handles concept tool calls
func (r *ToolRegistry) callConcept(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.ConceptRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if name, ok := args["name"].(string); ok {
		req.Name = name
	}

	items, err := api.Concept(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call concept API: %w", err)
	}

	result, err := r.formatItems(items, "concept")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// formatItems formats a slice of items for display
func (r *ToolRegistry) formatItems(items interface{}, itemType string) (*ToolResult, error) {
	val := reflect.ValueOf(items)
	if val.Kind() != reflect.Slice {
		return &ToolResult{
			Content: []Content{
				{
					Type: "text",
					Text: fmt.Sprintf("%+v", items),
				},
			},
		}, nil
	}

	length := val.Len()
	text := fmt.Sprintf("Found %d %s item(s):\n", length, itemType)

	// Show first few items
	maxItems := 10
	if length < maxItems {
		maxItems = length
	}

	for i := 0; i < maxItems; i++ {
		item := val.Index(i).Interface()
		text += r.formatItem(item) + "\n"
	}

	if length > maxItems {
		text += fmt.Sprintf("... and %d more items\n", length-maxItems)
	}

	return &ToolResult{
		Content: []Content{
			{
				Type: "text",
				Text: text,
			},
		},
	}, nil
}

// formatItem formats a single item for display
func (r *ToolRegistry) formatItem(item interface{}) string {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Struct {
		var parts []string
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			value := val.Field(i)

			// Skip zero values
			if r.isZero(value) {
				continue
			}

			parts = append(parts, fmt.Sprintf("%s: %v", field.Name, value.Interface()))
		}

		return fmt.Sprintf("{%s}", joinStrings(parts, ", "))
	}

	return fmt.Sprintf("%v", item)
}

// isZero checks if a value is zero
func (r *ToolRegistry) isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// joinStrings joins a slice of strings
func joinStrings(items []string, separator string) string {
	if len(items) == 0 {
		return ""
	}
	result := items[0]
	for i := 1; i < len(items); i++ {
		result += separator + items[i]
	}
	return result
}

// callWeekly handles weekly tool calls
func (r *ToolRegistry) callWeekly(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.WeeklyRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Weekly(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call weekly API: %w", err)
	}

	result, err := r.formatItems(items, "weekly data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callMonthly handles monthly tool calls
func (r *ToolRegistry) callMonthly(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.MonthlyRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Monthly(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call monthly API: %w", err)
	}

	result, err := r.formatItems(items, "monthly data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callProBar handles pro_bar tool calls
func (r *ToolRegistry) callProBar(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.ProBarRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}
	if freq, ok := args["freq"].(string); ok {
		req.Freq = freq
	}
	if asset, ok := args["asset"].(string); ok {
		req.Asset = asset
	}
	if adj, ok := args["adj"].(string); ok {
		req.Adj = adj
	}

	items, err := api.ProBar(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call pro_bar API: %w", err)
	}

	result, err := r.formatItems(items, "market data")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callIncome handles income tool calls
func (r *ToolRegistry) callIncome(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.IncomeRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if period, ok := args["period"].(string); ok {
		req.Period = period
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Income(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call income API: %w", err)
	}

	result, err := r.formatItems(items, "income statement")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callBalancesheet handles balancesheet tool calls
func (r *ToolRegistry) callBalancesheet(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.BalancesheetRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if period, ok := args["period"].(string); ok {
		req.Period = period
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Balancesheet(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call balancesheet API: %w", err)
	}

	result, err := r.formatItems(items, "balance sheet")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callFinaIndicator handles fina_indicator tool calls
func (r *ToolRegistry) callFinaIndicator(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.FinaIndicatorRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if period, ok := args["period"].(string); ok {
		req.Period = period
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.FinaIndicator(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call fina_indicator API: %w", err)
	}

	result, err := r.formatItems(items, "financial indicators")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callMoneyflow handles moneyflow tool calls
func (r *ToolRegistry) callMoneyflow(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.MoneyflowRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if tradeDate, ok := args["trade_date"].(string); ok {
		req.TradeDate = tradeDate
	}
	if startDate, ok := args["start_date"].(string); ok {
		req.StartDate = startDate
	}
	if endDate, ok := args["end_date"].(string); ok {
		req.EndDate = endDate
	}

	items, err := api.Moneyflow(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call moneyflow API: %w", err)
	}

	result, err := r.formatItems(items, "money flow")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callDividend handles dividend tool calls
func (r *ToolRegistry) callDividend(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.DividendRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if annDate, ok := args["ann_date"].(string); ok {
		req.AnnDate = annDate
	}
	if recordDate, ok := args["record_date"].(string); ok {
		req.RecordDate = recordDate
	}
	if exDate, ok := args["ex_date"].(string); ok {
		req.ExDate = exDate
	}

	items, err := api.Dividend(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call dividend API: %w", err)
	}

	result, err := r.formatItems(items, "dividend")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callTop10Holders handles top10_holders tool calls
func (r *ToolRegistry) callTop10Holders(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.Top10HoldersRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if period, ok := args["period"].(string); ok {
		req.Period = period
	}

	items, err := api.Top10Holders(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call top10_holders API: %w", err)
	}

	result, err := r.formatItems(items, "top 10 holders")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callHolderNumber handles holder_number tool calls
func (r *ToolRegistry) callHolderNumber(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.HolderNumberRequest{}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if period, ok := args["period"].(string); ok {
		req.Period = period
	}

	items, err := api.HolderNumber(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call holder_number API: %w", err)
	}

	result, err := r.formatItems(items, "holder number")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// callLimitList handles limit_list tool calls
func (r *ToolRegistry) callLimitList(ctx context.Context, client *sdk.Client, args map[string]interface{}) (*ToolResult, error) {
	req := &api.LimitListRequest{}
	if tradeDate, ok := args["trade_date"].(string); ok {
		req.TradeDate = tradeDate
	}
	if tsCode, ok := args["ts_code"].(string); ok {
		req.TsCode = tsCode
	}
	if limitType, ok := args["limit_type"].(string); ok {
		req.LimitType = limitType
	}

	items, err := api.LimitList(ctx, client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call limit_list API: %w", err)
	}

	result, err := r.formatItems(items, "limit list")
	if err != nil {
		return nil, err
	}
	return result, nil
}
