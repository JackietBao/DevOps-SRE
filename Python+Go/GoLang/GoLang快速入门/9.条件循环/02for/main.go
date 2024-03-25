package main

import "fmt"

func main() {
	//// 1、输出1到100
	///*
	//i := 1  赋值语句
	//i <= 100  条件
	//i++   i = i + 1
	// */
	//for i := 1; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	//// 2、将i定义到外部
	//i := 0
	//for i < 10 { // 容易出现死循环，一定要有结束的方法
	//	fmt.Println(i)
	//	i++ // i = i + 1 (不是从0开始打印，而是从 i = 0 + 1)
	//	// i++在fmt上面打印 1~10
	//	// i++在fmt下面打印 0~9
	//}

	//// 3、for模拟while循环 （while true）
	//k := 1
	//for {
	//	if k <= 10 {
	//		fmt.Println(k)
	//	} else {
	//		break
	//	}
	//	k++
	//}

	// 4、for range打印索引和值
	s := "hello world"
	for index, val := range s {
		fmt.Println(index, val)
	}
}
