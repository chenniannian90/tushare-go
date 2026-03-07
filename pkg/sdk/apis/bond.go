package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/bond"
)

type Bond interface {
	BondOc(ctx context.Context, req *bond.BondOcRequest) ([]bond.BondOcItem, error)
	BondRepurchase(ctx context.Context, req *bond.BondRepurchaseRequest) ([]bond.BondRepurchaseItem, error)
	Api272(ctx context.Context, req *bond.Api272Request) ([]bond.Api272Item, error)
	GlobalCalendar(ctx context.Context, req *bond.GlobalCalendarRequest) ([]bond.GlobalCalendarItem, error)
	CbInterest(ctx context.Context, req *bond.CbInterestRequest) ([]bond.CbInterestItem, error)
}

type bondImpl struct {
	client *sdk.Client
}

func (impl *bondImpl) BondOc(ctx context.Context, req *bond.BondOcRequest) ([]bond.BondOcItem, error) {
	return bond.BondOc(ctx, impl.client, req)
}

func (impl *bondImpl) BondRepurchase(ctx context.Context, req *bond.BondRepurchaseRequest) ([]bond.BondRepurchaseItem, error) {
	return bond.BondRepurchase(ctx, impl.client, req)
}

func (impl *bondImpl) Api272(ctx context.Context, req *bond.Api272Request) ([]bond.Api272Item, error) {
	return bond.Api272(ctx, impl.client, req)
}

func (impl *bondImpl) GlobalCalendar(ctx context.Context, req *bond.GlobalCalendarRequest) ([]bond.GlobalCalendarItem, error) {
	return bond.GlobalCalendar(ctx, impl.client, req)
}

func (impl *bondImpl) CbInterest(ctx context.Context, req *bond.CbInterestRequest) ([]bond.CbInterestItem, error) {
	return bond.CbInterest(ctx, impl.client, req)
}

func newBondImpl(client *sdk.Client) Bond {
	return &bondImpl{client: client}
}
