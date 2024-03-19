package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 第一步：定义一个计数器
//定义了一个全局变量 wg
//是一个 sync.WaitGroup 类型的变量
//用于等待一组 Goroutines 完成任务
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
	wg.Add(1) //协程计数器+1 第二步：开启一个协程计数器+1
	//增加了 WaitGroup 的计数器，将要开启的 Goroutines 数量加 1
	go test1() //表示开启一个协程
	//开启一个新的 Goroutine 来执行 test1() 函数
	wg.Add(1) //协程计数器+1
	//再次增加了 WaitGroup 的计数器，又要开启的 Goroutines 数量加 1。
	go test2() //表示开启一个协程
	//开启另一个 Goroutine 来执行 test2() 函数
	wg.Wait() //等待协程执行完毕... 第四步：计数器为0时推出
	//等待所有 Goroutines 完成任务，即等待 WaitGroup 的计数器归零。在这之后，程序将继续执行后面的代码
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
