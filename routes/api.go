package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Pong": "pong",
			})
		})
	}
}
