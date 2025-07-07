package main

//Если v - очень большая строка, а нам нужна только её небольшая часть, 
// то вся большая строка останется в памяти, пока существует указатель на justString. 
// Это может привести к утечке памяти.

// Корректный вариант:
var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = string([]rune(v)[:100]) // создаём новую строку, не ссылающуюся на исходную
}

func createHugeString(size int) string {
  return string(make([]rune, size))
}

func main() {
	someFunc()
}
