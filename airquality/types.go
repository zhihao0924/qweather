package airquality

import "github.com/zhihao0924/qweather/common"

type CurrentResponse struct {
	Metadata   common.Metadata `json:"metadata"`
	Indexes    []Index         `json:"indexes"`
	Pollutants []Pollutant     `json:"pollutants"`
	Stations   []Station       `json:"stations"`
}

type Index struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	AQI              float64          `json:"aqi"`
	AQIDisplay       string           `json:"aqiDisplay"`
	Level            string           `json:"level"`
	Category         string           `json:"category"`
	Color            common.RGBAColor `json:"color"`
	PrimaryPollutant PollutantInfo    `json:"primaryPollutant"`
	Health           Health           `json:"health"`
}

type Health struct {
	Effect string       `json:"effect"`
	Advice HealthAdvice `json:"advice"`
}

type HealthAdvice struct {
	GeneralPopulation   string `json:"generalPopulation"`
	SensitivePopulation string `json:"sensitivePopulation"`
}

type PollutantInfo struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type Pollutant struct {
	Code          string        `json:"code"`
	Name          string        `json:"name"`
	FullName      string        `json:"fullName"`
	Concentration Concentration `json:"concentration"`
	SubIndexes    []SubIndex    `json:"subIndexes"`
}

type Concentration struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type SubIndex struct {
	Code       string  `json:"code"`
	AQI        float64 `json:"aqi"`
	AQIDisplay string  `json:"aqiDisplay"`
}

type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
