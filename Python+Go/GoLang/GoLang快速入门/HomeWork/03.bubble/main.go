package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	//获取切片的长度，用于后续迭代
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换arr[j]和arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前：", arr)
	bubbleSort(arr)
	fmt.Println("排序后：", arr)
}
