package main

import "fmt"

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, x := range nums {
            out <- x
        }
        close(out)
    }()
    return out
}

func mult(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for x := range in {
            out <- x * 2
        }
        close(out)
    }()
    return out
}

func main() {
	//Пример массива
	arr := []int{1, 2, 3, 4, 5, 6}

	//Пайплайн
	ch1 := gen(arr...)
	ch2 := mult(ch1)

	// Чтение из второго канала завершается корректно, т.к. gen и mult закрывают каналы
	// после передачи всех данных, а при закрытии канала происходит выход из цикла for range
	for res := range ch2 {
		fmt.Println(res)
	}
}