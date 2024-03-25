package main

import (
	"encoding/json"
	"fmt"
)

// 序列化：将结构体转化成json字符串
func main() {
	// 2、初始化结构体
	s := Student{
		ID:      1,
		Name:    "zhangsan",
		Address: "bj",
		Age:     24,
	}
	fmt.Printf("%T %#v \n", s, s)
	// 3、将结构体转化为json数据
	/*
		json.Marshal方法返回两个值
		第一个值： sByte  是一个 []byte 对象（将结构体转化出来的数据）
		第二个值： err    json.Marshal转化失败，通过err接收
	*/
	sByte, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal err, ", err)
	}
	// 将string(sByte)类型转换为 string类型
	fmt.Println(string(sByte))
	/*
		结构体数据：main.Student main.Student{ID:1, Name:"zhangsan", Age:24, Address:"bj"}
		json数据：{"ID":1,"Name":"zhangsan","Age":24,"Address":"bj"}

	*/
}

// 1、定义结构体
type Student struct {
	ID      int
	Name    string
	Age     int
	Address string
}
