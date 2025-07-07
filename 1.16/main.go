package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 { return arr }
  
	left, right := 0, len(arr) - 1
  
	pivotIndex := len(arr)/2
  
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
  
	for i := range arr {
	  if arr[i] < arr[right] {
		arr[i], arr[left] = arr[left], arr[i]
		left++
	  }
	}
  
	arr[left], arr[right] = arr[right], arr[left]
  
	quicksort(arr[:left])
	quicksort(arr[left + 1:])
  
	return arr
}

func main() {
	arr := []int{5, 15, -25, 16, 12, -34}

	quicksort(arr)

	fmt.Println(arr)
}