package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/etf"
)

type ETF interface {
	Api127(ctx context.Context, req *etf.Api127Request) ([]etf.Api127Item, error)
	Api199(ctx context.Context, req *etf.Api199Request) ([]etf.Api199Item, error)
	Api385(ctx context.Context, req *etf.Api385Request) ([]etf.Api385Item, error)
	Api387(ctx context.Context, req *etf.Api387Request) ([]etf.Api387Item, error)
	Api400(ctx context.Context, req *etf.Api400Request) ([]etf.Api400Item, error)
	Api408(ctx context.Context, req *etf.Api408Request) ([]etf.Api408Item, error)
	Api416(ctx context.Context, req *etf.Api416Request) ([]etf.Api416Item, error)
}

type etfImpl struct {
	client *sdk.Client
}

func (impl *etfImpl) Api127(ctx context.Context, req *etf.Api127Request) ([]etf.Api127Item, error) {
	return etf.Api127(ctx, impl.client, req)
}

func (impl *etfImpl) Api199(ctx context.Context, req *etf.Api199Request) ([]etf.Api199Item, error) {
	return etf.Api199(ctx, impl.client, req)
}

func (impl *etfImpl) Api385(ctx context.Context, req *etf.Api385Request) ([]etf.Api385Item, error) {
	return etf.Api385(ctx, impl.client, req)
}

func (impl *etfImpl) Api387(ctx context.Context, req *etf.Api387Request) ([]etf.Api387Item, error) {
	return etf.Api387(ctx, impl.client, req)
}

func (impl *etfImpl) Api400(ctx context.Context, req *etf.Api400Request) ([]etf.Api400Item, error) {
	return etf.Api400(ctx, impl.client, req)
}

func (impl *etfImpl) Api408(ctx context.Context, req *etf.Api408Request) ([]etf.Api408Item, error) {
	return etf.Api408(ctx, impl.client, req)
}

func (impl *etfImpl) Api416(ctx context.Context, req *etf.Api416Request) ([]etf.Api416Item, error) {
	return etf.Api416(ctx, impl.client, req)
}

func newEtfImpl(client *sdk.Client) ETF {
	return &etfImpl{client: client}
}
