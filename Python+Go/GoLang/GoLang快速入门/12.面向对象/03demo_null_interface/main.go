package main

import "fmt"

func main() {
	//show(1)
	//show("aa")
	//// 第二种：切片实现空接口
	////s1 := []int{1,2,3,"zhangsan"}
	//s1 := []interface{}{1, 2, "zhangsan"}
	//fmt.Println(s1)

	//// 第三种：定义map的空接口
	//d := map[string]interface{}{
	//	"name": "zhangsan",
	//	"age": 24,
	//}

	// 第四种：断言使用
	var x interface{}
	x = 1
	fmt.Printf("%T %v \n", x, x)
	//x = x + 1
	// 断言的使用：返回两个值，一个是变量本身值，第二个，ok是否是当前类型
	v, ok := x.(int)
	fmt.Println(v, ok)
	v = v + 1
	fmt.Println()
}

/*
interface{} 是一个空接口，空接口可以表示任意数据类型
*/
// 第一种：空接口作为函数的参数
func show(val interface{}) {
	fmt.Printf("%T %v \n", val, val)
}
