package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/spot"
)

type Spot interface {
	SpotBasic(ctx context.Context, req *spot.SpotBasicRequest) ([]spot.SpotBasicItem, error)
	SpotDaily(ctx context.Context, req *spot.SpotDailyRequest) ([]spot.SpotDailyItem, error)
}

type spotImpl struct {
	client *sdk.Client
}

func (impl *spotImpl) SpotBasic(ctx context.Context, req *spot.SpotBasicRequest) ([]spot.SpotBasicItem, error) {
	return spot.SpotBasic(ctx, impl.client, req)
}

func (impl *spotImpl) SpotDaily(ctx context.Context, req *spot.SpotDailyRequest) ([]spot.SpotDailyItem, error) {
	return spot.SpotDaily(ctx, impl.client, req)
}

func newSpotImpl(client *sdk.Client) Spot {
	return &spotImpl{client: client}
}
