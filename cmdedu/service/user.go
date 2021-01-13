package service

import "cmdedu/model"

type UserService struct {
}

var (
	user *model.User
)

func init() {
	user = model.NewUser()
}

func (this *UserService) GetList() []*model.User {
	return user.All()
}
