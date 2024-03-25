package main

import "fmt"

func main() {
	// break: 跳出当前循环（循环结束）
	// continue: 结束当前这一次循环，进行下一次循环

	for i := 1; i < 10; i++ {
		if i == 5 {
			//break // 直接跳出了for循环
			continue
		}
		fmt.Println(i)
	}
}
