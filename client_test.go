package tushare

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	token := "test_token"
	api := New(token)

	if api == nil {
		t.Errorf("New() should return a non-nil TuShare instance")
	}
	if api.token != token {
		t.Errorf("New() token = %v, want %v", api.token, token)
	}
	if api.client == nil {
		t.Errorf("New() should initialize http.Client")
	}
}

func TestNewWithClient(t *testing.T) {
	token := "test_token"
	customClient := &http.Client{}
	api := NewWithClient(token, customClient)

	if api == nil {
		t.Errorf("NewWithClient() should return a non-nil TuShare instance")
	}
	if api.token != token {
		t.Errorf("NewWithClient() token = %v, want %v", api.token, token)
	}
	if api.client != customClient {
		t.Errorf("NewWithClient() client = %v, want %v", api.client, customClient)
	}
}

func TestToken(t *testing.T) {
	token := "test_token"
	api := New(token)

	if api.Token() != token {
		t.Errorf("Token() = %v, want %v", api.Token(), token)
	}
}

func TestSubClientsInitialized(t *testing.T) {
	api := New("test_token")

	// Check that all sub-clients are initialized
	if api.Basic == nil {
		t.Errorf("New() should initialize Basic client")
	}
	if api.Market == nil {
		t.Errorf("New() should initialize Market client")
	}
	if api.Finance == nil {
		t.Errorf("New() should initialize Finance client")
	}
	if api.Moneyflow == nil {
		t.Errorf("New() should initialize Moneyflow client")
	}
	if api.Margin == nil {
		t.Errorf("New() should initialize Margin client")
	}
	if api.Reference == nil {
		t.Errorf("New() should initialize Reference client")
	}
	if api.Special == nil {
		t.Errorf("New() should initialize Special client")
	}
	if api.Toplist == nil {
		t.Errorf("New() should initialize Toplist client")
	}
	if api.Index == nil {
		t.Errorf("New() should initialize Index client")
	}
	if api.Etf == nil {
		t.Errorf("New() should initialize Etf client")
	}
}

func TestEndpoint(t *testing.T) {
	if Endpoint == "" {
		t.Errorf("Endpoint should be defined")
	}
	expectedEndpoint := "http://api.tushare.pro"
	if Endpoint != expectedEndpoint {
		t.Errorf("Endpoint = %v, want %v", Endpoint, expectedEndpoint)
	}
}
