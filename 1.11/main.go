package main

import "fmt"

//о да, it's time to use a hashmap which allows us to get O(n) time complexity

func findIntersection(arr1, arr2 []int) []int {
	m := map[int]struct{}{}

	for _, e := range arr1 {
		m[e] = struct{}{}
	}

	res := make([]int, 0, len(arr1))

	for _, e := range arr2 {
		if _, ok := m[e]; ok {
			res = append(res, e)
		}
	}

	return res
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println(findIntersection(A, B))
}
