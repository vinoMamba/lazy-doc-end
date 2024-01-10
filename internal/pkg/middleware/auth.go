package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/core"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"github.com/vinoMamba/lazydoc/pkg/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := token.GetToken(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(known.XEmailKey, email)
		c.Next()
	}
}
