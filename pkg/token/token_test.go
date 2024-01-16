package token

import (
	"testing"
)

func TestGenJWT(t *testing.T) {
	Init("test", "userInfo")

	tokenInfo := &TokenInfo{
		Username: "test",
		Email:    "test@qq.com",
		ID:       1,
	}
	token, err := GenerateJWT(tokenInfo)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)

	tt, err := Parse(token, config.key)
	if err != nil {
		t.Error(err)
	}
	t.Log(tt.Email)
	t.Log(tt.Username)
	t.Log(tt.ID)
}
