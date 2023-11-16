package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/server"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func TestRegister(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()

	registerBody := request.UserRegisterRequest{
		Username:        "vino",
		Password:        "123456",
		ConfirmPassword: "123456",
	}

	body := strings.NewReader(utils.Marshal(registerBody))

	req := httptest.NewRequest("POST", "/user/register", body)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
