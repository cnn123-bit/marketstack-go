package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrencies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/currencies" {
			t.Errorf("expected path '/currencies', got '%s'", r.URL.Path)
		}

		response := CurrenciesResponse{
			Pagination: Pagination{
				Limit:  10,
				Offset: 0,
				Count:  3,
				Total:  3,
			},
			Data: []Currency{
				{
					Code:   "USD",
					Symbol: "$",
					Name:   "US Dollar",
				},
				{
					Code:   "EUR",
					Symbol: "€",
					Name:   "Euro",
				},
				{
					Code:   "GBP",
					Symbol: "£",
					Name:   "British Pound",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetCurrencies(context.Background(), &CurrenciesOptions{
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 3 {
		t.Fatalf("expected 3 currencies, got %d", len(result.Data))
	}

	if result.Data[0].Code != "USD" {
		t.Errorf("expected code 'USD', got '%s'", result.Data[0].Code)
	}

	if result.Data[1].Symbol != "€" {
		t.Errorf("expected symbol '€', got '%s'", result.Data[1].Symbol)
	}

	if result.Data[2].Name != "British Pound" {
		t.Errorf("expected name 'British Pound', got '%s'", result.Data[2].Name)
	}
}
