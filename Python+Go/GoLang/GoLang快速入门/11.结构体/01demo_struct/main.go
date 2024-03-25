package main

import "fmt"

/*
这一块大家学会两个事情
第一：如何定义一个结构体
第二：初始化结构体的三种方法
*/

func main() {
	// 结构体数据类型，可以存储json这种数据，绑定方法
	// 1、通过var关键字来实例化结构体
	var p1 Person
	p1.Age = 28
	p1.name = "zhangsan"
	fmt.Println(p1)
	fmt.Println(p1.name)

	//// 2、new方法实例化一个结构体(返回指针类型)
	//var p2 = new(Person)
	//p2.Age = 25
	//p2.name = "lisi"
	//fmt.Println(p2)

	//// 3、:= 键值对初始化
	//p3 := Person{
	//	name: "wangwu",
	//	Age:  35,
	//}
	//fmt.Println(p3)

}

/*
type: 关键字
Person： 结构体名字
struct：  定义结构体的关键字
*/
// golang中严格区分大小写（json序列化的时候，或者不同包访问是不能访问小写）
type Person struct {
	name string
	Age  int
}

/*
1、结构体，与函数啥区别啊。什么时候用结构体呢？
	结构体是go语言中实现类的一种方案，简单理解一个是类，一个是方法
*/
