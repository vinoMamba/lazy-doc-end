package biz

import (
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz/user"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
)

type IBiz interface {
	User() user.UserBiz
}

type Biz struct {
	ds store.IStore
}

func NewBiz(db store.IStore) *Biz {
	return &Biz{db}
}

func (b *Biz) User() user.UserBiz {
	return user.New(b.ds)
}
