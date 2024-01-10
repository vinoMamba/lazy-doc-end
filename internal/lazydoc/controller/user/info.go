package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
)

func (ctrl *UserController) UserInfoController(c *gin.Context) {
	email, _ := c.Get(known.XEmailKey)

	u, err := ctrl.b.User().GetUserInfoBiz(c, email.(string))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, u)
}
