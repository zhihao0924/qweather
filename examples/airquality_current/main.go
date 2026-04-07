package main

import (
	"context"
	"fmt"
	"log"

	"qweather/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.APIKeyClient()
	query := exampleutil.CoordinateQuery()

	resp, err := client.AirQuality.Current(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	exampleutil.PrintSection("Air Quality")
	for _, index := range resp.Indexes {
		fmt.Printf("%s AQI %.0f %s\n", index.Name, index.AQI, index.Category)
	}
}
