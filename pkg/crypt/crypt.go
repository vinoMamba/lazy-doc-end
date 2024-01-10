package crypt

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

var (
	salt = "lazy_doc"
)

func PasswordEncrypt(password string) string {
	bytes, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", string(bytes))
}

func ComparePassword(hashPassword, password string) bool {
	return subtle.ConstantTimeCompare([]byte(hashPassword), []byte(PasswordEncrypt(password))) == 1
}
