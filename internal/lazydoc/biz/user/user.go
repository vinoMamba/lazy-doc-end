package user

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/log"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	v1 "github.com/vinoMamba/lazydoc/pkg/api/v1"
)

type UserBiz interface {
	RegisterBiz(c context.Context, req *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds}
}

func (biz *userBiz) RegisterBiz(c context.Context, req *v1.CreateUserRequest) error {
	log.Infow("RegisterBiz", "req", req)
	var userM model.UserM
	_ = copier.Copy(&userM, req)
	log.Infow("RegisterBiz", "userM", userM)
	if err := biz.ds.Users().Create(c, &userM); err != nil {
		return err
	}
	return nil
}
