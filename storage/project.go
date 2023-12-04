package storage

import (
	"context"

	"github.com/vinoMamba/lazy-doc-end/models"
)

func CreateProject(c context.Context, p *models.Project) error {
	return DB.WithContext(c).Create(p).Error
}
