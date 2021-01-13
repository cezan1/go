package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"cmdedu/util"
	"strconv"
	"strings"
)

type Model interface {
	ToString() string
	Save() bool
}

// 用于读取文件的信息内容
func fdata(name, dataKey string, models map[string]Model) error {
	f, ferr := os.Open(path + name + suffix)
	if ferr != nil {
		fmt.Println("抛出的异常", ferr)
		return errors.New("文件找不到")
	}
	defer f.Close()
	// 放到缓冲区中
	buf := bufio.NewReader(f)

	// fmt.Println("数据库的路径地址 : path", path)
	field := make([]string, 0) // 记录模型的字段
	for {
		b, rerr := buf.ReadBytes('\n') // 读取到换行就分开
		// fmt.Println("读取内容", string(b))
		if rerr != nil {
			if rerr == io.EOF {
				break
			}
			fmt.Println("抛出的缓存异常", rerr)
		}
		// 字符串分割
		data := strings.Split(string(b), ",")
		// 根据信息进行属性设置
		if len(field) == 0 { // 如果没有信息先记录
			// 记录需要新增的数据字段
			field = data
			for k, v := range data {
				// 去除影响数据的因素
				field[k] = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
			}
		} else {
			// fmt.Println("尝试添加的数据", data)
			// 根据字段值匹配到结构体中
			toModel(name, dataKey, models, data, field)
		}
	}
	// fmt.Println("最终的models", models)

	return nil
}

func fwrite(name string, models map[string]Model) bool {
	// 获取到转化的数据内容
	content := getModelsToString(models)
	// 打开文件
	// 不管是 Unix 还是 Windows，都需要使用 0666
	outfile, outErr := os.OpenFile(path+name+suffix, os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Println("文件找不到")
		return false
	}
	defer outfile.Close()

	// 创建写入的缓冲区
	outbufwri := bufio.NewWriter(outfile)
	outbufwri.WriteString(content + "\n")
	outbufwri.Flush()
	return true
}

// 根据传递的数据，与模型进行转化
func toModel(name, dataKey string, models map[string]Model, data, field []string) error {
	var dataValue string
	// 再利用反射分别调用对应的方法
	if newmodels[name] == nil {
		return errors.New("不存在" + name)
	}
	// reflect.ValueOf(i interface{}) Value 反射包的Value为GO提供反射接口
	// 从 newmodels 获取对应的方法实例
	fnm := reflect.ValueOf(newmodels[name])
	// 默认不传递参数
	// 执行对象方法 并等到一个新的模型
	modelV := fnm.Call([]reflect.Value{})[0]
	// 利用反射获取模型对象
	// 因为模型是指针
	modelZ := reflect.Zero(modelV.Type().Elem())
	// 利用反射获取到所有的set方法并把参数信息放进去
	for k, v := range data {
		// 去除影响数据的因素
		v = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
		// 判断是否是指定的查询key如果是记录value
		if field[k] == dataKey {
			dataValue = v
		}
		// 根据字符串找到Set系列方法
		fset := modelV.MethodByName("Set" + util.Ucfirst(field[k]))
		// 执行set给model设置值
		fset.Call([]reflect.Value{
			reflect.ValueOf(toTypeValue(modelZ, field[k], v)),
		})
	}
	// 记录数据
	models[dataValue] = modelV.Interface().(Model)
	return nil
}

// 根据传递的类型进行转化
func toTypeValue(model reflect.Value, field, val string) interface{} {
	// FieldByName只支持非指针的类型调用
	// 根据传递的字符串查找字段类型
	typ := model.FieldByName(field).Type()
	// 根据类型匹配相应的类型并转化
	switch typ.Name() {
	case "int":
		b, _ := strconv.Atoi(val)
		return b
	}
	return string(val)
}

// 把模型数据源转化为字符串
func getModelsToString(models map[string]Model) string {
	// 记录字段内容
	var fields string
	// 循环处理数据
	var content string
	for _, model := range models {
		if fields == "" {
			// 利用反射获取字段内容
			// rmodel := reflect.TypeOf(model)
			rmodel := reflect.ValueOf(model)
			modelZ := reflect.Zero(rmodel.Type().Elem())
			for i := 0; i < modelZ.NumField(); i++ {
				fields = fields + modelZ.Type().Field(i).Name + ","
			}
			fields = strings.TrimSuffix(fields, ",")
		}
		// 记录数据内容
		content = content + model.ToString() + "\n"
	}
	// 最终内容
	return fields + "\n" + strings.TrimSuffix(content, "\n")
}
