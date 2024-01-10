package crypt

import (
	"fmt"
	"testing"
)

func TestPasswordEnCrypt(t *testing.T) {
	str := PasswordEncrypt("mihasdasdadsa")
	fmt.Println(str, len(str))

	fmt.Println(ComparePassword(str, "123456"))
}
