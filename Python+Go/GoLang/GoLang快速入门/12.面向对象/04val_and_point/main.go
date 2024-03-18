package main

import "fmt"

type Usb interface {
	//定义了一个接口 Usb，该接口包含了 Start() 和 Stop() 两个方法
	Start()
	Stop()
}
type Phone struct {
	//结构体 Phone，包含了一个字段 Name，代表手机的名称
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) Stop() {
	fmt.Println("phone 停止")
}

//Phone 结构体上定义了 Start() 和 Stop() 两个方法，以满足 Usb 接口的实现要求
func main() {
	phone1 := Phone{ // 一：实例化值类型
		Name: "小米手机",
	}
	var p1 Usb = phone1 //phone1 实现了 Usb 接口 phone1 是 Phone 类型
	//将 phone1 赋值给 Usb 类型的变量 p1
	p1.Start()
	phone2 := &Phone{ // 二：实例化指针类型
		Name: "苹果手机",
	}
	var p2 Usb = phone2 //phone2 实现了 Usb 接口 phone2 是 *Phone 类型
	p2.Start()          //苹果手机 开始工作
}
