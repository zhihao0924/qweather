package qweather

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type AirQualityCurrentResponse struct {
	Metadata   Metadata              `json:"metadata"`
	Indexes    []AirQualityIndex     `json:"indexes"`
	Pollutants []AirQualityPollutant `json:"pollutants"`
	Stations   []AirQualityStation   `json:"stations"`
}

type AirQualityIndex struct {
	Code             string                  `json:"code"`
	Name             string                  `json:"name"`
	AQI              float64                 `json:"aqi"`
	AQIDisplay       string                  `json:"aqiDisplay"`
	Level            string                  `json:"level"`
	Category         string                  `json:"category"`
	Color            RGBAColor               `json:"color"`
	PrimaryPollutant AirQualityPollutantInfo `json:"primaryPollutant"`
	Health           AirQualityHealth        `json:"health"`
}

type AirQualityHealth struct {
	Effect string                 `json:"effect"`
	Advice AirQualityHealthAdvice `json:"advice"`
}

type AirQualityHealthAdvice struct {
	GeneralPopulation   string `json:"generalPopulation"`
	SensitivePopulation string `json:"sensitivePopulation"`
}

type AirQualityPollutantInfo struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type AirQualityPollutant struct {
	Code          string                  `json:"code"`
	Name          string                  `json:"name"`
	FullName      string                  `json:"fullName"`
	Concentration AirQualityConcentration `json:"concentration"`
	SubIndexes    []AirQualitySubIndex    `json:"subIndexes"`
}

type AirQualityConcentration struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type AirQualitySubIndex struct {
	Code       string  `json:"code"`
	AQI        float64 `json:"aqi"`
	AQIDisplay string  `json:"aqiDisplay"`
}

type AirQualityStation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) AirQualityCurrent(ctx context.Context, latitude, longitude float64, lang string) (*AirQualityCurrentResponse, error) {
	var out AirQualityCurrentResponse
	path := fmt.Sprintf("/airquality/v1/current/%s/%s", formatCoordinate(latitude), formatCoordinate(longitude))
	if err := c.doJSON(ctx, http.MethodGet, path, map[string]string{
		"lang": lang,
	}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func formatCoordinate(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
