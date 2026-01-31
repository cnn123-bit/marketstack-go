package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("with API key", func(t *testing.T) {
		client := NewClient("test-key", nil)
		if client.apiKey != "test-key" {
			t.Errorf("expected API key to be 'test-key', got '%s'", client.apiKey)
		}
		if client.httpClient == nil {
			t.Error("expected httpClient to be set")
		}
	})

	t.Run("from environment variable", func(t *testing.T) {
		os.Setenv(envAPIKey, "env-test-key")
		defer os.Unsetenv(envAPIKey)

		client := NewClient("", nil)
		if client.apiKey != "env-test-key" {
			t.Errorf("expected API key from env to be 'env-test-key', got '%s'", client.apiKey)
		}
	})

	t.Run("with custom HTTP client", func(t *testing.T) {
		customClient := &http.Client{}
		client := NewClient("test-key", customClient)
		if client.httpClient != customClient {
			t.Error("expected custom HTTP client to be used")
		}
	})
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient("test-key", nil)
	client.SetBaseURL("https://custom.api.com/")
	if client.baseURL != "https://custom.api.com" {
		t.Errorf("expected base URL to be 'https://custom.api.com', got '%s'", client.baseURL)
	}
}

func TestDoRequestMissingAPIKey(t *testing.T) {
	client := NewClient("", nil)
	os.Unsetenv(envAPIKey)

	err := client.doRequest(context.Background(), "/test", nil, nil)
	if err == nil {
		t.Error("expected error for missing API key")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Errorf("expected APIError, got %T", err)
	}
	if apiErr.Code != "missing_api_key" {
		t.Errorf("expected error code 'missing_api_key', got '%s'", apiErr.Code)
	}
}

func TestDoRequestAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: &APIError{
				Code:    "invalid_access_key",
				Message: "You have not supplied a valid API Access Key.",
			},
		})
	}))
	defer server.Close()

	client := NewClient("invalid-key", nil)
	client.SetBaseURL(server.URL)

	var result map[string]interface{}
	err := client.doRequest(context.Background(), "/test", nil, &result)
	if err == nil {
		t.Error("expected error for invalid API key")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Errorf("expected APIError, got %T: %v", err, err)
	}
	if apiErr.Code != "invalid_access_key" {
		t.Errorf("expected error code 'invalid_access_key', got '%s'", apiErr.Code)
	}
}

func TestDoRequestSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("access_key") != "test-key" {
			t.Error("API key not found in query parameters")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
		})
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	var result map[string]string
	err := client.doRequest(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result["status"] != "success" {
		t.Errorf("expected status 'success', got '%s'", result["status"])
	}
}
