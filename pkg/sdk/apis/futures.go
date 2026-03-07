package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/futures"
)

type Futures interface {
	FutBasic(ctx context.Context, req *futures.FutBasicRequest) ([]futures.FutBasicItem, error)
	FutDaily(ctx context.Context, req *futures.FutDailyRequest) ([]futures.FutDailyItem, error)
	FutWeekly(ctx context.Context, req *futures.FutWeeklyRequest) ([]futures.FutWeeklyItem, error)
	FutSettlement(ctx context.Context, req *futures.FutSettlementRequest) ([]futures.FutSettlementItem, error)
	TradeCal(ctx context.Context, req *futures.TradeCalRequest) ([]futures.TradeCalItem, error)
}

type futuresImpl struct {
	client *sdk.Client
}

func (impl *futuresImpl) FutBasic(ctx context.Context, req *futures.FutBasicRequest) ([]futures.FutBasicItem, error) {
	return futures.FutBasic(ctx, impl.client, req)
}

func (impl *futuresImpl) FutDaily(ctx context.Context, req *futures.FutDailyRequest) ([]futures.FutDailyItem, error) {
	return futures.FutDaily(ctx, impl.client, req)
}

func (impl *futuresImpl) FutWeekly(ctx context.Context, req *futures.FutWeeklyRequest) ([]futures.FutWeeklyItem, error) {
	return futures.FutWeekly(ctx, impl.client, req)
}

func (impl *futuresImpl) FutSettlement(ctx context.Context, req *futures.FutSettlementRequest) ([]futures.FutSettlementItem, error) {
	return futures.FutSettlement(ctx, impl.client, req)
}

func (impl *futuresImpl) TradeCal(ctx context.Context, req *futures.TradeCalRequest) ([]futures.TradeCalItem, error) {
	return futures.TradeCal(ctx, impl.client, req)
}

func newFuturesImpl(client *sdk.Client) Futures {
	return &futuresImpl{client: client}
}
