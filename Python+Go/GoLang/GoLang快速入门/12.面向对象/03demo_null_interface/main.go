/*
package main

import "fmt"

//空接口作为函数的参数
func show(a interface{}) {
	//定义了一个函数 show
	//接受一个空接口类型 interface{} 作为参数
	//空接口类型可以表示任意类型的值
	fmt.Printf("值:%v 类型:%T\n", a, a)
}
func main() {
	show(20)         // 值:20 类型:int
	show("你好golang") // 值:你好golang 类型:string
	slice := []int{1, 2, 34, 4}
	show(slice) // 值:[1 2 34 4] 类型:[]int
}

package main

import "fmt"

func main() {
	var slice = []interface{}{"张三", 20, true, 32.2}
	//创建了一个切片 slice
	//包含了不同类型的元素，包括字符串、整数、布尔值和浮点数。这里使用了空接口类型
	//interface{}，因为它可以容纳任何类型的值
	fmt.Println(slice) // [张三 20 true 32.2]
}


package main

import "fmt"

func main() {
	// 空接口作为 map 值
	var studentInfo = make(map[string]interface{})
	//定义了一个空接口类型的 Map
	//键类型为字符串，值类型为接口类型。这意味着该 Map 可以接受任何类型的值作为其值
	studentInfo["name"] = "张三"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
	// [age:18 married:false name:张三]
}

*/

package main

import "fmt"

func main() {
	var x interface{}
	x = "Hello golnag"
	v, ok := x.(string)
	//进行类型断言，尝试将 x 中的值转换为字符串类型
	//将结果赋值给变量 v
	//ok 是一个布尔值，表示类型断言是否成功。如果成功，ok 将为 true，否则为 false
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("非字符串类型")
	}
}
