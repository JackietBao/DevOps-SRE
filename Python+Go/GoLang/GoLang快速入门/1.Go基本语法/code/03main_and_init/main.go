package main

import "fmt"

// 第一步：执行全局声明
var log string

// main函数有且只有一个
//func main()  {
//
//}

// 第二步：执行init函数
// 在main函数执行前自动被调用，常用来做初始化，比如初始化变量
func init() {
	fmt.Println("init")
	log = "log对象"
}

func init() {
	fmt.Println("init 222")
}

func SayHi2() {
	fmt.Println("Hello World222")
	fmt.Println()

}

// 第三步：最终从才会进入我们函数入口，执行你们写的代码
// go语言执行的函数入口，所有函数运行都是从main函数函数开始的
// main在一个包有且只有一个main函数
// 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用
func main() {
	SayHi() // 你写的所有代码逻辑，入口都在main函数
	fmt.Println("main")
	fmt.Println(log)
}

// 函数既变量：SayHi函数声明类似用变量
func SayHi() {
	fmt.Println("Hello World")
	SayHi2()
}

//sayHi()
