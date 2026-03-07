package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/hk_stock"
)

type HKStock interface {
	HkBasic(ctx context.Context, req *hk_stock.HkBasicRequest) ([]hk_stock.HkBasicItem, error)
	HkDaily(ctx context.Context, req *hk_stock.HkDailyRequest) ([]hk_stock.HkDailyItem, error)
	HkCal(ctx context.Context, req *hk_stock.HkCalRequest) ([]hk_stock.HkCalItem, error)
	HkMin(ctx context.Context, req *hk_stock.HkMinRequest) ([]hk_stock.HkMinItem, error)
	HkFactor(ctx context.Context, req *hk_stock.HkFactorRequest) ([]hk_stock.HkFactorItem, error)
}

type hkStockImpl struct {
	client *sdk.Client
}

func (impl *hkStockImpl) HkBasic(ctx context.Context, req *hk_stock.HkBasicRequest) ([]hk_stock.HkBasicItem, error) {
	return hk_stock.HkBasic(ctx, impl.client, req)
}

func (impl *hkStockImpl) HkDaily(ctx context.Context, req *hk_stock.HkDailyRequest) ([]hk_stock.HkDailyItem, error) {
	return hk_stock.HkDaily(ctx, impl.client, req)
}

func (impl *hkStockImpl) HkCal(ctx context.Context, req *hk_stock.HkCalRequest) ([]hk_stock.HkCalItem, error) {
	return hk_stock.HkCal(ctx, impl.client, req)
}

func (impl *hkStockImpl) HkMin(ctx context.Context, req *hk_stock.HkMinRequest) ([]hk_stock.HkMinItem, error) {
	return hk_stock.HkMin(ctx, impl.client, req)
}

func (impl *hkStockImpl) HkFactor(ctx context.Context, req *hk_stock.HkFactorRequest) ([]hk_stock.HkFactorItem, error) {
	return hk_stock.HkFactor(ctx, impl.client, req)
}

func newHkStockImpl(client *sdk.Client) HKStock {
	return &hkStockImpl{client: client}
}
