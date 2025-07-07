package main

import (
	"fmt"
)

func reverseRunes(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func reverseWordsInPlace(s string) string {
	runes := []rune(s)
	n := len(runes)

	reverseRunes(runes, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWordsInPlace(str))
}