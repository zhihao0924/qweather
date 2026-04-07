package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhihao0924/qweather"
	"github.com/zhihao0924/qweather/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.APIKeyClient()
	query := exampleutil.WeatherQuery()

	nowResp, err := client.Weather.Now(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	dailyResp, err := client.Weather.Daily(context.Background(), qweather.Daily3d, query)
	if err != nil {
		log.Fatal(err)
	}

	exampleutil.PrintSection("Now")
	fmt.Printf("%s %sC, humidity %s%%\n", nowResp.Now.Text, nowResp.Now.Temp, nowResp.Now.Humidity)

	exampleutil.PrintSection("Daily Forecast")
	for _, day := range dailyResp.Daily {
		fmt.Printf("%s %s/%sC %s\n", day.FxDate, day.TempMin, day.TempMax, day.TextDay)
	}
}
