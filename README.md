# Marketstack Go Client

![Marketstack Golang SDK](https://github.com/user-attachments/assets/453558ea-34f3-48bd-912a-13fba90885a0)

[![Go Reference](https://pkg.go.dev/badge/github.com/tigusigalpa/marketstack-go.svg)](https://pkg.go.dev/github.com/tigusigalpa/marketstack-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/tigusigalpa/marketstack-go)](https://goreportcard.com/report/github.com/tigusigalpa/marketstack-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive, idiomatic Go client library for
the [Marketstack API](https://docs.apilayer.com/marketstack/docs/api-documentation). This library provides
easy access to real-time, intraday, and historical stock market data from 70+ global exchanges.

## Features

- ✅ **Idiomatic Go**: Clean, well-documented API following Go best practices
- ✅ **Context Support**: All methods accept `context.Context` for cancellation and timeouts
- ✅ **Type-Safe**: Strongly typed request options and responses
- ✅ **Comprehensive**: Supports all Marketstack API endpoints
- ✅ **Well-Tested**: Extensive unit tests with >90% coverage
- ✅ **Error Handling**: Custom error types for API-specific issues
- ✅ **Easy Configuration**: Support for environment variables and custom HTTP clients

## Installation

```bash
go get github.com/tigusigalpa/marketstack-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/tigusigalpa/marketstack-go"
)

func main() {
    // Create a new client (API key from environment variable)
    client := marketstack.NewClient("", nil)
    
    // Or explicitly pass the API key
    // client := marketstack.NewClient("your-api-key", nil)

    // Get end-of-day data
    eodData, err := client.GetEOD(context.Background(), &marketstack.EODOptions{
        Symbols:  []string{"AAPL", "GOOG"},
        DateFrom: "2024-01-01",
        DateTo:   "2024-01-31",
        Limit:    100,
    })
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    for _, stock := range eodData.Data {
        fmt.Printf("Symbol: %s, Close: $%.2f, Date: %s\n", 
            stock.Symbol, stock.Close, stock.Date)
    }
}
```

## Configuration

### API Key

You can provide your API key in two ways:

1. **Environment Variable** (recommended):

```bash
export MARKETSTACK_API_KEY="your-api-key-here"
```

2. **Direct Parameter**:

```go
client := marketstack.NewClient("your-api-key-here", nil)
```

### Custom HTTP Client

You can provide a custom `*http.Client` for advanced configuration:

```go
import (
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    90 * time.Second,
    },
}

client := marketstack.NewClient("your-api-key", httpClient)
```

### Custom Base URL

For testing or using a different API endpoint:

```go
client := marketstack.NewClient("your-api-key", nil)
client.SetBaseURL("https://custom-api.example.com/v1")
```

## API Methods

### End-of-Day Data

Get historical end-of-day stock market data.

```go
// Get EOD data with filters
eodData, err := client.GetEOD(ctx, &marketstack.EODOptions{
    Symbols:  []string{"AAPL", "MSFT", "GOOG"},
    Exchange: "XNAS",
    Sort:     "DESC",
    DateFrom: "2024-01-01",
    DateTo:   "2024-01-31",
    Limit:    100,
    Offset:   0,
})

// Get latest EOD data
latestData, err := client.GetEODLatest(ctx, &marketstack.EODOptions{
    Symbols: []string{"AAPL"},
})

// Get EOD data for a specific date
dateData, err := client.GetEODByDate(ctx, "2024-01-15", &marketstack.EODOptions{
    Symbols: []string{"TSLA"},
})
```

**Available Options:**

- `Symbols` - Array of stock symbols (e.g., `[]string{"AAPL", "GOOG"}`)
- `Exchange` - Filter by exchange MIC code (e.g., `"XNAS"`)
- `Sort` - Sort order: `"ASC"` or `"DESC"`
- `DateFrom` - Start date in `YYYY-MM-DD` format
- `DateTo` - End date in `YYYY-MM-DD` format
- `Limit` - Number of results (default: 100, max: 1000)
- `Offset` - Pagination offset

### Intraday Data

Get real-time and intraday stock market data.

```go
// Get intraday data
intradayData, err := client.GetIntraday(ctx, &marketstack.IntradayOptions{
    Symbols:  []string{"AAPL"},
    Interval: "1min", // 1min, 5min, 10min, 15min, 30min, 1hour
    DateFrom: "2024-01-31T09:30:00+0000",
    DateTo:   "2024-01-31T16:00:00+0000",
    Limit:    100,
})

// Get latest intraday data
latestIntraday, err := client.GetIntradayLatest(ctx, &marketstack.IntradayOptions{
    Symbols:  []string{"GOOG"},
    Interval: "5min",
})
```

**Available Options:**

- `Symbols` - Array of stock symbols
- `Exchange` - Filter by exchange MIC code
- `Interval` - Time interval: `"1min"`, `"5min"`, `"10min"`, `"15min"`, `"30min"`, `"1hour"`
- `Sort` - Sort order: `"ASC"` or `"DESC"`
- `DateFrom` - Start datetime in ISO 8601 format
- `DateTo` - End datetime in ISO 8601 format
- `Limit` - Number of results
- `Offset` - Pagination offset

### Tickers

Get information about stock tickers.

```go
// Search for tickers
tickers, err := client.GetTickers(ctx, &marketstack.TickersOptions{
    Search:   "apple",
    Exchange: "XNAS",
    Limit:    10,
})

for _, ticker := range tickers.Data {
    fmt.Printf("Symbol: %s, Name: %s, Exchange: %s\n",
        ticker.Symbol, ticker.Name, ticker.StockExchange.Name)
}

// Get specific ticker information
ticker, err := client.GetTicker(ctx, "AAPL")
fmt.Printf("Name: %s, Country: %s, Has Intraday: %v\n",
    ticker.Name, ticker.Country, ticker.HasIntraday)
```

**Available Options:**

- `Search` - Search query for ticker name or symbol
- `Exchange` - Filter by exchange MIC code
- `Limit` - Number of results
- `Offset` - Pagination offset

### Exchanges

Get information about stock exchanges.

```go
// Get all exchanges
exchanges, err := client.GetExchanges(ctx, &marketstack.ExchangesOptions{
    Search: "nasdaq",
    Limit:  10,
})

for _, exchange := range exchanges.Data {
    fmt.Printf("Name: %s, MIC: %s, Country: %s\n",
        exchange.Name, exchange.MIC, exchange.Country)
}

// Get specific exchange by MIC code
exchange, err := client.GetExchange(ctx, "XNAS")
fmt.Printf("Exchange: %s, City: %s, Timezone: %s\n",
    exchange.Name, exchange.City, exchange.Timezone.Timezone)
```

**Available Options:**

- `Search` - Search query for exchange name
- `Limit` - Number of results
- `Offset` - Pagination offset

### Currencies

Get supported currencies.

```go
currencies, err := client.GetCurrencies(ctx, &marketstack.CurrenciesOptions{
    Limit: 50,
})

for _, currency := range currencies.Data {
    fmt.Printf("Code: %s, Symbol: %s, Name: %s\n",
        currency.Code, currency.Symbol, currency.Name)
}
```

**Available Options:**

- `Limit` - Number of results
- `Offset` - Pagination offset

### Timezones

Get supported timezones.

```go
timezones, err := client.GetTimezones(ctx, &marketstack.TimezonesOptions{
    Limit: 50,
})

for _, tz := range timezones.Data {
    fmt.Printf("Timezone: %s, Abbr: %s, DST: %s\n",
        tz.Timezone, tz.Abbr, tz.AbbrDST)
}
```

**Available Options:**

- `Limit` - Number of results
- `Offset` - Pagination offset

## Error Handling

The library provides custom error types for API-specific errors:

```go
eodData, err := client.GetEOD(ctx, &marketstack.EODOptions{
    Symbols: []string{"INVALID"},
})

if err != nil {
    if apiErr, ok := err.(*marketstack.APIError); ok {
        fmt.Printf("API Error Code: %s\n", apiErr.Code)
        fmt.Printf("API Error Message: %s\n", apiErr.Message)
        
        // Check for specific error codes
        switch apiErr.Code {
        case "invalid_access_key":
            fmt.Println("Invalid API key provided")
        case "missing_access_key":
            fmt.Println("API key is required")
        case "404_not_found":
            fmt.Println("Ticker not found")
        default:
            fmt.Printf("API error: %v\n", apiErr)
        }
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

## Complete Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/tigusigalpa/marketstack-go"
)

func main() {
    // Create client with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    client := marketstack.NewClient("", nil)

    // Get latest EOD data for tech stocks
    fmt.Println("=== Latest End-of-Day Data ===")
    eodData, err := client.GetEODLatest(ctx, &marketstack.EODOptions{
        Symbols: []string{"AAPL", "MSFT", "GOOG", "AMZN"},
    })
    if err != nil {
        log.Fatalf("Error getting EOD data: %v", err)
    }

    for _, stock := range eodData.Data {
        fmt.Printf("%s: $%.2f (Volume: %.0f)\n",
            stock.Symbol, stock.Close, stock.Volume)
    }

    // Get ticker information
    fmt.Println("\n=== Ticker Information ===")
    ticker, err := client.GetTicker(ctx, "TSLA")
    if err != nil {
        log.Fatalf("Error getting ticker: %v", err)
    }

    fmt.Printf("Company: %s\n", ticker.Name)
    fmt.Printf("Symbol: %s\n", ticker.Symbol)
    fmt.Printf("Exchange: %s (%s)\n",
        ticker.StockExchange.Name, ticker.StockExchange.MIC)
    fmt.Printf("Country: %s\n", ticker.Country)
    fmt.Printf("Has Intraday: %v\n", ticker.HasIntraday)

    // Search for exchanges
    fmt.Println("\n=== US Exchanges ===")
    exchanges, err := client.GetExchanges(ctx, &marketstack.ExchangesOptions{
        Search: "US",
        Limit:  5,
    })
    if err != nil {
        log.Fatalf("Error getting exchanges: %v", err)
    }

    for _, exchange := range exchanges.Data {
        fmt.Printf("- %s (%s) in %s, %s\n",
            exchange.Name, exchange.MIC, exchange.City, exchange.Country)
    }

    // Get historical data with date range
    fmt.Println("\n=== Historical Data (Last 7 Days) ===")
    endDate := time.Now().Format("2006-01-02")
    startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

    historicalData, err := client.GetEOD(ctx, &marketstack.EODOptions{
        Symbols:  []string{"AAPL"},
        DateFrom: startDate,
        DateTo:   endDate,
        Sort:     "DESC",
        Limit:    10,
    })
    if err != nil {
        log.Fatalf("Error getting historical data: %v", err)
    }

    for _, day := range historicalData.Data {
        fmt.Printf("%s: Open $%.2f, Close $%.2f, High $%.2f, Low $%.2f\n",
            day.Date, day.Open, day.Close, day.High, day.Low)
    }
}
```

## Response Structures

All API responses include pagination information:

```go
type Pagination struct {
    Limit  int // Number of results per page
    Offset int // Current offset
    Count  int // Number of results in current response
    Total  int // Total number of results available
}
```

### EOD Data Structure

```go
type EODData struct {
    Open        float64
    High        float64
    Low         float64
    Close       float64
    Volume      float64
    AdjHigh     float64
    AdjLow      float64
    AdjClose    float64
    AdjOpen     float64
    AdjVolume   float64
    SplitFactor float64
    Dividend    float64
    Symbol      string
    Exchange    string
    Date        string
}
```

## Testing

Run the test suite:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -v -cover ./...
```

## Requirements

- Go 1.21 or higher
- Valid Marketstack API key (get one at [marketstack.com](https://marketstack.com/))

## API Documentation

For complete API documentation, visit:

- [Marketstack API Documentation](https://docs.apilayer.com/marketstack/docs/api-documentation)
- [Marketstack Website](https://marketstack.com/)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

**Igor Sazonov**

- Email: sovletig@gmail.com
- GitHub: [@tigusigalpa](https://github.com/tigusigalpa)

## Related Projects

- [marketstack-php](https://github.com/tigusigalpa/marketstack-php) - PHP/Laravel client for Marketstack API

## Disclaimer

This library is not officially associated with Marketstack or APILayer. Use at your own risk.
