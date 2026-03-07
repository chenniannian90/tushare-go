package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	industrtmt "github.com/chenniannian90/tushare-go/pkg/sdk/api/industry/industry_tmt"
)

type Industry interface {
	MovieBoxoffice(ctx context.Context, req *industrtmt.MovieBoxofficeRequest) ([]industrtmt.MovieBoxofficeItem, error)
	MovieScript(ctx context.Context, req *industrtmt.MovieScriptRequest) ([]industrtmt.MovieScriptItem, error)
	TmtRevenue(ctx context.Context, req *industrtmt.TmtRevenueRequest) ([]industrtmt.TmtRevenueItem, error)
	TmtRevenueDetail(ctx context.Context, req *industrtmt.TmtRevenueDetailRequest) ([]industrtmt.TmtRevenueDetailItem, error)
}

type industryImpl struct {
	client *sdk.Client
}

func (impl *industryImpl) MovieBoxoffice(ctx context.Context, req *industrtmt.MovieBoxofficeRequest) ([]industrtmt.MovieBoxofficeItem, error) {
	return industrtmt.MovieBoxoffice(ctx, impl.client, req)
}

func (impl *industryImpl) MovieScript(ctx context.Context, req *industrtmt.MovieScriptRequest) ([]industrtmt.MovieScriptItem, error) {
	return industrtmt.MovieScript(ctx, impl.client, req)
}

func (impl *industryImpl) TmtRevenue(ctx context.Context, req *industrtmt.TmtRevenueRequest) ([]industrtmt.TmtRevenueItem, error) {
	return industrtmt.TmtRevenue(ctx, impl.client, req)
}

func (impl *industryImpl) TmtRevenueDetail(ctx context.Context, req *industrtmt.TmtRevenueDetailRequest) ([]industrtmt.TmtRevenueDetailItem, error) {
	return industrtmt.TmtRevenueDetail(ctx, impl.client, req)
}

func newIndustryImpl(client *sdk.Client) Industry {
	return &industryImpl{client: client}
}
