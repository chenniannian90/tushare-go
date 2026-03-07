package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/wealth/wealth_fund_sales"
)

type Wealth interface {
	WealthFundSales
}

type WealthFundSales interface {
	FundSales(ctx context.Context, req *wealth_fund_sales.FundSalesRequest) ([]wealth_fund_sales.FundSalesItem, error)
}

type wealthImpl struct {
	client *sdk.Client
	WealthFundSales
}

type wealthFundSalesImpl struct {
	client *sdk.Client
}

func (impl *wealthFundSalesImpl) FundSales(ctx context.Context, req *wealth_fund_sales.FundSalesRequest) ([]wealth_fund_sales.FundSalesItem, error) {
	return wealth_fund_sales.FundSales(ctx, impl.client, req)
}

func newWealthImpl(client *sdk.Client) Wealth {
	return wealthImpl{
		client:          client,
		WealthFundSales: &wealthFundSalesImpl{client: client},
	}
}
