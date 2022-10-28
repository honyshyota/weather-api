package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"

	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/sirupsen/logrus"
)

type httpClient struct {
	client *http.Client
	config *config.Config
}

func NewHttpClient(config *config.Config) *httpClient {
	return &httpClient{
		client: config.HttpClient,
		config: config,
	}
}

func (h *httpClient) GetCities(names []string) (models.CityArray, error) {
	cities := models.CityArray{}

	for _, name := range names {
		var city models.CityArray

		resp, err := h.client.Get(h.config.GetCityURI +
			name + "&appid=" + h.config.WeatherToken)
		if err != nil {
			logrus.Errorln("[http client] Failed give data from openweather, ", err)
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorln("[http client] Cannot read data from response body, ", err)
			return nil, err
		}

		err = json.Unmarshal(body, &city)
		if err != nil {
			logrus.Errorln("[http client] Cannot unmarshaling body to variable, ", err)
			return nil, err
		}

		cities = append(cities, city...)
	}

	return cities, nil
}

func (h *httpClient) GetForecast(cities []*models.City) ([]*models.CompleteWeather, error) {
	var weathers []*models.CompleteWeather
	resCh := make(chan *models.CompleteWeather, len(cities))
	errCh := make(chan error)

	for _, city := range cities {
		go func(city *models.City) {
			var weather models.FullForecast

			lat := fmt.Sprintf("%f", city.Lat)
			lon := fmt.Sprintf("%f", city.Lon)

			resp, err := http.Get(h.config.GetForecastURI +
				lat + "&lon=" + lon + "&appid=" + h.config.WeatherToken + "&units=metric")
			if err != nil {
				errCh <- err
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errCh <- err
				return
			}

			err = json.Unmarshal(body, &weather)
			if err != nil {
				errCh <- err
				return
			}

			weather.City.Name = city.Name

			var summTemp float64
			var numTemp int

			for i, val := range weather.List {
				summTemp += float64(val.Main.Temp)
				if numTemp < i {
					numTemp = i
				}
			}

			averageTemp := math.Round((summTemp/(float64(numTemp)+1))*100) / 100

			resultWeather := &models.CompleteWeather{
				Weather: weather,
				Temp:    averageTemp,
				Date:    time.Now(),
				Data:    body,
			}

			resCh <- resultWeather
		}(city)
	}

loop:
	for {
		select {
		case value := <-resCh:
			weathers = append(weathers, value)
			if len(weathers) == len(cities) {
				break loop
			}
		case err := <-errCh:
			return nil, err
		}
	}

	return weathers, nil
}
