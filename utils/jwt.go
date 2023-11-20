package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vinoMamba/lazy-doc-end/config"
)

func CreateJwt(email, username string) (string, error) {
	iat := time.Now()
	exp := iat.Add(time.Hour * 24 * 7)
	jwtKey := []byte(config.GetJwtSecret())

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
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
