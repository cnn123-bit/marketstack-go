package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTickers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tickers" {
			t.Errorf("expected path '/tickers', got '%s'", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("search") != "apple" {
			t.Errorf("expected search 'apple', got '%s'", query.Get("search"))
		}

		response := TickersResponse{
			Pagination: Pagination{
				Limit:  10,
				Offset: 0,
				Count:  1,
				Total:  1,
			},
			Data: []Ticker{
				{
					Name:        "Apple Inc",
					Symbol:      "AAPL",
					HasIntraday: true,
					HasEOD:      true,
					Country:     "US",
					StockExchange: &StockExchange{
						Name:    "NASDAQ",
						Acronym: "NASDAQ",
						MIC:     "XNAS",
					},
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetTickers(context.Background(), &TickersOptions{
		Search: "apple",
		Limit:  10,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 ticker, got %d", len(result.Data))
	}

	if result.Data[0].Symbol != "AAPL" {
		t.Errorf("expected symbol 'AAPL', got '%s'", result.Data[0].Symbol)
	}

	if result.Data[0].Name != "Apple Inc" {
		t.Errorf("expected name 'Apple Inc', got '%s'", result.Data[0].Name)
	}
}

func TestGetTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tickers/TSLA" {
			t.Errorf("expected path '/tickers/TSLA', got '%s'", r.URL.Path)
		}

		response := TickerResponse{
			Name:        "Tesla Inc",
			Symbol:      "TSLA",
			HasIntraday: true,
			HasEOD:      true,
			Country:     "US",
			StockExchange: &StockExchange{
				Name:    "NASDAQ",
				Acronym: "NASDAQ",
				MIC:     "XNAS",
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetTicker(context.Background(), "TSLA")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Symbol != "TSLA" {
		t.Errorf("expected symbol 'TSLA', got '%s'", result.Symbol)
	}

	if result.Name != "Tesla Inc" {
		t.Errorf("expected name 'Tesla Inc', got '%s'", result.Name)
	}

	if result.StockExchange == nil {
		t.Fatal("expected stock exchange to be set")
	}

	if result.StockExchange.MIC != "XNAS" {
		t.Errorf("expected MIC 'XNAS', got '%s'", result.StockExchange.MIC)
	}
}
