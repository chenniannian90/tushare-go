package types

// APIResponse tushare api response
type APIResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      struct {
		Fields []string        `json:"fields"`
		Items  [][]interface{} `json:"items"`
	} `json:"data"`
}

// PostFunc represents the function to post data to API
type PostFunc func(body map[string]interface{}) (*APIResponse, error)

// TokenFunc represents the function to get API token
type TokenFunc func() string
