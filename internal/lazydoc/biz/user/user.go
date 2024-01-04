package user

import (
	"context"

	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
)

type UserBiz interface {
	Create(ctx context.Context, user *model.UserM) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds}
}

func (b *userBiz) Create(ctx context.Context, user *model.UserM) error {
	return b.ds.Users().Create(ctx, user)
}
