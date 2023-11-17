package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
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
	log := logger.New(c)
	var body request.UserRegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "Bad Request",
			"data":    nil,
		})
		return
	}

	if ok := utils.VerifyEmail(body.Username); !ok {
		log.Errorln("email verify failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "email verify failed",
			"data":    nil,
		})
		return
	}

	_, err := storage.GetUserByEmail(body.Username)
	if err == nil {
		log.Errorln("email already exists")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "email already exists",
			"data":    nil,
		})
		return
	}

	if ok := utils.VerifyPassword(body.Password, body.ConfirmPassword); !ok {
		log.Errorln("password verify failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "password verify failed",
			"data":    nil,
		})
		return
	}

	hashPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		log.WithError(err).Errorln("Hash password failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
	}

	newUser := models.User{
		Username: body.Username,
		Password: hashPassword,
		Email:    body.Username,
	}

	uId, err := storage.CreateUser(c, &newUser)
	if err != nil {
		log.WithError(err).Errorln("Create user failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "Server Error",
			"data":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data": response.UserRegisterResponse{
			Avatar:   "",
			UserId:   uId,
			Username: body.Username,
			Email:    body.Username,
			Token:    "",
		},
	})
}

func userLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
