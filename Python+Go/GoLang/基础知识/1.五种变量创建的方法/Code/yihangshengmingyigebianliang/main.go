package main

//行代码指定了当前文件属于哪一个包。在Go语言中，每个Go文件都必须归属于一个包。
import "fmt"

func main() {
	var name string = "Python"
	//var关键字来声明变量，指定了变量的名称为name，类型为string，并将其初始化值设为"Python"。
	//这种声明变量的方式明确指出了变量的类型。在Go语言中，变量也可以通过类型推导来声明，但在这个例子中，使用了显式类型声明。
	fmt.Println(name)
}
