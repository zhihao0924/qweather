package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"qweather"
)

func main() {
	host := "m24t2cbqc7.re.qweatherapi.com"
	apiKey := "2d6102e12c174e799c959e5b125c1684"
	location := os.Getenv("QWEATHER_LOCATION")
	if location == "" {
		location = "101010100"
	}

	client, err := qweather.NewClient(qweather.Config{
		Host:   host,
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Weather.Now(context.Background(), qweather.WeatherQuery{
		LocationQuery: qweather.LocationQuery{
			Location: location,
			Lang:     "zh",
		},
		Unit: qweather.UnitMetric,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %sC\n", resp.Now.Text, resp.Now.Temp)
}
