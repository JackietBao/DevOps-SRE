package main

import "fmt"

func main() {
	// 解决我们有大量的 if else if ...
	score := "B"
	switch score {
	//switch 语句，它会根据 score 的值执行相应的 case 分支代码块
	case "A":
		//如果 score 的值为 "A"，则会执行这个 case 下的代码块，打印 "得分为A基本
		fmt.Println("得分为A基本")
	case "B":
		fmt.Println("得分为B等级")
	case "C":
		fmt.Println("得分为C")
	default:
		fmt.Println("得分为default")
	}
}
