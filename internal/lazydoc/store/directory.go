package store

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"gorm.io/gorm"
)

type DirectoryStore interface {
	CreateDir(c context.Context, dir *model.DirectoryM) error
	UpdateDir(c context.Context, dir *model.DirectoryM) error
	GetDirListByUserId(c context.Context, createdBy int64) ([]*model.DirectoryM, error)
	GetDirListByStatus(c context.Context, isPublic int) ([]*model.DirectoryM, error)
	GetDirById(c context.Context, id int64) (*model.DirectoryM, error)
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

func (d *directories) GetDirListByUserId(c context.Context, createdBy int64) ([]*model.DirectoryM, error) {
	var dirs []*model.DirectoryM
	err := d.db.WithContext(c).Where("created_by = ? and is_deleted = 0", createdBy).Find(&dirs).Error
	return dirs, err
}

func (d *directories) GetDirListByStatus(c context.Context, isPublic int) ([]*model.DirectoryM, error) {
	var dirs []*model.DirectoryM
	err := d.db.WithContext(c).Where("is_public = ? and is_deleted = 0", isPublic).Find(&dirs).Error
	return dirs, err
}

func (d *directories) GetDirById(c context.Context, id int64) (*model.DirectoryM, error) {
	var dir model.DirectoryM
	err := d.db.Where("id = ?", id).First(&dir).Error
	return &dir, err
}
