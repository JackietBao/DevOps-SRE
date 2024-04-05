package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, greater []int

	for _, value := range arr {
		if value < pivot {
			less = append(less, value)
		} else if value > pivot {
			greater = append(greater, value)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, pivot), greater...)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Unsorted array:", arr)
	arr = quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
