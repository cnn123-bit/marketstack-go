package marketstack

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "http://api.marketstack.com/v1"
	envAPIKey      = "MARKETSTACK_API_KEY"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewClient(apiKey string, httpClient *http.Client) *Client {
	if apiKey == "" {
		apiKey = os.Getenv(envAPIKey)
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    defaultBaseURL,
		httpClient: httpClient,
	}
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = strings.TrimSuffix(baseURL, "/")
}

func (c *Client) doRequest(ctx context.Context, endpoint string, params interface{}, result interface{}) error {
	if c.apiKey == "" {
		return &APIError{
			Code:    "missing_api_key",
			Message: "API key is required. Set it via NewClient or MARKETSTACK_API_KEY environment variable",
		}
	}

	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	q := u.Query()
	q.Set("access_key", c.apiKey)

	if params != nil {
		values, err := query.Values(params)
		if err != nil {
			return fmt.Errorf("failed to encode query parameters: %w", err)
		}
		for k, v := range values {
			for _, val := range v {
				q.Add(k, val)
			}
		}
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr ErrorResponse
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
		}
		if apiErr.Error != nil {
			return apiErr.Error
		}
		return fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
