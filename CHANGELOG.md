# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-31

### Added
- Initial release of Marketstack Go Client
- Support for End-of-Day (EOD) data endpoints
  - `GetEOD()` - Get historical EOD data with filters
  - `GetEODLatest()` - Get latest EOD data
  - `GetEODByDate()` - Get EOD data for specific date
- Support for Intraday data endpoints
  - `GetIntraday()` - Get intraday data with intervals
  - `GetIntradayLatest()` - Get latest intraday data
- Support for Tickers endpoints
  - `GetTickers()` - Search and list tickers
  - `GetTicker()` - Get specific ticker information
- Support for Exchanges endpoints
  - `GetExchanges()` - List all exchanges
  - `GetExchange()` - Get specific exchange details
- Support for Currencies endpoint
  - `GetCurrencies()` - List supported currencies
- Support for Timezones endpoint
  - `GetTimezones()` - List supported timezones
- Context support for all API methods
- Custom error types for API-specific errors
- Environment variable support for API key (MARKETSTACK_API_KEY)
- Comprehensive unit tests with 75.8% coverage
- Complete documentation with examples
- MIT License

### Features
- Idiomatic Go API design
- Type-safe request options and responses
- Automatic query parameter encoding
- Custom HTTP client support
- Configurable base URL for testing
- Pagination support for all list endpoints
- Well-documented public APIs with godoc comments

[1.0.0]: https://github.com/tigusigalpa/marketstack-go/releases/tag/v1.0.0
