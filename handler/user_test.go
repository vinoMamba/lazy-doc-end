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
		"code":    1,
		"message": "Bad Request",
		"data":    nil,
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}

func TestRegisterWithBadEmail(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()
	errReqBody := request.UserRegisterRequest{
		Username:        "test",
		Password:        "test",
		ConfirmPassword: "test",
	}
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(errReqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
	errRes := gin.H{
		"code":    1,
		"message": "email verify failed",
		"data":    nil,
	}
	assert.Equal(t, utils.Marshal(errRes), w.Body.String())
}

func TestRegisterWithRightEmail(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()
	errReqBody := request.UserRegisterRequest{
		Username:        "test@qq.com",
		Password:        "test",
		ConfirmPassword: "test",
	}
	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(errReqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	errRes := gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	}
	assert.Equal(t, utils.Marshal(errRes), w.Body.String())
}

func TestRegisterWithInValidPassword(t *testing.T) {
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
		"code":    1,
		"message": "password verify failed",
		"data":    nil,
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}

func TestRegisterWithValidPassword(t *testing.T) {
	r := server.SetupHttpServer()
	HandleUser(r)
	w := httptest.NewRecorder()
	reqBody := request.UserRegisterRequest{
		Username:        "test@test.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}

	req := httptest.NewRequest("POST", "/user/register", strings.NewReader(utils.Marshal(reqBody)))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	res := gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	}
	assert.Equal(t, utils.Marshal(res), w.Body.String())
}
