package main

import (
	"fmt"
	"time"
	//time 包提供了时间相关的函数和类型
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Millisecond * 100)
		//定义了一个名为 test 的函数。该函数会循环执行 10 次，
		//每次循环打印一行字符串 "test() 你好golang"，然后休眠 100 毫秒
	}
}
func main() {
	go test() //表示开启一个协程
	for i := 0; i < 2; i++ {
		fmt.Println("main() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

/*
main() 你好golang
test() 你好golang
main() 你好golang
test() 你好golang
test() 你好golang
*/
