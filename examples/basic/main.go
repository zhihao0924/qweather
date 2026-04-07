package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhihao0924/qweather/examples/internal/exampleutil"
)

func main() {
	client := exampleutil.APIKeyClient()

	resp, err := client.Weather.Now(context.Background(), exampleutil.WeatherQuery())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s %sC\n", resp.Now.Text, resp.Now.Temp)
}
