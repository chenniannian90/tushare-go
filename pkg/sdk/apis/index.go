package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/index"
)

type Index interface {
	IndexBasic(ctx context.Context, req *index.IndexBasicRequest) ([]index.IndexBasicItem, error)
	IndexDaily(ctx context.Context, req *index.IndexDailyRequest) ([]index.IndexDailyItem, error)
	IndexMember(ctx context.Context, req *index.IndexMemberRequest) ([]index.IndexMemberItem, error)
	IndexWeight(ctx context.Context, req *index.IndexWeightRequest) ([]index.IndexWeightItem, error)
	Api358(ctx context.Context, req *index.Api358Request) ([]index.Api358Item, error)
	IndexWeekly(ctx context.Context, req *index.IndexWeeklyRequest) ([]index.IndexWeeklyItem, error)
}

type indexImpl struct {
	client *sdk.Client
}

func (impl *indexImpl) IndexBasic(ctx context.Context, req *index.IndexBasicRequest) ([]index.IndexBasicItem, error) {
	return index.IndexBasic(ctx, impl.client, req)
}

func (impl *indexImpl) IndexDaily(ctx context.Context, req *index.IndexDailyRequest) ([]index.IndexDailyItem, error) {
	return index.IndexDaily(ctx, impl.client, req)
}

func (impl *indexImpl) IndexMember(ctx context.Context, req *index.IndexMemberRequest) ([]index.IndexMemberItem, error) {
	return index.IndexMember(ctx, impl.client, req)
}

func (impl *indexImpl) IndexWeight(ctx context.Context, req *index.IndexWeightRequest) ([]index.IndexWeightItem, error) {
	return index.IndexWeight(ctx, impl.client, req)
}

func (impl *indexImpl) Api358(ctx context.Context, req *index.Api358Request) ([]index.Api358Item, error) {
	return index.Api358(ctx, impl.client, req)
}

func (impl *indexImpl) IndexWeekly(ctx context.Context, req *index.IndexWeeklyRequest) ([]index.IndexWeeklyItem, error) {
	return index.IndexWeekly(ctx, impl.client, req)
}

func newIndexImpl(client *sdk.Client) Index {
	return &indexImpl{client: client}
}
