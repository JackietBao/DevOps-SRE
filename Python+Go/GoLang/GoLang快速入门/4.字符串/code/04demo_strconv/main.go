package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、将int转换成字符串
	num := 100
	fmt.Printf("%T %d \n", num, num)
	strNum := strconv.Itoa(num)
	//使用strconv.Itoa()函数将整数num转换为字符串，并将结果赋值给变量strNum
	//Itoa代表"Integer to ASCII"，它的作用是将整数转换为其对应的字符串表示形式
	fmt.Printf("%T %v \n", strNum, strNum) //string 100
	//%T用于格式化输出变量的类型
	//%v用于格式化输出变量的值
	//2、将字符串转为int
	//intNum接收转换后的数据，err接收转换如果报错结果
	intNum, err := strconv.Atoi(strNum)
	//使用strconv.Atoi()函数将字符串strNum转换为整数
	//如果转换成功，err将为nil
	//Atoi代表"ASCII to Integer"，它的作用是将字符串表示的整数转换为整数值
	if err != nil {
		//在错误处理方面，nil通常用于表示没有发生错误
		//err通常用于捕获函数执行过程中可能发生的错误
		fmt.Println(err)
	}
	fmt.Printf("%T: %v \n", intNum, intNum)
}
