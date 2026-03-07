package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// HealthStatus represents the server health status
type HealthStatus struct {
	Status     string    `json:"status"`
	Uptime     string    `json:"uptime"`
	ToolCount  int       `json:"tool_count"`
	RequestCount int64   `json:"request_count"`
	ErrorCount  int64     `json:"error_count"`
	LastRequest string   `json:"last_request"`
}

// HealthManager manages server health and lifecycle
type HealthManager struct {
	startTime    time.Time
	mu           sync.RWMutex
	requestCount int64
	errorCount   int64
	lastRequest  time.Time
}

// NewHealthManager creates a new health manager
func NewHealthManager() *HealthManager {
	return &HealthManager{
		startTime: time.Now(),
	}
}

// GetStatus returns the current health status
func (h *HealthManager) GetStatus(toolCount int) HealthStatus {
	h.mu.RLock()
	defer h.mu.RUnlock()

	uptime := time.Since(h.startTime)
	var lastRequestStr string
	if !h.lastRequest.IsZero() {
		lastRequestStr = h.lastRequest.Format(time.RFC3339)
	}

	return HealthStatus{
		Status:      "healthy",
		Uptime:      uptime.String(),
		ToolCount:   toolCount,
		RequestCount: h.requestCount,
		ErrorCount:  h.errorCount,
		LastRequest: lastRequestStr,
	}
}

// RecordRequest records a request
func (h *HealthManager) RecordRequest() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.requestCount++
	h.lastRequest = time.Now()
}

// RecordError records an error
func (h *HealthManager) RecordError() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.errorCount++
}

// LifecycleManager manages server lifecycle
type LifecycleManager struct {
	ctx        context.Context
	cancel     context.CancelFunc
	shutdownMu sync.Mutex
	shutdown   bool
}

// NewLifecycleManager creates a new lifecycle manager
func NewLifecycleManager() *LifecycleManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &LifecycleManager{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Context returns the server context
func (lm *LifecycleManager) Context() context.Context {
	return lm.ctx
}

// Shutdown initiates graceful shutdown
func (lm *LifecycleManager) Shutdown(timeout time.Duration) error {
	lm.shutdownMu.Lock()
	defer lm.shutdownMu.Unlock()

	if lm.shutdown {
		return fmt.Errorf("already shutdown")
	}

	lm.shutdown = true
	log.Printf("Initiating graceful shutdown (timeout: %v)...", timeout)

	// Cancel the context
	lm.cancel()

	// Wait for cleanup or timeout
	done := make(chan struct{})
	go func() {
		// Perform cleanup here if needed
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		log.Println("Graceful shutdown completed")
		return nil
	case <-time.After(timeout):
		log.Println("Shutdown timeout exceeded, forcing exit")
		return fmt.Errorf("shutdown timeout")
	}
}

// IsShutdown checks if the server is shutting down
func (lm *LifecycleManager) IsShutdown() bool {
	lm.shutdownMu.Lock()
	defer lm.shutdownMu.Unlock()
	return lm.shutdown
}
