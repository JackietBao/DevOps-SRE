package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	if len(s) <= 1 {
		return s
	}

	compressed := ""
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			compressed += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		}

		// Check if adding next character will make the string longer
		if len(compressed) >= len(s) {
			return s
		}
	}

	// Add the last character and its count
	compressed += string(s[len(s)-1]) + strconv.Itoa(count)

	// Check final length
	if len(compressed) >= len(s) {
		return s
	}

	return compressed
}

func main() {
	fmt.Println(compressString("aabcccccaaa")) // Output: "a2b1c5a3"
	fmt.Println(compressString("abbccd"))      // Output: "abbccd"
}
