package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

func (ctrl *UserController) RegisterController(c *gin.Context) {
	var r request.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.BadRequest, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.BadRequest.SetMsg(err.Error()), nil)

		return
	}

	if err := ctrl.b.User().RegisterBiz(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, errno.OK, nil)
}
