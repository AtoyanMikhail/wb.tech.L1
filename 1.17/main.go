package main

import "fmt"

func bsearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 5, 7, 8, 12, 40}

	fmt.Println(bsearch(arr, 12))
}
