package main

import "fmt"

func main() {
	//1、定义切片
	//1.1 定义一个int类型的切片
	num := []int{1, 2, 3, 4}
	//1.2 定义一个sting类型的切片
	strs := []string{"LiuHui", "TianBao"}
	fmt.Println(num, strs)
	fmt.Printf("长度：%v, 容量：%v \n", len(num), cap(num))
}
