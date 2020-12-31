package main

import (
	"fmt"
	"reflect"
)

type stu struct {
	name string
	age  int
}

func main() {
	// var i string = "10"
	// fmt.Println(i)
	i := stu{
		name: "shineyork",
		// age:  10,
	}

	ref(i)
}

// 不明确传递类型
func ref(i interface{}) {
	// 如何知道i类型是什么
	// i.(type)
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)
	fmt.Println("i 的数据类型", rt) // 反射一个数据 的数据类型
	fmt.Println("i 的数据指", rv)
	fmt.Println(rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		fmt.Println("i 元素中字段", rt.Field(i).Name)
	}

	// fmt.Printf("type rt: %T\n", rt)
	// fmt.Printf("type rv: %T\n", rv)
	//
	// rvi := rv.Interface()
	// // rtv := reflect.ValueOf(rt)
	// rtv := rt.Name()
	// fmt.Printf("type rtv: %s\n", rtv)
	// // reflect.TypeOf(i) ==>> reflect.ValueOf(i) ==>> Interface()
	// fmt.Printf("type rvi: %T\n", rvi)
}
