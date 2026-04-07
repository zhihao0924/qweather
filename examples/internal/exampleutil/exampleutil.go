package exampleutil

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"qweather"
)

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing required env %s", key)
	}
	return value
}

func Env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func MustFloatEnv(key string) float64 {
	value := MustEnv(key)
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("invalid float env %s=%q: %v", key, value, err)
	}
	return parsed
}

func APIKeyClient() *qweather.Client {
	client, err := qweather.NewClient(qweather.Config{
		Host:   MustEnv("QWEATHER_HOST"),
		APIKey: MustEnv("QWEATHER_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func JWTClient() *qweather.Client {
	privateKeyPEM, err := os.ReadFile(MustEnv("QWEATHER_PRIVATE_KEY_FILE"))
	if err != nil {
		log.Fatal(err)
	}

	tokenProvider, err := qweather.NewJWTTokenProvider(qweather.JWTConfig{
		CredentialID:  MustEnv("QWEATHER_CREDENTIAL_ID"),
		ProjectID:     MustEnv("QWEATHER_PROJECT_ID"),
		PrivateKeyPEM: privateKeyPEM,
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := qweather.NewClient(qweather.Config{
		Host:             MustEnv("QWEATHER_HOST"),
		JWTTokenProvider: tokenProvider,
	})
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func WeatherQuery() qweather.WeatherQuery {
	return qweather.WeatherQuery{
		LocationQuery: qweather.LocationQuery{
			Location: Env("QWEATHER_LOCATION", "101010100"),
			Lang:     Env("QWEATHER_LANG", "zh"),
		},
		Unit: qweather.UnitMetric,
	}
}

func CoordinateQuery() qweather.CoordinateQuery {
	return qweather.CoordinateQuery{
		Latitude:  MustFloatEnv("QWEATHER_LAT"),
		Longitude: MustFloatEnv("QWEATHER_LON"),
		Lang:      Env("QWEATHER_LANG", "zh"),
	}
}

func PrintSection(title string) {
	fmt.Printf("\n== %s ==\n", title)
}
