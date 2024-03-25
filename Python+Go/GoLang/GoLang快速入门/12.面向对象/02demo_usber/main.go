package main

import "fmt"

/*
interface这个接口是golang中的一个数据类型（int、string）
和框架写的api接口没有任何关系，只是名字相同而已
*/

/*
1.函数编程：每个函数之间没有任何关系
2.结构体绑定方法：实现一类事物通用的方法，这些方法是有关系
定义一个人：eat方法、say方法 这些方法是有关系，都是定义了一个人能做的事情
*/
func main() {
	// 第四步：初始化phone结构体
	p := Phone{
		Name: "Mate 30 pro",
	}

	// 第五步：通过接口来调用方法
	var p1 Usber
	p1 = p
	Run(p1)
	//p1.Start()
	//p1.Stop()
	//
	//c := Computer{
	//	Name: "苹果",
	//}
	//var c1 Usber
	//c1 =
	//	Run(p1)
	//c.Start()

}

/*
第一句话：接口（interface）定义了一个对象的行为规范，
`只定义规范不实现`，由具体的`对象来实现规范的细节`
第二句话：一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
*/
// 第一步：定义一个接口 , 接口的名字 xxxer结尾，只是一个规范
/*
Usber  是接口的名字
interface  定义接口的关键字
*/
type Usber interface {
	// type Usber interface 就是定义了一个接口
	// Start() / Stop() 定义必须要实现这两个方法（对象实现的细节）
	Start()
	Stop()
}

/* 应为 Phone没有实现Usber要求实现的规范：Stop、Start两个方法
Cannot use 'p' (type Phone) as the type Usber
Type does not implement 'Usber' as some methods are missing: Start()
*/

// 第二句话：一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
// 第二步：定义一个结构体
type Phone struct {
	Name string
}

// 第三步：Phone这个结构体实现了接口中定义的两个方法 Start、Stop
func (p Phone) Start() {
	fmt.Println("start ...")
}

func (p Phone) Stop() {
	fmt.Println("stop ...")
}

type Computer struct {
	Name string
}

// 第三步：Phone这个结构体实现了接口中定义的两个方法 Start、Stop
func (c Computer) Start() {
	fmt.Println("start ...")
}

func (c Computer) Stop() {
	fmt.Println("stop ...")
}

func Run(u Usber) {
	u.Stop()
	u.Start()
}
