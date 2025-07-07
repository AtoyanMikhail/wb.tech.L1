package main

import "fmt"

func Reverse(s string) string {
	res := []rune(s)

	for i := range(len(res)/2) {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}

func main() {
	s := "главрыба"

	fmt.Println(Reverse(s))
}