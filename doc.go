// Package marketstack provides a comprehensive Go client for the Marketstack API.
//
// Marketstack is a REST API that provides real-time, intraday, and historical
// stock market data from 70+ global exchanges. This client library offers an
// idiomatic Go interface to interact with all Marketstack API endpoints.
//
// # Installation
//
//	go get github.com/tigusigalpa/marketstack-go
//
// # Quick Start
//
// Create a new client and fetch end-of-day data:
//
//	client := marketstack.NewClient("your-api-key", nil)
//	eodData, err := client.GetEOD(context.Background(), &marketstack.EODOptions{
//	    Symbols:  []string{"AAPL", "GOOG"},
//	    DateFrom: "2024-01-01",
//	    DateTo:   "2024-01-31",
//	    Limit:    100,
//	})
//
// # Authentication
//
// The API key can be provided in two ways:
//
// 1. Directly to the constructor:
//
//	client := marketstack.NewClient("your-api-key", nil)
//
// 2. Via the MARKETSTACK_API_KEY environment variable:
//
//	export MARKETSTACK_API_KEY="your-api-key"
//	client := marketstack.NewClient("", nil)
//
// # Context Support
//
// All API methods accept a context.Context parameter for cancellation and timeouts:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	data, err := client.GetEOD(ctx, opts)
//
// # Error Handling
//
// API errors are returned as *APIError types:
//
//	data, err := client.GetTicker(ctx, "INVALID")
//	if err != nil {
//	    if apiErr, ok := err.(*marketstack.APIError); ok {
//	        fmt.Printf("API Error: %s - %s\n", apiErr.Code, apiErr.Message)
//	    }
//	}
//
// # Available Endpoints
//
// The client supports all major Marketstack API endpoints:
//
//   - End-of-Day Data: GetEOD, GetEODLatest, GetEODByDate
//   - Intraday Data: GetIntraday, GetIntradayLatest
//   - Tickers: GetTickers, GetTicker
//   - Exchanges: GetExchanges, GetExchange
//   - Currencies: GetCurrencies
//   - Timezones: GetTimezones
//
// For complete documentation, visit: https://docs.apilayer.com/marketstack
package marketstack
