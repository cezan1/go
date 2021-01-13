package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ReadfileFun func(a interface{})

func ReadJosn(filepath string, a interface{}) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return err
	}
	err = json.Unmarshal(bytes, a)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return err
	}
	return nil
}
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, nil
}

func WriteFile(path, file, data string) (bool, error) {
	// 判断是否存在目录
	if exist, _ := PathExist(path); !exist {
		// 不存在
		os.Mkdir(path, os.ModePerm)
	}
	// 打开文件
	// 0666 是文件的写入权限
	outputFile, outputError := os.OpenFile(path+file, os.O_WRONLY|os.O_CREATE, 0666)

	if outputError != nil {
		return false, outputError
	}
	defer outputFile.Close()
	// 创建写的缓冲区
	outputWriter := bufio.NewWriter(outputFile)
	// 写入信息
	outputWriter.WriteString(data)
	// 刷新
	outputWriter.Flush()
	return true, nil
}
