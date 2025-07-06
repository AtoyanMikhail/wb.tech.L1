package main

import "fmt"

func defineVariable(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return "idk"
	}
}

func main() {
	variables := []interface{}{
		42,
		"hello",
		true,
		false,
		make(chan interface{}),
		3.14,
	}

	for i, v := range variables {
		fmt.Printf("Элемент %d (%v): %s\n", i, v, defineVariable(v))
	}
}
