package qweather

import "context"

func (c *Client) CityLookup(ctx context.Context, params CityLookupParams) (*CityLookupResponse, error) {
	return c.Geo.Lookup(ctx, params)
}

func (c *Client) TopCities(ctx context.Context, params TopCitiesParams) (*TopCitiesResponse, error) {
	return c.Geo.TopCities(ctx, params)
}

func (c *Client) WeatherNow(ctx context.Context, query WeatherQuery) (*WeatherNowResponse, error) {
	return c.Weather.Now(ctx, query)
}

func (c *Client) WeatherDaily(ctx context.Context, span DailySpan, query WeatherQuery) (*WeatherDailyResponse, error) {
	return c.Weather.Daily(ctx, span, query)
}

func (c *Client) WeatherHourly(ctx context.Context, span HourlySpan, query WeatherQuery) (*WeatherHourlyResponse, error) {
	return c.Weather.Hourly(ctx, span, query)
}

func (c *Client) MinutelyPrecipitation(ctx context.Context, location, lang string) (*MinutelyResponse, error) {
	return c.Weather.MinutelyPrecipitation(ctx, LocationQuery{
		Location: location,
		Lang:     lang,
	})
}

func (c *Client) WeatherAlertCurrent(ctx context.Context, latitude, longitude float64, localTime bool, lang string) (*WeatherAlertCurrentResponse, error) {
	return c.Alerts.Current(ctx, AlertQuery{
		CoordinateQuery: CoordinateQuery{
			Latitude:  latitude,
			Longitude: longitude,
			Lang:      lang,
		},
		LocalTime: localTime,
	})
}

func (c *Client) AirQualityCurrent(ctx context.Context, latitude, longitude float64, lang string) (*AirQualityCurrentResponse, error) {
	return c.AirQuality.Current(ctx, CoordinateQuery{
		Latitude:  latitude,
		Longitude: longitude,
		Lang:      lang,
	})
}

func (c *Client) HistoricalWeather(ctx context.Context, params HistoricalWeatherParams) (*HistoricalWeatherResponse, error) {
	return c.History.Weather(ctx, params)
}
