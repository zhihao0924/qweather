package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhihao0924/qweather"
	"github.com/zhihao0924/qweather/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.JWTClient()

	resp, err := client.Geo.Lookup(context.Background(), qweather.CityLookupParams{
		Location: exampleutil.Env("QWEATHER_KEYWORD", "shanghai"),
		Lang:     exampleutil.Env("QWEATHER_LANG", "en"),
		Number:   3,
	})
	if err != nil {
		log.Fatal(err)
	}

	exampleutil.PrintSection("JWT City Lookup")
	for _, item := range resp.Location {
		fmt.Printf("%s (%s)\n", item.Name, item.ID)
	}
}
