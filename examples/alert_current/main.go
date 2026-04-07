package main

import (
	"context"
	"fmt"
	"log"

	"qweather"
	"qweather/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.APIKeyClient()
	query := exampleutil.CoordinateQuery()

	resp, err := client.Alerts.Current(context.Background(), qweather.AlertQuery{
		CoordinateQuery: query,
		LocalTime:       true,
	})
	if err != nil {
		log.Fatal(err)
	}

	exampleutil.PrintSection("Alerts")
	if len(resp.Alerts) == 0 {
		fmt.Println("no active alerts")
		return
	}

	for _, item := range resp.Alerts {
		fmt.Printf("%s %s %s\n", item.Severity, item.EventType.Name, item.Headline)
	}
}
