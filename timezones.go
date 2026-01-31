package marketstack

import "context"

type TimezonesOptions struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type TimezonesResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Timezone `json:"data"`
}

func (c *Client) GetTimezones(ctx context.Context, opts *TimezonesOptions) (*TimezonesResponse, error) {
	var result TimezonesResponse
	if err := c.doRequest(ctx, "/timezones", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
