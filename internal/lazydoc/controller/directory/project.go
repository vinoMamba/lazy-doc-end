package directory

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

type DirController struct {
	b biz.IBiz
}

func New(db store.IStore) *DirController {
	return &DirController{b: biz.NewBiz(db)}
}

func (ctrl *DirController) CreateDir(c *gin.Context) {
	var r request.CreateDirRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.BadRequest, nil)

		return
	}
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.BadRequest.SetMsg(err.Error()), nil)

		return
	}

	if err := ctrl.b.Directory().CreateDirBiz(c, &r); err != nil {
		core.WriteResponse(c, errno.InternalServerError, nil)

		return
	}

	core.WriteResponse(c, errno.OK, nil)
}
