package weather

import "qweather/common"

type Query struct {
	common.LocationQuery
	Unit common.Unit
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

type NowResponse struct {
	common.BaseResponse
	UpdateTime string       `json:"updateTime"`
	FxLink     string       `json:"fxLink"`
	Now        Now          `json:"now"`
	Refer      common.Refer `json:"refer"`
}

type Now struct {
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

type DailyResponse struct {
	common.BaseResponse
	UpdateTime string       `json:"updateTime"`
	FxLink     string       `json:"fxLink"`
	Daily      []Daily      `json:"daily"`
	Refer      common.Refer `json:"refer"`
}

type Daily struct {
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

type HourlyResponse struct {
	common.BaseResponse
	UpdateTime string       `json:"updateTime"`
	FxLink     string       `json:"fxLink"`
	Hourly     []Hourly     `json:"hourly"`
	Refer      common.Refer `json:"refer"`
}

type Hourly struct {
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
	common.BaseResponse
	UpdateTime string       `json:"updateTime"`
	FxLink     string       `json:"fxLink"`
	Summary    string       `json:"summary"`
	Minutely   []Minutely   `json:"minutely"`
	Refer      common.Refer `json:"refer"`
}

type Minutely struct {
	FxTime string `json:"fxTime"`
	Precip string `json:"precip"`
	Type   string `json:"type"`
}
