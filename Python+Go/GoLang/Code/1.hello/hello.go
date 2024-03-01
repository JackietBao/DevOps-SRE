package main

//表示hello.go文件所在包是main，在go中每个文件都必须归属于一个包
import "fmt"

//表示：引入一个包，包名fmt，引入该包后，就可以使用fmt包的函数，例如：fmt.Println
func main() {
	//func是一个关键字，表示一个函数
	//main是函数名，是一个主函数，即我们程序的入口
	fmt.Println("hello,world!")
	//表示调用fmt包的函数Println输出"hello,world"
}
