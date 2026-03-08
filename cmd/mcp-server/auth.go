package main

import (
	"context"
	"crypto/subtle"
	"net/http"
	"strings"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// validatedTokenKey is the context key for storing validated tokens
const validatedTokenKey contextKey = "validated_token"

// Security constants
const (
	// Minimum token length to prevent obviously invalid tokens
	minTokenLength = 32
	// Maximum token length to prevent potential DoS attacks
	maxTokenLength = 256
)

// AuthMiddleware creates an HTTP middleware that validates API tokens
func AuthMiddleware(validTokens []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If no tokens configured, allow all requests
		if len(validTokens) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		// Get token from X-API-Key header (only method)
		// Trim whitespace to prevent format issues
		token := strings.TrimSpace(r.Header.Get("X-API-Key"))

		// Check for missing token
		if token == "" {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Validate token format
		if !isValidTokenFormat(token) {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Validate token against allowed tokens
		if !isValidToken(token, validTokens) {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Token is valid, store it in request context for later use
		// This prevents token leakage in logs and is thread-safe
		ctx := context.WithValue(r.Context(), validatedTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isValidTokenFormat performs basic format validation of the token
func isValidTokenFormat(token string) bool {
	trimmed := strings.TrimSpace(token)
	length := len(trimmed)

	// Check token length is within reasonable bounds (after trimming)
	if length < minTokenLength || length > maxTokenLength {
		return false
	}

	// Validate character set - only allow printable ASCII characters (32-126)
	// Check original token before trimming to catch control characters
	for _, r := range token {
		if r < 32 || r > 126 { // ASCII printable characters range (32-126)
			return false
		}
	}

	return true
}

// isValidToken checks if the provided token is in the valid tokens list
// Uses constant-time comparison to prevent timing attacks
func isValidToken(token string, validTokens []string) bool {
	for _, validToken := range validTokens {
		// Use constant-time comparison to prevent timing attacks
		// This ensures that comparison time is independent of the input
		if subtle.ConstantTimeCompare([]byte(token), []byte(validToken)) == 1 {
			return true
		}
	}
	return false
}

// ExtractTokenFromRequest extracts the validated token from the request context
// This should be called after AuthMiddleware has validated the token
// Returns empty string if no valid token was found in context
func ExtractTokenFromRequest(r *http.Request) string {
	if token, ok := r.Context().Value(validatedTokenKey).(string); ok {
		return token
	}
	return ""
}
