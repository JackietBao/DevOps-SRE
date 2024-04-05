package main

import "fmt"

func main() {
	// 嵌套循环，外部循环控制乘数，内部循环控制被乘数
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			// 计算乘积
			result := i * j
			// 使用Printf输出格式化的乘法表
			fmt.Printf("%d * %d = %d\t", i, j, result)
		}
		// 内部循环结束后换行
		fmt.Println()
	}
}
