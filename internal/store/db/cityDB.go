package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type cityRepo struct {
	db *sqlx.Conn
}

func NewCityDB(config *config.Config) *cityRepo {
	return &cityRepo{db: config.DbConn}
}

func (c *cityRepo) Create(cities models.CityArray) {
	c.db.ExecContext(context.Background(), fmt.Sprintln("TRUNCATE city CASCADE"))

	for _, city := range cities {
		query := "INSERT INTO city (name, country, lat, lon) VALUES ($1, $2, $3, $4)"
		c.db.QueryRowxContext(context.Background(), query, city.Name, city.Country, city.Lat, city.Lon)
	}
}

func (c *cityRepo) GetAll() ([]*models.City, error) {
	var cities []*models.City

	rows, err := c.db.QueryxContext(context.Background(), "SELECT * FROM city")
	if err != nil {
		logrus.Errorln("[city repo] Failed download data from city DB, ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		rows.Scan(&city.Name, &city.Country, &city.Lat, &city.Lon)
		cities = append(cities, &city)
	}

	return cities, nil
}

func (c *cityRepo) GetByName(name string) (*models.City, error) {
	var city models.City

	err := c.db.QueryRowxContext(context.Background(), "SELECT name FROM city WHERE name = $1", name).Scan(&city.Name)
	if err != nil {
		logrus.Errorln("[city repo] Not found city name, ", err)
		return nil, err
	}

	if name != city.Name {
		return nil, errors.New("incorrect city name")
	}

	return &city, nil
}
