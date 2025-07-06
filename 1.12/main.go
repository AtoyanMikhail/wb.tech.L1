package main

import "fmt"

// Порядок недетерминирован в силу особенностей реализации мапы
func Set(arr []string) []string {
	m := map[string]struct{}{}

	for _, e := range arr {
		m[e] = struct{}{}
	}
	
	res := make([]string, 0, len(m))

	for e := range m {
		res = append(res, e)
	}

	return res
}

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(Set(s))
}