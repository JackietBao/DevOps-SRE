/*

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` //通过指定 tag 实现 json 序列化该字段时的 key
	Gender string `json:"gender"`
	// 使用了标签 json:"id" 和 json:"gender"。这些标签用于指定在 JSON 序列化时，对应字段的名称
	Name string
	Sno  string
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	// main.Student{ID:1, Gender:"男", Name:"李四", Sno:"s0001"}
	fmt.Printf("%#v\n", s1)
	var s, _ = json.Marshal(s1)
	//将结构体 s1 编码为 JSON 格式
	jsonStr := string(s)
	fmt.Println(jsonStr) // {"id":1,"gender":"男","Name":"李四","Sno":"s0001"}
}


*/

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` //通过指定 tag 实现 json 序列化该字段时的 key
	Gender string `json:"gender"`
	Name   string
	Sno    string
	//json:"id" 和 json:"gender" 是结构体字段的标签，用于指定在 JSON 序列化和反序列化时的键
}

func main() {
	var s2 Student
	//定义了一个 Student 类型的变量 s2
	var str = `{"id":1,"gender":"男","Name":"李四","Sno":"s0001"}`
	//JSON 格式的字符串 str，表示一个学生的信息
	err := json.Unmarshal([]byte(str), &s2)
	//使用 json.Unmarshal 函数将 JSON 格式的字符串解码为 Student 结构体类型的变量 s2
	//这个函数接受两个参数，第一个是 JSON 字节切片
	//第二个是目标变量的指针。解码成功后，目标变量 s2 将包含解析得到的值
	if err != nil {
		fmt.Println(err)
	}
	// main.Student{ID:1, Gender:"男", Name:"李四", Sno:"s0001"}
	fmt.Printf("%#v", s2)
}
