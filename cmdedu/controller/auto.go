package controller

import (
	"cmdedu/util"
	"errors"
	"fmt"
	"strconv"
)

//AutoController struct
type AutoController struct {
}

// Login 用户登入
func (a *AutoController) Login() {

	fmt.Print("输入你的用户名 : ")
	username := util.CRead()
	fmt.Print("输入你的密码 : ")
	password := util.CRead()

	fmt.Println("username", username)
	fmt.Println("passworld", password)

	if autoService.Login(username, password) {
		fmt.Println("登入成功")
		view = "index_view"
	} else {
		fmt.Println("登入失败")
		view = "login_view"
	}

}

// Register 用户注册
func (a *AutoController) Register() error {
	view = "login_view"
	fmt.Println("输入你需要注册的用户信息 username,password,age,sex")
	fmt.Print("输入你的用户名 username: ")
	username := util.CRead()
	fmt.Print("输入你的密码 : ")
	password := util.CRead()
	fmt.Print("确认密码 : ")
	password1 := util.CRead()
	fmt.Print("输入你的年龄 : ")
	age, _ := strconv.Atoi(util.CRead())
	fmt.Print("输入你的性别 : ")
	sex := util.CRead()

	if password != password1 {
		return errors.New("密码需要一致")
	}

	if autoService.Register(username, password, age, sex) {
		fmt.Println("注册成功")
		return nil
	} else {
		fmt.Println("注册失败")
		return nil
	}
	// return nil
}
