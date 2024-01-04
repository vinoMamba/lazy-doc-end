package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{
		b: biz.NewBiz(ds),
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	ctrl.b.Users().Create(c, nil)
	core.WriteResponse(c, nil, gin.H{
		"msg": "register",
	})
}
