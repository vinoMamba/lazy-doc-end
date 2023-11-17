package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/server"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func TestRegisterWithoutBody(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/user/register", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	res := gin.H{
		"message": "Bad Request",
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}

func TestRegisterVerifyEmail(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()
	reqBody := request.UserRegisterRequest{
		Username:        "test",
		Password:        "test",
		ConfirmPassword: "test",
	}

	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(reqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)

	res := gin.H{
		"message": "email verify failed",
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}

func TestRegisterVerifyPassword(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()
	reqBody := request.UserRegisterRequest{
		Username:        "test@test.com",
		Password:        "123",
		ConfirmPassword: "456",
	}

	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(reqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)

	res := gin.H{
		"message": "password verify failed",
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}

func TestDuplicateRegistration(t *testing.T) {
	reqBody := request.UserRegisterRequest{
		Username:        "test@vino.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}

	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(reqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(reqBody)))
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 500, w2.Code)
	res := gin.H{
		"message": "email already exists",
	}
	assert.Equal(t, utils.Marshal(res), w2.Body.String())
}
