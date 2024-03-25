package main

import "fmt"

func main() {
	// 2、实例化结构体（三种方法）
	p := Person{
		Name: "zhangsan",
		Age:  24,
	}
	p.setInfo()
	p.sayHi()
	fmt.Println("main", p.Name) // 理论上是lisi
	fmt.Println("main", p.Age)
}

// 1、定义一个结构体
type Person struct {
	Name string
	Age  int
}

// 3、指针接收者绑定方法（*Person传递的是一个指针）
// 指针接收者才类似于python语言的self
func (xx *Person) setInfo() {
	fmt.Println("通过结构体调用了方法")
	xx.Name = "lisi"
}

// 值接收者（Person） 只传递的是拷贝的一份数据
// 值接收者修改数据，只是修改的是拷贝的那份数据，原始数据
func (xx Person) sayHi() {
	fmt.Println("hello world")
	xx.Age = 100
	fmt.Println("sayHi", xx.Age)
}
