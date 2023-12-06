package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/models"
)

func CreateProject(c context.Context, p *models.Project) error {
	return DB.WithContext(c).Create(p).Error
}

func UpdateProject(c context.Context, p *models.Project) error {
	return DB.WithContext(c).Where("id = ?", p.Id).Updates(p).Error
}

func GetProjectById(c context.Context, id int) (*models.Project, error) {
	var p models.Project
	if err := DB.WithContext(c).Model(models.Project{}).Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func GetPorjectList(c context.Context) ([]models.Project, error) {
	var projects []models.Project
	err := DB.WithContext(c).Where("is_public = ?", "0").Find(&projects).Error
	return projects, err
}

func DeleteProjectById(c context.Context, id int) error {
	return DB.WithContext(c).Where("id = ?", id).Update("is_deleted", "0").Error
}
