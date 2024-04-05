package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	fmt.Println("101-200之间的素数有：")
	for i := 101; i <= 200; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			count++
		}
	}
	fmt.Printf("\n共有 %d 个素数\n", count)
}
