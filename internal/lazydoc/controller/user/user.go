package user

import (
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
)

type UserController struct {
	b biz.IBiz
}

func New(db store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(db)}
}
