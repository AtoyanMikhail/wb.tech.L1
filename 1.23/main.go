package main

import (
	"fmt"
)

func Delete[S ~[]E, E any](s S, i int) S {
	_ = s[i:] // bounds check

	oldlen := len(s)
	s = append(s[:i], s[i+1:]...)
	clear(s[len(s):oldlen])
	return s
}

func main() {
	s := []string{"a", "b", "c", "d"}
	s = Delete(s, 2) 
	fmt.Println(s)   // [a b d]
}
