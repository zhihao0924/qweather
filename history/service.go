package history

import (
	"context"
	"net/http"

	"github.com/zhihao0924/qweather/internal/sdk"
)

type Service struct {
	client *sdk.Client
}

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Weather(ctx context.Context, params WeatherParams) (*WeatherResponse, error) {
	var out WeatherResponse
	query := sdk.NewParams().
		Set("location", params.Location).
		Set("date", params.Date).
		Set("lang", params.Lang).
		Set("unit", string(params.Unit))
	if err := s.client.DoJSON(ctx, http.MethodGet, "/v7/historical/weather", query, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
