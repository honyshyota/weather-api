package usecase

import (
	"github.com/honyshyota/weather-api/internal/models"
	"github.com/sirupsen/logrus"
)

func (u *usecase) CreateUser(user *models.User) error {
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) FindUser(name string) (*models.User, error) {
	user, err := u.userRepo.FindByName(name)
	if err != nil {
		logrus.Errorln("[usecase] Failed received user data from userDB, ", err)
		return nil, err
	}
	return user, nil
}

func (u *usecase) UpdateFavCity(city, name string) error {
	_, err := u.cityRepo.GetByName(city)

	if err != nil {
		return err
	}

	u.userRepo.UpdateFavCity(city, name)
	return nil
}
