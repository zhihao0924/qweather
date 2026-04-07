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

	resp, err := client.Geo.Lookup(context.Background(), qweather.CityLookupParams{
		Location: exampleutil.Env("QWEATHER_KEYWORD", "beijing"),
		Lang:     exampleutil.Env("QWEATHER_LANG", "en"),
		Number:   5,
	})
	if err != nil {
		log.Fatal(err)
	}

	exampleutil.PrintSection("Cities")
	for _, item := range resp.Location {
		fmt.Printf("%s (%s) %s,%s\n", item.Name, item.ID, item.Lat, item.Lon)
	}
}
