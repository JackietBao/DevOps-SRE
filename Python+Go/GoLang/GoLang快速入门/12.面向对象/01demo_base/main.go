package main

import (
	"fmt"
)

type Cat struct {
	//定义了一个名为 Cat 的结构体类型
	//字段 Name，表示猫的名字
	Name string
}

func (c Cat) Say() string { return c.Name + "：喵喵喵" }

//为 Cat 结构体定义了一个方法 Say()
//它返回一个字符串，包含猫的名字和“喵喵喵”作为猫的叫声
type Dog struct {
	Name string
}

func (d Dog) Say() string { return d.Name + ": 汪汪汪" }
func main() {
	c := Cat{Name: "小白猫"} // 小白猫：喵喵喵
	//创建了一个名为 c 的 Cat 类型的变量，名字为“小白猫”
	fmt.Println(c.Say())
	d := Dog{"阿黄"}
	fmt.Println(d.Say()) // 阿黄: 汪汪汪
}
