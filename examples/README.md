# Examples

运行示例前先准备环境变量：

```bash
export QWEATHER_HOST="your-api-host.qweatherapi.com"
export QWEATHER_API_KEY="your-api-key"
export QWEATHER_LOCATION="101010100"
export QWEATHER_LANG="zh"
export QWEATHER_LAT="39.90"
export QWEATHER_LON="116.40"
```

JWT 示例还需要：

```bash
export QWEATHER_CREDENTIAL_ID="your-credential-id"
export QWEATHER_PROJECT_ID="your-project-id"
export QWEATHER_PRIVATE_KEY_FILE="./ed25519-private.pem"
```

可运行示例：

```bash
go run ./examples/basic
go run ./examples/geo_lookup
go run ./examples/weather_forecast
go run ./examples/airquality_current
go run ./examples/alert_current
go run ./examples/jwt
```
