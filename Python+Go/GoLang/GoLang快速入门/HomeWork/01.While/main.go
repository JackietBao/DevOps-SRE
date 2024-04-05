package main

import "fmt"

func main() {
	sum := 0
	number := 2
	positive := true

	for number <= 100 {
		if positive {
			sum += number
		} else {
			sum -= number
		}

		// Flip the sign for the next iteration
		positive = !positive

		// Move to the next number
		number++
	}

	fmt.Println("Sum:", sum)
}
