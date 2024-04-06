package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间", now)
	m, _ := time.ParseDuration("-10m")
	//解析一个表示时间段的字符串，-10m 表示减去 10 分钟
	//将字符串解析为持续时间（time.Duration），并将其赋值给变量 m。
	//函数的第二个返回值 _ 是一个空白标识符，表示忽略该值
	m1 := now.Add(m)
	//将当前时间 now 减去 m 所表示的时间段，得到一个新的时间，并将其赋值给变量 m1
	fmt.Println("前一分钟", m1)
}
