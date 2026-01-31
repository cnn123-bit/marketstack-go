package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEOD(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/eod" {
			t.Errorf("expected path '/eod', got '%s'", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("symbols") != "AAPL,GOOG" {
			t.Errorf("expected symbols 'AAPL,GOOG', got '%s'", query.Get("symbols"))
		}
		if query.Get("date_from") != "2023-01-01" {
			t.Errorf("expected date_from '2023-01-01', got '%s'", query.Get("date_from"))
		}
		if query.Get("limit") != "100" {
			t.Errorf("expected limit '100', got '%s'", query.Get("limit"))
		}

		response := EODResponse{
			Pagination: Pagination{
				Limit:  100,
				Offset: 0,
				Count:  2,
				Total:  2,
			},
			Data: []EODData{
				{
					Symbol: "AAPL",
					Close:  150.25,
					Date:   "2023-01-15",
				},
				{
					Symbol: "GOOG",
					Close:  2800.50,
					Date:   "2023-01-15",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetEOD(context.Background(), &EODOptions{
		Symbols:  []string{"AAPL", "GOOG"},
		DateFrom: "2023-01-01",
		Limit:    100,
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
		t.Errorf("expected first symbol 'AAPL', got '%s'", result.Data[0].Symbol)
	}

	if result.Data[0].Close != 150.25 {
		t.Errorf("expected first close 150.25, got %.2f", result.Data[0].Close)
	}
}

func TestGetEODLatest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/eod/latest" {
			t.Errorf("expected path '/eod/latest', got '%s'", r.URL.Path)
		}

		response := EODResponse{
			Pagination: Pagination{
				Limit:  1,
				Offset: 0,
				Count:  1,
				Total:  1,
			},
			Data: []EODData{
				{
					Symbol: "AAPL",
					Close:  175.50,
					Date:   "2024-01-31",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetEODLatest(context.Background(), &EODOptions{
		Symbols: []string{"AAPL"},
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 data item, got %d", len(result.Data))
	}

	if result.Data[0].Symbol != "AAPL" {
		t.Errorf("expected symbol 'AAPL', got '%s'", result.Data[0].Symbol)
	}
}

func TestGetEODByDate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/eod/2023-01-15" {
			t.Errorf("expected path '/eod/2023-01-15', got '%s'", r.URL.Path)
		}

		response := EODResponse{
			Pagination: Pagination{
				Limit:  100,
				Offset: 0,
				Count:  1,
				Total:  1,
			},
			Data: []EODData{
				{
					Symbol: "TSLA",
					Close:  185.75,
					Date:   "2023-01-15",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetEODByDate(context.Background(), "2023-01-15", &EODOptions{
		Symbols: []string{"TSLA"},
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 1 {
		t.Fatalf("expected 1 data item, got %d", len(result.Data))
	}

	if result.Data[0].Date != "2023-01-15" {
		t.Errorf("expected date '2023-01-15', got '%s'", result.Data[0].Date)
	}
}
