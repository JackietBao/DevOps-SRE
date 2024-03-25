package main

import (
	"encoding/json"
	"fmt"
)

// 反序列化：将json字符串转换成struct对象
/*
结构体的作用有两个
第一：绑定方法和接收者，然后通过结构体调用方法（类方法）
第二：可以和json字符串进行相互转
*/

/*
序列化： 将struct对象转换成 json的string格式
反序列化：将一个json字符串转换成struct对象
*/

func main() {
	//// 2、反序列化：结构体对应的json字符串数据(string)
	//// 当字符串 本身有双引号，或者 多行的时候使用 ``
	s := `{"ID":1,"Name":"zhangsan","Age":24,"Address":"bj"}`

	// 3、将json字符串转换为struct结构体
	var stu Student
	// 将string类型的数据转换为一个 []byte类型数据
	byteS := []byte(s)
	err := json.Unmarshal(byteS, &stu)
	if err != nil {
		fmt.Println("json.Unmarshal err, ", err)
	}
	//// %#v 是将数据展开
	//fmt.Printf("%T %#v \n", stu, stu)
	//fmt.Println(stu.Name, stu.Address)

	//// 4、结构体嵌套
	//p := Person{
	//	Id: 1,
	//	Stu: []Student{
	//		{ID: 1, Name: "zs", Age: 24, Address: "bj"},
	//		{ID: 2, Name: "lisi", Age: 24, Address: "bj"},
	//		{ID: 3, Name: "wangwu", Age: 24, Address: "bj"},
	//	},
	//}
	//
	//s, _ := json.Marshal(p.Stu)
	//fmt.Println(string(s))
}

type Person struct {
	Id  int
	Stu []Student
}

// 1、定义结构体
type Student struct {
	ID      int
	Name    string
	Age     int
	Address string
}
