package store

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(c context.Context, u *model.UserM) error
	GetUserByEmail(email string) (*model.UserM, error)
}

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (s *users) Create(c context.Context, user *model.UserM) error {
	return s.db.Create(&user).Error
}

func (s *users) GetUserByEmail(email string) (*model.UserM, error) {
	var user model.UserM
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
