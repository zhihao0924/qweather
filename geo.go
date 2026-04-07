package qweather

import (
	"context"
	"net/http"
	"strconv"
)

type CityLookupParams struct {
	Location string
	Adm      string
	Range    string
	Number   int
	Lang     string
}

type CityLookupResponse struct {
	baseResponse
	Location []Location `json:"location"`
	Refer    Refer      `json:"refer"`
}

type TopCitiesParams struct {
	Range  string
	Number int
	Lang   string
}

type TopCitiesResponse struct {
	baseResponse
	TopCityList []Location `json:"topCityList"`
	Refer       Refer      `json:"refer"`
}

func (c *Client) CityLookup(ctx context.Context, params CityLookupParams) (*CityLookupResponse, error) {
	query := map[string]string{
		"location": params.Location,
		"adm":      params.Adm,
		"range":    params.Range,
		"lang":     params.Lang,
	}
	if params.Number > 0 {
		query["number"] = strconv.Itoa(params.Number)
	}

	var out CityLookupResponse
	if err := c.doJSON(ctx, http.MethodGet, "/geo/v2/city/lookup", query, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) TopCities(ctx context.Context, params TopCitiesParams) (*TopCitiesResponse, error) {
	query := map[string]string{
		"range": params.Range,
		"lang":  params.Lang,
	}
	if params.Number > 0 {
		query["number"] = strconv.Itoa(params.Number)
	}

	var out TopCitiesResponse
	if err := c.doJSON(ctx, http.MethodGet, "/geo/v2/city/top", query, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
