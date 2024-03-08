package main

import "fmt"

func main() {
	//1、定义一个数组，工作中基本都是切片
	var a [5]int
	//声明了一个名为 a 的数组，数组的元素类型是整数 int，并且有5个元素
	//如果声明一个数组但没有指定初始值，那么数组中的元素会被赋予其类型的零值。
	//对于整数类型 int，零值就是 0
	fmt.Println(a)
	//2、向数组中制定索引添加一个元素
	a[0] = 100
	a[1] = 200
	//数组是有长度限制的，不会自增
	//a[10] = 200  // Invalid array index '10' (out of bounds for the 5-element array)
	fmt.Println(a)
	//3、让编译器自动识别我们数组长度
	var aa = [...]string{"bj", "sh", "gz"}
	//使用了 ... 语法来让编译器根据初始值的个数确定数组的长度
	//数组的元素类型是字符串 string
	//并且数组的长度会根据初始化值的个数来确定
	//由于有三个初始化值，因此数组 aa 的长度会被确定为3
	fmt.Println(aa)
}
