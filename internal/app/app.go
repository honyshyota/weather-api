package app

import (
	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/controller"
	"github.com/honyshyota/weather-api/internal/service"
	"github.com/honyshyota/weather-api/internal/store/db"
	"github.com/honyshyota/weather-api/internal/usecase"
)

func Run() {
	config := config.NewConfig()
	httpClient := service.NewHttpClient(config)
	cityRepo := db.NewCityDB(config)
	weatherRepo := db.NewWeatherDB(config)
	userRepo := db.NewUserDB(config)
	usecase := usecase.NewUsecase(cityRepo, weatherRepo, userRepo, httpClient, config)
	usecase.InitSetupApp()
	controller.Build(usecase, config)
}
