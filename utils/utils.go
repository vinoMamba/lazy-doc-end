package utils

import (
	"encoding/json"
	"regexp"
)

func Marshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func VerifyPassword(password, hashedPassword string) bool {
	return password == hashedPassword
}

func VerifyEmail(email string) bool {
	reg := regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`)
	return reg.MatchString(email)
}
