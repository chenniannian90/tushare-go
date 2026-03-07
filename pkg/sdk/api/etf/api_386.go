package etf

import (
	"context"

	"github.com/chenniannian90/tushare-go/pkg/sdk"
)

type Api386Request struct {
}

type Api386Item struct {
}

func Api386(ctx context.Context, client *sdk.Client, req *Api386Request) ([]Api386Item, error) {
	params := map[string]interface{}{}
	fields := []string{}
	var result struct {
		Fields []string                 `json:"fields"`
		Items  []map[string]interface{} `json:"items"`
	}

	if err := client.CallAPI(ctx, "api_386", params, fields, &result); err != nil {
		return nil, err
	}
	return []Api386Item{}, nil
}
