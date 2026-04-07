package qweather

import (
	"context"
	"net/http"

	"github.com/zhihao0924/qweather/airquality"
	"github.com/zhihao0924/qweather/alert"
	"github.com/zhihao0924/qweather/geo"
	"github.com/zhihao0924/qweather/history"
	"github.com/zhihao0924/qweather/internal/sdk"
	"github.com/zhihao0924/qweather/weather"
)

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
	client     *sdk.Client
	Geo        *geo.Service
	Weather    *weather.Service
	AirQuality *airquality.Service
	Alerts     *alert.Service
	History    *history.Service
}

func NewClient(cfg Config) (*Client, error) {
	inner, err := sdk.New(sdk.Config{
		Host:          cfg.Host,
		HTTPClient:    cfg.HTTPClient,
		UserAgent:     cfg.UserAgent,
		APIKey:        cfg.APIKey,
		APIKeyInQuery: cfg.APIKeyInQuery,
		TokenProvider: cfg.JWTTokenProvider,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client:     inner,
		Geo:        geo.New(inner),
		Weather:    weather.New(inner),
		AirQuality: airquality.New(inner),
		Alerts:     alert.New(inner),
		History:    history.New(inner),
	}, nil
}
