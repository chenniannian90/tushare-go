package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/options"
)

type Options interface {
	OptBasic(ctx context.Context, req *options.OptBasicRequest) ([]options.OptBasicItem, error)
	OptDaily(ctx context.Context, req *options.OptDailyRequest) ([]options.OptDailyItem, error)
	OptMin(ctx context.Context, req *options.OptMinRequest) ([]options.OptMinItem, error)
}

type optionsImpl struct {
	client *sdk.Client
}

func (impl *optionsImpl) OptBasic(ctx context.Context, req *options.OptBasicRequest) ([]options.OptBasicItem, error) {
	return options.OptBasic(ctx, impl.client, req)
}

func (impl *optionsImpl) OptDaily(ctx context.Context, req *options.OptDailyRequest) ([]options.OptDailyItem, error) {
	return options.OptDaily(ctx, impl.client, req)
}

func (impl *optionsImpl) OptMin(ctx context.Context, req *options.OptMinRequest) ([]options.OptMinItem, error) {
	return options.OptMin(ctx, impl.client, req)
}

func newOptionsImpl(client *sdk.Client) Options {
	return &optionsImpl{client: client}
}
