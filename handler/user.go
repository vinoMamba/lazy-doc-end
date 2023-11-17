package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/models"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/params/response"
	"github.com/vinoMamba/lazy-doc-end/storage"
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
	_, err := storage.GetUserByEmail(c, body.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "email already exists",
		})
		return
	}

	if ok := utils.VerifyPassword(body.Password, body.ConfirmPassword); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "password verify failed",
		})
		return
	}

	hasPwd, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "password hash failed",
		})
		return
	}
	newUser := models.User{
		Username: body.Username,
		Password: hasPwd,
		Email:    body.Username,
	}

	id, err := storage.CreateUser(c, &newUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "create user failed",
		})
		return
	}

	c.JSON(http.StatusOK, response.UserRegisterResponse{
		Avatar:   "",
		Username: body.Username,
		Email:    body.Username,
		UserId:   id,
		Token:    "123456",
	})
}

func userLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
