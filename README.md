# qweather

一个基于 Go 的和风天气 API SDK，面向服务端调用场景，支持：

- API Host 接入
- `API KEY` 认证
- `JWT(Ed25519)` 认证
- 城市搜索、热门城市
- 实时天气、逐日预报、逐小时预报、分钟级降水
- 极端天气预警 `weatheralert/v1`
- 空气质量 `airquality/v1`
- 历史天气 `v7/historical/weather`

## 安装

```bash
go get qweather
```

如果你准备发布到自己的仓库，把 `go.mod` 里的 module 路径改成你的实际仓库地址即可。

## 初始化

### 1. 使用 API KEY

```go
package main

import (
	"context"
	"fmt"
	"log"

	"qweather"
)

func main() {
	client, err := qweather.NewClient(qweather.Config{
		Host:   "your-api-host.qweatherapi.com",
		APIKey: "your-api-key",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.WeatherNow(context.Background(), qweather.WeatherQuery{
		Location: "101010100",
		Lang:     "zh",
		Unit:     qweather.UnitMetric,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Now.Text, resp.Now.Temp)
}
```

### 2. 使用 JWT

```go
package main

import (
	"context"
	"log"
	"os"

	"qweather"
)

func main() {
	privateKeyPEM, err := os.ReadFile("./ed25519-private.pem")
	if err != nil {
		log.Fatal(err)
	}

	tokenProvider, err := qweather.NewJWTTokenProvider(qweather.JWTConfig{
		CredentialID:  "your-credential-id",
		ProjectID:     "your-project-id",
		PrivateKeyPEM: privateKeyPEM,
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := qweather.NewClient(qweather.Config{
		Host:             "your-api-host.qweatherapi.com",
		JWTTokenProvider: tokenProvider,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.CityLookup(context.Background(), qweather.CityLookupParams{
		Location: "beijing",
		Lang:     "en",
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

## 已实现接口

```go
CityLookup(ctx, CityLookupParams)
TopCities(ctx, TopCitiesParams)
WeatherNow(ctx, WeatherQuery)
WeatherDaily(ctx, DailySpan, WeatherQuery)
WeatherHourly(ctx, HourlySpan, WeatherQuery)
MinutelyPrecipitation(ctx, location, lang)
WeatherAlertCurrent(ctx, latitude, longitude, localTime, lang)
AirQualityCurrent(ctx, latitude, longitude, lang)
HistoricalWeather(ctx, HistoricalWeatherParams)
```

## 设计说明

- `Host` 使用和风天气控制台里的独立 `API Host`，例如 `abc123.def.qweatherapi.com`
- 为了适配和风天气 2026 年后的域名策略，SDK 默认不会写死公共域名
- 气象类 `v7` 接口返回里包含 `code`，SDK 会在 `code != 200` 时返回 `*APIError`
- `airquality/v1` 和 `weatheralert/v1` 使用新的 v1 接口格式，因此不走 `code` 字段判断

## 测试

```bash
go test ./...
```
