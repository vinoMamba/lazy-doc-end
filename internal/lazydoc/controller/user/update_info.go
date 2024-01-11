package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

func (ctrl *UserController) UpdateUserInfo(c *gin.Context) {
	var r *request.UpdateUserInfoRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.BadRequest, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.BadRequest.SetMsg(err.Error()), nil)

		return
	}

	email, _ := c.Get(known.XEmailKey)
	ctrl.b.User().UpdateUserInfoBiz(c, email.(string), r)
}
