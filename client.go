package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/chenniannian90/tushare-go/stock"
	"github.com/chenniannian90/tushare-go/index"
	"github.com/chenniannian90/tushare-go/market"
	"github.com/chenniannian90/tushare-go/finance"
	"github.com/chenniannian90/tushare-go/hsgt"
	"github.com/chenniannian90/tushare-go/margin"
	"github.com/chenniannian90/tushare-go/toplist"
	"github.com/chenniannian90/tushare-go/holder"
	"github.com/chenniannian90/tushare-go/pledge"
	"github.com/chenniannian90/tushare-go/concept"
	"github.com/chenniannian90/tushare-go/ths"
	"github.com/chenniannian90/tushare-go/sw"
	"github.com/chenniannian90/tushare-go/limit"
	"github.com/chenniannian90/tushare-go/research"
	"github.com/chenniannian90/tushare-go/repurchase"
	"github.com/chenniannian90/tushare-go/realtime"
	"github.com/chenniannian90/tushare-go/fund"
	"github.com/chenniannian90/tushare-go/types"
)

// Endpoint URL
const Endpoint = "http://api.tushare.pro"

// TuShare instance
type TuShare struct {
	token  string
	client *http.Client

	// Sub-clients for different domains
	Stock      *stock.Client
	Index      *index.Client
	Market     *market.Client
	Finance    *finance.Client
	Hsgt       *hsgt.Client
	Margin     *margin.Client
	Toplist    *toplist.Client
	Holder     *holder.Client
	Pledge     *pledge.Client
	Concept    *concept.Client
	Ths        *ths.Client
	Sw         *sw.Client
	Limit      *limit.Client
	Research   *research.Client
	Repurchase *repurchase.Client
	Realtime   *realtime.Client
	Fund       *fund.Client
}

// New TuShare default client
func New(token string) *TuShare {
	return NewWithClient(token, http.DefaultClient)
}

// NewWithClient TuShare client with arguments
func NewWithClient(token string, httpClient *http.Client) *TuShare {
	api := &TuShare{
		token:  token,
		client: httpClient,
	}

	// Initialize sub-clients
	postFunc := api.PostData
	tokenFunc := api.Token

	api.Stock = stock.New(postFunc, tokenFunc)
	api.Index = index.New(postFunc, tokenFunc)
	api.Market = market.New(postFunc, tokenFunc)
	api.Finance = finance.New(postFunc, tokenFunc)
	api.Hsgt = hsgt.New(postFunc, tokenFunc)
	api.Margin = margin.New(postFunc, tokenFunc)
	api.Toplist = toplist.New(postFunc, tokenFunc)
	api.Holder = holder.New(postFunc, tokenFunc)
	api.Pledge = pledge.New(postFunc, tokenFunc)
	api.Concept = concept.New(postFunc, tokenFunc)
	api.Ths = ths.New(postFunc, tokenFunc)
	api.Sw = sw.New(postFunc, tokenFunc)
	api.Limit = limit.New(postFunc, tokenFunc)
	api.Research = research.New(postFunc, tokenFunc)
	api.Repurchase = repurchase.New(postFunc, tokenFunc)
	api.Realtime = realtime.New(postFunc, tokenFunc)
	api.Fund = fund.New(postFunc, tokenFunc)

	return api
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*types.APIResponse, error) {
	// Set http content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	//Handle network error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("oops! Network error")
	}


	// Read request
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check mime type of response
	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if mimeType != "application/json" {
		return nil, fmt.Errorf("could not execute request (%s)", fmt.Sprintf("response Content-Type is '%s', but should be 'application/json'.", mimeType))
	}

	// Parse Request
	var jsonData *types.APIResponse

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	// @TODO: handle API exception
	// Argument required
	if jsonData.Code == -2001 {
		return jsonData, fmt.Errorf("argument error: %s", jsonData.Msg)
	}

	// Permission deny
	if jsonData.Code == -2002 {
		return jsonData, fmt.Errorf("your point is not enough to use this api")
	}

	return jsonData, nil
}

// Token returns the API token
func (api *TuShare) Token() string {
	return api.token
}

// PostData sends a POST request to the API
func (api *TuShare) PostData(body map[string]interface{}) (*types.APIResponse, error) {
	req, err := api.request("POST", Endpoint, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
