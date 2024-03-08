package main

import (
	"fmt"
)

func main() {
	//定义 map
	userInfo := map[string]int{
		"LiuHui":   18,
		"TianBao":  70,
		"XiaoKeAi": 17,
	}
	//1、使用for range只遍历key
	for k := range userInfo {
		if val, ok := userInfo[k]; ok {
			fmt.Println(k, val)
		}
	}
	//2、遍历key和value
	for key1, value1 := range userInfo {
		fmt.Println(key1, value1)
	}
}
