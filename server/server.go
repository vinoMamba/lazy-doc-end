package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/handler"
)

func SetupHttpServer() {
	r := gin.Default()
	handler.HandleUser(r)
	r.Run(":3000")
}
