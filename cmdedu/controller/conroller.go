package controller

import (
	"cmdedu/service"
)

var (
	views       map[string][][3]string
	controllers map[string]interface{}
	autoService *service.AutoService
	userService *service.UserService
)

func init() {
	views = make(map[string][][3]string, 0)
	controllers = make(map[string]interface{}, 0)
	// 添加控制方法
	initView()
	addController()

	autoService = &service.AutoService{}
	userService = &service.UserService{}
}

func addController() {
	// addIndexController()
	// addUserController()
	// addAutoController()

	controllers["index"] = &IndexController{}
	controllers["user"] = &UserController{}
	controllers["auto"] = &AutoController{}
}

func initView() {
	views["login_view"] = [][3]string{
		0: {
			0: "auto",
			1: "Login",
			2: "登入系统",
		},
		1: {
			0: "auto",
			1: "Register",
			2: "注册用户",
		},
	}
	views["index_view"] = [][3]string{
		0: {
			0: "index",
			1: "Index",
			2: "进入首页",
		},
		1: {
			0: "user",
			1: "List",
			2: "展示用户信息",
		},
	}
}

// views[name]
// 0: {
// 	0: "auto",
// 	1: "Login",
// 	2: "登入系统",
// },
// 获取控制器与注册的方法
func getController(name string) [][3]string {
	return views[name]
}

// 获取方法
func getMethod(name, path string) string {
	var method string

	for _, methods := range views[name] {
		if methods[1] == path {
			method = methods[1]
			break
		}
	}
	return method
}

// [][3]string{
// 0: {
// 	0: "auto",
// 	1: "Login",
// 	2: "登入系统",
// },
// 1: {
// 	0: "auto",
// 	1: "Register",
// 	2: "注册用户",
// },
// 格式化输出方法和描述
func toModelFormate(methods [][3]string) ([]string, []string) {
	var method []string = make([]string, len(methods))
	var decs []string = make([]string, len(methods))

	for k, v := range methods {
		method[k] = v[0] + "::" + v[1]
		decs[k] = v[2]
	}

	return method, decs
}
