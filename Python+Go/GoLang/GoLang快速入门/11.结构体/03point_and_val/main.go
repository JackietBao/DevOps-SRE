package main

import "fmt"

type Person struct {
	name string
	age  int
}

//值类型接受者
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v\n", p.name, p.age) // 姓名:小王子 年龄:25
}

//指针类型接收者
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}
func main() {
	p1 := Person{
		name: "小王子",
		age:  25,
	}
	p1.printInfo() // 姓名:小王子 年龄:25
	p1.setInfo("张三", 20)
	//调用了 p1 变量的 setInfo() 方法，传递了新的姓名和年龄作为参数。
	//因为 setInfo() 方法的接收者是指针类型，所以在方法内部对接收者进行的修改会影响到原始变量
	p1.printInfo() // 姓名:张三 年龄:20
}
