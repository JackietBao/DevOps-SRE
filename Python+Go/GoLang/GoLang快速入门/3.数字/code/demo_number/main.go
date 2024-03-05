package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//a := 99  // := 只能在代码块中使用，不能在全局使用

func main() {
	//1、定义数字类型,一定要注意int8能表达最大的范围
	var a int8 = 4
	//var a2 int8 = 1000 //constant 1000 overflows int8
	var b int32 = 4 //不怎么用
	var c int64 = 4 //最常用

	fmt.Println(a, b, c)
	//2、查看数据类型
	//法1：reflect查看
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	//法2：Printf
	//%T：打印类型
	//%V：打印值
	//%s：打印字符串
	//%d：打印数字
	fmt.Printf("a: %T %v \n", a, a)
	//3、布尔值0，1代表True和False
	br := true
	var b2 = false
	fmt.Println(br, b2)

	//4、打印字节占用数
	fmt.Println(unsafe.Sizeof(br))
	fmt.Println(unsafe.Sizeof(a))

	//5、如何不指定数字类型，:= 赋值，会自动变成int类型
	n := 100
	fmt.Println(reflect.TypeOf(n)) //不指定类型默认识别int
	fmt.Println(unsafe.Sizeof(n))
	fmt.Println(unsafe.Sizeof(c))
}

//fmt.Print(111)  // 除了声明变量以外的语句，都要放到代码块里
