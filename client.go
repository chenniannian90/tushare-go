package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/chenniannian90/tushare-go/etf"
	"github.com/chenniannian90/tushare-go/index"
	"github.com/chenniannian90/tushare-go/stock/basic"
	"github.com/chenniannian90/tushare-go/stock/finance"
	"github.com/chenniannian90/tushare-go/stock/margin"
	"github.com/chenniannian90/tushare-go/stock/market"
	"github.com/chenniannian90/tushare-go/stock/moneyflow"
	"github.com/chenniannian90/tushare-go/stock/reference"
	"github.com/chenniannian90/tushare-go/stock/special"
	"github.com/chenniannian90/tushare-go/stock/toplist"
	"github.com/chenniannian90/tushare-go/types"
)

// Endpoint URL
const Endpoint = "http://api.tushare.pro"

// TuShare instance
type TuShare struct {
	token  string
	client *http.Client

	// Sub-clients for different domains
	StockBasic     *basic.Client
	StockMarket    *market.Client
	StockFinance   *finance.Client
	StockMoneyflow *moneyflow.Client
	StockMargin    *margin.Client
	StockReference *reference.Client
	StockSpecial   *special.Client
	StockToplist   *toplist.Client
	Index          *index.Client
	Etf            *etf.Client
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

	api.StockBasic = basic.New(postFunc, tokenFunc)
	api.StockMarket = market.New(postFunc, tokenFunc)
	api.StockFinance = finance.New(postFunc, tokenFunc)
	api.StockMoneyflow = moneyflow.New(postFunc, tokenFunc)
	api.StockMargin = margin.New(postFunc, tokenFunc)
	api.StockReference = reference.New(postFunc, tokenFunc)
	api.StockSpecial = special.New(postFunc, tokenFunc)
	api.StockToplist = toplist.New(postFunc, tokenFunc)
	api.Index = index.New(postFunc, tokenFunc)
	api.Etf = etf.New(postFunc, tokenFunc)

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
