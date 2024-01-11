package model

import (
	"time"

	"github.com/vinoMamba/lazydoc/pkg/id"
	"gorm.io/gorm"
)

type ProjectM struct {
	ID          int64     `gorm:"column:id;primary_key"`
	ProjectName string    `gorm:"column:project_name;not null"`
	ProjectDesc string    `gorm:"column:project_desc"`
	Status      int       `gorm:"column:status;default:0"`
	IsDeleted   int       `gorm:"column:is_deleted;default:0"`
	CreatedBy   string    `gorm:"column:created_by"`
	UpdatedBy   string    `gorm:"column:updated_by"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (p *ProjectM) TableName() string {
	return "projects"
}

func (p *ProjectM) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = id.GenSnowflakeID()
	return nil
}
