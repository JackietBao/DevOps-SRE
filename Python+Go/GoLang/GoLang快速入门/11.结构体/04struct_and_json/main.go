package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string //私有属性不能被 json 包访问
	Sno    string
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v\n", s1) // main.Student{ID:1, Gender:"男", name:"李四", Sno:"s0001"}
	var s, _ = json.Marshal(s1)
	jsonStr := string(s)
	//将 s（即序列化后的 JSON 字节切片）转换为字符串类型，并将结果赋值给变量 jsonStr
	fmt.Println(jsonStr) // {"ID":1,"Gender":"男","Sno":"s0001"}
}
