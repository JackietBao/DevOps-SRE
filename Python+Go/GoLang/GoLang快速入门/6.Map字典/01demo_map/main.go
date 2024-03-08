package main

import "fmt"

func main() {
	//map是一种无序的基于 key-value的数据结构
	//1、定义一个map字典
	//[string] : 表示key的数据类型
	//第二个string: 表示value值的数据类型
	userInfo := map[string]int{ //类似Python中字典 key:str   value: int
		"LiuHui":  18,
		"TianBao": 70,
	}
	fmt.Println(userInfo)
	//2、判断一个key是否在map中(唯一依据是以ok1第二个返回值true，false来判断)
	//会返回两个数据
	//val：返回时查询可以的值
	//ok1：返回是一个bool值，true表示存在这个key，false不存在这个key
	val1, ok1 := userInfo["LiuHui"]
	//val1=18 取出key="LiuHui"的值，ok=true true表示有这个key，false表示没有key
	fmt.Println(val1, ok1) //18 true
	val2, ok2 := userInfo["GouDan"]
	fmt.Println(val2, ok2) //0 false
	//3、和if一起常用写法
	if _, ok := userInfo["LiuHui"]; ok {
		//用于检查一个名为 userInfo 的map中是否存在一个键为"LiuHui"的条目。
		//如果存在，ok 变量将被设置为 true，否则将被设置为 false。
		//_ 是一个占位符，用于忽略变量
		fmt.Println(userInfo["LiuHui"])
	} else {
		fmt.Println("map中不存在这个key")
	}
	//4、删除map中的一个数据
	fmt.Println(userInfo)
	delete(userInfo, "TianBao")
	fmt.Println(userInfo)
}
