package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

// 用了指针，只要任意一个地方修改，其他所有地方都会改是吧
func (p *Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) Stop() {
	fmt.Println("phone 停止")
}
func main() {
	//phone1 := Phone{ // 一：实例化值类型
	//	Name: "小米手机",
	//}
	//var p1 Usb = phone1 //phone1 实现了 Usb 接口 phone1 是 Phone 类型
	//p1.Start()

	// 如果接口中使用的是指针类型作为接收者，那么只能使用指针类型的结构体去调用
	// 不能使用值类型的结构体调用
	phone2 := &Phone{ // 二：实例化指针类型
		Name: "苹果手机",
	}
	var p2 Usb = phone2 //phone2 实现了 Usb 接口 phone2 是 *Phone 类型
	p2.Start()          //苹果手机 开始工作
}
