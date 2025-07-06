package main

import "fmt"

func Swap(a, b *int) {
	*a = *a ^ *b
	*b = *b ^ *a
	*a = *a ^ *b
}

func main() {
	a, b := 9, 11
	Swap(&a, &b)

	fmt.Printf("%d %d\n", a, b)
}