package main

import "fmt"

// 要求：只要能理解go语言中有一种调用方法的方式，可以把一个方法绑定到一个结构体中
// 然后通过结构体就可以调用这个方法

// 函数式编程和类编程只是两种编程思想，如果个别同学还没发理解这种思维
// - 先按照讲的四步把代码写出来，会写
// - 会用：看场景，看别人的代码
// 结构体用法：用来模拟其他语言中的类
// 函数调用方式：直接通过函数名字调用
// 结构体方法调用：第一、初始化一个结构体   第二、通过结构体这对象.方法名调用

func main() {
	//第二步：实例化一个结构体（实例化一个类）
	pxx := Person{
		Name: "zhangsan",
		Age:  28,
	}
	//fmt.Println(pxx)

	// 第四步：调用结构的方法（通过类调用方法）
	pxx.printInfo()

	//px2 := Person{
	//	Name: "lisi",
	//	Age:  26,
	//}
	//px2.printInfo()
}

// PersonInfo

// 第一步：定义结构体(定义一个类)
type Person struct {
	Name string
	Age  int
}

// 第三步：给结构体绑定方法和接收者
//func  printInfo()  { // 这样定义的是一个函数
//
//}
/*
(p Person) : 接收者
	- p （变量名字）类似于我们python中类函数中的 self,或者js中类中的this
	- Person  接收者的结构体
printInfo() : 一个方法绑定了结构体之后就叫做 结构体方法
*/
func (xx Person) printInfo() { // 这样定义的是一个函数
	fmt.Println("name: ", xx.Name, xx.Age)
}
