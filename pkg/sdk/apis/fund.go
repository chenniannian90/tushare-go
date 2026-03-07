package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/fund"
)

type Fund interface {
	FundBasic(ctx context.Context, req *fund.FundBasicRequest) ([]fund.FundBasicItem, error)
	FundNav(ctx context.Context, req *fund.FundNavRequest) ([]fund.FundNavItem, error)
	FundDiv(ctx context.Context, req *fund.FundDivRequest) ([]fund.FundDivItem, error)
	FundManager(ctx context.Context, req *fund.FundManagerRequest) ([]fund.FundManagerItem, error)
	Api359(ctx context.Context, req *fund.Api359Request) ([]fund.Api359Item, error)
}

type fundImpl struct {
	client *sdk.Client
}

func (impl *fundImpl) FundBasic(ctx context.Context, req *fund.FundBasicRequest) ([]fund.FundBasicItem, error) {
	return fund.FundBasic(ctx, impl.client, req)
}

func (impl *fundImpl) FundNav(ctx context.Context, req *fund.FundNavRequest) ([]fund.FundNavItem, error) {
	return fund.FundNav(ctx, impl.client, req)
}

func (impl *fundImpl) FundDiv(ctx context.Context, req *fund.FundDivRequest) ([]fund.FundDivItem, error) {
	return fund.FundDiv(ctx, impl.client, req)
}

func (impl *fundImpl) FundManager(ctx context.Context, req *fund.FundManagerRequest) ([]fund.FundManagerItem, error) {
	return fund.FundManager(ctx, impl.client, req)
}

func (impl *fundImpl) Api359(ctx context.Context, req *fund.Api359Request) ([]fund.Api359Item, error) {
	return fund.Api359(ctx, impl.client, req)
}

func newFundImpl(client *sdk.Client) Fund {
	return &fundImpl{client: client}
}
