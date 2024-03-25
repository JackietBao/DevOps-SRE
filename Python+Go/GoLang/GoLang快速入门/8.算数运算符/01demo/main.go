package main

import "fmt"

func main() {
	//// 1、算数运算符
	//fmt.Println("1 + 1 = ", 1+1)
	//fmt.Println("2 * 3 = ", 2*3)
	//fmt.Println("3 / 2 = ", 3/2) // 应该是1.5结果却是1
	//// 数字转换
	//fmt.Println("3 / 2 = ", float64(3)/2) // 1.5
	//// i++
	//i := 1
	//i++ // i++ => i = i + 1
	//fmt.Println(i)

	//// 2、关系运算符
	//n1 := 1
	//n2 := 2
	//fmt.Println(n1 == n2) // false
	//fmt.Println(n1 > n2)
	//fmt.Println(n1 < n2)

	//// 3、逻辑运算符
	//n1 := 10
	//n2 := 20
	//// and条件：所有的田间全部为true才执行
	//if n1 > 15 && n2 > 15 {
	//	fmt.Println("n1,n2都大于5")
	//}

	//// 或条件：其中有一个条件为true就执行
	//if n1 > 15 || n2 > 15 {
	//	fmt.Println("n1,n2有一个大于15")
	//}

	//// 非：当条件不成立时才执行
	//n := 10
	//if !(n > 20) {
	//	fmt.Println("n不大于20")
	//}

	// 赋值语句
	//n := 1
	//n++   n = n + 1
	//n += 2    // n = n + 2
	//n -= 2 // n = n - 2  -1
	n2 := 5
	n2 %= 4 //  n2 = 5 % 4

	fmt.Println(n2)
}
