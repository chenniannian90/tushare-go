package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/forex"
)

type Forex interface {
	Api178(ctx context.Context, req *forex.Api178Request) ([]forex.Api178Item, error)
	ForexDaily(ctx context.Context, req *forex.ForexDailyRequest) ([]forex.ForexDailyItem, error)
}

type forexImpl struct {
	client *sdk.Client
}

func (impl *forexImpl) Api178(ctx context.Context, req *forex.Api178Request) ([]forex.Api178Item, error) {
	return forex.Api178(ctx, impl.client, req)
}

func (impl *forexImpl) ForexDaily(ctx context.Context, req *forex.ForexDailyRequest) ([]forex.ForexDailyItem, error) {
	return forex.ForexDaily(ctx, impl.client, req)
}

func newForexImpl(client *sdk.Client) Forex {
	return &forexImpl{client: client}
}
