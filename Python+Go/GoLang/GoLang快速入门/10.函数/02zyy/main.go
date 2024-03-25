package main

import "fmt"

// go语言执行顺序
/*
首先执行 定义全局变量
然后执行 init方法
最后执行 main方法，调用我们写的逻辑
*/

// 1、全局变量
// 在函数外部直接定义的变量是全局变量（特点：可以在其他函数中使用）
var num int = 64

func main() {
	TestFunc()
	fmt.Println("main", num) // 这里打印的是全局的num
	//fmt.Println(num2)  // 局部变量在函数内部生效，不能在其他函数调用
}

func TestFunc() {
	//fmt.Println(num) // 全局变量可以在其他函数中使用
	//// 2、局部变量（在函数内部定义的变量，不能在其他函数中访问）
	//var num2 int = 1000
	//fmt.Println(num2)
	//
	//// 3、for循环中定义的变量
	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//}
	////fmt.Println(i)  // 在for后面定义的变量是自己这个代码块的，在外部不能使用

	//fmt.Println("test", num) // 全局变量可以在其他函数中使用
	//num = 3000

	// 局部作用域优先级更高，局部和全局定义了相同的变量名，局部在函数内部生效的，全局的被覆盖了
	num := 7 // := 类似var 声明一个新的变量(在内存中开启新空间)     = 是赋值
	fmt.Println("test", num)
	num = 444 // 函数内部如果修改，只修改的是自己，不影响全局的num
	fmt.Println("test", num)

}
