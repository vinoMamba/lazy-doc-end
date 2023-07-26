package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vinoMamba/lazy-doc-end/routes"
)

func SetupRoute(r *gin.Engine) {
	registerGlobalMiddleWare(r)

	routes.RegisterApiRoutes(r)

	registerNotFoundRoute(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func registerNotFoundRoute(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
