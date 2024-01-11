package lazydoc

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/controller/user"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/middleware"
)

func registerAllApis(g *gin.Engine) error {

	register404Route(g)

	uc := user.New(store.Ds)
	userGroup := g.Group("/user")
	userGroup.POST("/register", uc.RegisterController)
	userGroup.POST("/login", uc.LoginController)
	userGroup.Use(middleware.Auth())
	userGroup.GET("/info", uc.UserInfoController)
	userGroup.PUT("/password", uc.UpdatePassword)
	userGroup.PUT("/info", uc.UpdateUserInfo)

	return nil
}

func register404Route(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.NotFound, nil)
	})
}
