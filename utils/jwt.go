package utils

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba/lazy-doc-end/config"
)

func CreateJwt(userId int64, email, username string) (string, error) {
	iat := time.Now()
	exp := iat.Add(time.Hour * 24 * 7)
	jwtKey := []byte(config.GetJwtSecret())

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"id":       strconv.FormatInt(userId, 10),
			"email":    email,
			"username": username,
		},
		"iat": iat.Unix(),
		"exp": exp.Unix(),
	})
	return t.SignedString(jwtKey)
}

func VerifyJwt(tokenString string) (*jwt.MapClaims, bool, error) {
	var cliams jwt.MapClaims
	t, err := jwt.ParseWithClaims(tokenString, &cliams, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtSecret()), nil
	})
	if err != nil {
		return nil, false, err
	}
	if t.Valid {
		return &cliams, true, nil
	} else {
		return nil, false, nil
	}
}

func GetCurrentUsername(c *gin.Context) string {
	mapClaims := c.MustGet("cliams").(*jwt.MapClaims)
	return (*mapClaims)["user"].(map[string]interface{})["username"].(string)
}

func GetCurrentEmail(c *gin.Context) string {
	mapClaims := c.MustGet("cliams").(*jwt.MapClaims)
	return (*mapClaims)["user"].(map[string]interface{})["email"].(string)
}

func GetCurrentUserId(c *gin.Context) int64 {
	mapClaims := c.MustGet("cliams").(*jwt.MapClaims)
	idStr := (*mapClaims)["user"].(map[string]interface{})["id"].(string)
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}
