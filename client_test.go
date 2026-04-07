package qweather

import (
	"context"
	"github.com/zhihao0924/qweather/internal/testutil"
	"net/http"
	"strings"
	"testing"
)

func TestClientUsesAPIKeyHeader(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		Host:   "https://api.example.com",
		APIKey: "test-key",
		HTTPClient: &http.Client{
			Transport: testutil.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
				if got := r.Header.Get("X-QW-Api-Key"); got != "test-key" {
					t.Fatalf("unexpected api key header: %q", got)
				}
				if got := r.URL.Path; got != "/v7/weather/now" {
					t.Fatalf("unexpected path: %q", got)
				}
				if got := r.URL.Query().Get("location"); got != "101010100" {
					t.Fatalf("unexpected location query: %q", got)
				}
				return testutil.JSONResponse(`{"code":"200","updateTime":"2026-04-07T10:00+08:00","fxLink":"https://example.com","now":{"obsTime":"2026-04-07T09:50+08:00","temp":"24"},"refer":{"sources":["QWeather"],"license":["QWeather Developers License"]}}`), nil
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	resp, err := client.WeatherNow(context.Background(), WeatherQuery{
		LocationQuery: LocationQuery{Location: "101010100"},
	})
	if err != nil {
		t.Fatalf("WeatherNow() error = %v", err)
	}
	if resp.Now.Temp != "24" {
		t.Fatalf("unexpected temp: %q", resp.Now.Temp)
	}
}

func TestClientUsesAPIKeyQuery(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		Host:          "https://api.example.com",
		APIKey:        "test-key",
		APIKeyInQuery: true,
		HTTPClient: &http.Client{
			Transport: testutil.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
				if got := r.URL.Query().Get("key"); got != "test-key" {
					t.Fatalf("unexpected api key query: %q", got)
				}
				return testutil.JSONResponse(`{"code":"200","location":[],"refer":{}}`), nil
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if _, err := client.CityLookup(context.Background(), CityLookupParams{Location: "beijing"}); err != nil {
		t.Fatalf("CityLookup() error = %v", err)
	}
}

func TestClientUsesJWTProvider(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		Host:             "https://api.example.com",
		JWTTokenProvider: staticTokenProvider("jwt-token"),
		HTTPClient: &http.Client{
			Transport: testutil.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
				auth := r.Header.Get("Authorization")
				if !strings.HasPrefix(auth, "Bearer ") {
					t.Fatalf("unexpected auth header: %q", auth)
				}
				return testutil.JSONResponse(`{"metadata":{"tag":"abc"},"alerts":[]}`), nil
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if _, err := client.WeatherAlertCurrent(context.Background(), 39.90, 116.40, false, "en"); err != nil {
		t.Fatalf("WeatherAlertCurrent() error = %v", err)
	}
}

func TestClientReturnsAPIErrorForNon200Code(t *testing.T) {
	t.Parallel()

	client, err := NewClient(Config{
		Host:   "https://api.example.com",
		APIKey: "test-key",
		HTTPClient: &http.Client{
			Transport: testutil.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
				return testutil.JSONResponse(`{"code":"401"}`), nil
			}),
		},
	})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = client.WeatherNow(context.Background(), WeatherQuery{
		LocationQuery: LocationQuery{Location: "101010100"},
	})
	if err == nil {
		t.Fatal("expected error")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Code != "401" {
		t.Fatalf("unexpected api error code: %q", apiErr.Code)
	}
}

type staticTokenProvider string

func (s staticTokenProvider) Token(context.Context) (string, error) {
	return string(s), nil
}
