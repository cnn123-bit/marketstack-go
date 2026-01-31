# Project Structure

This document describes the organization and structure of the Marketstack Go Client library.

## Directory Layout

```
marketstack-go/
├── client.go              # Core client implementation
├── errors.go              # Custom error types
├── types.go               # Common data structures
├── doc.go                 # Package documentation
├── eod.go                 # End-of-Day data endpoints
├── intraday.go            # Intraday data endpoints
├── tickers.go             # Tickers endpoints
├── exchanges.go           # Exchanges endpoints
├── currencies.go          # Currencies endpoint
├── timezones.go           # Timezones endpoint
├── client_test.go         # Client tests
├── eod_test.go            # EOD endpoint tests
├── intraday_test.go       # Intraday endpoint tests
├── tickers_test.go        # Tickers endpoint tests
├── exchanges_test.go      # Exchanges endpoint tests
├── currencies_test.go     # Currencies endpoint tests
├── timezones_test.go      # Timezones endpoint tests
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── README.md              # Main documentation
├── LICENSE                # MIT License
├── CONTRIBUTING.md        # Contribution guidelines
├── CHANGELOG.md           # Version history
├── .gitignore             # Git ignore rules
└── examples/
    ├── basic/
    │   └── main.go        # Basic usage example
    └── advanced/
        └── main.go        # Advanced usage example
```

## Core Components

### client.go
- `Client` struct - Main client type
- `NewClient()` - Constructor function
- `SetBaseURL()` - Base URL configuration
- `doRequest()` - Internal HTTP request handler

### errors.go
- `APIError` - Custom error type for API errors
- `ErrorResponse` - API error response wrapper

### types.go
Common data structures used across endpoints:
- `Pagination` - Pagination metadata
- `EODData` - End-of-day stock data
- `IntradayData` - Intraday stock data
- `Ticker` - Stock ticker information
- `StockExchange` - Exchange information
- `Currency` - Currency information
- `Timezone` - Timezone information

## Endpoint Files

Each endpoint is implemented in its own file:

### eod.go
- `EODOptions` - Request parameters
- `EODResponse` - Response structure
- `GetEOD()` - Get historical EOD data
- `GetEODLatest()` - Get latest EOD data
- `GetEODByDate()` - Get EOD data for specific date

### intraday.go
- `IntradayOptions` - Request parameters
- `IntradayResponse` - Response structure
- `GetIntraday()` - Get intraday data
- `GetIntradayLatest()` - Get latest intraday data

### tickers.go
- `TickersOptions` - Request parameters
- `TickersResponse` - Response structure
- `TickerResponse` - Single ticker response
- `GetTickers()` - Search tickers
- `GetTicker()` - Get specific ticker

### exchanges.go
- `ExchangesOptions` - Request parameters
- `ExchangesResponse` - Response structure
- `ExchangeResponse` - Single exchange response
- `GetExchanges()` - List exchanges
- `GetExchange()` - Get specific exchange

### currencies.go
- `CurrenciesOptions` - Request parameters
- `CurrenciesResponse` - Response structure
- `GetCurrencies()` - List currencies

### timezones.go
- `TimezonesOptions` - Request parameters
- `TimezonesResponse` - Response structure
- `GetTimezones()` - List timezones

## Test Files

Each endpoint has corresponding test files using `httptest` to mock API responses:
- Verify correct URL construction
- Test query parameter encoding
- Validate response deserialization
- Check error handling
- Test edge cases

## Examples

### basic/main.go
Simple example demonstrating:
- Client initialization
- Getting EOD data
- Getting ticker information

### advanced/main.go
Comprehensive example showing:
- Custom HTTP client configuration
- Context with timeout
- Error handling with type assertions
- Multiple endpoint usage
- Pagination handling
- Data processing

## Dependencies

- `github.com/google/go-querystring` - Query parameter encoding
- Standard library packages:
  - `context` - Request cancellation
  - `encoding/json` - JSON serialization
  - `net/http` - HTTP client
  - `net/http/httptest` - Testing utilities

## Design Principles

1. **Idiomatic Go** - Follows Go conventions and best practices
2. **Type Safety** - Strongly typed throughout
3. **Context Support** - All methods accept context.Context
4. **Error Handling** - Custom error types for API errors
5. **Testability** - Comprehensive test coverage using httptest
6. **Documentation** - All public APIs documented with godoc
7. **Extensibility** - Easy to add new endpoints
8. **Simplicity** - Clean, straightforward API design

## Adding New Endpoints

To add a new endpoint:

1. Create `newfeature.go` with:
   - Options struct with `url` tags
   - Response struct with `json` tags
   - Client method(s)

2. Create `newfeature_test.go` with:
   - Mock server using httptest
   - Test cases for success and error scenarios
   - Verify request parameters and response parsing

3. Update documentation:
   - Add section to README.md
   - Include usage examples
   - Update CHANGELOG.md

## Code Style

- Formatted with `gofmt`
- Linted with `go vet`
- Documented with godoc comments
- Tested with standard `testing` package
- 75.8% test coverage

## Version Control

- Git repository: https://github.com/tigusigalpa/marketstack-go
- Semantic versioning (SemVer)
- Changelog maintained in CHANGELOG.md
