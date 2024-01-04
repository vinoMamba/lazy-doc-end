package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		staus, code, msg := errno.DeCode(err)
		c.JSON(staus, ErrResponse{
			Code:    code,
			Message: msg,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
