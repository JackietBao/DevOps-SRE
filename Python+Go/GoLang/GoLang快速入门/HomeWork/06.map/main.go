package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建一个map
	m := map[string]int{
		"banana": 2,
		"apple":  1,
		"pear":   3,
	}

	// 将map的键复制到切片中
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 遍历有序的键，并输出键值对
	for _, k := range keys {
		fmt.Println(k, ":", m[k])
	}
}
