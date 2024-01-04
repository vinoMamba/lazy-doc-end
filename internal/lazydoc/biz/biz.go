package biz

import (
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz/user"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
)

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) *biz {
	return &biz{ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
