package marketstack

import "context"

type IntradayOptions struct {
	Symbols  []string `url:"symbols,omitempty,comma"`
	Exchange string   `url:"exchange,omitempty"`
	Sort     string   `url:"sort,omitempty"`
	DateFrom string   `url:"date_from,omitempty"`
	DateTo   string   `url:"date_to,omitempty"`
	Limit    int      `url:"limit,omitempty"`
	Offset   int      `url:"offset,omitempty"`
	Interval string   `url:"interval,omitempty"`
}

type IntradayResponse struct {
	Pagination Pagination     `json:"pagination"`
	Data       []IntradayData `json:"data"`
}

func (c *Client) GetIntraday(ctx context.Context, opts *IntradayOptions) (*IntradayResponse, error) {
	var result IntradayResponse
	if err := c.doRequest(ctx, "/intraday", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetIntradayLatest(ctx context.Context, opts *IntradayOptions) (*IntradayResponse, error) {
	var result IntradayResponse
	if err := c.doRequest(ctx, "/intraday/latest", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
