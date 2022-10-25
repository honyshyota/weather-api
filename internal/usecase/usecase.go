package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=usecase.go -destination=mock/mock.go

type (
	UserRepo interface {
		Create(user *models.User) error
		FindByName(name string) (*models.User, error)
		UpdateFavCity(string, string)
	}

	Usecase interface {
		GetAllCities() ([]string, error)
		GetShortForecast(string) (*models.ShortForecastResponse, error)
		GetForecastForTime(string, string) (*models.FullForecast, error)
		CreateUser(*models.User) error
		FindUser(string) (*models.User, error)
		UpdateFavCity(string, string) error
	}

	CityRepo interface {
		Create(cities models.CityArray)
		GetAll() ([]*models.City, error)
		GetByName(string) (*models.City, error)
	}

	WeatherRepo interface {
		Create(weathers []*models.CompleteWeather)
		Update(weathers []*models.CompleteWeather)
		GetByName(string) (*models.CompleteWeather, error)
		GetAll() ([]*models.CompleteWeather, error)
	}

	Client interface {
		GetCities([]string) (models.CityArray, error)
		GetForecast([]*models.City) ([]*models.CompleteWeather, error)
	}
)

type usecase struct {
	cityRepo    CityRepo
	weatherRepo WeatherRepo
	userRepo    UserRepo
	httpClient  Client
	config      *config.Config
}

func NewUsecase(cityRepo CityRepo, weatherRepo WeatherRepo, userRepo UserRepo, client Client, config *config.Config) *usecase {
	return &usecase{
		cityRepo:    cityRepo,
		weatherRepo: weatherRepo,
		userRepo:    userRepo,
		httpClient:  client,
		config:      config,
	}
}

func (u *usecase) InitSetupApp() {
	var wg sync.WaitGroup
	wg.Add(1)
	go u.UpdateWeatherRepo(context.Background(), u.config)
	wg.Done()
	wg.Wait()

	cities, err := u.cityRepo.GetAll()
	if err != nil {
		logrus.Errorln("[usecase] Failed give cities array from CityDB, ", err)
	} else if cities == nil {
		logrus.Println("[usecase] CityDB refresh from weather api")
		citiesRequest := u.InitCityList()
		u.cityRepo.Create(*citiesRequest)
		cities, err := u.cityRepo.GetAll()
		if err != nil {
			logrus.Errorln("[usecase] Failed give cities array from CityDB, ", err)
			return
		}

		forecasts, err := u.weatherRepo.GetAll()
		if err != nil {
			logrus.Errorln("[usecase] Failed give weathers array from WeatherDB, ", err)
			return
		}

		if forecasts == nil {
			u.createForecast(cities)
		}

		for _, forecast := range forecasts {
			if (time.Now().UTC().Minute() - forecast.Date.UTC().Minute()) > 30 {
				u.createForecast(cities)
			}
		}
	} else {
		forecasts, err := u.weatherRepo.GetAll()
		if err != nil {
			logrus.Errorln("[usecase] Failed give weathers array from WeatherDB, ", err)
			return
		}

		if forecasts == nil {
			u.createForecast(cities)
		}

		for _, forecast := range forecasts {
			if (time.Now().UTC().Minute() - forecast.Date.UTC().Minute()) > 1 {
				u.createForecast(cities)
				return
			}
		}
	}
}

func (u *usecase) createForecast(cities []*models.City) {
	logrus.Println("[usecase] WeatherDB refresh from weather api")
	forecasts, err := u.httpClient.GetForecast(cities)
	if err != nil {
		logrus.Errorln("[usecase] Failed give forecast from httpClient, ", err)
		return
	}
	u.weatherRepo.Create(forecasts)
}

func (u *usecase) InitCityList() *models.CityArray {
	citiesNameArray := u.config.CityList

	cities, err := u.httpClient.GetCities(citiesNameArray)
	if err != nil {
		logrus.Errorln("[usecase] Error give cities from httpClient, ", err)
	}

	return &cities
}

func (u *usecase) UpdateWeatherRepo(ctx context.Context, conf *config.Config) {

	ticker := time.NewTicker(time.Duration(conf.UpdateTime) * time.Minute)

	go func(tick *time.Ticker, conf *config.Config) {
		for {
			select {
			case <-ticker.C:
				cities, err := u.cityRepo.GetAll()
				if err != nil {
					logrus.Errorln("[usecase] Failed give cities array from CityDB, ", err)
				}
				logrus.Println("[usecase] Update WeatherDB")
				forecasts, err := u.httpClient.GetForecast(cities)
				if err != nil {
					logrus.Errorln("[usecase] Failed give forecast from httpClient, ", err)
					return
				}
				u.weatherRepo.Update(forecasts)
				tick.Reset(time.Duration(conf.UpdateTime) * time.Minute)
			case <-ctx.Done():
			}
		}
	}(ticker, conf)

}

func (u *usecase) GetAllCities() ([]string, error) {
	var result []string

	cities, err := u.cityRepo.GetAll()
	if err != nil {
		logrus.Errorln("[usecase] Failed to give all cities, ", err)

		return nil, err
	}

	for _, city := range cities {
		result = append(result, city.Name)
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result, nil
}

func (u *usecase) GetShortForecast(name string) (*models.ShortForecastResponse, error) {
	forecast, err := u.weatherRepo.GetByName(name)
	if err != nil {
		logrus.Errorln("[usecase] Failed to give forecast from weatherDB, ", err)
		return nil, err
	}

	var cityWeather *models.FullForecast

	err = json.Unmarshal(forecast.Data, &cityWeather)
	if err != nil {
		logrus.Errorln("[usecase] Cannot unmarshaling forecast, ", err)
		return nil, err
	}

	var dateSlice []string

	for _, val := range cityWeather.List {
		dateSlice = append(dateSlice, val.DtTxt)
	}

	return &models.ShortForecastResponse{
		Name:     forecast.Weather.City.Name,
		Country:  forecast.Weather.City.Country,
		AvgTemp:  forecast.Temp,
		DateList: dateSlice,
	}, nil
}

func (u *usecase) GetForecastForTime(name, date string) (*models.FullForecast, error) {
	data, err := u.weatherRepo.GetByName(name)
	if err != nil {
		logrus.Errorln("[usecase] Failed to give forecast from weatherDB, ", err)
		return nil, err
	}

	var forecasts models.FullForecast

	err = json.Unmarshal(data.Data, &forecasts)
	if err != nil {
		logrus.Println(err)
	}

	for _, forecast := range forecasts.List {
		if forecast.DtTxt == date {
			forecasts.List[0] = forecast
			forecasts.List = forecasts.List[:1]
		}
	}

	if len(forecasts.List) > 1 {
		return nil, errors.New("incorrect input")
	}

	return &forecasts, err
}
