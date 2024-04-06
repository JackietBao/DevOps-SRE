package main

import "fmt"

func main() {
	m := map[string]string{
		"username": "zhangsan",
		//创建了一个包含一个键值对的map，键是"username"，对应的值是"zhangsan"
	}
	fmt.Printf("%T %v \n", m, m)
	//输出map变量m的类型和值
	fmt.Printf("%T %+v \n", m, m)
	//输出map变量m的类型和每个键值对的详细信息
	fmt.Printf("%T %#v \n", m, m)
	//%#v会以Go语法格式打印变量的值，对于map而言，它会打印每个键值对的键和值
	fmt.Println("Println, 自动换行")
	fmt.Println("Println, 自动换行")
	name := "zhangsan"
	age := 24
	fmt.Printf("\n 姓名：%s age: %d \n", name, age)

}
