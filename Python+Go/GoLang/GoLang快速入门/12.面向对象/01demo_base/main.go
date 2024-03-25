package main

import (
	"fmt"
)

type Cat struct {
	Name string
}

func (c Cat) Say() string { return c.Name + "：喵喵喵" }

type Dog struct {
	Name string
}

func (d Dog) Say() string { return d.Name + ": 汪汪汪" }

func main() {
	c := Cat{Name: "小白猫"} // 小白猫：喵喵喵
	fmt.Println(c.Say())
	d := Dog{"阿黄"}
	fmt.Println(d.Say()) // 阿黄: 汪汪汪
}

/*
小白猫：喵喵喵
阿黄: 汪汪汪
*/
