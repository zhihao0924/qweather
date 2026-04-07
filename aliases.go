package qweather

import (
	"qweather/airquality"
	"qweather/alert"
	"qweather/common"
	"qweather/geo"
	"qweather/history"
	"qweather/internal/sdk"
	"qweather/weather"
)

type APIError = sdk.APIError

type Refer = common.Refer
type LocationQuery = common.LocationQuery
type CoordinateQuery = common.CoordinateQuery
type Location = common.Location
type Metadata = common.Metadata
type RGBAColor = common.RGBAColor
type Unit = common.Unit

const (
	UnitMetric   = common.UnitMetric
	UnitImperial = common.UnitImperial
)

type CityLookupParams = geo.CityLookupParams
type TopCitiesParams = geo.TopCitiesParams
type CityLookupResponse = geo.CityLookupResponse
type TopCitiesResponse = geo.TopCitiesResponse

type WeatherQuery = weather.Query
type DailySpan = weather.DailySpan
type HourlySpan = weather.HourlySpan

const (
	Daily3d  = weather.Daily3d
	Daily7d  = weather.Daily7d
	Daily10d = weather.Daily10d
	Daily15d = weather.Daily15d
	Daily30d = weather.Daily30d

	Hourly24h  = weather.Hourly24h
	Hourly72h  = weather.Hourly72h
	Hourly168h = weather.Hourly168h
)

type WeatherNowResponse = weather.NowResponse
type WeatherNow = weather.Now
type WeatherDailyResponse = weather.DailyResponse
type WeatherDaily = weather.Daily
type WeatherHourlyResponse = weather.HourlyResponse
type WeatherHourly = weather.Hourly
type MinutelyResponse = weather.MinutelyResponse
type MinutelyForecast = weather.Minutely

type AirQualityCurrentResponse = airquality.CurrentResponse
type AirQualityIndex = airquality.Index
type AirQualityHealth = airquality.Health
type AirQualityHealthAdvice = airquality.HealthAdvice
type AirQualityPollutantInfo = airquality.PollutantInfo
type AirQualityPollutant = airquality.Pollutant
type AirQualityConcentration = airquality.Concentration
type AirQualitySubIndex = airquality.SubIndex
type AirQualityStation = airquality.Station

type AlertQuery = alert.Query
type WeatherAlertCurrentResponse = alert.CurrentResponse
type WeatherAlertRecord = alert.Record
type WeatherAlertMessageType = alert.MessageType
type WeatherAlertEventType = alert.EventType

type HistoricalWeatherParams = history.WeatherParams
type HistoricalWeatherResponse = history.WeatherResponse
type HistoricalWeatherDaily = history.WeatherDaily
