package sdk

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"syscall"
)

// ErrorCode 表示不同类型的错误
type ErrorCode string

const (
	// 认证错误
	ErrInvalidToken      ErrorCode = "INVALID_TOKEN"
	ErrTokenExpired      ErrorCode = "TOKEN_EXPIRED"
	ErrAccessDenied      ErrorCode = "ACCESS_DENIED"

	// 参数错误
	ErrInvalidParameter  ErrorCode = "INVALID_PARAMETER"
	ErrMissingParameter  ErrorCode = "MISSING_PARAMETER"

	// 速率限制错误
	ErrRateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"

	// 网络错误
	ErrNetworkTimeout     ErrorCode = "NETWORK_TIMEOUT"
	ErrConnectionFailed   ErrorCode = "CONNECTION_FAILED"
	ErrNetworkUnreachable ErrorCode = "NETWORK_UNREACHABLE"

	// 服务器错误
	ErrInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"

	// 数据错误
	ErrDataNotFound       ErrorCode = "DATA_NOT_FOUND"
	ErrInvalidResponse    ErrorCode = "INVALID_RESPONSE"

	// 上下文错误
	ErrContextCanceled    ErrorCode = "CONTEXT_CANCELED"
	ErrContextDeadline    ErrorCode = "CONTEXT_DEADLINE"
)

// APIError 表示 Tushare API 返回的增强错误
type APIError struct {
	Code       ErrorCode
	HTTPStatus int
	APICode    int    // 原始 API 错误代码
	Message    string
	Detail     string // 额外的错误详情
	RequestID  string // 用于跟踪的请求 ID
	Err        error  // 底层错误
}

func (e *APIError) Error() string {
	base := fmt.Sprintf("%s: %s", e.Code, e.Message)
	if e.Detail != "" {
		base += fmt.Sprintf(" (详情: %s)", e.Detail)
	}
	if e.RequestID != "" {
		base += fmt.Sprintf(" [请求ID: %s]", e.RequestID)
	}
	if e.Err != nil {
		base += fmt.Sprintf(": %v", e.Err)
	}
	return base
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// IsTemporary 如果错误是临时的且可以重试，则返回 true
func (e *APIError) IsTemporary() bool {
	switch e.Code { //nolint:exhaustive
	case ErrNetworkTimeout, ErrConnectionFailed, ErrRateLimitExceeded,
		ErrServiceUnavailable, ErrNetworkUnreachable:
		return true
	default:
		return false
	}
}

// IsPermanent 如果错误是永久的且不应重试，则返回 true
func (e *APIError) IsPermanent() bool {
	switch e.Code { //nolint:exhaustive
	case ErrInvalidToken, ErrTokenExpired, ErrAccessDenied,
		ErrInvalidParameter, ErrMissingParameter:
		return true
	default:
		return false
	}
}

// IsNetworkRelated 如果错误与网络相关，则返回 true
func (e *APIError) IsNetworkRelated() bool {
	switch e.Code { //nolint:exhaustive
	case ErrNetworkTimeout, ErrConnectionFailed, ErrNetworkUnreachable,
		ErrServiceUnavailable:
		return true
	default:
		return false
	}
}

// ShouldRetry 如果操作应该重试，则返回 true
func (e *APIError) ShouldRetry() bool {
	return e.IsTemporary() && !e.IsPermanent()
}

// WrapAPIError 使用 API 错误代码包装错误
func WrapAPIError(err error, message string) *APIError {
	return &APIError{
		Code:    ErrInternalError,
		Message: message,
		Err:     err,
	}
}

// WrapAPIErrorWithCode 使用特定错误代码包装错误
func WrapAPIErrorWithCode(code ErrorCode, message string, apiCode int, err error) *APIError {
	return &APIError{
		Code:    code,
		APICode: apiCode,
		Message: message,
		Err:     err,
	}
}

// NetworkError 表示增强的网络级别错误
type NetworkError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *NetworkError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// IsTemporary 如果错误是临时的，则返回 true
func (e *NetworkError) IsTemporary() bool {
	// 大多数网络错误都是临时的
	return true
}

// ShouldRetry 如果操作应该重试，则返回 true
func (e *NetworkError) ShouldRetry() bool {
	return e.IsTemporary()
}

// WrapNetworkError 使用网络错误信息包装错误
func WrapNetworkError(err error, message string) *NetworkError {
	code := ErrConnectionFailed

	// 检测特定错误类型
	if errors.Is(err, context.Canceled) {
		code = ErrContextCanceled
	} else if errors.Is(err, context.DeadlineExceeded) {
		code = ErrNetworkTimeout
	} else if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			code = ErrNetworkTimeout
		}
	} else {
		// 检查操作系统级别的错误
		var sysErr syscall.Errno
		if errors.As(err, &sysErr) {
			switch sysErr { //nolint:exhaustive
			case syscall.ECONNREFUSED:
				code = ErrConnectionFailed
			case syscall.EHOSTUNREACH:
				code = ErrNetworkUnreachable
			case syscall.ETIMEDOUT:
				code = ErrNetworkTimeout
			}
		}
	}

	return &NetworkError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ClassifyAPIError 将 Tushare API 错误代码分类为错误类型
func ClassifyAPIError(apiCode int, message string) ErrorCode {
	switch apiCode {
	case -2001:
		return ErrInvalidParameter
	case -2002:
		return ErrMissingParameter
	case -2003:
		return ErrInvalidToken
	case -2004:
		return ErrTokenExpired
	case -2005:
		return ErrAccessDenied
	case -2006:
		return ErrRateLimitExceeded
	case -2007:
		return ErrDataNotFound
	case 40101:
		return ErrInvalidToken
	case 40203:
		return ErrAccessDenied
	case 40204:
		return ErrRateLimitExceeded
	case 40205:
		return ErrTokenExpired
	case 500:
		return ErrInternalError
	case 503:
		return ErrServiceUnavailable
	default:
		// 尝试按消息分类
		msgLower := strings.ToLower(message)
		if strings.Contains(msgLower, "timeout") || strings.Contains(msgLower, "timed out") {
			return ErrNetworkTimeout
		}
		if strings.Contains(msgLower, "connection") || strings.Contains(msgLower, "connect") {
			return ErrConnectionFailed
		}
		if strings.Contains(msgLower, "rate limit") || strings.Contains(msgLower, "too many requests") {
			return ErrRateLimitExceeded
		}
		// 检查消息中的 token 相关错误
		if strings.Contains(msgLower, "token") {
			return ErrInvalidToken
		}

		return ErrInternalError
	}
}

// IsTemporaryError 如果错误是临时的且可以重试，则返回 true
func IsTemporaryError(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.IsTemporary()
	}

	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return netErr.IsTemporary()
	}

	// 检查常见的临时错误
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	// 检查网络错误
	var netErr2 net.Error
	if errors.As(err, &netErr2) {
		return netErr2.Timeout()
	}

	return false
}

// ShouldRetryOperation 如果操作应该重试，则返回 true
func ShouldRetryOperation(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.ShouldRetry()
	}

	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return netErr.ShouldRetry()
	}

	return IsTemporaryError(err)
}

// IsPermanentError 如果错误是永久的，则返回 true
func IsPermanentError(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.IsPermanent()
	}

	// 上下文已取消是永久的
	if errors.Is(err, context.Canceled) {
		return true
	}

	return false
}

// IsNetworkError 如果错误与网络相关，则返回 true
func IsNetworkError(err error) bool {
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return true
	}

	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.IsNetworkRelated()
	}

	// 检查 net.Error 接口
	var netErr2 net.Error
	if errors.As(err, &netErr2) {
		return true
	}

	// 检查连接相关的错误
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	return false
}
