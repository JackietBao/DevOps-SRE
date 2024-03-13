package main

import "fmt"

func main() {

	/*

		//1、算数运算符
		fmt.Println("1 + 1 = ", 1+1)
		fmt.Println("2 * 3 = ", 2*3)
		fmt.Println("3 / 2 = ", 3/2)
		//2、数字转换
		fmt.Println("3 / 2 = ", float64(3)/2)
		i := 1
		i++
		fmt.Println(i)

		n1 := 1
		n2 := 2
		fmt.Println(n1 == n2)
		fmt.Println(n1 > n2)
		fmt.Println(n1 < n2)
	*/

	/*

		//3、逻辑运算符
			n1 := 10
			n2 := 20

			if n1 > 15 && n2 > 15 {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}

			if n1 > 15 || n2 > 15 {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}

			n := 10
			if !(n > 20) {
				fmt.Println("yes")
			}


	*/

	n := 1
	n++
	fmt.Println(n)
	n += 2
	fmt.Println(n)
	n -= 2
	fmt.Println(n)
	n2 := 5
	n2 %= 4
	fmt.Println(n2)

}
