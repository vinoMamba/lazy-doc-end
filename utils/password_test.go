package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pwd := "123456"
	hash, err := HashPassword(pwd)
	if err != nil {
		t.Error(err)
	}
	t.Log(hash)
	ok := CheckHashPassword(hash, pwd)
	assert.Equal(t, ok, true)
}
