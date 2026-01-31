package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIntraday(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/intraday" {
			t.Errorf("expected path '/intraday', got '%s'", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("symbols") != "AAPL" {
			t.Errorf("expected symbols 'AAPL', got '%s'", query.Get("symbols"))
		}
		if query.Get("interval") != "1min" {
			t.Errorf("expected interval '1min', got '%s'", query.Get("interval"))
		}

		response := IntradayResponse{
			Pagination: Pagination{
				Limit:  10,
				Offset: 0,
				Count:  2,
				Total:  2,
			},
			Data: []IntradayData{
				{
					Symbol: "AAPL",
					Last:   150.25,
					Date:   "2023-01-15T10:00:00+0000",
				},
				{
					Symbol: "AAPL",
					Last:   150.50,
					Date:   "2023-01-15T10:01:00+0000",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetIntraday(context.Background(), &IntradayOptions{
		Symbols:  []string{"AAPL"},
		Interval: "1min",
		Limit:    10,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Pagination.Count != 2 {
		t.Errorf("expected count 2, got %d", result.Pagination.Count)
	}

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 data items, got %d", len(result.Data))
	}

	if result.Data[0].Symbol != "AAPL" {
		t.Errorf("expected symbol 'AAPL', got '%s'", result.Data[0].Symbol)
	}
}

func TestGetIntradayLatest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/intraday/latest" {
			t.Errorf("expected path '/intraday/latest', got '%s'", r.URL.Path)
		}

		response := IntradayResponse{
			Pagination: Pagination{
				Limit:  1,
				Offset: 0,
				Count:  1,
				Total:  1,
			},
			Data: []IntradayData{
				{
					Symbol: "GOOG",
					Last:   2850.75,
					Date:   "2024-01-31T15:59:00+0000",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetIntradayLatest(context.Background(), &IntradayOptions{
		Symbols: []string{"GOOG"},
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 data item, got %d", len(result.Data))
	}

	if result.Data[0].Symbol != "GOOG" {
		t.Errorf("expected symbol 'GOOG', got '%s'", result.Data[0].Symbol)
	}
}
