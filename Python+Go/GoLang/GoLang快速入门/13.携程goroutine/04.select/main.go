package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义channel
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定义channel
	strChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChan <- "aaa"
	}

	for {
		select {
		case v := <-strChan:
			// 写业务逻辑
			fmt.Println("strChan", v)
			time.Sleep(time.Microsecond * 100)
		case v := <-intChan:
			fmt.Println("intChan", v)
			time.Sleep(time.Microsecond * 100)
		default:
			fmt.Println("所有数据取出")
			return
		}
	}
}
