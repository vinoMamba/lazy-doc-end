package store

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"gorm.io/gorm"
)

type DirectoryStore interface {
	CreateDir(c context.Context, dir *model.DirectoryM) error
	UpdateDir(c context.Context, dir *model.DirectoryM) error
	GetDirList(c context.Context) ([]*model.DirectoryM, error)
}
type directories struct {
	db *gorm.DB
}

var _ DirectoryStore = (*directories)(nil)

func newDirectories(db *gorm.DB) *directories {
	return &directories{db}
}

func (d *directories) CreateDir(c context.Context, dir *model.DirectoryM) error {
	return d.db.Create(&dir).Error
}

func (d *directories) UpdateDir(c context.Context, dir *model.DirectoryM) error {
	return d.db.Save(dir).Error
}

func (d *directories) GetDirList(c context.Context) ([]*model.DirectoryM, error) {
	var dirs []*model.DirectoryM
	err := d.db.Find(&dirs).Error
	return dirs, err
}
