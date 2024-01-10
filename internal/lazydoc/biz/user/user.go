package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"github.com/vinoMamba/lazydoc/pkg/request"
)

type UserBiz interface {
	RegisterBiz(c context.Context, req *request.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

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
