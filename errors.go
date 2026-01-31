package marketstack

import "fmt"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Context struct {
		Symbol string `json:"symbol,omitempty"`
	} `json:"context,omitempty"`
}

func (e *APIError) Error() string {
	if e.Context.Symbol != "" {
		return fmt.Sprintf("marketstack API error [%s]: %s (symbol: %s)", e.Code, e.Message, e.Context.Symbol)
	}
	return fmt.Sprintf("marketstack API error [%s]: %s", e.Code, e.Message)
}

type ErrorResponse struct {
	Error *APIError `json:"error"`
}
