package main

import "fmt"

func main() {
	//1、定义一个单号字符串
	s1 := "hello"
	fmt.Println(s1)
	//2、定义多行字符串
	s2 := `
	第一行
	第二行
	第三行
	`
	fmt.Println(s2)
	//3、go中str类型其实是一个只读数组，可以使用切片
	s := "hello world"
	fmt.Printf("%#v \n", s[:2]) // "he"
	//"%#v \n"是格式化字符串
	//%#v 是一个格式化动词，它表示将变量以 Go 语法的格式输出。
	//# 表示使用 Go 语法输出，v 表示将变量的值以默认格式输出
	//字符串的开头开始，取到索引为2之前（不包括索引2）的所有字符
	//s[0] = "abc" 不可以直接切片赋值
	//4、如何直接修改字符串
	/*
		注：byte类型还是rune类型都是我们go中的字符串
		byte：ASCII码(ASCII不可以表示中文)
		rune：UTF8码(每个中文3个字节，每个英文是一个字节)
	*/
	ss := "美国第一"
	/*
		第一步：需要将字符串转换成rune的byte数组：[]rune(ss)
		第二步：切片获取原有数据ss_rune[2:]
		第三步：string方法将byte数组转换成字符串：string(ss_rune[2:])
	*/
	//ss rune不是字符串，已经变成了一个rune切片
	ss_rune := []rune(ss)
	//将字符串ss转换为rune类型切片
	//rune是一个类型，用于表示Unicode代码点
	//[]rune(ss)将字符串ss转换为一个rune类型的切片，其中每个元素都是字符串中的一个Unicode字符
	fmt.Println(ss_rune)
	//string：将其他数据类型转换成string类型
	//:2切片(表示数组第二个元素往后的所有数据)
	fmt.Println(string(ss_rune[:2]))
	s_zh := "中国" + string(ss_rune[2:])
	fmt.Println(s_zh)
}
