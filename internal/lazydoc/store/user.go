package store

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (s *users) Create(ctx context.Context, user *model.UserM) error {
	return s.db.WithContext(ctx).Create(user).Error
}
