package token

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
)

type TokenInfo struct {
	Username string
	Email    string
	ID       int64
}

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

func GenerateJWT(tokenInfo *TokenInfo) (string, error) {
	tokenDuration := 24 * time.Hour
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: tokenInfo,
		"iat":              now.Unix(),
		"exp":              now.Add(tokenDuration).Unix(),
	})
	return t.SignedString([]byte(config.key))
}

func GetToken(c *gin.Context) (*TokenInfo, error) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return nil, errno.ErrTokenInvalid
	}
	var t string
	fmt.Sscanf(auth, "Bearer %s", &t)
	return Parse(t, config.key)
}

func Parse(token, key string) (*TokenInfo, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	var tokenInfo TokenInfo
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		t := claims[config.identityKey]
		tokenInfo.ID = int64(t.(map[string]interface{})["ID"].(float64))
		tokenInfo.Username = t.(map[string]interface{})["Username"].(string)
		tokenInfo.Email = t.(map[string]interface{})["Email"].(string)
	}
	return &tokenInfo, nil
}
