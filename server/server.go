package server

import (
	"github.com/gin-gonic/gin"
)

func SetupHttpServer() *gin.Engine {
	r := gin.Default()
	return r
}
