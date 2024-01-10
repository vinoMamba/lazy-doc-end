package token

import (
	"fmt"
	"testing"
)

func TestGenJWT(t *testing.T) {
	Init("secret", "email")
	token, err := GenerateJWT("1@qq.com")
	if err != nil {
		fmt.Println(err)
	}
	str, err := Parse(token, "secret")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
