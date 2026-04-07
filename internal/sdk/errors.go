package sdk

import (
	"fmt"
	"strings"
)

type APIError struct {
	Code       string
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	parts := []string{"qweather: api error"}
	if e.Code != "" {
		parts = append(parts, "code="+e.Code)
	}
	if e.StatusCode != 0 {
		parts = append(parts, fmt.Sprintf("http=%d", e.StatusCode))
	}
	if body := strings.TrimSpace(e.Body); body != "" {
		parts = append(parts, "body="+body)
	}
	return strings.Join(parts, " ")
}

func newHTTPError(statusCode int, body []byte) error {
	return &APIError{
		StatusCode: statusCode,
		Body:       strings.TrimSpace(string(body)),
	}
}
