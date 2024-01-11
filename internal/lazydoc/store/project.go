package store

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"gorm.io/gorm"
)

type ProjectStore interface {
	Create(c context.Context, project *model.ProjectM) error
	UpdateProject(c context.Context, project *model.ProjectM) error
	GetProject(c context.Context, id int) (*model.ProjectM, error)
}
type projects struct {
	db *gorm.DB
}

var _ ProjectStore = (*projects)(nil)

func newProjects(db *gorm.DB) *projects {
	return &projects{db}
}

func (p *projects) Create(c context.Context, prpject *model.ProjectM) error {
	return p.db.Create(&prpject).Error
}

func (p *projects) GetProject(c context.Context, id int) (*model.ProjectM, error) {
	var project model.ProjectM
	if err := p.db.Where("id = ", id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *projects) UpdateProject(c context.Context, prpject *model.ProjectM) error {
	return p.db.Save(prpject).Error
}
