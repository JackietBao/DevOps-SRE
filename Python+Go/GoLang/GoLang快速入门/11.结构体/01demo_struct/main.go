/*
package main

import "fmt"

type person struct {
	//定义了一个名为 person 的结构体类型
	name string
	city string
	age  int
}

func main() {
	var p1 person
	//声明了一个变量 p1，其类型为 person 结构体
	p1.name = "张三"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1) // p1={张三 北京 18}
	//%v 是一个占位符，表示按照变量的默认格式输出
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"张三", city:"北京", age:18}
	//%#v 是一个占位符，表示按照 Go 语法格式输出
}

*/

/*

package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

func main() {
	var p2 = new(person)
	p2.name = "张三"
	p2.age = 20
	p2.city = "北京"
	fmt.Printf("p2=%#v \n", p2) //p2=&main.person{name:"张三", city:"北京", age:20}





package main
import "fmt"
type person struct {
	name string
	city string
	age int
}
func main() {
	p4 := person{
		name: "zhangsan",
		city: "北京",
		age: 18,
	}
	// p4=main.person{name:"zhangsan", city:"北京", age:18}
	fmt.Printf("p4=%#v\n", p4)
}

}


 */



















