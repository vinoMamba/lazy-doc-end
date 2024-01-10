package model

import (
	"time"

	"github.com/vinoMamba/lazydoc/pkg/crypt"
	"github.com/vinoMamba/lazydoc/pkg/id"
	"gorm.io/gorm"
)

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email;not null unique"`
	Password  string    `gorm:"column:password;not null"`
	IsDeleted int       `gorm:"column:is_deleted;default:0"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserM) TableName() string {
	return "users"
}

func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = id.GenSnowflakeID()
	u.Username = u.Email
	u.Password = crypt.PasswordEncrypt(u.Password)
	return nil
}
