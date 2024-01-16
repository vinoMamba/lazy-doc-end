package model

import (
	"time"

	"github.com/vinoMamba/lazydoc/pkg/id"
	"gorm.io/gorm"
)

type DirectoryM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	ParentId  int64     `gorm:"column:parent_id;default:0"`
	DirName   string    `gorm:"column:dir_name"`
	IsPublic  int       `gorm:"column:is_public;default:0"`
	IsDeleted int       `gorm:"column:is_deleted;default:0"`
	CreatedBy int64     `gorm:"column:created_by"`
	UpdatedBy int64     `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (d *DirectoryM) TableName() string {
	return "directories"
}

func (d *DirectoryM) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = id.GenSnowflakeID()
	return nil
}
