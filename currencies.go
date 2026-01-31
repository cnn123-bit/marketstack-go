package marketstack

import "context"

type CurrenciesOptions struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type CurrenciesResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Currency `json:"data"`
}

func (c *Client) GetCurrencies(ctx context.Context, opts *CurrenciesOptions) (*CurrenciesResponse, error) {
	var result CurrenciesResponse
	if err := c.doRequest(ctx, "/currencies", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
