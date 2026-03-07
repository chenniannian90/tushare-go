package sdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_CallAPI(t *testing.T) {
	tests := []struct {
		name       string
		apiName    string
		params     map[string]interface{}
		fields     []string
		setupMock  func(*httptest.Server)
		wantErr    bool
		checkError func(error) bool
	}{
		{
			name:    "successful API call",
			apiName: "stock_basic",
			params:  map[string]interface{}{"limit": "1"},
			fields:  []string{"ts_code", "name"},
			setupMock: func(s *httptest.Server) {
				// Mock setup handled in test
			},
			wantErr: false,
		},
		{
			name:    "API error - insufficient privileges",
			apiName: "stock_basic",
			params:  map[string]interface{}{},
			fields:  []string{"ts_code"},
			setupMock: func(s *httptest.Server) {
				// Mock setup handled in test
			},
			wantErr: true,
			checkError: func(err error) bool {
				apiErr, ok := err.(*APIError)
				return ok && apiErr.Code == ErrAccessDenied && apiErr.APICode == 40203
			},
		},
		{
			name:    "network error - connection refused",
			apiName: "stock_basic",
			params:  map[string]interface{}{},
			fields:  []string{"ts_code"},
			setupMock: func(s *httptest.Server) {
				s.Close() // Close server to simulate connection error
			},
			wantErr: true,
			checkError: func(err error) bool {
				_, ok := err.(*NetworkError)
				return ok
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request method
				if r.Method != "POST" {
					t.Errorf("expected POST, got %s", r.Method)
				}

				// Verify content type
				if ct := r.Header.Get("Content-Type"); ct != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", ct)
				}

				// Parse request body
				var reqBody map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
					t.Errorf("failed to decode request body: %v", err)
					return
				}

				// Verify required fields
				if reqBody["api_name"] != tt.apiName {
					t.Errorf("expected api_name %s, got %v", tt.apiName, reqBody["api_name"])
				}
				if reqBody["token"] == nil {
					t.Error("token is required")
				}

				// Return mock response based on test case
				if tt.name == "API error - insufficient privileges" {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"code": 40203,
						"msg":  "insufficient privileges",
					})
				} else {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"code": 0,
						"msg":  "success",
						"data": map[string]interface{}{
							"fields": tt.fields,
							"items": []map[string]interface{}{
								{"ts_code": "000001.SZ", "name": "平安银行"},
							},
						},
					})
				}
			}))

			// Apply test-specific mock setup
			tt.setupMock(server)

			// Don't defer server.Close() for network error test (already closed)
			if tt.name != "network error - connection refused" {
				defer server.Close()
			}

			config, _ := NewConfig("test-token")
			config.Endpoint = server.URL
			client := NewClient(config)

			var result map[string]interface{}
			err := client.CallAPI(context.Background(), tt.apiName, tt.params, tt.fields, &result)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
					return
				}
				if tt.checkError != nil && !tt.checkError(err) {
					t.Errorf("error check failed: %v", err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestClient_CallAPI_RequestFormat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse request body to verify format
		var reqBody map[string]interface{}
		json.NewDecoder(r.Body).Decode(&reqBody)

		// Verify fields are comma-separated
		if fields, ok := reqBody["fields"].(string); ok {
			if !strings.Contains(fields, ",") {
				t.Errorf("expected comma-separated fields, got %s", fields)
			}
		} else {
			t.Error("fields should be a string")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"fields": []string{"ts_code", "name"},
				"items":  []map[string]interface{}{},
			},
		})
	}))
	defer server.Close()

	config, _ := NewConfig("test-token")
	config.Endpoint = server.URL
	client := NewClient(config)

	var result map[string]interface{}
	err := client.CallAPI(context.Background(), "stock_basic", map[string]interface{}{}, []string{"ts_code", "name"}, &result)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewClient(t *testing.T) {
	config, _ := NewConfig("test-token")
	client := NewClient(config)

	if client == nil {
		t.Error("NewClient() should not return nil")
	}

	if client.config != config {
		t.Error("NewClient() should store the config")
	}
}
