//go:build integration

package sdk

import (
	"context"
	"os"
	"testing"
	"time"
)

// APIResult is a common result structure for API responses
type APIResult struct {
	Fields []string        `json:"fields"`
	Items  []interface{}   `json:"items"`
}

// TestIntegration_RealAPI_Calls tests real Tushare API calls
func TestIntegration_RealAPI_Calls(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	config, err := NewConfig(token)
	if err != nil {
		t.Fatalf("failed to create config: %v", err)
	}

	client := NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test stock_basic API
	t.Run("stock_basic", func(t *testing.T) {
		req := map[string]interface{}{"ts_code": "000001.SZ"}
		fields := []string{"ts_code", "symbol", "name", "area", "industry"}

		var result APIResult
		err := client.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err != nil {
			t.Fatalf("failed to call stock_basic API: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected at least one stock, got empty result")
		}

		t.Logf("Successfully retrieved stock data: %d items", len(result.Items))
		for _, item := range result.Items {
			if itemMap, ok := item.(map[string]interface{}); ok {
				t.Logf("  - %s: %s", itemMap["ts_code"], itemMap["name"])
			}
		}
	})

	// Test trade_cal API
	t.Run("trade_cal", func(t *testing.T) {
		req := map[string]interface{}{
			"exchange":  "SSE",
			"start_date": "20240101",
			"end_date":   "20240110",
		}
		fields := []string{"exchange", "cal_date", "is_open"}

		var result APIResult
		err := client.CallAPI(ctx, "trade_cal", req, fields, &result)
		if err != nil {
			t.Fatalf("failed to call trade_cal API: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected calendar data, got empty result")
		}

		t.Logf("Successfully retrieved calendar data: %d items", len(result.Items))
	})

	// Test daily API
	t.Run("daily", func(t *testing.T) {
		req := map[string]interface{}{
			"ts_code":    "000001.SZ",
			"start_date": "20240101",
			"end_date":   "20240105",
		}
		fields := []string{"ts_code", "trade_date", "open", "high", "low", "close"}

		var result APIResult
		err := client.CallAPI(ctx, "daily", req, fields, &result)
		if err != nil {
			t.Fatalf("failed to call daily API: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected daily data, got empty result")
		}

		t.Logf("Successfully retrieved daily data: %d items", len(result.Items))
	})

	// Test index_basic API
	t.Run("index_basic", func(t *testing.T) {
		req := map[string]interface{}{"market": "SSE"}
		fields := []string{"ts_code", "name", "market"}

		var result APIResult
		err := client.CallAPI(ctx, "index_basic", req, fields, &result)
		if err != nil {
			t.Fatalf("failed to call index_basic API: %v", err)
		}

		if len(result.Items) == 0 {
			t.Error("expected index data, got empty result")
		}

		t.Logf("Successfully retrieved index data: %d items", len(result.Items))
	})
}

// TestIntegration_ErrorHandling tests error handling with real API
func TestIntegration_ErrorHandling(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	_ = token // Use token
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("invalid_token", func(t *testing.T) {
		invalidConfig, _ := NewConfig("invalid_token_12345")
		invalidClient := NewClient(invalidConfig)

		req := map[string]interface{}{}
		fields := []string{"ts_code"}

		var result APIResult
		err := invalidClient.CallAPI(ctx, "stock_basic", req, fields, &result)
		if err == nil {
			t.Error("expected error with invalid token, got nil")
			return
		}

		apiErr, ok := err.(*APIError)
		if !ok {
			t.Errorf("expected APIError, got %T", err)
			return
		}

		t.Logf("Got expected error: %s (code: %s, apiCode: %d)",
			apiErr.Message, apiErr.Code, apiErr.APICode)

		if apiErr.Code != ErrInvalidToken && apiErr.Code != ErrAccessDenied {
			t.Errorf("expected token/auth error, got %s", apiErr.Code)
		}

		if !apiErr.IsPermanent() {
			t.Error("token error should be permanent")
		}

		if apiErr.ShouldRetry() {
			t.Error("token error should not be retried")
		}
	})
}

// TestIntegration_ErrorClassification tests error classification
func TestIntegration_ErrorClassification(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	testCases := []struct {
		name     string
		apiCode  int
		message  string
		expected ErrorCode
	}{
		{"parameter_error", -2001, "parameter error", ErrInvalidParameter},
		{"access_denied", 40203, "insufficient privileges", ErrAccessDenied},
		{"rate_limit", 40204, "rate limit exceeded", ErrRateLimitExceeded},
		{"server_error", 500, "internal server error", ErrInternalError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code := ClassifyAPIError(tc.apiCode, tc.message)
			if code != tc.expected {
				t.Errorf("ClassifyAPIError(%d, %s) = %v, want %v",
					tc.apiCode, tc.message, code, tc.expected)
			} else {
				t.Logf("✓ Error code %d correctly classified as %s", tc.apiCode, code)
			}
		})
	}
}

// TestIntegration_ConcurrentCalls tests concurrent API calls
func TestIntegration_ConcurrentCalls(t *testing.T) {
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		t.Skip("TUSHARE_TOKEN environment variable not set")
	}

	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	config, _ := NewConfig(token)
	client := NewClient(config)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	results := make(chan error, 3)

	// Concurrent calls
	go func() {
		req := map[string]interface{}{"ts_code": "000001.SZ"}
		fields := []string{"ts_code", "name"}
		var result APIResult
		results <- client.CallAPI(ctx, "stock_basic", req, fields, &result)
	}()

	go func() {
		req := map[string]interface{}{"exchange": "SSE", "start_date": "20240101", "end_date": "20240105"}
		fields := []string{"exchange", "cal_date", "is_open"}
		var result APIResult
		results <- client.CallAPI(ctx, "trade_cal", req, fields, &result)
	}()

	go func() {
		req := map[string]interface{}{"market": "SSE"}
		fields := []string{"ts_code", "name"}
		var result APIResult
		results <- client.CallAPI(ctx, "index_basic", req, fields, &result)
	}()

	// Collect results
	successCount := 0
	for i := 0; i < 3; i++ {
		if err := <-results; err != nil {
			t.Errorf("concurrent call %d failed: %v", i+1, err)
		} else {
			successCount++
		}
	}

	t.Logf("Concurrent API calls: %d out of 3 succeeded", successCount)

	if successCount != 3 {
		t.Errorf("expected all 3 calls to succeed, got %d", successCount)
	}
}
