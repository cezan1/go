package service

import "cmdedu/model"

type AutoService struct {
}

func (this *AutoService) Login(username, password string) bool {
	user := model.GetUser(username)
	if user == nil {
		return false
	}

	if user.GetPassword() != password {
		return false
	} else {
		return true
	}
}

// 注册用户的信息
func (this *AutoService) Register(username, password string, age int, sex string) bool {
	user := model.NewUser()
	user.SetUsername(username)
	user.SetPassword(password)
	user.SetSex(sex)
	user.SetAge(age)
	user.Save()
	return true
}
