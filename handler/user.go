package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
	"github.com/vinoMamba/lazy-doc-end/models"
	"github.com/vinoMamba/lazy-doc-end/params/request"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Bad Request",
		})
		return
	}

	if ok := utils.VerifyEmail(body.Email); !ok {
		log.Errorln("Email verify failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Email verify failed",
		})
		return
	}

	if _, err := storage.GetUserByEmail(c, body.Email); err == nil {
		log.WithField("email", body.Email).Errorln("The email has been registered")
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    409,
			"message": "The email has been registered",
		})
		return
	}

	if ok := utils.VerifyPassword(body.Password, body.ConfirmPassword); !ok {
		log.Errorln("The password is not equal to confirm password")
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    400,
			"message": "The password is not equal to confirm password",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		log.WithError(err).Errorln("Hash password failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Server error",
		})
		return
	}

	u := models.User{
		Username: body.Email,
		Email:    body.Email,
		Password: hashedPassword,
	}

	if err := storage.CreateUser(c, &u); err != nil {
		log.Errorln("Create user failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func userLogin(c *gin.Context) {
	log := logger.New(c)
	var body request.UserLoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "Bad Request",
		})
		return
	}

	if ok := utils.VerifyEmail(body.Email); !ok {
		log.Errorln("Email verify failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Email verify failed",
		})
		return
	}

	u, err := storage.GetUserByEmail(c, body.Email)

	if err != nil {
		log.WithError(err).Errorln("Get user by email failed")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "User not found",
		})
		return
	}
	hashPwd := u.Password
	if ok := utils.CheckHashPassword(hashPwd, body.Password); !ok {
		log.Errorln("Password verify failed")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Password verify failed",
		})
		return
	}
	token, err := utils.CreateJwt(u.Id, u.Username, u.Email)
	if err != nil {
		log.WithError(err).Errorln("Create jwt failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Server error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"username": u.Username,
		"email":    u.Email,
		"token":    token,
	})

}
