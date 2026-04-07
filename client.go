package qweather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultUserAgent = "qweather-go-sdk/0.1.0"

type Config struct {
	Host             string
	HTTPClient       *http.Client
	UserAgent        string
	APIKey           string
	APIKeyInQuery    bool
	JWTTokenProvider TokenProvider
}

type TokenProvider interface {
	Token(context.Context) (string, error)
}

type Client struct {
	baseURL       *url.URL
	httpClient    *http.Client
	userAgent     string
	apiKey        string
	apiKeyInQuery bool
	tokenProvider TokenProvider
}

func NewClient(cfg Config) (*Client, error) {
	if strings.TrimSpace(cfg.Host) == "" {
		return nil, fmt.Errorf("qweather: host is required")
	}

	if strings.TrimSpace(cfg.APIKey) == "" && cfg.JWTTokenProvider == nil {
		return nil, fmt.Errorf("qweather: either api key or jwt token provider is required")
	}

	baseURL, err := normalizeHost(cfg.Host)
	if err != nil {
		return nil, err
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 15 * time.Second}
	}

	userAgent := strings.TrimSpace(cfg.UserAgent)
	if userAgent == "" {
		userAgent = defaultUserAgent
	}

	return &Client{
		baseURL:       baseURL,
		httpClient:    httpClient,
		userAgent:     userAgent,
		apiKey:        strings.TrimSpace(cfg.APIKey),
		apiKeyInQuery: cfg.APIKeyInQuery,
		tokenProvider: cfg.JWTTokenProvider,
	}, nil
}

func normalizeHost(host string) (*url.URL, error) {
	host = strings.TrimSpace(host)
	if host == "" {
		return nil, fmt.Errorf("qweather: empty host")
	}

	if !strings.Contains(host, "://") {
		host = "https://" + host
	}

	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("qweather: parse host: %w", err)
	}

	if u.Scheme == "" || u.Host == "" {
		return nil, fmt.Errorf("qweather: invalid host %q", host)
	}

	u.Path = strings.TrimRight(u.Path, "/")
	u.RawQuery = ""
	u.Fragment = ""
	return u, nil
}

func (c *Client) newRequest(ctx context.Context, method, path string, query map[string]string) (*http.Request, error) {
	u := *c.baseURL
	u.Path = strings.TrimRight(c.baseURL.Path, "/") + "/" + strings.TrimLeft(path, "/")

	values := u.Query()
	for key, value := range query {
		if strings.TrimSpace(value) == "" {
			continue
		}
		values.Set(key, value)
	}

	if c.apiKey != "" && c.apiKeyInQuery {
		values.Set("key", c.apiKey)
	}

	u.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("qweather: create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	if c.apiKey != "" && !c.apiKeyInQuery {
		req.Header.Set("X-QW-Api-Key", c.apiKey)
	}

	if c.tokenProvider != nil {
		token, err := c.tokenProvider.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("qweather: resolve jwt token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
	}

	return req, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, query map[string]string, out any) error {
	req, err := c.newRequest(ctx, method, path, query)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("qweather: send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("qweather: read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return newHTTPError(resp.StatusCode, body)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("qweather: decode response: %w", err)
	}

	if coded, ok := out.(codedResponse); ok {
		code := coded.getCode()
		if code != "" && code != "200" {
			return &APIError{
				Code:       code,
				StatusCode: resp.StatusCode,
				Body:       string(body),
			}
		}
	}

	return nil
}
