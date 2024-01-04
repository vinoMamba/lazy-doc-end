package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(known.XRequestIDKey)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(known.XRequestIDKey, requestId)
		c.Writer.Header().Set(known.XRequestIDKey, requestId)
		c.Next()
	}
}
