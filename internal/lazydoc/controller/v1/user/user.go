package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	v1 "github.com/vinoMamba/lazydoc/pkg/api/v1"
)

type UserController struct {
	b biz.IBiz
}

func New(db store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(db)}
}

func (ctrl *UserController) RegisterController(c *gin.Context) {
	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.BadRequest, nil)

		return
	}

	if err := ctrl.b.User().RegisterBiz(c, &r); err != nil {
		core.WriteResponse(c, errno.InternalServerError, nil)

		return
	}

	core.WriteResponse(c, errno.OK, nil)
}
