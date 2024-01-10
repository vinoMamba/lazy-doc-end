package token

import (
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	key         string
	identityKey string
}

var (
	config Config
	once   sync.Once
)

func Init(key, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

func GenerateJWT(identityKey string) (string, error) {
	tokenDuration := 24 * time.Hour
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"iat":              now.Unix(),
		"exp":              now.Add(tokenDuration).Unix(),
	})
	return t.SignedString([]byte(config.key))
}

func Parse(token, key string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	var identityKey string
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		identityKey = claims[config.identityKey].(string)
	}
	return identityKey, nil
}
