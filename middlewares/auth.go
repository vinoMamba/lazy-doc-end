package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func AuthMiddleware(c *gin.Context) {
	log := logger.New(c)
	authStr := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(authStr, "Bearer ")
	cliams, ok, err := utils.VerifyJwt(token)
	if err != nil || !ok {
		log.WithError(err).Errorln("Verify jwt failed")
		c.JSON(401, gin.H{
			"code":    1,
			"message": "Unauthorized",
			"data":    nil,
		})
		c.Abort()
		return
	}
	log.Infof("cliams: %v", utils.Marshal(cliams))
	c.Set("cliams", cliams)
	c.Next()
}
