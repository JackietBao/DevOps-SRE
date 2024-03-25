package main

import (
	"fmt"
)

func main() {
	// 解决我们有大量的 if else if ...
	score := "Bxxx"
	switch score {
	case "A":
		fmt.Println("得分为A基本")
	case "B":
		fmt.Println("得分为B等级")
	case "C":
		fmt.Println("得分为C")
	default: // 当上面的所有条件全不成立，会执行default
		fmt.Println("得分为default")
	}
}
