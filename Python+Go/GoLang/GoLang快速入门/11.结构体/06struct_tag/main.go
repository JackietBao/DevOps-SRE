package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s1 := Student{
		Name: "zhangsan",
		Age:  24,
	}
	// 转json
	s, _ := json.Marshal(s1)
	fmt.Println(string(s))
}

// 定义结构体
type Student struct {
	// `json:"xxx"` 定义tag 将Name字段给他与name做反射
	// json是关键字
	Name string `json:"xxx"`
	//name string `json:"name"`
	Age int `json:"age"`
}
