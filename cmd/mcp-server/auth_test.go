package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidToken(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		validTokens []string
		expected    bool
	}{
		{
			name:        "valid token",
			token:       "token123",
			validTokens: []string{"token123", "token456"},
			expected:    true,
		},
		{
			name:        "invalid token",
			token:       "invalid",
			validTokens: []string{"token123", "token456"},
			expected:    false,
		},
		{
			name:        "empty valid tokens list",
			token:       "any",
			validTokens: []string{},
			expected:    false,
		},
		{
			name:        "empty token",
			token:       "",
			validTokens: []string{"token123"},
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidToken(tt.token, tt.validTokens)
			if result != tt.expected {
				t.Errorf("isValidToken(%q, %v) = %v; want %v", tt.token, tt.validTokens, result, tt.expected)
			}
		})
	}
}

func TestIsValidTokenFormat(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  bool
	}{
		{
			name:  "Valid token length",
			token: "test-token-123456789012345678901234567890",
			want:  true,
		},
		{
			name:  "Token too short",
			token: "short",
			want:  false,
		},
		{
			name:  "Token exactly min length",
			token: "01234567890123456789012345678901", // 32 chars
			want:  true,
		},
		{
			name:  "Token with leading/trailing whitespace",
			token: "  test-token-123456789012345678901234567890  ",
			want:  true,
		},
		{
			name:  "Token only whitespace",
			token: "     ",
			want:  false,
		},
		{
			name:  "Empty token",
			token: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidTokenFormat(tt.token)
			if got != tt.want {
				t.Errorf("isValidTokenFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthMiddleware(t *testing.T) {
	validTokens := []string{"test-token-123456789012345678901234567890"}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := ExtractTokenFromRequest(r)
		if token != "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("authenticated")) // Don't expose token in response
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("no-token"))
		}
	})

	tests := []struct {
		name           string
		validTokens    []string
		headers        map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "valid token with X-API-Key header",
			validTokens: validTokens,
			headers: map[string]string{
				"X-API-Key": "test-token-123456789012345678901234567890",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "authenticated",
		},
		{
			name:        "Bearer token not supported",
			validTokens: validTokens,
			headers: map[string]string{
				"Authorization": "Bearer test-token-123456789012345678901234567890",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Authentication required\n",
		},
		{
			name:        "invalid token",
			validTokens: validTokens,
			headers: map[string]string{
				"X-API-Key": "invalid-token-123456789012345678901234567890",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Authentication required\n",
		},
		{
			name:           "missing token",
			validTokens:    validTokens,
			headers:        map[string]string{},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Authentication required\n",
		},
		{
			name:        "token too short (security check)",
			validTokens: validTokens,
			headers: map[string]string{
				"X-API-Key": "short",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Authentication required\n",
		},
		{
			name:           "no tokens configured - allow all",
			validTokens:    []string{},
			headers:        map[string]string{},
			expectedStatus: http.StatusOK,
			expectedBody:   "no-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := AuthMiddleware(tt.validTokens, nextHandler)

			req := httptest.NewRequest("GET", "/test", nil)
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			w := httptest.NewRecorder()
			middleware.ServeHTTP(w, req)

			// Check status
			if w.Code != tt.expectedStatus {
				t.Errorf("AuthMiddleware status = %d; want %d", w.Code, tt.expectedStatus)
			}

			// Check body
			body := w.Body.String()
			if body != tt.expectedBody {
				t.Errorf("AuthMiddleware body = %q; want %q", body, tt.expectedBody)
			}

			// Security check: ensure token is NOT in request headers
			if tt.name == "valid token with X-API-Key header" {
				if req.Header.Get("X-Validated-Token") != "" {
					t.Error("Security issue: token should not be stored in request headers")
				}
			}
		})
	}
}

func TestExtractTokenFromRequest(t *testing.T) {
	validToken := "test-token-123456789012345678901234567890"

	tests := []struct {
		name  string
		setup func(*http.Request)
		want  string
	}{
		{
			name: "Token in context",
			setup: func(r *http.Request) {
				ctx := context.WithValue(r.Context(), validatedTokenKey, validToken)
				*r = *r.WithContext(ctx)
			},
			want: validToken,
		},
		{
			name:  "No token in context",
			setup: func(r *http.Request) {},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/test", nil)
			tt.setup(req)

			got := ExtractTokenFromRequest(req)
			if got != tt.want {
				t.Errorf("ExtractTokenFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimingAttackProtection(t *testing.T) {
	// Test that token comparison uses constant-time comparison
	validTokens := []string{"test-token-123456789012345678901234567890"}

	// Test with various tokens to ensure comparison time is consistent
	testCases := []struct {
		name  string
		token string
	}{
		{"valid token", "test-token-123456789012345678901234567890"},
		{"invalid token - first char different", "Xest-token-123456789012345678901234567890"},
		{"invalid token - last char different", "test-token-12345678901234567890123456789X"},
		{"invalid token - middle different", "test-token-X23456789012345678901234567890"},
		{"completely different", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Just verify the function works correctly, timing tests require special setup
			result := isValidToken(tt.token, validTokens)
			expected := tt.token == "test-token-123456789012345678901234567890"
			if result != expected {
				t.Errorf("isValidToken() = %v, want %v", result, expected)
			}
		})
	}
}
