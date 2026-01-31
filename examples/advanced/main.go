package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tigusigalpa/marketstack-go"
)

func main() {
	apiKey := os.Getenv("MARKETSTACK_API_KEY")
	if apiKey == "" {
		log.Fatal("MARKETSTACK_API_KEY environment variable is required")
	}

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    90 * time.Second,
			DisableCompression: false,
		},
	}

	client := marketstack.NewClient(apiKey, httpClient)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("=== Marketstack Go Client - Advanced Example ===")

	fmt.Println("1. Getting Latest EOD Data for Tech Giants")
	eodData, err := client.GetEODLatest(ctx, &marketstack.EODOptions{
		Symbols: []string{"AAPL", "MSFT", "GOOG", "AMZN", "META"},
	})
	if err != nil {
		if apiErr, ok := err.(*marketstack.APIError); ok {
			log.Fatalf("API Error [%s]: %s", apiErr.Code, apiErr.Message)
		}
		log.Fatalf("Error: %v", err)
	}

	for _, stock := range eodData.Data {
		change := stock.Close - stock.Open
		changePercent := (change / stock.Open) * 100
		fmt.Printf("  %s: $%.2f (%.2f%%) | Volume: %.0fM\n",
			stock.Symbol, stock.Close, changePercent, stock.Volume/1000000)
	}

	fmt.Println("\n2. Searching for Exchanges")
	exchanges, err := client.GetExchanges(ctx, &marketstack.ExchangesOptions{
		Search: "New York",
		Limit:  3,
	})
	if err != nil {
		log.Fatalf("Error getting exchanges: %v", err)
	}

	for _, exchange := range exchanges.Data {
		fmt.Printf("  - %s (%s)\n", exchange.Name, exchange.MIC)
		fmt.Printf("    Location: %s, %s\n", exchange.City, exchange.Country)
	}

	fmt.Println("\n3. Getting Historical Data with Pagination")
	endDate := time.Now().Format("2006-01-02")
	startDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	historicalData, err := client.GetEOD(ctx, &marketstack.EODOptions{
		Symbols:  []string{"AAPL"},
		DateFrom: startDate,
		DateTo:   endDate,
		Sort:     "DESC",
		Limit:    10,
		Offset:   0,
	})
	if err != nil {
		log.Fatalf("Error getting historical data: %v", err)
	}

	fmt.Printf("  Showing %d of %d total records\n",
		historicalData.Pagination.Count, historicalData.Pagination.Total)

	var totalVolume float64
	for _, day := range historicalData.Data {
		totalVolume += day.Volume
		fmt.Printf("  %s: $%.2f | Range: $%.2f - $%.2f\n",
			day.Date, day.Close, day.Low, day.High)
	}

	avgVolume := totalVolume / float64(len(historicalData.Data))
	fmt.Printf("  Average Volume: %.0fM shares\n", avgVolume/1000000)

	fmt.Println("\n4. Getting Ticker Details")
	tickers, err := client.GetTickers(ctx, &marketstack.TickersOptions{
		Search: "tesla",
		Limit:  5,
	})
	if err != nil {
		log.Fatalf("Error searching tickers: %v", err)
	}

	for _, ticker := range tickers.Data {
		fmt.Printf("  %s - %s\n", ticker.Symbol, ticker.Name)
		fmt.Printf("    Exchange: %s | Intraday: %v | EOD: %v\n",
			ticker.StockExchange.Acronym, ticker.HasIntraday, ticker.HasEOD)
	}

	fmt.Println("\n5. Getting Supported Currencies")
	currencies, err := client.GetCurrencies(ctx, &marketstack.CurrenciesOptions{
		Limit: 5,
	})
	if err != nil {
		log.Fatalf("Error getting currencies: %v", err)
	}

	for _, currency := range currencies.Data {
		fmt.Printf("  %s (%s) - %s\n", currency.Code, currency.Symbol, currency.Name)
	}

	fmt.Println("\nExample completed successfully!")
}
