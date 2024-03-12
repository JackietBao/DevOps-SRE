package main

import "fmt"

func main() {
	// 1、声明引用类型，必须要使用make或者new方法申请内存空间
	//我们只是声明了一个变量，但是没有到内存中申请空间
	//var userInfo map[string]int
	//map[string]int，表示一个字符串键对应一个整数值的映射。
	//在这一步中，userInfo 被声明但是还没有被初始化，它的零值是 nil，即未分配内存的空映射
	//userInfo["zhangsan"] = 24
	//尝试向 userInfo 映射中添加一个键值对，将字符串 "zhangsan" 映射到整数 24。
	//由于 userInfo 是一个 nil 映射，这个操作会导致运行时错误
	//fmt.Println(userInfo)
	//2、使用make方法申请空间(返回值类型)
	//第一个 []int : 指定数据类型是一个 int切片
	//第二 3 10 ： 切片默认长度是3， 默认容量是 10

	a := make([]int, 3, 10)
	//创建了一个切片，长度为 3，容量为 10，其中元素类型为整数
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
	//len(a) 返回切片的长度
	//cap(a) 返回切片的容量

	b := new([]int)
	//new() 函数用于动态分配内存，并返回一个指向新分配的零值的指针
	//new([]int) 分配了一个长度为 0 的切片，并将其地址分配给变量 b。
	//b 是一个指向切片的指针，因为 new() 返回的是指针
	*b = append(*b, 2)
	//*b 解引用指针 b，获取指针指向的实际值，即切片
	//向切片指针 b 指向的切片中追加了元素 2
	fmt.Printf("make:%T  ==> new: %T \n", a, b)
	fmt.Println(b)
	//表示切片指针 b 指向的切片的值为 [2]，同时 b 本身是一个指向切片的指针，所以输出中使用了 & 来表示指针
}
