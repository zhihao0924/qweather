package qweather

import (
	"context"
	"net/http"
)

type HistoricalWeatherParams struct {
	Location string
	Date     string
	Lang     string
	Unit     Unit
}

type HistoricalWeatherResponse struct {
	baseResponse
	FxLink       string                 `json:"fxLink"`
	WeatherDaily HistoricalWeatherDaily `json:"weatherDaily"`
	Refer        Refer                  `json:"refer"`
}

type HistoricalWeatherDaily struct {
	Date         string `json:"date"`
	Sunrise      string `json:"sunrise"`
	Sunset       string `json:"sunset"`
	Moonrise     string `json:"moonrise"`
	Moonset      string `json:"moonset"`
	MoonPhase    string `json:"moonPhase"`
	TempMax      string `json:"tempMax"`
	TempMin      string `json:"tempMin"`
	Humidity     string `json:"humidity"`
	Precip       string `json:"precip"`
	Pressure     string `json:"pressure"`
	Vis          string `json:"vis"`
	Wind360Day   string `json:"wind360Day"`
	WindDirDay   string `json:"windDirDay"`
	WindScaleDay string `json:"windScaleDay"`
	WindSpeedDay string `json:"windSpeedDay"`
	WeatherDay   string `json:"weatherDay"`
	IconDay      string `json:"iconDay"`
	HourlyPrecip string `json:"hourlyPrecip"`
}

func (c *Client) HistoricalWeather(ctx context.Context, params HistoricalWeatherParams) (*HistoricalWeatherResponse, error) {
	var out HistoricalWeatherResponse
	if err := c.doJSON(ctx, http.MethodGet, "/v7/historical/weather", map[string]string{
		"location": params.Location,
		"date":     params.Date,
		"lang":     params.Lang,
		"unit":     string(params.Unit),
	}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
