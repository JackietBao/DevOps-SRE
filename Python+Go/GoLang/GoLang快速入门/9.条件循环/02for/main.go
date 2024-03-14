package main

import "fmt"

func main() {
	/*
		// 1、输出1到100
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
	*/

	//2、将i定义到外部
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// 3、for模拟while循环 （while true）
	k := 1
	for {
		if k <= 10 {
			fmt.Println(k)
		} else {
			break
		}
		k++
	}
	// 4、for range打印索引和值
	s := "hello world"
	for index, val := range s {
		//for 循环，通过 range 关键字遍历字符串 s 中的每个字符
		//每次迭代中，index 表示当前字符在字符串中的索引
		//val 则表示当前字符的Unicode码点值
		fmt.Println(index, val)

	}
}
