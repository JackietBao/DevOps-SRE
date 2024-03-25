package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine并行测试
	//test()
	// go中开启一个协程 go关键字 方法名()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Microsecond * 100)
	}

	time.Sleep(time.Second)
	fmt.Println("over")
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test", i)
		time.Sleep(time.Microsecond * 100)
	}
}

/*
main 0
test 0
main 1
test 1
test 2
main 2
test 3
main 3
test 4
test 5
main 4
main 5
test 6
main 6
test 7
main 7
test 8
main 8
main 9
test 9
over
*/
