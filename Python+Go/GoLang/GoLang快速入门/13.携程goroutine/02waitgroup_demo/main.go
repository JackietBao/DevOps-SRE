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


package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//声明了一个全局的等待组 wg，用于等待协程完成

func main() {
	wg.Add(2)
	//将等待组的计数器设置为2，表示有两个协程需要等待
	go test()
	//启动一个新的协程来执行 test() 函数
	go test()
	//再次启动一个新的协程来执行 test() 函数
	wg.Wait()
	//等待等待组中的计数器归零，即等待两个协程执行完毕
	fmt.Println("main over")
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Microsecond)
		//每次打印后暂停 100 微秒，模拟一些工作负载
	}
	wg.Done() // 标记协程已经完成工作，将等待组中的计数器减 1
}

//运行结果会是两个协程同时执行 test() 函数，
//每个协程都会打印数字 0 到 9，并且每次打印后都会暂停 100 微秒。
//由于两个协程同时执行，它们的输出可能交错。最后，在两个协程都完成后，main 函数会打印 "main over"
