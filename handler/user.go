package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/params/response"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func HandleUser(r *gin.Engine) {
	ug := r.Group("/user")
	ug.POST("/register", userRegister)
	ug.POST("/login", userLogin)
}

func userRegister(c *gin.Context) {
	var body request.UserRegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if ok := utils.VerifyEmail(body.Username); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "email verify failed",
		})
		return
	}

	if ok := utils.VerifyPassword(body.Password, body.ConfirmPassword); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "password verify failed",
		})
		return
	}

	c.JSON(http.StatusOK, response.UserRegisterResponse{
		Avatar:   "",
		Username: body.Username,
		Email:    "",
		UserId:   "123456",
		Token:    "123456",
	})
}

func userLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
