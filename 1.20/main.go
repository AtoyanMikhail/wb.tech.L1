package main

import (
	"fmt"
	"slices"
)

func reverseWordsInPlace(s string) string {
	runes := []rune(s)
	n := len(runes)

	slices.Reverse(runes)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			slices.Reverse(runes[start:i])
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWordsInPlace(str))
}