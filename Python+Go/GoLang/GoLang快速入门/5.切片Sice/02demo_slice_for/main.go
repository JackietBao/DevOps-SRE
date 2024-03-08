package main

import "fmt"

func main() {
	//定义一个切片
	a := []string{"bj", "sh", "gz"} // 定义切片
	// 数组一定在定义的时候是知道长度 要么自己生命[5]int [...]int 系统自动识别
	b := [...]int{1, 2, 3} //定义数组[]
	//1、普通方法循环切片
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	//2、for range循环切片
	for x, y := range b {
		fmt.Println(x, y)

	}

}
