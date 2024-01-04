package user

import "fmt"

type UserBiz interface {
	Create()
}

type userBiz struct {
}

var _ UserBiz = (*userBiz)(nil)

func New() *userBiz {
	return &userBiz{}
}

func (b *userBiz) Create() {
	fmt.Println("create user")
}
