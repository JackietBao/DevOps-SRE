package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println("输入:", strs1)
	fmt.Println("输出:", longestCommonPrefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println("输入:", strs2)
	fmt.Println("输出:", longestCommonPrefix(strs2))
}
