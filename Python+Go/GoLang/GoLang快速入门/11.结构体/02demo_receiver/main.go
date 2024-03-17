package main

import "fmt"

type Person struct {
	name string
	age  int8
}

func (p Person) printInfo() {
	//定义了一个名为 printInfo 的方法，其接收者(receiver)是 Person 类型的变量 p
	//printInfo() 方法在这里被称为 Person 类型的一个方法
	fmt.Printf("姓名:%v 年龄:%v", p.name, p.age) // 姓名:小王子 年龄:25

}
func main() {
	p1 := Person{
		name: "小王子",
		age:  25,
	}
	p1.printInfo() // 姓名:小王子 年龄:25
}
