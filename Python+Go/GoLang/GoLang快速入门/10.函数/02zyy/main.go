/*

全局变量
package main

import "fmt"

//定义全局变量 num
var num int64 = 10

func main() {
	fmt.Printf("num=%d\n", num) //num=10
}
//局部变量
package main

import "fmt"

func main() {
	// 这里name是函数test的局部变量，在其他函数无法访问
	//fmt.Println(name)
	test()
}
func test() {
	name := "zhangsan"
	fmt.Println(name)
}

//for 循环语句中定义的变量

*/

package main

import "fmt"

func main() {
	testLocalVar3()
}
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量 i 只在当前 for 语句块中生效
	}
	// fmt.Println(i) //此处无法使用变量 i
}
