package main

import (
	"fmt"
	"strings"
)

func main() {
	//定义一个字符串
	s := "hello world"
	s_zh := "中国"
	//1、获取字符串的长度len()
	fmt.Println(len(s))
	//"中国"这个汉字，字符串长度不是我们想象的2而是6
	//一个中文占用三个字节
	fmt.Println(len(s_zh))
	//2、字符串拼接"+"
	s_new := s_zh + "第一"
	fmt.Println(s_new) //中国第一
	//3、string.Split将字符串转换成数组
	str := "11+12+13"
	//查看某一个方法的源码： control + 点击鼠标
	arr1 := strings.Split(str, "+")
	//使用strings.Split函数将字符串str按照指定的分隔符"+"进行分割
	//分割后的部分存储到名为arr1的字符串切片中
	fmt.Printf("%T: %v \n", arr1, arr1)
	//使用fmt.Printf函数将arr1的类型和值打印到标准输出
	//%T用于格式化输出变量的类型
	//%v用于格式化输出变量的值
	//4、strings.Join将数组转换成字符串
	ss_new := strings.Join(arr1, "*")
	fmt.Printf("%T: %v \n", ss_new, ss_new)
	//5、单引只能表示一个字符
	s100 := 'a'
	//s200 := 'a' //错误写法
	//意味着s100是一个rune类型的变量，而不是字符串
	fmt.Println(s100)
	//打印字符'a'的Unicode码点
}
