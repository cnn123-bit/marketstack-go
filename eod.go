package marketstack

import "context"

type EODOptions struct {
	Symbols  []string `url:"symbols,omitempty,comma"`
	Exchange string   `url:"exchange,omitempty"`
	Sort     string   `url:"sort,omitempty"`
	DateFrom string   `url:"date_from,omitempty"`
	DateTo   string   `url:"date_to,omitempty"`
	Limit    int      `url:"limit,omitempty"`
	Offset   int      `url:"offset,omitempty"`
}

type EODResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []EODData  `json:"data"`
}

func (c *Client) GetEOD(ctx context.Context, opts *EODOptions) (*EODResponse, error) {
	var result EODResponse
	if err := c.doRequest(ctx, "/eod", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetEODLatest(ctx context.Context, opts *EODOptions) (*EODResponse, error) {
	var result EODResponse
	if err := c.doRequest(ctx, "/eod/latest", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetEODByDate(ctx context.Context, date string, opts *EODOptions) (*EODResponse, error) {
	var result EODResponse
	endpoint := "/eod/" + date
	if err := c.doRequest(ctx, endpoint, opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
