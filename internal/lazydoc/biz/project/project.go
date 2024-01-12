package project

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/log"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

type ProjectBiz interface {
	CreateProjectBiz(c context.Context, req *request.CreateProjectRequest) error
}

type projectBiz struct {
	ds store.IStore
}

var _ ProjectBiz = (*projectBiz)(nil)

func New(ds store.IStore) *projectBiz {
	return &projectBiz{ds}
}

func (p *projectBiz) CreateProjectBiz(c context.Context, req *request.CreateProjectRequest) error {
	log.C(c).Infow("create project", "req", req)
	var projectM model.ProjectM

	_ = copier.Copy(&projectM, req)

	if err := p.ds.Projects().Create(c, &projectM); err != nil {
		return err
	}

	return nil
}
