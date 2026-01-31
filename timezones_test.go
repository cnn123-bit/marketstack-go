package marketstack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTimezones(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/timezones" {
			t.Errorf("expected path '/timezones', got '%s'", r.URL.Path)
		}

		response := TimezonesResponse{
			Pagination: Pagination{
				Limit:  10,
				Offset: 0,
				Count:  3,
				Total:  3,
			},
			Data: []Timezone{
				{
					Timezone: "America/New_York",
					Abbr:     "EST",
					AbbrDST:  "EDT",
				},
				{
					Timezone: "Europe/London",
					Abbr:     "GMT",
					AbbrDST:  "BST",
				},
				{
					Timezone: "Asia/Tokyo",
					Abbr:     "JST",
					AbbrDST:  "JST",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient("test-key", nil)
	client.SetBaseURL(server.URL)

	result, err := client.GetTimezones(context.Background(), &TimezonesOptions{
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Data) != 3 {
		t.Fatalf("expected 3 timezones, got %d", len(result.Data))
	}

	if result.Data[0].Timezone != "America/New_York" {
		t.Errorf("expected timezone 'America/New_York', got '%s'", result.Data[0].Timezone)
	}

	if result.Data[1].Abbr != "GMT" {
		t.Errorf("expected abbr 'GMT', got '%s'", result.Data[1].Abbr)
	}

	if result.Data[2].AbbrDST != "JST" {
		t.Errorf("expected abbr_dst 'JST', got '%s'", result.Data[2].AbbrDST)
	}
}
