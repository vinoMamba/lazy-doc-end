package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPassword(t *testing.T) {
	pwd := "123456"
	cpwd := "123456"
	if !VerifyPassword(pwd, cpwd) {
		t.Error("VerifyPassword failed")
	}
	t.Log("VerifyPassword passed")
}

func TestVerifyEmail(t *testing.T) {
	var email string
	var isEmail bool
	email = "1111"
	isEmail = VerifyEmail(email)
	assert.Equal(t, false, isEmail)
	email = "vino@test"
	isEmail = VerifyEmail(email)
	assert.Equal(t, false, isEmail)
}
