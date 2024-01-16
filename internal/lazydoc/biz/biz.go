package biz

import (
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz/directory"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/biz/user"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
)

type IBiz interface {
	User() user.UserBiz
	Directory() directory.DirBiz
}

type Biz struct {
	ds store.IStore
}

var _ IBiz = (*Biz)(nil)

func NewBiz(db store.IStore) *Biz {
	return &Biz{db}
}

func (b *Biz) User() user.UserBiz {
	return user.New(b.ds)
}

func (b *Biz) Directory() directory.DirBiz {
	return directory.New(b.ds)
}
