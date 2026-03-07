package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_basic"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_board"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_feature"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_financial"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_fund_flow"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_margin"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_market"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/stock/stock_reference"
)

type Stock interface {
	StockBasic
	StockBoard
	StockFeature
	StockFinancial
	StockFundFlow
	StockMargin
	StockMarket
	StockReference
}

type StockBasic interface {
	BseCodeTranslate(ctx context.Context, req *stock_basic.BseCodeTranslateRequest) ([]stock_basic.BseCodeTranslateItem, error)
	BasicDailyBasic(ctx context.Context, req *stock_basic.DailyBasicRequest) ([]stock_basic.DailyBasicItem, error)
	HkHoldStock(ctx context.Context, req *stock_basic.HkHoldStockRequest) ([]stock_basic.HkHoldStockItem, error)
	Namechange(ctx context.Context, req *stock_basic.NamechangeRequest) ([]stock_basic.NamechangeItem, error)
	NewShare(ctx context.Context, req *stock_basic.NewShareRequest) ([]stock_basic.NewShareItem, error)
	StRiskWarning(ctx context.Context, req *stock_basic.StRiskWarningRequest) ([]stock_basic.StRiskWarningItem, error)
	StStockList(ctx context.Context, req *stock_basic.StStockListRequest) ([]stock_basic.StStockListItem, error)
	StkManagers(ctx context.Context, req *stock_basic.StkManagersRequest) ([]stock_basic.StkManagersItem, error)
	StkRewards(ctx context.Context, req *stock_basic.StkRewardsRequest) ([]stock_basic.StkRewardsItem, error)
	StockBasic(ctx context.Context, req *stock_basic.StockBasicRequest) ([]stock_basic.StockBasicItem, error)
	StockCompany(ctx context.Context, req *stock_basic.StockCompanyRequest) ([]stock_basic.StockCompanyItem, error)
	StockHistory(ctx context.Context, req *stock_basic.StockHistoryRequest) ([]stock_basic.StockHistoryItem, error)
	TradeCal(ctx context.Context, req *stock_basic.TradeCalRequest) ([]stock_basic.TradeCalItem, error)
}

type StockBoard interface {
	Api260(ctx context.Context, req *stock_board.Api260Request) ([]stock_board.Api260Item, error)
	Api347(ctx context.Context, req *stock_board.Api347Request) ([]stock_board.Api347Item, error)
	Api351(ctx context.Context, req *stock_board.Api351Request) ([]stock_board.Api351Item, error)
	Api369(ctx context.Context, req *stock_board.Api369Request) ([]stock_board.Api369Item, error)
	DragonDetail(ctx context.Context, req *stock_board.DragonDetailRequest) ([]stock_board.DragonDetailItem, error)
	DragonList(ctx context.Context, req *stock_board.DragonListRequest) ([]stock_board.DragonListItem, error)
	EmConcept(ctx context.Context, req *stock_board.EmConceptRequest) ([]stock_board.EmConceptItem, error)
	EmHot(ctx context.Context, req *stock_board.EmHotRequest) ([]stock_board.EmHotItem, error)
	EmIndex(ctx context.Context, req *stock_board.EmIndexRequest) ([]stock_board.EmIndexItem, error)
	BoardLimitList(ctx context.Context, req *stock_board.LimitListRequest) ([]stock_board.LimitListItem, error)
	LimitListD(ctx context.Context, req *stock_board.LimitListDRequest) ([]stock_board.LimitListDItem, error)
	LimitListSec(ctx context.Context, req *stock_board.LimitListSecRequest) ([]stock_board.LimitListSecItem, error)
	TdxIndex(ctx context.Context, req *stock_board.TdxIndexRequest) ([]stock_board.TdxIndexItem, error)
	TdxMember(ctx context.Context, req *stock_board.TdxMemberRequest) ([]stock_board.TdxMemberItem, error)
	TdxSector(ctx context.Context, req *stock_board.TdxSectorRequest) ([]stock_board.TdxSectorItem, error)
	ThsConcept(ctx context.Context, req *stock_board.ThsConceptRequest) ([]stock_board.ThsConceptItem, error)
	ThsHot(ctx context.Context, req *stock_board.ThsHotRequest) ([]stock_board.ThsHotItem, error)
	ThsMember(ctx context.Context, req *stock_board.ThsMemberRequest) ([]stock_board.ThsMemberItem, error)
	TopInst(ctx context.Context, req *stock_board.TopInstRequest) ([]stock_board.TopInstItem, error)
	TopList(ctx context.Context, req *stock_board.TopListRequest) ([]stock_board.TopListItem, error)
}

type StockFeature interface {
	AhLotSize(ctx context.Context, req *stock_feature.AhLotSizeRequest) ([]stock_feature.AhLotSizeItem, error)
	Api328(ctx context.Context, req *stock_feature.Api328Request) ([]stock_feature.Api328Item, error)
	Auction(ctx context.Context, req *stock_feature.AuctionRequest) ([]stock_feature.AuctionItem, error)
	CcassDetail(ctx context.Context, req *stock_feature.CcassDetailRequest) ([]stock_feature.CcassDetailItem, error)
	CcassStat(ctx context.Context, req *stock_feature.CcassStatRequest) ([]stock_feature.CcassStatItem, error)
	ChipDistribution(ctx context.Context, req *stock_feature.ChipDistributionRequest) ([]stock_feature.ChipDistributionItem, error)
	ForecastProfit(ctx context.Context, req *stock_feature.ForecastProfitRequest) ([]stock_feature.ForecastProfitItem, error)
	HkHold(ctx context.Context, req *stock_feature.HkHoldRequest) ([]stock_feature.HkHoldItem, error)
	Magic9(ctx context.Context, req *stock_feature.Magic9Request) ([]stock_feature.Magic9Item, error)
	MonthlyStock(ctx context.Context, req *stock_feature.MonthlyStockRequest) ([]stock_feature.MonthlyStockItem, error)
	OrgResearch(ctx context.Context, req *stock_feature.OrgResearchRequest) ([]stock_feature.OrgResearchItem, error)
}

type StockFinancial interface {
	Balancesheet(ctx context.Context, req *stock_financial.BalancesheetRequest) ([]stock_financial.BalancesheetItem, error)
	Cashflow(ctx context.Context, req *stock_financial.CashflowRequest) ([]stock_financial.CashflowItem, error)
	Dividend(ctx context.Context, req *stock_financial.DividendRequest) ([]stock_financial.DividendItem, error)
	Express(ctx context.Context, req *stock_financial.ExpressRequest) ([]stock_financial.ExpressItem, error)
	FinaAudit(ctx context.Context, req *stock_financial.FinaAuditRequest) ([]stock_financial.FinaAuditItem, error)
	FinaDisclosure(ctx context.Context, req *stock_financial.FinaDisclosureRequest) ([]stock_financial.FinaDisclosureItem, error)
	FinaIndicator(ctx context.Context, req *stock_financial.FinaIndicatorRequest) ([]stock_financial.FinaIndicatorItem, error)
	FinaMainbz(ctx context.Context, req *stock_financial.FinaMainbzRequest) ([]stock_financial.FinaMainbzItem, error)
	Forecast(ctx context.Context, req *stock_financial.ForecastRequest) ([]stock_financial.ForecastItem, error)
	Income(ctx context.Context, req *stock_financial.IncomeRequest) ([]stock_financial.IncomeItem, error)
}

type StockFundFlow interface {
	Api343(ctx context.Context, req *stock_fund_flow.Api343Request) ([]stock_fund_flow.Api343Item, error)
	Api344(ctx context.Context, req *stock_fund_flow.Api344Request) ([]stock_fund_flow.Api344Item, error)
	Api345(ctx context.Context, req *stock_fund_flow.Api345Request) ([]stock_fund_flow.Api345Item, error)
	Api348(ctx context.Context, req *stock_fund_flow.Api348Request) ([]stock_fund_flow.Api348Item, error)
	Api349(ctx context.Context, req *stock_fund_flow.Api349Request) ([]stock_fund_flow.Api349Item, error)
	Api371(ctx context.Context, req *stock_fund_flow.Api371Request) ([]stock_fund_flow.Api371Item, error)
	Moneyflow(ctx context.Context, req *stock_fund_flow.MoneyflowRequest) ([]stock_fund_flow.MoneyflowItem, error)
	MoneyflowHsgt(ctx context.Context, req *stock_fund_flow.MoneyflowHsgtRequest) ([]stock_fund_flow.MoneyflowHsgtItem, error)
}

type StockMargin interface {
	Api326(ctx context.Context, req *stock_margin.Api326Request) ([]stock_margin.Api326Item, error)
	Api332(ctx context.Context, req *stock_margin.Api332Request) ([]stock_margin.Api332Item, error)
	Api333(ctx context.Context, req *stock_margin.Api333Request) ([]stock_margin.Api333Item, error)
	Api334(ctx context.Context, req *stock_margin.Api334Request) ([]stock_margin.Api334Item, error)
	Margin(ctx context.Context, req *stock_margin.MarginRequest) ([]stock_margin.MarginItem, error)
	MarginDetail(ctx context.Context, req *stock_margin.MarginDetailRequest) ([]stock_margin.MarginDetailItem, error)
	MarginLend(ctx context.Context, req *stock_margin.MarginLendRequest) ([]stock_margin.MarginLendItem, error)
}

type StockMarket interface {
	AdjFactor(ctx context.Context, req *stock_market.AdjFactorRequest) ([]stock_market.AdjFactorItem, error)
	Api315(ctx context.Context, req *stock_market.Api315Request) ([]stock_market.Api315Item, error)
	Api316(ctx context.Context, req *stock_market.Api316Request) ([]stock_market.Api316Item, error)
	Api317(ctx context.Context, req *stock_market.Api317Request) ([]stock_market.Api317Item, error)
	Api336(ctx context.Context, req *stock_market.Api336Request) ([]stock_market.Api336Item, error)
	Api365(ctx context.Context, req *stock_market.Api365Request) ([]stock_market.Api365Item, error)
	Daily(ctx context.Context, req *stock_market.DailyRequest) ([]stock_market.DailyItem, error)
	MarketDailyBasic(ctx context.Context, req *stock_market.DailyBasicRequest) ([]stock_market.DailyBasicItem, error)
	DailyNow(ctx context.Context, req *stock_market.DailyNowRequest) ([]stock_market.DailyNowItem, error)
	GgDaily(ctx context.Context, req *stock_market.GgDailyRequest) ([]stock_market.GgDailyItem, error)
	GgMonthly(ctx context.Context, req *stock_market.GgMonthlyRequest) ([]stock_market.GgMonthlyItem, error)
	GgTop10(ctx context.Context, req *stock_market.GgTop10Request) ([]stock_market.GgTop10Item, error)
	HgTop10(ctx context.Context, req *stock_market.HgTop10Request) ([]stock_market.HgTop10Item, error)
	MarketLimitList(ctx context.Context, req *stock_market.LimitListRequest) ([]stock_market.LimitListItem, error)
	Monthly(ctx context.Context, req *stock_market.MonthlyRequest) ([]stock_market.MonthlyItem, error)
	StkMins(ctx context.Context, req *stock_market.StkMinsRequest) ([]stock_market.StkMinsItem, error)
	StkMinsNow(ctx context.Context, req *stock_market.StkMinsNowRequest) ([]stock_market.StkMinsNowItem, error)
	Suspend(ctx context.Context, req *stock_market.SuspendRequest) ([]stock_market.SuspendItem, error)
	Universal(ctx context.Context, req *stock_market.UniversalRequest) ([]stock_market.UniversalItem, error)
	Weekly(ctx context.Context, req *stock_market.WeeklyRequest) ([]stock_market.WeeklyItem, error)
}

type StockReference interface {
	Api164(ctx context.Context, req *stock_reference.Api164Request) ([]stock_reference.Api164Item, error)
	Api165(ctx context.Context, req *stock_reference.Api165Request) ([]stock_reference.Api165Item, error)
	BlockTrade(ctx context.Context, req *stock_reference.BlockTradeRequest) ([]stock_reference.BlockTradeItem, error)
	PledgeDetail(ctx context.Context, req *stock_reference.PledgeDetailRequest) ([]stock_reference.PledgeDetailItem, error)
	PledgeStat(ctx context.Context, req *stock_reference.PledgeStatRequest) ([]stock_reference.PledgeStatItem, error)
	Repurchase(ctx context.Context, req *stock_reference.RepurchaseRequest) ([]stock_reference.RepurchaseItem, error)
	StkChange(ctx context.Context, req *stock_reference.StkChangeRequest) ([]stock_reference.StkChangeItem, error)
	StkHoldernumber(ctx context.Context, req *stock_reference.StkHoldernumberRequest) ([]stock_reference.StkHoldernumberItem, error)
	Top10Floatholders(ctx context.Context, req *stock_reference.Top10FloatholdersRequest) ([]stock_reference.Top10FloatholdersItem, error)
	Top10Holders(ctx context.Context, req *stock_reference.Top10HoldersRequest) ([]stock_reference.Top10HoldersItem, error)
	UnlockShare(ctx context.Context, req *stock_reference.UnlockShareRequest) ([]stock_reference.UnlockShareItem, error)
}

type stockImpl struct {
	client           *sdk.Client
	stockBasic       StockBasic
	stockBoard       StockBoard
	stockFeature     StockFeature
	stockFinancial   StockFinancial
	stockFundFlow    StockFundFlow
	stockMargin      StockMargin
	stockMarket      StockMarket
	stockReference   StockReference
}

type stockBasicImpl struct {
	client *sdk.Client
}

func (impl *stockBasicImpl) BseCodeTranslate(ctx context.Context, req *stock_basic.BseCodeTranslateRequest) ([]stock_basic.BseCodeTranslateItem, error) {
	return stock_basic.BseCodeTranslate(ctx, impl.client, req)
}

func (impl *stockBasicImpl) BasicDailyBasic(ctx context.Context, req *stock_basic.DailyBasicRequest) ([]stock_basic.DailyBasicItem, error) {
	return stock_basic.DailyBasic(ctx, impl.client, req)
}

func (impl *stockBasicImpl) HkHoldStock(ctx context.Context, req *stock_basic.HkHoldStockRequest) ([]stock_basic.HkHoldStockItem, error) {
	return stock_basic.HkHoldStock(ctx, impl.client, req)
}

func (impl *stockBasicImpl) Namechange(ctx context.Context, req *stock_basic.NamechangeRequest) ([]stock_basic.NamechangeItem, error) {
	return stock_basic.Namechange(ctx, impl.client, req)
}

func (impl *stockBasicImpl) NewShare(ctx context.Context, req *stock_basic.NewShareRequest) ([]stock_basic.NewShareItem, error) {
	return stock_basic.NewShare(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StRiskWarning(ctx context.Context, req *stock_basic.StRiskWarningRequest) ([]stock_basic.StRiskWarningItem, error) {
	return stock_basic.StRiskWarning(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StStockList(ctx context.Context, req *stock_basic.StStockListRequest) ([]stock_basic.StStockListItem, error) {
	return stock_basic.StStockList(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StkManagers(ctx context.Context, req *stock_basic.StkManagersRequest) ([]stock_basic.StkManagersItem, error) {
	return stock_basic.StkManagers(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StkRewards(ctx context.Context, req *stock_basic.StkRewardsRequest) ([]stock_basic.StkRewardsItem, error) {
	return stock_basic.StkRewards(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StockBasic(ctx context.Context, req *stock_basic.StockBasicRequest) ([]stock_basic.StockBasicItem, error) {
	return stock_basic.StockBasic(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StockCompany(ctx context.Context, req *stock_basic.StockCompanyRequest) ([]stock_basic.StockCompanyItem, error) {
	return stock_basic.StockCompany(ctx, impl.client, req)
}

func (impl *stockBasicImpl) StockHistory(ctx context.Context, req *stock_basic.StockHistoryRequest) ([]stock_basic.StockHistoryItem, error) {
	return stock_basic.StockHistory(ctx, impl.client, req)
}

func (impl *stockBasicImpl) TradeCal(ctx context.Context, req *stock_basic.TradeCalRequest) ([]stock_basic.TradeCalItem, error) {
	return stock_basic.TradeCal(ctx, impl.client, req)
}

type stockBoardImpl struct {
	client *sdk.Client
}

func (impl *stockBoardImpl) Api260(ctx context.Context, req *stock_board.Api260Request) ([]stock_board.Api260Item, error) {
	return stock_board.Api260(ctx, impl.client, req)
}

func (impl *stockBoardImpl) Api347(ctx context.Context, req *stock_board.Api347Request) ([]stock_board.Api347Item, error) {
	return stock_board.Api347(ctx, impl.client, req)
}

func (impl *stockBoardImpl) Api351(ctx context.Context, req *stock_board.Api351Request) ([]stock_board.Api351Item, error) {
	return stock_board.Api351(ctx, impl.client, req)
}

func (impl *stockBoardImpl) Api369(ctx context.Context, req *stock_board.Api369Request) ([]stock_board.Api369Item, error) {
	return stock_board.Api369(ctx, impl.client, req)
}

func (impl *stockBoardImpl) DragonDetail(ctx context.Context, req *stock_board.DragonDetailRequest) ([]stock_board.DragonDetailItem, error) {
	return stock_board.DragonDetail(ctx, impl.client, req)
}

func (impl *stockBoardImpl) DragonList(ctx context.Context, req *stock_board.DragonListRequest) ([]stock_board.DragonListItem, error) {
	return stock_board.DragonList(ctx, impl.client, req)
}

func (impl *stockBoardImpl) EmConcept(ctx context.Context, req *stock_board.EmConceptRequest) ([]stock_board.EmConceptItem, error) {
	return stock_board.EmConcept(ctx, impl.client, req)
}

func (impl *stockBoardImpl) EmHot(ctx context.Context, req *stock_board.EmHotRequest) ([]stock_board.EmHotItem, error) {
	return stock_board.EmHot(ctx, impl.client, req)
}

func (impl *stockBoardImpl) EmIndex(ctx context.Context, req *stock_board.EmIndexRequest) ([]stock_board.EmIndexItem, error) {
	return stock_board.EmIndex(ctx, impl.client, req)
}

func (impl *stockBoardImpl) BoardLimitList(ctx context.Context, req *stock_board.LimitListRequest) ([]stock_board.LimitListItem, error) {
	return stock_board.LimitList(ctx, impl.client, req)
}

func (impl *stockBoardImpl) LimitListD(ctx context.Context, req *stock_board.LimitListDRequest) ([]stock_board.LimitListDItem, error) {
	return stock_board.LimitListD(ctx, impl.client, req)
}

func (impl *stockBoardImpl) LimitListSec(ctx context.Context, req *stock_board.LimitListSecRequest) ([]stock_board.LimitListSecItem, error) {
	return stock_board.LimitListSec(ctx, impl.client, req)
}

func (impl *stockBoardImpl) TdxIndex(ctx context.Context, req *stock_board.TdxIndexRequest) ([]stock_board.TdxIndexItem, error) {
	return stock_board.TdxIndex(ctx, impl.client, req)
}

func (impl *stockBoardImpl) TdxMember(ctx context.Context, req *stock_board.TdxMemberRequest) ([]stock_board.TdxMemberItem, error) {
	return stock_board.TdxMember(ctx, impl.client, req)
}

func (impl *stockBoardImpl) TdxSector(ctx context.Context, req *stock_board.TdxSectorRequest) ([]stock_board.TdxSectorItem, error) {
	return stock_board.TdxSector(ctx, impl.client, req)
}

func (impl *stockBoardImpl) ThsConcept(ctx context.Context, req *stock_board.ThsConceptRequest) ([]stock_board.ThsConceptItem, error) {
	return stock_board.ThsConcept(ctx, impl.client, req)
}

func (impl *stockBoardImpl) ThsHot(ctx context.Context, req *stock_board.ThsHotRequest) ([]stock_board.ThsHotItem, error) {
	return stock_board.ThsHot(ctx, impl.client, req)
}

func (impl *stockBoardImpl) ThsMember(ctx context.Context, req *stock_board.ThsMemberRequest) ([]stock_board.ThsMemberItem, error) {
	return stock_board.ThsMember(ctx, impl.client, req)
}

func (impl *stockBoardImpl) TopInst(ctx context.Context, req *stock_board.TopInstRequest) ([]stock_board.TopInstItem, error) {
	return stock_board.TopInst(ctx, impl.client, req)
}

func (impl *stockBoardImpl) TopList(ctx context.Context, req *stock_board.TopListRequest) ([]stock_board.TopListItem, error) {
	return stock_board.TopList(ctx, impl.client, req)
}

type stockFeatureImpl struct {
	client *sdk.Client
}

func (impl *stockFeatureImpl) AhLotSize(ctx context.Context, req *stock_feature.AhLotSizeRequest) ([]stock_feature.AhLotSizeItem, error) {
	return stock_feature.AhLotSize(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) Api328(ctx context.Context, req *stock_feature.Api328Request) ([]stock_feature.Api328Item, error) {
	return stock_feature.Api328(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) Auction(ctx context.Context, req *stock_feature.AuctionRequest) ([]stock_feature.AuctionItem, error) {
	return stock_feature.Auction(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) CcassDetail(ctx context.Context, req *stock_feature.CcassDetailRequest) ([]stock_feature.CcassDetailItem, error) {
	return stock_feature.CcassDetail(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) CcassStat(ctx context.Context, req *stock_feature.CcassStatRequest) ([]stock_feature.CcassStatItem, error) {
	return stock_feature.CcassStat(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) ChipDistribution(ctx context.Context, req *stock_feature.ChipDistributionRequest) ([]stock_feature.ChipDistributionItem, error) {
	return stock_feature.ChipDistribution(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) ForecastProfit(ctx context.Context, req *stock_feature.ForecastProfitRequest) ([]stock_feature.ForecastProfitItem, error) {
	return stock_feature.ForecastProfit(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) HkHold(ctx context.Context, req *stock_feature.HkHoldRequest) ([]stock_feature.HkHoldItem, error) {
	return stock_feature.HkHold(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) Magic9(ctx context.Context, req *stock_feature.Magic9Request) ([]stock_feature.Magic9Item, error) {
	return stock_feature.Magic9(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) MonthlyStock(ctx context.Context, req *stock_feature.MonthlyStockRequest) ([]stock_feature.MonthlyStockItem, error) {
	return stock_feature.MonthlyStock(ctx, impl.client, req)
}

func (impl *stockFeatureImpl) OrgResearch(ctx context.Context, req *stock_feature.OrgResearchRequest) ([]stock_feature.OrgResearchItem, error) {
	return stock_feature.OrgResearch(ctx, impl.client, req)
}

type stockFinancialImpl struct {
	client *sdk.Client
}

func (impl *stockFinancialImpl) Balancesheet(ctx context.Context, req *stock_financial.BalancesheetRequest) ([]stock_financial.BalancesheetItem, error) {
	return stock_financial.Balancesheet(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) Cashflow(ctx context.Context, req *stock_financial.CashflowRequest) ([]stock_financial.CashflowItem, error) {
	return stock_financial.Cashflow(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) Dividend(ctx context.Context, req *stock_financial.DividendRequest) ([]stock_financial.DividendItem, error) {
	return stock_financial.Dividend(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) Express(ctx context.Context, req *stock_financial.ExpressRequest) ([]stock_financial.ExpressItem, error) {
	return stock_financial.Express(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) FinaAudit(ctx context.Context, req *stock_financial.FinaAuditRequest) ([]stock_financial.FinaAuditItem, error) {
	return stock_financial.FinaAudit(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) FinaDisclosure(ctx context.Context, req *stock_financial.FinaDisclosureRequest) ([]stock_financial.FinaDisclosureItem, error) {
	return stock_financial.FinaDisclosure(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) FinaIndicator(ctx context.Context, req *stock_financial.FinaIndicatorRequest) ([]stock_financial.FinaIndicatorItem, error) {
	return stock_financial.FinaIndicator(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) FinaMainbz(ctx context.Context, req *stock_financial.FinaMainbzRequest) ([]stock_financial.FinaMainbzItem, error) {
	return stock_financial.FinaMainbz(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) Forecast(ctx context.Context, req *stock_financial.ForecastRequest) ([]stock_financial.ForecastItem, error) {
	return stock_financial.Forecast(ctx, impl.client, req)
}

func (impl *stockFinancialImpl) Income(ctx context.Context, req *stock_financial.IncomeRequest) ([]stock_financial.IncomeItem, error) {
	return stock_financial.Income(ctx, impl.client, req)
}

type stockFundFlowImpl struct {
	client *sdk.Client
}

func (impl *stockFundFlowImpl) Api343(ctx context.Context, req *stock_fund_flow.Api343Request) ([]stock_fund_flow.Api343Item, error) {
	return stock_fund_flow.Api343(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Api344(ctx context.Context, req *stock_fund_flow.Api344Request) ([]stock_fund_flow.Api344Item, error) {
	return stock_fund_flow.Api344(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Api345(ctx context.Context, req *stock_fund_flow.Api345Request) ([]stock_fund_flow.Api345Item, error) {
	return stock_fund_flow.Api345(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Api348(ctx context.Context, req *stock_fund_flow.Api348Request) ([]stock_fund_flow.Api348Item, error) {
	return stock_fund_flow.Api348(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Api349(ctx context.Context, req *stock_fund_flow.Api349Request) ([]stock_fund_flow.Api349Item, error) {
	return stock_fund_flow.Api349(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Api371(ctx context.Context, req *stock_fund_flow.Api371Request) ([]stock_fund_flow.Api371Item, error) {
	return stock_fund_flow.Api371(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) Moneyflow(ctx context.Context, req *stock_fund_flow.MoneyflowRequest) ([]stock_fund_flow.MoneyflowItem, error) {
	return stock_fund_flow.Moneyflow(ctx, impl.client, req)
}

func (impl *stockFundFlowImpl) MoneyflowHsgt(ctx context.Context, req *stock_fund_flow.MoneyflowHsgtRequest) ([]stock_fund_flow.MoneyflowHsgtItem, error) {
	return stock_fund_flow.MoneyflowHsgt(ctx, impl.client, req)
}

type stockMarginImpl struct {
	client *sdk.Client
}

func (impl *stockMarginImpl) Api326(ctx context.Context, req *stock_margin.Api326Request) ([]stock_margin.Api326Item, error) {
	return stock_margin.Api326(ctx, impl.client, req)
}

func (impl *stockMarginImpl) Api332(ctx context.Context, req *stock_margin.Api332Request) ([]stock_margin.Api332Item, error) {
	return stock_margin.Api332(ctx, impl.client, req)
}

func (impl *stockMarginImpl) Api333(ctx context.Context, req *stock_margin.Api333Request) ([]stock_margin.Api333Item, error) {
	return stock_margin.Api333(ctx, impl.client, req)
}

func (impl *stockMarginImpl) Api334(ctx context.Context, req *stock_margin.Api334Request) ([]stock_margin.Api334Item, error) {
	return stock_margin.Api334(ctx, impl.client, req)
}

func (impl *stockMarginImpl) Margin(ctx context.Context, req *stock_margin.MarginRequest) ([]stock_margin.MarginItem, error) {
	return stock_margin.Margin(ctx, impl.client, req)
}

func (impl *stockMarginImpl) MarginDetail(ctx context.Context, req *stock_margin.MarginDetailRequest) ([]stock_margin.MarginDetailItem, error) {
	return stock_margin.MarginDetail(ctx, impl.client, req)
}

func (impl *stockMarginImpl) MarginLend(ctx context.Context, req *stock_margin.MarginLendRequest) ([]stock_margin.MarginLendItem, error) {
	return stock_margin.MarginLend(ctx, impl.client, req)
}

type stockMarketImpl struct {
	client *sdk.Client
}

func (impl *stockMarketImpl) AdjFactor(ctx context.Context, req *stock_market.AdjFactorRequest) ([]stock_market.AdjFactorItem, error) {
	return stock_market.AdjFactor(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Api315(ctx context.Context, req *stock_market.Api315Request) ([]stock_market.Api315Item, error) {
	return stock_market.Api315(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Api316(ctx context.Context, req *stock_market.Api316Request) ([]stock_market.Api316Item, error) {
	return stock_market.Api316(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Api317(ctx context.Context, req *stock_market.Api317Request) ([]stock_market.Api317Item, error) {
	return stock_market.Api317(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Api336(ctx context.Context, req *stock_market.Api336Request) ([]stock_market.Api336Item, error) {
	return stock_market.Api336(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Api365(ctx context.Context, req *stock_market.Api365Request) ([]stock_market.Api365Item, error) {
	return stock_market.Api365(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Daily(ctx context.Context, req *stock_market.DailyRequest) ([]stock_market.DailyItem, error) {
	return stock_market.Daily(ctx, impl.client, req)
}

func (impl *stockMarketImpl) MarketDailyBasic(ctx context.Context, req *stock_market.DailyBasicRequest) ([]stock_market.DailyBasicItem, error) {
	return stock_market.DailyBasic(ctx, impl.client, req)
}

func (impl *stockMarketImpl) DailyNow(ctx context.Context, req *stock_market.DailyNowRequest) ([]stock_market.DailyNowItem, error) {
	return stock_market.DailyNow(ctx, impl.client, req)
}

func (impl *stockMarketImpl) GgDaily(ctx context.Context, req *stock_market.GgDailyRequest) ([]stock_market.GgDailyItem, error) {
	return stock_market.GgDaily(ctx, impl.client, req)
}

func (impl *stockMarketImpl) GgMonthly(ctx context.Context, req *stock_market.GgMonthlyRequest) ([]stock_market.GgMonthlyItem, error) {
	return stock_market.GgMonthly(ctx, impl.client, req)
}

func (impl *stockMarketImpl) GgTop10(ctx context.Context, req *stock_market.GgTop10Request) ([]stock_market.GgTop10Item, error) {
	return stock_market.GgTop10(ctx, impl.client, req)
}

func (impl *stockMarketImpl) HgTop10(ctx context.Context, req *stock_market.HgTop10Request) ([]stock_market.HgTop10Item, error) {
	return stock_market.HgTop10(ctx, impl.client, req)
}

func (impl *stockMarketImpl) MarketLimitList(ctx context.Context, req *stock_market.LimitListRequest) ([]stock_market.LimitListItem, error) {
	return stock_market.LimitList(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Monthly(ctx context.Context, req *stock_market.MonthlyRequest) ([]stock_market.MonthlyItem, error) {
	return stock_market.Monthly(ctx, impl.client, req)
}

func (impl *stockMarketImpl) StkMins(ctx context.Context, req *stock_market.StkMinsRequest) ([]stock_market.StkMinsItem, error) {
	return stock_market.StkMins(ctx, impl.client, req)
}

func (impl *stockMarketImpl) StkMinsNow(ctx context.Context, req *stock_market.StkMinsNowRequest) ([]stock_market.StkMinsNowItem, error) {
	return stock_market.StkMinsNow(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Suspend(ctx context.Context, req *stock_market.SuspendRequest) ([]stock_market.SuspendItem, error) {
	return stock_market.Suspend(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Universal(ctx context.Context, req *stock_market.UniversalRequest) ([]stock_market.UniversalItem, error) {
	return stock_market.Universal(ctx, impl.client, req)
}

func (impl *stockMarketImpl) Weekly(ctx context.Context, req *stock_market.WeeklyRequest) ([]stock_market.WeeklyItem, error) {
	return stock_market.Weekly(ctx, impl.client, req)
}

type stockReferenceImpl struct {
	client *sdk.Client
}

func (impl *stockReferenceImpl) Api164(ctx context.Context, req *stock_reference.Api164Request) ([]stock_reference.Api164Item, error) {
	return stock_reference.Api164(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) Api165(ctx context.Context, req *stock_reference.Api165Request) ([]stock_reference.Api165Item, error) {
	return stock_reference.Api165(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) BlockTrade(ctx context.Context, req *stock_reference.BlockTradeRequest) ([]stock_reference.BlockTradeItem, error) {
	return stock_reference.BlockTrade(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) PledgeDetail(ctx context.Context, req *stock_reference.PledgeDetailRequest) ([]stock_reference.PledgeDetailItem, error) {
	return stock_reference.PledgeDetail(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) PledgeStat(ctx context.Context, req *stock_reference.PledgeStatRequest) ([]stock_reference.PledgeStatItem, error) {
	return stock_reference.PledgeStat(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) Repurchase(ctx context.Context, req *stock_reference.RepurchaseRequest) ([]stock_reference.RepurchaseItem, error) {
	return stock_reference.Repurchase(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) StkChange(ctx context.Context, req *stock_reference.StkChangeRequest) ([]stock_reference.StkChangeItem, error) {
	return stock_reference.StkChange(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) StkHoldernumber(ctx context.Context, req *stock_reference.StkHoldernumberRequest) ([]stock_reference.StkHoldernumberItem, error) {
	return stock_reference.StkHoldernumber(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) Top10Floatholders(ctx context.Context, req *stock_reference.Top10FloatholdersRequest) ([]stock_reference.Top10FloatholdersItem, error) {
	return stock_reference.Top10Floatholders(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) Top10Holders(ctx context.Context, req *stock_reference.Top10HoldersRequest) ([]stock_reference.Top10HoldersItem, error) {
	return stock_reference.Top10Holders(ctx, impl.client, req)
}

func (impl *stockReferenceImpl) UnlockShare(ctx context.Context, req *stock_reference.UnlockShareRequest) ([]stock_reference.UnlockShareItem, error) {
	return stock_reference.UnlockShare(ctx, impl.client, req)
}

func (s stockImpl) BseCodeTranslate(ctx context.Context, req *stock_basic.BseCodeTranslateRequest) ([]stock_basic.BseCodeTranslateItem, error) {
	return s.stockBasic.BseCodeTranslate(ctx, req)
}

func (s stockImpl) BasicDailyBasic(ctx context.Context, req *stock_basic.DailyBasicRequest) ([]stock_basic.DailyBasicItem, error) {
	return s.stockBasic.BasicDailyBasic(ctx, req)
}

func (s stockImpl) HkHoldStock(ctx context.Context, req *stock_basic.HkHoldStockRequest) ([]stock_basic.HkHoldStockItem, error) {
	return s.stockBasic.HkHoldStock(ctx, req)
}

func (s stockImpl) Namechange(ctx context.Context, req *stock_basic.NamechangeRequest) ([]stock_basic.NamechangeItem, error) {
	return s.stockBasic.Namechange(ctx, req)
}

func (s stockImpl) NewShare(ctx context.Context, req *stock_basic.NewShareRequest) ([]stock_basic.NewShareItem, error) {
	return s.stockBasic.NewShare(ctx, req)
}

func (s stockImpl) StRiskWarning(ctx context.Context, req *stock_basic.StRiskWarningRequest) ([]stock_basic.StRiskWarningItem, error) {
	return s.stockBasic.StRiskWarning(ctx, req)
}

func (s stockImpl) StStockList(ctx context.Context, req *stock_basic.StStockListRequest) ([]stock_basic.StStockListItem, error) {
	return s.stockBasic.StStockList(ctx, req)
}

func (s stockImpl) StkManagers(ctx context.Context, req *stock_basic.StkManagersRequest) ([]stock_basic.StkManagersItem, error) {
	return s.stockBasic.StkManagers(ctx, req)
}

func (s stockImpl) StkRewards(ctx context.Context, req *stock_basic.StkRewardsRequest) ([]stock_basic.StkRewardsItem, error) {
	return s.stockBasic.StkRewards(ctx, req)
}

func (s stockImpl) StockBasic(ctx context.Context, req *stock_basic.StockBasicRequest) ([]stock_basic.StockBasicItem, error) {
	return s.stockBasic.StockBasic(ctx, req)
}

func (s stockImpl) StockCompany(ctx context.Context, req *stock_basic.StockCompanyRequest) ([]stock_basic.StockCompanyItem, error) {
	return s.stockBasic.StockCompany(ctx, req)
}

func (s stockImpl) StockHistory(ctx context.Context, req *stock_basic.StockHistoryRequest) ([]stock_basic.StockHistoryItem, error) {
	return s.stockBasic.StockHistory(ctx, req)
}

func (s stockImpl) TradeCal(ctx context.Context, req *stock_basic.TradeCalRequest) ([]stock_basic.TradeCalItem, error) {
	return s.stockBasic.TradeCal(ctx, req)
}

func (s stockImpl) Api260(ctx context.Context, req *stock_board.Api260Request) ([]stock_board.Api260Item, error) {
	return s.stockBoard.Api260(ctx, req)
}

func (s stockImpl) Api347(ctx context.Context, req *stock_board.Api347Request) ([]stock_board.Api347Item, error) {
	return s.stockBoard.Api347(ctx, req)
}

func (s stockImpl) Api351(ctx context.Context, req *stock_board.Api351Request) ([]stock_board.Api351Item, error) {
	return s.stockBoard.Api351(ctx, req)
}

func (s stockImpl) Api369(ctx context.Context, req *stock_board.Api369Request) ([]stock_board.Api369Item, error) {
	return s.stockBoard.Api369(ctx, req)
}

func (s stockImpl) DragonDetail(ctx context.Context, req *stock_board.DragonDetailRequest) ([]stock_board.DragonDetailItem, error) {
	return s.stockBoard.DragonDetail(ctx, req)
}

func (s stockImpl) DragonList(ctx context.Context, req *stock_board.DragonListRequest) ([]stock_board.DragonListItem, error) {
	return s.stockBoard.DragonList(ctx, req)
}

func (s stockImpl) EmConcept(ctx context.Context, req *stock_board.EmConceptRequest) ([]stock_board.EmConceptItem, error) {
	return s.stockBoard.EmConcept(ctx, req)
}

func (s stockImpl) EmHot(ctx context.Context, req *stock_board.EmHotRequest) ([]stock_board.EmHotItem, error) {
	return s.stockBoard.EmHot(ctx, req)
}

func (s stockImpl) EmIndex(ctx context.Context, req *stock_board.EmIndexRequest) ([]stock_board.EmIndexItem, error) {
	return s.stockBoard.EmIndex(ctx, req)
}

func (s stockImpl) BoardLimitList(ctx context.Context, req *stock_board.LimitListRequest) ([]stock_board.LimitListItem, error) {
	return s.stockBoard.BoardLimitList(ctx, req)
}

func (s stockImpl) LimitListD(ctx context.Context, req *stock_board.LimitListDRequest) ([]stock_board.LimitListDItem, error) {
	return s.stockBoard.LimitListD(ctx, req)
}

func (s stockImpl) LimitListSec(ctx context.Context, req *stock_board.LimitListSecRequest) ([]stock_board.LimitListSecItem, error) {
	return s.stockBoard.LimitListSec(ctx, req)
}

func (s stockImpl) TdxIndex(ctx context.Context, req *stock_board.TdxIndexRequest) ([]stock_board.TdxIndexItem, error) {
	return s.stockBoard.TdxIndex(ctx, req)
}

func (s stockImpl) TdxMember(ctx context.Context, req *stock_board.TdxMemberRequest) ([]stock_board.TdxMemberItem, error) {
	return s.stockBoard.TdxMember(ctx, req)
}

func (s stockImpl) TdxSector(ctx context.Context, req *stock_board.TdxSectorRequest) ([]stock_board.TdxSectorItem, error) {
	return s.stockBoard.TdxSector(ctx, req)
}

func (s stockImpl) ThsConcept(ctx context.Context, req *stock_board.ThsConceptRequest) ([]stock_board.ThsConceptItem, error) {
	return s.stockBoard.ThsConcept(ctx, req)
}

func (s stockImpl) ThsHot(ctx context.Context, req *stock_board.ThsHotRequest) ([]stock_board.ThsHotItem, error) {
	return s.stockBoard.ThsHot(ctx, req)
}

func (s stockImpl) ThsMember(ctx context.Context, req *stock_board.ThsMemberRequest) ([]stock_board.ThsMemberItem, error) {
	return s.stockBoard.ThsMember(ctx, req)
}

func (s stockImpl) TopInst(ctx context.Context, req *stock_board.TopInstRequest) ([]stock_board.TopInstItem, error) {
	return s.stockBoard.TopInst(ctx, req)
}

func (s stockImpl) TopList(ctx context.Context, req *stock_board.TopListRequest) ([]stock_board.TopListItem, error) {
	return s.stockBoard.TopList(ctx, req)
}

func (s stockImpl) AhLotSize(ctx context.Context, req *stock_feature.AhLotSizeRequest) ([]stock_feature.AhLotSizeItem, error) {
	return s.stockFeature.AhLotSize(ctx, req)
}

func (s stockImpl) Api328(ctx context.Context, req *stock_feature.Api328Request) ([]stock_feature.Api328Item, error) {
	return s.stockFeature.Api328(ctx, req)
}

func (s stockImpl) Auction(ctx context.Context, req *stock_feature.AuctionRequest) ([]stock_feature.AuctionItem, error) {
	return s.stockFeature.Auction(ctx, req)
}

func (s stockImpl) CcassDetail(ctx context.Context, req *stock_feature.CcassDetailRequest) ([]stock_feature.CcassDetailItem, error) {
	return s.stockFeature.CcassDetail(ctx, req)
}

func (s stockImpl) CcassStat(ctx context.Context, req *stock_feature.CcassStatRequest) ([]stock_feature.CcassStatItem, error) {
	return s.stockFeature.CcassStat(ctx, req)
}

func (s stockImpl) ChipDistribution(ctx context.Context, req *stock_feature.ChipDistributionRequest) ([]stock_feature.ChipDistributionItem, error) {
	return s.stockFeature.ChipDistribution(ctx, req)
}

func (s stockImpl) ForecastProfit(ctx context.Context, req *stock_feature.ForecastProfitRequest) ([]stock_feature.ForecastProfitItem, error) {
	return s.stockFeature.ForecastProfit(ctx, req)
}

func (s stockImpl) HkHold(ctx context.Context, req *stock_feature.HkHoldRequest) ([]stock_feature.HkHoldItem, error) {
	return s.stockFeature.HkHold(ctx, req)
}

func (s stockImpl) Magic9(ctx context.Context, req *stock_feature.Magic9Request) ([]stock_feature.Magic9Item, error) {
	return s.stockFeature.Magic9(ctx, req)
}

func (s stockImpl) MonthlyStock(ctx context.Context, req *stock_feature.MonthlyStockRequest) ([]stock_feature.MonthlyStockItem, error) {
	return s.stockFeature.MonthlyStock(ctx, req)
}

func (s stockImpl) OrgResearch(ctx context.Context, req *stock_feature.OrgResearchRequest) ([]stock_feature.OrgResearchItem, error) {
	return s.stockFeature.OrgResearch(ctx, req)
}

func (s stockImpl) Balancesheet(ctx context.Context, req *stock_financial.BalancesheetRequest) ([]stock_financial.BalancesheetItem, error) {
	return s.stockFinancial.Balancesheet(ctx, req)
}

func (s stockImpl) Cashflow(ctx context.Context, req *stock_financial.CashflowRequest) ([]stock_financial.CashflowItem, error) {
	return s.stockFinancial.Cashflow(ctx, req)
}

func (s stockImpl) Dividend(ctx context.Context, req *stock_financial.DividendRequest) ([]stock_financial.DividendItem, error) {
	return s.stockFinancial.Dividend(ctx, req)
}

func (s stockImpl) Express(ctx context.Context, req *stock_financial.ExpressRequest) ([]stock_financial.ExpressItem, error) {
	return s.stockFinancial.Express(ctx, req)
}

func (s stockImpl) FinaAudit(ctx context.Context, req *stock_financial.FinaAuditRequest) ([]stock_financial.FinaAuditItem, error) {
	return s.stockFinancial.FinaAudit(ctx, req)
}

func (s stockImpl) FinaDisclosure(ctx context.Context, req *stock_financial.FinaDisclosureRequest) ([]stock_financial.FinaDisclosureItem, error) {
	return s.stockFinancial.FinaDisclosure(ctx, req)
}

func (s stockImpl) FinaIndicator(ctx context.Context, req *stock_financial.FinaIndicatorRequest) ([]stock_financial.FinaIndicatorItem, error) {
	return s.stockFinancial.FinaIndicator(ctx, req)
}

func (s stockImpl) FinaMainbz(ctx context.Context, req *stock_financial.FinaMainbzRequest) ([]stock_financial.FinaMainbzItem, error) {
	return s.stockFinancial.FinaMainbz(ctx, req)
}

func (s stockImpl) Forecast(ctx context.Context, req *stock_financial.ForecastRequest) ([]stock_financial.ForecastItem, error) {
	return s.stockFinancial.Forecast(ctx, req)
}

func (s stockImpl) Income(ctx context.Context, req *stock_financial.IncomeRequest) ([]stock_financial.IncomeItem, error) {
	return s.stockFinancial.Income(ctx, req)
}

func (s stockImpl) Api343(ctx context.Context, req *stock_fund_flow.Api343Request) ([]stock_fund_flow.Api343Item, error) {
	return s.stockFundFlow.Api343(ctx, req)
}

func (s stockImpl) Api344(ctx context.Context, req *stock_fund_flow.Api344Request) ([]stock_fund_flow.Api344Item, error) {
	return s.stockFundFlow.Api344(ctx, req)
}

func (s stockImpl) Api345(ctx context.Context, req *stock_fund_flow.Api345Request) ([]stock_fund_flow.Api345Item, error) {
	return s.stockFundFlow.Api345(ctx, req)
}

func (s stockImpl) Api348(ctx context.Context, req *stock_fund_flow.Api348Request) ([]stock_fund_flow.Api348Item, error) {
	return s.stockFundFlow.Api348(ctx, req)
}

func (s stockImpl) Api349(ctx context.Context, req *stock_fund_flow.Api349Request) ([]stock_fund_flow.Api349Item, error) {
	return s.stockFundFlow.Api349(ctx, req)
}

func (s stockImpl) Api371(ctx context.Context, req *stock_fund_flow.Api371Request) ([]stock_fund_flow.Api371Item, error) {
	return s.stockFundFlow.Api371(ctx, req)
}

func (s stockImpl) Moneyflow(ctx context.Context, req *stock_fund_flow.MoneyflowRequest) ([]stock_fund_flow.MoneyflowItem, error) {
	return s.stockFundFlow.Moneyflow(ctx, req)
}

func (s stockImpl) MoneyflowHsgt(ctx context.Context, req *stock_fund_flow.MoneyflowHsgtRequest) ([]stock_fund_flow.MoneyflowHsgtItem, error) {
	return s.stockFundFlow.MoneyflowHsgt(ctx, req)
}

func (s stockImpl) Api326(ctx context.Context, req *stock_margin.Api326Request) ([]stock_margin.Api326Item, error) {
	return s.stockMargin.Api326(ctx, req)
}

func (s stockImpl) Api332(ctx context.Context, req *stock_margin.Api332Request) ([]stock_margin.Api332Item, error) {
	return s.stockMargin.Api332(ctx, req)
}

func (s stockImpl) Api333(ctx context.Context, req *stock_margin.Api333Request) ([]stock_margin.Api333Item, error) {
	return s.stockMargin.Api333(ctx, req)
}

func (s stockImpl) Api334(ctx context.Context, req *stock_margin.Api334Request) ([]stock_margin.Api334Item, error) {
	return s.stockMargin.Api334(ctx, req)
}

func (s stockImpl) Margin(ctx context.Context, req *stock_margin.MarginRequest) ([]stock_margin.MarginItem, error) {
	return s.stockMargin.Margin(ctx, req)
}

func (s stockImpl) MarginDetail(ctx context.Context, req *stock_margin.MarginDetailRequest) ([]stock_margin.MarginDetailItem, error) {
	return s.stockMargin.MarginDetail(ctx, req)
}

func (s stockImpl) MarginLend(ctx context.Context, req *stock_margin.MarginLendRequest) ([]stock_margin.MarginLendItem, error) {
	return s.stockMargin.MarginLend(ctx, req)
}

func (s stockImpl) AdjFactor(ctx context.Context, req *stock_market.AdjFactorRequest) ([]stock_market.AdjFactorItem, error) {
	return s.stockMarket.AdjFactor(ctx, req)
}

func (s stockImpl) Api315(ctx context.Context, req *stock_market.Api315Request) ([]stock_market.Api315Item, error) {
	return s.stockMarket.Api315(ctx, req)
}

func (s stockImpl) Api316(ctx context.Context, req *stock_market.Api316Request) ([]stock_market.Api316Item, error) {
	return s.stockMarket.Api316(ctx, req)
}

func (s stockImpl) Api317(ctx context.Context, req *stock_market.Api317Request) ([]stock_market.Api317Item, error) {
	return s.stockMarket.Api317(ctx, req)
}

func (s stockImpl) Api336(ctx context.Context, req *stock_market.Api336Request) ([]stock_market.Api336Item, error) {
	return s.stockMarket.Api336(ctx, req)
}

func (s stockImpl) Api365(ctx context.Context, req *stock_market.Api365Request) ([]stock_market.Api365Item, error) {
	return s.stockMarket.Api365(ctx, req)
}

func (s stockImpl) Daily(ctx context.Context, req *stock_market.DailyRequest) ([]stock_market.DailyItem, error) {
	return s.stockMarket.Daily(ctx, req)
}

func (s stockImpl) MarketDailyBasic(ctx context.Context, req *stock_market.DailyBasicRequest) ([]stock_market.DailyBasicItem, error) {
	return s.stockMarket.MarketDailyBasic(ctx, req)
}

func (s stockImpl) DailyNow(ctx context.Context, req *stock_market.DailyNowRequest) ([]stock_market.DailyNowItem, error) {
	return s.stockMarket.DailyNow(ctx, req)
}

func (s stockImpl) GgDaily(ctx context.Context, req *stock_market.GgDailyRequest) ([]stock_market.GgDailyItem, error) {
	return s.stockMarket.GgDaily(ctx, req)
}

func (s stockImpl) GgMonthly(ctx context.Context, req *stock_market.GgMonthlyRequest) ([]stock_market.GgMonthlyItem, error) {
	return s.stockMarket.GgMonthly(ctx, req)
}

func (s stockImpl) GgTop10(ctx context.Context, req *stock_market.GgTop10Request) ([]stock_market.GgTop10Item, error) {
	return s.stockMarket.GgTop10(ctx, req)
}

func (s stockImpl) HgTop10(ctx context.Context, req *stock_market.HgTop10Request) ([]stock_market.HgTop10Item, error) {
	return s.stockMarket.HgTop10(ctx, req)
}

func (s stockImpl) MarketLimitList(ctx context.Context, req *stock_market.LimitListRequest) ([]stock_market.LimitListItem, error) {
	return s.stockMarket.MarketLimitList(ctx, req)
}

func (s stockImpl) Monthly(ctx context.Context, req *stock_market.MonthlyRequest) ([]stock_market.MonthlyItem, error) {
	return s.stockMarket.Monthly(ctx, req)
}

func (s stockImpl) StkMins(ctx context.Context, req *stock_market.StkMinsRequest) ([]stock_market.StkMinsItem, error) {
	return s.stockMarket.StkMins(ctx, req)
}

func (s stockImpl) StkMinsNow(ctx context.Context, req *stock_market.StkMinsNowRequest) ([]stock_market.StkMinsNowItem, error) {
	return s.stockMarket.StkMinsNow(ctx, req)
}

func (s stockImpl) Suspend(ctx context.Context, req *stock_market.SuspendRequest) ([]stock_market.SuspendItem, error) {
	return s.stockMarket.Suspend(ctx, req)
}

func (s stockImpl) Universal(ctx context.Context, req *stock_market.UniversalRequest) ([]stock_market.UniversalItem, error) {
	return s.stockMarket.Universal(ctx, req)
}

func (s stockImpl) Weekly(ctx context.Context, req *stock_market.WeeklyRequest) ([]stock_market.WeeklyItem, error) {
	return s.stockMarket.Weekly(ctx, req)
}

func (s stockImpl) Api164(ctx context.Context, req *stock_reference.Api164Request) ([]stock_reference.Api164Item, error) {
	return s.stockReference.Api164(ctx, req)
}

func (s stockImpl) Api165(ctx context.Context, req *stock_reference.Api165Request) ([]stock_reference.Api165Item, error) {
	return s.stockReference.Api165(ctx, req)
}

func (s stockImpl) BlockTrade(ctx context.Context, req *stock_reference.BlockTradeRequest) ([]stock_reference.BlockTradeItem, error) {
	return s.stockReference.BlockTrade(ctx, req)
}

func (s stockImpl) PledgeDetail(ctx context.Context, req *stock_reference.PledgeDetailRequest) ([]stock_reference.PledgeDetailItem, error) {
	return s.stockReference.PledgeDetail(ctx, req)
}

func (s stockImpl) PledgeStat(ctx context.Context, req *stock_reference.PledgeStatRequest) ([]stock_reference.PledgeStatItem, error) {
	return s.stockReference.PledgeStat(ctx, req)
}

func (s stockImpl) Repurchase(ctx context.Context, req *stock_reference.RepurchaseRequest) ([]stock_reference.RepurchaseItem, error) {
	return s.stockReference.Repurchase(ctx, req)
}

func (s stockImpl) StkChange(ctx context.Context, req *stock_reference.StkChangeRequest) ([]stock_reference.StkChangeItem, error) {
	return s.stockReference.StkChange(ctx, req)
}

func (s stockImpl) StkHoldernumber(ctx context.Context, req *stock_reference.StkHoldernumberRequest) ([]stock_reference.StkHoldernumberItem, error) {
	return s.stockReference.StkHoldernumber(ctx, req)
}

func (s stockImpl) Top10Floatholders(ctx context.Context, req *stock_reference.Top10FloatholdersRequest) ([]stock_reference.Top10FloatholdersItem, error) {
	return s.stockReference.Top10Floatholders(ctx, req)
}

func (s stockImpl) Top10Holders(ctx context.Context, req *stock_reference.Top10HoldersRequest) ([]stock_reference.Top10HoldersItem, error) {
	return s.stockReference.Top10Holders(ctx, req)
}

func (s stockImpl) UnlockShare(ctx context.Context, req *stock_reference.UnlockShareRequest) ([]stock_reference.UnlockShareItem, error) {
	return s.stockReference.UnlockShare(ctx, req)
}

func newStockImpl(client *sdk.Client) Stock {
	return stockImpl{
		client:           client,
		stockBasic:       &stockBasicImpl{client: client},
		stockBoard:       &stockBoardImpl{client: client},
		stockFeature:     &stockFeatureImpl{client: client},
		stockFinancial:   &stockFinancialImpl{client: client},
		stockFundFlow:    &stockFundFlowImpl{client: client},
		stockMargin:      &stockMarginImpl{client: client},
		stockMarket:      &stockMarketImpl{client: client},
		stockReference:   &stockReferenceImpl{client: client},
	}
}