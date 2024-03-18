package main

import "fmt"

//1.接口是一个规范
type Usber interface {
	//Usber 是一个接口类型，它规定了两个方法：start() 和 stop()
	//接口的目的是描述一种行为规范，任何实现了这两个方法的类型都可以被视为 Usber 接口的实现者
	start()
	stop()
}

//2.如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	//Phone 结构体表示手机，拥有一个字段 Name，表示手机的名字
	Name string
}

//3.手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}
func main() {
	p := Phone{
		Name: "华为手机",
	}
	var p1 Usber // golang中接口就是一个数据类型
	//声明了一个变量 p1，类型为 Usber
	p1 = p // 表示手机实现Usb接口
	p1.start()
	p1.stop()
}

/*
华为手机 启动
华为手机 关机
*/
