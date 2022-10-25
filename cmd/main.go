package main

import (
	"github.com/honyshyota/weather-api/internal/app"
)

// @title Weather API
// @version 1.0
// @description API server weather forecast

// @host localhost:8181
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name weather_api

func main() {
	app.Run()
}
