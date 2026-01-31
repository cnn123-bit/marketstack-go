package marketstack

import (
	"context"
	"fmt"
)

type ExchangesOptions struct {
	Search string `url:"search,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

type ExchangesResponse struct {
	Pagination Pagination      `json:"pagination"`
	Data       []StockExchange `json:"data"`
}

type ExchangeResponse struct {
	Name        string `json:"name"`
	Acronym     string `json:"acronym"`
	MIC         string `json:"mic"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	Website     string `json:"website"`
	Timezone    struct {
		Timezone string `json:"timezone"`
		Abbr     string `json:"abbr"`
		AbbrDST  string `json:"abbr_dst"`
	} `json:"timezone"`
	Currency struct {
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"currency"`
}

func (c *Client) GetExchanges(ctx context.Context, opts *ExchangesOptions) (*ExchangesResponse, error) {
	var result ExchangesResponse
	if err := c.doRequest(ctx, "/exchanges", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetExchange(ctx context.Context, mic string) (*ExchangeResponse, error) {
	var result ExchangeResponse
	endpoint := fmt.Sprintf("/exchanges/%s", mic)
	if err := c.doRequest(ctx, endpoint, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
