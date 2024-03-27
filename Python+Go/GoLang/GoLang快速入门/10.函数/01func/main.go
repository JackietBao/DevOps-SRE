package main

import "fmt"

func main() {
	// 1、无参数、无返回值
	//fmt.Println(xxx)  // 0x4a7a80
	// go语言执行：执行入口在main方法
	//xxx()

	// 2、函数参数和返回值
	val := add(1, 1, "zhangsan") // val接收的是函数的返回值
	fmt.Println("函数执行结果：", val)

}

// 2、两数相加：函数参数和返回值
/*
func : 定义函数的关键字（go25个关键字）
add  : 函数的名字（类似定义的全局变量）
参数：x,y int
	- x y : 函数参数
	- int : 指定两个参数的类型（比如int、str）在调用函数时一定要使用相同的类型
返回值：(int)
	- 定义这个函数，执行结束后会返回一个int类型的数据
*/
func add(x, y int, c string) int {
	// x, y int 相同数据类型可以一起定义：“x,y int”  应该是 “x int ,y int” ?
	// , c string 不同数据类型需要分开定义
	fmt.Println("运行了add函数")
	return x + y
}

// 1、定义函数（无参数、无返回值）
/*
func : 定义函数的关键字（go25个关键字）
xxx  : 函数的名字（类似定义的全局变量）
*/
func xxx() {
	fmt.Println("hello world")
}
















package main

import "fmt"

func add(x, y int, c string) int {
//函数的签名是 func add(x, y int, c string) int
//这个函数叫做 add，接受两个整数类型的参数 x 和 y，以及一个字符串类型的参数 c
//并且返回一个整数类型的值
//它打印一条信息 "运行了add函数"，然后返回 x + y 的结果

	fmt.Println("运行了add函数")
	return x + y
}

func main() {
	val := add(1, 1, "zhangsan")
	//调用了 add 函数，并传入了参数 1, 1, "zhangsan"
	//意味着 x 的值为 1，y 的值为 1，c 的值为 "zhangsan"
	//main 函数会接收 add 函数的返回值，并将其赋给变量 val。
	//最后，它打印出 "函数执行结果：" + val
	fmt.Println("函数执行结果：", val)
}



