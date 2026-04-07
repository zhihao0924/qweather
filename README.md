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

## 目录结构

```text
.
├── common/               # 跨接口共享模型
├── auth/                 # JWT 鉴权实现
├── geo/                  # Geo API
├── weather/              # Weather API
├── airquality/           # Air Quality API
├── alert/                # Weather Alert API
├── history/              # Historical Weather API
├── internal/sdk/         # 通用 HTTP 请求与错误处理
├── internal/testutil/    # 测试辅助代码
├── client.go             # SDK 入口 Client
├── aliases.go            # 根包类型别名
├── auth.go               # JWT 对外包装
├── compat.go             # 向后兼容的 Client 直调方法
├── examples/             # 示例与示例说明
└── *_test.go             # 单元测试
```

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

	resp, err := client.Weather.Now(context.Background(), qweather.WeatherQuery{
		LocationQuery: qweather.LocationQuery{
			Location: "101010100",
			Lang:     "zh",
		},
		Unit: qweather.UnitMetric,
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

	_, err = client.Geo.Lookup(context.Background(), qweather.CityLookupParams{
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
client.Geo.Lookup(ctx, CityLookupParams)
client.Geo.TopCities(ctx, TopCitiesParams)
client.Weather.Now(ctx, WeatherQuery)
client.Weather.Daily(ctx, DailySpan, WeatherQuery)
client.Weather.Hourly(ctx, HourlySpan, WeatherQuery)
client.Weather.MinutelyPrecipitation(ctx, LocationQuery)
client.Alerts.Current(ctx, AlertQuery)
client.AirQuality.Current(ctx, CoordinateQuery)
client.History.Weather(ctx, HistoricalWeatherParams)
```

## 示例

先准备环境变量：

```bash
export QWEATHER_HOST="your-api-host.qweatherapi.com"
export QWEATHER_API_KEY="your-api-key"
export QWEATHER_LOCATION="101010100"
export QWEATHER_LANG="zh"
export QWEATHER_LAT="39.90"
export QWEATHER_LON="116.40"
```

运行示例：

```bash
go run ./examples/basic
go run ./examples/geo_lookup
go run ./examples/weather_forecast
go run ./examples/airquality_current
go run ./examples/alert_current
```

JWT 示例额外需要：

```bash
export QWEATHER_CREDENTIAL_ID="your-credential-id"
export QWEATHER_PROJECT_ID="your-project-id"
export QWEATHER_PRIVATE_KEY_FILE="./ed25519-private.pem"
go run ./examples/jwt
```

更多说明见 [examples/README.md](/Users/zhihao/Worker/qweather/examples/README.md)。

## 设计说明

- `Host` 使用和风天气控制台里的独立 `API Host`，例如 `abc123.def.qweatherapi.com`
- 为了适配和风天气 2026 年后的域名策略，SDK 默认不会写死公共域名
- 按领域拆成 `Geo / Weather / AirQuality / Alerts / History` service，`Client` 上保留兼容方法
- 气象类 `v7` 接口返回里包含 `code`，SDK 会在 `code != 200` 时返回 `*APIError`
- `airquality/v1` 和 `weatheralert/v1` 使用新的 v1 接口格式，因此不走 `code` 字段判断

## 测试

```bash
go test ./...
```
