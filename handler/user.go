package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
	"github.com/vinoMamba/lazy-doc-end/model"
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
	db := storage.GetDB()
	log := logger.New(c)
	var body request.UserRegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if ok := utils.VerifyEmail(body.Email); !ok {
		log.Errorln("Email verify failed")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var u model.User

	db.Where("email = ?", body.Email).First(&u)
	if u.ID != 0 {
		log.WithField("email", u.Email).Errorln("The email has been registered")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if ok := utils.VerifyPassword(body.Password, body.ConfirmPassword); !ok {
		log.Errorln("password verify failed")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		log.WithError(err).Errorln("Hash password failed")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	u = model.User{
		Username: body.Email,
		Email:    body.Email,
		Password: hashedPassword,
	}
	result := db.Create(&u)
	if result.Error != nil {
		log.WithError(result.Error).Errorln("Create user failed")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}

func userLogin(c *gin.Context) {
	db := storage.GetDB()
	log := logger.New(c)
	var body request.UserLoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "Bad Request",
			"data":    nil,
		})
		return
	}
	var u model.User
	db.Where("username = ?", body.Username).First(&u)
	if u.ID == 0 {
		log.WithField("email", body.Username).Warnln("the email has not been registered")
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "The email has not been registered",
			"data":    nil,
		})
		return
	}

	if ok := utils.CheckHashPassword(u.Password, body.Password); !ok {
		log.WithField("email", body.Username).Warnln("the password is wrong")
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "The password is wrong",
			"data":    nil,
		})
		return
	}

	token, err := utils.CreateJwt(int64(u.ID), u.Email, u.Username)
	if err != nil {
		log.WithError(err).Errorln("Create jwt failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "server error",
			"data":    nil,
		})
		return
	}
	res := response.UserLoginResponse{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})
}
