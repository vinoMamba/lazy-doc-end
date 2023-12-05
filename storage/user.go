package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/models"
)

func CreateUser(c context.Context, u *models.User) error {
	return DB.WithContext(c).Create(u).Error
}

func GetUserByEmail(c context.Context, email string) (*models.User, error) {
	var u models.User
	if err := DB.WithContext(c).Model(models.User{}).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func UpdateUser(c context.Context, email string, u *models.User) error {
	return DB.WithContext(c).Where("email = ?", email).Updates(u).Error
}
