package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	var jsonStr = `{"ID":1,"Gender":"男","Name":"李四","Sno":"s0001"}`
	//定义了一个 JSON 字符串 jsonStr
	var student Student //定义一个 Monster 实例
	//创建了一个 Student 类型的变量 student，用于存储反序列化后的数据。
	err := json.Unmarshal([]byte(jsonStr), &student)
	//json.Unmarshal:将 JSON 数据解析为 Go 数据结构。它的作用是将 JSON 格式的数据解析成 Go 语言中的数据类型
	//[]byte(jsonStr)：将字符串 jsonStr 转换为字节切片（即 byte slice）
	//在 Go 语言中，字符串和字节切片之间可以相互转换
	//&student：这是一个取址操作符，用于获取结构体变量 student 的地址
	if err != nil {
		//检查反序列化过程中是否发生了错误，如果有错误，则打印错误信息
		fmt.Printf("unmarshal err=%v\n", err)
	}

	fmt.Printf("反序列化后 student=%#v student.Name=%v \n", student, student.Name)
	//%#v 会输出变量的 Go 语法表示形式，包括字段名和值；%v 会只输出值
}
