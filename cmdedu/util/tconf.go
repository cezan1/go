package util

import (
	"flag"
	"fmt"
)

var (
	instance *uconfig
	conf     = flag.String("conf", "../etc/config.json", "cmdedu . config")
)

type uconfig struct {
	BasePath string `json:"base_path"`
	DatePath string `json:"date_path"`
}

func init() {
	fmt.Println("tcong->init-start ")
	flag.Parse()
	NewUConfigWithFile(*conf)
}

//NewUConfigWithFile 实例配置文件
func NewUConfigWithFile(name string) *uconfig {
	if instance == nil {
		c := &uconfig{}
		ReadJosn(name, c)
		instance = c // <--- NOT THREAD SAFE
	}
	return instance
}

//GetConfig 获取配置值 *uconfig
func GetConfig() *uconfig {
	fmt.Println(instance)
	return instance
}
func (u *uconfig) GetDatePath() string {
	return u.DatePath
}
