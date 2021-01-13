package model

import (
	"cmdedu/util"
	"fmt"
	"reflect"
)

var (
	path      string
	suffix    string = ".sql"
	newmodels map[string]interface{}
	config    Config
	// configFile = flag.String("config", "./etc/config.json", "cmdedu . config")
)

func init() {
	// fmt.Println("model -->> ")
	path = util.GetConfig().GetDatePath()
	// fmt.Println("path:", path)
	initData()
	initNewModel()
}

// 用于记录需要创建模型的方法
func initNewModel() {
	newmodels = make(map[string]interface{})
	// key 为表，value为模型
	newmodels["user"] = NewUser
	userDatas = make(map[string]Model, 0)
	fdata("user", userDataKey, userDatas)
	// fmt.Println("解析的数据：", userDatas)
}

// 初始化数据
func initData() {
	/*
		username,password,age,sex
		shineyork,123456,18,男
		sixstar,123456,18,男
		ppp,123456,99,男
	*/
	fmt.Println("初始化数据 ")
	flag, _ := util.PathExist(path + "user.sql")
	// fmt.Println("校验目录", path+"\\user.sql", "是否存在 ：", flag)
	if !flag {
		fmt.Println("初始化", path+"user.sql", "数据信息")
		// 创建文件并写入数据
		data := "username,password,age,sex\nshineyork,123456,18,男\nsixstar,123456,18,男\nppp,123456,99,男\n"
		b, err := util.WriteFile(path, "user.sql", data)
		fmt.Println("创建结果:", b, " error:", err)
	}
}

func ref() {
	m := NewUser()

	// 反射获取字段类型
	typeOfCat := reflect.TypeOf(m)
	fmt.Println("typeOfCat:", typeOfCat)
	temp := reflect.Zero(typeOfCat.Elem())
	fmt.Println("temp:", temp)
	typ := temp.FieldByName("sex").Type()

	fmt.Println(typeOfCat)
	fmt.Println(typ)
	switch typ.Name() {
	case "string":
		fmt.Println("ok")
	}

}
