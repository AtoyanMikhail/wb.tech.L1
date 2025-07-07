package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)
	a.SetString("12345678901234567890", 10)
	b.SetString("98765432109876543210", 10)

	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", diff)

	prod := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", prod)

	quot := new(big.Int).Div(a, b)
	fmt.Println("Частное:", quot)
}
