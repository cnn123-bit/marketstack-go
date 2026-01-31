package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tigusigalpa/marketstack-go"
)

func main() {
	apiKey := os.Getenv("MARKETSTACK_API_KEY")
	if apiKey == "" {
		log.Fatal("MARKETSTACK_API_KEY environment variable is required")
	}

	client := marketstack.NewClient(apiKey, nil)
	ctx := context.Background()

	fmt.Println("=== Marketstack Go Client - Basic Example ===")

	eodData, err := client.GetEOD(ctx, &marketstack.EODOptions{
		Symbols:  []string{"AAPL", "GOOG"},
		DateFrom: "2024-01-01",
		DateTo:   "2024-01-31",
		Sort:     "DESC",
		Limit:    5,
	})
	if err != nil {
		log.Fatalf("Error getting EOD data: %v", err)
	}

	fmt.Println("Recent Stock Prices:")
	for _, stock := range eodData.Data {
		fmt.Printf("  %s on %s: $%.2f (Volume: %.0f)\n",
			stock.Symbol, stock.Date, stock.Close, stock.Volume)
	}

	fmt.Println("\n=== Ticker Information ===")
	ticker, err := client.GetTicker(ctx, "TSLA")
	if err != nil {
		log.Fatalf("Error getting ticker info: %v", err)
	}

	fmt.Printf("Company: %s\n", ticker.Name)
	fmt.Printf("Symbol: %s\n", ticker.Symbol)
	fmt.Printf("Exchange: %s (%s)\n", ticker.StockExchange.Name, ticker.StockExchange.MIC)
	fmt.Printf("Country: %s\n", ticker.Country)
}
