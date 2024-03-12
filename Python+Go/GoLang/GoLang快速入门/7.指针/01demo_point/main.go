package main

import "fmt"

func main() {
	//第一步：声明一个变量a
	//第二步：在内存中申请了一块内存空间，用来存放 100这个值
	//第三步：要想用a变量取出100这个值，实际先获取到的是100的内存地址
	a := 100
	//1、指针地址概念
	fmt.Print(&a)
	fmt.Printf("%d \n", &a)
	b := &a

	//2、*&a取出值
	fmt.Println(*b)

	//3、获取指针地址
	fmt.Printf("%T: %v", &a, &a)
	//%v 来表示相应值的默认格式
}
