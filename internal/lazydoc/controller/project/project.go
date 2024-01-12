package project

import (
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
)

type ProjectController struct {
	b biz.IBiz
}

func New(db store.IStore) *ProjectController {
	return &ProjectController{b: biz.NewBiz(db)}
}
