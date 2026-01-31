package marketstack

import (
	"context"
	"fmt"
)

type TickersOptions struct {
	Search   string `url:"search,omitempty"`
	Exchange string `url:"exchange,omitempty"`
	Limit    int    `url:"limit,omitempty"`
	Offset   int    `url:"offset,omitempty"`
}

type TickersResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Ticker   `json:"data"`
}

type TickerResponse struct {
	Name          string         `json:"name"`
	Symbol        string         `json:"symbol"`
	HasIntraday   bool           `json:"has_intraday"`
	HasEOD        bool           `json:"has_eod"`
	Country       string         `json:"country"`
	StockExchange *StockExchange `json:"stock_exchange"`
}

func (c *Client) GetTickers(ctx context.Context, opts *TickersOptions) (*TickersResponse, error) {
	var result TickersResponse
	if err := c.doRequest(ctx, "/tickers", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTicker(ctx context.Context, symbol string) (*TickerResponse, error) {
	var result TickerResponse
	endpoint := fmt.Sprintf("/tickers/%s", symbol)
	if err := c.doRequest(ctx, endpoint, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
