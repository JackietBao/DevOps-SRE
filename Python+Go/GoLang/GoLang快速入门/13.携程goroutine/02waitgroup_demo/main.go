package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 第一步：定义一个计数器
func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1 // 第三步：协程执行完毕，计数器-1
}
func test2() {
	for i := 0; i < 2; i++ {
		fmt.Println("test2() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}
func main() {
	wg.Add(1)  //协程计数器+1 第二步：开启一个协程计数器+1
	go test1() //表示开启一个协程
	wg.Add(1)  //协程计数器+1
	go test2() //表示开启一个协程
	wg.Wait()  //等待协程执行完毕... 第四步：计数器为0时推出
	fmt.Println("主线程退出...")
}

/*
test2() 你好golang- 0
test1() 你好golang- 0
.....
test1() 你好golang- 8
test1() 你好golang- 9
主线程退出...
*/
