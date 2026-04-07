package qweather

import (
	"context"
	"fmt"
	"net/http"
)

type Unit string

const (
	UnitMetric   Unit = "m"
	UnitImperial Unit = "i"
)

type WeatherQuery struct {
	Location string
	Lang     string
	Unit     Unit
}

type DailySpan string

const (
	Daily3d  DailySpan = "3d"
	Daily7d  DailySpan = "7d"
	Daily10d DailySpan = "10d"
	Daily15d DailySpan = "15d"
	Daily30d DailySpan = "30d"
)

type HourlySpan string

const (
	Hourly24h  HourlySpan = "24h"
	Hourly72h  HourlySpan = "72h"
	Hourly168h HourlySpan = "168h"
)

type WeatherNowResponse struct {
	baseResponse
	UpdateTime string     `json:"updateTime"`
	FxLink     string     `json:"fxLink"`
	Now        WeatherNow `json:"now"`
	Refer      Refer      `json:"refer"`
}

type WeatherNow struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type WeatherDailyResponse struct {
	baseResponse
	UpdateTime string         `json:"updateTime"`
	FxLink     string         `json:"fxLink"`
	Daily      []WeatherDaily `json:"daily"`
	Refer      Refer          `json:"refer"`
}

type WeatherDaily struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UVIndex        string `json:"uvIndex"`
}

type WeatherHourlyResponse struct {
	baseResponse
	UpdateTime string          `json:"updateTime"`
	FxLink     string          `json:"fxLink"`
	Hourly     []WeatherHourly `json:"hourly"`
	Refer      Refer           `json:"refer"`
}

type WeatherHourly struct {
	FxTime    string `json:"fxTime"`
	Temp      string `json:"temp"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Pop       string `json:"pop"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type MinutelyResponse struct {
	baseResponse
	UpdateTime string             `json:"updateTime"`
	FxLink     string             `json:"fxLink"`
	Summary    string             `json:"summary"`
	Minutely   []MinutelyForecast `json:"minutely"`
	Refer      Refer              `json:"refer"`
}

type MinutelyForecast struct {
	FxTime string `json:"fxTime"`
	Precip string `json:"precip"`
	Type   string `json:"type"`
}

func (c *Client) WeatherNow(ctx context.Context, query WeatherQuery) (*WeatherNowResponse, error) {
	var out WeatherNowResponse
	if err := c.doJSON(ctx, http.MethodGet, "/v7/weather/now", weatherQueryMap(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) WeatherDaily(ctx context.Context, span DailySpan, query WeatherQuery) (*WeatherDailyResponse, error) {
	if span == "" {
		return nil, fmt.Errorf("qweather: daily span is required")
	}

	var out WeatherDailyResponse
	if err := c.doJSON(ctx, http.MethodGet, "/v7/weather/"+string(span), weatherQueryMap(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) WeatherHourly(ctx context.Context, span HourlySpan, query WeatherQuery) (*WeatherHourlyResponse, error) {
	if span == "" {
		return nil, fmt.Errorf("qweather: hourly span is required")
	}

	var out WeatherHourlyResponse
	if err := c.doJSON(ctx, http.MethodGet, "/v7/weather/"+string(span), weatherQueryMap(query), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) MinutelyPrecipitation(ctx context.Context, location, lang string) (*MinutelyResponse, error) {
	var out MinutelyResponse
	if err := c.doJSON(ctx, http.MethodGet, "/v7/minutely/5m", map[string]string{
		"location": location,
		"lang":     lang,
	}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func weatherQueryMap(query WeatherQuery) map[string]string {
	return map[string]string{
		"location": query.Location,
		"lang":     query.Lang,
		"unit":     string(query.Unit),
	}
}
