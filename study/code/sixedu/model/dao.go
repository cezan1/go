package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Model 统一的模型接口
type Model interface {
	ToString() string // 格式化输出数据信息
}

var (
	path   string                 = "D:\\MyGo\\src\\study\\code\\sixedu\\data\\" // 数据路径
	suffix string                 = ".sql"
	models map[string]interface{} // 记录标识 user =》 结构体
)

// "User" => User{} 不支持
func init() {
	// 标识绑定注册
	models = make(map[string]interface{})
	models["user"] = &User{}
}

// 数据库文件 -> 通过配置设置
// name 数据库名称   user,admin
// pirmay 查询主键
// models 存放数据
func rfdata(name, pirmay string, datas map[string]Model) error {
	// 1. 读取数据库文件 -》读取那个文件？
	f, ferr := os.Open(path + name + suffix) // 根据路径读取文件信息
	if ferr != nil {
		fmt.Println("文件读取异常,", ferr)
		return errors.New("文件查询不到 error")
	}
	// 关闭文件的资源流
	defer f.Close()
	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)
	// 2. 遍历数据  每一行的数据 字段根据 , 分割；数据通过 \n 分割
	field := make([]string, 0)
	data2 := make([]string, 0)
	for {
		row, rerr := buf.ReadBytes('\n') // 根据换行读取文件信息 , 返回的是byte[]
		if rerr != nil {
			if rerr == io.EOF { // 是否文件读取结束
				break
			}
			fmt.Println("抛出缓存读取文件异常", rerr)
		}
		// 去掉字符串，并分割数据
		data := strings.Split(strings.TrimSuffix(string(row), "\n"), ",")
		// fmt.Println("读取到的数据", data)

		// 根据数据判断操作 ； 是记录字段还是设置数据
		if len(field) == 0 {
			field = data // 存储字段
		} else {
			data2 = data
			//    2.2. 存储数据到 models
			//        2.2.1. 根据name得到模型
			//        2.2.2. 利用反射-》对模型赋值
			//        2.2.3. 再把模型存储在datas

		}
	}
	fmt.Println("\t存储的字段", field)
	fmt.Println("\t存储的值", data2)
	return nil

}
func rwdata() {

}
