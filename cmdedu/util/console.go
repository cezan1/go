package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//CInfo test
type CInfo interface {
	call() (bool, error)
}

//Cfunc test
type Cfunc func() (bool, error)

func (c Cfunc) call() (bool, error) {
	return c()
}

var (
	inputReader *bufio.Reader
)

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

//CRead 读取输入
func CRead() string {
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(strings.TrimSuffix(input, "\n"))
	return input
}

//CReturn 格式化输出
func CReturn(cf Cfunc) bool {
	fmt.Println("<<<<<<<<=========== start ===========>>>>>>>>")
	// flag, err := a.call()
	flag, err := cf()
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	fmt.Println("<<<<<<<<=========== end ===========>>>>>>>>")
	fmt.Println("SB")
	return flag
}

//Coper 输出指令
func Coper(operate []string) {
	// operate := [3]string{0: "登入", 1: "注册", 2: "退出"}
	for key, value := range operate {
		fmt.Printf("(%d) : %s \n", key, value)
	}
	fmt.Println("退出请输 x ")
}
