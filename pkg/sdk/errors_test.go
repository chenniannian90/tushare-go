package sdk

import (
	"context"
	"errors"
	"testing"
)

func TestWrapAPIError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message string
	}{
		{
			name:    "API error with code 40203",
			err:     errors.New("40203"),
			message: "insufficient privileges",
		},
		{
			name:    "API error with code -2001",
			err:     errors.New("-2001"),
			message: "parameter error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapAPIError(tt.err, tt.message)
			if got.Message != tt.message {
				t.Errorf("WrapAPIError().Message = %v, want %v", got.Message, tt.message)
			}
			if got.Err != tt.err {
				t.Errorf("WrapAPIError().Err = %v, want %v", got.Err, tt.err)
			}
			if got.Code != ErrInternalError {
				t.Errorf("WrapAPIError().Code = %v, want %v", got.Code, ErrInternalError)
			}
		})
	}
}

func TestWrapAPIErrorWithCode(t *testing.T) {
	err := errors.New("test error")
	got := WrapAPIErrorWithCode(ErrAccessDenied, "access denied", 40203, err)

	if got.Code != ErrAccessDenied {
		t.Errorf("WrapAPIErrorWithCode().Code = %v, want %v", got.Code, ErrAccessDenied)
	}
	if got.APICode != 40203 {
		t.Errorf("WrapAPIErrorWithCode().APICode = %v, want %v", got.APICode, 40203)
	}
	if got.Message != "access denied" {
		t.Errorf("WrapAPIErrorWithCode().Message = %v, want %v", got.Message, "access denied")
	}
	if got.Err != err {
		t.Errorf("WrapAPIErrorWithCode().Err = %v, want %v", got.Err, err)
	}
}

func TestAPIError_Temporary(t *testing.T) {
	tests := []struct {
		name     string
		err      *APIError
		expected bool
	}{
		{
			name: "network timeout is temporary",
			err: &APIError{
				Code:    ErrNetworkTimeout,
				Message: "timeout",
			},
			expected: true,
		},
		{
			name: "access denied is permanent",
			err: &APIError{
				Code:    ErrAccessDenied,
				Message: "access denied",
			},
			expected: false,
		},
		{
			name: "rate limit exceeded is temporary",
			err: &APIError{
				Code:    ErrRateLimitExceeded,
				Message: "rate limit",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsTemporary(); got != tt.expected {
				t.Errorf("APIError.IsTemporary() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAPIError_ShouldRetry(t *testing.T) {
	tests := []struct {
		name     string
		err      *APIError
		expected bool
	}{
		{
			name: "network timeout should retry",
			err: &APIError{
				Code:    ErrNetworkTimeout,
				Message: "timeout",
			},
			expected: true,
		},
		{
			name: "access denied should not retry",
			err: &APIError{
				Code:    ErrAccessDenied,
				Message: "access denied",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.ShouldRetry(); got != tt.expected {
				t.Errorf("APIError.ShouldRetry() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestClassifyAPIError(t *testing.T) {
	tests := []struct {
		name     string
		apiCode  int
		message  string
		expected ErrorCode
	}{
		{
			name:     "parameter error",
			apiCode:  -2001,
			expected: ErrInvalidParameter,
		},
		{
			name:     "invalid token",
			apiCode:  -2003,
			expected: ErrInvalidToken,
		},
		{
			name:     "access denied",
			apiCode:  40203,
			expected: ErrAccessDenied,
		},
		{
			name:     "rate limit",
			apiCode:  40204,
			expected: ErrRateLimitExceeded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyAPIError(tt.apiCode, tt.message)
			if got != tt.expected {
				t.Errorf("ClassifyAPIError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWrapNetworkError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode ErrorCode
	}{
		{
			name:         "context canceled",
			err:          context.Canceled,
			expectedCode: ErrContextCanceled,
		},
		{
			name:         "context deadline exceeded",
			err:          context.DeadlineExceeded,
			expectedCode: ErrNetworkTimeout,
		},
		{
			name:         "generic error",
			err:          errors.New("some error"),
			expectedCode: ErrConnectionFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WrapNetworkError(tt.err, "test message")
			if got.Code != tt.expectedCode {
				t.Errorf("WrapNetworkError().Code = %v, want %v", got.Code, tt.expectedCode)
			}
		})
	}
}

func TestIsTemporaryError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "temporary API error",
			err: &APIError{
				Code:    ErrNetworkTimeout,
				Message: "timeout",
			},
			expected: true,
		},
		{
			name:     "permanent API error",
			err: &APIError{
				Code:    ErrAccessDenied,
				Message: "access denied",
			},
			expected: false,
		},
		{
			name:     "network error",
			err: &NetworkError{
				Code:    ErrConnectionFailed,
				Message: "connection failed",
			},
			expected: true,
		},
		{
			name:     "context deadline exceeded",
			err:      context.DeadlineExceeded,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTemporaryError(tt.err)
			if got != tt.expected {
				t.Errorf("IsTemporaryError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsNetworkError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "network error",
			err: &NetworkError{
				Code:    ErrConnectionFailed,
				Message: "connection failed",
			},
			expected: true,
		},
		{
			name:     "API network error",
			err: &APIError{
				Code:    ErrNetworkTimeout,
				Message: "timeout",
			},
			expected: true,
		},
		{
			name:     "non-network error",
			err: &APIError{
				Code:    ErrAccessDenied,
				Message: "access denied",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNetworkError(tt.err)
			if got != tt.expected {
				t.Errorf("IsNetworkError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestShouldRetryOperation(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "temporary error should retry",
			err: &APIError{
				Code:    ErrNetworkTimeout,
				Message: "timeout",
			},
			expected: true,
		},
		{
			name:     "permanent error should not retry",
			err: &APIError{
				Code:    ErrAccessDenied,
				Message: "access denied",
			},
			expected: false,
		},
		{
			name:     "network error should retry",
			err: &NetworkError{
				Code:    ErrConnectionFailed,
				Message: "connection failed",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShouldRetryOperation(tt.err)
			if got != tt.expected {
				t.Errorf("ShouldRetryOperation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNetworkError_Unwrap(t *testing.T) {
	originalErr := errors.New("connection failed")
	netErr := WrapNetworkError(originalErr, "request failed")

	if got := netErr.Unwrap(); got != originalErr {
		t.Errorf("NetworkError.Unwrap() = %v, want %v", got, originalErr)
	}
}

func TestAPIError_Unwrap(t *testing.T) {
	originalErr := errors.New("original error")
	apiErr := WrapAPIErrorWithCode(ErrAccessDenied, "access denied", 40203, originalErr)

	if got := apiErr.Unwrap(); got != originalErr {
		t.Errorf("APIError.Unwrap() = %v, want %v", got, originalErr)
	}
}

