package db

import (
	"context"
	"fmt"

	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type weatherRepo struct {
	db *sqlx.Conn
}

func NewWeatherDB(config *config.Config) *weatherRepo {
	return &weatherRepo{db: config.DbConn}
}

func (w *weatherRepo) Create(weathers []*models.CompleteWeather) {
	w.db.ExecContext(context.Background(), fmt.Sprintln("TRUNCATE weather CASCADE"))

	for _, weather := range weathers {
		query := "INSERT INTO weather (name, country, lat, lon, temp, date, data) VALUES ($1, $2, $3, $4, $5, $6, $7::jsonb)"
		w.db.QueryRowxContext(context.Background(), query, weather.Weather.City.Name, weather.Weather.City.Country, weather.Weather.City.Coord.Lat,
			weather.Weather.City.Coord.Lon, weather.Temp, weather.Date, weather.Data)
	}
}

func (w *weatherRepo) GetByName(name string) (*models.CompleteWeather, error) {
	var result models.CompleteWeather

	err := w.db.QueryRowxContext(context.Background(), "SELECT name, country, temp, data FROM weather WHERE name = $1",
		name).Scan(&result.Weather.City.Name, &result.Weather.City.Country, &result.Temp, &result.Data)
	if err != nil {
		logrus.Errorln("[weather db] Failed download forecast, ", err)
		return nil, err
	}

	return &result, nil
}

func (w *weatherRepo) GetAll() ([]*models.CompleteWeather, error) {
	var weathers []*models.CompleteWeather

	rows, err := w.db.QueryxContext(context.Background(), "SELECT * FROM weather")
	if err != nil {
		logrus.Errorln("[weather repo] Failed download data from weather DB, ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var weather models.CompleteWeather
		rows.Scan(&weather.Weather.City.Name, &weather.Weather.City.Country, &weather.Weather.City.Coord.Lat,
			&weather.Weather.City.Coord.Lon, &weather.Temp, &weather.Date, &weather.Data)
		weathers = append(weathers, &weather)
	}

	return weathers, nil
}

func (w *weatherRepo) Update(weathers []*models.CompleteWeather) {
	for _, weather := range weathers {
		query := "UPDATE weather SET  temp = $1, date = $2, data = $3 WHERE name = $4"
		w.db.QueryRowxContext(context.Background(), query, weather.Temp, weather.Date, weather.Data, weather.Weather.City.Name)
	}
}
