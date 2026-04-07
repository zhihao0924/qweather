package weather

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zhihao0924/qweather/common"
	"github.com/zhihao0924/qweather/internal/sdk"
)

type Service struct {
	client *sdk.Client
}

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Now(ctx context.Context, query Query) (*NowResponse, error) {
	var out NowResponse
	if err := s.client.DoJSON(ctx, http.MethodGet, "/v7/weather/now", params(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *Service) Daily(ctx context.Context, span DailySpan, query Query) (*DailyResponse, error) {
	if span == "" {
		return nil, fmt.Errorf("qweather: daily span is required")
	}

	var out DailyResponse
	if err := s.client.DoJSON(ctx, http.MethodGet, "/v7/weather/"+string(span), params(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *Service) Hourly(ctx context.Context, span HourlySpan, query Query) (*HourlyResponse, error) {
	if span == "" {
		return nil, fmt.Errorf("qweather: hourly span is required")
	}

	var out HourlyResponse
	if err := s.client.DoJSON(ctx, http.MethodGet, "/v7/weather/"+string(span), params(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *Service) MinutelyPrecipitation(ctx context.Context, query common.LocationQuery) (*MinutelyResponse, error) {
	var out MinutelyResponse
	if err := s.client.DoJSON(ctx, http.MethodGet, "/v7/minutely/5m", locationParams(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func locationParams(query common.LocationQuery) sdk.Params {
	return sdk.NewParams().
		Set("location", query.Location).
		Set("lang", query.Lang)
}

func params(query Query) sdk.Params {
	return locationParams(query.LocationQuery).
		Set("unit", string(query.Unit))
}
