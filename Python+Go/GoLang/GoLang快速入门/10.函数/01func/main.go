/*


package main

import "fmt"

func main() {
	ret := intSum(1, 2)
	fmt.Println(ret) // 3
}
func intSum(x, y int) int {
	return x + y
}

//int: 这是函数的返回类型。在这个例子中，函数返回一个整数类型的值

*/

package main

import (
	"fmt"
)

func main() {
	p, s := calc(4, 5)
	fmt.Println(p) // 和为：9
	fmt.Println(s) // 差为：-1
}
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
