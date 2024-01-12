package project

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

func (ctrl *ProjectController) CreateProject(c *gin.Context) {
	var r request.CreateProjectRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.BadRequest, nil)

		return
	}
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.BadRequest.SetMsg(err.Error()), nil)

		return
	}

	if err := ctrl.b.Project().CreateProjectBiz(c, &r); err != nil {
		core.WriteResponse(c, errno.InternalServerError, nil)

		return
	}

	core.WriteResponse(c, errno.OK, nil)
}
