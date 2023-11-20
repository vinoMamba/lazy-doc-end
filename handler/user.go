package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
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
	db := storage.NewQuery()
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

	u, _ := db.GetUserByEmail(c, body.Username)
	if u.ID != 0 {
		log.WithField("email", u.Email).Warnln("the email has been registered")
		c.JSON(http.StatusConflict, gin.H{
			"code":    1,
			"message": "The email has been registered",
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
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		log.WithError(err).Errorln("Hash password failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "server error",
			"data":    nil,
		})
		return
	}

	cUser := storage.CreateUserParams{
		Username: body.Username,
		Email:    body.Username,
		Password: hashedPassword,
	}
	result, err := db.CreateUser(c, cUser)

	if err != nil {
		log.WithError(err).Errorln("Create user failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "server error",
			"data":    nil,
		})
		return
	}

	token, err := utils.CreateJwt(body.Username, body.Username)
	if err != nil {
		log.WithError(err).Errorln("Create jwt failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "server error",
			"data":    nil,
		})
		return
	}

	id, _ := result.LastInsertId()
	res := response.UserRegisterResponse{
		Username: body.Username,
		Email:    body.Username,
		UserId:   id,
		Token:    token,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})
}

func userLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
