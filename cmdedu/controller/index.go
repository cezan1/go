package controller

import "fmt"

//IndexController 首页控制器结构体
type IndexController struct {
}

//Index 首页
func (I *IndexController) Index() {
	view = "index_view"
	fmt.Println("这里是首页，没有什么信息")
}

// Welcome 欢迎页
func (I *IndexController) Welcome() {
	fmt.Println("欢迎来到 肖振的edu管理系统 cmdedu")
	fmt.Println("你的执行操作：")
	view = "login_view"
}
