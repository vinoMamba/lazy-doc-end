package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/model"
)

func CreateUser(c context.Context, u *model.User) error {
	return DB.WithContext(c).Create(u).Error
}

func GetUserByEmail(c context.Context, email string) (*model.User, error) {
	var u model.User
	if err := DB.WithContext(c).Model(model.User{}).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
