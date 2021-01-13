package controller

import (
	"encoding/json"
	"fmt"
	"testing"
)

var ar map[string][][3]string

func TestA(t *testing.T) {
	// ar = make(map[string][][3]string, 0)
	// ar["index"] = [][3]string{
	// 	0: {
	// 		0: "1",
	// 		1: "2",
	// 		2: "3",
	// 	},
	// }
	// fmt.Println(ar)
}

type Person struct {
	Name  string
	Hobby string
}

type Person2 struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility,string"`
}

func TestJson(t *testing.T) {
	p := Person{"shineyork", "男"}
	// type 转 json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "     ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
	// map 转 json
	student := make(map[string]interface{})
	student["name"] = "5lmh.com"
	student["age"] = 18
	student["sex"] = "man"
	b, err = json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// json 转 type
	c := []byte(`{"age":"18","name":"shineyork","marry":false}`)
	var p2 Person2
	err = json.Unmarshal(c, &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
