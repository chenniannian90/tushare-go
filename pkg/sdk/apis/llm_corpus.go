package apis

import (
	"context"
	"github.com/chenniannian90/tushare-go/pkg/sdk"
	"github.com/chenniannian90/tushare-go/pkg/sdk/api/llm_corpus"
)

type LLMCorpus interface {
	Announcement(ctx context.Context, req *llm_corpus.AnnouncementRequest) ([]llm_corpus.AnnouncementItem, error)
	Api143(ctx context.Context, req *llm_corpus.Api143Request) ([]llm_corpus.Api143Item, error)
	Api195(ctx context.Context, req *llm_corpus.Api195Request) ([]llm_corpus.Api195Item, error)
	Einteraction(ctx context.Context, req *llm_corpus.EinteractionRequest) ([]llm_corpus.EinteractionItem, error)
	NewsBroadcast(ctx context.Context, req *llm_corpus.NewsBroadcastRequest) ([]llm_corpus.NewsBroadcastItem, error)
	Policy(ctx context.Context, req *llm_corpus.PolicyRequest) ([]llm_corpus.PolicyItem, error)
	ResearchReport(ctx context.Context, req *llm_corpus.ResearchReportRequest) ([]llm_corpus.ResearchReportItem, error)
}

type llmCorpusImpl struct {
	client *sdk.Client
}

func (impl *llmCorpusImpl) Announcement(ctx context.Context, req *llm_corpus.AnnouncementRequest) ([]llm_corpus.AnnouncementItem, error) {
	return llm_corpus.Announcement(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) Api143(ctx context.Context, req *llm_corpus.Api143Request) ([]llm_corpus.Api143Item, error) {
	return llm_corpus.Api143(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) Api195(ctx context.Context, req *llm_corpus.Api195Request) ([]llm_corpus.Api195Item, error) {
	return llm_corpus.Api195(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) Einteraction(ctx context.Context, req *llm_corpus.EinteractionRequest) ([]llm_corpus.EinteractionItem, error) {
	return llm_corpus.Einteraction(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) NewsBroadcast(ctx context.Context, req *llm_corpus.NewsBroadcastRequest) ([]llm_corpus.NewsBroadcastItem, error) {
	return llm_corpus.NewsBroadcast(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) Policy(ctx context.Context, req *llm_corpus.PolicyRequest) ([]llm_corpus.PolicyItem, error) {
	return llm_corpus.Policy(ctx, impl.client, req)
}

func (impl *llmCorpusImpl) ResearchReport(ctx context.Context, req *llm_corpus.ResearchReportRequest) ([]llm_corpus.ResearchReportItem, error) {
	return llm_corpus.ResearchReport(ctx, impl.client, req)
}

func newLLMCorpusImpl(client *sdk.Client) LLMCorpus {
	return &llmCorpusImpl{client: client}
}
