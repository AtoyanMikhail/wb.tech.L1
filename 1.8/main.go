package main

import (
	"fmt"
)

func setBit(n int64, i uint, bit uint) int64 {
	if bit == 1 {
		return n | (1 << i)  // Установить бит в 1
	} else {
		return n &^ (1 << i) // Установить бит в 0
	}
}

func main() {
	var n int64 = 5              // 0b0101=5
	var i uint = 1               // Устанавливаем 1-й бит
	fmt.Println(setBit(n, i, 0)) // 0b0100=4
	fmt.Println(setBit(n, i, 1)) // 0b0111=7
}
