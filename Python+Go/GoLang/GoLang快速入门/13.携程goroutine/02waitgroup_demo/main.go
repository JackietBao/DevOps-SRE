package main

import (
	"fmt"
	"sync"
	"time"
)

/*
// 计数器怎么知道有一共有多少次的
goroutine开启的时候进行 wg.Add(1) 加一
goroutine结束是进行 wg.Done() 加1
wg.Wait() 会判断当前的goroutine是否为0，为0就退出

var wg sync.WaitGroup       // 第一步：定义一个计数器
wg.Add(1)               // 第二步：开启一个协程计数器+1
wg.Done()               // 第三步：协程执行完毕，计数器-1
wg.Wait()               // 第四步：计数器为0时推出
*/

var wg sync.WaitGroup // 第一步：定义一个计数器

func main() {
	wg.Add(2) // 第二步：开启一个协程计数器+1
	go test()
	go test()
	// wg.Wait()发现没有goroutine在使用而我们的wg不为0，那么就会触发异常
	wg.Wait() // 第四步：计数器为0时推出
	//time.Sleep(time.Second)
	fmt.Println("main over")
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Microsecond)
	}
	wg.Done() // 第三步：协程执行完毕，计数器-1
}
