package main

import "fmt"

func main() {
	s := "hello 张三"
	//1、for遍历字符串常见问题
	//for i := 0; i < len(s); i++ {
	//fmt.Printf("%T: %v: %c \n", s[i], s[i], s[i])
	//%T用于格式化输出变量的类型
	//%v用于格式化输出变量的值
	//%c用于格式化输出字符
	//}
	//使用range遍历字符串
	//下划线的意思是，占位符，下面不引用不会报错
	for _, val := range s {
		//range会返回两个值，第一个值是当前迭代的索引（在这里我们使用下划线 _来忽略索引）
		//第二个值是当前迭代的元素（即当前字符），这里用变量val来接收
		//val将依次取到s中的每个字符
		//_表示忽略索引值，而val则是每次迭代的字符值
		fmt.Println(val)
		fmt.Printf("%T: %v: %c \n", val, val, val)
	}
}
