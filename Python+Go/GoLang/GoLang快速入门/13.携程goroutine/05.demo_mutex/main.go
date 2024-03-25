package main

import (
	"fmt"
	"sync"
)

/*
当开启多个协程对同一个资源进行操作的时候，就出现了资源竞争
比如开两个协程对x同时加5000次正确结构是10000，实际结构是不相符
*/
var x int
var wg sync.WaitGroup
var lock sync.Mutex

func main() {

	fmt.Println(x)
	for i := 0; i < 4000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x) // 期待结构是10000
}

func add() {
	for i := 1; i <= 5000; i++ {
		// 当要对x操作时，先加锁，能获取锁再操作
		// 保证同一时刻只有一个协程能够对变量进行操作
		lock.Lock()
		x = x + 1
		lock.Unlock() // 操作完成释放锁
	}
	wg.Done()
}
