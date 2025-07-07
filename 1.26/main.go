package main

import (
	"fmt"
	"strings"
)

func CharsAreUniqe(s string) bool {
	m := map[rune]struct{}{}

	s = strings.ToLower(s)

	for _, r := range s {
		if _, ok := m[r]; ok {
			return false
		}

		m[r] = struct{}{}
	}

	return	true
}

func main(){
	fmt.Println(CharsAreUniqe("abcd"))
	fmt.Println(CharsAreUniqe("abCdefAaf"))
	fmt.Println(CharsAreUniqe("aabcd"))
}
