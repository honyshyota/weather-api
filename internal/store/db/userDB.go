package db

import (
	"github.com/honyshyota/weather-api/config"
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserDB(config *config.Config) *userRepo {
	return &userRepo{db: config.DbConn}
}

func (u *userRepo) Create(user *models.User) error {
	err := user.Validate()
	if err != nil {
		logrus.Println("[users db] Failed validation")
		return err
	}

	query := "INSERT INTO users (uuid, name, email, password) VALUES ($1, $2, $3, $4)"
	u.db.QueryRowx(query, user.Uuid, user.Name, user.Email, user.Password)

	return nil
}

func (u *userRepo) FindByName(name string) (*models.User, error) {
	var user models.User

	query := "SELECT name, email, password FROM users WHERE name = $1"
	err := u.db.QueryRowx(query, name).Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		logrus.Errorln("[users db] Failed download user, ", err)
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UpdateFavCity(city, name string) {
	query := "UPDATE users SET city = $1 WHERE name = $2"
	u.db.QueryRowx(query, city, name)
}
