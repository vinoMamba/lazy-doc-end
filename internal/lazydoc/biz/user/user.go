package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"github.com/vinoMamba/lazydoc/pkg/crypt"
	"github.com/vinoMamba/lazydoc/pkg/request"
	"github.com/vinoMamba/lazydoc/pkg/response"
	"github.com/vinoMamba/lazydoc/pkg/token"
)

type UserBiz interface {
	RegisterBiz(c context.Context, req *request.CreateUserRequest) error
	LoginBiz(c context.Context, req *request.LoginRequest) (*response.LoginResponse, error)
	GetUserInfoBiz(c context.Context, email string) (*response.UserInfoResponse, error)
	UpdatePasswordBiz(c context.Context, email string, req *request.UpdatePasswordRequest) error
	UpdateUserInfoBiz(c context.Context, email string, req *request.UpdateUserInfoRequest) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds}
}

func (biz *userBiz) RegisterBiz(c context.Context, req *request.CreateUserRequest) error {

	if req.ConfirmPassword != req.Password {
		return errno.ErrConfirmPassword
	}

	var userM model.UserM
	_ = copier.Copy(&userM, req)

	if err := biz.ds.Users().Create(c, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'users.email'", err.Error()); match {
			return errno.ErrEmailAlreadyInUse
		}
		return err
	}
	return nil
}

func (biz *userBiz) LoginBiz(c context.Context, req *request.LoginRequest) (*response.LoginResponse, error) {
	u, err := biz.ds.Users().GetUserByEmail(c, req.Email)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	if equal := crypt.ComparePassword(u.Password, req.Password); !equal {
		return nil, errno.ErrPassswordNotMatch
	}

	var tokenInfo token.TokenInfo
	_ = copier.Copy(&tokenInfo, u)

	t, err := token.GenerateJWT(&tokenInfo)

	if err != nil {
		return nil, errno.InternalServerError
	}

	var loginResponse response.LoginResponse
	_ = copier.Copy(&loginResponse, u)
	loginResponse.Token = t

	return &loginResponse, nil
}

func (biz *userBiz) GetUserInfoBiz(c context.Context, email string) (*response.UserInfoResponse, error) {
	u, err := biz.ds.Users().GetUserByEmail(c, email)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	var userInfo response.UserInfoResponse
	_ = copier.Copy(&userInfo, u)
	return &userInfo, nil
}

func (biz *userBiz) UpdatePasswordBiz(c context.Context, email string, req *request.UpdatePasswordRequest) error {
	u, err := biz.ds.Users().GetUserByEmail(c, email)
	if err != nil {
		return errno.ErrUserNotFound
	}

	if ok := crypt.ComparePassword(u.Password, req.OldPassword); !ok {
		return errno.ErrPassswordNotMatch
	}

	u.Password = crypt.PasswordEncrypt(req.NewPassword)
	if err := biz.ds.Users().UpdateUser(c, u); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (biz *userBiz) UpdateUserInfoBiz(c context.Context, email string, req *request.UpdateUserInfoRequest) error {
	u, err := biz.ds.Users().GetUserByEmail(c, email)
	if err != nil {
		return errno.ErrUserNotFound
	}

	copier.Copy(&u, req)

	if err := biz.ds.Users().UpdateUser(c, u); err != nil {
		return errno.InternalServerError
	}

	return nil
}
