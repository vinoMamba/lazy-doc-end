package biz

import "github.com/vinoMamba/lazydoc/internal/lazydoc/biz/user"

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
}

var _ IBiz = (*biz)(nil)

func NewBiz() *biz {
	return &biz{}
}

func (b *biz) Users() user.UserBiz {
	return user.New()
}
