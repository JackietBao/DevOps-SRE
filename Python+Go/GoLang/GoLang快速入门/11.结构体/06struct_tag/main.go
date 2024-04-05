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
	s, _ := json.Marshal(s1)
	//使用 json.Marshal() 函数将其编码为 JSON 格式的字节切片
	//由于 json.Marshal() 可能返回一个错误，但在这段代码中我们忽略了这个错误（使用了匿名变量 _）
	fmt.Println(string(s))
}

type Student struct {
	Name string `json:"xxx"`

	Age int `json:"age"`
	//字段的标签 json:"xxx" 指示编码时使用 xxx 作为 Name 字段的键，age 作为 Age 字段的键
}
