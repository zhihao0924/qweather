package history

import "github.com/zhihao0924/qweather/common"

type WeatherParams struct {
	Location string
	Date     string
	Lang     string
	Unit     common.Unit
}

type WeatherResponse struct {
	common.BaseResponse
	FxLink       string       `json:"fxLink"`
	WeatherDaily WeatherDaily `json:"weatherDaily"`
	Refer        common.Refer `json:"refer"`
}

type WeatherDaily struct {
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
