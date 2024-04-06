package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//获取当前的时间，并将其赋值给变量now，这个时间包含了日期和时间的信息
	fmt.Printf("%T %v \n", now, now)
	strTime := now.Format("2006-01-02 15:04:05")
	//将当前时间格式化为字符串，使用Format方法按照指定的格式"2006-01-02 15:04:05"转换为字符串，
	//并将其赋值给变量strTime
	fmt.Printf("%T %v \n", strTime, strTime)
	ts := now.Unix()
	//将当前时间转换为Unix时间戳，即从1970年1月1日起到当前时间的秒数，并将其赋值给变量ts
	fmt.Printf("%T %v \n", ts, ts)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//加载指定时区的地理位置信息，此处加载了"Asia/Shanghai"时区的信息，并将其赋值给变量loc
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)

	//使用指定的时区解析字符串格式的时间，将字符串strTime按照指定格式"2006-01-02 15:04:05"解析为时间对象，
	//并且应用了上一步加载的时区信息，将解析后的时间对象赋值给变量timeObj
	fmt.Printf("%T %v \n", timeObj, timeObj)
	fmt.Println(timeObj.Unix())
	time.Sleep(time.Second * 2)
	//暂停程序执行，让程序休眠2秒
	fmt.Println("over")
}
