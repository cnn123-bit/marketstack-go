package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchanges(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/exchanges" {
			t.Errorf("expected path '/exchanges', got '%s'", r.URL.Path)
		}

		response := ExchangesResponse{
			Pagination: Pagination{
				Limit:  10,
				Offset: 0,
				Count:  2,
				Total:  2,
			},
			Data: []StockExchange{
				{
					Name:        "NASDAQ",
					Acronym:     "NASDAQ",
					MIC:         "XNAS",
					Country:     "USA",
					CountryCode: "US",
					City:        "New York",
				},
				{
					Name:        "New York Stock Exchange",
					Acronym:     "NYSE",
					MIC:         "XNYS",
					Country:     "USA",
					CountryCode: "US",
					City:        "New York",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetExchanges(context.Background(), &ExchangesOptions{
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 2 {
		t.Fatalf("expected 2 exchanges, got %d", len(result.Data))
	}

	if result.Data[0].MIC != "XNAS" {
		t.Errorf("expected MIC 'XNAS', got '%s'", result.Data[0].MIC)
	}

	if result.Data[1].Acronym != "NYSE" {
		t.Errorf("expected acronym 'NYSE', got '%s'", result.Data[1].Acronym)
	}
}

func TestGetExchange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/exchanges/XNAS" {
			t.Errorf("expected path '/exchanges/XNAS', got '%s'", r.URL.Path)
		}

		response := ExchangeResponse{
			Name:        "NASDAQ",
			Acronym:     "NASDAQ",
			MIC:         "XNAS",
			Country:     "USA",
			CountryCode: "US",
			City:        "New York",
			Website:     "www.nasdaq.com",
			Timezone: struct {
				Timezone string `json:"timezone"`
				Abbr     string `json:"abbr"`
				AbbrDST  string `json:"abbr_dst"`
			}{
				Timezone: "America/New_York",
				Abbr:     "EST",
				AbbrDST:  "EDT",
			},
			Currency: struct {
				Code   string `json:"code"`
				Symbol string `json:"symbol"`
				Name   string `json:"name"`
			}{
				Code:   "USD",
				Symbol: "$",
				Name:   "US Dollar",
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetExchange(context.Background(), "XNAS")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.MIC != "XNAS" {
		t.Errorf("expected MIC 'XNAS', got '%s'", result.MIC)
	}

	if result.Timezone.Timezone != "America/New_York" {
		t.Errorf("expected timezone 'America/New_York', got '%s'", result.Timezone.Timezone)
	}

	if result.Currency.Code != "USD" {
		t.Errorf("expected currency code 'USD', got '%s'", result.Currency.Code)
	}
}
