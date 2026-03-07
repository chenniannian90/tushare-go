package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/us_stock"
)

type USStock interface {
	UsBasic(ctx context.Context, req *us_stock.UsBasicRequest) ([]us_stock.UsBasicItem, error)
	UsDaily(ctx context.Context, req *us_stock.UsDailyRequest) ([]us_stock.UsDailyItem, error)
	UsCal(ctx context.Context, req *us_stock.UsCalRequest) ([]us_stock.UsCalItem, error)
	UsFactor(ctx context.Context, req *us_stock.UsFactorRequest) ([]us_stock.UsFactorItem, error)
}

type usStockImpl struct {
	client *sdk.Client
}

func (impl *usStockImpl) UsBasic(ctx context.Context, req *us_stock.UsBasicRequest) ([]us_stock.UsBasicItem, error) {
	return us_stock.UsBasic(ctx, impl.client, req)
}

func (impl *usStockImpl) UsDaily(ctx context.Context, req *us_stock.UsDailyRequest) ([]us_stock.UsDailyItem, error) {
	return us_stock.UsDaily(ctx, impl.client, req)
}

func (impl *usStockImpl) UsCal(ctx context.Context, req *us_stock.UsCalRequest) ([]us_stock.UsCalItem, error) {
	return us_stock.UsCal(ctx, impl.client, req)
}

func (impl *usStockImpl) UsFactor(ctx context.Context, req *us_stock.UsFactorRequest) ([]us_stock.UsFactorItem, error) {
	return us_stock.UsFactor(ctx, impl.client, req)
}

func newUsStockImpl(client *sdk.Client) USStock {
	return &usStockImpl{client: client}
}
