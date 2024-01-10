package lazydoc

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/controller/user"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
)

func registerAllApis(g *gin.Engine) error {

	register404Route(g)

	uc := user.New(store.Ds)
	g.POST("/user/register", uc.RegisterController)

	return nil
}

func register404Route(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.NotFound, nil)
	})
}
