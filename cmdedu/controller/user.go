package controller

import (
	"fmt"
)

//UserController 用户控制
type UserController struct {
}

//List 展示用户
func (u *UserController) List() {
	view = "index_view"
	users := userService.GetList()
	fmt.Println("| username  | password | age | sex |")
	for _, user := range users {
		fmt.Printf("| %s | %s | %d | %s |\n", user.GetUsername(), user.GetPassword(), user.GetAge(), user.GetSex())
	}
}

//Update 修改用户
func (u *UserController) Update() {

}
