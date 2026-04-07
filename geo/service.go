package geo

import (
	"context"
	"net/http"

	"qweather/internal/sdk"
)

type Service struct {
	client *sdk.Client
}

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Lookup(ctx context.Context, params CityLookupParams) (*CityLookupResponse, error) {
	var out CityLookupResponse
	query := sdk.NewParams().
		Set("location", params.Location).
		Set("adm", params.Adm).
		Set("range", params.Range).
		Set("lang", params.Lang).
		SetInt("number", params.Number)
	if err := s.client.DoJSON(ctx, http.MethodGet, "/geo/v2/city/lookup", query, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *Service) TopCities(ctx context.Context, params TopCitiesParams) (*TopCitiesResponse, error) {
	var out TopCitiesResponse
	query := sdk.NewParams().
		Set("range", params.Range).
		Set("lang", params.Lang).
		SetInt("number", params.Number)
	if err := s.client.DoJSON(ctx, http.MethodGet, "/geo/v2/city/top", query, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
