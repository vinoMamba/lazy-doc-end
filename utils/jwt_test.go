package utils

import (
	"fmt"
	"testing"

	"github.com/vinoMamba/lazy-doc-end/config"
)

func TestMain(m *testing.M) {
	config.LoadConfig("../")
	m.Run()
}

func TestJwt(t *testing.T) {
	token, err := CreateJwt("vino", "vino@qq.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	c, ok, err := VerifyJwt(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c)
	if !ok {
		t.Error("verify failed")
	}

}
