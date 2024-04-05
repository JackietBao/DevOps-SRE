package main

import (
	"fmt"
)

func reverseString(input string) string {
	// 将字符串转换为rune数组以支持修改
	characters := []rune(input)

	// 使用两个指针进行反转
	left := 0
	right := len(characters) - 1

	for left < right {
		// 交换左右指针所指的字符
		characters[left], characters[right] = characters[right], characters[left]

		// 移动指针
		left++
		right--
	}

	// 将rune数组转换为字符串并返回
	return string(characters)
}

func main() {
	input := "Hello, 世界"
	reversed := reverseString(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}
