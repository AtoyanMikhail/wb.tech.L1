package main

import "fmt"

// int() отбрасывает знаки после запятой, а "/ 10 * 10" обнуляет младший разряд
func determineRange(temperature float64) int {
	return int(temperature) / 10 * 10
}

func main() {
	res := map[int][]float64{}

	//arr := []float64{-19.9, -20, -29.9, 9.9, 10.9, 11.9, 19.999}
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, e := range arr {
		rng := determineRange(e)
		res[rng] = append(res[rng], e)
	}

	fmt.Println(res)
}
