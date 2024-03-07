package main

import (
	"fmt"
)

func main() {
	//1、定义一个数组
	/*
		fmt.Println(a)
		//2、普通方法通过索引遍历数组
		fmt.Println(len(a))
		for i := 0; i < len(a); i++ {
			fmt.Println(i, a[i])
	*/
	var a = [...]string{"bj", "sh", "gz"}
	//3、for range方法循环数组
	/*
		i: 是索引，只是一个变量（数组的索引）
		value： 获取的值
		:= range 是go中自己实现的方法，会返回两个值
		   - i
		   - value
	*/
	for x, y := range a {
		fmt.Println(x, y)

	}
}
